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
)

func TestLoginCmd(t *testing.T) {
	// Create temporary directory for test config
	tempDir, err := os.MkdirTemp("", "agentbay-login-test")
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

	t.Run("login command should have correct metadata", func(t *testing.T) {
		assert.Equal(t, "login", cmd.LoginCmd.Use)
		assert.Equal(t, "Log in to AgentBay", cmd.LoginCmd.Short)
		assert.Equal(t, "core", cmd.LoginCmd.GroupID)
		assert.True(t, strings.Contains(cmd.LoginCmd.Long, "OAuth"))
		assert.NotNil(t, cmd.LoginCmd.RunE)
	})

	t.Run("login command should accept no arguments", func(t *testing.T) {
		// Test that Args is set to NoArgs
		assert.NotNil(t, cmd.LoginCmd.Args)

		// Test with no arguments (should be valid)
		err := cmd.LoginCmd.Args(cmd.LoginCmd, []string{})
		assert.NoError(t, err)

		// Test with arguments (should be invalid)
		err = cmd.LoginCmd.Args(cmd.LoginCmd, []string{"extra", "args"})
		assert.Error(t, err)
	})

	// Note: We don't test the actual OAuth flow here as it requires
	// browser interaction and external services. Integration tests
	// would be more appropriate for that.
	t.Run("login command should be executable", func(t *testing.T) {
		// Just verify the command can be created and has the right structure
		assert.NotNil(t, cmd.LoginCmd.RunE)

		// Verify it's a valid cobra command
		assert.NoError(t, cmd.LoginCmd.ValidateArgs([]string{}))
	})
}
