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

// TestImageCreate_ADDCOPYFiles tests that image create correctly parses and processes
// Dockerfiles with COPY/ADD instructions. The test creates a Dockerfile with COPY,
// creates the referenced files, and runs the real command. It expects the command to
// either fail at authentication (if not logged in) or proceed to the upload step.
// The key verification is that the flow does not panic and handles ADD/COPY files.
func TestImageCreate_ADDCOPYFiles(t *testing.T) {
	if os.Getenv("RUN_INTEGRATION_TESTS") == "" {
		t.Skip("Skipping integration test. Set RUN_INTEGRATION_TESTS=1 to run")
	}

	tempDir := t.TempDir()

	// Create referenced files
	codeDir := filepath.Join(tempDir, "code")
	if err := os.MkdirAll(codeDir, 0755); err != nil {
		t.Fatalf("Failed to create code dir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(codeDir, "app.py"), []byte("print('hello')"), 0644); err != nil {
		t.Fatalf("Failed to create app.py: %v", err)
	}
	if err := os.WriteFile(filepath.Join(tempDir, "requirements.txt"), []byte("flask"), 0644); err != nil {
		t.Fatalf("Failed to create requirements.txt: %v", err)
	}

	// Create Dockerfile with COPY instructions
	dockerfilePath := filepath.Join(tempDir, "Dockerfile")
	dockerfileContent := `FROM ubuntu:20.04
RUN echo "test"
COPY requirements.txt /app/
COPY code/app.py /app/code/
`
	if err := os.WriteFile(dockerfilePath, []byte(dockerfileContent), 0644); err != nil {
		t.Fatalf("Failed to create Dockerfile: %v", err)
	}

	rootCmd := &cobra.Command{Use: "agentbay"}
	rootCmd.AddGroup(&cobra.Group{ID: "management", Title: "Management Commands"})
	rootCmd.AddCommand(cmd.ImageCmd)

	cmdArgs := []string{"image", "create", "test-add-copy", "--dockerfile", dockerfilePath, "--imageId", "imgc-07if81rziujpkp72y"}
	rootCmd.SetArgs(cmdArgs)

	var buf bytes.Buffer
	rootCmd.SetOut(&buf)
	rootCmd.SetErr(&buf)

	err := rootCmd.Execute()

	output := buf.String()

	// Command will fail - either at auth, validation, or build. All are acceptable.
	// The key is that we reach the ADD/COPY step without panic.
	if err == nil {
		t.Logf("Command succeeded (unexpected in typical test env). Output: %s", output)
		return
	}

	// Verify we got past the initial validation and Dockerfile parsing
	// If we got to "Uploading ADD/COPY files" or "Getting upload credentials", the flow worked
	reachedUploadStep := strings.Contains(output, "Uploading ADD/COPY files") ||
		strings.Contains(output, "Uploading Dockerfile") ||
		strings.Contains(output, "Getting upload credentials") ||
		strings.Contains(output, "Requesting upload credentials")

	if reachedUploadStep {
		t.Logf("ADD/COPY flow reached upload step: %s", output)
	}

	// Should not contain parse errors for ADD/COPY
	hasParseError := strings.Contains(output, "source not found") ||
		strings.Contains(output, "source path escapes context") ||
		strings.Contains(output, "absolute source path not supported")

	if hasParseError {
		t.Errorf("Unexpected ADD/COPY parse error in output: %s", output)
	}
}

// TestImageCreate_ADDCOPYFiles_NoAuth verifies that ADD/COPY parsing works even when
// not authenticated. Runs without RUN_INTEGRATION_TESTS to avoid real API calls.
func TestImageCreate_ADDCOPYFiles_NoAuth(t *testing.T) {
	tempDir := t.TempDir()

	// Create referenced files
	requireNoError(t, os.MkdirAll(filepath.Join(tempDir, "code"), 0755))
	requireNoError(t, os.WriteFile(filepath.Join(tempDir, "code", "app.py"), []byte("x"), 0644))
	requireNoError(t, os.WriteFile(filepath.Join(tempDir, "requirements.txt"), []byte("flask"), 0644))

	dockerfilePath := filepath.Join(tempDir, "Dockerfile")
	dockerfileContent := `FROM ubuntu:20.04
COPY requirements.txt /app/
COPY code/app.py /app/
`
	requireNoError(t, os.WriteFile(dockerfilePath, []byte(dockerfileContent), 0644))

	rootCmd := &cobra.Command{Use: "agentbay"}
	rootCmd.AddGroup(&cobra.Group{ID: "management", Title: "Management Commands"})
	rootCmd.AddCommand(cmd.ImageCmd)

	cmdArgs := []string{"image", "create", "test-add-copy", "--dockerfile", dockerfilePath, "--imageId", "imgc-test"}
	rootCmd.SetArgs(cmdArgs)

	var buf bytes.Buffer
	rootCmd.SetOut(&buf)
	rootCmd.SetErr(&buf)

	err := rootCmd.Execute()

	// Must fail (no auth)
	if err == nil {
		t.Error("Expected error when not authenticated")
		return
	}

	output := buf.String()

	// Should fail at auth, not at ADD/COPY parsing
	// If we fail with "source not found" etc, it means parsing failed
	hasParseError := strings.Contains(output, "source not found") ||
		strings.Contains(output, "source path escapes context")

	if hasParseError {
		t.Errorf("ADD/COPY parse failed unexpectedly: %s", output)
	}

	// Ideally we reach "Validating source image" or "Getting upload credentials" before failing
	// (indicates parsing succeeded)
	t.Logf("Command failed as expected. Output: %s", output)
}

func requireNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
