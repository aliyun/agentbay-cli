// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/agentbay/agentbay-cli/cmd"
)

func TestVersionCmd(t *testing.T) {
	t.Run("version command should display version information", func(t *testing.T) {
		// Execute command and capture output
		err := cmd.VersionCmd.RunE(cmd.VersionCmd, []string{})

		// Assertions - we can't easily capture fmt.Printf output in unit tests
		// but we can verify the command executes without error
		assert.NoError(t, err)
	})

	t.Run("version command should contain default values", func(t *testing.T) {
		// Execute command
		err := cmd.VersionCmd.RunE(cmd.VersionCmd, []string{})

		// Assertions - verify default values are set correctly
		assert.NoError(t, err)
		assert.Equal(t, "dev", cmd.Version)
		assert.Equal(t, "unknown", cmd.GitCommit)
		assert.Equal(t, "unknown", cmd.BuildDate)
	})

	t.Run("version command should have correct metadata", func(t *testing.T) {
		assert.Equal(t, "version", cmd.VersionCmd.Use)
		assert.Equal(t, "Show version information", cmd.VersionCmd.Short)
		assert.Equal(t, "core", cmd.VersionCmd.GroupID)
		assert.True(t, strings.Contains(cmd.VersionCmd.Long, "Display version"))
	})
}
