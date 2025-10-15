// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/agentbay/agentbay-cli/internal/agentbay"
	"github.com/agentbay/agentbay-cli/internal/config"
)

func TestClient(t *testing.T) {
	t.Run("NewClient should create client with correct configuration", func(t *testing.T) {
		cfg := &config.Config{}
		apiConfig := &config.APIConfig{
			Endpoint:  "api.example.com",
			TimeoutMs: 30000,
		}

		c := agentbay.NewClient(apiConfig, cfg)
		assert.NotNil(t, c)
	})

	t.Run("NewClientFromConfig should create client with default configuration", func(t *testing.T) {
		cfg := &config.Config{}

		c := agentbay.NewClientFromConfig(cfg)
		assert.NotNil(t, c)
	})
}
