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

	t.Run("apikey has subcommands create, concurrency, enable, disable, and delete", func(t *testing.T) {
		children := cmd.ApiKeyCmd.Commands()
		names := make([]string, len(children))
		for i, c := range children {
			names[i] = c.Name()
		}
		assert.Contains(t, names, "create")
		assert.Contains(t, names, "concurrency")
		assert.Contains(t, names, "enable")
		assert.Contains(t, names, "disable")
		assert.Contains(t, names, "delete")
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
		assert.Equal(t, "create [name]", createCmd.Use)
		assert.Equal(t, "Create a new API key", createCmd.Short)
		assert.True(t, strings.Contains(createCmd.Long, "API key"))
	})

	t.Run("create command has optional --name flag", func(t *testing.T) {
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
		// --name is now optional; name can also be provided as a positional argument
		assert.False(t, strings.Contains(nameFlag.Usage, "required"))
	})

	t.Run("create command accepts positional argument", func(t *testing.T) {
		var createCmd *cobra.Command
		for _, c := range cmd.ApiKeyCmd.Commands() {
			if c.Name() == "create" {
				createCmd = c
				break
			}
		}

		assert.NotNil(t, createCmd)
		// Use field should indicate optional positional argument
		assert.Contains(t, createCmd.Use, "[name]")
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

func TestApiKeyEnableCmd(t *testing.T) {
	t.Run("enable command has correct metadata", func(t *testing.T) {
		var enableCmd *cobra.Command
		for _, c := range cmd.ApiKeyCmd.Commands() {
			if c.Name() == "enable" {
				enableCmd = c
				break
			}
		}

		assert.NotNil(t, enableCmd)
		assert.Equal(t, "enable <api-key>", enableCmd.Use)
		assert.Equal(t, "Enable an API key", enableCmd.Short)
		assert.True(t, strings.Contains(enableCmd.Long, "akm-"))
	})

	t.Run("enable command uses positional argument", func(t *testing.T) {
		var enableCmd *cobra.Command
		for _, c := range cmd.ApiKeyCmd.Commands() {
			if c.Name() == "enable" {
				enableCmd = c
				break
			}
		}

		assert.NotNil(t, enableCmd)
		assert.Contains(t, enableCmd.Use, "<api-key>")
		// No --api-key flag; the key is a positional argument
		apiKeyFlag := enableCmd.Flags().Lookup("api-key")
		assert.Nil(t, apiKeyFlag)
	})
}

func TestApiKeyDisableCmd(t *testing.T) {
	t.Run("disable command has correct metadata", func(t *testing.T) {
		var disableCmd *cobra.Command
		for _, c := range cmd.ApiKeyCmd.Commands() {
			if c.Name() == "disable" {
				disableCmd = c
				break
			}
		}

		assert.NotNil(t, disableCmd)
		assert.Equal(t, "disable <api-key>", disableCmd.Use)
		assert.Equal(t, "Disable an API key", disableCmd.Short)
		assert.True(t, strings.Contains(disableCmd.Long, "akm-"))
	})

	t.Run("disable command uses positional argument", func(t *testing.T) {
		var disableCmd *cobra.Command
		for _, c := range cmd.ApiKeyCmd.Commands() {
			if c.Name() == "disable" {
				disableCmd = c
				break
			}
		}

		assert.NotNil(t, disableCmd)
		assert.Contains(t, disableCmd.Use, "<api-key>")
		// No --api-key flag; the key is a positional argument
		apiKeyFlag := disableCmd.Flags().Lookup("api-key")
		assert.Nil(t, apiKeyFlag)
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
		assert.True(t, strings.Contains(setCmd.Long, "--api-key"))
		assert.True(t, strings.Contains(setCmd.Long, "--api-key-id"))
	})

	t.Run("set command has --api-key flag as recommended", func(t *testing.T) {
		var setCmd *cobra.Command
		for _, c := range cmd.ApiKeyCmd.Commands() {
			if c.Name() == "concurrency" {
				for _, sc := range c.Commands() {
					if sc.Name() == "set" {
						setCmd = sc
						break
					}
				}
				break
			}
		}

		assert.NotNil(t, setCmd)

		apiKeyFlag := setCmd.Flags().Lookup("api-key")
		assert.NotNil(t, apiKeyFlag)
		assert.Equal(t, "", apiKeyFlag.DefValue)
		assert.True(t, strings.Contains(apiKeyFlag.Usage, "recommended"))
	})

	t.Run("set command has --api-key-id flag with prefer --api-key usage", func(t *testing.T) {
		var setCmd *cobra.Command
		for _, c := range cmd.ApiKeyCmd.Commands() {
			if c.Name() == "concurrency" {
				for _, sc := range c.Commands() {
					if sc.Name() == "set" {
						setCmd = sc
						break
					}
				}
				break
			}
		}

		assert.NotNil(t, setCmd)

		apiKeyIdFlag := setCmd.Flags().Lookup("api-key-id")
		assert.NotNil(t, apiKeyIdFlag)
		assert.Equal(t, "", apiKeyIdFlag.DefValue)
		assert.False(t, strings.Contains(apiKeyIdFlag.Usage, "required"))
		assert.True(t, strings.Contains(apiKeyIdFlag.Usage, "--api-key"))
	})

	t.Run("set command has --concurrency as required flag", func(t *testing.T) {
		var setCmd *cobra.Command
		for _, c := range cmd.ApiKeyCmd.Commands() {
			if c.Name() == "concurrency" {
				for _, sc := range c.Commands() {
					if sc.Name() == "set" {
						setCmd = sc
						break
					}
				}
				break
			}
		}

		assert.NotNil(t, setCmd)

		concurrencyFlag := setCmd.Flags().Lookup("concurrency")
		assert.NotNil(t, concurrencyFlag)
		assert.Equal(t, "0", concurrencyFlag.DefValue)
		assert.True(t, strings.Contains(concurrencyFlag.Usage, "required"))
	})
}

func TestApikeyDeleteCmd(t *testing.T) {
	var deleteCmd *cobra.Command
	for _, c := range cmd.ApiKeyCmd.Commands() {
		if c.Name() == "delete" {
			deleteCmd = c
			break
		}
	}

	t.Run("delete command exists", func(t *testing.T) {
		assert.NotNil(t, deleteCmd)
	})

	t.Run("delete command has correct metadata", func(t *testing.T) {
		assert.NotNil(t, deleteCmd)
		assert.Equal(t, "delete <api-key>", deleteCmd.Use)
		assert.Equal(t, "Delete an API key", deleteCmd.Short)
		assert.True(t, strings.Contains(deleteCmd.Long, "akm-"))
	})

	t.Run("delete command uses positional argument", func(t *testing.T) {
		assert.NotNil(t, deleteCmd)
		assert.Contains(t, deleteCmd.Use, "<api-key>")
		// No --api-key flag; the key is a positional argument
		apiKeyFlag := deleteCmd.Flags().Lookup("api-key")
		assert.Nil(t, apiKeyFlag)
	})

	t.Run("delete command has --yes flag", func(t *testing.T) {
		assert.NotNil(t, deleteCmd)
		yesFlag := deleteCmd.Flags().Lookup("yes")
		assert.NotNil(t, yesFlag)
		assert.Equal(t, "false", yesFlag.DefValue)
		assert.Equal(t, "y", yesFlag.Shorthand)
	})
}
