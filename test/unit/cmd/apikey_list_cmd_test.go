// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd_test

import (
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	"github.com/agentbay/agentbay-cli/cmd"
)

func TestApiKeyListCmd(t *testing.T) {
	t.Run("list command has correct metadata", func(t *testing.T) {
		var listCmd *cobra.Command
		for _, c := range cmd.ApiKeyCmd.Commands() {
			if c.Name() == "list" {
				listCmd = c
				break
			}
		}

		assert.NotNil(t, listCmd)
		assert.Equal(t, "list", listCmd.Use)
		assert.Equal(t, "List API keys", listCmd.Short)
		assert.True(t, strings.Contains(listCmd.Long, "api-key"))
		assert.True(t, strings.Contains(listCmd.Long, "max-results"))
		assert.True(t, strings.Contains(listCmd.Long, "next-token"))
	})

	t.Run("list command has --max-results flag with default 10", func(t *testing.T) {
		var listCmd *cobra.Command
		for _, c := range cmd.ApiKeyCmd.Commands() {
			if c.Name() == "list" {
				listCmd = c
				break
			}
		}

		assert.NotNil(t, listCmd)
		maxResultsFlag := listCmd.Flags().Lookup("max-results")
		assert.NotNil(t, maxResultsFlag)
		assert.Equal(t, "10", maxResultsFlag.DefValue)
	})

	t.Run("list command has --api-key flag", func(t *testing.T) {
		var listCmd *cobra.Command
		for _, c := range cmd.ApiKeyCmd.Commands() {
			if c.Name() == "list" {
				listCmd = c
				break
			}
		}

		assert.NotNil(t, listCmd)
		apiKeyFlag := listCmd.Flags().Lookup("api-key")
		assert.NotNil(t, apiKeyFlag)
		assert.Equal(t, "", apiKeyFlag.DefValue)
		assert.True(t, strings.Contains(apiKeyFlag.Usage, "akm-"))
	})

	t.Run("list command has --api-key-id flag", func(t *testing.T) {
		var listCmd *cobra.Command
		for _, c := range cmd.ApiKeyCmd.Commands() {
			if c.Name() == "list" {
				listCmd = c
				break
			}
		}

		assert.NotNil(t, listCmd)
		apiKeyIdFlag := listCmd.Flags().Lookup("api-key-id")
		assert.NotNil(t, apiKeyIdFlag)
		assert.Equal(t, "", apiKeyIdFlag.DefValue)
		assert.True(t, strings.Contains(apiKeyIdFlag.Usage, "--api-key"))
	})

	t.Run("list command has --next-token flag", func(t *testing.T) {
		var listCmd *cobra.Command
		for _, c := range cmd.ApiKeyCmd.Commands() {
			if c.Name() == "list" {
				listCmd = c
				break
			}
		}

		assert.NotNil(t, listCmd)
		nextTokenFlag := listCmd.Flags().Lookup("next-token")
		assert.NotNil(t, nextTokenFlag)
		assert.Equal(t, "", nextTokenFlag.DefValue)
	})

	t.Run("list command is registered under apikey", func(t *testing.T) {
		children := cmd.ApiKeyCmd.Commands()
		names := make([]string, len(children))
		for i, c := range children {
			names[i] = c.Name()
		}
		assert.Contains(t, names, "list")
	})
}
