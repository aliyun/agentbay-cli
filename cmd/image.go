// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/agentbay/agentbay-cli/internal/agentbay"
	"github.com/agentbay/agentbay-cli/internal/client"
	"github.com/agentbay/agentbay-cli/internal/config"
	"github.com/alibabacloud-go/tea/dara"
)

// printErrorMessage prints multi-line error messages by printing each line separately
// This avoids Windows line ending issues
func printErrorMessage(lines ...string) error {
	// Print each line to stderr for immediate display
	for _, line := range lines {
		fmt.Fprintln(os.Stderr, line)
	}
	// Return a simple error for the command framework
	return fmt.Errorf("command failed")
}

var ImageCmd = &cobra.Command{
	Use:     "image",
	Short:   "Manage AgentBay images",
	Long:    "Create, build, and manage custom AgentBay images",
	GroupID: "management",
}

var imageCreateCmd = &cobra.Command{
	Use:   "create <image-name>",
	Short: "Create a new AgentBay image",
	Long: `Create a new AgentBay image from a Dockerfile.

This command builds a custom image that can be used in AgentBay environments.
The image will be built from the specified Dockerfile and based on the provided source image.

Examples:
  # Create an image with a custom Dockerfile
  agentbay image create my-custom-image --dockerfile ./Dockerfile --imageId code_latest
  
  # Short form
  agentbay image create my-image -f ./Dockerfile -i code_latest`,
	Args: cobra.ExactArgs(1),
	RunE: runImageCreate,
}

var imageListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available AgentBay images",
	Long: `List available AgentBay images that can be used as base images for custom builds.

This command queries the AgentBay platform for available User images by default and displays
their details including image ID, name, type, and description.

OS types:
  Linux   - Linux-based images
  Android - Android-based images  
  Windows - Windows-based images

Examples:
  # List user images (default)
  agentbay image list
  
  # Include system images
  agentbay image list --include-system
  
  # Show only system images
  agentbay image list --system-only
  
  # List Linux images only
  agentbay image list --os-type Linux
  
  # List images with pagination
  agentbay image list --page 2 --size 5`,
	RunE: runImageList,
}

var imageActivateCmd = &cobra.Command{
	Use:   "activate <image-id>",
	Short: "Activate a User image",
	Long: `Activate a User image to make it available for use.

This command creates a resource group for the specified User image, making it 
available for deployment. Only User type images can be activated.

Supported CPU and Memory combinations:
  2c4g  - 2 CPU cores with 4 GB memory
  4c8g  - 4 CPU cores with 8 GB memory
  8c16g - 8 CPU cores with 16 GB memory

If no CPU/memory is specified, default resources will be used.

Examples:
  # Activate a user image with default resources
  agentbay image activate imgc-xxxxxxxxxxxxxx

  # Activate with specific CPU and memory
  agentbay image activate imgc-xxxxxxxxxxxxxx --cpu 2 --memory 4

  # Activate with verbose output
  agentbay image activate imgc-xxxxxxxxxxxxxx --cpu 4 --memory 8 --verbose`,
	Args: cobra.ExactArgs(1),
	RunE: runImageActivate,
}

var imageDeactivateCmd = &cobra.Command{
	Use:   "deactivate <image-id>",
	Short: "Deactivate an activated User image",
	Long: `Deactivate an activated User image to stop its resource group.

This command deletes the resource group for the specified User image, making it 
unavailable for deployment. Only activated User type images can be deactivated.

Examples:
  # Deactivate a user image
  agentbay image deactivate imgc-xxxxxxxxxxxxxx
  
  # Deactivate with verbose output
  agentbay image deactivate imgc-xxxxxxxxxxxxxx --verbose`,
	Args: cobra.ExactArgs(1),
	RunE: runImageDeactivate,
}

func init() {
	// Add flags to image create command
	imageCreateCmd.Flags().StringP("dockerfile", "f", "", "Path to the Dockerfile (required)")
	imageCreateCmd.Flags().StringP("imageId", "i", "", "Source image ID to build from (required)")

	// Mark required flags
	imageCreateCmd.MarkFlagRequired("dockerfile")
	imageCreateCmd.MarkFlagRequired("imageId")

	// Add flags to image activate command
	imageActivateCmd.Flags().IntP("cpu", "c", 0, "CPU cores (e.g., 2, 4, 8)")
	imageActivateCmd.Flags().IntP("memory", "m", 0, "Memory in GB (e.g., 4, 8, 16)")

	// Add flags to image list command
	imageListCmd.Flags().StringP("os-type", "o", "", "Filter by OS type: Linux, Android, or Windows (optional)")
	imageListCmd.Flags().Bool("include-system", false, "Include system images in addition to user images")
	imageListCmd.Flags().Bool("system-only", false, "Show only system images")
	imageListCmd.Flags().IntP("page", "p", 1, "Page number (default: 1)")
	imageListCmd.Flags().IntP("size", "s", 10, "Page size (default: 10)")

	// Add subcommands to image command
	ImageCmd.AddCommand(imageCreateCmd)
	ImageCmd.AddCommand(imageListCmd)
	ImageCmd.AddCommand(imageActivateCmd)
	ImageCmd.AddCommand(imageDeactivateCmd)
}

func runImageCreate(cmd *cobra.Command, args []string) error {
	imageName := args[0]
	dockerfilePath, _ := cmd.Flags().GetString("dockerfile")
	sourceImageId, _ := cmd.Flags().GetString("imageId")

	// Validate required flags with friendly messages
	if dockerfilePath == "" {
		return printErrorMessage(
			fmt.Sprintf("[ERROR] Missing required flag: --dockerfile for %s", imageName),
			"",
			fmt.Sprintf("[TIP] Usage: agentbay image create %s --dockerfile <path> --imageId <id>", imageName),
			fmt.Sprintf("[NOTE] Example: agentbay image create %s --dockerfile ./Dockerfile --imageId code_latest", imageName),
			fmt.Sprintf("[NOTE] Short form: agentbay image create %s -f ./Dockerfile -i code_latest", imageName),
		)
	}
	if sourceImageId == "" {
		return printErrorMessage(
			fmt.Sprintf("[ERROR] Missing required flag: --imageId for %s", imageName),
			"",
			fmt.Sprintf("[TIP] Usage: agentbay image create %s --dockerfile <path> --imageId <id>", imageName),
			fmt.Sprintf("[NOTE] Example: agentbay image create %s --dockerfile ./Dockerfile --imageId code_latest", imageName),
			fmt.Sprintf("[NOTE] Short form: agentbay image create %s -f ./Dockerfile -i code_latest", imageName),
		)
	}

	// Validate dockerfile path
	if !filepath.IsAbs(dockerfilePath) {
		var err error
		dockerfilePath, err = filepath.Abs(dockerfilePath)
		if err != nil {
			return fmt.Errorf("failed to resolve dockerfile path: %w", err)
		}
	}

	if _, err := os.Stat(dockerfilePath); os.IsNotExist(err) {
		return fmt.Errorf("dockerfile not found: %s", dockerfilePath)
	}

	fmt.Printf("[BUILD] Creating image '%s'...\n", imageName)

	// Load configuration and check authentication
	cfg, err := config.GetConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Failed to load configuration: %v\n", err)
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	if !cfg.IsAuthenticated() {
		fmt.Fprintf(os.Stderr, "[ERROR] Not authenticated. Please run 'agentbay login' first\n")
		return fmt.Errorf("not authenticated. Please run 'agentbay login' first")
	}

	// Create API client
	apiClient := agentbay.NewClientFromConfig(cfg)
	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Minute)
	defer cancel()

	// Validate source image ID exists before proceeding
	fmt.Printf("Validating source image ID '%s'...\n", sourceImageId)
	validateCtx, validateCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer validateCancel()

	_, err = GetImageInfo(validateCtx, apiClient, sourceImageId)
	if err != nil {
		// Check if the error is an authentication error
		if IsAuthenticationError(err) {
			return printErrorMessage(
				fmt.Sprintf("[ERROR] Authentication failed. Please run 'agentbay login' first."),
				"",
			)
		}
		return printErrorMessage(
			fmt.Sprintf("[ERROR] Source image not found: %s", sourceImageId),
			"",
			fmt.Sprintf("[TIP] The specified source image ID '%s' does not exist or is not accessible.", sourceImageId),
			"[TIP] Use 'agentbay image list' to see available images that can be used as base images.",
			"[NOTE] Example: agentbay image list",
		)
	}
	fmt.Printf(" Done.\n")

	// Step 1: Get Docker file store credentials
	fmt.Printf("[STEP 1/4] Getting upload credentials...\n")
	sourceAgentBay := "AgentBay"
	credReq := &client.GetDockerFileStoreCredentialRequest{
		Source: &sourceAgentBay,
	}

	// Debug: Print credential request (simplified)
	if log.GetLevel() >= log.DebugLevel {
		log.Debugf("[DEBUG] GetDockerFileStoreCredential Request:")
		if credReq.Source != nil {
			log.Debugf("[DEBUG] - Source: %s", *credReq.Source)
		}
	}

	log.Debugf("[DEBUG] Making GetDockerFileStoreCredential API call...")
	fmt.Printf("Requesting upload credentials...")
	credResp, err := apiClient.GetDockerFileStoreCredential(ctx, credReq)
	if err != nil {
		log.Debugf("[DEBUG] GetDockerFileStoreCredential API call failed: %v", err)
		// Show user-friendly error message in non-verbose mode
		fmt.Printf("[ERROR] Failed to get upload credentials. Please check your authentication and try again.\n")
		if log.GetLevel() >= log.DebugLevel {
			fmt.Printf("[DEBUG] Error details: %v\n", err)
		}
		return fmt.Errorf("failed to get upload credentials: %w", err)
	}
	fmt.Printf(" Done.\n")
	log.Debugf("[DEBUG] GetDockerFileStoreCredential API call completed successfully")

	// Debug: Print credential response (simplified)
	if log.GetLevel() >= log.DebugLevel && credResp.Body != nil && credResp.Body.Data != nil {
		taskId := credResp.Body.Data.GetTaskId()
		ossUrl := credResp.Body.Data.GetOssUrl()

		log.Debugf("[DEBUG] GetDockerFileStoreCredential Response:")
		if taskId != nil {
			log.Debugf("[DEBUG] - TaskId: %s", *taskId)
		}
		if ossUrl != nil {
			log.Debugf("[DEBUG] - OssUrl: %s", *ossUrl)
		}
	}

	if credResp.Body == nil || credResp.Body.Data == nil {
		return fmt.Errorf("invalid response: missing upload credentials")
	}

	ossUrl := credResp.Body.Data.GetOssUrl()
	taskId := credResp.Body.Data.GetTaskId()

	if ossUrl == nil || taskId == nil {
		return fmt.Errorf("invalid response: missing OSS URL or task ID")
	}

	fmt.Printf("[STEP 2/4] Uploading Dockerfile...\n")

	// Step 2: Upload Dockerfile to OSS
	fmt.Printf("Uploading file...")
	err = uploadDockerfile(dockerfilePath, *ossUrl)
	if err != nil {
		// Show user-friendly error message in non-verbose mode
		fmt.Printf("[ERROR] Failed to upload Dockerfile. Please check your network connection and try again.\n")
		if log.GetLevel() >= log.DebugLevel {
			fmt.Printf("[DEBUG] Error details: %v\n", err)
		}
		return fmt.Errorf("failed to upload Dockerfile: %w", err)
	}
	fmt.Printf(" Done.\n")

	fmt.Printf("[STEP 3/4] Creating Docker image task...\n")

	// Step 3: Create Docker image task
	createReq := &client.CreateDockerImageTaskRequest{
		ImageName:     &imageName,
		Source:        &sourceAgentBay,
		SourceImageId: &sourceImageId,
		TaskId:        taskId,
	}

	// Debug: Print create task request (simplified)
	if log.GetLevel() >= log.DebugLevel {
		log.Debugf("[DEBUG] CreateDockerImageTask Request:")
		if createReq.ImageName != nil {
			log.Debugf("[DEBUG] - ImageName: %s", *createReq.ImageName)
		}
		if createReq.Source != nil {
			log.Debugf("[DEBUG] - Source: %s", *createReq.Source)
		}
		if createReq.SourceImageId != nil {
			log.Debugf("[DEBUG] - SourceImageId: %s", *createReq.SourceImageId)
		}
		if createReq.TaskId != nil {
			log.Debugf("[DEBUG] - TaskId: %s", *createReq.TaskId)
		}
	}

	fmt.Printf("Creating image task...")
	createResp, err := apiClient.CreateDockerImageTask(ctx, createReq)
	if err != nil {
		// Show user-friendly error message in non-verbose mode
		fmt.Printf("[ERROR] Failed to create Docker image task. Please try again.\n")
		if log.GetLevel() >= log.DebugLevel {
			fmt.Printf("[DEBUG] Error details: %v\n", err)
		}
		// Try to extract Request ID from response if available
		if createResp != nil && createResp.Body != nil && createResp.Body.GetRequestId() != nil {
			fmt.Printf("[DEBUG] Request ID: %s\n", *createResp.Body.GetRequestId())
		}
		return fmt.Errorf("failed to create Docker image task: %w", err)
	}
	fmt.Printf(" Done.\n")

	// Debug: Print create task response (simplified)
	if log.GetLevel() >= log.DebugLevel && createResp.Body != nil && createResp.Body.Data != nil {
		taskId := createResp.Body.Data.GetTaskId()

		log.Debugf("[DEBUG] CreateDockerImageTask Response:")
		if taskId != nil {
			log.Debugf("[DEBUG] - TaskId: %s", *taskId)
		}
	}

	if createResp.Body == nil || createResp.Body.Data == nil {
		// Print Request ID for debugging if available
		if createResp != nil && createResp.Body != nil && createResp.Body.GetRequestId() != nil {
			fmt.Printf("[DEBUG] Request ID: %s\n", *createResp.Body.GetRequestId())
		}
		return fmt.Errorf("invalid response: missing task data")
	}

	finalTaskId := createResp.Body.Data.GetTaskId()
	if finalTaskId == nil {
		// Print Request ID for debugging
		if createResp.Body.GetRequestId() != nil {
			fmt.Printf("[DEBUG] Request ID: %s\n", *createResp.Body.GetRequestId())
		}
		return fmt.Errorf("invalid response: missing final task ID")
	}

	fmt.Printf("[STEP 4/4] Building image (Task ID: %s)...\n", *finalTaskId)

	// Step 4: Poll for task completion
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("build timeout: %w", ctx.Err())
		case <-ticker.C:
			sourceAgentBay := "AgentBay"
			taskReq := &client.GetDockerImageTaskRequest{
				Source: &sourceAgentBay,
				TaskId: finalTaskId,
			}

			// Debug: Print polling request (simplified)
			if log.GetLevel() >= log.DebugLevel {
				log.Debugf("[DEBUG] GetDockerImageTask Request:")
				if taskReq.Source != nil {
					log.Debugf("[DEBUG] - Source: %s", *taskReq.Source)
				}
				if taskReq.TaskId != nil {
					log.Debugf("[DEBUG] - TaskId: %s", *taskReq.TaskId)
				}
			}

			taskResp, err := apiClient.GetDockerImageTask(ctx, taskReq)
			if err != nil {
				log.Debugf("[DEBUG] GetDockerImageTask Polling Error: %v", err)
				fmt.Printf("[WARN] Warning: Failed to check task status: %v\n", err)
				// Try to extract Request ID from response if available
				if taskResp != nil && taskResp.Body != nil && taskResp.Body.GetRequestId() != nil {
					fmt.Printf("[DEBUG] Request ID: %s\n", *taskResp.Body.GetRequestId())
				}
				continue // Continue polling on API errors
			}

			// Debug: Print polling response (simplified)
			if log.GetLevel() >= log.DebugLevel && taskResp.Body != nil && taskResp.Body.Data != nil {
				status := taskResp.Body.Data.GetStatus()
				taskMsg := taskResp.Body.Data.GetTaskMsg()
				imageId := taskResp.Body.Data.GetImageId()

				log.Debugf("[DEBUG] GetDockerImageTask Response:")
				if status != nil {
					log.Debugf("[DEBUG] - Status: %s", *status)
				}
				if taskMsg != nil && *taskMsg != "" {
					log.Debugf("[DEBUG] - Message: %s", *taskMsg)
				}
				if imageId != nil && *imageId != "" {
					log.Debugf("[DEBUG] - ImageId: %s", *imageId)
				}
			}

			if taskResp.Body == nil || taskResp.Body.Data == nil {
				fmt.Printf("[WARN] Warning: Invalid response format\n")
				// Print Request ID for debugging if available
				if taskResp != nil && taskResp.Body != nil && taskResp.Body.GetRequestId() != nil {
					fmt.Printf("[DEBUG] Request ID: %s\n", *taskResp.Body.GetRequestId())
				}
				continue
			}

			status := taskResp.Body.Data.GetStatus()
			taskMsg := taskResp.Body.Data.GetTaskMsg()
			imageId := taskResp.Body.Data.GetImageId()

			if status == nil {
				fmt.Printf("[WARN] Warning: Missing status in response\n")
				// Print Request ID for debugging
				if taskResp.Body.GetRequestId() != nil {
					fmt.Printf("[DEBUG] Request ID: %s\n", *taskResp.Body.GetRequestId())
				}
				continue
			}

			fmt.Printf("[STATUS] Build status: %s\n", *status)

			if taskMsg != nil && *taskMsg != "" {
				fmt.Printf("[MESSAGE] %s\n", *taskMsg)
			}

			switch *status {
			case "SUCCESS", "Finished":
				fmt.Printf("[SUCCESS] ✅ Image '%s' created successfully!\n", imageName)
				if imageId != nil && *imageId != "" {
					fmt.Printf("[RESULT] Image ID: %s\n", *imageId)
				}
				fmt.Printf("[DOC] Task ID: %s\n", *finalTaskId)
				return nil
			case "FAILED", "Failed":
				fmt.Printf("[ERROR] ❌ Image build failed\n")
				if taskMsg != nil && *taskMsg != "" {
					fmt.Printf("[ERROR] Error details: %s\n", *taskMsg)
				}
				// Print Request ID for debugging
				if taskResp.Body.GetRequestId() != nil {
					fmt.Printf("[DEBUG] Request ID: %s\n", *taskResp.Body.GetRequestId())
				}
				fmt.Printf("[DOC] Task ID: %s\n", *finalTaskId)
				return fmt.Errorf("image build failed")
			case "RUNNING", "PENDING", "Preparing":
				// Continue polling
				continue
			default:
				fmt.Printf("[WARN] Warning: Unknown status: %s\n", *status)
				continue
			}
		}
	}
}

func runImageList(cmd *cobra.Command, args []string) error {
	// Get flag values
	osType, _ := cmd.Flags().GetString("os-type")
	includeSystem, _ := cmd.Flags().GetBool("include-system")
	systemOnly, _ := cmd.Flags().GetBool("system-only")
	page, _ := cmd.Flags().GetInt("page")
	pageSize, _ := cmd.Flags().GetInt("size")

	// Determine what type of images to fetch
	var fetchMessage string
	if systemOnly {
		fetchMessage = "[LIST] Fetching available AgentBay system images...\n"
	} else if includeSystem {
		fetchMessage = "[LIST] Fetching available AgentBay images (user + system)...\n"
	} else {
		fetchMessage = "[LIST] Fetching available AgentBay user images...\n"
	}
	fmt.Printf(fetchMessage)

	// Load configuration and check authentication
	cfg, err := config.GetConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Failed to load configuration: %v\n", err)
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	if !cfg.IsAuthenticated() {
		fmt.Fprintf(os.Stderr, "[ERROR] Not authenticated. Please run 'agentbay login' first\n")
		return fmt.Errorf("not authenticated. Please run 'agentbay login' first")
	}

	// Create API client
	apiClient := agentbay.NewClientFromConfig(cfg)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Prepare request
	req := &client.ListMcpImagesRequest{}

	// Handle different image type queries
	if includeSystem {
		// For include-system, we need to make two API calls and merge results
		return runImageListWithBothTypes(ctx, apiClient, osType, page, pageSize)
	}

	// Single query for system-only or user-only (default)
	var imageType string
	if systemOnly {
		imageType = "System"
	} else {
		// Default behavior: only user images
		imageType = "User"
	}

	// Prepare request
	req = &client.ListMcpImagesRequest{}
	req.ImageType = &imageType
	if osType != "" {
		req.OsType = &osType
	}
	if pageSize > 0 {
		pageSizeInt32 := int32(pageSize)
		req.PageSize = &pageSizeInt32
	}
	if page > 0 {
		pageInt32 := int32(page)
		req.PageStart = &pageInt32
	}

	// Debug: Print request details
	if log.GetLevel() >= log.DebugLevel {
		log.Debugf("[DEBUG] ListMcpImages Request:")
		if imageType != "" {
			log.Debugf("[DEBUG] - ImageType: %s", imageType)
		} else {
			log.Debugf("[DEBUG] - ImageType: (all types)")
		}
		if req.OsType != nil {
			log.Debugf("[DEBUG] - OsType: %s", *req.OsType)
		}
		if req.PageSize != nil {
			log.Debugf("[DEBUG] - PageSize: %d", *req.PageSize)
		}
		if req.PageStart != nil {
			log.Debugf("[DEBUG] - PageStart: %d", *req.PageStart)
		}
	}

	// Make API call
	fmt.Printf("Requesting image list...")
	resp, err := apiClient.ListMcpImages(ctx, req)
	if err != nil {
		log.Debugf("[DEBUG] ListMcpImages API call failed: %v", err)
		fmt.Printf("[ERROR] Failed to fetch image list. Please check your authentication and try again.\n")
		if log.GetLevel() >= log.DebugLevel {
			fmt.Printf("[DEBUG] Error details: %v\n", err)
		}
		return fmt.Errorf("failed to fetch image list: %w", err)
	}
	fmt.Printf(" Done.\n")

	// Debug: Print response details
	if log.GetLevel() >= log.DebugLevel && resp.Body != nil {
		log.Debugf("[DEBUG] ListMcpImages Response:")
		if resp.Body.GetRequestId() != nil {
			log.Debugf("[DEBUG] - RequestId: %s", *resp.Body.GetRequestId())
		}
		if resp.Body.GetSuccess() != nil {
			log.Debugf("[DEBUG] - Success: %t", *resp.Body.GetSuccess())
		}
		if resp.Body.GetTotalCount() != nil {
			log.Debugf("[DEBUG] - TotalCount: %d", *resp.Body.GetTotalCount())
		}
	}

	// Validate response
	if resp.Body == nil {
		return fmt.Errorf("invalid response: missing response body")
	}

	if resp.Body.GetSuccess() != nil && !*resp.Body.GetSuccess() {
		errorMsg := "unknown error"
		if resp.Body.GetMessage() != nil {
			errorMsg = *resp.Body.GetMessage()
		}
		return fmt.Errorf("API request failed: %s", errorMsg)
	}

	images := resp.Body.GetData()
	if len(images) == 0 {
		fmt.Printf("\n[EMPTY] No images found.\n")
		return nil
	}

	// Display results
	fmt.Printf("\n[OK] Found %d images", len(images))
	if resp.Body.GetTotalCount() != nil {
		fmt.Printf(" (Total: %d)", *resp.Body.GetTotalCount())
	}
	fmt.Printf("\n")

	if resp.Body.GetPageStart() != nil && resp.Body.GetPageSize() != nil && resp.Body.GetTotalCount() != nil {
		pageSize := *resp.Body.GetPageSize()
		if pageSize > 0 {
			totalPages := (*resp.Body.GetTotalCount() + pageSize - 1) / pageSize
			fmt.Printf("[PAGE] Page %d of %d (Page Size: %d)\n\n", *resp.Body.GetPageStart(), totalPages, pageSize)
		}
	}

	if len(images) == 0 {
		fmt.Println("[EMPTY] No images found.")
		return nil
	}

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

	for _, image := range images {
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
			padString(truncateString(imageType, 20), 20),
			padString(truncateString(status, 15), 15),
			padString(truncateString(osInfo, 18), 18),
			truncateString(applyScene, 15)) // 最后一列不需要填充
	}

	return nil
}

// Helper functions for formatting table output

// getStringValue safely extracts string value from pointer, returns empty string if nil
func getStringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// truncateString truncates a string to the specified length, adding "..." if truncated
func truncateString(s string, maxLen int) string {
	if displayWidth(s) <= maxLen {
		return s
	}
	if maxLen <= 3 {
		return s[:maxLen]
	}

	// 逐字符截断直到达到合适的显示宽度
	runes := []rune(s)
	width := 0
	for i, r := range runes {
		charWidth := runeDisplayWidth(r)
		if width+charWidth+3 > maxLen { // 为 "..." 预留3个字符
			return string(runes[:i]) + "..."
		}
		width += charWidth
	}
	return s
}

// padString 填充字符串到指定显示宽度，支持中文字符
func padString(s string, width int) string {
	currentWidth := displayWidth(s)
	if currentWidth >= width {
		return s
	}
	// 添加空格填充到指定宽度
	padding := width - currentWidth
	return s + strings.Repeat(" ", padding)
}

// displayWidth 计算字符串的显示宽度，中文字符算2个宽度
func displayWidth(s string) int {
	width := 0
	for _, r := range s {
		width += runeDisplayWidth(r)
	}
	return width
}

// runeDisplayWidth 计算单个字符的显示宽度
func runeDisplayWidth(r rune) int {
	// 中文字符、全角字符等占用2个显示宽度
	if r >= 0x4e00 && r <= 0x9fff || // CJK统一汉字
		r >= 0x3400 && r <= 0x4dbf || // CJK扩展A
		r >= 0x20000 && r <= 0x2a6df || // CJK扩展B
		r >= 0x2a700 && r <= 0x2b73f || // CJK扩展C
		r >= 0x2b740 && r <= 0x2b81f || // CJK扩展D
		r >= 0x2b820 && r <= 0x2ceaf || // CJK扩展E
		r >= 0xf900 && r <= 0xfaff || // CJK兼容汉字
		r >= 0x2f800 && r <= 0x2fa1f || // CJK兼容汉字补充
		r >= 0xff00 && r <= 0xffef { // 全角ASCII、半角片假名、全角符号
		return 2
	}
	return 1
}

// formatImageStatus formats image status for better readability
func formatImageStatus(status string) string {
	switch status {
	case "IMAGE_CREATING":
		return "Creating"
	case "IMAGE_CREATE_FAILED":
		return "Create Failed"
	case "IMAGE_AVAILABLE":
		return "Available"
	case "RESOURCE_DEPLOYING":
		return "Activating"
	case "RESOURCE_PUBLISHED":
		return "Activated"
	case "RESOURCE_DELETING":
		return "Deactivating"
	case "RESOURCE_FAILED":
		return "Activate Failed"
	case "RESOURCE_CEASED":
		return "Ceased"
	default:
		if status == "" {
			return "-"
		}
		return status
	}
}

// ValidateCPUMemoryCombo validates that CPU and memory combination is supported
func ValidateCPUMemoryCombo(cpu, memory int) error {
	// If both are 0, use default (no validation needed)
	if cpu == 0 && memory == 0 {
		return nil
	}

	// If only one is specified, both must be specified
	if (cpu == 0 && memory > 0) || (cpu > 0 && memory == 0) {
		return fmt.Errorf("Both CPU and memory must be specified together. Supported combinations: 2c4g (--cpu 2 --memory 4), 4c8g (--cpu 4 --memory 8), 8c16g (--cpu 8 --memory 16)")
	}

	// Check supported combinations
	validCombos := map[int]int{
		2: 4,  // 2c4g
		4: 8,  // 4c8g
		8: 16, // 8c16g
	}

	expectedMemory, exists := validCombos[cpu]
	if !exists || expectedMemory != memory {
		return fmt.Errorf("Invalid CPU/Memory combination: %dc%dg. Supported combinations: 2c4g (--cpu 2 --memory 4), 4c8g (--cpu 4 --memory 8), 8c16g (--cpu 8 --memory 16)", cpu, memory)
	}

	return nil
}

// printCPUMemoryValidationError prints validation error with nice formatting
func printCPUMemoryValidationError(err error) error {
	if err == nil {
		return nil
	}
	// Print formatted error message to stderr
	lines := []string{
		"[ERROR] " + err.Error(),
	}
	return printErrorMessage(lines...)
}

// formatOSInfo formats OS information for compact display
func formatOSInfo(imageInfo *client.ListMcpImagesResponseBodyDataImageInfo) string {
	if imageInfo == nil {
		return "-"
	}

	osName := getStringValue(imageInfo.GetOsName())
	osVersion := getStringValue(imageInfo.GetOsVersion())

	if osName == "" && osVersion == "" {
		return "-"
	}

	if osName == "" {
		return osVersion
	}

	if osVersion == "" {
		return osName
	}

	// Format as "OS Version" (e.g., "Linux Debian", "Windows 2022")
	if osName == "Linux" && (osVersion == "Debian" || osVersion == "Ubuntu 2204") {
		return osName + " " + osVersion
	}

	if osName == "Windows" && osVersion == "Windows Server 2022" {
		return "Windows 2022"
	}

	if osName == "Android" {
		return osVersion // "Android 14", "Android 12"
	}

	return osName
}

// uploadDockerfile uploads the Dockerfile to the provided OSS URL with retry mechanism
func uploadDockerfile(dockerfilePath, ossUrl string) error {
	log.Debugf("[DEBUG] Starting Dockerfile upload...")
	log.Debugf("[DEBUG] - Dockerfile path: %s", dockerfilePath)
	log.Debugf("[DEBUG] - OSS URL: %s", ossUrl)

	// Read the Dockerfile
	dockerfileContent, err := os.ReadFile(dockerfilePath)
	if err != nil {
		log.Debugf("[DEBUG] Failed to read Dockerfile: %v", err)
		return fmt.Errorf("failed to read Dockerfile: %w", err)
	}

	log.Debugf("[DEBUG] Dockerfile content length: %d bytes", len(dockerfileContent))
	previewLen := 100
	if len(dockerfileContent) < previewLen {
		previewLen = len(dockerfileContent)
	}
	log.Debugf("[DEBUG] Dockerfile content preview: %s", string(dockerfileContent[:previewLen]))

	// Create retry configuration for upload
	retryConfig := &client.RetryConfig{
		MaxRetries:    3,
		InitialDelay:  1 * time.Second,
		MaxDelay:      10 * time.Second,
		BackoffFactor: 2.0,
	}

	var lastErr error
	delay := retryConfig.InitialDelay

	for attempt := 0; attempt <= retryConfig.MaxRetries; attempt++ {
		fmt.Printf("[UPLOAD] Dockerfile upload attempt %d/%d...\n", attempt+1, retryConfig.MaxRetries+1)
		log.Debugf("[DEBUG] Attempt %d/%d for uploading Dockerfile", attempt+1, retryConfig.MaxRetries+1)

		// Create HTTP PUT request for each attempt
		req, err := http.NewRequest(http.MethodPut, ossUrl, strings.NewReader(string(dockerfileContent)))
		if err != nil {
			log.Debugf("[DEBUG] Failed to create upload request: %v", err)
			return fmt.Errorf("failed to create upload request: %w", err)
		}

		// Set appropriate headers
		req.Header.Set("Content-Type", "application/octet-stream")
		req.Header.Set("User-Agent", "AgentBay-CLI/1.0")
		req.ContentLength = int64(len(dockerfileContent))

		log.Debugf("[DEBUG] Upload request headers:")
		for key, values := range req.Header {
			for _, value := range values {
				log.Debugf("[DEBUG] - %s: %s", key, value)
			}
		}

		// Perform the upload
		httpClient := &http.Client{
			Timeout: 60 * time.Second,
		}

		log.Debugf("[DEBUG] Sending upload request...")
		resp, err := httpClient.Do(req)

		// Success case
		if err == nil && resp.StatusCode >= 200 && resp.StatusCode < 300 {
			resp.Body.Close()
			log.Debugf("[DEBUG] Upload response received:")
			log.Debugf("[DEBUG] - Status: %s", resp.Status)
			log.Debugf("[DEBUG] - Status Code: %d", resp.StatusCode)
			if attempt > 0 {
				fmt.Printf("[OK] Dockerfile upload succeeded on attempt %d\n", attempt+1)
				log.Infof("[RETRY] Request succeeded on attempt %d", attempt+1)
			}
			return nil
		}

		// Handle error cases
		if err != nil {
			lastErr = fmt.Errorf("failed to upload dockerfile: %w", err)
			log.Debugf("[DEBUG] Attempt %d failed with error: %v", attempt+1, err)
		} else {
			// Read response body for error details
			body, _ := io.ReadAll(resp.Body)
			resp.Body.Close()
			lastErr = fmt.Errorf("upload failed with status %d: %s", resp.StatusCode, string(body))
			log.Debugf("[DEBUG] Attempt %d failed with HTTP status: %d", attempt+1, resp.StatusCode)
			log.Debugf("[DEBUG] Response body: %s", string(body))
		}

		// Don't retry if this is the last attempt
		if attempt == retryConfig.MaxRetries {
			break
		}

		// Check if the error is retryable
		shouldRetry := false
		if err != nil {
			// Use the same retry logic as the API client
			shouldRetry = client.IsRetryableError(err)
		} else if resp != nil {
			// Check if HTTP status is retryable
			shouldRetry = client.IsRetryableHTTPStatus(resp.StatusCode)
		}

		if !shouldRetry {
			log.Debugf("[DEBUG] Error is not retryable, stopping attempts")
			fmt.Printf("[WARN] Upload error is not retryable, stopping attempts\n")
			break
		}

		// Wait before retrying
		fmt.Printf("[RETRY] Upload failed (attempt %d/%d), retrying in %v...\n",
			attempt+1, retryConfig.MaxRetries+1, delay)
		log.Infof("[RETRY] Request failed (attempt %d/%d), retrying in %v...",
			attempt+1, retryConfig.MaxRetries+1, delay)

		time.Sleep(delay)

		// Calculate next delay with exponential backoff
		delay = time.Duration(float64(delay) * retryConfig.BackoffFactor)
		if delay > retryConfig.MaxDelay {
			delay = retryConfig.MaxDelay
		}
	}

	fmt.Printf("[ERROR] All %d upload attempts failed\n", retryConfig.MaxRetries+1)
	log.Warnf("[RETRY] All %d attempts failed, giving up", retryConfig.MaxRetries+1)
	return fmt.Errorf("dockerfile upload failed after %d attempts, last error: %w",
		retryConfig.MaxRetries+1, lastErr)
}

func runImageActivate(cmd *cobra.Command, args []string) error {
	imageId := args[0]
	cpu, _ := cmd.Flags().GetInt("cpu")
	memory, _ := cmd.Flags().GetInt("memory")

	// Validate CPU and memory combination
	if err := ValidateCPUMemoryCombo(cpu, memory); err != nil {
		return printCPUMemoryValidationError(err)
	}

	fmt.Printf("[ACTIVATE] Activating image '%s'...\n", imageId)
	if cpu > 0 || memory > 0 {
		fmt.Printf("[RESOURCE] CPU: %d cores, Memory: %d GB\n", cpu, memory)
	}

	// Load configuration and check authentication
	cfg, err := config.GetConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Failed to load configuration: %v\n", err)
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	if !cfg.IsAuthenticated() {
		fmt.Fprintf(os.Stderr, "[ERROR] Not authenticated. Please run 'agentbay login' first\n")
		return fmt.Errorf("not authenticated. Please run 'agentbay login' first")
	}

	// Create API client
	apiClient := agentbay.NewClientFromConfig(cfg)

	// Use longer timeout for status check (not for the full polling)
	statusCtx, statusCancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer statusCancel()

	// Check current image status and type using GetMcpImageInfo
	fmt.Printf("Checking current image status...")
	imageInfo, err := GetImageInfo(statusCtx, apiClient, imageId)
	if err != nil {
		fmt.Printf(" Failed.\n")
		return fmt.Errorf("failed to get image info: %w", err)
	}
	fmt.Printf(" Done.\n")
	fmt.Printf("[INFO] Image Type: %s\n", imageInfo.ImageType)
	fmt.Printf("[INFO] Current Status: %s\n", TranslateImageResourceStatus(imageInfo.ResourceStatus))

	// Check if this is a System image
	if IsSystemImage(imageInfo.ImageType) {
		fmt.Printf("[INFO] This is a System image.\n")
		fmt.Printf("[INFO] System images are always available and do not need to be activated.\n")
		fmt.Printf("[INFO] You can use this image directly without activation.\n")
		fmt.Printf("[INFO] Image ID: %s\n", imageId)
		return nil
	}

	// Check if this is a User image
	if !IsUserImage(imageInfo.ImageType) {
		return fmt.Errorf("unknown image type: %s (expected 'User' or 'System')", imageInfo.ImageType)
	}

	// Check if image is already activated
	if IsActivated(imageInfo.ResourceStatus) {
		fmt.Printf("[OK] Image is already activated! No action needed.\n")
		fmt.Printf("[INFO] Image ID: %s\n", imageId)
		return nil
	}

	// Check if image is currently activating
	shouldCreateResourceGroup := true
	if IsActivating(imageInfo.ResourceStatus) {
		fmt.Printf("[INFO] Image is currently activating, waiting for completion...\n")
		shouldCreateResourceGroup = false
	} else if IsDeactivated(imageInfo.ResourceStatus) {
		// Image is deactivated, proceed with activation
		shouldCreateResourceGroup = true
	} else {
		// Image is in an unexpected state
		return fmt.Errorf("cannot activate image in current state: %s", TranslateImageResourceStatus(imageInfo.ResourceStatus))
	}

	// Create resource group if needed
	if shouldCreateResourceGroup {
		fmt.Printf("Creating resource group...")
		createReq := &client.CreateResourceGroupRequest{
			ImageId: dara.String(imageId),
		}

		// Add CPU and Memory if specified
		if cpu > 0 {
			log.Debugf("[DEBUG] Setting Cpu to %d", cpu)
			createReq.SetCpu(int32(cpu))
			log.Debugf("[DEBUG] After SetCpu, createReq.Cpu = %v", createReq.Cpu)
		}
		if memory > 0 {
			log.Debugf("[DEBUG] Setting Memory to %d", memory)
			createReq.SetMemory(int32(memory))
			log.Debugf("[DEBUG] After SetMemory, createReq.Memory = %v", createReq.Memory)
		}

		// Debug: Print request details
		if log.GetLevel() >= log.DebugLevel {
			log.Debugf("[DEBUG] CreateResourceGroup Request:")
			if createReq.ImageId != nil {
				log.Debugf("[DEBUG] - ImageId: %s", *createReq.ImageId)
			}
			if createReq.Cpu != nil {
				log.Debugf("[DEBUG] - Cpu: %d", *createReq.Cpu)
			} else {
				log.Debugf("[DEBUG] - Cpu: nil")
			}
			if createReq.Memory != nil {
				log.Debugf("[DEBUG] - Memory: %d", *createReq.Memory)
			} else {
				log.Debugf("[DEBUG] - Memory: nil")
			}
		}

		// Use a separate context for the create operation
		createCtx, createCancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer createCancel()

		createResp, err := apiClient.CreateResourceGroup(createCtx, createReq)
		if err != nil {
			fmt.Printf(" Failed.\n")
			return fmt.Errorf("failed to create resource group: %w", err)
		}

		// Check response
		if createResp.Body == nil {
			fmt.Printf(" Failed.\n")
			return fmt.Errorf("invalid response from server")
		}

		// Log Request ID for debugging
		if createResp.Body.GetRequestId() != nil {
			log.Debugf("[DEBUG] CreateResourceGroup Request ID: %s", *createResp.Body.GetRequestId())
		}

		success := createResp.Body.GetSuccess()
		if success == nil || !*success {
			fmt.Printf(" Failed.\n")
			code := createResp.Body.GetCode()
			message := createResp.Body.GetMessage()
			if code != nil && message != nil {
				if log.GetLevel() >= log.DebugLevel && createResp.Body.GetRequestId() != nil {
					return fmt.Errorf("failed to create resource group: %s - %s (Request ID: %s)", *code, *message, *createResp.Body.GetRequestId())
				}
				return fmt.Errorf("failed to create resource group: %s - %s", *code, *message)
			}
			if log.GetLevel() >= log.DebugLevel && createResp.Body.GetRequestId() != nil {
				return fmt.Errorf("failed to create resource group (Request ID: %s)", *createResp.Body.GetRequestId())
			}
			return fmt.Errorf("failed to create resource group")
		}

		fmt.Printf(" Done.\n")
	}

	// Poll for activation completion
	fmt.Printf("Waiting for activation to complete...\n")
	pollingCtx := context.Background() // Don't use timeout context, polling has its own timeout
	config := DefaultActivatePollingConfig()

	if err := PollForActivation(pollingCtx, apiClient, imageId, config); err != nil {
		return fmt.Errorf("activation failed: %w", err)
	}

	fmt.Printf("[SUCCESS] Image activated successfully!\n")
	fmt.Printf("[INFO] Image ID: %s\n", imageId)

	return nil
}

func runImageDeactivate(cmd *cobra.Command, args []string) error {
	imageId := args[0]

	fmt.Printf("[DEACTIVATE] Deactivating image '%s'...\n", imageId)

	// Load configuration and check authentication
	cfg, err := config.GetConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Failed to load configuration: %v\n", err)
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	if !cfg.IsAuthenticated() {
		fmt.Fprintf(os.Stderr, "[ERROR] Not authenticated. Please run 'agentbay login' first\n")
		return fmt.Errorf("not authenticated. Please run 'agentbay login' first")
	}

	// Create API client
	apiClient := agentbay.NewClientFromConfig(cfg)

	// Use longer timeout for status check (not for the full polling)
	statusCtx, statusCancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer statusCancel()

	// Check current image status and type using GetMcpImageInfo
	fmt.Printf("Checking current image status...")
	imageInfo, err := GetImageInfo(statusCtx, apiClient, imageId)
	if err != nil {
		fmt.Printf(" Failed.\n")
		return fmt.Errorf("failed to get image info: %w", err)
	}
	fmt.Printf(" Done.\n")
	fmt.Printf("[INFO] Image Type: %s\n", imageInfo.ImageType)
	fmt.Printf("[INFO] Current Status: %s\n", TranslateImageResourceStatus(imageInfo.ResourceStatus))

	// Check if this is a System image
	if IsSystemImage(imageInfo.ImageType) {
		fmt.Printf("[INFO] This is a System image.\n")
		fmt.Printf("[INFO] System images cannot be deactivated as they are always available.\n")
		fmt.Printf("[INFO] Image ID: %s\n", imageId)
		return nil
	}

	// Check if this is a User image
	if !IsUserImage(imageInfo.ImageType) {
		return fmt.Errorf("unknown image type: %s (expected 'User' or 'System')", imageInfo.ImageType)
	}

	// Check if image is already deactivated
	if IsDeactivated(imageInfo.ResourceStatus) {
		fmt.Printf("[OK] Image is already deactivated! No action needed.\n")
		fmt.Printf("[INFO] Image ID: %s\n", imageId)
		return nil
	}

	// Check if image is currently deactivating
	shouldDeleteResourceGroup := true
	if IsDeactivating(imageInfo.ResourceStatus) {
		fmt.Printf("[INFO] Image is currently deactivating, waiting for completion...\n")
		shouldDeleteResourceGroup = false
	} else if IsActivated(imageInfo.ResourceStatus) {
		// Image is activated, proceed with deactivation
		shouldDeleteResourceGroup = true
	} else {
		// Image is in an unexpected state
		return fmt.Errorf("cannot deactivate image in current state: %s", TranslateImageResourceStatus(imageInfo.ResourceStatus))
	}

	// Delete resource group if needed
	if shouldDeleteResourceGroup {
		fmt.Printf("Deleting resource group...")
		deleteReq := &client.DeleteResourceGroupRequest{}
		deleteReq.SetImageId(imageId)

		// Debug: Print request details
		if log.GetLevel() >= log.DebugLevel {
			log.Debugf("[DEBUG] DeleteResourceGroup Request:")
			log.Debugf("[DEBUG] - ImageId: %s", imageId)
		}

		// Use a separate context for the delete operation
		deleteCtx, deleteCancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer deleteCancel()

		deleteResp, err := apiClient.DeleteResourceGroup(deleteCtx, deleteReq)
		if err != nil {
			fmt.Printf(" Failed.\n")
			log.Debugf("[DEBUG] DeleteResourceGroup API call failed: %v", err)
			return fmt.Errorf("failed to delete resource group: %w", err)
		}

		// Check response
		if deleteResp.Body == nil {
			fmt.Printf(" Failed.\n")
			return fmt.Errorf("invalid response from server")
		}

		// Log Request ID for debugging
		if deleteResp.Body.GetRequestId() != nil {
			log.Debugf("[DEBUG] DeleteResourceGroup Request ID: %s", *deleteResp.Body.GetRequestId())
		}

		success := deleteResp.Body.GetSuccess()
		if success == nil || !*success {
			fmt.Printf(" Failed.\n")
			code := deleteResp.Body.GetCode()
			message := deleteResp.Body.GetMessage()
			if code != nil && message != nil {
				if log.GetLevel() >= log.DebugLevel && deleteResp.Body.GetRequestId() != nil {
					return fmt.Errorf("failed to delete resource group: %s - %s (Request ID: %s)", *code, *message, *deleteResp.Body.GetRequestId())
				}
				return fmt.Errorf("failed to delete resource group: %s - %s", *code, *message)
			}
			if log.GetLevel() >= log.DebugLevel && deleteResp.Body.GetRequestId() != nil {
				return fmt.Errorf("failed to delete resource group (Request ID: %s)", *deleteResp.Body.GetRequestId())
			}
			return fmt.Errorf("failed to delete resource group")
		}

		fmt.Printf(" Done.\n")
	}

	// Poll for deactivation completion
	fmt.Printf("Waiting for deactivation to complete...\n")
	pollingCtx := context.Background() // Don't use timeout context, polling has its own timeout
	config := DefaultDeactivatePollingConfig()

	if err := PollForDeactivation(pollingCtx, apiClient, imageId, config); err != nil {
		return fmt.Errorf("deactivation failed: %w", err)
	}

	fmt.Printf("[SUCCESS] Image deactivated successfully!\n")
	fmt.Printf("[INFO] Image ID: %s\n", imageId)

	return nil
}

// pollImageDeactivationStatus polls the image deactivation status until completion or failure
func pollImageDeactivationStatus(ctx context.Context, apiClient agentbay.Client, imageId string) error {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	// Create a new context with longer timeout for polling
	pollCtx, pollCancel := context.WithTimeout(context.Background(), 45*time.Minute)
	defer pollCancel()

	for {
		select {
		case <-pollCtx.Done():
			fmt.Printf("[INFO] Image ID: %s\n", imageId)
			return fmt.Errorf("timeout waiting for image deactivation to complete")
		case <-ticker.C:
			// Query specific image status using ListMcpImages
			listReq := &client.ListMcpImagesRequest{
				ImageType: dara.String("User"),
				PageSize:  dara.Int32(100),
			}

			listResp, err := apiClient.ListMcpImages(pollCtx, listReq)
			if err != nil {
				fmt.Printf("[WARN] Warning: Failed to check image status: %v\n", err)
				fmt.Printf("[INFO] Image ID: %s\n", imageId)
				continue // Continue polling on API errors
			}

			if listResp.Body == nil || listResp.Body.Data == nil {
				fmt.Printf("[WARN] Warning: Invalid response format\n")
				fmt.Printf("[INFO] Image ID: %s\n", imageId)
				continue
			}

			// Find the specified image
			var targetImage *client.ListMcpImagesResponseBodyData
			for _, image := range listResp.Body.Data {
				if image.ImageId != nil && *image.ImageId == imageId {
					targetImage = image
					break
				}
			}

			if targetImage == nil {
				fmt.Printf("[WARN] Warning: Image not found: %s\n", imageId)
				continue // Continue polling
			}

			status := getStringValue(targetImage.ImageResourceStatus)
			formattedStatus := formatImageStatus(status)

			fmt.Printf("[STATUS] Status: %s\n", formattedStatus)

			switch status {
			case "IMAGE_AVAILABLE":
				fmt.Printf("[SUCCESS] Image deactivated successfully! Image ID: %s\n", imageId)
				fmt.Printf("[INFO] Final Status: %s\n", formattedStatus)
				return nil
			case "RESOURCE_FAILED":
				fmt.Printf("[INFO] Image ID: %s\n", imageId)
				if listResp.Body.GetRequestId() != nil {
					fmt.Printf("[INFO] Request ID: %s\n", *listResp.Body.GetRequestId())
				}
				return fmt.Errorf("image deactivation failed with status: %s", formattedStatus)
			case "RESOURCE_DELETING":
				// Continue polling - deactivation in progress
				continue
			case "RESOURCE_PUBLISHED":
				// Image is still activated, continue polling in case deactivation is delayed
				fmt.Printf("[REFRESH] Image still activated, continuing to monitor deactivation...\n")
				continue
			default:
				fmt.Printf("[REFRESH] Unknown status '%s', continuing to monitor...\n", formattedStatus)
				continue
			}
		}
	}
}
