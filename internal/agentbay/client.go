// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package agentbay

import (
	"context"
	"fmt"
	"net/http"
	"time"

	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
	log "github.com/sirupsen/logrus"

	"github.com/agentbay/agentbay-cli/internal/auth"
	"github.com/agentbay/agentbay-cli/internal/client"
	"github.com/agentbay/agentbay-cli/internal/config"
)

// debugTransport wraps http.RoundTripper for OAuth client configuration.
type debugTransport struct {
	base http.RoundTripper
}

// RoundTrip implements http.RoundTripper interface.
func (dt *debugTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := dt.base.RoundTrip(req)
	if err != nil && log.GetLevel() >= log.DebugLevel {
		log.Debugf("[DEBUG] HTTP request failed: %v", err)
	}
	return resp, err
}

// debugHttpClient implements dara.HttpClient interface with debug logging
type debugHttpClient struct {
	client *http.Client
}

// Call implements dara.HttpClient interface
func (dhc *debugHttpClient) Call(request *http.Request, transport *http.Transport) (*http.Response, error) {
	// Use our debug client to make the request
	return dhc.client.Do(request)
}

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
	// Market Skill & Group
	GetMarketSkillCredential(ctx context.Context, request *client.GetMarketSkillCredentialRequest) (*client.GetMarketSkillCredentialResponse, error)
	CreateMarketSkill(ctx context.Context, request *client.CreateMarketSkillRequest) (*client.CreateMarketSkillResponse, error)
	DescribeMarketSkillDetail(ctx context.Context, request *client.DescribeMarketSkillDetailRequest) (*client.DescribeMarketSkillDetailResponse, error)
	CreateMarketSkillGroup(ctx context.Context, request *client.CreateMarketSkillGroupRequest) (*client.CreateMarketSkillGroupResponse, error)
	ListMarketGroupSkill(ctx context.Context, request *client.ListMarketGroupSkillRequest) (*client.ListMarketGroupSkillResponse, error)
	AddMarketGroupSkill(ctx context.Context, request *client.AddMarketGroupSkillRequest) (*client.AddMarketGroupSkillResponse, error)
	RemoveMarketGroupSkill(ctx context.Context, request *client.RemoveMarketGroupSkillRequest) (*client.RemoveMarketGroupSkillResponse, error)
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

	// Custom HTTP client for XML response caching (fallback parsing)
	// Create a custom transport that wraps the default transport
	baseTransport := http.DefaultTransport
	if baseTransport == nil {
		baseTransport = &http.Transport{}
	}

	debugTransport := &debugTransport{
		base: baseTransport,
	}

	// Create a custom HTTP client with our debug transport
	httpClient := &http.Client{
		Transport: debugTransport,
	}

	// Create a debug HTTP client that implements dara.HttpClient interface
	debugClient := &debugHttpClient{
		client: httpClient,
	}

	// Set the custom HTTP client in OpenAPI config
	openapiConfig.HttpClient = debugClient

	sdkClient, err := client.NewClient(openapiConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create API client: %w", err)
	}

	return sdkClient, nil
}

// getRuntimeOptions returns runtime options with debug enabled in verbose mode
func (cw *clientWrapper) getRuntimeOptions() *dara.RuntimeOptions {
	runtimeOptions := &dara.RuntimeOptions{}

	// Note: The detailed HTTP logging is now handled by the custom HTTP client
	// set in the OpenAPI config during client creation

	return runtimeOptions
}

// GetDockerFileStoreCredential wraps the SDK client method
func (cw *clientWrapper) GetDockerFileStoreCredential(ctx context.Context, request *client.GetDockerFileStoreCredentialRequest) (*client.GetDockerFileStoreCredentialResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	runtimeOptions := cw.getRuntimeOptions()
	if log.GetLevel() >= log.DebugLevel {
		log.Debugf("[DEBUG] Making GetDockerFileStoreCredential request...")
	}
	resp, err := sdkClient.GetDockerFileStoreCredentialWithOptions(request, runtimeOptions)
	if log.GetLevel() >= log.DebugLevel {
		if err != nil {
			log.Debugf("[DEBUG] GetDockerFileStoreCredential API error: %v", err)
		} else {
			log.Debugf("[DEBUG] GetDockerFileStoreCredential request completed successfully")
		}
	}
	return resp, err
}

// GetMarketSkillCredential wraps the SDK client method
func (cw *clientWrapper) GetMarketSkillCredential(ctx context.Context, request *client.GetMarketSkillCredentialRequest) (*client.GetMarketSkillCredentialResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	runtimeOptions := cw.getRuntimeOptions()
	return sdkClient.GetMarketSkillCredentialWithOptions(request, runtimeOptions)
}

// CreateMarketSkill wraps the SDK client method
func (cw *clientWrapper) CreateMarketSkill(ctx context.Context, request *client.CreateMarketSkillRequest) (*client.CreateMarketSkillResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	runtimeOptions := cw.getRuntimeOptions()
	return sdkClient.CreateMarketSkillWithOptions(request, runtimeOptions)
}

// DescribeMarketSkillDetail wraps the SDK client method
func (cw *clientWrapper) DescribeMarketSkillDetail(ctx context.Context, request *client.DescribeMarketSkillDetailRequest) (*client.DescribeMarketSkillDetailResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	runtimeOptions := cw.getRuntimeOptions()
	return sdkClient.DescribeMarketSkillDetailWithOptions(request, runtimeOptions)
}

// CreateMarketSkillGroup wraps the SDK client method
func (cw *clientWrapper) CreateMarketSkillGroup(ctx context.Context, request *client.CreateMarketSkillGroupRequest) (*client.CreateMarketSkillGroupResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	runtimeOptions := cw.getRuntimeOptions()
	return sdkClient.CreateMarketSkillGroupWithOptions(request, runtimeOptions)
}

// ListMarketGroupSkill wraps the SDK client method
func (cw *clientWrapper) ListMarketGroupSkill(ctx context.Context, request *client.ListMarketGroupSkillRequest) (*client.ListMarketGroupSkillResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	runtimeOptions := cw.getRuntimeOptions()
	return sdkClient.ListMarketGroupSkillWithOptions(request, runtimeOptions)
}

// AddMarketGroupSkill wraps the SDK client method
func (cw *clientWrapper) AddMarketGroupSkill(ctx context.Context, request *client.AddMarketGroupSkillRequest) (*client.AddMarketGroupSkillResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	runtimeOptions := cw.getRuntimeOptions()
	return sdkClient.AddMarketGroupSkillWithOptions(request, runtimeOptions)
}

// RemoveMarketGroupSkill wraps the SDK client method
func (cw *clientWrapper) RemoveMarketGroupSkill(ctx context.Context, request *client.RemoveMarketGroupSkillRequest) (*client.RemoveMarketGroupSkillResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	runtimeOptions := cw.getRuntimeOptions()
	return sdkClient.RemoveMarketGroupSkillWithOptions(request, runtimeOptions)
}

// CreateDockerImageTask wraps the SDK client method
func (cw *clientWrapper) CreateDockerImageTask(ctx context.Context, request *client.CreateDockerImageTaskRequest) (*client.CreateDockerImageTaskResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	runtimeOptions := cw.getRuntimeOptions()
	if log.GetLevel() >= log.DebugLevel {
		log.Debugf("[DEBUG] Making CreateDockerImageTask request...")
	}
	resp, err := sdkClient.CreateDockerImageTaskWithContext(ctx, request, runtimeOptions)
	if log.GetLevel() >= log.DebugLevel {
		if err != nil {
			log.Debugf("[DEBUG] CreateDockerImageTask API error: %v", err)
		} else {
			log.Debugf("[DEBUG] CreateDockerImageTask request completed successfully")
		}
	}
	return resp, err
}

// GetDockerImageTask wraps the SDK client method
func (cw *clientWrapper) GetDockerImageTask(ctx context.Context, request *client.GetDockerImageTaskRequest) (*client.GetDockerImageTaskResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	runtimeOptions := cw.getRuntimeOptions()
	if log.GetLevel() >= log.DebugLevel {
		log.Debugf("[DEBUG] Making GetDockerImageTask request...")
	}
	resp, err := sdkClient.GetDockerImageTaskWithContext(ctx, request, runtimeOptions)
	if log.GetLevel() >= log.DebugLevel {
		if err != nil {
			log.Debugf("[DEBUG] GetDockerImageTask API error: %v", err)
		} else {
			log.Debugf("[DEBUG] GetDockerImageTask request completed successfully")
		}
	}
	return resp, err
}

// ListMcpImages wraps the SDK client method
func (cw *clientWrapper) ListMcpImages(ctx context.Context, request *client.ListMcpImagesRequest) (*client.ListMcpImagesResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	runtimeOptions := cw.getRuntimeOptions()
	if log.GetLevel() >= log.DebugLevel {
		log.Debugf("[DEBUG] Making ListMcpImages request...")
	}
	resp, err := sdkClient.ListMcpImagesWithContext(ctx, request, runtimeOptions)
	if log.GetLevel() >= log.DebugLevel {
		if err != nil {
			log.Debugf("[DEBUG] ListMcpImages API error: %v", err)
		} else {
			log.Debugf("[DEBUG] ListMcpImages request completed successfully")
		}
	}
	return resp, err
}

// CreateResourceGroup wraps the SDK client method
func (cw *clientWrapper) CreateResourceGroup(ctx context.Context, request *client.CreateResourceGroupRequest) (*client.CreateResourceGroupResponse, error) {
	log.Debugf("[DEBUG] ClientWrapper: CreateResourceGroup called")
	if log.GetLevel() >= log.DebugLevel {
		log.Debugf("[DEBUG] CreateResourceGroup request parameters:")
		if request.ImageId != nil {
			log.Debugf("[DEBUG]   - ImageId: %s", *request.ImageId)
		}
		if request.Cpu != nil {
			log.Debugf("[DEBUG]   - Cpu: %d", *request.Cpu)
		}
		if request.Memory != nil {
			log.Debugf("[DEBUG]   - Memory: %d", *request.Memory)
		}
		if request.BizRegionId != nil {
			log.Debugf("[DEBUG]   - BizRegionId: %s", *request.BizRegionId)
		}
		if request.RegionId != nil {
			log.Debugf("[DEBUG]   - RegionId: %s", *request.RegionId)
		}
	}
	sdkClient, err := cw.getClient()
	if err != nil {
		log.Debugf("[DEBUG] ClientWrapper: Failed to get SDK client: %v", err)
		return nil, err
	}
	runtimeOptions := cw.getRuntimeOptions()
	if log.GetLevel() >= log.DebugLevel {
		log.Debugf("[DEBUG] Making CreateResourceGroup request...")
	}
	resp, err := sdkClient.CreateResourceGroupWithContext(ctx, request, runtimeOptions)
	if err != nil {
		log.Debugf("[DEBUG] ClientWrapper: CreateResourceGroup SDK call failed: %v", err)
		return nil, err
	}
	log.Debugf("[DEBUG] ClientWrapper: CreateResourceGroup completed successfully")
	return resp, nil
}

// DeleteResourceGroup wraps the SDK client method
func (cw *clientWrapper) DeleteResourceGroup(ctx context.Context, request *client.DeleteResourceGroupRequest) (*client.DeleteResourceGroupResponse, error) {
	log.Debugf("[DEBUG] ClientWrapper: DeleteResourceGroup called")
	if log.GetLevel() >= log.DebugLevel {
		log.Debugf("[DEBUG] ClientWrapper: Request ImageId = %v", request.GetImageId())
		if request.GetImageId() != nil {
			log.Debugf("[DEBUG] ClientWrapper: ImageId value = %s", *request.GetImageId())
		}
	}
	sdkClient, err := cw.getClient()
	if err != nil {
		log.Debugf("[DEBUG] ClientWrapper: Failed to get SDK client: %v", err)
		return nil, err
	}
	runtimeOptions := cw.getRuntimeOptions()
	if log.GetLevel() >= log.DebugLevel {
		log.Debugf("[DEBUG] Making DeleteResourceGroup request...")
	}
	resp, err := sdkClient.DeleteResourceGroupWithContext(ctx, request, runtimeOptions)
	if err != nil {
		log.Debugf("[DEBUG] ClientWrapper: DeleteResourceGroup SDK call failed: %v", err)
		return nil, err
	}
	log.Debugf("[DEBUG] ClientWrapper: DeleteResourceGroup completed successfully")
	return resp, nil
}

// GetMcpImageInfo wraps the SDK client method
func (cw *clientWrapper) GetMcpImageInfo(ctx context.Context, request *client.GetMcpImageInfoRequest) (*client.GetMcpImageInfoResponse, error) {
	log.Debugf("[DEBUG] ClientWrapper: GetMcpImageInfo called")
	if log.GetLevel() >= log.DebugLevel {
		log.Debugf("[DEBUG] ClientWrapper: Request ImageId = %v", request.GetImageId())
		if request.GetImageId() != nil {
			log.Debugf("[DEBUG] ClientWrapper: ImageId value = %s", *request.GetImageId())
		}
	}
	sdkClient, err := cw.getClient()
	if err != nil {
		log.Debugf("[DEBUG] ClientWrapper: Failed to get SDK client: %v", err)
		return nil, err
	}
	runtimeOptions := cw.getRuntimeOptions()
	if log.GetLevel() >= log.DebugLevel {
		log.Debugf("[DEBUG] Making GetMcpImageInfo request...")
	}
	resp, err := sdkClient.GetMcpImageInfoWithContext(ctx, request, runtimeOptions)
	if err != nil {
		log.Debugf("[DEBUG] ClientWrapper: GetMcpImageInfo SDK call failed: %v", err)
		return nil, err
	}
	if log.GetLevel() >= log.DebugLevel {
		log.Debugf("[DEBUG] ClientWrapper: GetMcpImageInfo completed successfully")
	}
	return resp, nil
}

// GetDockerfileTemplate wraps the SDK client method
func (cw *clientWrapper) GetDockerfileTemplate(ctx context.Context, request *client.GetDockerfileTemplateRequest) (*client.GetDockerfileTemplateResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}
	runtimeOptions := cw.getRuntimeOptions()
	if log.GetLevel() >= log.DebugLevel {
		log.Debugf("[DEBUG] Making GetDockerfileTemplate request...")
	}
	resp, err := sdkClient.GetDockerfileTemplateWithContext(ctx, request, runtimeOptions)
	if log.GetLevel() >= log.DebugLevel {
		if err != nil {
			log.Debugf("[DEBUG] GetDockerfileTemplate API error: %v", err)
		} else {
			log.Debugf("[DEBUG] ClientWrapper: GetDockerfileTemplate completed successfully")
		}
	}
	return resp, err
}
