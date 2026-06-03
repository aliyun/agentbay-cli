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
	DeleteMcpImage(ctx context.Context, request *client.DeleteMcpImageRequest) (*client.DeleteMcpImageResponse, error)
	GetDockerfileTemplate(ctx context.Context, request *client.GetDockerfileTemplateRequest) (*client.GetDockerfileTemplateResponse, error)
	// Market Skill
	GetMarketSkillCredential(ctx context.Context, request *client.GetMarketSkillCredentialRequest) (*client.GetMarketSkillCredentialResponse, error)
	CreateMarketSkill(ctx context.Context, request *client.CreateMarketSkillRequest) (*client.CreateMarketSkillResponse, error)
	UpdateMarketSkill(ctx context.Context, request *client.UpdateMarketSkillRequest) (*client.CreateMarketSkillResponse, error)
	DescribeMarketSkillDetail(ctx context.Context, request *client.DescribeMarketSkillDetailRequest) (*client.DescribeMarketSkillDetailResponse, error)
	ListMarketSkillByPage(ctx context.Context, request *client.ListMarketSkillByPageRequest) (*client.ListMarketSkillByPageResponse, error)
	DeleteMarketSkill(ctx context.Context, request *client.DeleteMarketSkillRequest) (*client.DeleteMarketSkillResponse, error)
	// Tags
	ListTag(ctx context.Context) (*client.ListTagResponse, error)
	CreateTag(ctx context.Context, request *client.CreateTagRequest) (*client.CreateTagResponse, error)
	// API Key
	CreateApiKey(ctx context.Context, request *client.CreateApiKeyRequest) (*client.CreateApiKeyResponse, error)
	ModifyMcpApiKeyConfig(ctx context.Context, request *client.ModifyMcpApiKeyConfigRequest) (*client.ModifyMcpApiKeyConfigResponse, error)
	DescribeMcpApiKey(ctx context.Context, request *client.DescribeMcpApiKeyRequest) (*client.DescribeMcpApiKeyResponse, error)
	ModifyApiKeyStatus(ctx context.Context, request *client.ModifyApiKeyStatusRequest) (*client.ModifyApiKeyStatusResponse, error)
	DeleteApiKey(ctx context.Context, request *client.DeleteApiKeyRequest) (*client.DeleteApiKeyResponse, error)
	DescribeApiKeys(ctx context.Context, request *client.DescribeApiKeysRequest) (*client.DescribeApiKeysResponse, error)
	DescribeKeyContent(ctx context.Context, request *client.DescribeKeyContentRequest) (*client.DescribeKeyContentResponse, error)
	// Advanced Network
	DescribeInstanceTypes(ctx context.Context, request *client.DescribeInstanceTypesRequest) (*client.DescribeInstanceTypesResponse, error)
	DescribeMcpPolicyData(ctx context.Context, request *client.DescribeMcpPolicyDataRequest) (*client.DescribeMcpPolicyDataResponse, error)
	SaveMcpPolicyData(ctx context.Context, request *client.SaveMcpPolicyDataRequest) (*client.SaveMcpPolicyDataResponse, error)
	DescribeOfficeSites(ctx context.Context, request *client.DescribeOfficeSitesRequest) (*client.DescribeOfficeSitesResponse, error)
	// Policy Data Create/Modify
	CreateMcpPolicyData(ctx context.Context, request *client.CreateModifyMcpPolicyDataRequest) (*client.CreateMcpPolicyDataResponse, error)
	ModifyMcpPolicyData(ctx context.Context, request *client.CreateModifyMcpPolicyDataRequest) (*client.ModifyMcpPolicyDataResponse, error)
	// Network Packages
	DescribeNetworkPackages(ctx context.Context, request *client.DescribeNetworkPackagesRequest) (*client.DescribeNetworkPackagesResponse, error)
	// Resource Group Max Session
	BatchCreateHideResourceGroupsWithMaxSession(ctx context.Context, request *client.BatchCreateHideResourceGroupsWithMaxSessionRequest) (*client.BatchCreateHideResourceGroupsWithMaxSessionResponse, error)
	// WarmUp Status
	DescribeWarmUpStatusOpen(ctx context.Context, request *client.DescribeWarmUpStatusOpenRequest) (*client.DescribeWarmUpStatusOpenResponse, error)
	// Docker Repo Sharing
	ShareDockerRepo(ctx context.Context, request *client.ShareDockerRepoRequest) (*client.ShareDockerRepoResponse, error)
	UnshareDockerRepo(ctx context.Context, request *client.UnshareDockerRepoRequest) (*client.UnshareDockerRepoResponse, error)
	ListSharedDockerRepos(ctx context.Context, request *client.ListSharedDockerReposRequest) (*client.ListSharedDockerReposResponse, error)
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

// UpdateMarketSkill wraps the SDK client method
func (cw *clientWrapper) UpdateMarketSkill(ctx context.Context, request *client.UpdateMarketSkillRequest) (*client.CreateMarketSkillResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	resp, err := sdkClient.UpdateMarketSkillWithOptions(request, cw.getRuntimeOptions())
	if err != nil {
		return nil, err
	}
	// Convert UpdateMarketSkillResponse to CreateMarketSkillResponse (same Body structure)
	return &client.CreateMarketSkillResponse{
		Headers:    resp.Headers,
		StatusCode: resp.StatusCode,
		Body:       resp.Body,
		RawBody:    resp.RawBody,
	}, nil
}

// DescribeMarketSkillDetail wraps the SDK client method
func (cw *clientWrapper) DescribeMarketSkillDetail(ctx context.Context, request *client.DescribeMarketSkillDetailRequest) (*client.DescribeMarketSkillDetailResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	return sdkClient.DescribeMarketSkillDetailWithOptions(request, cw.getRuntimeOptions())
}

// ListTag wraps the SDK client method
func (cw *clientWrapper) ListTag(ctx context.Context) (*client.ListTagResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	return sdkClient.ListTagWithOptions(cw.getRuntimeOptions())
}

// ListMarketSkillByPage wraps the SDK client method
func (cw *clientWrapper) ListMarketSkillByPage(ctx context.Context, request *client.ListMarketSkillByPageRequest) (*client.ListMarketSkillByPageResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	return sdkClient.ListMarketSkillByPageWithOptions(request, cw.getRuntimeOptions())
}

// DeleteMarketSkill wraps the SDK client method
func (cw *clientWrapper) DeleteMarketSkill(ctx context.Context, request *client.DeleteMarketSkillRequest) (*client.DeleteMarketSkillResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	return sdkClient.DeleteMarketSkillWithOptions(request, cw.getRuntimeOptions())
}

// CreateTag wraps the SDK client method
func (cw *clientWrapper) CreateTag(ctx context.Context, request *client.CreateTagRequest) (*client.CreateTagResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	return sdkClient.CreateTagWithContext(ctx, request, cw.getRuntimeOptions())
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

// DeleteMcpImage wraps the SDK client method
func (cw *clientWrapper) DeleteMcpImage(ctx context.Context, request *client.DeleteMcpImageRequest) (*client.DeleteMcpImageResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	return sdkClient.DeleteMcpImageWithContext(ctx, request, cw.getRuntimeOptions())
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

// ModifyMcpApiKeyConfig wraps the SDK client method
func (cw *clientWrapper) ModifyMcpApiKeyConfig(ctx context.Context, request *client.ModifyMcpApiKeyConfigRequest) (*client.ModifyMcpApiKeyConfigResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	return sdkClient.ModifyMcpApiKeyConfigWithContext(ctx, request, cw.getRuntimeOptions())
}

// DescribeInstanceTypes wraps the SDK client method
func (cw *clientWrapper) DescribeInstanceTypes(ctx context.Context, request *client.DescribeInstanceTypesRequest) (*client.DescribeInstanceTypesResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	return sdkClient.DescribeInstanceTypesWithContext(ctx, request, cw.getRuntimeOptions())
}

// DescribeMcpPolicyData wraps the SDK client method
func (cw *clientWrapper) DescribeMcpPolicyData(ctx context.Context, request *client.DescribeMcpPolicyDataRequest) (*client.DescribeMcpPolicyDataResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	return sdkClient.DescribeMcpPolicyDataWithContext(ctx, request, cw.getRuntimeOptions())
}

// SaveMcpPolicyData wraps the SDK client method
func (cw *clientWrapper) SaveMcpPolicyData(ctx context.Context, request *client.SaveMcpPolicyDataRequest) (*client.SaveMcpPolicyDataResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	return sdkClient.SaveMcpPolicyDataWithContext(ctx, request, cw.getRuntimeOptions())
}

// DescribeOfficeSites wraps the SDK client method
func (cw *clientWrapper) DescribeOfficeSites(ctx context.Context, request *client.DescribeOfficeSitesRequest) (*client.DescribeOfficeSitesResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	return sdkClient.DescribeOfficeSitesWithContext(ctx, request, cw.getRuntimeOptions())
}

// CreateMcpPolicyData wraps the SDK client method
func (cw *clientWrapper) CreateMcpPolicyData(ctx context.Context, request *client.CreateModifyMcpPolicyDataRequest) (*client.CreateMcpPolicyDataResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	return sdkClient.CreateMcpPolicyDataWithContext(ctx, request, cw.getRuntimeOptions())
}

// ModifyMcpPolicyData wraps the SDK client method
func (cw *clientWrapper) ModifyMcpPolicyData(ctx context.Context, request *client.CreateModifyMcpPolicyDataRequest) (*client.ModifyMcpPolicyDataResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	return sdkClient.ModifyMcpPolicyDataWithContext(ctx, request, cw.getRuntimeOptions())
}

// DescribeNetworkPackages wraps the SDK client method
func (cw *clientWrapper) DescribeNetworkPackages(ctx context.Context, request *client.DescribeNetworkPackagesRequest) (*client.DescribeNetworkPackagesResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	return sdkClient.DescribeNetworkPackagesWithContext(ctx, request, cw.getRuntimeOptions())
}

// BatchCreateHideResourceGroupsWithMaxSession wraps the SDK client method
func (cw *clientWrapper) BatchCreateHideResourceGroupsWithMaxSession(ctx context.Context, request *client.BatchCreateHideResourceGroupsWithMaxSessionRequest) (*client.BatchCreateHideResourceGroupsWithMaxSessionResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	return sdkClient.BatchCreateHideResourceGroupsWithMaxSessionWithContext(ctx, request, cw.getRuntimeOptions())
}

// DescribeMcpApiKey wraps the SDK client method
func (cw *clientWrapper) DescribeMcpApiKey(ctx context.Context, request *client.DescribeMcpApiKeyRequest) (*client.DescribeMcpApiKeyResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	return sdkClient.DescribeMcpApiKeyWithContext(ctx, request, cw.getRuntimeOptions())
}

// ModifyApiKeyStatus wraps the SDK client method
func (cw *clientWrapper) ModifyApiKeyStatus(ctx context.Context, request *client.ModifyApiKeyStatusRequest) (*client.ModifyApiKeyStatusResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	return sdkClient.ModifyApiKeyStatusWithContext(ctx, request, cw.getRuntimeOptions())
}

// DeleteApiKey wraps the SDK client method
func (cw *clientWrapper) DeleteApiKey(ctx context.Context, request *client.DeleteApiKeyRequest) (*client.DeleteApiKeyResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	return sdkClient.DeleteApiKeyWithContext(ctx, request, cw.getRuntimeOptions())
}

// DescribeApiKeys wraps the SDK client method
func (cw *clientWrapper) DescribeApiKeys(ctx context.Context, request *client.DescribeApiKeysRequest) (*client.DescribeApiKeysResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	return sdkClient.DescribeApiKeysWithContext(ctx, request, cw.getRuntimeOptions())
}

// DescribeKeyContent wraps the SDK client method
func (cw *clientWrapper) DescribeKeyContent(ctx context.Context, request *client.DescribeKeyContentRequest) (*client.DescribeKeyContentResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	return sdkClient.DescribeKeyContentWithContext(ctx, request, cw.getRuntimeOptions())
}

// DescribeWarmUpStatusOpen wraps the SDK client method
func (cw *clientWrapper) DescribeWarmUpStatusOpen(ctx context.Context, request *client.DescribeWarmUpStatusOpenRequest) (*client.DescribeWarmUpStatusOpenResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	return sdkClient.DescribeWarmUpStatusOpenWithOptions(request, cw.getRuntimeOptions())
}

// ShareDockerRepo wraps the SDK client method
func (cw *clientWrapper) ShareDockerRepo(ctx context.Context, request *client.ShareDockerRepoRequest) (*client.ShareDockerRepoResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	return sdkClient.ShareDockerRepoWithContext(ctx, request, cw.getRuntimeOptions())
}

// UnshareDockerRepo wraps the SDK client method
func (cw *clientWrapper) UnshareDockerRepo(ctx context.Context, request *client.UnshareDockerRepoRequest) (*client.UnshareDockerRepoResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	return sdkClient.UnshareDockerRepoWithContext(ctx, request, cw.getRuntimeOptions())
}

// ListSharedDockerRepos wraps the SDK client method
func (cw *clientWrapper) ListSharedDockerRepos(ctx context.Context, request *client.ListSharedDockerReposRequest) (*client.ListSharedDockerReposResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	return sdkClient.ListSharedDockerReposWithContext(ctx, request, cw.getRuntimeOptions())
}
