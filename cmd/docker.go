// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

// docker.go implements "agentbay docker login|tag|push" commands.
// These wrap native docker CLI commands, with ACR credential management
// via the GetACRRepoCredential POP Action (raw HTTP, POP RPC V1).

package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

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
