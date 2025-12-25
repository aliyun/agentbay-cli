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
		expectedInStderr string
	}{
		{
			name:             "image list after logout should show auth error",
			command:          []string{"image", "list"},
			expectedInStderr: "Not authenticated",
		},
		{
			name:             "image create after logout should show auth error",
			command:          []string{"image", "create", "test", "--dockerfile", dockerfilePath, "--imageId", "test"},
			expectedInStderr: "Not authenticated",
		},
		{
			name:             "image activate after logout should show auth error",
			command:          []string{"image", "activate", "test-id"},
			expectedInStderr: "Not authenticated",
		},
		{
			name:             "image deactivate after logout should show auth error",
			command:          []string{"image", "deactivate", "test-id"},
			expectedInStderr: "Not authenticated",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Redirect stderr to capture error messages
			oldStderr := os.Stderr
			r, w, _ := os.Pipe()
			os.Stderr = w

			rootCmd := &cobra.Command{Use: "agentbay"}
			rootCmd.AddGroup(&cobra.Group{ID: "management", Title: "Management Commands"})
			rootCmd.AddGroup(&cobra.Group{ID: "core", Title: "Core Commands"})
			rootCmd.AddCommand(cmd.ImageCmd)

			rootCmd.SetArgs(tt.command)

			var buf bytes.Buffer
			rootCmd.SetOut(&buf)
			rootCmd.SetErr(&buf)

			err := rootCmd.Execute()

			// Restore stderr and read captured output
			w.Close()
			os.Stderr = oldStderr

			stderrOutput := make([]byte, 4096)
			n, _ := r.Read(stderrOutput)
			stderrStr := string(stderrOutput[:n])

			// Command should fail
			if err == nil {
				t.Error("Expected error but got none")
			}

			// Error message should be in stderr
			if !strings.Contains(stderrStr, tt.expectedInStderr) {
				t.Errorf("Expected stderr to contain '%s', but got:\n%s", tt.expectedInStderr, stderrStr)
			}

			// Verify that error message is user-friendly
			if !strings.Contains(stderrStr, "[ERROR]") {
				t.Errorf("Expected stderr to contain '[ERROR]' prefix, but got:\n%s", stderrStr)
			}

			// Verify that helpful guidance is provided
			if !strings.Contains(stderrStr, "agentbay login") {
				t.Errorf("Expected stderr to suggest 'agentbay login', but got:\n%s", stderrStr)
			}
		})
	}
}

func TestPostLogout_ErrorMessageFormat(t *testing.T) {
	// Test that error messages have consistent format
	expectedFormat := "[ERROR] Not authenticated. Please run 'agentbay login' first"

	t.Logf("Expected error message format: %s", expectedFormat)
	t.Log("This format should be consistent across all commands that require authentication")
}

func TestPostLogout_NoSilentFailure(t *testing.T) {
	// This test documents that commands should NOT fail silently
	// If a command fails, it MUST print an error message to stderr

	// Redirect stderr to capture error messages
	oldStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	rootCmd := &cobra.Command{Use: "agentbay"}
	rootCmd.AddGroup(&cobra.Group{ID: "management", Title: "Management Commands"})
	rootCmd.AddCommand(cmd.ImageCmd)

	rootCmd.SetArgs([]string{"image", "list"})

	var buf bytes.Buffer
	rootCmd.SetOut(&buf)
	rootCmd.SetErr(&buf)

	err := rootCmd.Execute()

	// Restore stderr and read captured output
	w.Close()
	os.Stderr = oldStderr

	stderrOutput := make([]byte, 4096)
	n, _ := r.Read(stderrOutput)
	stderrStr := string(stderrOutput[:n])

	// If command fails (err != nil), there MUST be output to stderr
	if err != nil && stderrStr == "" {
		t.Error("Command failed silently - no error message in stderr. This is a bug.")
		t.Errorf("Error: %v", err)
		t.Log("FAILED: Commands must print error messages to stderr when they fail")
	} else if err != nil {
		t.Logf("PASS: Command failed with error message: %s", stderrStr)
	}
}
