// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"net"
	"net/http"
	"strings"
	"syscall"
	"time"
)

// RetryConfig defines retry behavior
type RetryConfig struct {
	MaxRetries    int
	InitialDelay  time.Duration
	MaxDelay      time.Duration
	BackoffFactor float64
}

// DefaultRetryConfig returns a sensible default retry configuration
func DefaultRetryConfig() *RetryConfig {
	return &RetryConfig{
		MaxRetries:    3,
		InitialDelay:  1 * time.Second,
		MaxDelay:      10 * time.Second,
		BackoffFactor: 2.0,
	}
}

// IsRetryableError determines if an error should trigger a retry
func IsRetryableError(err error) bool {
	if err == nil {
		return false
	}

	// Network connection errors that are typically transient
	errorStr := err.Error()

	// First check for non-retryable errors
	nonRetryableErrors := []string{
		"invalid uri for request",
		"parse",
		"malformed",
		"unsupported protocol",
	}

	for _, nonRetryableErr := range nonRetryableErrors {
		if strings.Contains(strings.ToLower(errorStr), nonRetryableErr) {
			return false
		}
	}

	// Common transient network errors
	retryableErrors := []string{
		"bad file descriptor",
		"connection refused",
		"connection reset by peer",
		"no such host",
		"network is unreachable",
		"timeout",
		"deadline exceeded",
		"temporary failure",
		"i/o timeout",
		"broken pipe",
	}

	for _, retryableErr := range retryableErrors {
		if strings.Contains(strings.ToLower(errorStr), retryableErr) {
			return true
		}
	}

	// Check for specific error types
	if netErr, ok := err.(net.Error); ok {
		return netErr.Temporary() || netErr.Timeout()
	}

	// Check for syscall errors
	if opErr, ok := err.(*net.OpError); ok {
		if syscallErr, ok := opErr.Err.(*syscall.Errno); ok {
			switch *syscallErr {
			case syscall.ECONNREFUSED, syscall.ECONNRESET, syscall.ETIMEDOUT:
				return true
			}
		}
	}

	return false
}

// IsRetryableHTTPStatus determines if an HTTP status code should trigger a retry
func IsRetryableHTTPStatus(statusCode int) bool {
	// Retry on server errors and some client errors
	switch statusCode {
	case http.StatusRequestTimeout, // 408
		http.StatusTooManyRequests,     // 429
		http.StatusInternalServerError, // 500
		http.StatusBadGateway,          // 502
		http.StatusServiceUnavailable,  // 503
		http.StatusGatewayTimeout:      // 504
		return true
	}
	return false
}

// IsTransientGatewayError returns true for network-layer failures and common
// gateway / overload signals often wrapped by HTTP SDKs. Used for CLI retries
// on API calls where the operation is safe to attempt again.
func IsTransientGatewayError(err error) bool {
	if err == nil {
		return false
	}
	if IsRetryableError(err) {
		return true
	}
	msg := strings.ToLower(err.Error())
	keywords := []string{
		"timeout", "timed out", "connection reset", "connection refused",
		"eof", "broken pipe", "tls handshake", "server closed",
		"bad gateway", "service unavailable", "gateway timeout",
		"too many requests", "throttl",
		" status code 408", " status code 429",
		" status code 500", " status code 502", " status code 503", " status code 504",
	}
	for _, k := range keywords {
		if strings.Contains(msg, k) {
			return true
		}
	}
	return false
}
