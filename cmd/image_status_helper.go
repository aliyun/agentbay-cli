// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/agentbay/agentbay-cli/internal/agentbay"
	"github.com/agentbay/agentbay-cli/internal/client"
	"github.com/alibabacloud-go/tea/dara"
	log "github.com/sirupsen/logrus"
)

// ImageResourceStatus represents the possible states of an image resource
type ImageResourceStatus string

const (
	StatusImageCreating     ImageResourceStatus = "IMAGE_CREATING"
	StatusImageCreateFailed ImageResourceStatus = "IMAGE_CREATE_FAILED"
	StatusImageAvailable    ImageResourceStatus = "IMAGE_AVAILABLE"
	StatusResourceDeploying ImageResourceStatus = "RESOURCE_DEPLOYING"
	StatusResourcePublished ImageResourceStatus = "RESOURCE_PUBLISHED"
	StatusResourceDeleting  ImageResourceStatus = "RESOURCE_DELETING"
	StatusResourceFailed    ImageResourceStatus = "RESOURCE_FAILED"
	StatusResourceCeased    ImageResourceStatus = "RESOURCE_CEASED"
)

// TranslateImageResourceStatus translates the raw status to a human-readable format
func TranslateImageResourceStatus(status string) string {
	switch ImageResourceStatus(status) {
	case StatusImageCreating:
		return "Creating"
	case StatusImageCreateFailed:
		return "Create Failed"
	case StatusImageAvailable:
		return "Available (Deactivated)"
	case StatusResourceDeploying:
		return "Activating"
	case StatusResourcePublished:
		return "Activated"
	case StatusResourceDeleting:
		return "Deactivating"
	case StatusResourceFailed:
		return "Activation Failed"
	case StatusResourceCeased:
		return "Ceased"
	default:
		if status == "" {
			return "Unknown"
		}
		return status
	}
}

// IsActivated checks if the image is in activated state
func IsActivated(status string) bool {
	return ImageResourceStatus(status) == StatusResourcePublished
}

// IsDeactivated checks if the image is in deactivated state
func IsDeactivated(status string) bool {
	return ImageResourceStatus(status) == StatusImageAvailable
}

// IsActivating checks if the image is currently being activated
func IsActivating(status string) bool {
	return ImageResourceStatus(status) == StatusResourceDeploying
}

// IsDeactivating checks if the image is currently being deactivated
func IsDeactivating(status string) bool {
	return ImageResourceStatus(status) == StatusResourceDeleting
}

// IsFailed checks if the operation has failed
func IsFailed(status string) bool {
	s := ImageResourceStatus(status)
	return s == StatusImageCreateFailed || s == StatusResourceFailed
}

// IsTerminalState checks if the status is a final state (success or failure)
func IsTerminalState(status string) bool {
	s := ImageResourceStatus(status)
	return s == StatusResourcePublished ||
		s == StatusImageAvailable ||
		s == StatusImageCreateFailed ||
		s == StatusResourceFailed ||
		s == StatusResourceCeased
}

// IsAuthenticationError checks if the error is an authentication error
// This helps distinguish between "not authenticated" and "image not found" errors
func IsAuthenticationError(err error) bool {
	if err == nil {
		return false
	}

	errMsg := strings.ToLower(err.Error())

	// Check for various authentication-related error indicators
	authErrorPatterns := []string{
		"401",
		"403",
		"unauthorized",
		"authentication",
		"authentication failed",
		"not authenticated",
		"token",
		"expired",
		"invalid token",
		"access denied",
		"invalidaccesskeyid",
		"invalidsecretkey",
		"signaturedoesnotmatch",
	}

	for _, pattern := range authErrorPatterns {
		if strings.Contains(errMsg, strings.ToLower(pattern)) {
			log.Debugf("[DEBUG] Detected authentication error pattern '%s' in: %s", pattern, err.Error())
			return true
		}
	}

	return false
}

// PollingConfig defines the configuration for status polling
type PollingConfig struct {
	MaxAttempts     int           // Maximum number of polling attempts
	InitialInterval time.Duration // Initial polling interval
	MaxInterval     time.Duration // Maximum polling interval
	Timeout         time.Duration // Overall timeout for the polling operation
}

// DefaultActivatePollingConfig returns the default polling configuration for activation
func DefaultActivatePollingConfig() PollingConfig {
	return PollingConfig{
		MaxAttempts:     60,               // 60 attempts
		InitialInterval: 5 * time.Second,  // Start with 5 seconds
		MaxInterval:     30 * time.Second, // Max 30 seconds between polls
		Timeout:         30 * time.Minute, // 30 minutes total timeout
	}
}

// DefaultDeactivatePollingConfig returns the default polling configuration for deactivation
func DefaultDeactivatePollingConfig() PollingConfig {
	return PollingConfig{
		MaxAttempts:     40,               // 40 attempts
		InitialInterval: 5 * time.Second,  // Start with 5 seconds
		MaxInterval:     20 * time.Second, // Max 20 seconds between polls
		Timeout:         20 * time.Minute, // 20 minutes total timeout
	}
}

// ImageInfo contains image status and type information
type ImageInfo struct {
	ResourceStatus string // IMAGE_AVAILABLE, RESOURCE_PUBLISHED, etc.
	ImageType      string // User or System
}

// GetImageInfo retrieves the current status and type for the given image ID
func GetImageInfo(ctx context.Context, apiClient agentbay.Client, imageId string) (*ImageInfo, error) {
	log.Debugf("[DEBUG] GetImageInfo: Querying info for image %s", imageId)

	request := &client.GetMcpImageInfoRequest{}
	request.SetImageId(imageId)

	resp, err := apiClient.GetMcpImageInfo(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("failed to get image info: %w", err)
	}

	if resp == nil || resp.Body == nil {
		return nil, fmt.Errorf("invalid response from GetMcpImageInfo")
	}

	if !dara.BoolValue(resp.Body.Success) {
		code := dara.StringValue(resp.Body.Code)
		message := dara.StringValue(resp.Body.Message)
		return nil, fmt.Errorf("GetMcpImageInfo failed: %s - %s", code, message)
	}

	if resp.Body.Data == nil {
		return nil, fmt.Errorf("no data in GetMcpImageInfo response")
	}

	info := &ImageInfo{}

	// Try to get ImageResourceStatus from response headers (stored during XML parsing)
	if resp.Headers != nil {
		if resourceStatus, ok := resp.Headers["X-Image-Resource-Status"]; ok && resourceStatus != nil {
			info.ResourceStatus = dara.StringValue(resourceStatus)
			log.Debugf("[DEBUG] Got ImageResourceStatus from header: %s", info.ResourceStatus)
		}

		if imageType, ok := resp.Headers["X-Image-Type"]; ok && imageType != nil {
			info.ImageType = dara.StringValue(imageType)
			log.Debugf("[DEBUG] Got ImageType from header: %s", info.ImageType)
		}
	}

	// Fallback to ImageInfo.Status if ImageResourceStatus not available
	if info.ResourceStatus == "" && resp.Body.Data.ImageInfo != nil && resp.Body.Data.ImageInfo.Status != nil {
		imageInfoStatus := dara.StringValue(resp.Body.Data.ImageInfo.Status)
		log.Debugf("[DEBUG] Using ImageInfo.Status as fallback: %s", imageInfoStatus)
		info.ResourceStatus = imageInfoStatus
	}

	if info.ResourceStatus == "" {
		return nil, fmt.Errorf("no status information available in response")
	}

	log.Debugf("[DEBUG] GetImageInfo: Returning status=%s, type=%s", info.ResourceStatus, info.ImageType)
	return info, nil
}

// GetImageResourceStatus retrieves the current ImageResourceStatus for the given image ID
// Deprecated: Use GetImageInfo instead for more complete information
func GetImageResourceStatus(ctx context.Context, apiClient agentbay.Client, imageId string) (string, error) {
	info, err := GetImageInfo(ctx, apiClient, imageId)
	if err != nil {
		return "", err
	}
	return info.ResourceStatus, nil
}

// GetResourceGroupIdForImage fetches ResourceGroupId for the given image from ListMcpImages (user images).
// Returns empty string if not found or if the image has no ResourceGroupId.
func GetResourceGroupIdForImage(ctx context.Context, apiClient agentbay.Client, imageId string) (string, error) {
	req := &client.ListMcpImagesRequest{}
	userImageType := "User"
	req.ImageType = &userImageType

	// Use a larger page size to reduce pagination; we only need to find one image
	pageSize := int32(100)
	req.PageSize = &pageSize
	pageStart := int32(0)
	req.PageStart = &pageStart

	for {
		resp, err := apiClient.ListMcpImages(ctx, req)
		if err != nil {
			return "", fmt.Errorf("failed to list images: %w", err)
		}
		if resp == nil || resp.Body == nil || resp.Body.Data == nil {
			return "", nil
		}

		for _, img := range resp.Body.Data {
			if img == nil {
				continue
			}
			imgId := ""
			if img.ImageId != nil {
				imgId = *img.ImageId
			}
			if imgId != imageId {
				continue
			}
			rgInfo := img.GetImageResourceGroupInfo()
			if rgInfo != nil && rgInfo.ResourceGroupId != nil && *rgInfo.ResourceGroupId != "" {
				return *rgInfo.ResourceGroupId, nil
			}
			return "", nil // found image but no ResourceGroupId
		}

		// Check for more pages
		if resp.Body.NextToken == nil || *resp.Body.NextToken == "" {
			break
		}
		req.NextToken = resp.Body.NextToken
	}

	return "", nil // image not found
}

// IsUserImage checks if the image is a User type image
func IsUserImage(imageType string) bool {
	return imageType == "User"
}

// IsSystemImage checks if the image is a System type image
func IsSystemImage(imageType string) bool {
	return imageType == "System"
}

// PollForActivation polls the image status until it reaches RESOURCE_PUBLISHED or fails
func PollForActivation(ctx context.Context, apiClient agentbay.Client, imageId string, config PollingConfig) error {
	return pollForStatus(ctx, apiClient, imageId, config, []ImageResourceStatus{StatusResourcePublished}, "activation")
}

// PollForDeactivation polls the image status until it reaches IMAGE_AVAILABLE or fails
func PollForDeactivation(ctx context.Context, apiClient agentbay.Client, imageId string, config PollingConfig) error {
	return pollForStatus(ctx, apiClient, imageId, config, []ImageResourceStatus{StatusImageAvailable}, "deactivation")
}

// pollForStatus is the generic polling function
func pollForStatus(
	ctx context.Context,
	apiClient agentbay.Client,
	imageId string,
	config PollingConfig,
	expectedStatuses []ImageResourceStatus,
	operationName string,
) error {
	startTime := time.Now()
	interval := config.InitialInterval
	attempts := 0

	// Create a timeout context
	timeoutCtx, cancel := context.WithTimeout(ctx, config.Timeout)
	defer cancel()

	for {
		attempts++

		// Check if we've exceeded max attempts
		if attempts > config.MaxAttempts {
			return fmt.Errorf("%s polling exceeded maximum attempts (%d)", operationName, config.MaxAttempts)
		}

		// Check if context is cancelled or timed out
		select {
		case <-timeoutCtx.Done():
			return fmt.Errorf("%s polling timed out after %v", operationName, time.Since(startTime))
		default:
		}

		// Get current status
		status, err := GetImageResourceStatus(timeoutCtx, apiClient, imageId)
		if err != nil {
			log.Debugf("[DEBUG] Failed to get status (attempt %d/%d): %v", attempts, config.MaxAttempts, err)
			// Don't fail immediately on API errors, continue polling
		} else {
			currentStatus := ImageResourceStatus(status)
			translatedStatus := TranslateImageResourceStatus(status)

			log.Debugf("[DEBUG] Polling attempt %d/%d: Current status = %s (%s)",
				attempts, config.MaxAttempts, status, translatedStatus)

			// Check if we've reached a target status
			for _, expectedStatus := range expectedStatuses {
				if currentStatus == expectedStatus {
					fmt.Printf("[SUCCESS] %s completed! Current status: %s\n",
						operationName, translatedStatus)
					return nil
				}
			}

			// Check if we've reached a failed state
			if IsFailed(status) {
				return fmt.Errorf("%s failed with status: %s", operationName, translatedStatus)
			}

			// Update user with current status
			fmt.Printf("  Status: %s (elapsed: %v, attempt: %d/%d)\n",
				translatedStatus, time.Since(startTime).Round(time.Second), attempts, config.MaxAttempts)
		}

		// Wait before next attempt
		select {
		case <-timeoutCtx.Done():
			return fmt.Errorf("%s polling timed out after %v", operationName, time.Since(startTime))
		case <-time.After(interval):
			// Exponential backoff with max interval
			interval = interval * 3 / 2 // 1.5x multiplier
			if interval > config.MaxInterval {
				interval = config.MaxInterval
			}
		}
	}
}
