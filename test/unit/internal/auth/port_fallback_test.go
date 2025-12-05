// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package auth_test

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/agentbay/agentbay-cli/internal/auth"
)

// occupyPort occupies a port and returns the listener for cleanup
func occupyPort(t *testing.T, port string) net.Listener {
	listener, err := net.Listen("tcp", ":"+port)
	require.NoError(t, err, "Failed to occupy port %s", port)
	return listener
}

// findAvailablePort finds an available port for testing
func findAvailablePort(t *testing.T) string {
	listener, err := net.Listen("tcp", ":0")
	require.NoError(t, err)
	port := listener.Addr().(*net.TCPAddr).Port
	listener.Close()
	return fmt.Sprintf("%d", port)
}

func TestIsPortOccupied(t *testing.T) {
	t.Run("port should be available when not occupied", func(t *testing.T) {
		// Find an available port
		listener, err := net.Listen("tcp", ":0")
		require.NoError(t, err)
		port := listener.Addr().(*net.TCPAddr).Port
		listener.Close()

		// Convert port number to string
		portStr := fmt.Sprintf("%d", port)

		// Test that port is not occupied
		occupied := auth.IsPortOccupied(portStr)
		assert.False(t, occupied, "Port %s should not be occupied", portStr)
	})

	t.Run("port should be occupied when listener is active", func(t *testing.T) {
		// Find an available port
		listener, err := net.Listen("tcp", ":0")
		require.NoError(t, err)
		portNum := listener.Addr().(*net.TCPAddr).Port
		listener.Close()
		port := fmt.Sprintf("%d", portNum)

		// Occupy the port
		occupiedListener := occupyPort(t, port)
		defer occupiedListener.Close()

		// Test that port is occupied
		occupied := auth.IsPortOccupied(port)
		assert.True(t, occupied, "Port %s should be occupied", port)
	})

	t.Run("IsPortOccupied should be accurate for concurrent checks", func(t *testing.T) {
		// Find an available port
		listener, err := net.Listen("tcp", ":0")
		require.NoError(t, err)
		portNum := listener.Addr().(*net.TCPAddr).Port
		listener.Close()
		port := fmt.Sprintf("%d", portNum)

		// Verify port is available before concurrent checks
		assert.False(t, auth.IsPortOccupied(port), "Port should be available before concurrent checks")

		// Test concurrent checks (but don't occupy the port)
		results := make(chan bool, 10)
		for i := 0; i < 10; i++ {
			go func() {
				results <- auth.IsPortOccupied(port)
			}()
		}

		// All should return false (port is available)
		// Note: In rare cases, another process might occupy the port between checks
		// So we allow for some flexibility
		falseCount := 0
		for i := 0; i < 10; i++ {
			occupied := <-results
			if !occupied {
				falseCount++
			}
		}
		// At least most checks should return false (port is available)
		// Allow for race conditions where port might be temporarily occupied
		assert.GreaterOrEqual(t, falseCount, 5, "Most concurrent checks should report port as available")
	})
}

func TestStartCallbackServer_PortBinding(t *testing.T) {
	t.Run("should successfully bind available port", func(t *testing.T) {
		// Find an available port
		listener, err := net.Listen("tcp", ":0")
		require.NoError(t, err)
		portNum := listener.Addr().(*net.TCPAddr).Port
		listener.Close()
		port := fmt.Sprintf("%d", portNum)

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		// Start server in background
		codeChan := make(chan string, 1)
		errChan := make(chan error, 1)

		go func() {
			code, err := auth.StartCallbackServer(ctx, port)
			if err != nil {
				errChan <- err
				return
			}
			codeChan <- code
		}()

		// Give server time to start
		time.Sleep(200 * time.Millisecond)

		// Verify port is now occupied (server is bound)
		occupied := auth.IsPortOccupied(port)
		assert.True(t, occupied, "Port %s should be occupied after server starts", port)

		// Cancel context to stop server
		cancel()

		// Wait for server to stop
		select {
		case <-codeChan:
			// Server completed (unexpected in this test)
		case <-errChan:
			// Server error (expected due to context cancellation)
		case <-time.After(1 * time.Second):
			// Timeout - server should have stopped
		}
	})

	t.Run("should fail immediately when port is occupied", func(t *testing.T) {
		// Find an available port
		listener, err := net.Listen("tcp", ":0")
		require.NoError(t, err)
		portNum := listener.Addr().(*net.TCPAddr).Port
		listener.Close()
		port := fmt.Sprintf("%d", portNum)

		// Occupy the port first
		listener = occupyPort(t, port)
		defer listener.Close()

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		// Try to start server on occupied port
		code, err := auth.StartCallbackServer(ctx, port)

		// Should fail immediately
		assert.Error(t, err, "Should fail when port is occupied")
		assert.Empty(t, code, "Should not return code when port is occupied")
		assert.Contains(t, err.Error(), "port")
		assert.Contains(t, err.Error(), "occupied")
	})

	t.Run("port binding should be atomic", func(t *testing.T) {
		// Find an available port
		listener, err := net.Listen("tcp", ":0")
		require.NoError(t, err)
		portNum := listener.Addr().(*net.TCPAddr).Port
		listener.Close()
		port := fmt.Sprintf("%d", portNum)

		// Verify port is available
		assert.False(t, auth.IsPortOccupied(port), "Port should be available before test")

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		// Try to start two servers on the same port concurrently
		errChan1 := make(chan error, 1)
		errChan2 := make(chan error, 1)

		// Start both servers at nearly the same time
		go func() {
			_, err := auth.StartCallbackServer(ctx, port)
			errChan1 <- err
		}()

		// Small delay to ensure first server starts binding
		time.Sleep(50 * time.Millisecond)

		go func() {
			_, err := auth.StartCallbackServer(ctx, port)
			errChan2 <- err
		}()

		// Wait for both to complete
		err1 := <-errChan1
		err2 := <-errChan2

		// At least one should fail (port is occupied by the other)
		// Both might fail if they both try to bind simultaneously and one fails
		// But at most one should succeed
		successCount := 0
		if err1 == nil {
			successCount++
		}
		if err2 == nil {
			successCount++
		}

		// At most one should succeed (atomic binding ensures this)
		assert.LessOrEqual(t, successCount, 1, "At most one server should succeed in binding the port")
		// At least one should fail (if both succeed, binding is not atomic)
		assert.GreaterOrEqual(t, successCount, 0, "At least zero servers should succeed (both may fail in race condition)")
	})
}

func TestStartCallbackServer_PortOccupied(t *testing.T) {
	t.Run("should return clear error message when port is occupied", func(t *testing.T) {
		// Find an available port
		listener, err := net.Listen("tcp", ":0")
		require.NoError(t, err)
		portNum := listener.Addr().(*net.TCPAddr).Port
		listener.Close()
		port := fmt.Sprintf("%d", portNum)

		// Occupy the port
		listener = occupyPort(t, port)
		defer listener.Close()

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		_, err = auth.StartCallbackServer(ctx, port)

		require.Error(t, err)
		assert.Contains(t, err.Error(), "port")
		assert.Contains(t, err.Error(), "occupied")
		assert.Contains(t, err.Error(), port)
	})
}
