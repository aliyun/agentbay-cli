// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/agentbay/agentbay-cli/internal/client"
)

// withTransientRetry runs fn until success or retries are exhausted. Uses
// exponential backoff from cfg between attempts after the first failure.
// The retry loop is context-aware: once ctx is cancelled or deadlined, the
// loop exits immediately without further sleeps.
func withTransientRetry(ctx context.Context, cfg *client.RetryConfig, opName string, fn func() error) error {
	if cfg == nil {
		cfg = client.DefaultRetryConfig()
	}
	delay := cfg.InitialDelay
	var lastErr error
	for attempt := 0; attempt <= cfg.MaxRetries; attempt++ {
		// Short-circuit: if the context is already done, return immediately
		// without further sleeps or attempts.
		if ctx.Err() != nil {
			return ctx.Err()
		}
		if attempt > 0 {
			if log.GetLevel() >= log.DebugLevel {
				log.Debugf("[DEBUG] %s: retry %d/%d after %v: %v", opName, attempt+1, cfg.MaxRetries+1, delay, lastErr)
			}
			select {
			case <-ctx.Done():
				return lastErr
			case <-time.After(delay):
			}
			nd := time.Duration(float64(delay) * cfg.BackoffFactor)
			if nd > cfg.MaxDelay {
				nd = cfg.MaxDelay
			}
			delay = nd
		}
		lastErr = fn()
		if lastErr == nil {
			return nil
		}
		if attempt == cfg.MaxRetries || !client.IsTransientGatewayError(lastErr) {
			return lastErr
		}
	}
	return lastErr
}
