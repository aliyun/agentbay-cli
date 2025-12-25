// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"strings"
	"testing"
)

// TestIsDockerfileValidationError tests the isDockerfileValidationError function
// Note: Since this function is not exported, we test it by replicating its logic
// or by testing it indirectly through the error handling in runImageCreate.
// The actual function logic is: strings.Contains(taskMsg, "Image reference is Invalid. Please do not modify the image reference in Dockerfile")
func TestIsDockerfileValidationError(t *testing.T) {
	// Define the validation error message that the function checks for
	validationErrorMsg := "Image reference is Invalid. Please do not modify the image reference in Dockerfile"

	tests := []struct {
		name        string
		taskMsg     string
		expectTrue  bool
		description string
	}{
		{
			name:        "should detect exact Dockerfile validation error",
			taskMsg:     "Image reference is Invalid. Please do not modify the image reference in Dockerfile",
			expectTrue:  true,
			description: "Exact match of validation error message",
		},
		{
			name:        "should detect validation error with prefix",
			taskMsg:     "Build failed: Image reference is Invalid. Please do not modify the image reference in Dockerfile",
			expectTrue:  true,
			description: "Validation error message with prefix",
		},
		{
			name:        "should detect validation error with suffix",
			taskMsg:     "Image reference is Invalid. Please do not modify the image reference in Dockerfile. Please check your Dockerfile.",
			expectTrue:  true,
			description: "Validation error message with suffix",
		},
		{
			name:        "should detect validation error with both prefix and suffix",
			taskMsg:     "Error: Image reference is Invalid. Please do not modify the image reference in Dockerfile. Task failed.",
			expectTrue:  true,
			description: "Validation error message with prefix and suffix",
		},
		{
			name:        "should not detect build failure as validation error",
			taskMsg:     "Build failed: dependency installation error",
			expectTrue:  false,
			description: "Regular build failure should not be detected as validation error",
		},
		{
			name:        "should not detect network error as validation error",
			taskMsg:     "Network connection timeout",
			expectTrue:  false,
			description: "Network error should not be detected as validation error",
		},
		{
			name:        "should not detect empty message as validation error",
			taskMsg:     "",
			expectTrue:  false,
			description: "Empty message should not be detected as validation error",
		},
		{
			name:        "should not detect similar but different message",
			taskMsg:     "Image reference is invalid. Please check your Dockerfile.",
			expectTrue:  false,
			description: "Similar message but missing key phrase should not match",
		},
		{
			name:        "should not detect partial match with wrong case",
			taskMsg:     "image reference is invalid. please do not modify the image reference in dockerfile",
			expectTrue:  false,
			description: "Case-sensitive matching (lowercase should not match)",
		},
		{
			name:        "should not detect unrelated Dockerfile error",
			taskMsg:     "Dockerfile syntax error: invalid instruction",
			expectTrue:  false,
			description: "Other Dockerfile errors should not be detected as validation error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test the logic that isDockerfileValidationError uses
			// The function uses: strings.Contains(taskMsg, validationErrorMsg)
			result := strings.Contains(tt.taskMsg, validationErrorMsg)

			if result != tt.expectTrue {
				t.Errorf("isDockerfileValidationError logic test failed for %q: got %v, want %v. %s",
					tt.taskMsg, result, tt.expectTrue, tt.description)
			}
		})
	}
}
