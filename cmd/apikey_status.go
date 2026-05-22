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

var (
	apikeyStatusApiKey    string
	apikeyStatusApiKeyId string
)

var apikeyEnableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable an API key",
	Long: `Enable a disabled API key so it can be used for authentication again.

You can identify the API key either by its user-visible value (--api-key, akm-xxx)
or by its internal ID (--api-key-id, ak-xxx). Using --api-key is recommended.

Examples:
  # Enable an API key using the user-visible API Key (recommended)
  agentbay apikey enable --api-key akm-xxxxxxxxxxxxxxxx

  # Enable an API key using the internal API Key ID
  agentbay apikey enable --api-key-id ak-xxxxxxxxxxxxxxxx

  # Enable with verbose output
  agentbay apikey enable --api-key akm-xxxxxxxxxxxxxxxx -v`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runApiKeyStatusChange(cmd, "ENABLED")
	},
}

var apikeyDisableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable an API key",
	Long: `Disable an API key so it can no longer be used for authentication.

You can identify the API key either by its user-visible value (--api-key, akm-xxx)
or by its internal ID (--api-key-id, ak-xxx). Using --api-key is recommended.

Examples:
  # Disable an API key using the user-visible API Key (recommended)
  agentbay apikey disable --api-key akm-xxxxxxxxxxxxxxxx

  # Disable an API key using the internal API Key ID
  agentbay apikey disable --api-key-id ak-xxxxxxxxxxxxxxxx

  # Disable with verbose output
  agentbay apikey disable --api-key akm-xxxxxxxxxxxxxxxx -v`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runApiKeyStatusChange(cmd, "DISABLED")
	},
}

func init() {
	apikeyEnableCmd.Flags().StringVar(&apikeyStatusApiKey, "api-key", "", "User-visible API Key (akm-xxx format, recommended)")
	apikeyEnableCmd.Flags().StringVar(&apikeyStatusApiKeyId, "api-key-id", "", "Internal API Key ID (ak-xxx). Prefer --api-key for normal usage")

	apikeyDisableCmd.Flags().StringVar(&apikeyStatusApiKey, "api-key", "", "User-visible API Key (akm-xxx format, recommended)")
	apikeyDisableCmd.Flags().StringVar(&apikeyStatusApiKeyId, "api-key-id", "", "Internal API Key ID (ak-xxx). Prefer --api-key for normal usage")

	ApiKeyCmd.AddCommand(apikeyEnableCmd)
	ApiKeyCmd.AddCommand(apikeyDisableCmd)
}

func runApiKeyStatusChange(cmd *cobra.Command, targetStatus string) error {
	// Validate mutual exclusivity: --api-key and --api-key-id
	if apikeyStatusApiKey == "" && apikeyStatusApiKeyId == "" {
		return fmt.Errorf("[ERROR] Either --api-key or --api-key-id must be specified. Using --api-key is recommended")
	}
	if apikeyStatusApiKey != "" && apikeyStatusApiKeyId != "" {
		return fmt.Errorf("[ERROR] --api-key and --api-key-id are mutually exclusive; please specify only one")
	}

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

	var apiKeyId string
	var keyName string

	if apikeyStatusApiKey != "" {
		// --api-key path: 2 steps (lookup via DescribeMcpApiKey, then modify)
		fmt.Printf("[STEP 1/2] Looking up API key...\n")

		descResp, err := apiClient.DescribeMcpApiKey(ctx, &client.DescribeMcpApiKeyRequest{
			ApiKey: &apikeyStatusApiKey,
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

		apiKeyId = data.GetApiKeyId()
		if apiKeyId == "" {
			return fmt.Errorf("[ERROR] Invalid response: missing ApiKeyId")
		}

		currentStatus := data.GetStatus()
		keyName = data.GetName()

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

		fmt.Printf("[STEP 2/2] %s API key...\n", action)
	} else {
		// --api-key-id path: 1 step (skip lookup, directly modify)
		apiKeyId = apikeyStatusApiKeyId
		fmt.Printf("[STEP 1/1] %s API key...\n", action)
	}

	fmt.Printf("  ApiKeyId: %s\n", apiKeyId)

	// Call ModifyApiKeyStatus to change the status
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
