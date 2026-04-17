// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"strings"
	"testing"

	"github.com/agentbay/agentbay-cli/cmd"
)

// TestValidateCPUMemoryCombo tests the CPU and memory validation logic
// Note: Specific combinations are not validated here - the backend
// DescribeInstanceTypes API will handle that with dynamic error messages
func TestValidateCPUMemoryCombo(t *testing.T) {
	tests := []struct {
		name        string
		cpu         int
		memory      int
		expectError bool
		errorMsg    string
	}{
		// Valid cases - any positive combination is allowed (backend will validate)
		{
			name:        "valid_2c4g",
			cpu:         2,
			memory:      4,
			expectError: false,
		},
		{
			name:        "valid_4c8g",
			cpu:         4,
			memory:      8,
			expectError: false,
		},
		{
			name:        "valid_8c16g",
			cpu:         8,
			memory:      16,
			expectError: false,
		},
		{
			name:        "valid_16c32g",
			cpu:         16,
			memory:      32,
			expectError: false,
		},
		{
			name:        "valid_any_combination",
			cpu:         3,
			memory:      6,
			expectError: false, // Backend will validate actual combinations
		},
		{
			name:        "default_resources",
			cpu:         0,
			memory:      0,
			expectError: false,
		},
		// Invalid cases - only one specified
		{
			name:        "invalid_cpu_only",
			cpu:         2,
			memory:      0,
			expectError: true,
			errorMsg:    "both CPU and memory must be specified together",
		},
		{
			name:        "invalid_memory_only",
			cpu:         0,
			memory:      4,
			expectError: true,
			errorMsg:    "both CPU and memory must be specified together",
		},
		// Invalid cases - negative values
		{
			name:        "negative_cpu",
			cpu:         -1,
			memory:      4,
			expectError: true,
			errorMsg:    "CPU and memory must be positive values",
		},
		{
			name:        "negative_memory",
			cpu:         2,
			memory:      -1,
			expectError: true,
			errorMsg:    "CPU and memory must be positive values",
		},
		{
			name:        "negative_both",
			cpu:         -2,
			memory:      -4,
			expectError: true,
			errorMsg:    "CPU and memory must be positive values",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := cmd.ValidateCPUMemoryCombo(tt.cpu, tt.memory)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				} else {
					if !strings.Contains(err.Error(), tt.errorMsg) {
						t.Errorf("Expected error message to contain '%s', got: %s", tt.errorMsg, err.Error())
					}
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got: %v", err)
				}
			}
		})
	}
}

// TestValidateCPUMemoryCombo_ErrorMessages tests error message formatting
func TestValidateCPUMemoryCombo_ErrorMessages(t *testing.T) {
	tests := []struct {
		name          string
		cpu           int
		memory        int
		expectedInMsg []string
	}{
		{
			name:   "cpu_only_error",
			cpu:    2,
			memory: 0,
			expectedInMsg: []string{
				"both CPU and memory must be specified together",
			},
		},
		{
			name:   "negative_cpu_error",
			cpu:    -1,
			memory: 4,
			expectedInMsg: []string{
				"CPU and memory must be positive values",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := cmd.ValidateCPUMemoryCombo(tt.cpu, tt.memory)

			if err == nil {
				t.Fatal("Expected error but got none")
			}

			errMsg := err.Error()
			for _, expected := range tt.expectedInMsg {
				if !strings.Contains(errMsg, expected) {
					t.Errorf("Error message should contain '%s', got: %s", expected, errMsg)
				}
			}
		})
	}
}

// TestValidateCPUMemoryCombo_BoundaryValues tests boundary values
func TestValidateCPUMemoryCombo_BoundaryValues(t *testing.T) {
	tests := []struct {
		name        string
		cpu         int
		memory      int
		expectError bool
	}{
		{
			name:        "zero_values",
			cpu:         0,
			memory:      0,
			expectError: false,
		},
		{
			name:        "large_values",
			cpu:         100,
			memory:      200,
			expectError: false, // Backend will validate actual limits
		},
		{
			name:        "one_and_one",
			cpu:         1,
			memory:      1,
			expectError: false, // Backend will validate actual limits
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := cmd.ValidateCPUMemoryCombo(tt.cpu, tt.memory)

			if tt.expectError && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}
