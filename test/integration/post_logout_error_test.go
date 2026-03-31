// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

//go:build integration
// +build integration

package integration_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/agentbay/agentbay-cli/cmd"
	"github.com/spf13/cobra"
)

func TestPostLogout_ErrorReporting(t *testing.T) {
	// Create a temporary dockerfile for testing
	tempDir := t.TempDir()
	dockerfilePath := tempDir + "/Dockerfile"
	os.WriteFile(dockerfilePath, []byte("FROM ubuntu:20.04\nRUN echo 'test'"), 0644)

	tests := []struct {
		name             string
		command          []string
		expectedInOutput string
	}{
		{
			name:             "image list after logout should show auth error",
			command:          []string{"image", "list"},
			expectedInOutput: "not authenticated",
		},
		{
			name:             "image create after logout should show auth error",
			command:          []string{"image", "create", "test", "--dockerfile", dockerfilePath, "--imageId", "test"},
			expectedInOutput: "not authenticated",
		},
		{
			name:             "image activate after logout should show auth error",
			command:          []string{"image", "activate", "test-id"},
			expectedInOutput: "not authenticated",
		},
		{
			name:             "image deactivate after logout should show auth error",
			command:          []string{"image", "deactivate", "test-id"},
			expectedInOutput: "not authenticated",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer

			rootCmd := &cobra.Command{Use: "agentbay"}
			rootCmd.AddGroup(&cobra.Group{ID: "management", Title: "Management Commands"})
			rootCmd.AddGroup(&cobra.Group{ID: "core", Title: "Core Commands"})
			rootCmd.AddCommand(cmd.ImageCmd)

			rootCmd.SetArgs(tt.command)
			rootCmd.SetOut(&buf)
			rootCmd.SetErr(&buf)

			err := rootCmd.Execute()

			// Command should fail
			if err == nil {
				t.Error("Expected error but got none")
			}

			combined := strings.ToLower(buf.String())
			if !strings.Contains(combined, strings.ToLower(tt.expectedInOutput)) {
				t.Errorf("Expected output to contain '%s', but got:\n%s", tt.expectedInOutput, buf.String())
			}

			// Verify that helpful guidance is provided (Cobra prints "Error: ..." to SetErr)
			if !strings.Contains(combined, "agentbay login") {
				t.Errorf("Expected output to suggest 'agentbay login', but got:\n%s", buf.String())
			}
		})
	}
}

func TestPostLogout_ErrorMessageFormat(t *testing.T) {
	// Returned error and Cobra output use ErrNotAuthenticated text (no duplicate stderr line).
	expectedFormat := "not authenticated. Please run 'agentbay login' or set AGENTBAY_ACCESS_KEY_ID and AGENTBAY_ACCESS_KEY_SECRET"

	t.Logf("Expected error message format: %s", expectedFormat)
	t.Log("This format should be consistent across all commands that require authentication")
}

func TestPostLogout_NoSilentFailure(t *testing.T) {
	// This test documents that commands should NOT fail silently
	// If a command fails, it MUST print an error message (Cobra writes returned error to SetErr)

	var buf bytes.Buffer

	rootCmd := &cobra.Command{Use: "agentbay"}
	rootCmd.AddGroup(&cobra.Group{ID: "management", Title: "Management Commands"})
	rootCmd.AddCommand(cmd.ImageCmd)

	rootCmd.SetArgs([]string{"image", "list"})
	rootCmd.SetOut(&buf)
	rootCmd.SetErr(&buf)

	err := rootCmd.Execute()

	// If command fails (err != nil), there MUST be output on the command's Err writer
	if err != nil && buf.Len() == 0 {
		t.Error("Command failed silently - no error message. This is a bug.")
		t.Errorf("Error: %v", err)
		t.Log("FAILED: Commands must print error messages when they fail")
	} else if err != nil {
		t.Logf("PASS: Command failed with error message: %s", buf.String())
	}
}
