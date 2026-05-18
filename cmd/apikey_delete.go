// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/agentbay/agentbay-cli/internal/agentbay"
	"github.com/agentbay/agentbay-cli/internal/client"
	"github.com/agentbay/agentbay-cli/internal/config"
)

var apikeyDeleteCmd = &cobra.Command{
	Use:   "delete <api-key>",
	Short: "Delete an API key",
	Long: `Delete an API key permanently.

The command looks up the API key by its user-visible value (akm-xxx),
checks its current status, and deletes it.

Rules:
  - Only DISABLED API keys can be deleted directly.
  - If the API key is ENABLED, you will be asked to disable it first.
  - A confirmation prompt is shown before deletion.
  - Use --yes to skip all confirmation prompts (for scripts/CI).

Examples:
  # Delete an API key (interactive, with confirmation prompts)
  agentbay apikey delete akm-xxxxxxxxxxxxxxxx

  # Delete without confirmation prompts (for scripts/CI)
  agentbay apikey delete akm-xxxxxxxxxxxxxxxx --yes
  agentbay apikey delete akm-xxxxxxxxxxxxxxxx -y`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runApiKeyDelete(cmd, args[0])
	},
}

func init() {
	apikeyDeleteCmd.Flags().BoolP("yes", "y", false, "Skip all confirmation prompts (for non-interactive use)")
	ApiKeyCmd.AddCommand(apikeyDeleteCmd)
}

func runApiKeyDelete(cmd *cobra.Command, apiKey string) error {
	autoYes, _ := cmd.Flags().GetBool("yes")

	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}
	apiClient := agentbay.NewClientFromConfig(cfg)
	ctx := context.Background()

	// Step 1/3: Look up API key info
	fmt.Printf("[STEP 1/3] Looking up API key...\n")

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

	if reqID := descResp.Body.GetRequestId(); reqID != "" {
		fmt.Printf("[INFO] DescribeMcpApiKey Request ID: %s\n", reqID)
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

	// Validate status
	if currentStatus != "ENABLED" && currentStatus != "DISABLED" {
		return fmt.Errorf("[ERROR] API key status is '%s', cannot delete (expected ENABLED or DISABLED)", currentStatus)
	}

	// Step 2/3 (only if ENABLED): Disable the API key first
	if currentStatus == "ENABLED" {
		fmt.Printf("\n[INFO] This API key is currently ENABLED. It must be disabled before deletion.\n")

		confirmed, err := ConfirmPrompt("Disable it now? [y/N]: ", autoYes)
		if err != nil {
			return fmt.Errorf("[ERROR] %w", err)
		}
		if !confirmed {
			fmt.Printf("[INFO] Operation cancelled.\n")
			return nil
		}

		fmt.Printf("[STEP 2/3] Disabling API key...\n")
		disabledStatus := "DISABLED"
		modResp, err := apiClient.ModifyApiKeyStatus(ctx, &client.ModifyApiKeyStatusRequest{
			ApiKey: &apiKeyId,
			Status: &disabledStatus,
		})
		if err != nil {
			printReqIDFromErr(err)
			return fmt.Errorf("[ERROR] Failed to disable API key: %w", err)
		}

		if modResp.Body == nil {
			return fmt.Errorf("[ERROR] Invalid response: missing body")
		}

		if reqID := modResp.Body.GetRequestId(); reqID != "" {
			fmt.Printf("[INFO] ModifyApiKeyStatus Request ID: %s\n", reqID)
		}

		if !modResp.Body.GetSuccess() {
			code := modResp.Body.GetCode()
			msg := ""
			if modResp.Body.Message != nil {
				msg = *modResp.Body.Message
			}
			return fmt.Errorf("[ERROR] Failed to disable API key: Code=%s, Message=%s", code, msg)
		}

		fmt.Printf("[INFO] API key has been disabled.\n")
	} else {
		fmt.Printf("[STEP 2/3] API key is already DISABLED, skipping disable step.\n")
	}

	// Step 3/3: Confirm and delete
	fmt.Printf("[STEP 3/3] Preparing to delete API key...\n")
	fmt.Printf("  ApiKeyId: %s\n", apiKeyId)
	if keyName != "" {
		fmt.Printf("  Name:     %s\n", keyName)
	}
	fmt.Printf("  Status:   DISABLED\n")
	fmt.Println()

	confirmed, err := ConfirmPrompt("Are you sure you want to permanently delete this API key? [y/N]: ", autoYes)
	if err != nil {
		return fmt.Errorf("[ERROR] %w", err)
	}
	if !confirmed {
		fmt.Printf("[INFO] Operation cancelled.\n")
		return nil
	}

	// Build KeyIdListJson: ["ak-xxx"]
	keyIds := []string{apiKeyId}
	keyIdListJSON, err := json.Marshal(keyIds)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to build request: %w", err)
	}
	keyIdListStr := string(keyIdListJSON)

	deleteResp, err := apiClient.DeleteApiKey(ctx, &client.DeleteApiKeyRequest{
		KeyIdListJson: &keyIdListStr,
	})
	if err != nil {
		printReqIDFromErr(err)
		return fmt.Errorf("[ERROR] Failed to delete API key: %w", err)
	}

	if deleteResp.Body == nil {
		return fmt.Errorf("[ERROR] Invalid response: missing body")
	}

	if reqID := deleteResp.Body.GetRequestId(); reqID != "" {
		fmt.Printf("[INFO] DeleteApiKey Request ID: %s\n", reqID)
	}

	// DeleteApiKey API may not return the Success field; use Code as primary indicator.
	// Treat as failure only if Success is explicitly false, or Code is non-empty and not "ok".
	code := deleteResp.Body.GetCode()
	successPtr := deleteResp.Body.Success
	if (successPtr != nil && !*successPtr) || (code != "" && !strings.EqualFold(code, "ok")) {
		msg := deleteResp.Body.GetMessage()
		return fmt.Errorf("[ERROR] Failed to delete API key: Code=%s, Message=%s", code, msg)
	}

	fmt.Println()
	fmt.Printf("[SUCCESS] API key has been deleted.\n")
	fmt.Printf("  ApiKeyId: %s\n", apiKeyId)
	if keyName != "" {
		fmt.Printf("  Name:     %s\n", keyName)
	}

	return nil
}
