// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd_test

import (
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	"github.com/agentbay/agentbay-cli/cmd"
)

func Test{CommandGroup}Cmd(t *testing.T) {
	t.Run("{command-group} command has correct metadata", func(t *testing.T) {
		assert.Equal(t, "{command-group}", cmd.{CommandGroup}Cmd.Use)
		assert.Equal(t, "Manage {feature}", cmd.{CommandGroup}Cmd.Short)
		assert.True(t, strings.Contains(cmd.{CommandGroup}Cmd.Long, "{keyword}"))
	})

	t.Run("{command-group} has {subcommand} subcommand", func(t *testing.T) {
		children := cmd.{CommandGroup}Cmd.Commands()
		names := make([]string, len(children))
		for i, c := range children {
			names[i] = c.Name()
		}
		assert.Contains(t, names, "{subcommand}")
	})
}

func Test{SubCommand}Cmd(t *testing.T) {
	t.Run("{subcommand} command has correct metadata", func(t *testing.T) {
		var {subcommand}Cmd *cobra.Command
		for _, c := range cmd.{CommandGroup}Cmd.Commands() {
			if c.Name() == "{subcommand}" {
				{subcommand}Cmd = c
				break
			}
		}
		
		assert.NotNil(t, {subcommand}Cmd)
		assert.Equal(t, "{subcommand}", {subcommand}Cmd.Use)
		assert.Equal(t, "{Short description}", {subcommand}Cmd.Short)
	})

	t.Run("{subcommand} command has required flags", func(t *testing.T) {
		var {subcommand}Cmd *cobra.Command
		for _, c := range cmd.{CommandGroup}Cmd.Commands() {
			if c.Name() == "{subcommand}" {
				{subcommand}Cmd = c
				break
			}
		}
		
		assert.NotNil(t, {subcommand}Cmd)
		
		{Param1}Flag := {subcommand}Cmd.Flags().Lookup("{param1}")
		assert.NotNil(t, {Param1}Flag)
		assert.Equal(t, "", {Param1}Flag.DefValue)
		assert.True(t, strings.Contains({Param1}Flag.Usage, "required"))
		
		{Param2}Flag := {subcommand}Cmd.Flags().Lookup("{param2}")
		assert.NotNil(t, {Param2}Flag)
		assert.Equal(t, "0", {Param2}Flag.DefValue)
		assert.True(t, strings.Contains({Param2}Flag.Usage, "required"))
	})
}
