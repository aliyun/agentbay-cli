// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/agentbay/agentbay-cli/internal/agentbay"
	"github.com/agentbay/agentbay-cli/internal/client"
	"github.com/agentbay/agentbay-cli/internal/config"
)

var imageWarmupStatusCmd = &cobra.Command{
	Use:   "warmup-status",
	Short: "Show warm-up quota and image status for the current account",
	Long: `Query the warm-up status for the current account.

This command displays the session quota, image quota, and details of warm-up images
for the current authenticated account.

Examples:
  # Query warm-up status
  agentbay image warmup-status`,
	Args: cobra.NoArgs,
	RunE: runImageWarmupStatus,
}

func runImageWarmupStatus(cmd *cobra.Command, args []string) error {
	// Load configuration and check authentication
	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to load configuration: %w", err)
	}

	if !cfg.IsAuthenticated() {
		return config.ErrNotAuthenticated()
	}

	// Create API client
	apiClient := agentbay.NewClientFromConfig(cfg)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Call API
	fmt.Printf("[WARMUP-STATUS] Querying warm-up status for current account...\n")
	req := &client.DescribeWarmUpStatusOpenRequest{}
	resp, err := apiClient.DescribeWarmUpStatusOpen(ctx, req)
	if err != nil {
		if reqId := extractRequestIDFromErr(err); reqId != "" {
			fmt.Printf("[INFO] DescribeWarmUpStatusOpen Request ID: %s\n", reqId)
		}
		return fmt.Errorf("[ERROR] Failed to query warm-up status: %w", err)
	}

	// Print RequestId
	if resp != nil && resp.Body != nil {
		if reqId := resp.Body.GetRequestId(); reqId != nil && *reqId != "" {
			fmt.Printf("[INFO] DescribeWarmUpStatusOpen Request ID: %s\n", *reqId)
		}
	}

	// Validate response
	if resp.Body == nil {
		return fmt.Errorf("[ERROR] Invalid response: missing response body")
	}

	if resp.Body.GetSuccess() != nil && !*resp.Body.GetSuccess() {
		errorMsg := "unknown error"
		if resp.Body.GetMessage() != nil {
			errorMsg = *resp.Body.GetMessage()
		}
		return fmt.Errorf("[ERROR] API request failed: %s", errorMsg)
	}

	data := resp.Body.GetData()
	if data == nil {
		fmt.Println("[INFO] No warm-up data available.")
		return nil
	}

	fmt.Println()

	// Session Quota
	fmt.Println("[QUOTA] Session Quota:")
	fmt.Printf("  Max Session Limit:       %d\n", data.GetMaxSessionNumLimit())
	fmt.Printf("  Total Used Session:      %d\n", data.GetTotalUsedSessionQuota())
	fmt.Printf("  Available Session:       %d\n", data.GetAvailableSessionQuota())
	fmt.Println()

	// Image Quota
	fmt.Println("[QUOTA] Image Quota:")
	fmt.Printf("  Max Image Count:         %d\n", data.GetMaxImageCount())
	fmt.Printf("  Current Image Count:     %d\n", data.GetCurrentImageCount())
	fmt.Println()

	// Images
	images := data.GetImages()
	if len(images) == 0 {
		fmt.Println("[INFO] No warm-up images found.")
		return nil
	}

	fmt.Printf("[IMAGES] Warm-up Images (%d):\n\n", len(images))
	fmt.Printf("  %s %s %s\n",
		padString("IMAGE ID", 25),
		padString("TOTAL MAX SIZE", 18),
		padString("GROUP COUNT", 14))
	fmt.Printf("  %s %s %s\n",
		padString("--------", 25),
		padString("--------------", 18),
		padString("-----------", 14))
	for _, img := range images {
		fmt.Printf("  %s %s %s\n",
			padString(truncateString(img.GetImageId(), 25), 25),
			padString(fmt.Sprintf("%d", img.GetTotalMaxSize()), 18),
			padString(fmt.Sprintf("%d", img.GetGroupCount()), 14))
	}

	return nil
}

func extractRequestIDFromErr(err error) string {
	if err == nil {
		return ""
	}
	if e, ok := err.(*client.ErrWithRequestID); ok {
		return e.RequestID
	}
	return ""
}
