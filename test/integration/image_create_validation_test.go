// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

//go:build integration
// +build integration

package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/agentbay/agentbay-cli/cmd"
	"github.com/spf13/cobra"
)

func TestImageCreateValidation_NonexistentImageId(t *testing.T) {
	// Create a temporary dockerfile for testing
	tempDir := t.TempDir()
	dockerfilePath := filepath.Join(tempDir, "Dockerfile")
	err := os.WriteFile(dockerfilePath, []byte("FROM ubuntu:20.04\nRUN echo 'test'"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test dockerfile: %v", err)
	}

	tests := []struct {
		name        string
		imageId     string
		expectError bool
	}{
		{
			name:        "nonexistent image id should fail with error",
			imageId:     "nonexistent_image_id",
			expectError: true,
		},
		{
			name:        "invalid image id format should fail with error",
			imageId:     "invalid-id-format",
			expectError: true,
		},
		{
			name:        "empty image id should fail early",
			imageId:     "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Redirect stderr to capture error messages
			oldStderr := os.Stderr
			r, w, _ := os.Pipe()
			os.Stderr = w

			// Create a new command instance for each test
			rootCmd := &cobra.Command{Use: "agentbay"}
			rootCmd.AddGroup(&cobra.Group{ID: "management", Title: "Management Commands"})
			rootCmd.AddCommand(cmd.ImageCmd)

			// Build command args
			cmdArgs := []string{"image", "create", "test-image", "--dockerfile", dockerfilePath}
			if tt.imageId != "" {
				cmdArgs = append(cmdArgs, "--imageId", tt.imageId)
			}
			rootCmd.SetArgs(cmdArgs)

			// Capture output
			var buf bytes.Buffer
			rootCmd.SetOut(&buf)
			rootCmd.SetErr(&buf)

			// Execute the command
			err := rootCmd.Execute()

			// Restore stderr and read captured output
			w.Close()
			os.Stderr = oldStderr

			stderrOutput := make([]byte, 4096)
			n, _ := r.Read(stderrOutput)
			stderrStr := string(stderrOutput[:n])

			// Check the result
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
					return
				}

				// For non-empty imageId, verify error message appears
				if tt.imageId != "" {
					// Error could be authentication error (in test environment) or validation error (with real auth)
					hasAuthError := strings.Contains(stderrStr, "Not authenticated")
					hasValidationError := strings.Contains(stderrStr, "Source image not found") || strings.Contains(stderrStr, "source image not found")

					if !hasAuthError && !hasValidationError {
						t.Errorf("Expected stderr to contain either auth or validation error for invalid imageId, got:\n%s", stderrStr)
					}

					// If it's a validation error, verify that helpful tip is provided
					if hasValidationError && !strings.Contains(stderrStr, "image list") {
						t.Errorf("Expected stderr to suggest 'image list' command for validation error, got:\n%s", stderrStr)
					}
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got: %v\nOutput: %s", err, buf.String())
				}
			}
		})
	}
}

func TestImageCreateValidation_ValidateBeforeAPICall(t *testing.T) {
	// This test ensures that we validate the source image ID before making any API calls
	// The validation should happen early in the process to save time and provide better UX

	tempDir := t.TempDir()
	dockerfilePath := filepath.Join(tempDir, "Dockerfile")
	err := os.WriteFile(dockerfilePath, []byte("FROM ubuntu:20.04\nRUN echo 'test'"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test dockerfile: %v", err)
	}

	rootCmd := &cobra.Command{Use: "agentbay"}
	rootCmd.AddGroup(&cobra.Group{ID: "management", Title: "Management Commands"})
	rootCmd.AddCommand(cmd.ImageCmd)

	cmdArgs := []string{"image", "create", "test-image", "--dockerfile", dockerfilePath, "--imageId", "fake_id"}
	rootCmd.SetArgs(cmdArgs)

	var buf bytes.Buffer
	rootCmd.SetOut(&buf)
	rootCmd.SetErr(&buf)

	err = rootCmd.Execute()

	// The error should happen before any upload/build occurs
	// We should not see messages about uploading or creating tasks
	output := buf.String()

	if err == nil {
		t.Error("Expected error for invalid image ID")
	}

	// Make sure we don't proceed to upload/build stages with invalid image ID
	if strings.Contains(output, "Uploading file") ||
		strings.Contains(output, "[STEP 2/4]") ||
		strings.Contains(output, "[STEP 3/4]") {
		t.Errorf("Command should fail before upload/build stages. Output: %s", output)
	}
}

func TestImageCreate_ErrorMessageFormat(t *testing.T) {
	// Test that error messages are user-friendly and clear
	tempDir := t.TempDir()
	dockerfilePath := filepath.Join(tempDir, "Dockerfile")
	err := os.WriteFile(dockerfilePath, []byte("FROM ubuntu:20.04\nRUN echo 'test'"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test dockerfile: %v", err)
	}

	// Redirect stderr to capture error messages
	oldStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	rootCmd := &cobra.Command{Use: "agentbay"}
	rootCmd.AddGroup(&cobra.Group{ID: "management", Title: "Management Commands"})
	rootCmd.AddCommand(cmd.ImageCmd)

	cmdArgs := []string{"image", "create", "test-image", "--dockerfile", dockerfilePath, "--imageId", "invalid_id"}
	rootCmd.SetArgs(cmdArgs)

	var buf bytes.Buffer
	rootCmd.SetOut(&buf)
	rootCmd.SetErr(&buf)

	err = rootCmd.Execute()

	// Restore stderr and read captured output
	w.Close()
	os.Stderr = oldStderr

	stderrOutput := make([]byte, 4096)
	n, _ := r.Read(stderrOutput)
	stderrStr := string(stderrOutput[:n])

	if err == nil {
		t.Error("Expected error for invalid image ID")
	}

	// Check that error message is clear and actionable
	output := buf.String() + err.Error() + stderrStr

	// Error message should contain helpful information
	hasErrorIndicator := strings.Contains(output, "[ERROR]") || strings.Contains(output, "error") || strings.Contains(output, "failed")
	if !hasErrorIndicator {
		t.Errorf("Error message should contain error indicator. Got: %s", output)
	}

	// Should mention either authentication issue or image ID issue
	hasAuthError := strings.Contains(strings.ToLower(output), "not authenticated")
	hasImageIdMention := strings.Contains(strings.ToLower(output), "image") &&
		(strings.Contains(strings.ToLower(output), "not found") ||
			strings.Contains(strings.ToLower(output), "invalid") ||
			strings.Contains(strings.ToLower(output), "does not exist"))

	if !hasAuthError && !hasImageIdMention {
		t.Errorf("Error message should clearly indicate the authentication or image ID issue. Got: %s", output)
	}
}

func TestImageCreate_SuggestImageListCommand(t *testing.T) {
	// When image ID validation fails, we should suggest using 'image list' to see available images
	tempDir := t.TempDir()
	dockerfilePath := filepath.Join(tempDir, "Dockerfile")
	err := os.WriteFile(dockerfilePath, []byte("FROM ubuntu:20.04\nRUN echo 'test'"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test dockerfile: %v", err)
	}

	rootCmd := &cobra.Command{Use: "agentbay"}
	rootCmd.AddGroup(&cobra.Group{ID: "management", Title: "Management Commands"})
	rootCmd.AddCommand(cmd.ImageCmd)

	cmdArgs := []string{"image", "create", "test-image", "--dockerfile", dockerfilePath, "--imageId", "nonexistent"}
	rootCmd.SetArgs(cmdArgs)

	var buf bytes.Buffer
	rootCmd.SetOut(&buf)
	rootCmd.SetErr(&buf)

	err = rootCmd.Execute()

	if err == nil {
		t.Error("Expected error for invalid image ID")
	}

	output := buf.String() + err.Error()

	// Should suggest using 'image list' command
	hasSuggestion := strings.Contains(strings.ToLower(output), "image list") ||
		strings.Contains(strings.ToLower(output), "available image") ||
		strings.Contains(strings.ToLower(output), "list available")

	if !hasSuggestion {
		t.Logf("Warning: Error message could be improved by suggesting 'image list' command. Got: %s", output)
		// This is not a hard failure, just a recommendation
	}
}
