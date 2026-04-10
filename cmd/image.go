// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/agentbay/agentbay-cli/internal/agentbay"
	"github.com/agentbay/agentbay-cli/internal/client"
	"github.com/agentbay/agentbay-cli/internal/config"
	"github.com/alibabacloud-go/tea/dara"
)

// printErrorMessage returns a multi-line error for Cobra to print once (avoids duplicate stderr + "Error:" lines).
func printErrorMessage(lines ...string) error {
	return fmt.Errorf("%s", strings.Join(lines, "\n"))
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
  2c4g  - 2 CPU cores with 4 GB memory (default)
  4c8g  - 4 CPU cores with 8 GB memory
  8c16g - 8 CPU cores with 16 GB memory

If no CPU/memory is specified, 2c4g (2 CPU, 4 GB memory) will be used by default.

Examples:
  # Activate with default resources (2c4g)
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

var imageInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Download a Dockerfile template from the cloud",
	Long: `Download a Dockerfile template from the cloud to the local root directory.

This command fetches a Dockerfile template from AgentBay and saves it as 'Dockerfile' 
in the current directory. The source image ID must be specified via --sourceImageId flag.

Examples:
  # Download Dockerfile template with source image ID
  agentbay image init --sourceImageId code-space-debian-12
  
  # Short form
  agentbay image init -i code-space-debian-12`,
	Args: cobra.NoArgs,
	RunE: runImageInit,
}

var imageStatusCmd = &cobra.Command{
	Use:   "status <image-id>",
	Short: "Show resource status for an image",
	Long: `Query lifecycle / deployment status for an image by ID (GetMcpImageInfo).

This is the image resource status used by activate/deactivate, not the Docker build
task status shown during 'agentbay image create'.

Common resource status values (API):
  IMAGE_CREATING       — Image is being created
  IMAGE_CREATE_FAILED  — Image creation failed
  IMAGE_AVAILABLE      — Available, not activated (deactivated)
  RESOURCE_DEPLOYING   — Activation in progress
  RESOURCE_PUBLISHED   — Activated
  RESOURCE_DELETING    — Deactivation in progress
  RESOURCE_FAILED      — Activation or resource operation failed
  RESOURCE_CEASED      — Resource ceased

Examples:
  agentbay image status imgc-xxxxxxxxxxxxxx`,
	Args: cobra.ExactArgs(1),
	RunE: runImageStatus,
}

func init() {
	// Add flags to image create command
	imageCreateCmd.Flags().StringP("dockerfile", "f", "", "Path to the Dockerfile (required)")
	imageCreateCmd.Flags().StringP("imageId", "i", "", "Source image ID to build from (required)")

	// Mark required flags
	imageCreateCmd.MarkFlagRequired("dockerfile")
	imageCreateCmd.MarkFlagRequired("imageId")

	// Add flags to image activate command
	imageActivateCmd.Flags().IntP("cpu", "c", 0, "CPU cores (2, 4, or 8; default: 2 when not specified)")
	imageActivateCmd.Flags().IntP("memory", "m", 0, "Memory in GB (4, 8, or 16; default: 4 when not specified)")

	// Add flags to image list command
	imageListCmd.Flags().StringP("os-type", "o", "", "Filter by OS type: Linux, Android, or Windows (optional)")
	imageListCmd.Flags().Bool("include-system", false, "Include system images in addition to user images")
	imageListCmd.Flags().Bool("system-only", false, "Show only system images")
	imageListCmd.Flags().IntP("page", "p", 1, "Page number (default: 1)")
	imageListCmd.Flags().IntP("size", "s", 10, "Page size (default: 10)")

	// Add required flag for image init command - use sourceImageId to match API field name
	imageInitCmd.Flags().StringP("sourceImageId", "i", "", "Source image ID (required)")

	// Mark required flag
	imageInitCmd.MarkFlagRequired("sourceImageId")

	// Add subcommands to image command
	ImageCmd.AddCommand(imageCreateCmd)
	ImageCmd.AddCommand(imageListCmd)
	ImageCmd.AddCommand(imageActivateCmd)
	ImageCmd.AddCommand(imageDeactivateCmd)
	ImageCmd.AddCommand(imageInitCmd)
	ImageCmd.AddCommand(imageStatusCmd)
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

	dockerfileContent, err := os.ReadFile(dockerfilePath)
	if err != nil {
		return fmt.Errorf("failed to read Dockerfile: %w", err)
	}
	contextDir := filepath.Dir(dockerfilePath)
	addCopyFiles, err := ParseCOPYADDSources(dockerfileContent, contextDir)
	if err != nil {
		return err
	}
	if err := ValidateCopyAddSourceFileSizes(contextDir, addCopyFiles); err != nil {
		return printErrorMessage(
			fmt.Sprintf("[ERROR] COPY/ADD file too large: %v", err),
			"",
			"[TIP] Each file referenced by COPY or ADD must be at most 1 MB (1,048,576 bytes).",
		)
	}

	fmt.Printf("[BUILD] Creating image '%s'...\n", imageName)

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
				"[ERROR] Authentication failed. Please run 'agentbay login' first.",
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

	fmt.Printf("[STEP 1/4] Getting upload credentials...\n")
	sourceAgentBay := "AgentBay"
	credReq := &client.GetDockerFileStoreCredentialRequest{
		Source:       &sourceAgentBay,
		FilePath:     dara.String("Dockerfile"),
		IsDockerfile: dara.String("true"),
	}
	if log.GetLevel() >= log.DebugLevel {
		log.Debugf("[DEBUG] GetDockerFileStoreCredential Request: Source=%s FilePath=%s IsDockerfile=%s", *credReq.Source, *credReq.FilePath, *credReq.IsDockerfile)
	}
	fmt.Printf("Requesting upload credentials...")
	credResp, err := apiClient.GetDockerFileStoreCredential(ctx, credReq)
	if err != nil {
		log.Debugf("[DEBUG] GetDockerFileStoreCredential API call failed: %v", err)
		if log.GetLevel() >= log.DebugLevel {
			fmt.Printf("[DEBUG] Error details: %v\n", err)
		}
		return fmt.Errorf("[ERROR] Failed to get upload credentials. Please check your authentication and try again: %w", err)
	}
	fmt.Printf(" Done.\n")
	if credResp.Body == nil || credResp.Body.Data == nil {
		return fmt.Errorf("invalid response: missing upload credentials")
	}
	ossUrl := credResp.Body.Data.GetOssUrl()
	taskId := credResp.Body.Data.GetTaskId()
	if ossUrl == nil || taskId == nil {
		return fmt.Errorf("invalid response: missing OSS URL or task ID")
	}

	fmt.Printf("[STEP 2/4] Uploading Dockerfile...\n")
	fmt.Printf("Uploading file...")
	if err = uploadFileToOSS(dockerfilePath, *ossUrl); err != nil {
		if log.GetLevel() >= log.DebugLevel {
			fmt.Printf("[DEBUG] Error details: %v\n", err)
		}
		return fmt.Errorf("[ERROR] Failed to upload Dockerfile. Please check your network connection and try again: %w", err)
	}
	fmt.Printf(" Done.\n")

	if len(addCopyFiles) > 0 {
		fmt.Printf("[STEP 3/4] Uploading ADD/COPY files (%d files)...\n", len(addCopyFiles))
		type fileItem struct{ absPath, relPath string }
		var files []fileItem
		for _, absPath := range addCopyFiles {
			relPath, err := RelativePathForUpload(contextDir, absPath)
			if err != nil {
				return fmt.Errorf("failed to get relative path for %s: %w", absPath, err)
			}
			files = append(files, fileItem{absPath: absPath, relPath: relPath})
		}
		const maxConcurrent = 10
		sem := make(chan struct{}, maxConcurrent)
		var credsMu sync.Mutex
		creds := make(map[string]string)
		var firstCredErr error
		var credWg sync.WaitGroup
		fmt.Printf("Requesting upload credentials for %d files (parallel)...\n", len(files))
		for _, f := range files {
			f := f
			credWg.Add(1)
			go func() {
				defer credWg.Done()
				sem <- struct{}{}
				defer func() { <-sem }()
				credReq := &client.GetDockerFileStoreCredentialRequest{
					Source:       &sourceAgentBay,
					FilePath:     &f.relPath,
					IsDockerfile: dara.String("false"),
					TaskId:       taskId,
				}
				resp, err := apiClient.GetDockerFileStoreCredential(ctx, credReq)
				if err != nil {
					credsMu.Lock()
					if firstCredErr == nil {
						firstCredErr = fmt.Errorf("failed to get upload credentials for %s: %w", f.relPath, err)
					}
					credsMu.Unlock()
					return
				}
				if resp.Body == nil || resp.Body.Data == nil {
					credsMu.Lock()
					if firstCredErr == nil {
						firstCredErr = fmt.Errorf("invalid response: missing upload credentials for %s", f.relPath)
					}
					credsMu.Unlock()
					return
				}
				ossUrl := resp.Body.Data.GetOssUrl()
				if ossUrl == nil || *ossUrl == "" {
					credsMu.Lock()
					if firstCredErr == nil {
						firstCredErr = fmt.Errorf("invalid response: missing OSS URL for %s", f.relPath)
					}
					credsMu.Unlock()
					return
				}
				credsMu.Lock()
				creds[f.absPath] = *ossUrl
				credsMu.Unlock()
			}()
		}
		credWg.Wait()
		if firstCredErr != nil {
			return fmt.Errorf("[ERROR] %w", firstCredErr)
		}
		var firstUploadErr error
		var uploadWg sync.WaitGroup
		fmt.Printf("Uploading %d files (parallel)...\n", len(files))
		for _, f := range files {
			f := f
			ossUrl := creds[f.absPath]
			uploadWg.Add(1)
			go func() {
				defer uploadWg.Done()
				sem <- struct{}{}
				defer func() { <-sem }()
				if err := uploadFileToOSS(f.absPath, ossUrl); err != nil {
					credsMu.Lock()
					if firstUploadErr == nil {
						firstUploadErr = fmt.Errorf("failed to upload %s: %w", f.relPath, err)
					}
					credsMu.Unlock()
				}
			}()
		}
		uploadWg.Wait()
		if firstUploadErr != nil {
			return fmt.Errorf("[ERROR] %w", firstUploadErr)
		}
		fmt.Printf(" Done.\n")
	}

	fmt.Printf("[STEP 4/4] Creating Docker image task...\n")

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
		if log.GetLevel() >= log.DebugLevel {
			fmt.Printf("[DEBUG] Error details: %v\n", err)
		}
		// Try to extract Request ID from response if available
		if createResp != nil && createResp.Body != nil && createResp.Body.GetRequestId() != nil {
			fmt.Printf("[DEBUG] Request ID: %s\n", *createResp.Body.GetRequestId())
		}
		return fmt.Errorf("[ERROR] Failed to create Docker image task. Please try again: %w", err)
	}
	fmt.Printf(" Done.\n")
	if createResp.Body != nil && createResp.Body.GetRequestId() != nil {
		printRequestIDIfVerbose(cmd, *createResp.Body.GetRequestId())
	}

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
				// Check if this is a Dockerfile validation error
				isValidationError := false
				if taskMsg != nil && *taskMsg != "" {
					isValidationError = isDockerfileValidationError(*taskMsg)
				}

				if isValidationError {
					lines := []string{
						"[ERROR] ❌ Dockerfile validation failed",
					}
					if taskMsg != nil && *taskMsg != "" {
						lines = append(lines, "[ERROR] Validation error: "+*taskMsg)
					}
					lines = append(lines,
						"[TIP] Please check your Dockerfile and ensure you haven't modified system-defined lines.",
						"[TIP] Use 'agentbay image init' to download a valid template.",
					)
					if taskResp.Body.GetRequestId() != nil {
						lines = append(lines, fmt.Sprintf("[DEBUG] Request ID: %s", *taskResp.Body.GetRequestId()))
					}
					lines = append(lines, fmt.Sprintf("[DOC] Task ID: %s", *finalTaskId))
					return printErrorMessage(lines...)
				}
				lines := []string{"[ERROR] ❌ Image build failed"}
				if taskMsg != nil && *taskMsg != "" {
					lines = append(lines, "[ERROR] Error details: "+*taskMsg)
				}
				if taskResp.Body.GetRequestId() != nil {
					lines = append(lines, fmt.Sprintf("[DEBUG] Request ID: %s", *taskResp.Body.GetRequestId()))
				}
				lines = append(lines, fmt.Sprintf("[DOC] Task ID: %s", *finalTaskId))
				return printErrorMessage(lines...)
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
	fmt.Print(fetchMessage)

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
		if log.GetLevel() >= log.DebugLevel {
			fmt.Printf("[DEBUG] Error details: %v\n", err)
		}
		return fmt.Errorf("[ERROR] Failed to fetch image list. Please check your authentication and try again: %w", err)
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
		return fmt.Errorf("both CPU and memory must be specified together. Supported combinations: 2c4g (--cpu 2 --memory 4), 4c8g (--cpu 4 --memory 8), 8c16g (--cpu 8 --memory 16)")
	}

	// Check supported combinations
	validCombos := map[int]int{
		2: 4,  // 2c4g
		4: 8,  // 4c8g
		8: 16, // 8c16g
	}

	expectedMemory, exists := validCombos[cpu]
	if !exists || expectedMemory != memory {
		return fmt.Errorf("invalid CPU/Memory combination: %dc%dg. Supported combinations: 2c4g (--cpu 2 --memory 4), 4c8g (--cpu 4 --memory 8), 8c16g (--cpu 8 --memory 16)", cpu, memory)
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

// ossPutContent performs a single PUT of content to a pre-signed OSS URL.
func ossPutContent(ossUrl string, content []byte) (statusCode int, respBody string, transportErr error) {
	req, err := http.NewRequest(http.MethodPut, ossUrl, bytes.NewReader(content))
	if err != nil {
		return 0, "", err
	}
	// OSS may return 307/308 redirect; without GetBody the redirect request sends an empty body and the object is stored empty.
	req.GetBody = func() (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewReader(content)), nil
	}
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("User-Agent", "AgentBay-CLI/1.0")
	req.ContentLength = int64(len(content))
	httpClient := &http.Client{Timeout: 60 * time.Second}
	resp, err := httpClient.Do(req)
	if err != nil {
		return 0, "", err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return resp.StatusCode, string(body), nil
}

func uploadFileToOSS(localPath, ossUrl string) error {
	log.Debugf("[DEBUG] Starting file upload: %s", localPath)
	content, err := os.ReadFile(localPath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}
	cfg := client.DefaultRetryConfig()
	delay := cfg.InitialDelay
	var lastErr error
	for attempt := 0; attempt <= cfg.MaxRetries; attempt++ {
		if attempt > 0 {
			if log.GetLevel() >= log.DebugLevel {
				log.Debugf("[DEBUG] OSS upload retry %d/%d after %v: %v", attempt+1, cfg.MaxRetries+1, delay, lastErr)
			}
			time.Sleep(delay)
			nd := time.Duration(float64(delay) * cfg.BackoffFactor)
			if nd > cfg.MaxDelay {
				nd = cfg.MaxDelay
			}
			delay = nd
		}
		status, body, transportErr := ossPutContent(ossUrl, content)
		if transportErr != nil {
			lastErr = fmt.Errorf("failed to upload: %w", transportErr)
			if !client.IsRetryableError(transportErr) || attempt == cfg.MaxRetries {
				return lastErr
			}
			continue
		}
		if status >= 200 && status < 300 {
			return nil
		}
		lastErr = fmt.Errorf("upload failed with status %d: %s", status, body)
		if !client.IsRetryableHTTPStatus(status) || attempt == cfg.MaxRetries {
			return lastErr
		}
	}
	return lastErr
}

// DefaultActivateCPU and DefaultActivateMemory are the default resource allocation when user does not specify --cpu/--memory
const DefaultActivateCPU = 2
const DefaultActivateMemory = 4

func runImageActivate(cmd *cobra.Command, args []string) error {
	imageId := args[0]
	cpu, _ := cmd.Flags().GetInt("cpu")
	memory, _ := cmd.Flags().GetInt("memory")

	// Apply default 2c4g when not specified
	if cpu == 0 && memory == 0 {
		cpu = DefaultActivateCPU
		memory = DefaultActivateMemory
	}

	// Validate CPU and memory combination
	if err := ValidateCPUMemoryCombo(cpu, memory); err != nil {
		return printCPUMemoryValidationError(err)
	}

	fmt.Printf("[ACTIVATE] Activating image '%s'...\n", imageId)
	fmt.Printf("[RESOURCE] CPU: %d cores, Memory: %d GB\n", cpu, memory)

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
		return fmt.Errorf("[ERROR] Failed to load configuration: %w", err)
	}

	if !cfg.IsAuthenticated() {
		return config.ErrNotAuthenticated()
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
	} else if IsFailed(imageInfo.ResourceStatus) {
		// Activation failed - cannot delete without ResourceGroupId
		fmt.Printf("[INFO] Image is in Activation Failed state.\n")
		fmt.Printf("[INFO] The image may recover automatically to Available state. Please try again later.\n")
		fmt.Printf("[INFO] Alternatively, use the web console to deactivate this image.\n")
		return nil
	} else {
		// Image is in an unexpected state
		return fmt.Errorf("cannot deactivate image in current state: %s", TranslateImageResourceStatus(imageInfo.ResourceStatus))
	}

	// Delete resource group if needed
	if shouldDeleteResourceGroup {
		fmt.Printf("Fetching resource group info...")
		resourceGroupId, err := GetResourceGroupIdForImage(statusCtx, apiClient, imageId)
		if err != nil {
			fmt.Printf(" Failed.\n")
			log.Debugf("[DEBUG] GetResourceGroupIdForImage failed: %v", err)
			return fmt.Errorf("failed to get resource group info: %w", err)
		}
		fmt.Printf(" Done.\n")

		if resourceGroupId == "" {
			fmt.Printf("[WARN] Could not find ResourceGroupId for this image.\n")
			fmt.Printf("[INFO] The image may recover automatically to Available state. Please try again later.\n")
			fmt.Printf("[INFO] Alternatively, use the web console to deactivate this image.\n")
			return nil
		}

		fmt.Printf("Deleting resource group...")
		deleteReq := &client.DeleteResourceGroupRequest{}
		deleteReq.SetImageId(imageId)
		deleteReq.SetResourceGroupId(resourceGroupId)

		// Debug: Print request details
		if log.GetLevel() >= log.DebugLevel {
			log.Debugf("[DEBUG] DeleteResourceGroup Request:")
			log.Debugf("[DEBUG] - ImageId: %s", imageId)
			log.Debugf("[DEBUG] - ResourceGroupId: %s", resourceGroupId)
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

func runImageInit(cmd *cobra.Command, args []string) error {
	fmt.Printf("[INIT] Downloading Dockerfile template...\n")

	// Source is always AgentBay
	source := "AgentBay"

	// Get sourceImageId from command line flag (required)
	sourceImageId, _ := cmd.Flags().GetString("sourceImageId")
	if sourceImageId == "" {
		return printErrorMessage(
			"[ERROR] Missing required flag: --sourceImageId",
			"",
			"[TIP] Usage: agentbay image init --sourceImageId <image-id>",
			"[NOTE] Example: agentbay image init --sourceImageId code-space-debian-12",
			"[NOTE] Short form: agentbay image init -i code-space-debian-12",
		)
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

	// Prepare request - Source and SourceImageId are required
	req := &client.GetDockerfileTemplateRequest{
		Source:        &source,
		SourceImageId: &sourceImageId,
	}

	// Debug: Print request details
	if log.GetLevel() >= log.DebugLevel {
		log.Debugf("[DEBUG] GetDockerfileTemplate Request:")
		if req.Source != nil {
			log.Debugf("[DEBUG] - Source: %s", *req.Source)
		}
		if req.SourceImageId != nil {
			log.Debugf("[DEBUG] - SourceImageId: %s", *req.SourceImageId)
		}
	}

	// Make API call to get Dockerfile template
	fmt.Printf("Requesting Dockerfile template...")
	resp, err := apiClient.GetDockerfileTemplate(ctx, req)
	if err != nil {
		log.Debugf("[DEBUG] GetDockerfileTemplate API call failed: %v", err)

		return fmt.Errorf("[ERROR] Failed to get Dockerfile template: %w", err)
	}
	fmt.Printf(" Done.\n")

	// Validate response
	if resp.Body == nil || resp.Body.Data == nil {
		return fmt.Errorf("invalid response: missing template data")
	}

	var dockerfileContent []byte

	// Prefer DockerfileContent if available, otherwise fall back to OSS download
	dockerfileContentStr := resp.Body.Data.GetDockerfileContent()
	if dockerfileContentStr != nil && *dockerfileContentStr != "" {
		log.Debugf("[DEBUG] Using DockerfileContent from response")
		dockerfileContent = []byte(*dockerfileContentStr)
	} else {
		// Fall back to OSS download
		ossUrl := resp.Body.Data.GetOssDownloadUrl()
		if ossUrl == nil || *ossUrl == "" {
			return fmt.Errorf("invalid response: missing both DockerfileContent and OSS download URL")
		}

		log.Debugf("[DEBUG] OSS Download URL: %s", *ossUrl)

		// Download Dockerfile from OSS URL
		fmt.Printf("Downloading Dockerfile from OSS...")
		var err error
		dockerfileContent, err = downloadDockerfileFromOSS(*ossUrl)
		if err != nil {
			fmt.Printf(" Failed.\n")
			return fmt.Errorf("failed to download Dockerfile from OSS: %w", err)
		}
		fmt.Printf(" Done.\n")
	}

	// Get NonEditLineNum if available
	nonEditLineNum := resp.Body.Data.GetNonEditLineNum()
	if nonEditLineNum != nil {
		log.Debugf("[DEBUG] NonEditLineNum: %d", *nonEditLineNum)
	} else {
		log.Debugf("[DEBUG] NonEditLineNum is nil or not present in response")
	}

	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %w", err)
	}

	// Write Dockerfile to current directory
	dockerfilePath := filepath.Join(cwd, "Dockerfile")
	fmt.Printf("Writing Dockerfile to %s...", dockerfilePath)

	// Check if Dockerfile already exists
	if _, err := os.Stat(dockerfilePath); err == nil {
		fmt.Printf("\n[WARN] Dockerfile already exists at %s\n", dockerfilePath)
		fmt.Printf("[INFO] The existing file will be overwritten.\n")
	}

	// Write the content to file
	err = os.WriteFile(dockerfilePath, dockerfileContent, 0644)
	if err != nil {
		fmt.Printf(" Failed.\n")
		return fmt.Errorf("failed to write Dockerfile: %w", err)
	}
	fmt.Printf(" Done.\n")

	fmt.Printf("[SUCCESS] ✅ Dockerfile template downloaded successfully!\n")
	fmt.Printf("[INFO] Dockerfile saved to: %s\n", dockerfilePath)

	// Display non-editable lines information if available
	if nonEditLineNum != nil && *nonEditLineNum > 0 {
		fmt.Printf("[IMPORTANT] The first %d line(s) of the Dockerfile are system-defined and cannot be modified.\n", *nonEditLineNum)
		fmt.Printf("[IMPORTANT] Please only modify content after line %d.\n", *nonEditLineNum)
	}

	return nil
}

func runImageStatus(_ *cobra.Command, args []string) error {
	imageId := args[0]

	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to load configuration: %w", err)
	}
	if !cfg.IsAuthenticated() {
		return config.ErrNotAuthenticated()
	}

	apiClient := agentbay.NewClientFromConfig(cfg)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	fmt.Printf("[STATUS] Querying image '%s'...\n", imageId)
	imageInfo, err := GetImageInfo(ctx, apiClient, imageId)
	if err != nil {
		return fmt.Errorf("failed to get image info: %w", err)
	}

	typeLabel := imageInfo.ImageType
	if typeLabel == "" {
		typeLabel = "(unknown)"
	}

	fmt.Printf("[INFO] Image ID: %s\n", imageId)
	fmt.Printf("[INFO] Image type: %s\n", typeLabel)
	fmt.Printf("[INFO] Resource status (API): %s\n", imageInfo.ResourceStatus)
	fmt.Printf("[INFO] Resource status (display): %s\n", TranslateImageResourceStatus(imageInfo.ResourceStatus))
	fmt.Printf("[INFO] Deployment: %s\n", summarizeDeploymentState(imageInfo.ResourceStatus))

	if IsSystemImage(imageInfo.ImageType) {
		fmt.Printf("[NOTE] System images do not use activate/deactivate; status is informational.\n")
	}

	return nil
}

func summarizeDeploymentState(resourceStatus string) string {
	switch {
	case IsActivated(resourceStatus):
		return "Activated"
	case IsDeactivated(resourceStatus):
		return "Not activated"
	case IsActivating(resourceStatus):
		return "Activating"
	case IsDeactivating(resourceStatus):
		return "Deactivating"
	case IsFailed(resourceStatus):
		return "Failed"
	default:
		return TranslateImageResourceStatus(resourceStatus)
	}
}

// ossGetOnce performs a single GET from an OSS URL and returns the full body on success.
func ossGetOnce(ossUrl string) (statusCode int, content []byte, transportErr error) {
	req, err := http.NewRequest(http.MethodGet, ossUrl, nil)
	if err != nil {
		return 0, nil, err
	}
	req.Header.Set("User-Agent", "AgentBay-CLI/1.0")
	httpClient := &http.Client{Timeout: 60 * time.Second}
	resp, err := httpClient.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return resp.StatusCode, nil, readErr
	}
	return resp.StatusCode, body, nil
}

// downloadDockerfileFromOSS downloads Dockerfile content from OSS URL
func downloadDockerfileFromOSS(ossUrl string) ([]byte, error) {
	log.Debugf("[DEBUG] Downloading from OSS URL: %s", ossUrl)
	cfg := client.DefaultRetryConfig()
	delay := cfg.InitialDelay
	var lastErr error
	for attempt := 0; attempt <= cfg.MaxRetries; attempt++ {
		if attempt > 0 {
			if log.GetLevel() >= log.DebugLevel {
				log.Debugf("[DEBUG] OSS download retry %d/%d after %v: %v", attempt+1, cfg.MaxRetries+1, delay, lastErr)
			}
			time.Sleep(delay)
			nd := time.Duration(float64(delay) * cfg.BackoffFactor)
			if nd > cfg.MaxDelay {
				nd = cfg.MaxDelay
			}
			delay = nd
		}
		status, content, transportErr := ossGetOnce(ossUrl)
		if transportErr != nil {
			lastErr = fmt.Errorf("failed to download from OSS: %w", transportErr)
			if !client.IsRetryableError(transportErr) || attempt == cfg.MaxRetries {
				return nil, lastErr
			}
			continue
		}
		if status >= 200 && status < 300 {
			log.Debugf("[DEBUG] Downloaded %d bytes from OSS", len(content))
			return content, nil
		}
		lastErr = fmt.Errorf("download failed with status %d: %s", status, string(content))
		if !client.IsRetryableHTTPStatus(status) || attempt == cfg.MaxRetries {
			return nil, lastErr
		}
	}
	return nil, lastErr
}

// isDockerfileValidationError checks if the error message indicates a Dockerfile validation failure
func isDockerfileValidationError(taskMsg string) bool {
	// Check for the specific Dockerfile validation error message
	validationErrorMsg := "Image reference is Invalid. Please do not modify the image reference in Dockerfile"
	return strings.Contains(taskMsg, validationErrorMsg)
}
