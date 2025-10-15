// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package integration_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/agentbay/agentbay-cli/cmd"
	"github.com/agentbay/agentbay-cli/internal/agentbay"
	"github.com/agentbay/agentbay-cli/internal/auth"
	"github.com/agentbay/agentbay-cli/internal/client"
	"github.com/agentbay/agentbay-cli/internal/config"
)

// TestTokenRefreshIntegration tests the automatic token refresh mechanism
// Note: This test requires valid authentication tokens to work properly
func TestTokenRefreshIntegration(t *testing.T) {
	// Skip if not in integration test mode
	if os.Getenv("RUN_INTEGRATION_TESTS") == "" {
		t.Skip("Skipping integration test. Set RUN_INTEGRATION_TESTS=1 to run")
	}

	// Create temporary directory for test config
	tempDir, err := os.MkdirTemp("", "agentbay-token-refresh-test")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Set environment variable to use temp directory
	originalConfigDir := os.Getenv("AGENTBAY_CLI_CONFIG_DIR")
	os.Setenv("AGENTBAY_CLI_CONFIG_DIR", tempDir)
	defer func() {
		if originalConfigDir == "" {
			os.Unsetenv("AGENTBAY_CLI_CONFIG_DIR")
		} else {
			os.Setenv("AGENTBAY_CLI_CONFIG_DIR", originalConfigDir)
		}
	}()

	t.Run("token should be automatically refreshed when approaching expiry", func(t *testing.T) {
		// This test would require a real OAuth setup with valid tokens
		// For now, we'll test the mechanism with a simulated scenario

		cfg, err := config.GetConfig()
		require.NoError(t, err)

		// Save a token that's about to expire (4 minutes from now)
		err = cfg.SaveTokens(
			"test-access-token",
			"Bearer",
			240, // 4 minutes (will trigger refresh as it's < 5 min threshold)
			"test-refresh-token",
			"test-id-token",
		)
		require.NoError(t, err)

		// Verify token is saved
		assert.True(t, cfg.IsAuthenticated())

		// Try to refresh (will fail with test tokens but we can verify the attempt)
		err = auth.RefreshTokenIfNeeded(cfg, cmd.GetClientID())

		// We expect an error because test tokens are not valid
		// But we can verify that the refresh mechanism was triggered
		assert.Error(t, err, "expected error with test tokens")
		assert.Contains(t, err.Error(), "token refresh failed", "error should indicate refresh was attempted")
	})

	t.Run("token should not be refreshed when still valid", func(t *testing.T) {
		cfg, err := config.GetConfig()
		require.NoError(t, err)

		// Save a token that won't expire soon (10 minutes from now)
		err = cfg.SaveTokens(
			"test-access-token",
			"Bearer",
			600, // 10 minutes
			"test-refresh-token",
			"test-id-token",
		)
		require.NoError(t, err)

		// Try to refresh
		err = auth.RefreshTokenIfNeeded(cfg, cmd.GetClientID())

		// Should not attempt refresh, so no error
		assert.NoError(t, err, "should not refresh when token is still valid")
	})

	t.Run("client should handle token refresh automatically", func(t *testing.T) {
		cfg, err := config.GetConfig()
		require.NoError(t, err)

		// Clear any existing tokens first
		err = cfg.ClearTokens()
		require.NoError(t, err)

		// Save a valid token (not expiring soon)
		err = cfg.SaveTokens(
			"test-access-token",
			"Bearer",
			3600, // 1 hour
			"test-refresh-token",
			"test-id-token",
		)
		require.NoError(t, err)

		// Create a client (this should trigger token refresh check)
		agbClient := agentbay.NewClientFromConfig(cfg)
		assert.NotNil(t, agbClient)

		// Try to make an API call (will fail with test tokens, but verifies the mechanism)
		ctx := context.Background()
		request := &client.ListMcpImagesRequest{}
		_, err = agbClient.ListMcpImages(ctx, request)

		// We expect an error because test tokens are not valid
		// But the token refresh mechanism should have been invoked
		assert.Error(t, err, "expected error with test tokens")
	})
}

// TestTokenRefreshWithExpiredToken tests the behavior when token has already expired
func TestTokenRefreshWithExpiredToken(t *testing.T) {
	// Skip if not in integration test mode
	if os.Getenv("RUN_INTEGRATION_TESTS") == "" {
		t.Skip("Skipping integration test. Set RUN_INTEGRATION_TESTS=1 to run")
	}

	// Create temporary directory for test config
	tempDir, err := os.MkdirTemp("", "agentbay-expired-token-test")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Set environment variable to use temp directory
	originalConfigDir := os.Getenv("AGENTBAY_CLI_CONFIG_DIR")
	os.Setenv("AGENTBAY_CLI_CONFIG_DIR", tempDir)
	defer func() {
		if originalConfigDir == "" {
			os.Unsetenv("AGENTBAY_CLI_CONFIG_DIR")
		} else {
			os.Setenv("AGENTBAY_CLI_CONFIG_DIR", originalConfigDir)
		}
	}()

	t.Run("should attempt refresh with expired token", func(t *testing.T) {
		cfg, err := config.GetConfig()
		require.NoError(t, err)

		// Manually set an expired token
		cfg.Token = &config.Token{
			AccessToken:  "expired-access-token",
			TokenType:    "Bearer",
			ExpiresIn:    3600,
			RefreshToken: "test-refresh-token",
			IDToken:      "test-id-token",
			ExpiresAt:    time.Now().Add(-1 * time.Hour), // Expired 1 hour ago
		}
		err = cfg.Save()
		require.NoError(t, err)

		// Verify token is expired
		assert.True(t, cfg.IsTokenExpired())

		// Try to refresh
		err = auth.RefreshTokenIfNeeded(cfg, cmd.GetClientID())

		// We expect an error because test tokens are not valid
		assert.Error(t, err, "expected error with test tokens")
		assert.Contains(t, err.Error(), "token refresh failed", "error should indicate refresh was attempted")
	})
}
