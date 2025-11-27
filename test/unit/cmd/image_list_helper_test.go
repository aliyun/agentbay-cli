// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd_test

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	"github.com/agentbay/agentbay-cli/cmd"
)

// TestImageListCommandFlags tests that the new flags for image list command are properly defined
func TestImageListCommandFlags(t *testing.T) {
	// Find list command
	var listCmd *cobra.Command
	for _, subCmd := range cmd.ImageCmd.Commands() {
		if subCmd.Use == "list" {
			listCmd = subCmd
			break
		}
	}

	require := assert.New(t)
	require.NotNil(listCmd, "image list command should exist")

	t.Run("image list command should have include-system flag", func(t *testing.T) {
		includeSystemFlag := listCmd.Flags().Lookup("include-system")
		assert.NotNil(t, includeSystemFlag, "include-system flag should exist")
		if includeSystemFlag != nil {
			assert.Equal(t, "include-system", includeSystemFlag.Name)
			assert.Equal(t, "Include system images in addition to user images", includeSystemFlag.Usage)
		}
	})

	t.Run("image list command should have system-only flag", func(t *testing.T) {
		systemOnlyFlag := listCmd.Flags().Lookup("system-only")
		assert.NotNil(t, systemOnlyFlag, "system-only flag should exist")
		if systemOnlyFlag != nil {
			assert.Equal(t, "system-only", systemOnlyFlag.Name)
			assert.Equal(t, "Show only system images", systemOnlyFlag.Usage)
		}
	})

	t.Run("image list command should have os-type flag", func(t *testing.T) {
		osTypeFlag := listCmd.Flags().Lookup("os-type")
		assert.NotNil(t, osTypeFlag, "os-type flag should exist")
	})

	t.Run("image list command should have page flag", func(t *testing.T) {
		pageFlag := listCmd.Flags().Lookup("page")
		assert.NotNil(t, pageFlag, "page flag should exist")
	})

	t.Run("image list command should have size flag", func(t *testing.T) {
		sizeFlag := listCmd.Flags().Lookup("size")
		assert.NotNil(t, sizeFlag, "size flag should exist")
	})
}

// TestImageListCommandStructure tests the structure of the image list command
func TestImageListCommandStructure(t *testing.T) {
	// Find list command
	var listCmd *cobra.Command
	for _, subCmd := range cmd.ImageCmd.Commands() {
		if subCmd.Use == "list" {
			listCmd = subCmd
			break
		}
	}

	require := assert.New(t)
	require.NotNil(listCmd, "image list command should exist")

	t.Run("image list command should have correct metadata", func(t *testing.T) {
		assert.Equal(t, "list", listCmd.Use)
		assert.Contains(t, listCmd.Short, "List")
		assert.NotNil(t, listCmd.RunE, "list command should have RunE function")
	})

	t.Run("image list command should be part of image command", func(t *testing.T) {
		assert.Equal(t, "image", cmd.ImageCmd.Use)
		assert.Contains(t, cmd.ImageCmd.Short, "Manage")
	})
}

// Note: The following functions are private in the cmd package and cannot be tested directly:
// - runImageListWithBothTypes: Tests would require mocking agentbay.Client and testing through
//   the public runImageList function, which requires authentication setup. This is better suited
//   for integration tests.
// - printImageTable: This function outputs to stdout and is tested indirectly through command
//   execution. Direct unit testing would require exporting the function or moving tests to the
//   cmd package, which would lose test isolation.
//
// For comprehensive testing of these functions, consider:
// 1. Integration tests that test the full command execution flow
// 2. Moving helper functions to a separate testable package
// 3. Exporting functions with a Test prefix (e.g., TestPrintImageTable) for testing purposes
