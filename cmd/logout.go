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

	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	token, tokenErr := cfg.GetToken()
	if tokenErr == nil && token.RefreshToken != "" {
		fmt.Println("Revoking server tokens...")
		err = auth.RevokeTokenWithHint(GetClientID(), token.RefreshToken, "refresh_token")
		if err != nil {
			fmt.Printf("Warning: Could not revoke refresh token: %v\n", err)
		} else {
			fmt.Println("Refresh token revoked successfully")
		}
	} else if tokenErr == nil && token.AccessToken != "" {
		fmt.Println("No refresh token to revoke; clearing local OAuth data only")
	} else {
		fmt.Println("No OAuth tokens in local config")
	}

	fmt.Println("Clearing local authentication data...")
	err = cfg.ClearTokens()
	if err != nil {
		return fmt.Errorf("failed to clear local authentication data: %w", err)
	}

	if config.HasAccessKeyFromEnv() {
		fmt.Printf("Note: %s and %s are still set; unset them to stop using access key authentication.\n",
			config.EnvAccessKeyID, config.EnvAccessKeySecret)
	}

	fmt.Println("Successfully logged out from AgentBay")
	return nil
}
