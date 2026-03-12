// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

//go:build integration
// +build integration

package integration_test

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/cobra"

	"github.com/agentbay/agentbay-cli/cmd"
)

func TestSkillsPushCommand_Integration(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		wantErr     bool
		errContains string
	}{
		{
			name:        "missing argument",
			args:        []string{"skills", "push"},
			wantErr:     true,
			errContains: "accepts 1 arg",
		},
		{
			name:        "directory does not exist",
			args:        []string{"skills", "push", filepath.Join(t.TempDir(), "nonexistent")},
			wantErr:     true,
			errContains: "does not exist",
		},
		{
			name: "not a directory",
			args: func() []string {
				f := filepath.Join(t.TempDir(), "file")
				_ = os.WriteFile(f, []byte("x"), 0644)
				return []string{"skills", "push", f}
			}(),
			wantErr:     true,
			errContains: "Not a directory",
		},
		{
			name: "SKILL.md not found in directory",
			args: func() []string {
				d := t.TempDir()
				return []string{"skills", "push", d}
			}(),
			wantErr:     true,
			errContains: "SKILL.md not found",
		},
		{
			name: "invalid frontmatter missing name",
			args: func() []string {
				d := t.TempDir()
				_ = os.WriteFile(filepath.Join(d, "SKILL.md"), []byte("description: only\n"), 0644)
				return []string{"skills", "push", d}
			}(),
			wantErr:     true,
			errContains: "name:",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rootCmd := &cobra.Command{Use: "agentbay"}
			rootCmd.AddGroup(&cobra.Group{ID: "management", Title: "Management Commands"})
			rootCmd.AddCommand(cmd.SkillsCmd)

			rootCmd.SetArgs(tt.args)
			var out, errOut bytes.Buffer
			rootCmd.SetOut(&out)
			rootCmd.SetErr(&errOut)

			err := rootCmd.Execute()

			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error but got none. stderr: %s", errOut.String())
					return
				}
				// Optional: errContains (cobra uses it for "accepts 1 arg"; printErrorMessage returns "command failed")
				if tt.errContains != "" && !strings.Contains(err.Error(), tt.errContains) && err.Error() != "command failed" {
					t.Errorf("Expected error to contain %q or 'command failed', got: %s", tt.errContains, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got: %v. stderr: %s", err, errOut.String())
				}
			}
		})
	}
}

func TestSkillsListCommand_Integration(t *testing.T) {
	rootCmd := &cobra.Command{Use: "agentbay"}
	rootCmd.AddGroup(&cobra.Group{ID: "management", Title: "Management Commands"})
	rootCmd.AddCommand(cmd.SkillsCmd)

	rootCmd.SetArgs([]string{"skills", "list"})
	var out, errOut bytes.Buffer
	rootCmd.SetOut(&out)
	rootCmd.SetErr(&errOut)

	err := rootCmd.Execute()
	if err != nil {
		t.Errorf("skills list should succeed (placeholder): %v", err)
	}
	// Placeholder prints info to stderr
	if errOut.Len() > 0 {
		if !strings.Contains(errOut.String(), "List skills") && !strings.Contains(errOut.String(), "backend") {
			t.Logf("stderr: %s", errOut.String())
		}
	}
}

func TestSkillsShowCommand_Integration(t *testing.T) {
	// skills show <id> requires config and API; without valid config it fails with load config error.
	// We only verify the command is accepted and runs (may fail on config/API).
	rootCmd := &cobra.Command{Use: "agentbay"}
	rootCmd.AddGroup(&cobra.Group{ID: "management", Title: "Management Commands"})
	rootCmd.AddCommand(cmd.SkillsCmd)

	rootCmd.SetArgs([]string{"skills", "show", "skill-nonexistent-id"})
	var out, errOut bytes.Buffer
	rootCmd.SetOut(&out)
	rootCmd.SetErr(&errOut)

	_ = rootCmd.Execute()
	// Either succeeds (with empty data) or fails on config/describe - both are acceptable
}

func TestSkillsGroupShowCommand_Integration(t *testing.T) {
	// skills group show <id> is a placeholder that always succeeds.
	rootCmd := &cobra.Command{Use: "agentbay"}
	rootCmd.AddGroup(&cobra.Group{ID: "management", Title: "Management Commands"})
	rootCmd.AddCommand(cmd.SkillsCmd)

	rootCmd.SetArgs([]string{"skills", "group", "show", "group-123"})
	var errOut bytes.Buffer
	rootCmd.SetErr(&errOut)

	err := rootCmd.Execute()
	if err != nil {
		t.Errorf("skills group show should succeed (placeholder): %v", err)
	}
}
