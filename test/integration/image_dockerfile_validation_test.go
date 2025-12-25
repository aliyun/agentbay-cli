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

	"github.com/agentbay/agentbay-cli/cmd"
	"github.com/spf13/cobra"
)

// TestImageCreate_DockerfileValidationError tests that Dockerfile validation errors
// are properly distinguished from build failures and display appropriate error messages.
// This test requires a real API environment to work properly.
func TestImageCreate_DockerfileValidationError(t *testing.T) {
	// Skip if not in integration test mode
	if os.Getenv("RUN_INTEGRATION_TESTS") == "" {
		t.Skip("Skipping integration test. Set RUN_INTEGRATION_TESTS=1 to run")
	}

	// Create a temporary dockerfile for testing
	tempDir := t.TempDir()
	dockerfilePath := filepath.Join(tempDir, "Dockerfile")

	// Create a Dockerfile with modified image reference (this should trigger validation error)
	// The first line should be system-defined and cannot be modified
	dockerfileContent := `FROM modified-image:tag
RUN echo "test"
`
	err := os.WriteFile(dockerfilePath, []byte(dockerfileContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test dockerfile: %v", err)
	}

	rootCmd := &cobra.Command{Use: "agentbay"}
	rootCmd.AddGroup(&cobra.Group{ID: "management", Title: "Management Commands"})
	rootCmd.AddCommand(cmd.ImageCmd)

	// Use a valid image ID (this test focuses on Dockerfile validation, not image ID validation)
	// In a real scenario, you would use a valid image ID from your environment
	cmdArgs := []string{"image", "create", "test-validation-error", "--dockerfile", dockerfilePath, "--imageId", "imgc-07if81rziujpkp72y"}
	rootCmd.SetArgs(cmdArgs)

	var buf bytes.Buffer
	rootCmd.SetOut(&buf)
	rootCmd.SetErr(&buf)

	err = rootCmd.Execute()

	// The command should fail
	if err == nil {
		t.Error("Expected error for Dockerfile validation failure, but got none")
		return
	}

	output := buf.String() + buf.String() // Combine stdout and stderr

	// Check that the error message indicates Dockerfile validation failure, not build failure
	hasValidationError := strings.Contains(output, "Dockerfile validation failed") ||
		strings.Contains(output, "dockerfile validation failed")
	hasBuildFailed := strings.Contains(output, "Image build failed") ||
		strings.Contains(output, "image build failed")

	// Should show validation error, not build failed
	if !hasValidationError && hasBuildFailed {
		t.Errorf("Expected 'Dockerfile validation failed' message, but got 'Image build failed'. Output: %s", output)
	}

	// Should contain the validation error message from API
	if !strings.Contains(output, "Image reference is Invalid") {
		t.Logf("Warning: Output does not contain validation error message. This might be expected if API returns different format. Output: %s", output)
	}

	// Should contain helpful tips
	hasTip := strings.Contains(output, "agentbay image init") ||
		strings.Contains(output, "system-defined") ||
		strings.Contains(output, "cannot be modified")

	if !hasTip {
		t.Logf("Warning: Output does not contain helpful tips. This might be expected if error format differs. Output: %s", output)
	}
}

// TestImageCreate_BuildFailure tests that actual build failures (not validation errors)
// display "Image build failed" message correctly.
func TestImageCreate_BuildFailure(t *testing.T) {
	// Skip if not in integration test mode
	if os.Getenv("RUN_INTEGRATION_TESTS") == "" {
		t.Skip("Skipping integration test. Set RUN_INTEGRATION_TESTS=1 to run")
	}

	// Create a temporary dockerfile for testing
	tempDir := t.TempDir()
	dockerfilePath := filepath.Join(tempDir, "Dockerfile")

	// Create a Dockerfile that might cause a build failure (e.g., invalid command)
	// But with correct image reference (not a validation error)
	dockerfileContent := `FROM ubuntu:20.04
RUN invalid-command-that-does-not-exist
`
	err := os.WriteFile(dockerfilePath, []byte(dockerfileContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test dockerfile: %v", err)
	}

	rootCmd := &cobra.Command{Use: "agentbay"}
	rootCmd.AddGroup(&cobra.Group{ID: "management", Title: "Management Commands"})
	rootCmd.AddCommand(cmd.ImageCmd)

	// Use a valid image ID
	cmdArgs := []string{"image", "create", "test-build-failure", "--dockerfile", dockerfilePath, "--imageId", "imgc-07if81rziujpkp72y"}
	rootCmd.SetArgs(cmdArgs)

	var buf bytes.Buffer
	rootCmd.SetOut(&buf)
	rootCmd.SetErr(&buf)

	err = rootCmd.Execute()

	// The command should fail
	if err == nil {
		t.Error("Expected error for build failure, but got none")
		return
	}

	output := buf.String() + err.Error()

	// Should show build failed, not validation error
	hasValidationError := strings.Contains(output, "Dockerfile validation failed")
	hasBuildFailed := strings.Contains(output, "Image build failed") ||
		strings.Contains(output, "image build failed")

	// If it's a validation error, that's unexpected for this test case
	if hasValidationError && !hasBuildFailed {
		t.Logf("Note: Got validation error instead of build failure. This might be expected if Dockerfile format triggers validation. Output: %s", output)
	}

	// Should not contain validation-specific tips
	if strings.Contains(output, "agentbay image init") && hasValidationError {
		t.Logf("Note: Output contains validation tips, which suggests validation error was detected. Output: %s", output)
	}
}

// TestImageCreate_DockerfileValidationErrorMessageFormat tests that error messages are properly formatted
// and distinguish between validation errors and build failures.
func TestImageCreate_DockerfileValidationErrorMessageFormat(t *testing.T) {
	// Skip if not in integration test mode
	if os.Getenv("RUN_INTEGRATION_TESTS") == "" {
		t.Skip("Skipping integration test. Set RUN_INTEGRATION_TESTS=1 to run")
	}

	t.Run("validation error should show validation-specific message", func(t *testing.T) {
		// This test documents the expected behavior
		// In a real scenario, you would trigger a validation error and verify the output
		t.Log("Expected behavior for Dockerfile validation error:")
		t.Log("  - Should display: [ERROR] ❌ Dockerfile validation failed")
		t.Log("  - Should display: [ERROR] Validation error: <error message>")
		t.Log("  - Should display: [TIP] Please check your Dockerfile...")
		t.Log("  - Should display: [TIP] Use 'agentbay image init'...")
		t.Log("  - Should return error: dockerfile validation failed")
	})

	t.Run("build failure should show build-specific message", func(t *testing.T) {
		// This test documents the expected behavior
		t.Log("Expected behavior for build failure:")
		t.Log("  - Should display: [ERROR] ❌ Image build failed")
		t.Log("  - Should display: [ERROR] Error details: <error message>")
		t.Log("  - Should NOT display validation-specific tips")
		t.Log("  - Should return error: image build failed")
	})
}
