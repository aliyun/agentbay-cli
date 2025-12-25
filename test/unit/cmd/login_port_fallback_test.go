// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/agentbay/agentbay-cli/cmd"
)

func TestGetRedirectURI(t *testing.T) {
	t.Run("should generate correct redirect URI for given port", func(t *testing.T) {
		port := "3001"
		expected := "http://localhost:3001/callback"
		actual := cmd.GetRedirectURI(port)
		assert.Equal(t, expected, actual)
	})

	t.Run("should generate different URIs for different ports", func(t *testing.T) {
		port1 := "3001"
		port2 := "51153"
		uri1 := cmd.GetRedirectURI(port1)
		uri2 := cmd.GetRedirectURI(port2)
		assert.NotEqual(t, uri1, uri2)
		assert.Equal(t, "http://localhost:3001/callback", uri1)
		assert.Equal(t, "http://localhost:51153/callback", uri2)
	})

	t.Run("should handle all callback ports", func(t *testing.T) {
		expectedPorts := []string{"3001", "51153", "53153", "55153", "57153"}
		for _, port := range expectedPorts {
			uri := cmd.GetRedirectURI(port)
			expected := "http://localhost:" + port + "/callback"
			assert.Equal(t, expected, uri, "Redirect URI for port %s should be correct", port)
		}
	})
}

func TestCallbackPorts(t *testing.T) {
	t.Run("should contain all expected ports", func(t *testing.T) {
		expectedPorts := []string{"3001", "51153", "53153", "55153", "57153"}
		actualPorts := cmd.CallbackPorts

		assert.Equal(t, len(expectedPorts), len(actualPorts), "CallbackPorts should have correct length")

		for i, expected := range expectedPorts {
			if i < len(actualPorts) {
				assert.Equal(t, expected, actualPorts[i], "Port at index %d should match", i)
			}
		}
	})

	t.Run("should have ports in correct order", func(t *testing.T) {
		expectedOrder := []string{"3001", "51153", "53153", "55153", "57153"}
		actualPorts := cmd.CallbackPorts

		for i, expected := range expectedOrder {
			if i < len(actualPorts) {
				assert.Equal(t, expected, actualPorts[i], "Port at index %d should be %s", i, expected)
			}
		}
	})

	t.Run("default callback port should be first port", func(t *testing.T) {
		assert.Equal(t, cmd.CallbackPorts[0], cmd.DefaultCallbackPort, "DefaultCallbackPort should be the first port in CallbackPorts")
		assert.Equal(t, "3001", cmd.DefaultCallbackPort, "DefaultCallbackPort should be 3001")
	})

	t.Run("should have at least one port", func(t *testing.T) {
		assert.Greater(t, len(cmd.CallbackPorts), 0, "CallbackPorts should contain at least one port")
	})
}
