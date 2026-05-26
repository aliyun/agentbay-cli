// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/agentbay/agentbay-cli/internal/agentbay"
	"github.com/agentbay/agentbay-cli/internal/client"
)

// imageItemJSON is the JSON representation of a single image entry.
type imageItemJSON struct {
	ImageId       string `json:"imageId"`
	ImageName     string `json:"imageName"`
	Type          string `json:"type"`
	Status        string `json:"status"`
	StatusDisplay string `json:"statusDisplay"`
	OsName        string `json:"osName"`
	OsVersion     string `json:"osVersion"`
	OsDisplay     string `json:"osDisplay"`
	ApplyScene    string `json:"applyScene"`
}

// printImagesAsJSON outputs image list data as indented JSON.
func printImagesAsJSON(images []*client.ListMcpImagesResponseBodyData, totalCount int32) error {
	type output struct {
		TotalCount int32           `json:"totalCount"`
		Images     []imageItemJSON `json:"images"`
	}
	out := output{TotalCount: totalCount}
	for _, image := range images {
		if image == nil {
			continue
		}
		item := imageItemJSON{
			ImageId:       getStringValue(image.GetImageId()),
			ImageName:     getStringValue(image.GetImageName()),
			Type:          getStringValue(image.GetImageBuildType()),
			Status:        getStringValue(image.GetImageResourceStatus()),
			StatusDisplay: formatImageStatus(getStringValue(image.GetImageResourceStatus())),
			ApplyScene:    getStringValue(image.GetImageApplyScene()),
		}
		if imgInfo := image.GetImageInfo(); imgInfo != nil {
			item.OsName = getStringValue(imgInfo.GetOsName())
			item.OsVersion = getStringValue(imgInfo.GetOsVersion())
			item.OsDisplay = formatOSInfo(imgInfo)
		}
		out.Images = append(out.Images, item)
	}
	if out.Images == nil {
		out.Images = []imageItemJSON{}
	}
	b, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return fmt.Errorf("json marshal: %w", err)
	}
	fmt.Println(string(b))
	return nil
}

// runImageListWithBothTypes handles querying both user and system images
func runImageListWithBothTypes(ctx context.Context, apiClient agentbay.Client, osType string, page, pageSize int, outputFmt string) error {
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

	// Separate user and system images
	var userImages []*client.ListMcpImagesResponseBodyData
	var systemImages []*client.ListMcpImagesResponseBodyData

	for _, image := range allImages {
		if image == nil {
			continue
		}
		// Check if this is a user image (starts with "imgc-") or system image
		imageId := getStringValue(image.GetImageId())
		if strings.HasPrefix(imageId, "imgc-") {
			userImages = append(userImages, image)
		} else {
			systemImages = append(systemImages, image)
		}
	}

	// Display results
	if len(allImages) == 0 {
		fmt.Printf("\n[EMPTY] No images found.\n")
		return nil
	}

	// JSON output mode
	if strings.EqualFold(outputFmt, "json") {
		return printImagesAsJSON(allImages, totalCount)
	}

	fmt.Printf("\n[OK] Found %d images (Total: %d)\n", len(allImages), totalCount)

	// Display user images first
	if len(userImages) > 0 {
		fmt.Printf("\n=== USER IMAGES (%d) ===\n", len(userImages))
		printImageTable(userImages)
	}

	// Display system images
	if len(systemImages) > 0 {
		fmt.Printf("\n=== SYSTEM IMAGES (%d) ===\n", len(systemImages))
		printImageTable(systemImages)
	}

	return nil
}

// printImageTable prints a formatted table of images
func printImageTable(images []*client.ListMcpImagesResponseBodyData) {
	// Print header
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
	for _, image := range images {
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
}
