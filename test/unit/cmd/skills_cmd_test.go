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

func TestSkillsCmd(t *testing.T) {
	t.Run("skills command has correct metadata", func(t *testing.T) {
		assert.Equal(t, "skills", cmd.SkillsCmd.Use)
		assert.Equal(t, "Manage AgentBay skills", cmd.SkillsCmd.Short)
		assert.Equal(t, "management", cmd.SkillsCmd.GroupID)
		assert.True(t, strings.Contains(cmd.SkillsCmd.Long, "Push"))
		assert.True(t, strings.Contains(cmd.SkillsCmd.Long, "skill"))
	})

	t.Run("skills has subcommands push list show group", func(t *testing.T) {
		children := cmd.SkillsCmd.Commands()
		names := make([]string, len(children))
		for i, c := range children {
			names[i] = c.Name()
		}
		assert.Contains(t, names, "push")
		assert.Contains(t, names, "list")
		assert.Contains(t, names, "show")
		assert.Contains(t, names, "group")
	})

	t.Run("skills push requires one argument", func(t *testing.T) {
		pushCmd := cmd.SkillsCmd.Commands()
		var push *cobra.Command
		for _, c := range pushCmd {
			if c.Name() == "push" {
				push = c
				break
			}
		}
		requireNotNil(t, push)
		err := push.Args(push, []string{})
		assert.Error(t, err)
		err = push.Args(push, []string{"/some/dir"})
		assert.NoError(t, err)
		err = push.Args(push, []string{"a", "b"})
		assert.Error(t, err)
	})

	t.Run("skills list accepts no arguments", func(t *testing.T) {
		var listCmd *cobra.Command
		for _, c := range cmd.SkillsCmd.Commands() {
			if c.Name() == "list" {
				listCmd = c
				break
			}
		}
		requireNotNil(t, listCmd)
		assert.NoError(t, listCmd.Args(listCmd, []string{}))
		assert.Error(t, listCmd.Args(listCmd, []string{"extra"}))
	})

	t.Run("skills show requires one argument", func(t *testing.T) {
		var showCmd *cobra.Command
		for _, c := range cmd.SkillsCmd.Commands() {
			if c.Name() == "show" {
				showCmd = c
				break
			}
		}
		requireNotNil(t, showCmd)
		assert.Error(t, showCmd.Args(showCmd, []string{}))
		assert.NoError(t, showCmd.Args(showCmd, []string{"skill-123"}))
		assert.Error(t, showCmd.Args(showCmd, []string{"a", "b"}))
	})

	t.Run("skills group has subcommands", func(t *testing.T) {
		var groupCmd *cobra.Command
		for _, c := range cmd.SkillsCmd.Commands() {
			if c.Name() == "group" {
				groupCmd = c
				break
			}
		}
		requireNotNil(t, groupCmd)
		groupChildren := groupCmd.Commands()
		names := make([]string, len(groupChildren))
		for i, c := range groupChildren {
			names[i] = c.Name()
		}
		assert.Contains(t, names, "create")
		assert.Contains(t, names, "list")
		assert.Contains(t, names, "show")
		assert.Contains(t, names, "add-skill")
		assert.Contains(t, names, "remove-skill")
	})
}

// requireNotNil helps avoid importing cmd package twice for *cobra.Command type.
func requireNotNil(t *testing.T, c interface{}) {
	t.Helper()
	if c == nil {
		t.Fatal("expected non-nil command")
	}
}
