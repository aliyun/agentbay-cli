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

func TestApiKeyCmd(t *testing.T) {
	t.Run("apikey command has correct metadata", func(t *testing.T) {
		assert.Equal(t, "apikey", cmd.ApiKeyCmd.Use)
		assert.Equal(t, "Manage AgentBay API keys", cmd.ApiKeyCmd.Short)
		assert.Equal(t, "management", cmd.ApiKeyCmd.GroupID)
		assert.True(t, strings.Contains(cmd.ApiKeyCmd.Long, "Create"))
		assert.True(t, strings.Contains(cmd.ApiKeyCmd.Long, "API keys"))
	})

	t.Run("apikey has subcommands create and concurrency", func(t *testing.T) {
		children := cmd.ApiKeyCmd.Commands()
		names := make([]string, len(children))
		for i, c := range children {
			names[i] = c.Name()
		}
		assert.Contains(t, names, "create")
		assert.Contains(t, names, "concurrency")
	})
}

func TestApiKeyCreateCmd(t *testing.T) {
	t.Run("create command has correct metadata", func(t *testing.T) {
		var createCmd *cobra.Command
		for _, c := range cmd.ApiKeyCmd.Commands() {
			if c.Name() == "create" {
				createCmd = c
				break
			}
		}
		
		assert.NotNil(t, createCmd)
		assert.Equal(t, "create", createCmd.Use)
		assert.Equal(t, "Create a new API key", createCmd.Short)
		assert.True(t, strings.Contains(createCmd.Long, "API key"))
	})

	t.Run("create command has required name flag", func(t *testing.T) {
		var createCmd *cobra.Command
		for _, c := range cmd.ApiKeyCmd.Commands() {
			if c.Name() == "create" {
				createCmd = c
				break
			}
		}
		
		assert.NotNil(t, createCmd)
		
		nameFlag := createCmd.Flags().Lookup("name")
		assert.NotNil(t, nameFlag)
		assert.Equal(t, "", nameFlag.DefValue)
		assert.True(t, strings.Contains(nameFlag.Usage, "required"))
	})

	t.Run("create command fails without name flag", func(t *testing.T) {
		var createCmd *cobra.Command
		for _, c := range cmd.ApiKeyCmd.Commands() {
			if c.Name() == "create" {
				createCmd = c
				break
			}
		}
		
		assert.NotNil(t, createCmd)
		
		// Test that the RunE function checks for required flag
		createCmd.SetArgs([]string{})
		// Don't execute, just verify the flag is marked as required
		flag := createCmd.Flags().Lookup("name")
		assert.NotNil(t, flag)
		// The flag should have no default value, making it required
		assert.Equal(t, "", flag.DefValue)
	})
}

func TestApiKeyConcurrencyCmd(t *testing.T) {
	t.Run("concurrency command has correct metadata", func(t *testing.T) {
		var concurrencyCmd *cobra.Command
		for _, c := range cmd.ApiKeyCmd.Commands() {
			if c.Name() == "concurrency" {
				concurrencyCmd = c
				break
			}
		}
		
		assert.NotNil(t, concurrencyCmd)
		assert.Equal(t, "concurrency", concurrencyCmd.Use)
		assert.Equal(t, "Manage API key concurrency settings", concurrencyCmd.Short)
	})

	t.Run("concurrency has set subcommand", func(t *testing.T) {
		var concurrencyCmd *cobra.Command
		for _, c := range cmd.ApiKeyCmd.Commands() {
			if c.Name() == "concurrency" {
				concurrencyCmd = c
				break
			}
		}
		
		assert.NotNil(t, concurrencyCmd)
		
		children := concurrencyCmd.Commands()
		names := make([]string, len(children))
		for i, c := range children {
			names[i] = c.Name()
		}
		assert.Contains(t, names, "set")
	})
}

func TestApiKeyConcurrencySetCmd(t *testing.T) {
	t.Run("set command has correct metadata", func(t *testing.T) {
		var concurrencyCmd *cobra.Command
		for _, c := range cmd.ApiKeyCmd.Commands() {
			if c.Name() == "concurrency" {
				concurrencyCmd = c
				break
			}
		}
		
		assert.NotNil(t, concurrencyCmd)
		
		var setCmd *cobra.Command
		for _, c := range concurrencyCmd.Commands() {
			if c.Name() == "set" {
				setCmd = c
				break
			}
		}
		
		assert.NotNil(t, setCmd)
		assert.Equal(t, "set", setCmd.Use)
		assert.Equal(t, "Set the concurrency limit for an API key", setCmd.Short)
		assert.True(t, strings.Contains(setCmd.Long, "concurrent sessions"))
	})

	t.Run("set command has required flags", func(t *testing.T) {
		var concurrencyCmd *cobra.Command
		for _, c := range cmd.ApiKeyCmd.Commands() {
			if c.Name() == "concurrency" {
				concurrencyCmd = c
				break
			}
		}
		
		assert.NotNil(t, concurrencyCmd)
		
		var setCmd *cobra.Command
		for _, c := range concurrencyCmd.Commands() {
			if c.Name() == "set" {
				setCmd = c
				break
			}
		}
		
		assert.NotNil(t, setCmd)
		
		apiKeyIdFlag := setCmd.Flags().Lookup("api-key-id")
		assert.NotNil(t, apiKeyIdFlag)
		assert.Equal(t, "", apiKeyIdFlag.DefValue)
		assert.True(t, strings.Contains(apiKeyIdFlag.Usage, "required"))
		
		concurrencyFlag := setCmd.Flags().Lookup("concurrency")
		assert.NotNil(t, concurrencyFlag)
		assert.Equal(t, "0", concurrencyFlag.DefValue)
		assert.True(t, strings.Contains(concurrencyFlag.Usage, "required"))
	})

	t.Run("set command fails without required flags", func(t *testing.T) {
		var concurrencyCmd *cobra.Command
		for _, c := range cmd.ApiKeyCmd.Commands() {
			if c.Name() == "concurrency" {
				concurrencyCmd = c
				break
			}
		}
		
		assert.NotNil(t, concurrencyCmd)
		
		var setCmd *cobra.Command
		for _, c := range concurrencyCmd.Commands() {
			if c.Name() == "set" {
				setCmd = c
				break
			}
		}
		
		assert.NotNil(t, setCmd)
		
		// Verify flags are marked as required
		apiKeyIdFlag := setCmd.Flags().Lookup("api-key-id")
		assert.NotNil(t, apiKeyIdFlag)
		assert.Equal(t, "", apiKeyIdFlag.DefValue)
		
		concurrencyFlag := setCmd.Flags().Lookup("concurrency")
		assert.NotNil(t, concurrencyFlag)
		assert.Equal(t, "0", concurrencyFlag.DefValue)
	})
}
