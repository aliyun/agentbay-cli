// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/agentbay/agentbay-cli/internal/agentbay"
	"github.com/agentbay/agentbay-cli/internal/client"
	"github.com/agentbay/agentbay-cli/internal/config"
)

var imageDescribePreOpenCmd = &cobra.Command{
	Use:   "describe-pre-open",
	Short: "Query pre-open (reserveMinAmount) values for ACS images",
	Long: `Query the pre-open (reserveMinAmount) configuration values for ACS images.

This command calls the DescribeImageReserveMinAmount API to read DB configuration
values directly, covering both default and hidden resource groups. Unlike
warmup-status, this command does not depend on set-max-session and returns
per-resource-group details (reserve, max, type, status).

Supports batch query with multiple --image-id flags, pagination via
--next-token / --max-results, and JSON output via --output json.

Examples:
  # Query a single image
  agentbay image describe-pre-open --image-id imgc-xxxxxxxxxxxxxx

  # Query multiple images
  agentbay image describe-pre-open --image-id imgc-aaa --image-id imgc-bbb

  # Query all images (first page, default page size 20)
  agentbay image describe-pre-open

  # Get the next page of results
  agentbay image describe-pre-open --next-token eyJpZCI6...

  # Use a larger page size
  agentbay image describe-pre-open --max-results 50

  # JSON output
  agentbay image describe-pre-open --output json`,
	Args: cobra.NoArgs,
	RunE: runImageDescribePreOpen,
}

func init() {
	imageDescribePreOpenCmd.Flags().StringArray("image-id", nil, "Image ID (optional, repeatable for batch query)")
	imageDescribePreOpenCmd.Flags().String("next-token", "", "Pagination token from a previous response")
	imageDescribePreOpenCmd.Flags().Int32("max-results", 20, "Page size (number of images per page, default 20, max 500)")
	imageDescribePreOpenCmd.Flags().StringP("output", "o", "", `Output format. Use "json" for machine-readable output (e.g. for AI/scripts)`)
}

func runImageDescribePreOpen(cmd *cobra.Command, args []string) error {
	imageIds, _ := cmd.Flags().GetStringArray("image-id")
	nextToken, _ := cmd.Flags().GetString("next-token")
	maxResults, _ := cmd.Flags().GetInt32("max-results")

	if maxResults < 1 {
		return fmt.Errorf("--max-results must be greater than or equal to 1")
	}
	if maxResults > 500 {
		return fmt.Errorf("--max-results must be less than or equal to 500")
	}

	// Mirror backend MAX_DESCRIBE_IMAGE_RESERVE_MIN_AMOUNT_IMAGE_IDS = 100
	if len(imageIds) > 100 {
		return fmt.Errorf("--image-id supports at most 100 image IDs")
	}

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

	// Build request
	req := &client.DescribeImageReserveMinAmountRequest{}
	if len(imageIds) > 0 {
		req.SetImageIds(imageIds)
	}
	if nextToken != "" {
		req.SetNextToken(nextToken)
	}
	req.SetMaxResults(maxResults)

	// Call API
	fmt.Printf("[DESCRIBE-PRE-OPEN] Querying pre-open values")
	if len(imageIds) > 0 {
		fmt.Printf(" for images: %s", strings.Join(imageIds, ", "))
	} else {
		fmt.Printf(" for all images")
	}
	if nextToken != "" {
		fmt.Printf(" (page token: %s)", nextToken)
	}
	fmt.Println("...")

	var resp *client.DescribeImageReserveMinAmountResponse
	err = withTransientRetry(ctx, client.DefaultRetryConfig(), "DescribeImageReserveMinAmount", func() error {
		var e error
		resp, e = apiClient.DescribeImageReserveMinAmount(ctx, req)
		return e
	})
	if err != nil {
		if reqId := extractRequestIDFromErr(err); reqId != "" {
			fmt.Printf("[INFO] DescribeImageReserveMinAmount Request ID: %s\n", reqId)
		}
		return fmt.Errorf("[ERROR] Failed to query pre-open values: %w", err)
	}

	// Print RequestId
	if resp != nil && resp.Body != nil {
		if reqId := resp.Body.GetRequestId(); reqId != "" {
			fmt.Printf("[INFO] DescribeImageReserveMinAmount Request ID: %s\n", reqId)
		}
	}

	// Validate response
	if resp == nil || resp.Body == nil {
		return fmt.Errorf("[ERROR] Invalid response: missing response body")
	}

	code := resp.Body.GetCode()
	if (resp.Body.Success != nil && !*resp.Body.Success) || (code != "" && !strings.EqualFold(code, "ok")) {
		message := ""
		if resp.Body.Message != nil {
			message = *resp.Body.Message
		}
		return fmt.Errorf("[ERROR] API request failed: code=%s, message=%s", code, message)
	}

	data := resp.Body.GetData()
	if data == nil {
		fmt.Println("[INFO] No pre-open data available.")
		return nil
	}

	images := data.GetImages()
	if len(images) == 0 {
		fmt.Println("No images found.")
		return nil
	}

	fmt.Println()

	outputFmt, _ := cmd.Flags().GetString("output")
	if strings.EqualFold(outputFmt, "json") {
		type resourceGroupJSON struct {
			ResourceGroupId    string `json:"resourceGroupId"`
			AppInstanceGroupId string `json:"appInstanceGroupId"`
			ReserveMinAmount   int32  `json:"reserveMinAmount"`
			MaxAmount          int32  `json:"maxAmount"`
			ResourceGroupType  string `json:"resourceGroupType"`
			Status             string `json:"status"`
		}
		type imageJSON struct {
			ImageId        string              `json:"imageId"`
			GroupCount     int                 `json:"groupCount"`
			ResourceGroups []resourceGroupJSON `json:"resourceGroups"`
		}
		type outputJSON struct {
			TotalCount int         `json:"totalCount"`
			NextToken  string      `json:"nextToken,omitempty"`
			Images     []imageJSON `json:"images"`
		}
		out := outputJSON{TotalCount: len(images)}
		if nt := data.GetNextToken(); nt != "" {
			out.NextToken = nt
		}
		for _, img := range images {
			if img == nil {
				continue
			}
			entry := imageJSON{
				ImageId:    img.GetImageId(),
				GroupCount: len(img.GetResourceGroups()),
			}
			for _, rg := range img.GetResourceGroups() {
				if rg == nil {
					continue
				}
				entry.ResourceGroups = append(entry.ResourceGroups, resourceGroupJSON{
					ResourceGroupId:    rg.GetResourceGroupId(),
					AppInstanceGroupId: rg.GetAppInstanceGroupId(),
					ReserveMinAmount:   rg.GetReserveMinAmount(),
					MaxAmount:          rg.GetMaxAmount(),
					ResourceGroupType:  rg.GetResourceGroupType(),
					Status:             rg.GetStatus(),
				})
			}
			if entry.ResourceGroups == nil {
				entry.ResourceGroups = []resourceGroupJSON{}
			}
			out.Images = append(out.Images, entry)
		}
		if out.Images == nil {
			out.Images = []imageJSON{}
		}
		b, jerr := json.MarshalIndent(out, "", "  ")
		if jerr != nil {
			return fmt.Errorf("json marshal: %w", jerr)
		}
		fmt.Println(string(b))
		return nil
	}

	// Print resource group details for each image
	for _, img := range images {
		groups := img.GetResourceGroups()
		if len(groups) == 0 {
			continue
		}

		fmt.Println()
		fmt.Printf("Resource Group Details for Image: %s (Total: %d)\n", img.GetImageId(), len(groups))
		fmt.Printf("  %s %s %s %s %s %s\n",
			padString("Type", 10),
			padString("Pool ID", 50),
			padString("Reserve", 10),
			padString("Max", 8),
			padString("Group ID", 25),
			padString("Status", 20))
		fmt.Printf("  %s %s %s %s %s %s\n",
			padString("----", 10),
			padString("-------", 50),
			padString("-------", 10),
			padString("---", 8),
			padString("--------", 25),
			padString("------", 20))
		for _, rg := range groups {
			fmt.Printf("  %s %s %s %s %s %s\n",
				padString(truncateString(rg.GetResourceGroupType(), 10), 10),
				padString(truncateString(rg.GetAppInstanceGroupId(), 50), 50),
				func() string {
					if rg.ReserveMinAmount == nil {
						return padString("-", 10)
					}
					return padString(fmt.Sprintf("%d", rg.GetReserveMinAmount()), 10)
				}(),
				padString(fmt.Sprintf("%d", rg.GetMaxAmount()), 8),
				padString(truncateString(rg.GetResourceGroupId(), 25), 25),
				padString(truncateString(rg.GetStatus(), 20), 20))
		}
	}

	// Print pagination info
	nextTokenStr := data.GetNextToken()
	if nextTokenStr != "" {
		fmt.Println()
		fmt.Printf("(NextToken: %s, use --next-token to get the next page)\n", nextTokenStr)
	}

	return nil
}
