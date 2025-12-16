// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"strings"
	"testing"

	"github.com/agentbay/agentbay-cli/cmd"
)

// TestValidateCPUMemoryCombo tests the CPU and memory validation logic
func TestValidateCPUMemoryCombo(t *testing.T) {
	tests := []struct {
		name        string
		cpu         int
		memory      int
		expectError bool
		errorMsg    string
	}{
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
			name:        "default_resources",
			cpu:         0,
			memory:      0,
			expectError: false,
		},
		{
			name:        "invalid_cpu_only",
			cpu:         2,
			memory:      0,
			expectError: true,
			errorMsg:    "Both CPU and memory must be specified together",
		},
		{
			name:        "invalid_memory_only",
			cpu:         0,
			memory:      4,
			expectError: true,
			errorMsg:    "Both CPU and memory must be specified together",
		},
		{
			name:        "invalid_combination_2c8g",
			cpu:         2,
			memory:      8,
			expectError: true,
			errorMsg:    "Invalid CPU/Memory combination: 2c8g",
		},
		{
			name:        "invalid_combination_4c4g",
			cpu:         4,
			memory:      4,
			expectError: true,
			errorMsg:    "Invalid CPU/Memory combination: 4c4g",
		},
		{
			name:        "invalid_cpu_value",
			cpu:         3,
			memory:      6,
			expectError: true,
			errorMsg:    "Invalid CPU/Memory combination: 3c6g",
		},
		{
			name:        "invalid_large_values",
			cpu:         16,
			memory:      32,
			expectError: true,
			errorMsg:    "Invalid CPU/Memory combination: 16c32g",
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
		name           string
		cpu            int
		memory         int
		expectedInMsg  []string
	}{
		{
			name:   "cpu_only_error_contains_hints",
			cpu:    2,
			memory: 0,
			expectedInMsg: []string{
				"Both CPU and memory must be specified together",
				"2c4g",
				"4c8g",
				"8c16g",
			},
		},
		{
			name:   "invalid_combo_error_contains_hints",
			cpu:    2,
			memory: 8,
			expectedInMsg: []string{
				"Invalid CPU/Memory combination: 2c8g",
				"2c4g",
				"4c8g",
				"8c16g",
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
			name:        "negative_cpu",
			cpu:         -1,
			memory:      4,
			expectError: true,
		},
		{
			name:        "negative_memory",
			cpu:         2,
			memory:      -1,
			expectError: true,
		},
		{
			name:        "very_large_cpu",
			cpu:         1000,
			memory:      2000,
			expectError: true,
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

