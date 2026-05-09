// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/agentbay/agentbay-cli/cmd"
)

func TestConfirmPromptAutoYes(t *testing.T) {
	// When autoYes is true, should return true immediately without reading stdin
	confirmed, err := cmd.ConfirmPrompt("Delete?", true)
	assert.NoError(t, err)
	assert.True(t, confirmed)
}

func TestConfirmPromptNonTTY(t *testing.T) {
	// When autoYes is false, the behavior depends on whether stdin is a TTY:
	// - Non-TTY (CI): returns error about --yes
	// - TTY (local terminal): tries to read stdin and gets EOF
	// Both cases should return (false, error)
	confirmed, err := cmd.ConfirmPrompt("Delete?", false)
	assert.Error(t, err)
	assert.False(t, confirmed)
}

func TestIsConfirmInput(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		// Accepted inputs
		{"y", true},
		{"Y", true},
		{"yes", true},
		{"YES", true},
		// With whitespace
		{"y\n", true},
		{"  y  ", true},
		{"yes\n", true},
		// Rejected inputs
		{"n", false},
		{"N", false},
		{"no", false},
		{"NO", false},
		{"", false},
		{" ", false},
		{"yep", false},
		{"Yeah", false},
		{"Yes", false}, // Mixed case not accepted
		{"yEs", false}, // Mixed case not accepted
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := cmd.IsConfirmInput(tt.input)
			assert.Equal(t, tt.expected, result, "IsConfirmInput(%q) = %v, want %v", tt.input, result, tt.expected)
		})
	}
}
