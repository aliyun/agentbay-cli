// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package main_test

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	"github.com/agentbay/agentbay-cli/cmd"
)

func TestRootCmd(t *testing.T) {
	// Create a root command for testing
	rootCmd := &cobra.Command{
		Use:               "agentbay",
		Short:             "AgentBay CLI",
		Long:              "Command line interface for AgentBay services",
		DisableAutoGenTag: true,
		SilenceUsage:      true,
		SilenceErrors:     true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	// Add command groups
	rootCmd.AddGroup(&cobra.Group{ID: "core", Title: "Core Commands"})
	rootCmd.AddGroup(&cobra.Group{ID: "management", Title: "Management Commands"})

	// Add commands
	rootCmd.AddCommand(cmd.VersionCmd)
	rootCmd.AddCommand(cmd.LoginCmd)
	rootCmd.AddCommand(cmd.LogoutCmd)

	// Global flags
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	rootCmd.PersistentFlags().BoolP("help", "", false, "help for agentbay")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose output")
	rootCmd.Flags().BoolP("version", "", false, "Display the version of AgentBay CLI")

	t.Run("root command should display help by default", func(t *testing.T) {
		// Execute command
		err := rootCmd.RunE(rootCmd, []string{})

		// Assertions
		assert.NoError(t, err)
	})

	t.Run("root command should have correct metadata", func(t *testing.T) {
		assert.Equal(t, "agentbay", rootCmd.Use)
		assert.Equal(t, "AgentBay CLI", rootCmd.Short)
		assert.Equal(t, "Command line interface for AgentBay services", rootCmd.Long)
		assert.True(t, rootCmd.SilenceUsage)
		assert.True(t, rootCmd.SilenceErrors)
		assert.True(t, rootCmd.DisableAutoGenTag)
	})

	t.Run("root command should have required flags", func(t *testing.T) {
		// Check persistent flags
		helpFlag := rootCmd.PersistentFlags().Lookup("help")
		assert.NotNil(t, helpFlag)
		assert.Equal(t, "bool", helpFlag.Value.Type())

		verboseFlag := rootCmd.PersistentFlags().Lookup("verbose")
		assert.NotNil(t, verboseFlag)
		assert.Equal(t, "bool", verboseFlag.Value.Type())

		// Check local flags
		versionFlag := rootCmd.Flags().Lookup("version")
		assert.NotNil(t, versionFlag)
		assert.Equal(t, "bool", versionFlag.Value.Type())
	})

	t.Run("root command should have command groups", func(t *testing.T) {
		groups := rootCmd.Groups()
		assert.Len(t, groups, 2)

		// Check for core group
		var coreGroup *cobra.Group
		var managementGroup *cobra.Group
		for _, group := range groups {
			if group.ID == "core" {
				coreGroup = group
			}
			if group.ID == "management" {
				managementGroup = group
			}
		}

		assert.NotNil(t, coreGroup)
		assert.Equal(t, "Core Commands", coreGroup.Title)
		assert.NotNil(t, managementGroup)
		assert.Equal(t, "Management Commands", managementGroup.Title)
	})

	t.Run("root command should have required subcommands", func(t *testing.T) {
		commands := rootCmd.Commands()
		assert.Len(t, commands, 3)

		// Check that we have the expected commands
		commandNames := make([]string, len(commands))
		for i, cmd := range commands {
			commandNames[i] = cmd.Use
		}

		assert.Contains(t, commandNames, "version")
		assert.Contains(t, commandNames, "login")
		assert.Contains(t, commandNames, "logout")
	})
}
