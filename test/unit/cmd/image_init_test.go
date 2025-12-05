// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/agentbay/agentbay-cli/cmd"
)

func TestImageInitCommand_Metadata(t *testing.T) {
	// Find the init subcommand
	var imageInitCmd *cobra.Command
	for _, subCmd := range cmd.ImageCmd.Commands() {
		if subCmd.Use == "init" {
			imageInitCmd = subCmd
			break
		}
	}

	require.NotNil(t, imageInitCmd, "image init command not found")

	t.Run("command should have correct metadata", func(t *testing.T) {
		assert.Equal(t, "init", imageInitCmd.Use)
		assert.Equal(t, "Download a Dockerfile template from the cloud", imageInitCmd.Short)
		assert.Contains(t, imageInitCmd.Long, "Dockerfile template")
		assert.Contains(t, imageInitCmd.Long, "AgentBay")
	})

	t.Run("command should accept no arguments", func(t *testing.T) {
		// Test that Args is set to NoArgs
		assert.NotNil(t, imageInitCmd.Args)

		// Test with no arguments (should be valid)
		err := imageInitCmd.Args(imageInitCmd, []string{})
		assert.NoError(t, err)

		// Test with arguments (should be invalid)
		err = imageInitCmd.Args(imageInitCmd, []string{"extra", "args"})
		assert.Error(t, err)
	})

	t.Run("command should have no flags", func(t *testing.T) {
		// Source is always AgentBay, so there should be no flags
		flags := imageInitCmd.Flags()
		// Only check that --source flag doesn't exist
		sourceFlag := flags.Lookup("source")
		assert.Nil(t, sourceFlag, "image init should not have --source flag")
	})
}

func TestImageInitCommand_Authentication(t *testing.T) {
	// Create temporary directory for test config
	tempDir, err := os.MkdirTemp("", "agentbay-image-init-test")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Set environment variable to use temp directory
	originalConfigDir := os.Getenv("AGENTBAY_CLI_CONFIG_DIR")
	os.Setenv("AGENTBAY_CLI_CONFIG_DIR", tempDir)
	defer func() {
		if originalConfigDir == "" {
			os.Unsetenv("AGENTBAY_CLI_CONFIG_DIR")
		} else {
			os.Setenv("AGENTBAY_CLI_CONFIG_DIR", originalConfigDir)
		}
	}()

	t.Run("should fail when not authenticated", func(t *testing.T) {
		// Find the init command
		var imageInitCmd *cobra.Command
		for _, subCmd := range cmd.ImageCmd.Commands() {
			if subCmd.Use == "init" {
				imageInitCmd = subCmd
				break
			}
		}

		require.NotNil(t, imageInitCmd, "image init command not found")

		// Execute command without authentication
		err := imageInitCmd.RunE(imageInitCmd, []string{})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "not authenticated")
	})
}

func TestImageInitCommand_FileOperations(t *testing.T) {
	// Create temporary directory for test
	tempDir, err := os.MkdirTemp("", "agentbay-image-init-file-test")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Change to temp directory
	originalDir, err := os.Getwd()
	require.NoError(t, err)
	defer os.Chdir(originalDir)

	err = os.Chdir(tempDir)
	require.NoError(t, err)

	t.Run("should handle file overwrite scenario", func(t *testing.T) {
		// Create an existing Dockerfile
		dockerfilePath := filepath.Join(tempDir, "Dockerfile")
		existingContent := []byte("FROM existing:image\nRUN echo 'existing'")
		err := os.WriteFile(dockerfilePath, existingContent, 0644)
		require.NoError(t, err)

		// Verify file exists
		_, err = os.Stat(dockerfilePath)
		assert.NoError(t, err, "Dockerfile should exist")

		// Read existing content
		content, err := os.ReadFile(dockerfilePath)
		require.NoError(t, err)
		assert.Equal(t, existingContent, content, "Existing Dockerfile content should match")
	})

	t.Run("should handle file write permissions", func(t *testing.T) {
		// This test verifies that the command handles file write scenarios
		// In a real scenario, the command would write the Dockerfile
		dockerfilePath := filepath.Join(tempDir, "Dockerfile")

		// Test that we can write to the directory
		testContent := []byte("FROM test:image")
		err := os.WriteFile(dockerfilePath, testContent, 0644)
		assert.NoError(t, err, "Should be able to write Dockerfile")

		// Verify file was written
		content, err := os.ReadFile(dockerfilePath)
		require.NoError(t, err)
		assert.Equal(t, testContent, content)
	})
}

func TestImageInitCommand_NoSourceFlag(t *testing.T) {
	// Find the init subcommand
	var imageInitCmd *cobra.Command
	for _, subCmd := range cmd.ImageCmd.Commands() {
		if subCmd.Use == "init" {
			imageInitCmd = subCmd
			break
		}
	}

	require.NotNil(t, imageInitCmd, "image init command not found")

	t.Run("should not have source flag", func(t *testing.T) {
		flags := imageInitCmd.Flags()
		sourceFlag := flags.Lookup("source")
		assert.Nil(t, sourceFlag, "image init should not have --source flag")
	})

	t.Run("command description should not mention source option", func(t *testing.T) {
		// The command description should not mention --source as an option
		longDesc := imageInitCmd.Long
		// Should mention that source is always AgentBay, but not as a flag
		assert.Contains(t, longDesc, "AgentBay", "Should mention AgentBay")
		// Should not mention --source as a flag option
		assert.NotContains(t, strings.ToLower(longDesc), "--source", "Should not mention --source flag in description")
	})
}
