// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/agentbay/agentbay-cli/internal/agentbay"
	"github.com/agentbay/agentbay-cli/internal/client"
)

// runImageListWithBothTypes handles querying both user and system images
func runImageListWithBothTypes(ctx context.Context, apiClient *agentbay.Client, osType string, page, pageSize int) error {
	var allImages []*client.ListMcpImagesResponseBodyData
	var totalCount int32

	// First, get user images
	fmt.Printf("Requesting user images...")
	userReq := &client.ListMcpImagesRequest{}
	userImageType := "User"
	userReq.ImageType = &userImageType

	if osType != "" {
		userReq.OsType = &osType
	}
	if pageSize > 0 {
		pageSizeInt32 := int32(pageSize)
		userReq.PageSize = &pageSizeInt32
	}
	if page > 0 {
		pageInt32 := int32(page)
		userReq.PageStart = &pageInt32
	}

	userResp, err := apiClient.ListMcpImages(ctx, userReq)
	if err != nil {
		fmt.Printf(" Failed.\n")
		log.Debugf("[DEBUG] Failed to get user images: %v", err)
		return fmt.Errorf("failed to get user images: %w", err)
	}
	fmt.Printf(" Done.")

	// Process user images response
	if userResp != nil && userResp.Body != nil && userResp.Body.Data != nil {
		allImages = append(allImages, userResp.Body.Data...)
		if userResp.Body.TotalCount != nil {
			totalCount += *userResp.Body.TotalCount
		}
	}

	// Then, get system images
	fmt.Printf(" Requesting system images...")
	systemReq := &client.ListMcpImagesRequest{}
	systemImageType := "System"
	systemReq.ImageType = &systemImageType

	if osType != "" {
		systemReq.OsType = &osType
	}
	if pageSize > 0 {
		pageSizeInt32 := int32(pageSize)
		systemReq.PageSize = &pageSizeInt32
	}
	if page > 0 {
		pageInt32 := int32(page)
		systemReq.PageStart = &pageInt32
	}

	systemResp, err := apiClient.ListMcpImages(ctx, systemReq)
	if err != nil {
		fmt.Printf(" Failed.\n")
		log.Debugf("[DEBUG] Failed to get system images: %v", err)
		// Don't fail completely if system images fail, just show user images
		fmt.Printf("[WARN] Failed to fetch system images, showing user images only\n")
	} else {
		fmt.Printf(" Done.\n")
		// Process system images response
		if systemResp != nil && systemResp.Body != nil && systemResp.Body.Data != nil {
			allImages = append(allImages, systemResp.Body.Data...)
			if systemResp.Body.TotalCount != nil {
				totalCount += *systemResp.Body.TotalCount
			}
		}
	}

	// Display merged results
	if len(allImages) == 0 {
		fmt.Printf("\n[EMPTY] No images found.\n")
		return nil
	}

	fmt.Printf("\n[OK] Found %d images (Total: %d)\n", len(allImages), totalCount)

	// Display image table with consistent formatting
	fmt.Printf("%s %s %s %s %s %s\n",
		padString("IMAGE ID", 25),
		padString("IMAGE NAME", 30),
		padString("TYPE", 20),
		padString("STATUS", 15),
		padString("OS", 18),
		"APPLY SCENE")
	fmt.Printf("%s %s %s %s %s %s\n",
		padString("--------", 25),
		padString("----------", 30),
		padString("----", 20),
		padString("------", 15),
		padString("--", 18),
		"-----------")

	// Print each image
	for _, image := range allImages {
		if image == nil {
			continue
		}

		imageId := getStringValue(image.GetImageId())
		imageName := getStringValue(image.GetImageName())
		imageType := getStringValue(image.GetImageBuildType())
		status := formatImageStatus(getStringValue(image.GetImageResourceStatus()))
		osInfo := formatOSInfo(image.GetImageInfo())
		applyScene := getStringValue(image.GetImageApplyScene())

		// 使用支持中文的填充和截断函数，手动控制列间距
		fmt.Printf("%s %s %s %s %s %s\n",
			padString(truncateString(imageId, 25), 25),
			padString(truncateString(imageName, 30), 30),
			padString(imageType, 20),
			padString(status, 15),
			padString(osInfo, 18),
			applyScene)
	}

	return nil
}