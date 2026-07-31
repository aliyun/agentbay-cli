// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/agentbay/agentbay-cli/internal/agentbay"
	"github.com/agentbay/agentbay-cli/internal/client"
	"github.com/agentbay/agentbay-cli/internal/config"
	"github.com/alibabacloud-go/tea/dara"
)

var imageSetPreOpenCmd = &cobra.Command{
	Use:   "set-pre-open",
	Short: "Set the pre-open (reserveMinAmount) for an activated ACS image",
	Long: `Set the pre-open (reserveMinAmount) for an activated ACS image.

This command directly sets the reserveMinAmount (pre-open value) for all
resource groups (default + hidden) of the specified ACS image. The image
must be a User type image in activated state (RESOURCE_PUBLISHED) and
using ACS (advanced network).

Note: Only ACS images support this feature. If the image is not an ACS
image, the server will return an error.

This operation requires whitelist approval. If your account is not
whitelisted, the server will return an error.

Examples:
  # Set pre-open to 10
  agentbay image set-pre-open --image-id imgc-xxxxxxxxxxxxxx --pre-open 10

  # Set pre-open to 1 (minimum pre-open instances)
  agentbay image set-pre-open --image-id imgc-xxxxxxxxxxxxxx --pre-open 1`,
	RunE: runImageSetPreOpen,
}

func init() {
	imageSetPreOpenCmd.Flags().String("image-id", "", "Image ID (required)")
	imageSetPreOpenCmd.Flags().Int32("pre-open", 0, "Pre-open value (reserveMinAmount, required, must be >= 1)")

	imageSetPreOpenCmd.MarkFlagRequired("image-id")
	imageSetPreOpenCmd.MarkFlagRequired("pre-open")
}

func runImageSetPreOpen(cmd *cobra.Command, args []string) error {
	imageId, _ := cmd.Flags().GetString("image-id")
	preOpen, _ := cmd.Flags().GetInt32("pre-open")

	if preOpen < 1 {
		return fmt.Errorf("--pre-open must be greater than or equal to 1")
	}

	fmt.Printf("[SET-PRE-OPEN] Setting pre-open to %d for image '%s'...\n", preOpen, imageId)

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
		return fmt.Errorf("only User images support set-pre-open (current type: %s)", imageInfo.ImageType)
	}

	// Must be activated (RESOURCE_PUBLISHED)
	if !IsActivated(imageInfo.ResourceStatus) {
		return fmt.Errorf("image must be in activated state to set pre-open (current status: %s)", TranslateImageResourceStatus(imageInfo.ResourceStatus))
	}

	// Step 2: Call UpdateImageReserveMinAmount
	fmt.Printf("Setting pre-open to %d...\n", preOpen)

	apiCtx, apiCancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer apiCancel()

	request := &client.UpdateImageReserveMinAmountRequest{}
	request.SetImageId(imageId)
	request.SetReserveMinAmount(preOpen)

	resp, err := apiClient.UpdateImageReserveMinAmount(apiCtx, request)
	if err != nil {
		return fmt.Errorf("failed to set pre-open: %w", err)
	}

	// Print RequestId regardless of success/failure
	if resp != nil && resp.Body != nil {
		requestId := resp.Body.GetRequestId()
		if requestId != "" {
			fmt.Printf("[INFO] UpdateImageReserveMinAmount Request ID: %s\n", requestId)
		}

		code := resp.Body.GetCode()
		successPtr := resp.Body.Success
		if (successPtr != nil && !*successPtr) || (code != "" && !strings.EqualFold(code, "ok")) {
			message := ""
			if resp.Body.Message != nil {
				message = dara.StringValue(resp.Body.Message)
			}
			return fmt.Errorf("server returned error: code=%s, message=%s", code, message)
		}
	}

	fmt.Printf("[OK] Pre-open has been set to %d for image '%s'.\n", preOpen, imageId)
	fmt.Printf("[INFO] Resource expansion/shrinkage is processing asynchronously. Use 'agentbay image warmup-status' to check the actual status.\n")
	return nil
}
