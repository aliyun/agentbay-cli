// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd_test

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/agentbay/agentbay-cli/cmd"
	"github.com/agentbay/agentbay-cli/internal/config"
)

func TestLogoutCmd(t *testing.T) {
	// Create temporary directory for test config
	tempDir, err := os.MkdirTemp("", "agentbay-logout-test")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Set environment variable to use temp directory
	originalConfigDir := os.Getenv("AGENTBAY_CLI_CONFIG_DIR")
	os.Setenv("AGENTBAY_CLI_CONFIG_DIR", tempDir)
	defer func() {
		if originalConfigDir == "" {
			os.Unsetenv("AGENTBAY_CLI_CONFIG_DIR")
		} else {
			os.Setenv("AGENTBAY_CLI_CONFIG_DIR", originalConfigDir)
		}
	}()

	t.Run("logout command should have correct metadata", func(t *testing.T) {
		assert.Equal(t, "logout", cmd.LogoutCmd.Use)
		assert.Equal(t, "Log out from AgentBay", cmd.LogoutCmd.Short)
		assert.Equal(t, "core", cmd.LogoutCmd.GroupID)
		assert.True(t, strings.Contains(cmd.LogoutCmd.Long, "invalidating"))
		assert.NotNil(t, cmd.LogoutCmd.RunE)
	})

	t.Run("logout command should accept no arguments", func(t *testing.T) {
		// Test that Args is set to NoArgs
		assert.NotNil(t, cmd.LogoutCmd.Args)

		// Test with no arguments (should be valid)
		err := cmd.LogoutCmd.Args(cmd.LogoutCmd, []string{})
		assert.NoError(t, err)

		// Test with arguments (should be invalid)
		err = cmd.LogoutCmd.Args(cmd.LogoutCmd, []string{"extra", "args"})
		assert.Error(t, err)
	})

	t.Run("logout should work when no tokens exist", func(t *testing.T) {
		// Ensure no tokens exist
		cfg, err := config.GetConfig()
		require.NoError(t, err)
		err = cfg.ClearTokens()
		require.NoError(t, err)

		// Run logout command
		err = cmd.LogoutCmd.RunE(cmd.LogoutCmd, []string{})
		assert.NoError(t, err)
	})

	t.Run("logout should clear tokens when they exist", func(t *testing.T) {
		// First save some tokens
		cfg, err := config.GetConfig()
		require.NoError(t, err)
		err = cfg.SaveTokens("test-token", "Bearer", 3600, "refresh-token", "id-token")
		require.NoError(t, err)
		assert.True(t, cfg.IsAuthenticated())

		// Run logout command
		err = cmd.LogoutCmd.RunE(cmd.LogoutCmd, []string{})
		assert.NoError(t, err)

		// Verify tokens are cleared
		cfg, err = config.GetConfig()
		require.NoError(t, err)
		assert.False(t, cfg.IsAuthenticated())
	})

	t.Run("logout command should be executable", func(t *testing.T) {
		// Just verify the command can be created and has the right structure
		assert.NotNil(t, cmd.LogoutCmd.RunE)

		// Verify it's a valid cobra command
		assert.NoError(t, cmd.LogoutCmd.ValidateArgs([]string{}))
	})
}
