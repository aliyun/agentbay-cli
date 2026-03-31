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

func findImageStatusSubcommand(t *testing.T) *cobra.Command {
	t.Helper()
	for _, sub := range cmd.ImageCmd.Commands() {
		if sub.Name() == "status" {
			return sub
		}
	}
	return nil
}

func TestImageStatusCommand(t *testing.T) {
	t.Run("status subcommand metadata", func(t *testing.T) {
		statusCmd := findImageStatusSubcommand(t)
		require.NotNil(t, statusCmd, "image status subcommand not found")
		assert.Equal(t, "status", statusCmd.Name())
		assert.Contains(t, strings.ToLower(statusCmd.Short), "status")
		assert.Contains(t, strings.ToLower(statusCmd.Long), "resource")
		assert.NotNil(t, statusCmd.RunE)
	})

	t.Run("status requires exactly one argument", func(t *testing.T) {
		statusCmd := findImageStatusSubcommand(t)
		require.NotNil(t, statusCmd)
		assert.NotNil(t, statusCmd.Args)
		assert.NoError(t, statusCmd.Args(statusCmd, []string{"imgc-test"}))
		assert.Error(t, statusCmd.Args(statusCmd, []string{}))
		assert.Error(t, statusCmd.Args(statusCmd, []string{"a", "b"}))
	})

	t.Run("status fails when not authenticated", func(t *testing.T) {
		tempDir, err := os.MkdirTemp("", "agentbay-image-status-test")
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

		statusCmd := findImageStatusSubcommand(t)
		require.NotNil(t, statusCmd)
		err = statusCmd.RunE(statusCmd, []string{"imgc-any"})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "not authenticated")
	})
}
