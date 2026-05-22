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

var (
	apikeyDeleteApiKey    string
	apikeyDeleteApiKeyId string
)

var apikeyDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an API key",
	Long: `Delete an API key permanently.

You can identify the API key either by its user-visible value (--api-key, akm-xxx)
or by its internal ID (--api-key-id, ak-xxx). Using --api-key is recommended.

Rules:
  - Only DISABLED API keys can be deleted directly.
  - If the API key is ENABLED, you will be asked to disable it first.
  - A confirmation prompt is shown before deletion.
  - Use --yes to skip all confirmation prompts (for scripts/CI).

Examples:
  # Delete an API key using the user-visible API Key (interactive, with confirmation prompts)
  agentbay apikey delete --api-key akm-xxxxxxxxxxxxxxxx

  # Delete an API key using the internal API Key ID
  agentbay apikey delete --api-key-id ak-xxxxxxxxxxxxxxxx

  # Delete without confirmation prompts (for scripts/CI)
  agentbay apikey delete --api-key akm-xxxxxxxxxxxxxxxx --yes
  agentbay apikey delete --api-key-id ak-xxxxxxxxxxxxxxxx -y`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runApiKeyDelete(cmd)
	},
}

func init() {
	apikeyDeleteCmd.Flags().StringVar(&apikeyDeleteApiKey, "api-key", "", "User-visible API Key (akm-xxx format, recommended)")
	apikeyDeleteCmd.Flags().StringVar(&apikeyDeleteApiKeyId, "api-key-id", "", "Internal API Key ID (ak-xxx). Prefer --api-key for normal usage")
	apikeyDeleteCmd.Flags().BoolP("yes", "y", false, "Skip all confirmation prompts (for non-interactive use)")
	ApiKeyCmd.AddCommand(apikeyDeleteCmd)
}

func runApiKeyDelete(cmd *cobra.Command) error {
	// Validate mutual exclusivity: --api-key and --api-key-id
	if apikeyDeleteApiKey == "" && apikeyDeleteApiKeyId == "" {
		return fmt.Errorf("[ERROR] Either --api-key or --api-key-id must be specified. Using --api-key is recommended")
	}
	if apikeyDeleteApiKey != "" && apikeyDeleteApiKeyId != "" {
		return fmt.Errorf("[ERROR] --api-key and --api-key-id are mutually exclusive; please specify only one")
	}

	autoYes, _ := cmd.Flags().GetBool("yes")

	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}
	apiClient := agentbay.NewClientFromConfig(cfg)
	ctx := context.Background()

	var apiKeyId string
	var keyName string
	var currentStatus string

	if apikeyDeleteApiKey != "" {
		// --api-key path: lookup via DescribeMcpApiKey
		fmt.Printf("[STEP 1/3] Looking up API key...\n")

		descResp, err := apiClient.DescribeMcpApiKey(ctx, &client.DescribeMcpApiKeyRequest{
			ApiKey: &apikeyDeleteApiKey,
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

		apiKeyId = data.GetApiKeyId()
		if apiKeyId == "" {
			return fmt.Errorf("[ERROR] Invalid response: missing ApiKeyId")
		}

		currentStatus = data.GetStatus()
		keyName = data.GetName()
	} else {
		// --api-key-id path: lookup via DescribeApiKeys
		fmt.Printf("[STEP 1/3] Looking up API key...\n")

		listResp, err := apiClient.DescribeApiKeys(ctx, &client.DescribeApiKeysRequest{
			KeyIds: []string{apikeyDeleteApiKeyId},
		})
		if err != nil {
			printReqIDFromErr(err)
			return fmt.Errorf("[ERROR] Failed to look up API key: %w", err)
		}

		if listResp.Body == nil {
			return fmt.Errorf("[ERROR] Invalid response: missing body")
		}

		if reqID := listResp.Body.GetRequestId(); reqID != "" {
			fmt.Printf("[INFO] DescribeApiKeys Request ID: %s\n", reqID)
		}

		code := listResp.Body.GetCode()
		successPtr := listResp.Body.Success
		if (successPtr != nil && !*successPtr) || (code != "" && !isSuccessCode(code)) {
			msg := listResp.Body.GetMessage()
			return fmt.Errorf("[ERROR] Failed to look up API key: Code=%s, Message=%s", code, msg)
		}

		listData := listResp.Body.GetData()
		if listData == nil || len(listData.GetApiKeys()) == 0 {
			return fmt.Errorf("[ERROR] API key not found for the given API Key ID: %s", apikeyDeleteApiKeyId)
		}

		keyInfo := listData.GetApiKeys()[0]
		if keyInfo == nil {
			return fmt.Errorf("[ERROR] Invalid response: missing key info")
		}

		apiKeyId = keyInfo.GetKeyId()
		currentStatus = keyInfo.GetStatus()
		keyName = keyInfo.GetName()
	}

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
