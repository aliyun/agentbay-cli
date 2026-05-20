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

// Output style: label width for apikey show (no emoji).
const apikeyDetailLabelW = 10

var ApiKeyCmd = &cobra.Command{
	Use:     "apikey",
	Short:   "Manage AgentBay API keys",
	Long:    "Create and manage API keys for AgentBay services.",
	GroupID: "management",
}

var apikeyCreateCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create a new API key",
	Long: `Create a new API key with the specified name.

The API key is used to authenticate requests to AgentBay services.
Each key must have a unique name.

Examples:
  # Create an API key (positional argument, recommended)
  agentbay apikey create "my-api-key"

  # Create an API key (--name flag, backward compatible)
  agentbay apikey create --name "my-api-key"

  # Create with verbose output
  agentbay apikey create "production-key" -v`,
	Args: cobra.MaximumNArgs(1),
	RunE: runApikeyCreate,
}

var apikeyCreateName string

func init() {
	apikeyCreateCmd.Flags().StringVar(&apikeyCreateName, "name", "", "API key name (can also be provided as positional argument)")

	ApiKeyCmd.AddCommand(apikeyCreateCmd)
	ApiKeyCmd.AddCommand(ApiKeyConcurrencyCmd)
}

func runApikeyCreate(cmd *cobra.Command, args []string) error {
	var name string
	if len(args) > 0 {
		name = args[0] // positional argument takes priority
	} else {
		name = apikeyCreateName // fall back to --name flag
	}
	if name == "" {
		return fmt.Errorf("[ERROR] API key name is required. Usage: agentbay apikey create <name>  or  --name <name>")
	}
	
	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}
	apiClient := agentbay.NewClientFromConfig(cfg)
	ctx := context.Background()

	fmt.Printf("[STEP 1/1] Creating API key...\n")
	
	req := &client.CreateApiKeyRequest{Name: &name}
	resp, err := apiClient.CreateApiKey(ctx, req)
	if err != nil {
		printRequestIDFromErrIfVerbose(cmd, err)
		
		// Check for specific error codes
		if resp != nil && resp.Body != nil {
			if code := resp.Body.GetCode(); code == "ApiKey.CreateError.NeedCertified" {
				return fmt.Errorf("[ERROR] Failed to create API key: account requires real-name verification")
			}
		}
		
		return fmt.Errorf("[ERROR] Failed to create API key: %w", err)
	}
	
	if resp.Body == nil {
		return fmt.Errorf("[ERROR] Invalid response: missing body")
	}
	
	verbose, _ := cmd.Flags().GetBool("verbose")
	if verbose && resp.Body.RequestId != nil && *resp.Body.RequestId != "" {
		printRequestIDIfVerbose(cmd, *resp.Body.RequestId)
	}
	
	keyId := resp.Body.GetData()
	if keyId == "" {
		return fmt.Errorf("[ERROR] Invalid response: missing KeyId")
	}
	
	fmt.Println()
	fmt.Printf("[SUCCESS] ✅ API key created successfully!\n")
	fmt.Printf("%-*s %s\n", apikeyDetailLabelW, "KeyId:", keyId)
	fmt.Printf("%-*s %s\n", apikeyDetailLabelW, "Name:", name)
	
	return nil
}
