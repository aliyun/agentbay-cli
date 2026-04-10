// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"errors"
	"testing"
	"time"

	"github.com/agentbay/agentbay-cli/internal/client"
	"github.com/stretchr/testify/require"
)

func TestWithTransientRetry_recoverAfterTransient(t *testing.T) {
	cfg := &client.RetryConfig{
		MaxRetries:    3,
		InitialDelay:  1 * time.Millisecond,
		MaxDelay:      5 * time.Millisecond,
		BackoffFactor: 2.0,
	}
	var n int
	err := withTransientRetry(cfg, "test", func() error {
		n++
		if n < 2 {
			return errors.New("connection reset by peer")
		}
		return nil
	})
	require.NoError(t, err)
	require.Equal(t, 2, n)
}

func TestWithTransientRetry_nonRetryableStopsImmediately(t *testing.T) {
	var n int
	err := withTransientRetry(client.DefaultRetryConfig(), "test", func() error {
		n++
		return errors.New("InvalidParameter: not transient")
	})
	require.Error(t, err)
	require.Equal(t, 1, n)
}
