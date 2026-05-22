// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/agentbay/agentbay-cli/internal/agentbay"
	"github.com/agentbay/agentbay-cli/internal/client"
	"github.com/agentbay/agentbay-cli/internal/config"
)

var apikeyDescribeKeyContentApiKeyId string

var apikeyDescribeKeyContentCmd = &cobra.Command{
	Use:   "describe-key-content",
	Short: "Retrieve the plaintext API key by API key ID",
	Long: `Retrieve the plaintext API key (akm-xxx format) for a given API key ID (ak-xxx).

This command calls the DescribeKeyContent API and returns the user-visible API key
associated with the specified internal API key ID.

Examples:
  # Retrieve the plaintext API key for a given API key ID
  agentbay apikey describe-key-content --api-key-id ak-xxxxxxxxxxxxxxxx`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runApikeyDescribeKeyContent(cmd)
	},
}

func init() {
	apikeyDescribeKeyContentCmd.Flags().StringVar(&apikeyDescribeKeyContentApiKeyId, "api-key-id", "", "Internal API key ID (ak-xxx format)")
	_ = apikeyDescribeKeyContentCmd.MarkFlagRequired("api-key-id")

	ApiKeyCmd.AddCommand(apikeyDescribeKeyContentCmd)
}

func runApikeyDescribeKeyContent(cmd *cobra.Command) error {
	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}
	apiClient := agentbay.NewClientFromConfig(cfg)
	ctx := context.Background()

	fmt.Printf("[STEP 1/1] Fetching API key content...\n")

	req := &client.DescribeKeyContentRequest{
		KeyId: &apikeyDescribeKeyContentApiKeyId,
	}
	resp, err := apiClient.DescribeKeyContent(ctx, req)
	if err != nil {
		printReqIDFromErr(err)
		return fmt.Errorf("[ERROR] Failed to describe key content: %w", err)
	}

	if resp.Body == nil {
		return fmt.Errorf("[ERROR] Invalid response: missing body")
	}

	// Print RequestID first for troubleshooting
	if reqID := resp.Body.GetRequestId(); reqID != "" {
		fmt.Printf("[INFO] DescribeKeyContent Request ID: %s\n", reqID)
	}

	// Success determination SOP for new interfaces:
	// - Use Code field as primary: "ok" (case-insensitive) or "200" = success
	// - Success field: only explicit false = failure; nil = treat as success
	code := resp.Body.GetCode()
	successPtr := resp.Body.Success
	isOk := code == "" || strings.EqualFold(code, "ok") || code == "200"
	if (successPtr != nil && !*successPtr) || !isOk {
		msg := resp.Body.GetMessage()
		return fmt.Errorf("[ERROR] Failed to describe key content: Code=%s, Message=%s", code, msg)
	}

	if resp.Body.Data == nil {
		return fmt.Errorf("[ERROR] Invalid response: missing data")
	}

	apiKey := resp.Body.Data.GetApiKey()
	if apiKey == "" {
		return fmt.Errorf("[ERROR] Invalid response: missing ApiKey in data")
	}

	fmt.Println()
	fmt.Printf("[SUCCESS] API key content retrieved successfully!\n")
	fmt.Printf("%-*s %s\n", apikeyDetailLabelW, "ApiKey:", apiKey)
	fmt.Printf("%-*s %s\n", apikeyDetailLabelW, "ApiKeyId:", apikeyDescribeKeyContentApiKeyId)

	return nil
}
