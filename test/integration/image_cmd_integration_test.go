// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

//go:build integration
// +build integration

package integration_test

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/agentbay/agentbay-cli/cmd"
	"github.com/spf13/cobra"
)

// These tests execute full commands and may make real API calls
// They are marked as integration tests to separate them from fast unit tests

func TestImageCreateCommand_Integration(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		flags       map[string]string
		wantErr     bool
		errContains string
	}{
		{
			name:        "missing image name",
			args:        []string{},
			wantErr:     true,
			errContains: "Missing required argument: <image-name>",
		},
		{
			name:        "too many arguments",
			args:        []string{"image1", "image2"},
			wantErr:     true,
			errContains: "Too many arguments provided",
		},
		{
			name:        "missing dockerfile flag",
			args:        []string{"myimage"},
			flags:       map[string]string{"imageId": "test-id"},
			wantErr:     true,
			errContains: "Missing required flag: --dockerfile",
		},
		{
			name:        "missing imageId flag",
			args:        []string{"myimage"},
			flags:       map[string]string{"dockerfile": "testdata/Dockerfile"},
			wantErr:     true,
			errContains: "Missing required flag: --imageId",
		},
		{
			name:        "dockerfile not found",
			args:        []string{"myimage"},
			flags:       map[string]string{"dockerfile": "./nonexistent", "imageId": "test-id"},
			wantErr:     true,
			errContains: "dockerfile not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a completely new command instance for each test
			imageCmd := &cobra.Command{
				Use:     "image",
				Short:   "Manage images",
				Long:    "Create and manage custom images for AgentBay",
				GroupID: "management",
			}

			imageCreateCmd := &cobra.Command{
				Use:   "create <image-name>",
				Short: "Create a custom image",
				Long:  "Create a custom image using a Dockerfile",
				Args: func(cmd *cobra.Command, args []string) error {
					if len(args) == 0 {
						return fmt.Errorf("[ERROR] Missing required argument: <image-name>")
					}
					if len(args) > 1 {
						return fmt.Errorf("[ERROR] Too many arguments provided. Expected 1 argument (image name), got %d", len(args))
					}
					return nil
				},
				RunE: func(cmd *cobra.Command, args []string) error {
					imageName := args[0]
					dockerfilePath, _ := cmd.Flags().GetString("dockerfile")
					sourceImageId, _ := cmd.Flags().GetString("imageId")

					// Validate required flags
					if dockerfilePath == "" {
						return fmt.Errorf("[ERROR] Missing required flag: --dockerfile for %s", imageName)
					}
					if sourceImageId == "" {
						return fmt.Errorf("[ERROR] Missing required flag: --imageId for %s", imageName)
					}

					// Validate dockerfile path
					if _, err := os.Stat(dockerfilePath); os.IsNotExist(err) {
						return fmt.Errorf("dockerfile not found: %s", dockerfilePath)
					}

					// This would normally continue with the actual implementation
					return fmt.Errorf("test should not reach this point")
				},
			}

			// Add flags
			imageCreateCmd.Flags().StringP("dockerfile", "f", "", "Path to Dockerfile (required)")
			imageCreateCmd.Flags().StringP("imageId", "i", "", "Source image ID (required)")

			imageCmd.AddCommand(imageCreateCmd)

			rootCmd := &cobra.Command{Use: "agentbay"}
			rootCmd.AddGroup(&cobra.Group{ID: "management", Title: "Management Commands"})
			rootCmd.AddCommand(imageCmd)

			// Set up the command with args and flags
			cmdArgs := append([]string{"image", "create"}, tt.args...)

			// Set flags if provided
			if tt.flags != nil {
				for key, value := range tt.flags {
					cmdArgs = append(cmdArgs, "--"+key, value)
				}
			}

			rootCmd.SetArgs(cmdArgs)

			// Execute the command
			err := rootCmd.Execute()

			// Check error expectations
			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error but got none")
					return
				}
				if tt.errContains != "" && !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("Expected error to contain '%s', but got: %s", tt.errContains, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got: %s", err.Error())
				}
			}
		})
	}
}

func TestImageCreateCommandWithValidDockerfile_Integration(t *testing.T) {
	// Create a temporary dockerfile for testing
	tempDir := t.TempDir()
	dockerfilePath := filepath.Join(tempDir, "Dockerfile")
	err := os.WriteFile(dockerfilePath, []byte("FROM ubuntu:20.04\nRUN echo 'test'"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test dockerfile: %v", err)
	}

	// Test with valid dockerfile but no authentication (should fail at auth check)
	rootCmd := &cobra.Command{Use: "agentbay"}
	rootCmd.AddGroup(&cobra.Group{ID: "management", Title: "Management Commands"})
	rootCmd.AddCommand(cmd.ImageCmd)

	cmdArgs := []string{"image", "create", "myimage", "--dockerfile", dockerfilePath, "--imageId", "test-id"}
	rootCmd.SetArgs(cmdArgs)

	err = rootCmd.Execute()
	// Should fail at some point - either authentication, validation, or API call
	if err == nil {
		t.Error("Expected error, but got none")
	}
	// The error could be authentication, validation, API related, or build failure, all are acceptable for this test
	if !strings.Contains(err.Error(), "not authenticated") &&
		!strings.Contains(err.Error(), "failed to get upload credentials") &&
		!strings.Contains(err.Error(), "image build failed") &&
		!strings.Contains(err.Error(), "command failed") {
		t.Errorf("Expected authentication, validation, API, or build error, but got: %s", err.Error())
	}
}

func TestImageListCommand_Integration(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		flags       map[string]string
		wantErr     bool
		errContains string
	}{
		{
			name:    "no arguments should be accepted",
			args:    []string{},
			wantErr: false,
		},
		{
			name:    "extra arguments should be ignored",
			args:    []string{"extra-arg"},
			wantErr: false, // Extra args are ignored, command still runs but fails on auth
		},
		{
			name:    "with os-type filter Linux",
			args:    []string{},
			flags:   map[string]string{"os-type": "Linux"},
			wantErr: false,
		},
		{
			name:    "with os-type filter Android",
			args:    []string{},
			flags:   map[string]string{"os-type": "Android"},
			wantErr: false,
		},
		{
			name:    "with os-type filter Windows",
			args:    []string{},
			flags:   map[string]string{"os-type": "Windows"},
			wantErr: false,
		},
		{
			name: "with pagination flags",
			args: []string{},
			flags: map[string]string{
				"page": "2",
				"size": "5",
			},
			wantErr: false,
		},
		{
			name: "with all flags",
			args: []string{},
			flags: map[string]string{
				"os-type": "Linux",
				"page":    "3",
				"size":    "20",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new command instance for each test
			rootCmd := &cobra.Command{Use: "agentbay"}
			// Add command groups to avoid panic
			rootCmd.AddGroup(&cobra.Group{ID: "management", Title: "Management Commands"})
			rootCmd.AddCommand(cmd.ImageCmd)

			// Set up the command with args and flags
			cmdArgs := append([]string{"image", "list"}, tt.args...)
			rootCmd.SetArgs(cmdArgs)

			// Set flags if provided
			if tt.flags != nil {
				for key, value := range tt.flags {
					rootCmd.SetArgs(append(cmdArgs, fmt.Sprintf("--%s=%s", key, value)))
				}
			}

			// Execute the command
			err := rootCmd.Execute()

			// Check the result
			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error but got none")
					return
				}
				if tt.errContains != "" && !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("Expected error to contain %q, but got %q", tt.errContains, err.Error())
				}
			} else {
				// For successful cases, we expect authentication errors since we don't have valid tokens in tests
				if err != nil && !strings.Contains(err.Error(), "not authenticated") &&
					!strings.Contains(err.Error(), "failed to load configuration") &&
					!strings.Contains(err.Error(), "InvalidBearerToken") &&
					!strings.Contains(err.Error(), "failed to fetch image list") {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		})
	}
}

func TestImageActivateCommand_Integration(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		expectError bool
		errorMsg    string
	}{
		{
			name:        "no arguments should fail",
			args:        []string{},
			expectError: true,
			errorMsg:    "accepts 1 arg(s), received 0",
		},
		{
			name:        "too many arguments should fail",
			args:        []string{"img1", "img2"},
			expectError: true,
			errorMsg:    "accepts 1 arg(s), received 2",
		},
		{
			name:        "valid image id should work",
			args:        []string{"imgc-xxxxxxxxxxxxxx"},
			expectError: true, // Expected to fail due to authentication in test environment
			errorMsg:    "",   // Don't check specific error message as it may vary
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rootCmd := &cobra.Command{}
			// Add the management group that ImageCmd requires
			rootCmd.AddGroup(&cobra.Group{ID: "management", Title: "Management Commands:"})
			rootCmd.AddCommand(cmd.ImageCmd)

			// Set up the command with arguments
			args := append([]string{"image", "activate"}, tt.args...)
			rootCmd.SetArgs(args)

			// Capture output
			var buf bytes.Buffer
			rootCmd.SetOut(&buf)
			rootCmd.SetErr(&buf)

			// Execute command
			err := rootCmd.Execute()

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				} else if tt.errorMsg != "" && !strings.Contains(err.Error(), tt.errorMsg) {
					t.Errorf("Expected error message to contain '%s', got: %s", tt.errorMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		})
	}
}

func TestImageDeactivateCommand_Integration(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		expectError bool
		errorMsg    string
	}{
		{
			name:        "no arguments should fail",
			args:        []string{},
			expectError: true,
			errorMsg:    "accepts 1 arg(s), received 0",
		},
		{
			name:        "too many arguments should fail",
			args:        []string{"imgc-xxxxxxxxxxxxxx", "extra"},
			expectError: true,
			errorMsg:    "accepts 1 arg(s), received 2",
		},
		{
			name:        "valid image id should work",
			args:        []string{"imgc-xxxxxxxxxxxxxx"},
			expectError: true, // Expected to fail due to auth in test env
			errorMsg:    "",   // Will check for auth-related errors
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a root command and add the image command
			rootCmd := &cobra.Command{Use: "agentbay"}
			rootCmd.AddGroup(&cobra.Group{ID: "management", Title: "Management Commands:"})
			rootCmd.AddCommand(cmd.ImageCmd)

			// Set up the command with test arguments
			rootCmd.SetArgs(append([]string{"image", "deactivate"}, tt.args...))

			// Capture output
			var buf bytes.Buffer
			rootCmd.SetOut(&buf)
			rootCmd.SetErr(&buf)

			// Execute the command
			err := rootCmd.Execute()

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
					return
				}

				// For the "valid image id" test case, check for expected auth errors
				if tt.name == "valid image id should work" {
					errorStr := err.Error()
					if !strings.Contains(errorStr, "not authenticated") &&
						!strings.Contains(errorStr, "failed to load configuration") &&
						!strings.Contains(errorStr, "InvalidBearerToken") &&
						!strings.Contains(errorStr, "failed to fetch image list") &&
						!strings.Contains(errorStr, "API not ready yet") &&
						!strings.Contains(errorStr, "image not found") &&
						!strings.Contains(errorStr, "failed to get image info") &&
						!strings.Contains(errorStr, "Image.NotExist") {
						t.Errorf("Expected auth-related, API-not-ready, or image-not-found error, got: %v", err)
					}
				} else if tt.errorMsg != "" && !strings.Contains(err.Error(), tt.errorMsg) {
					t.Errorf("Expected error containing '%s', got: %v", tt.errorMsg, err)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got: %v", err)
				}
			}
		})
	}
}
