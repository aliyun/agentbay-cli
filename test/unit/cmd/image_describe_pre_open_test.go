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

func findDescribePreOpenSubcommand(t *testing.T) *cobra.Command {
	t.Helper()
	for _, sub := range cmd.ImageCmd.Commands() {
		if sub.Name() == "describe-pre-open" {
			return sub
		}
	}
	return nil
}

func TestImageDescribePreOpenCmd(t *testing.T) {
	t.Run("describe-pre-open command has correct metadata", func(t *testing.T) {
		describeCmd := findDescribePreOpenSubcommand(t)
		require.NotNil(t, describeCmd, "image describe-pre-open subcommand not found")
		assert.Equal(t, "describe-pre-open", describeCmd.Use)
		assert.Contains(t, describeCmd.Short, "pre-open")
		assert.Contains(t, describeCmd.Long, "reserveMinAmount")
	})

	t.Run("describe-pre-open command has expected flags", func(t *testing.T) {
		describeCmd := findDescribePreOpenSubcommand(t)
		require.NotNil(t, describeCmd)

		imageIdFlag := describeCmd.Flags().Lookup("image-id")
		assert.NotNil(t, imageIdFlag, "image-id flag should exist")

		nextTokenFlag := describeCmd.Flags().Lookup("next-token")
		assert.NotNil(t, nextTokenFlag, "next-token flag should exist")

		maxResultsFlag := describeCmd.Flags().Lookup("max-results")
		assert.NotNil(t, maxResultsFlag, "max-results flag should exist")
		assert.Equal(t, "20", maxResultsFlag.DefValue)

		outputFlag := describeCmd.Flags().Lookup("output")
		assert.NotNil(t, outputFlag, "output flag should exist")
		assert.Equal(t, "o", outputFlag.Shorthand)
		assert.Equal(t, "", outputFlag.DefValue)
	})

	t.Run("describe-pre-open requires no arguments", func(t *testing.T) {
		describeCmd := findDescribePreOpenSubcommand(t)
		require.NotNil(t, describeCmd)
		err := describeCmd.Args(describeCmd, []string{})
		assert.NoError(t, err)
	})

	t.Run("describe-pre-open rejects positional arguments", func(t *testing.T) {
		describeCmd := findDescribePreOpenSubcommand(t)
		require.NotNil(t, describeCmd)
		err := describeCmd.Args(describeCmd, []string{"unexpected-arg"})
		assert.Error(t, err)
	})

	t.Run("describe-pre-open is registered under ImageCmd", func(t *testing.T) {
		children := cmd.ImageCmd.Commands()
		names := make([]string, len(children))
		for i, c := range children {
			names[i] = c.Name()
		}
		assert.Contains(t, names, "describe-pre-open")
	})
}
