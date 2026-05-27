// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"archive/zip"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/term"

	"github.com/agentbay/agentbay-cli/internal/agentbay"
	"github.com/agentbay/agentbay-cli/internal/client"
	"github.com/agentbay/agentbay-cli/internal/config"
)

const skillFileName = "SKILL.md"

// Output style: label width for skill show (no emoji).
const skillDetailLabelW = 14

var SkillsCmd = &cobra.Command{
	Use:     "skills",
	Short:   "Manage AgentBay skills",
	Long:    "Push, update, and list skills; show details.",
	GroupID: "management",
}

var skillsPushCmd = &cobra.Command{
	Use:   "push <skill-dir>|<skill.zip>",
	Short: "Push a local skill directory or zip to the cloud",
	Long: `Push a local skill to the cloud (upload zip, create/update skill).
Accepts either a directory (with SKILL.md and frontmatter name/description) or a .zip file.
Directory: packed to zip then uploaded. Zip: uploaded as-is.`,
	Args: cobra.ExactArgs(1),
	RunE: runSkillsPush,
}

var skillsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List cloud skills",
	Long:  `List skills visible to you (yours and public), with optional filters for name and tags.`,
	Args:  cobra.NoArgs,
	RunE:  runSkillsList,
}

var skillsShowCmd = &cobra.Command{
	Use:   "show <skill-id>",
	Short: "Show skill details",
	Long:  `Show details for a skill by ID.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runSkillsShow,
}

var skillsUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing skill in the cloud",
	Long:  `Update an existing skill by ID. Upload a new zip file and optionally update tags or icon.`,
	Args:  cobra.NoArgs,
	RunE:  runSkillsUpdate,
}

func init() {
	SkillsCmd.AddCommand(skillsPushCmd)
	SkillsCmd.AddCommand(skillsListCmd)
	SkillsCmd.AddCommand(skillsShowCmd)
	SkillsCmd.AddCommand(skillsUpdateCmd)
	SkillsCmd.AddCommand(skillsDeleteCmd)

	skillsPushCmd.Flags().StringArray("tag", nil, "Tag name for the skill (can be specified multiple times, e.g. --tag \"tag1\" --tag \"tag2\")")
	skillsPushCmd.Flags().String("icon", "https://img.alicdn.com/imgextra/i4/O1CN01syuoCy1qhsZxbwuBz_!!6000000005528-2-tps-100-100.png", "Icon for the skill (URL or identifier); uses the default AgentBay icon if not specified")

	skillsListCmd.Flags().Int("page", 1, "Page number (default: 1)")
	skillsListCmd.Flags().Int("size", 10, "Page size (default: 10)")
	skillsListCmd.Flags().String("name", "", "Filter by skill name (optional)")
	skillsListCmd.Flags().StringArray("tag", nil, "Filter by tag name (can be specified multiple times, e.g. --tag test --tag aliyun)")
	skillsListCmd.Flags().StringP("output", "o", "", `Output format. Use "json" for machine-readable output (e.g. for AI/scripts)`)

	skillsUpdateCmd.Flags().String("skill-id", "", "Skill ID to update (required)")
	_ = skillsUpdateCmd.MarkFlagRequired("skill-id")
	skillsUpdateCmd.Flags().String("file", "", "Path to skill directory or .zip file (required)")
	_ = skillsUpdateCmd.MarkFlagRequired("file")
	skillsUpdateCmd.Flags().StringArray("tag", nil, `Tag name for the skill (can be specified multiple times, e.g. --tag "tag1" --tag "tag2")`)
	skillsUpdateCmd.Flags().String("icon", "", "Icon for the skill (e.g. URL or identifier)")
	skillsUpdateCmd.Flags().Bool("clear-tags", false, "Remove all tags from the skill")

	skillsDeleteCmd.Flags().String("skill-id", "", "Skill ID to delete")
	skillsDeleteCmd.Flags().BoolP("yes", "y", false, "Skip confirmation prompt and skill detail lookup (for non-interactive use)")
}

// parseSkillFrontmatter parses --- name: x description: y --- from SKILL.md content.
// name is required; description is optional.
func parseSkillFrontmatter(content []byte) (name, description string, err error) {
	re := regexp.MustCompile(`(?m)^name:\s*(.+)$`)
	if m := re.FindSubmatch(content); len(m) >= 2 {
		name = strings.TrimSpace(string(m[1]))
	}
	if name == "" {
		return "", "", fmt.Errorf("SKILL.md must contain frontmatter with 'name:'")
	}
	reDesc := regexp.MustCompile(`(?m)^description:\s*(.+)$`)
	if m := reDesc.FindSubmatch(content); len(m) >= 2 {
		description = strings.TrimSpace(string(m[1]))
	}
	return name, description, nil
}

func runSkillsPush(cmd *cobra.Command, args []string) error {
	pathInput := filepath.Clean(args[0])
	info, err := os.Stat(pathInput)
	if err != nil {
		if os.IsNotExist(err) {
			return printErrorMessage(
				fmt.Sprintf("[ERROR] Path does not exist: %s", pathInput),
				"",
				fmt.Sprintf("[TIP] Usage: agentbay skills push <skill-dir> or agentbay skills push <skill.zip>"),
			)
		}
		return fmt.Errorf("path: %w", err)
	}

	var skillZipName string
	var zipPath string // path to zip file to upload (temp file for dir, or pathInput for .zip)

	if info.IsDir() {
		skillDir := pathInput
		skillMdPath := filepath.Join(skillDir, skillFileName)
		skillMd, err := os.ReadFile(skillMdPath)
		if err != nil {
			if os.IsNotExist(err) {
				return printErrorMessage(
					fmt.Sprintf("[ERROR] %s not found in %s", skillFileName, skillDir),
					"",
					fmt.Sprintf("[TIP] Create %s with frontmatter: name: <skill-name>", skillFileName),
				)
			}
			return fmt.Errorf("reading %s: %w", skillFileName, err)
		}
		_, _, err = parseSkillFrontmatter(skillMd)
		if err != nil {
			return printErrorMessage(
				fmt.Sprintf("[ERROR] %s", err.Error()),
				"",
				fmt.Sprintf("[TIP] Add frontmatter to %s: ---", skillFileName),
				"      name: my-skill",
				"      description: Optional description",
				"      ---",
			)
		}
		skillZipName = skillDirToZipFileName(skillDir)
		// zipPath will be set after we pack to temp file below
	} else {
		// Regular file: must be .zip
		if !strings.HasSuffix(strings.ToLower(pathInput), ".zip") {
			return printErrorMessage(
				fmt.Sprintf("[ERROR] Not a directory or .zip file: %s", pathInput),
				"",
				fmt.Sprintf("[TIP] Usage: agentbay skills push <skill-dir> or agentbay skills push <skill.zip>"),
			)
		}
		skillZipName = filepath.Base(pathInput)
		zipPath = pathInput
	}

	// Parse tags flag
	tagsFlag, _ := cmd.Flags().GetStringArray("tag")
	var tags []string
	for _, t := range tagsFlag {
		trimmed := strings.TrimSpace(t)
		if trimmed != "" {
			tags = append(tags, trimmed)
		}
	}

	// Parse icon flag
	iconInput, _ := cmd.Flags().GetString("icon")

	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}
	apiClient := agentbay.NewClientFromConfig(cfg)
	ctx := context.Background()

	// Dynamic step numbering
	stepIdx := 1
	totalSteps := 3
	if len(tags) > 0 {
		totalSteps = 4
	}

	// Step 1 (optional): Process tags before getting credential to avoid expiry
	if len(tags) > 0 {
		fmt.Printf("[STEP %d/%d] Processing tags...\n", stepIdx, totalSteps)
		stepIdx++

		listResp, err := apiClient.ListTag(ctx)
		if err != nil {
			printRequestIDFromErrIfVerbose(cmd, err)
			return fmt.Errorf("[ERROR] Failed to list tags: %w", err)
		}
		if listResp.Body != nil && listResp.Body.GetRequestId() != "" {
			fmt.Printf("[INFO] ListTag RequestId: %s\n", listResp.Body.GetRequestId())
		}

		existingTagNames := map[string]bool{}
		if listResp.Body != nil && listResp.Body.Data != nil {
			for _, item := range listResp.Body.Data {
				if item.TagName != nil {
					existingTagNames[*item.TagName] = true
				}
			}
		}

		var missingTags []string
		for _, tagName := range tags {
			if !existingTagNames[tagName] {
				missingTags = append(missingTags, tagName)
			}
		}

		if len(missingTags) > 0 {
			fmt.Printf("[INFO] Tags not found: %s, creating...\n", strings.Join(missingTags, ", "))
			createReq := &client.CreateTagRequest{TagList: missingTags}
			createTagResp, err := apiClient.CreateTag(ctx, createReq)
			if err != nil {
				printRequestIDFromErrIfVerbose(cmd, err)
				return fmt.Errorf("[ERROR] Failed to create tags: %w", err)
			}
			if createTagResp.Body != nil && createTagResp.Body.GetRequestId() != "" {
				fmt.Printf("[INFO] CreateTag RequestId: %s\n", createTagResp.Body.GetRequestId())
			}
			fmt.Printf("[INFO] Tags created successfully.\n")
		} else {
			fmt.Printf("[INFO] All tags already exist.\n")
		}
	}

	// Next step: Get upload credential
	fmt.Printf("[STEP %d/%d] Getting upload credential...\n", stepIdx, totalSteps)
	stepIdx++
	credReq := &client.GetMarketSkillCredentialRequest{FileName: &skillZipName}
	var credResp *client.GetMarketSkillCredentialResponse
	err = withTransientRetry(client.DefaultRetryConfig(), "GetMarketSkillCredential", func() error {
		var e error
		credResp, e = apiClient.GetMarketSkillCredential(ctx, credReq)
		return e
	})
	if err != nil {
		printRequestIDFromErrIfVerbose(cmd, err)
		return fmt.Errorf("[ERROR] Failed to get upload credential: %w", err)
	}
	if credResp.Body == nil || credResp.Body.Data == nil {
		return fmt.Errorf("invalid response: missing credential data")
	}
	if credResp.Body != nil && credResp.Body.RequestId != nil && *credResp.Body.RequestId != "" {
		fmt.Printf("[INFO] GetMarketSkillCredential RequestId: %s\n", *credResp.Body.RequestId)
	}
	uploadURLStr := ""
	if u := credResp.Body.Data.GetOssUrl(); u != nil && *u != "" {
		uploadURLStr = *u
	}
	if uploadURLStr == "" {
		if u := credResp.Body.Data.GetUrl(); u != nil && *u != "" {
			uploadURLStr = *u
		}
	}
	if uploadURLStr == "" {
		return fmt.Errorf("invalid response: missing OSS upload URL")
	}
	if verbose, _ := cmd.Flags().GetBool("verbose"); verbose {
		fmt.Fprintf(os.Stderr, "[DEBUG] OSS upload URL: %s\n", uploadURLStr)
	}

	// For step 3: prefer credential response OssBucket/OssFilePath (backend format); fallback to parsing upload URL.
	createBucket, createOssFilePath, err := parseBucketAndPathForCreate(credResp.Body.Data, uploadURLStr)
	if err != nil {
		return err
	}

	if zipPath != "" {
		// Direct zip file: upload as-is
		fmt.Printf("[STEP %d/%d] Uploading skill zip...\n", stepIdx, totalSteps)
		if verbose, _ := cmd.Flags().GetBool("verbose"); verbose {
			if fi, err := os.Stat(zipPath); err == nil {
				fmt.Fprintf(os.Stderr, "[DEBUG] Upload size: %d bytes, file: %s\n", fi.Size(), zipPath)
			}
		}
		if err := uploadFileToOSS(zipPath, uploadURLStr); err != nil {
			return fmt.Errorf("[ERROR] Failed to upload: %w", err)
		}
	} else {
		// Directory: pack then upload
		fmt.Printf("[STEP %d/%d] Packing and uploading skill...\n", stepIdx, totalSteps)
		zipBuf, err := zipSkillDir(pathInput)
		if err != nil {
			return fmt.Errorf("pack skill: %w", err)
		}
		tmpFile, err := os.CreateTemp("", "agentbay-skill-*.zip")
		if err != nil {
			return fmt.Errorf("create temp zip: %w", err)
		}
		tmpPath := tmpFile.Name()
		defer os.Remove(tmpPath)
		if _, err := tmpFile.Write(zipBuf.Bytes()); err != nil {
			_ = tmpFile.Close()
			return fmt.Errorf("write temp zip: %w", err)
		}
		if err := tmpFile.Close(); err != nil {
			return fmt.Errorf("close temp zip: %w", err)
		}
		if verbose, _ := cmd.Flags().GetBool("verbose"); verbose {
			fmt.Fprintf(os.Stderr, "[DEBUG] Upload size: %d bytes, temp file: %s\n", zipBuf.Len(), tmpPath)
		}
		if err := uploadFileToOSS(tmpPath, uploadURLStr); err != nil {
			return fmt.Errorf("[ERROR] Failed to upload: %w", err)
		}
	}
	stepIdx++

	fmt.Printf("[STEP %d/%d] Creating skill...\n", stepIdx, totalSteps)
	// Pre-release credential URL path is often "null/<id>"; some backends reject "null" in OssFilePath.
	// Pass only the suffix for CreateMarketSkill when path is exactly "null/<suffix>" so backend receives a valid path.
	createOssPath := createOssFilePath
	if strings.HasPrefix(createOssPath, "null/") && len(createOssPath) > 5 {
		createOssPath = createOssPath[5:] // len("null/")
	}
	createReq := &client.CreateMarketSkillRequest{
		OssBucket:   &createBucket,
		OssFilePath: &createOssPath,
	}
	if len(tags) > 0 {
		createReq.TagList = tags
	}
	if iconInput != "" {
		createReq.Icon = &iconInput
	}
	var createResp *client.CreateMarketSkillResponse
	err = withTransientRetry(client.DefaultRetryConfig(), "CreateMarketSkill", func() error {
		var e error
		createResp, e = apiClient.CreateMarketSkill(ctx, createReq)
		return e
	})
	if err != nil {
		if createResp != nil && createResp.RawBody != "" {
			fmt.Fprintf(os.Stderr, "[DEBUG] Raw response: %s\n", createResp.RawBody)
		}
		printRequestIDFromErrIfVerbose(cmd, err)
		return fmt.Errorf("[ERROR] Failed to create skill: %w", err)
	}
	if createResp.Body != nil && createResp.Body.RequestId != nil && *createResp.Body.RequestId != "" {
		fmt.Printf("[INFO] CreateMarketSkill RequestId: %s\n", *createResp.Body.RequestId)
	}
	var skillId string
	if createResp.Body != nil && createResp.Body.Data != nil && createResp.Body.Data.SkillId != nil {
		skillId = *createResp.Body.Data.SkillId
	}
	if skillId == "" {
		skillId = "<unknown>"
	}
	fmt.Println()
	fmt.Printf("[SUCCESS] ✅ Skill created successfully!\n")
	fmt.Printf("[RESULT] Skill ID: %s\n", skillId)
	return nil
}

func runSkillsUpdate(cmd *cobra.Command, args []string) error {
	skillId, _ := cmd.Flags().GetString("skill-id")
	fileInput, _ := cmd.Flags().GetString("file")
	iconInput, _ := cmd.Flags().GetString("icon")
	clearTags, _ := cmd.Flags().GetBool("clear-tags")

	// Parse tags flag
	tagsFlag, _ := cmd.Flags().GetStringArray("tag")
	var tags []string
	for _, t := range tagsFlag {
		trimmed := strings.TrimSpace(t)
		if trimmed != "" {
			tags = append(tags, trimmed)
		}
	}

	// Validate: --tag and --clear-tags are mutually exclusive
	if clearTags && len(tags) > 0 {
		return fmt.Errorf("[ERROR] --clear-tags and --tag cannot be used together")
	}

	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}
	apiClient := agentbay.NewClientFromConfig(cfg)
	ctx := context.Background()

	// Calculate steps
	hasFile := fileInput != ""
	stepIdx := 1
	totalSteps := 0
	if len(tags) > 0 {
		totalSteps++ // ListTag/CreateTag step
	}
	if hasFile {
		totalSteps += 2 // credential + upload
	}
	totalSteps++ // UpdateMarketSkill call

	// Step: Process tags (if any)
	if len(tags) > 0 {
		fmt.Printf("[STEP %d/%d] Processing tags...\n", stepIdx, totalSteps)
		stepIdx++

		listResp, err := apiClient.ListTag(ctx)
		if err != nil {
			printRequestIDFromErrIfVerbose(cmd, err)
			return fmt.Errorf("[ERROR] Failed to list tags: %w", err)
		}
		if listResp.Body != nil && listResp.Body.GetRequestId() != "" {
			fmt.Printf("[INFO] ListTag RequestId: %s\n", listResp.Body.GetRequestId())
		}

		existingTagNames := map[string]bool{}
		if listResp.Body != nil && listResp.Body.Data != nil {
			for _, item := range listResp.Body.Data {
				if item.TagName != nil {
					existingTagNames[*item.TagName] = true
				}
			}
		}

		var missingTags []string
		for _, tagName := range tags {
			if !existingTagNames[tagName] {
				missingTags = append(missingTags, tagName)
			}
		}

		if len(missingTags) > 0 {
			fmt.Printf("[INFO] Tags not found: %s, creating...\n", strings.Join(missingTags, ", "))
			createReq := &client.CreateTagRequest{TagList: missingTags}
			createTagResp, err := apiClient.CreateTag(ctx, createReq)
			if err != nil {
				printRequestIDFromErrIfVerbose(cmd, err)
				return fmt.Errorf("[ERROR] Failed to create tags: %w", err)
			}
			if createTagResp.Body != nil && createTagResp.Body.GetRequestId() != "" {
				fmt.Printf("[INFO] CreateTag RequestId: %s\n", createTagResp.Body.GetRequestId())
			}
			fmt.Printf("[INFO] Tags created successfully.\n")
		} else {
			fmt.Printf("[INFO] All tags already exist.\n")
		}
	}

	var ossBucket, ossFilePath string

	// Steps: Get credential + upload (if --file provided)
	if hasFile {
		pathInput := filepath.Clean(fileInput)
		info, err := os.Stat(pathInput)
		if err != nil {
			if os.IsNotExist(err) {
				return printErrorMessage(
					fmt.Sprintf("[ERROR] Path does not exist: %s", pathInput),
					"",
					"[TIP] Provide a valid skill directory or .zip file via --file",
				)
			}
			return fmt.Errorf("path: %w", err)
		}

		var skillZipName string
		var zipPath string

		if info.IsDir() {
			skillDir := pathInput
			skillMdPath := filepath.Join(skillDir, skillFileName)
			skillMd, err := os.ReadFile(skillMdPath)
			if err != nil {
				if os.IsNotExist(err) {
					return printErrorMessage(
						fmt.Sprintf("[ERROR] %s not found in %s", skillFileName, skillDir),
						"",
						fmt.Sprintf("[TIP] Create %s with frontmatter: name: <skill-name>", skillFileName),
					)
				}
				return fmt.Errorf("reading %s: %w", skillFileName, err)
			}
			_, _, err = parseSkillFrontmatter(skillMd)
			if err != nil {
				return printErrorMessage(
					fmt.Sprintf("[ERROR] %s", err.Error()),
					"",
					fmt.Sprintf("[TIP] Add frontmatter to %s: ---", skillFileName),
					"      name: my-skill",
					"      description: Optional description",
					"      ---",
				)
			}
			skillZipName = skillDirToZipFileName(skillDir)
		} else {
			if !strings.HasSuffix(strings.ToLower(pathInput), ".zip") {
				return printErrorMessage(
					fmt.Sprintf("[ERROR] Not a directory or .zip file: %s", pathInput),
					"",
					"[TIP] Provide a skill directory or .zip file via --file",
				)
			}
			skillZipName = filepath.Base(pathInput)
			zipPath = pathInput
		}

		// Get upload credential
		fmt.Printf("[STEP %d/%d] Getting upload credential...\n", stepIdx, totalSteps)
		stepIdx++
		credReq := &client.GetMarketSkillCredentialRequest{FileName: &skillZipName}
		var credResp *client.GetMarketSkillCredentialResponse
		err = withTransientRetry(client.DefaultRetryConfig(), "GetMarketSkillCredential", func() error {
			var e error
			credResp, e = apiClient.GetMarketSkillCredential(ctx, credReq)
			return e
		})
		if err != nil {
			printRequestIDFromErrIfVerbose(cmd, err)
			return fmt.Errorf("[ERROR] Failed to get upload credential: %w", err)
		}
		if credResp.Body == nil || credResp.Body.Data == nil {
			return fmt.Errorf("invalid response: missing credential data")
		}
		if credResp.Body != nil && credResp.Body.RequestId != nil && *credResp.Body.RequestId != "" {
			fmt.Printf("[INFO] GetMarketSkillCredential RequestId: %s\n", *credResp.Body.RequestId)
		}
		uploadURLStr := ""
		if u := credResp.Body.Data.GetOssUrl(); u != nil && *u != "" {
			uploadURLStr = *u
		}
		if uploadURLStr == "" {
			if u := credResp.Body.Data.GetUrl(); u != nil && *u != "" {
				uploadURLStr = *u
			}
		}
		if uploadURLStr == "" {
			return fmt.Errorf("invalid response: missing OSS upload URL")
		}
		if verbose, _ := cmd.Flags().GetBool("verbose"); verbose {
			fmt.Fprintf(os.Stderr, "[DEBUG] OSS upload URL: %s\n", uploadURLStr)
		}

		createBucket, createOssFilePath, err := parseBucketAndPathForCreate(credResp.Body.Data, uploadURLStr)
		if err != nil {
			return err
		}

		// Upload
		if zipPath != "" {
			fmt.Printf("[STEP %d/%d] Uploading skill zip...\n", stepIdx, totalSteps)
			if verbose, _ := cmd.Flags().GetBool("verbose"); verbose {
				if fi, err := os.Stat(zipPath); err == nil {
					fmt.Fprintf(os.Stderr, "[DEBUG] Upload size: %d bytes, file: %s\n", fi.Size(), zipPath)
				}
			}
			if err := uploadFileToOSS(zipPath, uploadURLStr); err != nil {
				return fmt.Errorf("[ERROR] Failed to upload: %w", err)
			}
		} else {
			fmt.Printf("[STEP %d/%d] Packing and uploading skill...\n", stepIdx, totalSteps)
			zipBuf, err := zipSkillDir(pathInput)
			if err != nil {
				return fmt.Errorf("pack skill: %w", err)
			}
			tmpFile, err := os.CreateTemp("", "agentbay-skill-*.zip")
			if err != nil {
				return fmt.Errorf("create temp zip: %w", err)
			}
			tmpPath := tmpFile.Name()
			defer os.Remove(tmpPath)
			if _, err := tmpFile.Write(zipBuf.Bytes()); err != nil {
				_ = tmpFile.Close()
				return fmt.Errorf("write temp zip: %w", err)
			}
			if err := tmpFile.Close(); err != nil {
				return fmt.Errorf("close temp zip: %w", err)
			}
			if verbose, _ := cmd.Flags().GetBool("verbose"); verbose {
				fmt.Fprintf(os.Stderr, "[DEBUG] Upload size: %d bytes, temp file: %s\n", zipBuf.Len(), tmpPath)
			}
			if err := uploadFileToOSS(tmpPath, uploadURLStr); err != nil {
				return fmt.Errorf("[ERROR] Failed to upload: %w", err)
			}
		}
		stepIdx++

		// Prepare OSS fields for UpdateMarketSkill
		ossBucket = createBucket
		ossFilePath = createOssFilePath
		// Pre-release credential URL path is often "null/<id>"; some backends reject "null" in OssFilePath.
		if strings.HasPrefix(ossFilePath, "null/") && len(ossFilePath) > 5 {
			ossFilePath = ossFilePath[5:]
		}
	}

	// Final step: Call UpdateMarketSkill
	fmt.Printf("[STEP %d/%d] Updating skill...\n", stepIdx, totalSteps)
	updateReq := &client.UpdateMarketSkillRequest{
		SkillId: &skillId,
	}
	if hasFile {
		updateReq.OssBucket = &ossBucket
		updateReq.OssFilePath = &ossFilePath
	}
	if len(tags) > 0 {
		updateReq.TagList = tags
	} else if clearTags {
		updateReq.TagList = []string{} // explicit empty slice → API clears all tags
	}
	// neither → TagList remains nil → API preserves existing tags
	if iconInput != "" {
		updateReq.Icon = &iconInput
	}

	var updateResp *client.CreateMarketSkillResponse
	err = withTransientRetry(client.DefaultRetryConfig(), "UpdateMarketSkill", func() error {
		var e error
		updateResp, e = apiClient.UpdateMarketSkill(ctx, updateReq)
		return e
	})
	if err != nil {
		if updateResp != nil && updateResp.RawBody != "" {
			fmt.Fprintf(os.Stderr, "[DEBUG] Raw response: %s\n", updateResp.RawBody)
		}
		printRequestIDFromErrIfVerbose(cmd, err)
		return fmt.Errorf("[ERROR] Failed to update skill: %w", err)
	}
	if updateResp.Body != nil && updateResp.Body.RequestId != nil && *updateResp.Body.RequestId != "" {
		fmt.Printf("[INFO] UpdateMarketSkill RequestId: %s\n", *updateResp.Body.RequestId)
	}
	fmt.Println()
	fmt.Printf("[SUCCESS] ✅ Skill updated successfully!\n")
	fmt.Printf("[RESULT] Skill ID: %s\n", skillId)
	return nil
}

// skillDirToZipFileName returns the zip filename to use for upload, derived from the skill directory path.
// Example: /path/to/xlsx -> "xlsx.zip", ./pdf -> "pdf.zip". Falls back to "skill.zip" if base is empty or ".".
func skillDirToZipFileName(skillDir string) string {
	base := filepath.Base(skillDir)
	if base == "" || base == "." {
		return "skill.zip"
	}
	return base + ".zip"
}

// parseBucketAndPathForCreate returns bucket and path for CreateMarketSkill (step 3).
// Prefers credential response OssBucket/OssFilePath when both are set; otherwise parses from upload URL.
// Path is always trimmed of trailing slash so backend receives e.g. "1762926266827681/5gwKWy4Y" not "1762926266827681/5gwKWy4Y/".
func parseBucketAndPathForCreate(data *client.GetMarketSkillCredentialResponseBodyData, uploadURLStr string) (createBucket, createOssFilePath string, err error) {
	if data != nil {
		if b := data.GetOssBucket(); b != nil && *b != "" {
			if p := data.GetOssFilePath(); p != nil && *p != "" {
				createBucket = *b
				createOssFilePath = strings.TrimSuffix(*p, "/")
				return createBucket, createOssFilePath, nil
			}
		}
	}
	bucket, path, err := parseOSSBucketAndPath(uploadURLStr)
	if err != nil {
		return "", "", fmt.Errorf("parse OSS URL for create: %w", err)
	}
	createBucket = bucket
	createOssFilePath = strings.TrimSuffix(path, "/")
	return createBucket, createOssFilePath, nil
}

// parseOSSBucketAndPath extracts bucket and object path from an OSS pre-signed URL.
// Example: https://bucket.oss-cn-hangzhou.aliyuncs.com/prefix/key?Expires=... -> bucket, "prefix/key"
func parseOSSBucketAndPath(ossURL string) (bucket, objectPath string, err error) {
	u, err := url.Parse(ossURL)
	if err != nil {
		return "", "", err
	}
	host := u.Host
	if host == "" {
		return "", "", fmt.Errorf("missing host in URL")
	}
	// Bucket: first label of host, e.g. "agentbay-market-skill" from "agentbay-market-skill.oss-cn-hangzhou.aliyuncs.com"
	parts := strings.SplitN(host, ".", 2)
	if len(parts) < 2 {
		return "", "", fmt.Errorf("host not in bucket.oss-... form")
	}
	bucket = parts[0]
	objectPath = strings.TrimPrefix(u.Path, "/")
	return bucket, objectPath, nil
}

func zipSkillDir(dir string) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		rel, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)
		if strings.HasPrefix(rel, "..") {
			return nil
		}
		fh, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		fh.Name = rel
		fh.Method = zip.Deflate // Backend Java ZipInputStream requires DEFLATED entries when EXT descriptor is present.
		fw, err := w.CreateHeader(fh)
		if err != nil {
			return err
		}
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = io.Copy(fw, f)
		return err
	})
	if err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return buf, nil
}

func runSkillsList(cmd *cobra.Command, args []string) error {
	page, _ := cmd.Flags().GetInt("page")
	size, _ := cmd.Flags().GetInt("size")
	name, _ := cmd.Flags().GetString("name")
	tags, _ := cmd.Flags().GetStringArray("tag")
	outputFmt, _ := cmd.Flags().GetString("output")

	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}
	if !cfg.IsAuthenticated() {
		return fmt.Errorf("[ERROR] Not authenticated. Run 'agentbay login' or set AGENTBAY_ACCESS_KEY_ID/AGENTBAY_ACCESS_KEY_SECRET")
	}
	apiClient := agentbay.NewClientFromConfig(cfg)
	ctx := context.Background()

	req := &client.ListMarketSkillByPageRequest{}
	if page > 0 {
		pageNo := int32(page)
		req.PageNo = &pageNo
	}
	if size > 0 {
		pageSize := int32(size)
		req.PageSize = &pageSize
	}
	if name != "" {
		req.SkillName = &name
	}
	if len(tags) > 0 {
		req.TagList = tags
	}

	resp, err := apiClient.ListMarketSkillByPage(ctx, req)
	if err != nil {
		if reqID := extractRequestIDFromErr(err); reqID != "" {
			fmt.Printf("[INFO] ListMarketSkillByPage Request ID: %s\n", reqID)
		}
		return fmt.Errorf("[ERROR] Failed to list skills: %w", err)
	}

	if resp != nil && resp.Body != nil {
		if reqId := resp.Body.GetRequestId(); reqId != nil && *reqId != "" {
			fmt.Printf("[INFO] ListMarketSkillByPage Request ID: %s\n", *reqId)
		}
	}

	// Success check: Code != "ok" = failure
	if resp.Body != nil {
		code := ""
		if resp.Body.Code != nil {
			code = *resp.Body.Code
		}
		successPtr := resp.Body.Success
		if (successPtr != nil && !*successPtr) || (code != "" && !strings.EqualFold(code, "ok")) {
			msg := ""
			if resp.Body.Message != nil {
				msg = *resp.Body.Message
			}
			return fmt.Errorf("[ERROR] Failed to list skills: Code=%s, Message=%s", code, msg)
		}
	}

	data := resp.Body.GetData()
	if data == nil {
		fmt.Println("[INFO] No skills found.")
		return nil
	}

	// JSON output mode
	if strings.EqualFold(outputFmt, "json") {
		type skillJSON struct {
			SkillId     string   `json:"skillId"`
			SkillName   string   `json:"skillName"`
			Description string   `json:"description"`
			Status      string   `json:"status"`
			Tags        []string `json:"tags"`
			Icon        string   `json:"icon"`
			GmtModified string   `json:"gmtModified"`
			GmtCreate   string   `json:"gmtCreate"`
		}
		type pageJSON struct {
			TotalCount int32       `json:"totalCount"`
			TotalPage  int32       `json:"totalPage"`
			PageSize   int32       `json:"pageSize"`
			PageNumber int32       `json:"pageNumber"`
			Result     []skillJSON `json:"result"`
		}
		var pg pageJSON
		if data.TotalCount != nil {
			pg.TotalCount = *data.TotalCount
		}
		if data.TotalPage != nil {
			pg.TotalPage = *data.TotalPage
		}
		if data.PageSize != nil {
			pg.PageSize = *data.PageSize
		}
		if data.PageNumber != nil {
			pg.PageNumber = *data.PageNumber
		}
		for _, item := range data.GetResult() {
			s := skillJSON{
				Tags: item.TenantTags,
			}
			if item.SkillId != nil {
				s.SkillId = *item.SkillId
			}
			if item.SkillName != nil {
				s.SkillName = *item.SkillName
			}
			if item.Description != nil {
				s.Description = *item.Description
			}
			if item.SkillStatus != nil {
				s.Status = *item.SkillStatus
			}
			if item.Icon != nil {
				s.Icon = *item.Icon
			}
			if item.GmtModified != nil {
				s.GmtModified = *item.GmtModified
			}
			if item.GmtCreate != nil {
				s.GmtCreate = *item.GmtCreate
			}
			if s.Tags == nil {
				s.Tags = []string{}
			}
			pg.Result = append(pg.Result, s)
		}
		if pg.Result == nil {
			pg.Result = []skillJSON{}
		}
		out, jerr := json.MarshalIndent(pg, "", "  ")
		if jerr != nil {
			return fmt.Errorf("json marshal: %w", jerr)
		}
		fmt.Println(string(out))
		return nil
	}

	// Print pagination info
	totalCount := int32(0)
	totalPage := int32(0)
	pageNum := int32(1)
	pageDisplaySize := int32(size)
	if data.TotalCount != nil {
		totalCount = *data.TotalCount
	}
	if data.TotalPage != nil {
		totalPage = *data.TotalPage
	}
	if data.PageNumber != nil {
		pageNum = *data.PageNumber
	}
	if data.PageSize != nil {
		pageDisplaySize = *data.PageSize
	}
	fmt.Printf("[PAGE] Page %d of %d (Page Size: %d, Total: %d)\n\n", pageNum, totalPage, pageDisplaySize, totalCount)

	results := data.GetResult()
	if len(results) == 0 {
		fmt.Println("[INFO] No skills found.")
		return nil
	}

	// Compute dynamic column widths based on terminal width.
	// Priority: SKILL NAME > SKILL ID > STATUS > TAGS > MODIFIED
	termWidth := 120
	if w, _, err := term.GetSize(int(os.Stdout.Fd())); err == nil && w > 0 {
		termWidth = w
	}

	const (
		colSkillID      = 32 // skill-xxx… always fixed
		colStatus       = 13 // VERIFY_PASSED length
		colSepPerCol    = 2  // two-space separator between columns
		colSkillNameMin = 15
		colSkillNameMax = 30
		colTagsMin      = 20
		colModifiedFull = 30
	)

	// Fixed budget: ID + STATUS + separators for 4-column base (name|id|status|tags)
	// separator count = (numCols - 1) * 2; we try 5 cols first, then 4, then 3.
	fixedBase := colSkillID + colStatus // 45

	// Determine SKILL NAME width
	colSkillName := colSkillNameMax
	if termWidth < colSkillNameMax+fixedBase+colSepPerCol*2+colTagsMin+colSepPerCol {
		colSkillName = colSkillNameMin
	}

	// Remaining space after name + id + status + separators
	// We always have at least: name | id | status (3 cols, 2 separators)
	baseUsed := colSkillName + colSkillID + colStatus + colSepPerCol*2
	remaining := termWidth - baseUsed

	// Decide which optional columns to show and their widths
	showModified := false
	showTags := false
	colTags := 0
	colModified := 0

	if remaining >= colSepPerCol+colTagsMin+colSepPerCol+colModifiedFull {
		// Enough for both TAGS and MODIFIED
		showTags = true
		showModified = true
		// Allocate MODIFIED its full width, give rest to TAGS
		colModified = colModifiedFull
		colTags = remaining - colSepPerCol*2 - colModified
		if colTags < colTagsMin {
			colTags = colTagsMin
		}
	} else if remaining >= colSepPerCol+colTagsMin {
		// Only enough for TAGS
		showTags = true
		colTags = remaining - colSepPerCol
		if colTags < colTagsMin {
			colTags = colTagsMin
		}
	}
	// else: only NAME | ID | STATUS

	// Build header
	header := fmt.Sprintf("%-*s  %-*s  %-*s",
		colSkillName, "SKILL NAME",
		colSkillID, "SKILL ID",
		colStatus, "STATUS",
	)
	separator := strings.Repeat("-", colSkillName) + "  " +
		strings.Repeat("-", colSkillID) + "  " +
		strings.Repeat("-", colStatus)
	if showTags {
		header += fmt.Sprintf("  %-*s", colTags, "TAGS")
		separator += "  " + strings.Repeat("-", colTags)
	}
	if showModified {
		header += fmt.Sprintf("  %-*s", colModified, "MODIFIED")
		separator += "  " + strings.Repeat("-", colModified)
	}
	fmt.Println(header)
	fmt.Println(separator)

	for _, item := range results {
		skillId := ""
		if item.SkillId != nil {
			skillId = *item.SkillId
		}
		skillName := ""
		if item.SkillName != nil {
			skillName = *item.SkillName
		}
		status := ""
		if item.SkillStatus != nil {
			status = *item.SkillStatus
		}
		tagsStr := strings.Join(item.TenantTags, ", ")
		modified := ""
		if item.GmtModified != nil {
			modified = *item.GmtModified
		}

		row := fmt.Sprintf("%-*s  %-*s  %-*s",
			colSkillName, truncateStr(skillName, colSkillName),
			colSkillID, truncateStr(skillId, colSkillID),
			colStatus, truncateStr(status, colStatus),
		)
		if showTags {
			row += fmt.Sprintf("  %-*s", colTags, truncateStr(tagsStr, colTags))
		}
		if showModified {
			row += fmt.Sprintf("  %-*s", colModified, truncateStr(modified, colModified))
		}
		fmt.Println(row)
	}

	// Show next page tip if there are more pages
	if pageNum < totalPage {
		fmt.Printf("\n[TIP] Use --page %d to view the next page.\n", pageNum+1)
	}

	return nil
}

// truncateStr truncates s to maxLen runes, appending "..." if truncated.
func truncateStr(s string, maxLen int) string {
	runes := []rune(s)
	if len(runes) <= maxLen {
		return s
	}
	if maxLen <= 3 {
		return string(runes[:maxLen])
	}
	return string(runes[:maxLen-3]) + "..."
}

func runSkillsShow(cmd *cobra.Command, args []string) error {
	skillId := args[0]
	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}
	apiClient := agentbay.NewClientFromConfig(cfg)
	ctx := context.Background()

	req := &client.DescribeMarketSkillDetailRequest{SkillId: &skillId}
	resp, err := apiClient.DescribeMarketSkillDetail(ctx, req)
	if err != nil {
		printRequestIDFromErrIfVerbose(cmd, err)
		return fmt.Errorf("[ERROR] Failed to get skill details: %w", err)
	}
	if resp.Body == nil || resp.Body.Data == nil {
		fmt.Fprintf(os.Stderr, "[INFO] No details for skill %s\n", skillId)
		return nil
	}
	if resp.Body.RequestId != nil && *resp.Body.RequestId != "" {
		fmt.Printf("[INFO] DescribeMarketSkillDetail RequestId: %s\n", *resp.Body.RequestId)
	}
	d := resp.Body.Data
	displaySkillId := strPtr(d.GetSkillId())
	if displaySkillId == "" {
		displaySkillId = skillId
	}
	fmt.Printf("%-*s %s\n", skillDetailLabelW, "SkillId:", displaySkillId)
	fmt.Printf("%-*s %s\n", skillDetailLabelW, "Name:", strPtr(d.GetName()))
	if tags := d.GetTenantTags(); len(tags) > 0 {
		fmt.Printf("%-*s %s\n", skillDetailLabelW, "Tags:", strings.Join(tags, ", "))
	}
	desc := strPtr(d.GetDescription())
	if desc != "" {
		fmt.Printf("%-*s\n", skillDetailLabelW, "Description:")
		fmt.Println(wrapText(desc, 72, "  "))
	}
	return nil
}

var skillsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a skill from the cloud",
	Long: `Delete a skill permanently from the cloud by its skill ID.

The skill ID can be passed as a positional argument or via --skill-id flag.

Without --yes, the command first fetches the skill details and shows them before asking for confirmation.
With --yes, the skill detail lookup is skipped and the deletion is performed directly.

Examples:
  # Delete a skill using positional argument (interactive)
  agentbay skills delete skill-xxxxxxxxxxxxxxxx

  # Delete using positional argument, skip confirmation (for scripts/CI)
  agentbay skills delete skill-xxxxxxxxxxxxxxxx --yes
  agentbay skills delete skill-xxxxxxxxxxxxxxxx -y

  # Delete using named flag (compatible)
  agentbay skills delete --skill-id skill-xxxxxxxxxxxxxxxx
  agentbay skills delete --skill-id skill-xxxxxxxxxxxxxxxx --yes`,
	Args: cobra.MaximumNArgs(1),
	RunE: runSkillsDelete,
}

func runSkillsDelete(cmd *cobra.Command, args []string) error {
	skillId, _ := cmd.Flags().GetString("skill-id")
	// Support positional argument as alternative to --skill-id
	if skillId == "" && len(args) > 0 {
		skillId = args[0]
	}
	if skillId == "" {
		return fmt.Errorf("[ERROR] skill ID is required: provide it as a positional argument or via --skill-id")
	}
	autoYes, _ := cmd.Flags().GetBool("yes")

	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}
	apiClient := agentbay.NewClientFromConfig(cfg)
	ctx := context.Background()

	if !autoYes {
		// STEP 1/2: fetch skill detail to show info before confirmation
		fmt.Printf("[STEP 1/2] Fetching skill details...\n")
		req := &client.DescribeMarketSkillDetailRequest{SkillId: &skillId}
		resp, err := apiClient.DescribeMarketSkillDetail(ctx, req)
		if err != nil {
			printRequestIDFromErrIfVerbose(cmd, err)
			return fmt.Errorf("[ERROR] Failed to get skill details: %w", err)
		}
		if resp.Body != nil && resp.Body.RequestId != nil && *resp.Body.RequestId != "" {
			fmt.Printf("[INFO] DescribeMarketSkillDetail Request ID: %s\n", *resp.Body.RequestId)
		}
		if resp.Body == nil || resp.Body.Data == nil {
			return fmt.Errorf("[ERROR] Skill not found: %s", skillId)
		}
		d := resp.Body.Data
		skillName := strPtr(d.GetName())
		fmt.Printf("  SkillId: %s\n", skillId)
		if skillName != "" {
			fmt.Printf("  Name:    %s\n", skillName)
		}
		fmt.Println()

		// STEP 2/2: confirm deletion
		confirmed, err := ConfirmPrompt("Are you sure you want to permanently delete this skill? [y/N]: ", autoYes)
		if err != nil {
			return fmt.Errorf("[ERROR] %w", err)
		}
		if !confirmed {
			fmt.Printf("[INFO] Operation cancelled.\n")
			return nil
		}
	} else {
		fmt.Printf("[INFO] --yes specified, skipping skill detail lookup.\n")
	}

	// Delete
	deleteResp, err := apiClient.DeleteMarketSkill(ctx, &client.DeleteMarketSkillRequest{
		SkillId: &skillId,
	})
	if err != nil {
		printRequestIDFromErrIfVerbose(cmd, err)
		return fmt.Errorf("[ERROR] Failed to delete skill: %w", err)
	}

	if deleteResp.Body == nil {
		return fmt.Errorf("[ERROR] Invalid response: missing body")
	}

	if reqID := deleteResp.Body.GetRequestId(); reqID != "" {
		fmt.Printf("[INFO] DeleteMarketSkill Request ID: %s\n", reqID)
	}

	// Success判定：以 Code 为主依据，兼容 Success 缺失
	code := deleteResp.Body.GetCode()
	successPtr := deleteResp.Body.Success
	if (successPtr != nil && !*successPtr) || (code != "" && !strings.EqualFold(code, "ok")) {
		msg := deleteResp.Body.GetMessage()
		return fmt.Errorf("[ERROR] Failed to delete skill: Code=%s, Message=%s", code, msg)
	}

	fmt.Println()
	fmt.Printf("[SUCCESS] Skill has been deleted.\n")
	fmt.Printf("  SkillId: %s\n", skillId)

	return nil
}

func strPtr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// wrapText wraps s at word boundaries to at most width runes per line; each line is prefixed with indent.
func wrapText(s string, width int, indent string) string {
	if width <= 0 || s == "" {
		return indent + s
	}
	var b strings.Builder
	runes := []rune(s)
	start := 0
	for start < len(runes) {
		b.WriteString(indent)
		end := start + width
		if end > len(runes) {
			b.WriteString(string(runes[start:]))
			break
		}
		lastSpace := -1
		for i := start; i < end && i < len(runes); i++ {
			if runes[i] == ' ' || runes[i] == '\t' {
				lastSpace = i
			}
		}
		if lastSpace >= start {
			end = lastSpace + 1
		}
		b.WriteString(string(runes[start:end]))
		b.WriteByte('\n')
		start = end
		for start < len(runes) && (runes[start] == ' ' || runes[start] == '\t') {
			start++
		}
	}
	return strings.TrimSuffix(b.String(), "\n")
}
