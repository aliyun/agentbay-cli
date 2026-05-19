// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/agentbay/agentbay-cli/internal/agentbay"
	"github.com/agentbay/agentbay-cli/internal/client"
	"github.com/agentbay/agentbay-cli/internal/config"
)

// ApiKeyConcurrencyCmd is the subcommand under apikey for concurrency settings
var ApiKeyConcurrencyCmd = &cobra.Command{
	Use:   "concurrency",
	Short: "Manage API key concurrency settings",
	Long:  "Configure the concurrent session limit for API keys.",
}

var apiKeyConcurrencySetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the concurrency limit for an API key",
	Long: `Set the maximum number of concurrent sessions for an API key.

The concurrency limit controls how many sessions can run simultaneously
for the specified API key.

You can identify the API key either by its user-visible value (--api-key, akm-xxx)
or by its internal ID (--api-key-id, ak-xxx). Using --api-key is recommended.

Examples:
  # Set concurrency using the user-visible API Key (recommended)
  agentbay apikey concurrency set --api-key "akm-xxx" --concurrency 10

  # Set concurrency using the internal API Key ID
  agentbay apikey concurrency set --api-key-id "ak-xxx" --concurrency 10

  # Set with verbose output
  agentbay apikey concurrency set --api-key "akm-xxx" --concurrency 5 -v`,
	RunE: runApiKeyConcurrencySet,
}

var apiKeyConcurrencySetApiKeyID string
var apiKeyConcurrencySetApiKey string
var apiKeyConcurrencySetValue int32

func init() {
	apiKeyConcurrencySetCmd.Flags().StringVar(&apiKeyConcurrencySetApiKeyID, "api-key-id", "", "Internal API Key ID (ak-xxx). Prefer --api-key for normal usage")
	apiKeyConcurrencySetCmd.Flags().StringVar(&apiKeyConcurrencySetApiKey, "api-key", "", "User-visible API Key (akm-xxx format, recommended)")
	apiKeyConcurrencySetCmd.Flags().Int32Var(&apiKeyConcurrencySetValue, "concurrency", 0, "Maximum concurrent sessions (required, must be >= 1)")
	apiKeyConcurrencySetCmd.MarkFlagRequired("concurrency")

	ApiKeyConcurrencyCmd.AddCommand(apiKeyConcurrencySetCmd)
}

func runApiKeyConcurrencySet(cmd *cobra.Command, args []string) error {
	// Validate mutual exclusivity: --api-key and --api-key-id
	if apiKeyConcurrencySetApiKey == "" && apiKeyConcurrencySetApiKeyID == "" {
		return fmt.Errorf("[ERROR] Either --api-key or --api-key-id must be specified. Using --api-key is recommended")
	}
	if apiKeyConcurrencySetApiKey != "" && apiKeyConcurrencySetApiKeyID != "" {
		return fmt.Errorf("[ERROR] --api-key and --api-key-id are mutually exclusive; please specify only one")
	}

	// Validate concurrency
	if apiKeyConcurrencySetValue < 1 {
		return fmt.Errorf("[ERROR] Concurrency must be greater than or equal to 1")
	}

	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}
	apiClient := agentbay.NewClientFromConfig(cfg)
	ctx := context.Background()

	var apiKeyId string

	// Resolve apiKeyId based on which flag was provided
	if apiKeyConcurrencySetApiKey != "" {
		// Two-step flow: lookup API key first
		fmt.Printf("[STEP 1/2] Looking up API key...\n")

		descResp, err := apiClient.DescribeMcpApiKey(ctx, &client.DescribeMcpApiKeyRequest{
			ApiKey: &apiKeyConcurrencySetApiKey,
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

		keyName := data.GetName()

		fmt.Printf("  ApiKeyId: %s\n", apiKeyId)
		if keyName != "" {
			fmt.Printf("  Name:     %s\n", keyName)
		}

		fmt.Printf("[STEP 2/2] Setting concurrency for API key...\n")
	} else {
		// Single-step flow: use --api-key-id directly
		apiKeyId = apiKeyConcurrencySetApiKeyID
		fmt.Printf("[STEP 1/1] Setting concurrency for API key...\n")
	}

	fmt.Printf("  ApiKeyId:    %s\n", apiKeyId)
	fmt.Printf("  Concurrency: %d\n", apiKeyConcurrencySetValue)

	req := &client.ModifyMcpApiKeyConfigRequest{
		ApiKeyId:    &apiKeyId,
		Concurrency: &apiKeyConcurrencySetValue,
	}

	resp, err := apiClient.ModifyMcpApiKeyConfig(ctx, req)
	if err != nil {
		printReqIDFromErr(err)
		return fmt.Errorf("[ERROR] Failed to set concurrency: %w", err)
	}

	if resp.Body == nil {
		return fmt.Errorf("[ERROR] Invalid response: missing body")
	}

	if resp.Body.RequestId != nil && *resp.Body.RequestId != "" {
		fmt.Printf("[INFO] ModifyMcpApiKeyConfig Request ID: %s\n", *resp.Body.RequestId)
	}

	if !resp.Body.GetSuccess() {
		code := resp.Body.GetCode()
		message := ""
		if resp.Body.Message != nil {
			message = *resp.Body.Message
		}
		return fmt.Errorf("[ERROR] Failed to set concurrency: Code=%s, Message=%s", code, message)
	}

	fmt.Println()
	fmt.Printf("[SUCCESS] Concurrency updated successfully!\n")
	fmt.Printf("%-*s %s\n", 14, "ApiKeyId:", apiKeyId)
	fmt.Printf("%-*s %d\n", 14, "Concurrency:", apiKeyConcurrencySetValue)
	if apiKeyConcurrencySetApiKey != "" {
		fmt.Printf("%-*s %s\n", 14, "ApiKey:", apiKeyConcurrencySetApiKey)
	}

	return nil
}
