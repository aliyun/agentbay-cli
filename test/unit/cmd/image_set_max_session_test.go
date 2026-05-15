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

func findSetMaxSessionSubcommand(t *testing.T) *cobra.Command {
	t.Helper()
	for _, sub := range cmd.ImageCmd.Commands() {
		if sub.Name() == "set-max-session" {
			return sub
		}
	}
	return nil
}

func TestImageSetMaxSessionCmd(t *testing.T) {
	t.Run("set-max-session command has correct metadata", func(t *testing.T) {
		setMaxSessionCmd := findSetMaxSessionSubcommand(t)
		require.NotNil(t, setMaxSessionCmd, "image set-max-session subcommand not found")
		assert.Equal(t, "set-max-session", setMaxSessionCmd.Use)
		assert.Equal(t, "Set the maximum concurrent session count for an activated User image", setMaxSessionCmd.Short)
		assert.Contains(t, setMaxSessionCmd.Long, "maximum number of concurrent sessions")
	})

	t.Run("set-max-session command has required flags", func(t *testing.T) {
		setMaxSessionCmd := findSetMaxSessionSubcommand(t)
		require.NotNil(t, setMaxSessionCmd)

		imageIdFlag := setMaxSessionCmd.Flags().Lookup("image-id")
		assert.NotNil(t, imageIdFlag, "image-id flag should exist")
		assert.Equal(t, "", imageIdFlag.DefValue)

		maxSessionNumFlag := setMaxSessionCmd.Flags().Lookup("max-session-num")
		assert.NotNil(t, maxSessionNumFlag, "max-session-num flag should exist")
		assert.Equal(t, "0", maxSessionNumFlag.DefValue)
	})

	t.Run("set-max-session is registered under ImageCmd", func(t *testing.T) {
		children := cmd.ImageCmd.Commands()
		names := make([]string, len(children))
		for i, c := range children {
			names[i] = c.Name()
		}
		assert.Contains(t, names, "set-max-session")
	})
}
