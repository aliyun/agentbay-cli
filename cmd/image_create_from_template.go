// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

// image_create_from_template.go implements "agentbay image create-from-template".
// Uses raw HTTP POP RPC V1 to call CreateImageFromTemplate.

package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/agentbay/agentbay-cli/internal/agentbay"
	"github.com/agentbay/agentbay-cli/internal/client"
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

const supportedSourceImageNamespace = "customer_cli"

type sourceImageRef struct {
	Original        string
	Registry        string
	Namespace       string
	RepoAliUID      int64
	RepoAliUIDText  string
	Tag             string
	PhysicalImageID string
	IsFullPath      bool
}

type sourceImageAuthorization struct {
	DisplaySourceImage string
	SourceType         string
	OwnRepository      bool
}

type listSharedDockerReposFunc func(context.Context, *client.ListSharedDockerReposRequest) (*client.ListSharedDockerReposResponse, error)

// --- Implementation ---

func runImageCreateFromTemplate(cmd *cobra.Command, args []string) error {
	sourceImage, _ := cmd.Flags().GetString("source-image")
	imageName, _ := cmd.Flags().GetString("name")
	templateImageId, _ := cmd.Flags().GetString("imageId")

	sourceRef, err := parseSourceImageRef(sourceImage)
	if err != nil {
		return fmt.Errorf("[ERROR] invalid source-image: %w", err)
	}

	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to load configuration: %w", err)
	}
	if !cfg.IsAuthenticated() {
		return config.ErrNotAuthenticated()
	}

	cache, _ := loadACRCredential()
	apiClient := agentbay.NewClientFromConfig(cfg)
	ctx := context.Background()
	authorization, err := authorizeSourceImage(ctx, sourceRef, cache, apiClient.ListSharedDockerRepos)
	if err != nil {
		return err
	}
	physicalImageId := sourceRef.PhysicalImageID

	fmt.Println("[IMAGE] Creating custom image from template...")
	fmt.Printf("  SourceImage:      %s\n", authorization.DisplaySourceImage)
	fmt.Printf("  SourceType:       %s\n", authorization.SourceType)
	fmt.Printf("  PhysicalImageId:  %s\n", physicalImageId)
	fmt.Printf("  Name:             %s\n", imageName)
	fmt.Printf("  ImageId:          %s\n", templateImageId)

	acsClient, err := newACSClientFromConfig(cfg)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create HTTP client: %w", err)
	}

	params := map[string]string{
		"PhysicalImageId": physicalImageId,
		"ImageName":       imageName,
		"TemplateImageId": templateImageId,
	}

	fmt.Printf("Requesting CreateImageFromTemplate...")
	body, statusCode, err := acsClient.callRPC("CreateImageFromTemplate", params)
	if err != nil {
		return fmt.Errorf("[ERROR] Request failed: %w", err)
	}
	fmt.Printf(" Done. (HTTP %d)\n", statusCode)

	if statusCode < 200 || statusCode >= 300 {
		if reqID := extractCreateFromTemplateRequestID(body); reqID != "" {
			fmt.Printf("[INFO] CreateImageFromTemplate Request ID: %s\n", reqID)
		}
		return fmt.Errorf("[ERROR] API returned HTTP %d: %s", statusCode, string(body))
	}

	var resp createFromTemplateResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		if reqID := extractCreateFromTemplateRequestID(body); reqID != "" {
			fmt.Printf("[INFO] CreateImageFromTemplate Request ID: %s\n", reqID)
		}
		return fmt.Errorf("[ERROR] Failed to parse response: %w (body=%s)", err, truncateBody(body, 512))
	}
	if reqID := ptrStr(resp.RequestId); reqID != "" {
		fmt.Printf("[INFO] CreateImageFromTemplate Request ID: %s\n", reqID)
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

func parseSourceImageRef(sourceImage string) (sourceImageRef, error) {
	trimmed := strings.TrimSpace(sourceImage)
	if trimmed == "" {
		return sourceImageRef{}, fmt.Errorf("source-image is required")
	}

	ref := sourceImageRef{Original: trimmed}
	pathPart := trimmed
	if strings.HasPrefix(trimmed, "/") {
		pathPart = strings.TrimPrefix(trimmed, "/")
	} else {
		marker := "/" + supportedSourceImageNamespace + "/"
		idx := strings.Index(trimmed, marker)
		if idx < 0 {
			return sourceImageRef{}, fmt.Errorf("must use /%s/<aliuid>:<tag> or <registry>/%s/<aliuid>:<tag>", supportedSourceImageNamespace, supportedSourceImageNamespace)
		}
		ref.Registry = trimmed[:idx]
		if strings.TrimSpace(ref.Registry) == "" {
			return sourceImageRef{}, fmt.Errorf("registry path is empty")
		}
		ref.IsFullPath = true
		pathPart = trimmed[idx+1:]
	}

	parts := strings.SplitN(pathPart, "/", 2)
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return sourceImageRef{}, fmt.Errorf("must use /%s/<aliuid>:<tag> or <registry>/%s/<aliuid>:<tag>", supportedSourceImageNamespace, supportedSourceImageNamespace)
	}
	ref.Namespace = parts[0]
	if ref.Namespace != supportedSourceImageNamespace {
		return sourceImageRef{}, fmt.Errorf("unsupported namespace %q; expected %q", ref.Namespace, supportedSourceImageNamespace)
	}

	repoAndTag := parts[1]
	tagIdx := strings.LastIndex(repoAndTag, ":")
	if tagIdx <= 0 || tagIdx == len(repoAndTag)-1 {
		return sourceImageRef{}, fmt.Errorf("source-image must include a tag, e.g. /%s/<aliuid>:<tag>", supportedSourceImageNamespace)
	}
	ref.RepoAliUIDText = repoAndTag[:tagIdx]
	ref.Tag = repoAndTag[tagIdx+1:]
	if strings.Contains(ref.RepoAliUIDText, "/") {
		return sourceImageRef{}, fmt.Errorf("repo must be an AliUID under /%s", supportedSourceImageNamespace)
	}
	aliUID, err := strconv.ParseInt(ref.RepoAliUIDText, 10, 64)
	if err != nil || aliUID <= 0 {
		return sourceImageRef{}, fmt.Errorf("repo AliUID %q must be a positive integer", ref.RepoAliUIDText)
	}
	ref.RepoAliUID = aliUID
	ref.PhysicalImageID = fmt.Sprintf("/%s/%s:%s", ref.Namespace, ref.RepoAliUIDText, ref.Tag)
	return ref, nil
}

func authorizeSourceImage(ctx context.Context, ref sourceImageRef, cache *acrCredentialCache, listShares listSharedDockerReposFunc) (sourceImageAuthorization, error) {
	if cache != nil && sourceImageMatchesACRCache(ref, cache) {
		displaySource := ref.Original
		if !ref.IsFullPath {
			displaySource = cache.RegistryURL + ref.PhysicalImageID
		}
		return sourceImageAuthorization{
			DisplaySourceImage: displaySource,
			SourceType:         "Own repository",
			OwnRepository:      true,
		}, nil
	}

	if listShares == nil {
		return sourceImageAuthorization{}, fmt.Errorf("[ERROR] unable to verify source-image authorization: ListSharedDockerRepos client is not available")
	}
	if err := verifySharedDockerRepoAuthorization(ctx, listShares, ref); err != nil {
		return sourceImageAuthorization{}, err
	}
	displaySource := ref.Original
	if !ref.IsFullPath {
		displaySource = ref.PhysicalImageID
	}
	return sourceImageAuthorization{
		DisplaySourceImage: displaySource,
		SourceType:         fmt.Sprintf("Shared repository (owner AliUID: %s)", maskAliUID(ref.RepoAliUIDText)),
		OwnRepository:      false,
	}, nil
}

func sourceImageMatchesACRCache(ref sourceImageRef, cache *acrCredentialCache) bool {
	if cache == nil {
		return false
	}
	if ref.Namespace != cache.Namespace || ref.RepoAliUIDText != cache.RepoName {
		return false
	}
	return !ref.IsFullPath || ref.Registry == cache.RegistryURL
}

func verifySharedDockerRepoAuthorization(ctx context.Context, listShares listSharedDockerReposFunc, ref sourceImageRef) error {
	direction := "Incoming"
	pageStart := int32(1)
	pageSize := int32(10)
	req := &client.ListSharedDockerReposRequest{
		Direction:   &direction,
		QueryAliUid: &ref.RepoAliUID,
		PageStart:   &pageStart,
		PageSize:    &pageSize,
	}
	resp, err := listShares(ctx, req)
	if err != nil {
		if reqID := extractRequestIDFromErr(err); reqID != "" {
			fmt.Printf("[INFO] ListSharedDockerRepos Request ID: %s\n", reqID)
		}
		return fmt.Errorf("[ERROR] Failed to verify shared Docker repo authorization: %w", err)
	}
	if resp != nil && resp.Body != nil {
		if reqID := resp.Body.GetRequestId(); reqID != "" {
			fmt.Printf("[INFO] ListSharedDockerRepos Request ID: %s\n", reqID)
		}
	}
	if resp == nil || resp.Body == nil {
		return fmt.Errorf("[ERROR] Invalid ListSharedDockerRepos response: missing body")
	}
	code := resp.Body.GetCode()
	successPtr := resp.Body.Success
	if (successPtr != nil && !*successPtr) || (code != "" && !strings.EqualFold(code, "ok")) {
		return fmt.Errorf("[ERROR] Failed to verify shared Docker repo authorization: Code=%s, Message=%s", code, resp.Body.GetMessage())
	}
	if len(resp.Body.Data) == 0 {
		return fmt.Errorf("[ERROR] source-image '%s' is not owned by the current ACR cache and no incoming Docker repo sharing authorization was found for AliUID %s. If this is your own image, run 'agentbay docker login' and use the returned registry path. If this is a shared image, ask the owner to run 'agentbay docker share --target-uid <your-uid>' and verify with 'agentbay docker list-shares --direction Incoming --aliuid %s'", ref.PhysicalImageID, ref.RepoAliUIDText, ref.RepoAliUIDText)
	}
	return nil
}

func maskAliUID(uid string) string {
	if len(uid) <= 4 {
		return "****" + uid
	}
	return "****" + uid[len(uid)-4:]
}

func extractCreateFromTemplateRequestID(body []byte) string {
	var payload struct {
		RequestId *string `json:"RequestId"`
		RequestID *string `json:"RequestID"`
	}
	if err := json.Unmarshal(body, &payload); err != nil {
		return ""
	}
	if payload.RequestId != nil {
		return *payload.RequestId
	}
	if payload.RequestID != nil {
		return *payload.RequestID
	}
	return ""
}
