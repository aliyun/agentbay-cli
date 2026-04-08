// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package agentbay

import (
	"context"
	"fmt"
	"time"

	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"

	"github.com/agentbay/agentbay-cli/internal/auth"
	"github.com/agentbay/agentbay-cli/internal/client"
	"github.com/agentbay/agentbay-cli/internal/config"
)

// Client interface defines the methods available for AgentBay API operations
type Client interface {
	GetDockerFileStoreCredential(ctx context.Context, request *client.GetDockerFileStoreCredentialRequest) (*client.GetDockerFileStoreCredentialResponse, error)
	CreateDockerImageTask(ctx context.Context, request *client.CreateDockerImageTaskRequest) (*client.CreateDockerImageTaskResponse, error)
	GetDockerImageTask(ctx context.Context, request *client.GetDockerImageTaskRequest) (*client.GetDockerImageTaskResponse, error)
	ListMcpImages(ctx context.Context, request *client.ListMcpImagesRequest) (*client.ListMcpImagesResponse, error)
	GetMcpImageInfo(ctx context.Context, request *client.GetMcpImageInfoRequest) (*client.GetMcpImageInfoResponse, error)
	CreateResourceGroup(ctx context.Context, request *client.CreateResourceGroupRequest) (*client.CreateResourceGroupResponse, error)
	DeleteResourceGroup(ctx context.Context, request *client.DeleteResourceGroupRequest) (*client.DeleteResourceGroupResponse, error)
	GetDockerfileTemplate(ctx context.Context, request *client.GetDockerfileTemplateRequest) (*client.GetDockerfileTemplateResponse, error)
	// Market Skill
	GetMarketSkillCredential(ctx context.Context, request *client.GetMarketSkillCredentialRequest) (*client.GetMarketSkillCredentialResponse, error)
	CreateMarketSkill(ctx context.Context, request *client.CreateMarketSkillRequest) (*client.CreateMarketSkillResponse, error)
	DescribeMarketSkillDetail(ctx context.Context, request *client.DescribeMarketSkillDetailRequest) (*client.DescribeMarketSkillDetailResponse, error)
	// API Key
	CreateApiKey(ctx context.Context, request *client.CreateApiKeyRequest) (*client.CreateApiKeyResponse, error)
}

// clientWrapper wraps the generated SDK client with additional functionality
type clientWrapper struct {
	apiConfig *config.APIConfig
	config    *config.Config
	client    *client.Client
}

// NewClient creates a new client wrapper with the given API configuration and config
func NewClient(apiConfig *config.APIConfig, cfg *config.Config) Client {
	return &clientWrapper{
		apiConfig: apiConfig,
		config:    cfg,
	}
}

// NewClientFromConfig creates a new client wrapper using default API configuration
func NewClientFromConfig(cfg *config.Config) Client {
	apiConfig := config.LoadAPIConfig(nil)
	return &clientWrapper{
		apiConfig: &apiConfig,
		config:    cfg,
	}
}

// getClient returns the underlying SDK client, creating it if necessary
func (cw *clientWrapper) getClient() (*client.Client, error) {
	if ak, sk, session, ok := config.AccessKeyFromEnv(); ok {
		return newSDKClientWithAccessKeys(cw.apiConfig, ak, sk, session)
	}

	// Refresh token if needed (checks expiry and refreshes automatically)
	// Create an adapter to bridge config.Config to auth.TokenConfig
	tokenCfgAdapter := auth.NewConfigAdapter(
		func() (string, string, time.Time, error) {
			return cw.config.GetTokens()
		},
		cw.config.RefreshTokens,
		cw.config.IsTokenExpired,
		cw.config.ClearTokens,
	)

	err := auth.RefreshTokenIfNeeded(tokenCfgAdapter, config.GetClientID())
	if err != nil {
		return nil, fmt.Errorf("failed to ensure valid token: %w", err)
	}

	token, err := cw.config.GetToken()
	if err != nil {
		return nil, fmt.Errorf("failed to get authentication token: %w", err)
	}

	endpoint := cw.apiConfig.Endpoint
	openapiConfig := &openapiutil.Config{
		// Use BearerToken for OAuth authentication (backend now supports BearerToken)
		BearerToken:    dara.String(token.AccessToken),
		Endpoint:       dara.String(endpoint),
		ReadTimeout:    dara.Int(cw.apiConfig.TimeoutMs),
		ConnectTimeout: dara.Int(cw.apiConfig.TimeoutMs),
		UserAgent:      dara.String("AgentBay-CLI/1.0"),
	}

	sdkClient, err := client.NewClient(openapiConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create API client: %w", err)
	}

	return sdkClient, nil
}

// getRuntimeOptions returns default runtime options for SDK calls.
func (cw *clientWrapper) getRuntimeOptions() *dara.RuntimeOptions {
	return &dara.RuntimeOptions{}
}

// GetDockerFileStoreCredential wraps the SDK client method
func (cw *clientWrapper) GetDockerFileStoreCredential(ctx context.Context, request *client.GetDockerFileStoreCredentialRequest) (*client.GetDockerFileStoreCredentialResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	return sdkClient.GetDockerFileStoreCredentialWithOptions(request, cw.getRuntimeOptions())
}

// GetMarketSkillCredential wraps the SDK client method
func (cw *clientWrapper) GetMarketSkillCredential(ctx context.Context, request *client.GetMarketSkillCredentialRequest) (*client.GetMarketSkillCredentialResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	return sdkClient.GetMarketSkillCredentialWithOptions(request, cw.getRuntimeOptions())
}

// CreateMarketSkill wraps the SDK client method
func (cw *clientWrapper) CreateMarketSkill(ctx context.Context, request *client.CreateMarketSkillRequest) (*client.CreateMarketSkillResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	return sdkClient.CreateMarketSkillWithOptions(request, cw.getRuntimeOptions())
}

// DescribeMarketSkillDetail wraps the SDK client method
func (cw *clientWrapper) DescribeMarketSkillDetail(ctx context.Context, request *client.DescribeMarketSkillDetailRequest) (*client.DescribeMarketSkillDetailResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	return sdkClient.DescribeMarketSkillDetailWithOptions(request, cw.getRuntimeOptions())
}

// CreateDockerImageTask wraps the SDK client method
func (cw *clientWrapper) CreateDockerImageTask(ctx context.Context, request *client.CreateDockerImageTaskRequest) (*client.CreateDockerImageTaskResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	return sdkClient.CreateDockerImageTaskWithContext(ctx, request, cw.getRuntimeOptions())
}

// GetDockerImageTask wraps the SDK client method
func (cw *clientWrapper) GetDockerImageTask(ctx context.Context, request *client.GetDockerImageTaskRequest) (*client.GetDockerImageTaskResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	return sdkClient.GetDockerImageTaskWithContext(ctx, request, cw.getRuntimeOptions())
}

// ListMcpImages wraps the SDK client method
func (cw *clientWrapper) ListMcpImages(ctx context.Context, request *client.ListMcpImagesRequest) (*client.ListMcpImagesResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	return sdkClient.ListMcpImagesWithContext(ctx, request, cw.getRuntimeOptions())
}

// CreateResourceGroup wraps the SDK client method
func (cw *clientWrapper) CreateResourceGroup(ctx context.Context, request *client.CreateResourceGroupRequest) (*client.CreateResourceGroupResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	return sdkClient.CreateResourceGroupWithContext(ctx, request, cw.getRuntimeOptions())
}

// DeleteResourceGroup wraps the SDK client method
func (cw *clientWrapper) DeleteResourceGroup(ctx context.Context, request *client.DeleteResourceGroupRequest) (*client.DeleteResourceGroupResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	return sdkClient.DeleteResourceGroupWithContext(ctx, request, cw.getRuntimeOptions())
}

// GetMcpImageInfo wraps the SDK client method
func (cw *clientWrapper) GetMcpImageInfo(ctx context.Context, request *client.GetMcpImageInfoRequest) (*client.GetMcpImageInfoResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	return sdkClient.GetMcpImageInfoWithContext(ctx, request, cw.getRuntimeOptions())
}

// GetDockerfileTemplate wraps the SDK client method
func (cw *clientWrapper) GetDockerfileTemplate(ctx context.Context, request *client.GetDockerfileTemplateRequest) (*client.GetDockerfileTemplateResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	return sdkClient.GetDockerfileTemplateWithContext(ctx, request, cw.getRuntimeOptions())
}

// CreateApiKey wraps the SDK client method
func (cw *clientWrapper) CreateApiKey(ctx context.Context, request *client.CreateApiKeyRequest) (*client.CreateApiKeyResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	return sdkClient.CreateApiKeyWithContext(ctx, request, cw.getRuntimeOptions())
}
