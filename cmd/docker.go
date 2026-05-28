// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

// docker.go implements "agentbay docker login|tag|push|share|unshare|list-shares" commands.
// These wrap native docker CLI commands, with ACR credential management
// via the GetACRRepoCredential POP Action (raw HTTP, POP RPC V1).

package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/agentbay/agentbay-cli/internal/agentbay"
	"github.com/agentbay/agentbay-cli/internal/client"
	"github.com/agentbay/agentbay-cli/internal/config"
)

// ---------------------------------------------------------------------------
// Constants
// ---------------------------------------------------------------------------

const defaultRegistryURL = "ai-container-pre-9543-registry.cn-hangzhou.cr.aliyuncs.com"

// ---------------------------------------------------------------------------
// ACR credential cache
// ---------------------------------------------------------------------------

// acrCredentialCache is the on-disk structure saved after "docker login".
type acrCredentialCache struct {
	TempUsername       string `json:"temp_username"`
	AuthorizationToken string `json:"authorization_token"`
	Namespace          string `json:"namespace"`
	RepoName           string `json:"repo_name"`
	RegistryURL        string `json:"registry_url"`
	ImageTag           string `json:"image_tag"`
	ExpireTime         int64  `json:"expire_time"`
	CachedAt           string `json:"cached_at"`
}

func acrCachePath() (string, error) {
	dir, err := config.ConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "acr_credential.json"), nil
}

func saveACRCredential(c *acrCredentialCache) error {
	p, err := acrCachePath()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(p), 0755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(p, data, 0600)
}

func loadACRCredential() (*acrCredentialCache, error) {
	p, err := acrCachePath()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(p)
	if err != nil {
		return nil, fmt.Errorf("no cached ACR credential found. Run 'agentbay docker login' first: %w", err)
	}
	var c acrCredentialCache
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, fmt.Errorf("failed to parse cached ACR credential: %w", err)
	}
	return &c, nil
}

// ---------------------------------------------------------------------------
// Command tree
// ---------------------------------------------------------------------------

var DockerCmd = &cobra.Command{
	Use:     "docker",
	Short:   "Docker image build & push operations",
	Long:    "Manage Docker images for AgentBay: login to ACR, tag images, push images.",
	GroupID: "management",
}

// --- docker login ---

var dockerLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to the AgentBay ACR registry",
	Long: `Login to the AgentBay ACR image registry using temporary credentials.

This calls the GetACRRepoCredential API to obtain temporary Docker credentials,
then executes "docker login" with those credentials. The credential info
(RegistryUrl, Namespace, RepoName, ImageTag, etc.) is cached for subsequent
tag/push commands.

Examples:
  agentbay docker login`,
	Args: cobra.NoArgs,
	RunE: runDockerLogin,
}

// --- docker tag ---

var dockerTagCmd = &cobra.Command{
	Use:   "tag <source-image> <target-tag>",
	Short: "Tag a local image for the AgentBay ACR registry",
	Long: `Tag a local Docker image for pushing to the AgentBay ACR registry.

The target image name is automatically constructed as:
  $RegistryUrl/$Namespace/$RepoName:<target-tag>

You must run "agentbay docker login" first.

Examples:
  agentbay docker tag myapp:latest v1.0`,
	Args: cobra.ExactArgs(2),
	RunE: runDockerTag,
}

// --- docker push ---

var dockerPushCmd = &cobra.Command{
	Use:   "push <image>",
	Short: "Push a tagged image to the AgentBay ACR registry",
	Long: `Push a Docker image to the AgentBay ACR registry.

The image name must match the pattern $RegistryUrl/$Namespace/$RepoName[:tag].
If it does not, the push is rejected to prevent accidental pushes to wrong repos.

You must run "agentbay docker login" first.

Examples:
  agentbay docker push ai-container-pre-9543-registry.cn-hangzhou.cr.aliyuncs.com/ns/repo:v1.0`,
	Args: cobra.ExactArgs(1),
	RunE: runDockerPush,
}

func init() {
	DockerCmd.AddCommand(dockerLoginCmd)
	DockerCmd.AddCommand(dockerTagCmd)
	DockerCmd.AddCommand(dockerPushCmd)
	DockerCmd.AddCommand(dockerShareCmd)
	DockerCmd.AddCommand(dockerUnshareCmd)
	DockerCmd.AddCommand(dockerListSharesCmd)

	dockerShareCmd.Flags().Int64("target-uid", 0, "Target Alibaba Cloud account UID to share the Docker repo with")

	dockerUnshareCmd.Flags().Int64("target-uid", 0, "Target Alibaba Cloud account UID to cancel sharing with")

	dockerListSharesCmd.Flags().String("direction", "", `Sharing direction: "Outgoing" (repos you shared) or "Incoming" (repos shared with you)`)
	_ = dockerListSharesCmd.MarkFlagRequired("direction")
	dockerListSharesCmd.Flags().StringP("output", "o", "", `Output format. Use "json" for machine-readable output (e.g. for AI/scripts)`)
	dockerListSharesCmd.Flags().Int("page", 1, "Page number (default: 1)")
	dockerListSharesCmd.Flags().Int("size", 10, "Page size (default: 10)")
}

// ---------------------------------------------------------------------------
// GetACRRepoCredential response (reused from acr.go — but we add RegistryUrl)
// ---------------------------------------------------------------------------

type dockerACRCredentialResponse struct {
	RequestId      *string                          `json:"RequestId"`
	Code           *string                          `json:"Code"`
	Message        *string                          `json:"Message"`
	HttpStatusCode *int                             `json:"HttpStatusCode"`
	Success        *bool                            `json:"Success"`
	Data           *dockerACRCredentialResponseData `json:"Data"`
}

type dockerACRCredentialResponseData struct {
	IsSuccess          *bool   `json:"IsSuccess"`
	Code               *string `json:"Code"`
	RequestId          *string `json:"RequestId"`
	TempUsername       *string `json:"TempUsername"`
	AuthorizationToken *string `json:"AuthorizationToken"`
	Namespace          *string `json:"Namespace"`
	RepoName           *string `json:"RepoName"`
	RegistryUrl        *string `json:"RegistryUrl"`
	ImageTag           *string `json:"ImageTag"`
	ExpireTime         *int64  `json:"ExpireTime"`
}

// ---------------------------------------------------------------------------
// docker login
// ---------------------------------------------------------------------------

func runDockerLogin(cobraCmd *cobra.Command, args []string) error {
	// 1. Auth & API call
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

	body, statusCode, err := client.callRPC("GetACRRepoCredential", map[string]string{})
	if err != nil {
		return fmt.Errorf("[ERROR] Request failed: %w", err)
	}

	if statusCode < 200 || statusCode >= 300 {
		return fmt.Errorf("[ERROR] API returned HTTP %d: %s", statusCode, string(body))
	}

	var resp dockerACRCredentialResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return fmt.Errorf("[ERROR] Failed to parse response: %w", err)
	}

	if resp.Success != nil && !*resp.Success {
		return fmt.Errorf("[ERROR] API error: Code=%s, Message=%s", ptrStr(resp.Code), ptrStr(resp.Message))
	}
	if resp.Data == nil {
		return fmt.Errorf("[ERROR] API returned empty Data")
	}

	d := resp.Data
	tempUsername := ptrStr(d.TempUsername)
	authToken := ptrStr(d.AuthorizationToken)
	namespace := ptrStr(d.Namespace)
	repoName := ptrStr(d.RepoName)
	imageTag := ptrStr(d.ImageTag)
	registryURL := ptrStr(d.RegistryUrl)

	// Fallback to default RegistryUrl
	if registryURL == "" || registryURL == "<nil>" {
		registryURL = defaultRegistryURL
		log.Debugf("[DOCKER LOGIN] RegistryUrl not returned, using default: %s", registryURL)
	}

	var expireTime int64
	if d.ExpireTime != nil {
		expireTime = *d.ExpireTime
	}

	// 2. Print key credential info
	fullRegistryPath := fmt.Sprintf("%s/%s/%s", registryURL, namespace, repoName)
	if expireTime > 0 {
		expireAt := time.Unix(expireTime/1000, 0)
		fmt.Printf("Credential expires at: %s\n", expireAt.Format("2006-01-02 15:04:05"))
	}
	fmt.Printf("Image registry path:   %s\n", fullRegistryPath)

	// 3. Cache credentials
	cache := &acrCredentialCache{
		TempUsername:       tempUsername,
		AuthorizationToken: authToken,
		Namespace:          namespace,
		RepoName:           repoName,
		RegistryURL:        registryURL,
		ImageTag:           imageTag,
		ExpireTime:         expireTime,
		CachedAt:           time.Now().Format(time.RFC3339),
	}
	if err := saveACRCredential(cache); err != nil {
		log.Debugf("Failed to cache credential: %v", err)
	}

	// 4. Execute: echo "$AuthorizationToken" | docker login $RegistryUrl -u "$TempUsername" --password-stdin
	fmt.Println("[DOCKER LOGIN] Logging in via 'docker'...")
	dockerCmd := exec.Command("docker", "login", registryURL, "-u", tempUsername, "--password-stdin")
	dockerCmd.Stdin = strings.NewReader(authToken)
	dockerCmd.Stdout = os.Stdout
	dockerCmd.Stderr = os.Stderr

	if err := dockerCmd.Run(); err != nil {
		return fmt.Errorf("docker login failed: %w", err)
	}

	// Also try: echo "$AuthorizationToken" | sudo -n docker login $RegistryUrl -u "$TempUsername" --password-stdin
	// Compatible with rootful docker daemon. Failure here is non-fatal.
	if _, lookErr := exec.LookPath("sudo"); lookErr == nil {
		fmt.Println("[DOCKER LOGIN] Logging in via 'sudo docker'...")
		sudoDockerCmd := exec.Command("sudo", "-n", "docker", "login", registryURL, "-u", tempUsername, "--password-stdin")
		sudoDockerCmd.Stdin = strings.NewReader(authToken)
		sudoDockerCmd.Stdout = os.Stdout
		sudoDockerCmd.Stderr = os.Stderr

		if err := sudoDockerCmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "[WARN] sudo docker login failed (skipped): %v\n", err)
		}
	}

	fmt.Println()
	if expireTime > 0 {
		fmt.Println("Note: Credentials will expire after the time above. You can run 'agentbay docker login' again to refresh.")
	}
	fmt.Printf("Note: When tagging images, use: %s:<your-tag>\n", fullRegistryPath)
	return nil
}

// ---------------------------------------------------------------------------
// docker tag
// ---------------------------------------------------------------------------

func runDockerTag(cobraCmd *cobra.Command, args []string) error {
	sourceImage := args[0]
	targetTag := args[1]

	// Load cached credentials
	cache, err := loadACRCredential()
	if err != nil {
		return fmt.Errorf("[ERROR] %w", err)
	}

	// Construct target: $RegistryUrl/$Namespace/$RepoName:<tag>
	targetImage := fmt.Sprintf("%s/%s/%s:%s", cache.RegistryURL, cache.Namespace, cache.RepoName, targetTag)

	fmt.Printf("[DOCKER TAG] Tagging image...\n")
	fmt.Printf("  Source: %s\n", sourceImage)
	fmt.Printf("  Target: %s\n", targetImage)

	// Debug: print full registry path
	fmt.Printf("\n[DEBUG] Full registry path: %s/%s/%s\n", cache.RegistryURL, cache.Namespace, cache.RepoName)

	// Execute: docker tag <source> <target>
	dockerCmd := exec.Command("docker", "tag", sourceImage, targetImage)
	dockerCmd.Stdout = os.Stdout
	dockerCmd.Stderr = os.Stderr

	if err := dockerCmd.Run(); err != nil {
		return fmt.Errorf("[ERROR] docker tag failed: %w", err)
	}

	fmt.Printf("\n[SUCCESS] Image tagged as: %s\n", targetImage)
	return nil
}

// ---------------------------------------------------------------------------
// docker push
// ---------------------------------------------------------------------------

func runDockerPush(cobraCmd *cobra.Command, args []string) error {
	pushImage := args[0]

	// Load cached credentials
	cache, err := loadACRCredential()
	if err != nil {
		return fmt.Errorf("[ERROR] %w", err)
	}

	// Validate: image must match $RegistryUrl/$Namespace/$RepoName
	expectedPrefix := fmt.Sprintf("%s/%s/%s", cache.RegistryURL, cache.Namespace, cache.RepoName)
	if !strings.HasPrefix(pushImage, expectedPrefix) {
		return fmt.Errorf("[ERROR] Image name '%s' does not match the authorized registry path.\n"+
			"  Expected prefix: %s\n"+
			"  Use 'agentbay docker tag' to tag your image correctly first",
			pushImage, expectedPrefix)
	}

	fmt.Printf("[DOCKER PUSH] Pushing image...\n")
	fmt.Printf("  Image: %s\n", pushImage)

	// Execute: docker push <image>
	dockerCmd := exec.Command("docker", "push", pushImage)
	dockerCmd.Stdout = os.Stdout
	dockerCmd.Stderr = os.Stderr

	if err := dockerCmd.Run(); err != nil {
		return fmt.Errorf("[ERROR] docker push failed: %w", err)
	}

	fmt.Println("\n[SUCCESS] Image pushed.")
	return nil
}

// ---------------------------------------------------------------------------
// docker share
// ---------------------------------------------------------------------------

var dockerShareCmd = &cobra.Command{
	Use:   "share [<target-uid>]",
	Short: "Share the Docker repo with another Alibaba Cloud account",
	Long: `Share your AgentBay Docker image repository with another Alibaba Cloud account.

The target account will be able to pull images from your repository.

Examples:
  agentbay docker share 1234567890
  agentbay docker share --target-uid 1234567890`,
	Args: cobra.MaximumNArgs(1),
	RunE: runDockerShare,
}

func runDockerShare(cobraCmd *cobra.Command, args []string) error {
	targetUID, _ := cobraCmd.Flags().GetInt64("target-uid")

	// Prefer positional arg if provided, otherwise fall back to flag
	if len(args) > 0 {
		parsed, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("[ERROR] Invalid target-uid %q: %w", args[0], err)
		}
		targetUID = parsed
	}

	if targetUID == 0 {
		return fmt.Errorf("[ERROR] target-uid is required. Provide it as a positional argument or use --target-uid")
	}

	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to load configuration: %w", err)
	}

	apiClient := agentbay.NewClientFromConfig(cfg)
	ctx := context.Background()

	fmt.Printf("[STEP 1/1] Sharing Docker repo with UID %d...\n", targetUID)

	req := &client.ShareDockerRepoRequest{TargetAliUid: &targetUID}
	resp, err := apiClient.ShareDockerRepo(ctx, req)
	if err != nil {
		if reqID := extractRequestIDFromErr(err); reqID != "" {
			fmt.Printf("[INFO] ShareDockerRepo Request ID: %s\n", reqID)
		}
		return fmt.Errorf("[ERROR] Failed to share Docker repo: %w", err)
	}

	if resp.Body != nil {
		if reqID := resp.Body.GetRequestId(); reqID != "" {
			fmt.Printf("[INFO] ShareDockerRepo Request ID: %s\n", reqID)
		}
	}

	if resp.Body == nil {
		return fmt.Errorf("[ERROR] Invalid response: missing body")
	}

	code := resp.Body.GetCode()
	successPtr := resp.Body.Success
	if (successPtr != nil && !*successPtr) || (code != "" && !strings.EqualFold(code, "ok")) {
		msg := resp.Body.GetMessage()
		return fmt.Errorf("[ERROR] Failed to share Docker repo: Code=%s, Message=%s", code, msg)
	}

	fmt.Println()
	fmt.Printf("[SUCCESS] Docker repo shared successfully!\n")
	if resp.Body.Data != nil {
		d := resp.Body.Data
		if d.TargetAliUid != nil {
			fmt.Printf("  TargetAliUid : %d\n", *d.TargetAliUid)
		}
		if d.OwnerAliUid != nil {
			fmt.Printf("  OwnerAliUid  : %d\n", *d.OwnerAliUid)
		}
		if d.AcrRepoName != nil {
			fmt.Printf("  AcrRepoName  : %s\n", *d.AcrRepoName)
		}
		if d.Status != nil {
			fmt.Printf("  Status       : %s\n", *d.Status)
		}
	}
	return nil
}

// ---------------------------------------------------------------------------
// docker unshare
// ---------------------------------------------------------------------------

var dockerUnshareCmd = &cobra.Command{
	Use:   "unshare [<target-uid>]",
	Short: "Cancel sharing the Docker repo with another Alibaba Cloud account",
	Long: `Cancel sharing your AgentBay Docker image repository with a specific Alibaba Cloud account.

Examples:
  agentbay docker unshare 1234567890
  agentbay docker unshare --target-uid 1234567890`,
	Args: cobra.MaximumNArgs(1),
	RunE: runDockerUnshare,
}

func runDockerUnshare(cobraCmd *cobra.Command, args []string) error {
	targetUID, _ := cobraCmd.Flags().GetInt64("target-uid")

	// Prefer positional arg if provided, otherwise fall back to flag
	if len(args) > 0 {
		parsed, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("[ERROR] Invalid target-uid %q: %w", args[0], err)
		}
		targetUID = parsed
	}

	if targetUID == 0 {
		return fmt.Errorf("[ERROR] target-uid is required. Provide it as a positional argument or use --target-uid")
	}

	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to load configuration: %w", err)
	}

	apiClient := agentbay.NewClientFromConfig(cfg)
	ctx := context.Background()

	fmt.Printf("[STEP 1/1] Cancelling Docker repo sharing with UID %d...\n", targetUID)

	req := &client.UnshareDockerRepoRequest{TargetAliUid: &targetUID}
	resp, err := apiClient.UnshareDockerRepo(ctx, req)
	if err != nil {
		if reqID := extractRequestIDFromErr(err); reqID != "" {
			fmt.Printf("[INFO] UnshareDockerRepo Request ID: %s\n", reqID)
		}
		return fmt.Errorf("[ERROR] Failed to cancel Docker repo sharing: %w", err)
	}

	if resp.Body != nil {
		if reqID := resp.Body.GetRequestId(); reqID != "" {
			fmt.Printf("[INFO] UnshareDockerRepo Request ID: %s\n", reqID)
		}
	}

	if resp.Body == nil {
		return fmt.Errorf("[ERROR] Invalid response: missing body")
	}

	code := resp.Body.GetCode()
	successPtr := resp.Body.Success
	if (successPtr != nil && !*successPtr) || (code != "" && !strings.EqualFold(code, "ok")) {
		msg := resp.Body.GetMessage()
		return fmt.Errorf("[ERROR] Failed to cancel Docker repo sharing: Code=%s, Message=%s", code, msg)
	}

	revoked := false
	if resp.Body.Data != nil && resp.Body.Data.Revoked != nil {
		revoked = *resp.Body.Data.Revoked
	}

	fmt.Println()
	fmt.Printf("[SUCCESS] Docker repo sharing cancelled.\n")
	fmt.Printf("  Revoked : %v\n", revoked)
	return nil
}

// ---------------------------------------------------------------------------
// docker list-shares
// ---------------------------------------------------------------------------

var dockerListSharesCmd = &cobra.Command{
	Use:   "list-shares",
	Short: "List Docker repo sharing information",
	Long: `List Docker image repository sharing information.

Use --direction to specify:
  Outgoing  Repos you have shared with other accounts
  Incoming  Repos that other accounts have shared with you

Use --page and --size to paginate results (both optional, default page=1, size=10).

Examples:
  agentbay docker list-shares --direction Outgoing
  agentbay docker list-shares --direction Incoming
  agentbay docker list-shares --direction Outgoing --page 2 --size 5
  agentbay docker list-shares --direction Outgoing --output json`,
	Args: cobra.NoArgs,
	RunE: runDockerListShares,
}

func runDockerListShares(cobraCmd *cobra.Command, args []string) error {
	direction, _ := cobraCmd.Flags().GetString("direction")
	outputFmt, _ := cobraCmd.Flags().GetString("output")
	page, _ := cobraCmd.Flags().GetInt("page")
	size, _ := cobraCmd.Flags().GetInt("size")

	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to load configuration: %w", err)
	}

	apiClient := agentbay.NewClientFromConfig(cfg)
	ctx := context.Background()

	req := &client.ListSharedDockerReposRequest{Direction: &direction}
	if size > 0 {
		sizeInt32 := int32(size)
		req.PageSize = &sizeInt32
	}
	if page > 0 {
		pageInt32 := int32(page)
		req.PageStart = &pageInt32
	}
	resp, err := apiClient.ListSharedDockerRepos(ctx, req)
	if err != nil {
		if reqID := extractRequestIDFromErr(err); reqID != "" {
			fmt.Printf("[INFO] ListSharedDockerRepos Request ID: %s\n", reqID)
		}
		return fmt.Errorf("[ERROR] Failed to list shared Docker repos: %w", err)
	}

	if resp.Body != nil {
		if reqID := resp.Body.GetRequestId(); reqID != "" {
			fmt.Printf("[INFO] ListSharedDockerRepos Request ID: %s\n", reqID)
		}
	}

	if resp.Body == nil {
		return fmt.Errorf("[ERROR] Invalid response: missing body")
	}

	code := resp.Body.GetCode()
	successPtr := resp.Body.Success
	if (successPtr != nil && !*successPtr) || (code != "" && !strings.EqualFold(code, "ok")) {
		msg := resp.Body.GetMessage()
		return fmt.Errorf("[ERROR] Failed to list shared Docker repos: Code=%s, Message=%s", code, msg)
	}

	items := resp.Body.Data
	if items == nil {
		items = []*client.ListSharedDockerReposResponseBodyDataItem{}
	}

	if strings.EqualFold(outputFmt, "json") {
		type itemJSON struct {
			PeerAliUid int64  `json:"peerAliUid"`
			Status     string `json:"status"`
		}
		type outputJSON struct {
			TotalCount int        `json:"totalCount"`
			PageNumber int        `json:"pageNumber"`
			PageSize   int        `json:"pageSize"`
			Items      []itemJSON `json:"items"`
		}
		out := outputJSON{
			TotalCount: len(items),
			PageNumber: page,
			PageSize:   size,
			Items:      []itemJSON{},
		}
		for _, item := range items {
			j := itemJSON{}
			if item.PeerAliUid != nil {
				j.PeerAliUid = *item.PeerAliUid
			}
			if item.Status != nil {
				j.Status = *item.Status
			}
			out.Items = append(out.Items, j)
		}
		b, err := json.MarshalIndent(out, "", "  ")
		if err != nil {
			return fmt.Errorf("json marshal: %w", err)
		}
		fmt.Println(string(b))
		return nil
	}

	// Table output
	if len(items) == 0 {
		fmt.Printf("No shared Docker repos found for direction: %s\n", direction)
		return nil
	}
	fmt.Printf("%-20s  %-15s\n", "PeerAliUid", "Status")
	fmt.Printf("%-20s  %-15s\n", "--------------------", "---------------")
	for _, item := range items {
		uid := int64(0)
		if item.PeerAliUid != nil {
			uid = *item.PeerAliUid
		}
		status := ""
		if item.Status != nil {
			status = *item.Status
		}
		fmt.Printf("%-20d  %-15s\n", uid, status)
	}
	fmt.Printf("\nTotal: %d\n", len(items))
	return nil
}
