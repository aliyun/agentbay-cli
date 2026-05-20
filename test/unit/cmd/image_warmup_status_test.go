// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd_test

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/agentbay/agentbay-cli/cmd"
)

func findWarmupStatusSubcommand(t *testing.T) *cobra.Command {
	t.Helper()
	for _, sub := range cmd.ImageCmd.Commands() {
		if sub.Name() == "warmup-status" {
			return sub
		}
	}
	return nil
}

func TestImageWarmupStatusCmd(t *testing.T) {
	t.Run("warmup-status command has correct metadata", func(t *testing.T) {
		warmupStatusCmd := findWarmupStatusSubcommand(t)
		require.NotNil(t, warmupStatusCmd, "image warmup-status subcommand not found")
		assert.Equal(t, "warmup-status", warmupStatusCmd.Use)
		assert.Contains(t, warmupStatusCmd.Short, "warm-up")
		assert.Contains(t, warmupStatusCmd.Long, "quota")
	})

	t.Run("warmup-status requires no arguments", func(t *testing.T) {
		warmupStatusCmd := findWarmupStatusSubcommand(t)
		require.NotNil(t, warmupStatusCmd)
		err := warmupStatusCmd.Args(warmupStatusCmd, []string{})
		assert.NoError(t, err)
	})

	t.Run("warmup-status is registered under ImageCmd", func(t *testing.T) {
		children := cmd.ImageCmd.Commands()
		names := make([]string, len(children))
		for i, c := range children {
			names[i] = c.Name()
		}
		assert.Contains(t, names, "warmup-status")
	})
}
