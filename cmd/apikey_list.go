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

var apikeyListCmd = &cobra.Command{
	Use:   "list",
	Short: "List API keys",
	Long: `List API keys with optional filtering and pagination.

You can filter by either the user-visible API Key (--api-key, akm-xxx) or
the internal API Key ID (--api-key-id, ak-xxx). These flags are mutually exclusive.

Examples:
  # List up to 10 API keys (default)
  agentbay apikey list

  # List up to 20 API keys
  agentbay apikey list --max-results 20

  # Query a specific API key by its user-visible value (recommended)
  agentbay apikey list --api-key akm-xxxxxxxxxxxxxxxx

  # Query a specific API key by its internal ID
  agentbay apikey list --api-key-id ak-xxxxxxxxxxxxxxxx

  # Fetch the next page of results
  agentbay apikey list --next-token AAAAAV3MpHK1AP0pfERHZN5pu6mUZcGrgQ3JzaYuUyH0MyLn`,
	RunE: runApikeyList,
}

var (
	apikeyListMaxResults int32
	apikeyListApiKey     string
	apikeyListApiKeyId   string
	apikeyListNextToken  string
)

func init() {
	apikeyListCmd.Flags().Int32Var(&apikeyListMaxResults, "max-results", 10, "Number of results per query")
	apikeyListCmd.Flags().StringVar(&apikeyListApiKey, "api-key", "", "User-visible API key (akm-xxx format, recommended) to filter")
	apikeyListCmd.Flags().StringVar(&apikeyListApiKeyId, "api-key-id", "", "Internal API Key ID (ak-xxx) to filter. Prefer --api-key for normal usage")
	apikeyListCmd.Flags().StringVar(&apikeyListNextToken, "next-token", "", "Pagination token from previous query")

	ApiKeyCmd.AddCommand(apikeyListCmd)
}

func runApikeyList(cmd *cobra.Command, args []string) error {
	// Validate mutual exclusivity: --api-key and --api-key-id
	if apikeyListApiKey != "" && apikeyListApiKeyId != "" {
		return fmt.Errorf("[ERROR] --api-key and --api-key-id are mutually exclusive; please specify only one")
	}

	maxResults := apikeyListMaxResults
	apiKey := apikeyListApiKey
	apiKeyIdFlag := apikeyListApiKeyId
	nextToken := apikeyListNextToken

	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}
	apiClient := agentbay.NewClientFromConfig(cfg)
	ctx := context.Background()

	var apiKeyId string
	totalSteps := 1

	// If --api-key is provided, look up the internal KeyId first
	if apiKey != "" {
		totalSteps = 2
		fmt.Printf("[STEP 1/%d] Looking up API key...\n", totalSteps)

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

		apiKeyId = data.GetApiKeyId()
		if apiKeyId == "" {
			return fmt.Errorf("[ERROR] Invalid response: missing ApiKeyId")
		}

		fmt.Printf("  ApiKeyId: %s\n", apiKeyId)
	} else if apiKeyIdFlag != "" {
		// --api-key-id path: use the ID directly, no lookup needed
		apiKeyId = apiKeyIdFlag
	}

	// Call DescribeApiKeys
	fmt.Printf("[STEP %d/%d] Listing API keys...\n", totalSteps, totalSteps)

	req := &client.DescribeApiKeysRequest{
		MaxResults: &maxResults,
	}
	if nextToken != "" {
		req.NextToken = &nextToken
	}
	if apiKeyId != "" {
		req.KeyIds = []string{apiKeyId}
	}

	resp, err := apiClient.DescribeApiKeys(ctx, req)
	if err != nil {
		printReqIDFromErr(err)
		return fmt.Errorf("[ERROR] Failed to list API keys: %w", err)
	}

	if resp.Body == nil {
		return fmt.Errorf("[ERROR] Invalid response: missing body")
	}

	if reqID := resp.Body.GetRequestId(); reqID != "" {
		fmt.Printf("[INFO] DescribeApiKeys Request ID: %s\n", reqID)
	}

	// Success determination (Code-based SOP — new API).
	// The DescribeApiKeys API returns code="200" (HTTP status) instead of "ok",
	// so we accept both "ok" and HTTP 2xx codes as success indicators.
	code := resp.Body.GetCode()
	successPtr := resp.Body.Success
	if (successPtr != nil && !*successPtr) || (code != "" && !isSuccessCode(code)) {
		msg := resp.Body.GetMessage()
		return fmt.Errorf("[ERROR] Failed to list API keys: Code=%s, Message=%s", code, msg)
	}

	data := resp.Body.GetData()
	if data == nil || len(data.GetApiKeys()) == 0 {
		fmt.Printf("\n[EMPTY] No API keys found.\n")
		return nil
	}

	apiKeys := data.GetApiKeys()
	fmt.Printf("\n[OK] Found %d API key(s)\n\n", len(apiKeys))
	printApiKeyTable(apiKeys)

	// Print pagination hint
	if nextTokenVal := data.GetNextToken(); nextTokenVal != "" {
		fmt.Printf("\n[INFO] More results available. Use --next-token %s to fetch the next page.\n", nextTokenVal)
	}

	return nil
}

func printApiKeyTable(apiKeys []*client.DescribeApiKeysResponseBodyDataApiKey) {
	// Print header
	fmt.Printf("%s %s %s %s %s %s\n",
		padString("NAME", 20),
		padString("STATUS", 12),
		padString("CONCURRENCY", 14),
		padString("KEY ID", 25),
		padString("CREATED", 22),
		"LAST USED")
	fmt.Printf("%s %s %s %s %s %s\n",
		padString("----", 20),
		padString("------", 12),
		padString("------------", 14),
		padString("------", 25),
		padString("-------", 22),
		"----------")

	// Print each API key
	for _, key := range apiKeys {
		if key == nil {
			continue
		}

		name := key.GetName()
		status := key.GetStatus()
		keyId := key.GetKeyId()
		created := truncateDateOffset(key.GetGmtCreate())
		lastUsed := truncateDateOffset(key.GetLastUseDate())

		var concurrency string
		if key.Concurrency != nil {
			concurrency = fmt.Sprintf("%d", *key.Concurrency)
		} else {
			concurrency = "-"
		}

		fmt.Printf("%s %s %s %s %s %s\n",
			padString(truncateString(name, 20), 20),
			padString(status, 12),
			padString(concurrency, 14),
			padString(truncateString(keyId, 25), 25),
			padString(truncateString(created, 22), 22),
			truncateString(lastUsed, 22))
	}
}

// truncateDateOffset removes the timezone offset from a date string for cleaner display
func truncateDateOffset(s string) string {
	if s == "" {
		return ""
	}
	// Find the + or - timezone offset at the end
	idx := strings.LastIndex(s, "+")
	if idx == -1 {
		idx = strings.LastIndex(s, "-")
		// Make sure it's a timezone offset, not a negative number in the date
		// Timezone offsets appear after 'T' in ISO8601 format
		if idx != -1 && idx < strings.Index(s, "T") {
			return s
		}
	}
	if idx > 0 {
		return s[:idx]
	}
	return s
}

// isSuccessCode checks whether a Code field value indicates success.
// Accepts "ok" (case-insensitive) per SOP, and also HTTP 2xx status codes
// (e.g. "200") since some APIs return HTTP status codes instead of "ok".
func isSuccessCode(code string) bool {
	if strings.EqualFold(code, "ok") {
		return true
	}
	// Accept HTTP 2xx status codes (200, 201, 204, etc.)
	if len(code) == 3 && code[0] == '2' {
		return true
	}
	return false
}
