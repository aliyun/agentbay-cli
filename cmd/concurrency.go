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

Examples:
  # Set concurrency to 10 for an API key
  agentbay apikey concurrency set --api-key-id "ak-xxx" --concurrency 10
  
  # Set with verbose output
  agentbay apikey concurrency set --api-key-id "ak-xxx" --concurrency 5 -v`,
	RunE: runApiKeyConcurrencySet,
}

var apiKeyConcurrencySetApiKeyID string
var apiKeyConcurrencySetValue int32

func init() {
	apiKeyConcurrencySetCmd.Flags().StringVar(&apiKeyConcurrencySetApiKeyID, "api-key-id", "", "API Key ID (required)")
	apiKeyConcurrencySetCmd.Flags().Int32Var(&apiKeyConcurrencySetValue, "concurrency", 0, "Maximum concurrent sessions (required, must be >= 1)")
	apiKeyConcurrencySetCmd.MarkFlagRequired("api-key-id")
	apiKeyConcurrencySetCmd.MarkFlagRequired("concurrency")
	
	ApiKeyConcurrencyCmd.AddCommand(apiKeyConcurrencySetCmd)
}

func runApiKeyConcurrencySet(cmd *cobra.Command, args []string) error {
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

	fmt.Printf("[STEP 1/1] Setting concurrency for API key...\n")
	fmt.Printf("  ApiKeyId:    %s\n", apiKeyConcurrencySetApiKeyID)
	fmt.Printf("  Concurrency: %d\n", apiKeyConcurrencySetValue)
	
	req := &client.ModifyMcpApiKeyConfigRequest{
		ApiKeyId:    &apiKeyConcurrencySetApiKeyID,
		Concurrency: &apiKeyConcurrencySetValue,
	}
	
	resp, err := apiClient.ModifyMcpApiKeyConfig(ctx, req)
	if err != nil {
		printRequestIDFromErrIfVerbose(cmd, err)
		return fmt.Errorf("[ERROR] Failed to set concurrency: %w", err)
	}
	
	if resp.Body == nil {
		return fmt.Errorf("[ERROR] Invalid response: missing body")
	}
	
	verbose, _ := cmd.Flags().GetBool("verbose")
	if verbose && resp.Body.RequestId != nil && *resp.Body.RequestId != "" {
		printRequestIDIfVerbose(cmd, *resp.Body.RequestId)
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
	fmt.Printf("[SUCCESS] ✅ Concurrency updated successfully!\n")
	fmt.Printf("%-*s %s\n", 14, "ApiKeyId:", apiKeyConcurrencySetApiKeyID)
	fmt.Printf("%-*s %d\n", 14, "Concurrency:", apiKeyConcurrencySetValue)
	
	return nil
}
