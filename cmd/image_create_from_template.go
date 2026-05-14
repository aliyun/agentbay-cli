// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

// image_create_from_template.go implements "agentbay image create-from-template".
// Uses raw HTTP POP RPC V1 to call CreateImageFromTemplate.

package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/agentbay/agentbay-cli/internal/config"
)

var imageCreateFromTemplateCmd = &cobra.Command{
	Use:   "create-from-template",
	Short: "Create a custom image from a system image template",
	Long: `Create a custom image by specifying a source image (repository address with tag),
image name, and a system (template) image ID.

This calls the CreateImageFromTemplate POP Action.

Examples:
  agentbay image create-from-template --source-image registry.cn-hangzhou.aliyuncs.com/myrepo/myimage:v1.0 --name my-custom-image --imageId <id>
  agentbay image create-from-template -s registry.cn-hangzhou.aliyuncs.com/myrepo/myimage:v1.0 -n my-custom-image -i <id>`,
	Args: cobra.NoArgs,
	RunE: runImageCreateFromTemplate,
}

func init() {
	imageCreateFromTemplateCmd.Flags().StringP("source-image", "s", "", "Source image reference, e.g. registry.cn-hangzhou.aliyuncs.com/myrepo/myimage:v1.0 (required)")
	imageCreateFromTemplateCmd.Flags().StringP("name", "n", "", "Name for the custom image (required)")
	imageCreateFromTemplateCmd.Flags().StringP("imageId", "i", "", "System (template) image ID (required)")
	imageCreateFromTemplateCmd.MarkFlagRequired("source-image")
	imageCreateFromTemplateCmd.MarkFlagRequired("name")
	imageCreateFromTemplateCmd.MarkFlagRequired("imageId")

	ImageCmd.AddCommand(imageCreateFromTemplateCmd)
}

// --- Response model ---

type createFromTemplateResponse struct {
	RequestId      *string                         `json:"RequestId"`
	Code           *string                         `json:"Code"`
	Message        *string                         `json:"Message"`
	HttpStatusCode *int                            `json:"HttpStatusCode"`
	Success        *bool                           `json:"Success"`
	Data           *createFromTemplateResponseData `json:"Data"`
}

type createFromTemplateResponseData struct {
	ImageId *string `json:"ImageId"`
}

// --- Implementation ---

func runImageCreateFromTemplate(cmd *cobra.Command, args []string) error {
	sourceImage, _ := cmd.Flags().GetString("source-image")
	imageName, _ := cmd.Flags().GetString("name")
	templateImageId, _ := cmd.Flags().GetString("imageId")

	// Load cached ACR credential to validate source-image prefix
	cache, err := loadACRCredential()
	if err != nil {
		return fmt.Errorf("[ERROR] %w\nPlease run 'agentbay docker login' first to obtain registry credentials", err)
	}

	// Validate: source-image must start with the cached registry path prefix
	// Expected prefix: $RegistryUrl/$Namespace/$RepoName
	expectedPrefix := fmt.Sprintf("%s/%s/%s", cache.RegistryURL, cache.Namespace, cache.RepoName)
	if !strings.HasPrefix(sourceImage, expectedPrefix) {
		return fmt.Errorf("[ERROR] source-image '%s' does not match the authorized registry path.\n"+
			"  Expected prefix: %s\n"+
			"  Please use the image tagged via 'agentbay docker tag' command", sourceImage, expectedPrefix)
	}

	// Truncate: strip the registry URL, only send /$Namespace/$RepoName:<tag> to backend
	// e.g. "ai-container-pre-9543-registry.cn-hangzhou.cr.aliyuncs.com/customer_cli/1160165251879674:v1"
	//    → "/customer_cli/1160165251879674:v1"
	physicalImageId := sourceImage[len(cache.RegistryURL):]

	fmt.Println("[IMAGE] Creating custom image from template...")
	fmt.Printf("  SourceImage:      %s\n", sourceImage)
	fmt.Printf("  PhysicalImageId:  %s\n", physicalImageId)
	fmt.Printf("  Name:             %s\n", imageName)
	fmt.Printf("  ImageId:          %s\n", templateImageId)

	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to load configuration: %w", err)
	}
	if !cfg.IsAuthenticated() {
		return config.ErrNotAuthenticated()
	}

	client, err := newACSClientFromConfig(cfg)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create HTTP client: %w", err)
	}

	params := map[string]string{
		"PhysicalImageId": physicalImageId,
		"ImageName":       imageName,
		"TemplateImageId": templateImageId,
	}

	fmt.Printf("Requesting CreateImageFromTemplate...")
	body, statusCode, err := client.callRPC("CreateImageFromTemplate", params)
	if err != nil {
		return fmt.Errorf("[ERROR] Request failed: %w", err)
	}
	fmt.Printf(" Done. (HTTP %d)\n", statusCode)

	if statusCode < 200 || statusCode >= 300 {
		return fmt.Errorf("[ERROR] API returned HTTP %d: %s", statusCode, string(body))
	}

	var resp createFromTemplateResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return fmt.Errorf("[ERROR] Failed to parse response: %w (body=%s)", err, truncateBody(body, 512))
	}

	fmt.Printf("\n[RESPONSE]\n")
	fmt.Printf("  RequestId:      %s\n", ptrStr(resp.RequestId))
	fmt.Printf("  Code:           %s\n", ptrStr(resp.Code))
	fmt.Printf("  Message:        %s\n", ptrStr(resp.Message))
	if resp.Success != nil {
		fmt.Printf("  Success:        %t\n", *resp.Success)
	}
	if resp.HttpStatusCode != nil {
		fmt.Printf("  HttpStatusCode: %d\n", *resp.HttpStatusCode)
	}

	if resp.Success != nil && !*resp.Success {
		return fmt.Errorf("[ERROR] API error: Code=%s, Message=%s", ptrStr(resp.Code), ptrStr(resp.Message))
	}

	if resp.Data != nil {
		fmt.Printf("\n[DATA]\n")
		fmt.Printf("  ImageId: %s\n", ptrStr(resp.Data.ImageId))
	} else {
		fmt.Println("\n[DATA] (empty)")
	}

	fmt.Println("\n[SUCCESS] CreateImageFromTemplate call completed.")
	return nil
}
