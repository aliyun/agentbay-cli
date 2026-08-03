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

func findSetPreOpenSubcommand(t *testing.T) *cobra.Command {
	t.Helper()
	for _, sub := range cmd.ImageCmd.Commands() {
		if sub.Name() == "set-pre-open" {
			return sub
		}
	}
	return nil
}

func TestImageSetPreOpenCmd(t *testing.T) {
	t.Run("set-pre-open command has correct metadata", func(t *testing.T) {
		setPreOpenCmd := findSetPreOpenSubcommand(t)
		require.NotNil(t, setPreOpenCmd, "image set-pre-open subcommand not found")
		assert.Equal(t, "set-pre-open", setPreOpenCmd.Use)
		assert.Equal(t, "Set the pre-open (reserveMinAmount) for an activated ACS image", setPreOpenCmd.Short)
		assert.Contains(t, setPreOpenCmd.Long, "reserveMinAmount")
	})

	t.Run("set-pre-open command has required flags", func(t *testing.T) {
		setPreOpenCmd := findSetPreOpenSubcommand(t)
		require.NotNil(t, setPreOpenCmd)

		imageIdFlag := setPreOpenCmd.Flags().Lookup("image-id")
		assert.NotNil(t, imageIdFlag, "image-id flag should exist")
		assert.Equal(t, "", imageIdFlag.DefValue)

		preOpenFlag := setPreOpenCmd.Flags().Lookup("pre-open")
		assert.NotNil(t, preOpenFlag, "pre-open flag should exist")
		assert.Equal(t, "0", preOpenFlag.DefValue)
	})

	t.Run("set-pre-open is registered under ImageCmd", func(t *testing.T) {
		children := cmd.ImageCmd.Commands()
		names := make([]string, len(children))
		for i, c := range children {
			names[i] = c.Name()
		}
		assert.Contains(t, names, "set-pre-open")
	})
}
