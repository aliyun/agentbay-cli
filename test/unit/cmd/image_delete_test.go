// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd_test

import (
	"os"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/agentbay/agentbay-cli/cmd"
	"github.com/agentbay/agentbay-cli/internal/config"
)

func findImageDeleteSubcommand(t *testing.T) *cobra.Command {
	t.Helper()
	for _, sub := range cmd.ImageCmd.Commands() {
		if sub.Name() == "delete" {
			return sub
		}
	}
	return nil
}

func TestImageDeleteCommand(t *testing.T) {
	t.Run("delete subcommand metadata", func(t *testing.T) {
		deleteCmd := findImageDeleteSubcommand(t)
		require.NotNil(t, deleteCmd, "image delete subcommand not found")
		assert.Equal(t, "delete", deleteCmd.Name())
		assert.Contains(t, strings.ToLower(deleteCmd.Short), "delete")
		assert.Contains(t, strings.ToLower(deleteCmd.Long), "permanently")
		assert.NotNil(t, deleteCmd.RunE)
	})

	t.Run("delete requires exactly one argument", func(t *testing.T) {
		deleteCmd := findImageDeleteSubcommand(t)
		require.NotNil(t, deleteCmd)
		assert.NotNil(t, deleteCmd.Args)
		// Should accept exactly one argument
		assert.NoError(t, deleteCmd.Args(deleteCmd, []string{"imgc-test"}))
		// Should reject zero arguments
		assert.Error(t, deleteCmd.Args(deleteCmd, []string{}))
		// Should reject two arguments
		assert.Error(t, deleteCmd.Args(deleteCmd, []string{"a", "b"}))
	})

	t.Run("delete has --yes flag", func(t *testing.T) {
		deleteCmd := findImageDeleteSubcommand(t)
		require.NotNil(t, deleteCmd)

		yesFlag := deleteCmd.Flags().Lookup("yes")
		require.NotNil(t, yesFlag, "--yes flag not found")
		assert.Equal(t, "false", yesFlag.DefValue)
		assert.Equal(t, "y", yesFlag.Shorthand)
	})

	t.Run("delete fails when not authenticated", func(t *testing.T) {
		tempDir, err := os.MkdirTemp("", "agentbay-image-delete-test")
		require.NoError(t, err)
		defer os.RemoveAll(tempDir)

		origDir := os.Getenv("AGENTBAY_CLI_CONFIG_DIR")
		os.Setenv("AGENTBAY_CLI_CONFIG_DIR", tempDir)
		defer func() {
			if origDir == "" {
				os.Unsetenv("AGENTBAY_CLI_CONFIG_DIR")
			} else {
				os.Setenv("AGENTBAY_CLI_CONFIG_DIR", origDir)
			}
		}()
		_ = os.Unsetenv(config.EnvAccessKeyID)
		_ = os.Unsetenv(config.EnvAccessKeySecret)
		_ = os.Unsetenv(config.EnvAccessKeySessionToken)

		deleteCmd := findImageDeleteSubcommand(t)
		require.NotNil(t, deleteCmd)
		err = deleteCmd.RunE(deleteCmd, []string{"imgc-any"})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "not authenticated")
	})
}

func TestIsDeletable(t *testing.T) {
	tests := []struct {
		status   string
		expected bool
	}{
		// Non-deletable states
		{"IMAGE_CREATING", false},
		{"RESOURCE_DEPLOYING", false},
		{"RESOURCE_DELETING", false},
		{"RESOURCE_PUBLISHED", false},
		{"RESOURCE_FAILED", false},
		{"RESOURCE_MAINTAINING", false},
		// Deletable states
		{"IMAGE_AVAILABLE", true},
		{"IMAGE_CREATE_FAILED", true},
		{"RESOURCE_CEASED", true},
		// Unknown status defaults to deletable
		{"UNKNOWN_STATUS", true},
		{"", true},
	}

	for _, tt := range tests {
		t.Run(tt.status, func(t *testing.T) {
			result := cmd.IsDeletable(tt.status)
			assert.Equal(t, tt.expected, result, "IsDeletable(%q) = %v, want %v", tt.status, result, tt.expected)
		})
	}
}
