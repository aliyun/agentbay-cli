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

	t.Run("skills has subcommands push update list show delete", func(t *testing.T) {
		children := cmd.SkillsCmd.Commands()
		names := make([]string, len(children))
		for i, c := range children {
			names[i] = c.Name()
		}
		assert.Contains(t, names, "push")
		assert.Contains(t, names, "update")
		assert.Contains(t, names, "list")
		assert.Contains(t, names, "show")
		assert.Contains(t, names, "delete")
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

	t.Run("skills push has --tag flag", func(t *testing.T) {
		var push *cobra.Command
		for _, c := range cmd.SkillsCmd.Commands() {
			if c.Name() == "push" {
				push = c
				break
			}
		}
		requireNotNil(t, push)
		tagFlag := push.Flags().Lookup("tag")
		assert.NotNil(t, tagFlag)
		assert.Equal(t, "[]", tagFlag.DefValue)
	})

	t.Run("skills push has --icon flag with default value", func(t *testing.T) {
		var push *cobra.Command
		for _, c := range cmd.SkillsCmd.Commands() {
			if c.Name() == "push" {
				push = c
				break
			}
		}
		requireNotNil(t, push)
		iconFlag := push.Flags().Lookup("icon")
		assert.NotNil(t, iconFlag)
		assert.Equal(t, "https://img.alicdn.com/imgextra/i4/O1CN01syuoCy1qhsZxbwuBz_!!6000000005528-2-tps-100-100.png", iconFlag.DefValue)
	})

	t.Run("skills update accepts no positional arguments", func(t *testing.T) {
		var updateCmd *cobra.Command
		for _, c := range cmd.SkillsCmd.Commands() {
			if c.Name() == "update" {
				updateCmd = c
				break
			}
		}
		requireNotNil(t, updateCmd)
		assert.NoError(t, updateCmd.Args(updateCmd, []string{}))
		assert.Error(t, updateCmd.Args(updateCmd, []string{"extra"}))
	})

	t.Run("skills update has --skill-id flag (required)", func(t *testing.T) {
		var updateCmd *cobra.Command
		for _, c := range cmd.SkillsCmd.Commands() {
			if c.Name() == "update" {
				updateCmd = c
				break
			}
		}
		requireNotNil(t, updateCmd)
		skillIdFlag := updateCmd.Flags().Lookup("skill-id")
		assert.NotNil(t, skillIdFlag)
		assert.Equal(t, "", skillIdFlag.DefValue)
	})

	t.Run("skills update has --file flag (required)", func(t *testing.T) {
		var updateCmd *cobra.Command
		for _, c := range cmd.SkillsCmd.Commands() {
			if c.Name() == "update" {
				updateCmd = c
				break
			}
		}
		requireNotNil(t, updateCmd)
		fileFlag := updateCmd.Flags().Lookup("file")
		assert.NotNil(t, fileFlag)
		assert.Equal(t, "", fileFlag.DefValue)
		// Verify --file is marked as required
		requiredAnnotation := fileFlag.Annotations[cobra.BashCompOneRequiredFlag]
		assert.Equal(t, []string{"true"}, requiredAnnotation)
	})

	t.Run("skills update has --tag flag", func(t *testing.T) {
		var updateCmd *cobra.Command
		for _, c := range cmd.SkillsCmd.Commands() {
			if c.Name() == "update" {
				updateCmd = c
				break
			}
		}
		requireNotNil(t, updateCmd)
		tagFlag := updateCmd.Flags().Lookup("tag")
		assert.NotNil(t, tagFlag)
		assert.Equal(t, "[]", tagFlag.DefValue)
	})

	t.Run("skills update has --icon flag", func(t *testing.T) {
		var updateCmd *cobra.Command
		for _, c := range cmd.SkillsCmd.Commands() {
			if c.Name() == "update" {
				updateCmd = c
				break
			}
		}
		requireNotNil(t, updateCmd)
		iconFlag := updateCmd.Flags().Lookup("icon")
		assert.NotNil(t, iconFlag)
		assert.Equal(t, "", iconFlag.DefValue)
	})

	t.Run("skills list has --page flag with default 1", func(t *testing.T) {
		var listCmd *cobra.Command
		for _, c := range cmd.SkillsCmd.Commands() {
			if c.Name() == "list" {
				listCmd = c
				break
			}
		}
		requireNotNil(t, listCmd)
		pageFlag := listCmd.Flags().Lookup("page")
		assert.NotNil(t, pageFlag)
		assert.Equal(t, "1", pageFlag.DefValue)
	})

	t.Run("skills list has --size flag with default 10", func(t *testing.T) {
		var listCmd *cobra.Command
		for _, c := range cmd.SkillsCmd.Commands() {
			if c.Name() == "list" {
				listCmd = c
				break
			}
		}
		requireNotNil(t, listCmd)
		sizeFlag := listCmd.Flags().Lookup("size")
		assert.NotNil(t, sizeFlag)
		assert.Equal(t, "10", sizeFlag.DefValue)
	})

	t.Run("skills list has --name flag", func(t *testing.T) {
		var listCmd *cobra.Command
		for _, c := range cmd.SkillsCmd.Commands() {
			if c.Name() == "list" {
				listCmd = c
				break
			}
		}
		requireNotNil(t, listCmd)
		nameFlag := listCmd.Flags().Lookup("name")
		assert.NotNil(t, nameFlag)
		assert.Equal(t, "", nameFlag.DefValue)
	})

	t.Run("skills list has --tag flag", func(t *testing.T) {
		var listCmd *cobra.Command
		for _, c := range cmd.SkillsCmd.Commands() {
			if c.Name() == "list" {
				listCmd = c
				break
			}
		}
		requireNotNil(t, listCmd)
		tagFlag := listCmd.Flags().Lookup("tag")
		assert.NotNil(t, tagFlag)
		assert.Equal(t, "[]", tagFlag.DefValue)
	})

	t.Run("skills delete accepts up to 1 positional argument", func(t *testing.T) {
		var deleteCmd *cobra.Command
		for _, c := range cmd.SkillsCmd.Commands() {
			if c.Name() == "delete" {
				deleteCmd = c
				break
			}
		}
		requireNotNil(t, deleteCmd)
		// No args is valid (--skill-id flag may be used instead)
		assert.NoError(t, deleteCmd.Args(deleteCmd, []string{}))
		// Single positional arg is valid
		assert.NoError(t, deleteCmd.Args(deleteCmd, []string{"skill-xxxxxxxxxxxxxxxx"}))
		// Two or more positional args is invalid
		assert.Error(t, deleteCmd.Args(deleteCmd, []string{"skill-1", "skill-2"}))
	})

	t.Run("skills delete has --skill-id flag (optional)", func(t *testing.T) {
		var deleteCmd *cobra.Command
		for _, c := range cmd.SkillsCmd.Commands() {
			if c.Name() == "delete" {
				deleteCmd = c
				break
			}
		}
		requireNotNil(t, deleteCmd)
		skillIdFlag := deleteCmd.Flags().Lookup("skill-id")
		assert.NotNil(t, skillIdFlag)
		assert.Equal(t, "", skillIdFlag.DefValue)
		// --skill-id is now optional (positional argument is an alternative)
		requiredAnnotation := skillIdFlag.Annotations[cobra.BashCompOneRequiredFlag]
		assert.Nil(t, requiredAnnotation)
	})

	t.Run("skills delete has --yes flag with shorthand y", func(t *testing.T) {
		var deleteCmd *cobra.Command
		for _, c := range cmd.SkillsCmd.Commands() {
			if c.Name() == "delete" {
				deleteCmd = c
				break
			}
		}
		requireNotNil(t, deleteCmd)
		yesFlag := deleteCmd.Flags().Lookup("yes")
		assert.NotNil(t, yesFlag)
		assert.Equal(t, "false", yesFlag.DefValue)
		assert.Equal(t, "y", yesFlag.Shorthand)
	})
}

// requireNotNil helps avoid importing cmd package twice for *cobra.Command type.
func requireNotNil(t *testing.T, c interface{}) {
	t.Helper()
	if c == nil {
		t.Fatal("expected non-nil command")
	}
}
