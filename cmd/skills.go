// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"

	"github.com/agentbay/agentbay-cli/internal/agentbay"
	"github.com/agentbay/agentbay-cli/internal/client"
	"github.com/agentbay/agentbay-cli/internal/config"
)

const skillFileName = "SKILL.md"

// Output style: label width for skill show.
const skillDetailLabelW = 14

var SkillsCmd = &cobra.Command{
	Use:     "skills",
	Short:   "Manage AgentBay skills",
	Long:    "Push skills and show details.",
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

var skillsShowCmd = &cobra.Command{
	Use:   "show <skill-id>",
	Short: "Show skill details",
	Long:  `Show details for a skill by ID.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runSkillsShow,
}

func init() {
	SkillsCmd.AddCommand(skillsPushCmd)
	SkillsCmd.AddCommand(skillsShowCmd)
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

	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}
	apiClient := agentbay.NewClientFromConfig(cfg)
	ctx := context.Background()

	fmt.Printf("[STEP 1/3] Getting upload credential...\n")
	credReq := &client.GetMarketSkillCredentialRequest{FileName: &skillZipName}
	credResp, err := apiClient.GetMarketSkillCredential(ctx, credReq)
	if err != nil {
		printRequestIDFromErrIfVerbose(cmd, err)
		fmt.Fprintf(os.Stderr, "[ERROR] Failed to get upload credential: %v\n", err)
		return fmt.Errorf("get credential: %w", err)
	}
	if credResp.Body == nil || credResp.Body.Data == nil {
		return fmt.Errorf("invalid response: missing credential data")
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
		fmt.Printf("[STEP 2/3] Uploading skill zip...\n")
		if verbose, _ := cmd.Flags().GetBool("verbose"); verbose {
			if fi, err := os.Stat(zipPath); err == nil {
				fmt.Fprintf(os.Stderr, "[DEBUG] Upload size: %d bytes, file: %s\n", fi.Size(), zipPath)
			}
		}
		if err := uploadFileToOSS(zipPath, uploadURLStr); err != nil {
			fmt.Fprintf(os.Stderr, "[ERROR] Failed to upload: %v\n", err)
			return fmt.Errorf("upload: %w", err)
		}
	} else {
		// Directory: pack then upload
		fmt.Printf("[STEP 2/3] Packing and uploading skill...\n")
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
			fmt.Fprintf(os.Stderr, "[ERROR] Failed to upload: %v\n", err)
			return fmt.Errorf("upload: %w", err)
		}
	}

	fmt.Printf("[STEP 3/3] Creating skill...\n")
	createReq := &client.CreateMarketSkillRequest{
		OssBucket:   &createBucket,
		OssFilePath: &createOssFilePath,
	}
	createResp, err := apiClient.CreateMarketSkill(ctx, createReq)
	if err != nil {
		if createResp != nil && createResp.RawBody != "" {
			fmt.Fprintf(os.Stderr, "[DEBUG] Raw response: %s\n", createResp.RawBody)
		}
		printRequestIDFromErrIfVerbose(cmd, err)
		fmt.Fprintf(os.Stderr, "[ERROR] Failed to create skill: %v\n", err)
		return fmt.Errorf("create skill: %w", err)
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
		fmt.Fprintf(os.Stderr, "[ERROR] Failed to get skill details: %v\n", err)
		return fmt.Errorf("describe skill: %w", err)
	}
	if resp.Body == nil || resp.Body.Data == nil {
		fmt.Fprintf(os.Stderr, "[INFO] No details for skill %s\n", skillId)
		return nil
	}
	verbose, _ := cmd.Flags().GetBool("verbose")
	if verbose && resp.Body.RequestId != nil && *resp.Body.RequestId != "" {
		printRequestIDIfVerbose(cmd, *resp.Body.RequestId)
	}
	d := resp.Body.Data
	displaySkillId := strPtr(d.GetSkillId())
	if displaySkillId == "" {
		displaySkillId = skillId
	}
	fmt.Printf("%-*s %s\n", skillDetailLabelW, "SkillId:", displaySkillId)
	fmt.Printf("%-*s %s\n", skillDetailLabelW, "Name:", strPtr(d.GetName()))
	desc := strPtr(d.GetDescription())
	if desc != "" {
		fmt.Printf("%-*s\n", skillDetailLabelW, "Description:")
		fmt.Println(wrapText(desc, 72, "  "))
	}
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

