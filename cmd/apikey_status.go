// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"context"
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/agentbay/agentbay-cli/internal/agentbay"
	"github.com/agentbay/agentbay-cli/internal/client"
	"github.com/agentbay/agentbay-cli/internal/config"
)

var apikeyEnableCmd = &cobra.Command{
	Use:   "enable <api-key>",
	Short: "Enable an API key",
	Long: `Enable a disabled API key so it can be used for authentication again.

The command looks up the API key by its user-visible value (akm-xxx),
retrieves the internal key ID, and then enables it.

Examples:
  # Enable an API key
  agentbay apikey enable akm-xxxxxxxxxxxxxxxx

  # Enable with verbose output
  agentbay apikey enable akm-xxxxxxxxxxxxxxxx -v`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runApiKeyStatusChange(cmd, args[0], "ENABLED")
	},
}

var apikeyDisableCmd = &cobra.Command{
	Use:   "disable <api-key>",
	Short: "Disable an API key",
	Long: `Disable an API key so it can no longer be used for authentication.

The command looks up the API key by its user-visible value (akm-xxx),
retrieves the internal key ID, and then disables it.

Examples:
  # Disable an API key
  agentbay apikey disable akm-xxxxxxxxxxxxxxxx

  # Disable with verbose output
  agentbay apikey disable akm-xxxxxxxxxxxxxxxx -v`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runApiKeyStatusChange(cmd, args[0], "DISABLED")
	},
}

func init() {
	ApiKeyCmd.AddCommand(apikeyEnableCmd)
	ApiKeyCmd.AddCommand(apikeyDisableCmd)
}

func runApiKeyStatusChange(cmd *cobra.Command, apiKey string, targetStatus string) error {
	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}
	apiClient := agentbay.NewClientFromConfig(cfg)
	ctx := context.Background()

	action := "Enabling"
	if targetStatus == "DISABLED" {
		action = "Disabling"
	}

	fmt.Printf("[STEP 1/2] Looking up API key...\n")

	// Step 1: Call DescribeMcpApiKey to get the internal ApiKeyId
	descResp, err := apiClient.DescribeMcpApiKey(ctx, &client.DescribeMcpApiKeyRequest{
		ApiKey: &apiKey,
	})
	if err != nil {
		printReqIDFromErr(err)
		return fmt.Errorf("[ERROR] Failed to look up API key: %w", err)
	}

	if descResp.Body == nil {
		return fmt.Errorf("[ERROR] Invalid response: missing body")
	}

	descRequestId := descResp.Body.GetRequestId()
	if descRequestId != "" {
		fmt.Printf("[INFO] DescribeMcpApiKey Request ID: %s\n", descRequestId)
	}

	if !descResp.Body.GetSuccess() {
		code := descResp.Body.GetCode()
		msg := ""
		if descResp.Body.Message != nil {
			msg = *descResp.Body.Message
		}
		return fmt.Errorf("[ERROR] Failed to look up API key: Code=%s, Message=%s", code, msg)
	}

	data := descResp.Body.GetData()
	if data == nil {
		return fmt.Errorf("[ERROR] Invalid response: missing data")
	}

	apiKeyId := data.GetApiKeyId()
	if apiKeyId == "" {
		return fmt.Errorf("[ERROR] Invalid response: missing ApiKeyId")
	}

	currentStatus := data.GetStatus()
	keyName := data.GetName()

	fmt.Printf("  ApiKeyId: %s\n", apiKeyId)
	if keyName != "" {
		fmt.Printf("  Name:     %s\n", keyName)
	}
	fmt.Printf("  Status:   %s\n", currentStatus)

	// Check if already in target status
	if currentStatus == targetStatus {
		fmt.Printf("\n[INFO] API key is already %s, no action needed.\n", targetStatus)
		return nil
	}

	// Step 2: Call ModifyApiKeyStatus to change the status
	fmt.Printf("[STEP 2/2] %s API key...\n", action)

	modResp, err := apiClient.ModifyApiKeyStatus(ctx, &client.ModifyApiKeyStatusRequest{
		ApiKey: &apiKeyId,
		Status: &targetStatus,
	})
	if err != nil {
		printReqIDFromErr(err)
		return fmt.Errorf("[ERROR] Failed to modify API key status: %w", err)
	}

	if modResp.Body == nil {
		return fmt.Errorf("[ERROR] Invalid response: missing body")
	}

	modRequestId := modResp.Body.GetRequestId()
	if modRequestId != "" {
		fmt.Printf("[INFO] ModifyApiKeyStatus Request ID: %s\n", modRequestId)
	}

	if !modResp.Body.GetSuccess() {
		code := modResp.Body.GetCode()
		msg := ""
		if modResp.Body.Message != nil {
			msg = *modResp.Body.Message
		}
		return fmt.Errorf("[ERROR] Failed to modify API key status: Code=%s, Message=%s", code, msg)
	}

	fmt.Println()
	fmt.Printf("[SUCCESS] API key has been %s.\n", targetStatus)
	fmt.Printf("  ApiKeyId: %s\n", apiKeyId)
	if keyName != "" {
		fmt.Printf("  Name:     %s\n", keyName)
	}
	fmt.Printf("  Status:   %s\n", targetStatus)

	return nil
}

// printReqIDFromErr extracts and prints RequestId from an error (always, not verbose-gated).
func printReqIDFromErr(err error) {
	if err == nil {
		return
	}
	var errWithID *client.ErrWithRequestID
	if errors.As(err, &errWithID) && errWithID.RequestID != "" {
		fmt.Printf("[INFO] Request ID: %s\n", errWithID.RequestID)
	}
}
