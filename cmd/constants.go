// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd

import "github.com/agentbay/agentbay-cli/internal/config"

// OAuth constants
const (
	RedirectURI  = "http://localhost:3001/callback"
	CallbackPort = "3001"
)

// GetClientID returns the OAuth client ID for the current environment
func GetClientID() string {
	return config.GetClientID()
}
