// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// IsConfirmInput checks if the input string is a valid confirmation response.
// Only "y", "Y", "yes", "YES" are accepted as confirmation.
func IsConfirmInput(input string) bool {
	s := strings.TrimSpace(input)
	switch s {
	case "y", "Y", "yes", "YES":
		return true
	default:
		return false
	}
}

// isStdinTerminal checks whether stdin is connected to a terminal (TTY).
func isStdinTerminal() bool {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return (fi.Mode() & os.ModeCharDevice) != 0
}

// ConfirmPrompt displays a confirmation prompt and reads user input.
// If autoYes is true, it returns true immediately without prompting.
// If stdin is not a TTY and autoYes is false, it returns an error.
// Only y/Y/yes/YES are accepted as confirmation; all other input (including empty) is treated as decline.
func ConfirmPrompt(prompt string, autoYes bool) (bool, error) {
	if autoYes {
		return true, nil
	}

	if !isStdinTerminal() {
		return false, fmt.Errorf("non-interactive environment detected: use --yes to confirm")
	}

	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return false, fmt.Errorf("failed to read input: %w", err)
	}

	return IsConfirmInput(input), nil
}
