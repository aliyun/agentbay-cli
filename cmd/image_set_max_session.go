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
	"github.com/alibabacloud-go/tea/dara"
)

var imageSetMaxSessionCmd = &cobra.Command{
	Use:   "set-max-session",
	Short: "Set the maximum concurrent session count for an activated User image",
	Long: `Set the maximum concurrent session count for an activated User image.

This command configures the maximum number of concurrent sessions allowed for
the specified image. The image must be a User type image in activated state
(RESOURCE_PUBLISHED) and using advanced network.

Note: Only images with advanced network support this feature. If the image
uses default network, the server will return an error.

After setting, the command will poll until the resource group is ready
(typically around 5 minutes).

Examples:
  # Set max session count to 10
  agentbay image set-max-session --image-id imgc-xxxxxxxxxxxxxx --max-session-num 10

  # Set max session count to 5
  agentbay image set-max-session --image-id imgc-xxxxxxxxxxxxxx --max-session-num 5`,
	RunE: runImageSetMaxSession,
}

func init() {
	imageSetMaxSessionCmd.Flags().String("image-id", "", "Image ID (required)")
	imageSetMaxSessionCmd.Flags().Int32("max-session-num", 0, "Maximum concurrent session count (required, must be >= 1)")

	imageSetMaxSessionCmd.MarkFlagRequired("image-id")
	imageSetMaxSessionCmd.MarkFlagRequired("max-session-num")
}

func runImageSetMaxSession(cmd *cobra.Command, args []string) error {
	imageId, _ := cmd.Flags().GetString("image-id")
	maxSessionNum, _ := cmd.Flags().GetInt32("max-session-num")

	if maxSessionNum < 1 {
		return fmt.Errorf("--max-session-num must be greater than or equal to 1")
	}

	fmt.Printf("[SET-MAX-SESSION] Setting max session count to %d for image '%s'...\n", maxSessionNum, imageId)

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

	// Step 1: Validate image status
	statusCtx, statusCancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer statusCancel()

	fmt.Printf("Checking current image status...")
	imageInfo, err := GetImageInfo(statusCtx, apiClient, imageId)
	if err != nil {
		fmt.Printf(" Failed.\n")
		return fmt.Errorf("failed to get image info: %w", err)
	}
	fmt.Printf(" Done.\n")
	if imageInfo.RequestId != "" {
		fmt.Printf("[INFO] GetMcpImageInfo Request ID: %s\n", imageInfo.RequestId)
	}
	fmt.Printf("[INFO] Image Type: %s\n", imageInfo.ImageType)
	fmt.Printf("[INFO] Current Status: %s\n", TranslateImageResourceStatus(imageInfo.ResourceStatus))

	// Must be User image
	if !IsUserImage(imageInfo.ImageType) {
		return fmt.Errorf("only User images support set-max-session (current type: %s)", imageInfo.ImageType)
	}

	// Must be activated (RESOURCE_PUBLISHED)
	if !IsActivated(imageInfo.ResourceStatus) {
		return fmt.Errorf("image must be in activated state to set max session (current status: %s)", TranslateImageResourceStatus(imageInfo.ResourceStatus))
	}

	// Step 2: Call BatchCreateHideResourceGroupsWithMaxSession
	fmt.Printf("Setting max session count to %d...\n", maxSessionNum)

	apiCtx, apiCancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer apiCancel()

	request := &client.BatchCreateHideResourceGroupsWithMaxSessionRequest{}
	request.SetImageId(imageId)
	request.SetMaxSessionNum(maxSessionNum)

	resp, err := apiClient.BatchCreateHideResourceGroupsWithMaxSession(apiCtx, request)
	if err != nil {
		return fmt.Errorf("failed to set max session: %w", err)
	}

	// Print RequestId regardless of success/failure
	if resp != nil && resp.Body != nil {
		requestId := resp.Body.GetRequestId()
		if requestId != "" {
			fmt.Printf("[INFO] BatchCreateHideResourceGroupsWithMaxSession Request ID: %s\n", requestId)
		}

		if !resp.Body.GetSuccess() {
			code := resp.Body.GetCode()
			message := ""
			if resp.Body.Message != nil {
				message = dara.StringValue(resp.Body.Message)
			}
			return fmt.Errorf("server returned error: code=%s, message=%s", code, message)
		}
	}

	fmt.Printf("[OK] Max session count set successfully. Waiting for resource group to be ready...\n")

	// Step 3: Poll for ResourceGroupReady
	pollingCtx := context.Background()
	pollingConfig := DefaultSetMaxSessionPollingConfig()

	if err := PollForResourceGroupReady(pollingCtx, apiClient, imageId, pollingConfig); err != nil {
		return fmt.Errorf("set-max-session polling failed: %w", err)
	}

	fmt.Printf("[DONE] Image '%s' max session count has been set to %d.\n", imageId, maxSessionNum)
	return nil
}
