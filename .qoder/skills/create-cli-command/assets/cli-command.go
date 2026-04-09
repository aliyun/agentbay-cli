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

// {CommandGroup}Cmd is the subcommand under {Parent} for {Feature}
var {CommandGroup}Cmd = &cobra.Command{
	Use:   "{command-group}",
	Short: "Manage {feature}",
	Long:  "Configure and manage {feature}.",
}

var {SubCommand}Cmd = &cobra.Command{
	Use:   "{subcommand}",
	Short: "{Short description}",
	Long: `{Long description with examples.

Examples:
  # Basic usage
  agentbay {parent} {command-group} {subcommand} --{param1} "<value1>"
  
  # With verbose output
  agentbay {parent} {command-group} {subcommand} --{param1} "<value1>" -v`,
	RunE: run{SubCommand},
}

var {ParamVar1} string
var {ParamVar2} int32

func init() {
	{SubCommand}Cmd.Flags().StringVar(&{ParamVar1}, "{param1}", "", "{Param1 description} (required)")
	{SubCommand}Cmd.Flags().Int32Var(&{ParamVar2}, "{param2}", 0, "{Param2 description} (required)")
	{SubCommand}Cmd.MarkFlagRequired("{param1}")
	{SubCommand}Cmd.MarkFlagRequired("{param2}")
	
	{CommandGroup}Cmd.AddCommand({SubCommand}Cmd)
}

func run{SubCommand}(cmd *cobra.Command, args []string) error {
	// Validate parameters
	if {ParamVar2} < 1 {
		return fmt.Errorf("[ERROR] {param2} must be greater than or equal to 1")
	}
	
	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}
	apiClient := agentbay.NewClientFromConfig(cfg)
	ctx := context.Background()

	fmt.Printf("[STEP 1/1] <Action description>...\n")
	fmt.Printf("  {Param1}: %s\n", {ParamVar1})
	fmt.Printf("  {Param2}: %d\n", {ParamVar2})
	
	req := &client.{Action}Request{
		{Field1}: &{ParamVar1},
		{Field2}: &{ParamVar2},
	}
	
	resp, err := apiClient.{Action}(ctx, req)
	if err != nil {
		printRequestIDFromErrIfVerbose(cmd, err)
		return fmt.Errorf("[ERROR] Failed to <action>: %w", err)
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
		return fmt.Errorf("[ERROR] Failed to <action>: Code=%s, Message=%s", code, message)
	}
	
	fmt.Println()
	fmt.Printf("[SUCCESS] ✅ <Success message>!\n")
	fmt.Printf("%-*s %s\n", 14, "{Param1}:", {ParamVar1})
	fmt.Printf("%-*s %d\n", 14, "{Param2}:", {ParamVar2})
	
	return nil
}
