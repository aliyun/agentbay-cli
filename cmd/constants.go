// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd

import "github.com/agentbay/agentbay-cli/internal/config"

// OAuth constants
var (
	// CallbackPorts is the list of ports to try in order
	CallbackPorts = []string{"3001", "51153", "53153", "55153", "57153"}
	// DefaultCallbackPort is the first port to try
	DefaultCallbackPort = CallbackPorts[0]
)

// GetRedirectURI returns the redirect URI for a given port
func GetRedirectURI(port string) string {
	return "http://localhost:" + port + "/callback"
}

// GetClientID returns the OAuth client ID for the current environment
func GetClientID() string {
	return config.GetClientID()
}
