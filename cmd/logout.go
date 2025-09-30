// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/agentbay/agentbay-cli/internal/auth"
	"github.com/agentbay/agentbay-cli/internal/config"
)

var LogoutCmd = &cobra.Command{
	Use:     "logout",
	Short:   "Log out from AgentBay",
	Long:    "Log out from AgentBay by invalidating server session and clearing local authentication data",
	Args:    cobra.NoArgs,
	GroupID: "core",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runLogout(cmd)
	},
}

func runLogout(cmd *cobra.Command) error {
	fmt.Println("Logging out from AgentBay...")

	// Load configuration
	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	// Check if we have valid tokens for API logout
	hasValidTokens := cfg.IsAuthenticated()

	if hasValidTokens {
		// Attempt to revoke server tokens
		fmt.Println("Revoking server tokens...")

		token, err := cfg.GetTokens()
		if err != nil {
			fmt.Printf("Warning: Could not get tokens for revocation: %v\n", err)
		} else {
			// Try to revoke access token
			err = auth.RevokeToken(ClientID, token.AccessToken)
			if err != nil {
				fmt.Printf("Warning: Could not revoke access token: %v\n", err)
			} else {
				fmt.Println("Access token revoked successfully")
			}

			// Try to revoke refresh token if different
			if token.RefreshToken != "" && token.RefreshToken != token.AccessToken {
				err = auth.RevokeToken(ClientID, token.RefreshToken)
				if err != nil {
					fmt.Printf("Warning: Could not revoke refresh token: %v\n", err)
				} else {
					fmt.Println("Refresh token revoked successfully")
				}
			}
		}
	} else {
		fmt.Println("No active session found")
	}

	// Always perform local cleanup
	fmt.Println("Clearing local authentication data...")

	// Clear tokens from config
	err = cfg.ClearTokens()
	if err != nil {
		return fmt.Errorf("failed to clear local authentication data: %w", err)
	}

	// Success message
	if hasValidTokens {
		fmt.Println("Successfully logged out from AgentBay")
	} else {
		fmt.Println("Successfully logged out from AgentBay (local session cleared)")
	}

	return nil
}
