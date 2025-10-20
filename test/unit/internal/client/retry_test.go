// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"errors"
	"net"
	"net/http"
	"testing"

	"github.com/agentbay/agentbay-cli/internal/client"
	"github.com/stretchr/testify/assert"
)

// TestIsRetryableError tests the IsRetryableError function
func TestIsRetryableError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "nil error is not retryable",
			err:      nil,
			expected: false,
		},
		{
			name:     "bad file descriptor error is retryable",
			err:      errors.New("bad file descriptor"),
			expected: true,
		},
		{
			name:     "connection refused is retryable",
			err:      errors.New("connection refused"),
			expected: true,
		},
		{
			name:     "connection reset by peer is retryable",
			err:      errors.New("connection reset by peer"),
			expected: true,
		},
		{
			name:     "timeout error is retryable",
			err:      errors.New("timeout"),
			expected: true,
		},
		{
			name:     "i/o timeout is retryable",
			err:      errors.New("i/o timeout"),
			expected: true,
		},
		{
			name:     "network is unreachable is retryable",
			err:      errors.New("network is unreachable"),
			expected: true,
		},
		{
			name:     "broken pipe is retryable",
			err:      errors.New("broken pipe"),
			expected: true,
		},
		{
			name:     "parse error is not retryable",
			err:      errors.New("parse error"),
			expected: false,
		},
		{
			name:     "malformed error is not retryable",
			err:      errors.New("malformed URL"),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := client.IsRetryableError(tt.err)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// TestIsRetryableHTTPStatus tests the IsRetryableHTTPStatus function
func TestIsRetryableHTTPStatus(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		expected   bool
	}{
		{
			name:       "200 OK is not retryable",
			statusCode: http.StatusOK,
			expected:   false,
		},
		{
			name:       "400 Bad Request is not retryable",
			statusCode: http.StatusBadRequest,
			expected:   false,
		},
		{
			name:       "401 Unauthorized is not retryable",
			statusCode: http.StatusUnauthorized,
			expected:   false,
		},
		{
			name:       "404 Not Found is not retryable",
			statusCode: http.StatusNotFound,
			expected:   false,
		},
		{
			name:       "408 Request Timeout is retryable",
			statusCode: http.StatusRequestTimeout,
			expected:   true,
		},
		{
			name:       "429 Too Many Requests is retryable",
			statusCode: http.StatusTooManyRequests,
			expected:   true,
		},
		{
			name:       "500 Internal Server Error is retryable",
			statusCode: http.StatusInternalServerError,
			expected:   true,
		},
		{
			name:       "502 Bad Gateway is retryable",
			statusCode: http.StatusBadGateway,
			expected:   true,
		},
		{
			name:       "503 Service Unavailable is retryable",
			statusCode: http.StatusServiceUnavailable,
			expected:   true,
		},
		{
			name:       "504 Gateway Timeout is retryable",
			statusCode: http.StatusGatewayTimeout,
			expected:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := client.IsRetryableHTTPStatus(tt.statusCode)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// TestDefaultRetryConfig tests the DefaultRetryConfig function
func TestDefaultRetryConfig(t *testing.T) {
	config := client.DefaultRetryConfig()

	assert.NotNil(t, config)
	assert.Equal(t, 3, config.MaxRetries)
	assert.Equal(t, 1, int(config.InitialDelay.Seconds()))
	assert.Equal(t, 10, int(config.MaxDelay.Seconds()))
	assert.Equal(t, 2.0, config.BackoffFactor)
}

// TestNetErrorRetryable tests that net.Error with Temporary() or Timeout() is retryable
func TestNetErrorRetryable(t *testing.T) {
	// Create a mock net.Error
	testErr := &net.OpError{
		Op:  "read",
		Net: "tcp",
		Err: errors.New("i/o timeout"),
	}

	result := client.IsRetryableError(testErr)
	// This should be retryable because it contains "i/o timeout"
	assert.True(t, result)
}
