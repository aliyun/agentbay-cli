// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package agentbay

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
	log "github.com/sirupsen/logrus"

	"github.com/agentbay/agentbay-cli/internal/auth"
	"github.com/agentbay/agentbay-cli/internal/client"
	"github.com/agentbay/agentbay-cli/internal/config"
)

// xmlResponseCache stores the last XML response for fallback parsing
var xmlResponseCache string

// formatXMLForDisplay formats XML string for better readability in logs
func formatXMLForDisplay(xmlStr string) string {
	// Simple XML formatting - add newlines after major tags
	formatted := []byte(xmlStr)
	formatted = bytes.ReplaceAll(formatted, []byte("><"), []byte(">\n<"))

	// Add indentation for nested tags
	lines := bytes.Split(formatted, []byte("\n"))
	var result []string
	indent := 0

	for _, line := range lines {
		lineStr := string(bytes.TrimSpace(line))
		if lineStr == "" {
			continue
		}

		// Decrease indent for closing tags
		if bytes.HasPrefix(bytes.TrimSpace(line), []byte("</")) {
			indent--
		}

		// Add indentation
		indentStr := ""
		for i := 0; i < indent; i++ {
			indentStr += "  "
		}
		result = append(result, indentStr+lineStr)

		// Increase indent for opening tags (but not self-closing or closing tags)
		if bytes.HasPrefix(bytes.TrimSpace(line), []byte("<")) &&
			!bytes.HasPrefix(bytes.TrimSpace(line), []byte("</")) &&
			!bytes.HasSuffix(bytes.TrimSpace(line), []byte("/>")) &&
			!bytes.HasPrefix(bytes.TrimSpace(line), []byte("<?")) {
			indent++
		}
	}

	return strings.Join(result, "\n")
}

// debugTransport wraps http.RoundTripper to log request/response details in verbose mode
type debugTransport struct {
	base http.RoundTripper
}

// RoundTrip implements http.RoundTripper interface
func (dt *debugTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Make the actual request
	resp, err := dt.base.RoundTrip(req)
	if err != nil {
		if log.GetLevel() >= log.DebugLevel {
			log.Debugf("[DEBUG] HTTP request failed: %v", err)
		}
		return resp, err
	}

	// Cache XML responses for fallback parsing (always needed)
	if resp.Body != nil {
		body, err := io.ReadAll(resp.Body)
		if err == nil {
			// Cache XML responses for fallback parsing
			if bytes.Contains(body, []byte("<?xml")) && (bytes.Contains(body, []byte("GetDockerFileStoreCredentialResponse")) || bytes.Contains(body, []byte("CreateDockerImageTaskResponse")) || bytes.Contains(body, []byte("GetDockerImageTaskResponse")) || bytes.Contains(body, []byte("ListMcpImagesResponse")) || bytes.Contains(body, []byte("GetMcpImageInfoResponse")) || bytes.Contains(body, []byte("CreateResourceGroupResponse")) || bytes.Contains(body, []byte("DeleteResourceGroupResponse")) || bytes.Contains(body, []byte("GetDockerfileTemplateResponse"))) {
				xmlResponseCache = string(body)
				log.Debugf("[DEBUG] Cached XML response for fallback parsing")
			}

			// Restore the body for normal processing
			resp.Body = io.NopCloser(bytes.NewReader(body))
		}
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
	log.Debugf("[DEBUG] getClient: Creating new SDK client...")

	// Refresh token if needed (checks expiry and refreshes automatically)
	log.Debugf("[DEBUG] getClient: Checking if token refresh is needed...")

	// Create an adapter to bridge config.Config to auth.TokenConfig
	tokenCfgAdapter := auth.NewConfigAdapter(
		func() (string, string, time.Time, error) {
			return cw.config.GetTokens()
		},
		cw.config.RefreshTokens,
		cw.config.IsTokenExpired,
		cw.config.ClearTokens,
	)

	err := auth.RefreshTokenIfNeeded(tokenCfgAdapter, auth.DefaultClientID)
	if err != nil {
		log.Debugf("[DEBUG] getClient: Token refresh check failed: %v", err)
		return nil, fmt.Errorf("failed to ensure valid token: %w", err)
	}

	// Get authentication token (now guaranteed to be valid or refreshed)
	log.Debugf("[DEBUG] getClient: Getting authentication token...")
	token, err := cw.config.GetToken()
	if err != nil {
		log.Debugf("[DEBUG] getClient: Failed to get authentication token: %v", err)
		return nil, fmt.Errorf("failed to get authentication token: %w", err)
	}
	log.Debugf("[DEBUG] getClient: Authentication token obtained")
	log.Debugf("[DEBUG] getClient: Access token length: %d", len(token.AccessToken))
	if len(token.AccessToken) > 20 {
		log.Debugf("[DEBUG] getClient: Access token preview: %s...", token.AccessToken[:20])
	} else {
		log.Debugf("[DEBUG] getClient: Full access token (short): '%s'", token.AccessToken)
	}

	// Create OpenAPI config
	// For Alibaba Cloud SDK, we should pass only the hostname, not the full URL
	endpoint := cw.apiConfig.Endpoint // Use raw endpoint without https:// prefix
	log.Debugf("[DEBUG] getClient: Creating OpenAPI config with endpoint: %s", endpoint)
	log.Debugf("[DEBUG] getClient: Timeout: %d ms", cw.apiConfig.TimeoutMs)

	openapiConfig := &openapiutil.Config{
		// Use BearerToken for OAuth authentication (backend now supports BearerToken)
		BearerToken:    dara.String(token.AccessToken),
		Endpoint:       dara.String(endpoint),
		ReadTimeout:    dara.Int(cw.apiConfig.TimeoutMs),
		ConnectTimeout: dara.Int(cw.apiConfig.TimeoutMs),
		UserAgent:      dara.String("AgentBay-CLI/1.0"),
	}

	// Set custom HTTP client for XML response caching (always needed for fallback parsing)
	log.Debugf("[DEBUG] getClient: Setting up HTTP transport for XML response caching")

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

	// Create the client using the generated SDK client constructor
	log.Debugf("[DEBUG] getClient: Calling client.NewClient...")
	sdkClient, err := client.NewClient(openapiConfig)
	if err != nil {
		log.Debugf("[DEBUG] getClient: client.NewClient failed: %v", err)
		return nil, fmt.Errorf("failed to create API client: %w", err)
	}
	log.Debugf("[DEBUG] getClient: SDK client created successfully")

	// Don't cache the client to ensure token refresh check on each API call
	return sdkClient, nil
}

// getRuntimeOptions returns runtime options with debug enabled in verbose mode
func (cw *clientWrapper) getRuntimeOptions() *dara.RuntimeOptions {
	runtimeOptions := &dara.RuntimeOptions{}

	// Note: The detailed HTTP logging is now handled by the custom HTTP client
	// set in the OpenAPI config during client creation

	return runtimeOptions
}

// XMLGetDockerFileStoreCredentialResponse represents the XML response structure
type XMLGetDockerFileStoreCredentialResponse struct {
	XMLName        xml.Name `xml:"GetDockerFileStoreCredentialResponse"`
	RequestId      string   `xml:"RequestId"`
	HttpStatusCode int      `xml:"HttpStatusCode"`
	Data           struct {
		TaskId string `xml:"TaskId"`
		OssUrl string `xml:"OssUrl"`
	} `xml:"Data"`
	Code    string `xml:"Code"`
	Success bool   `xml:"Success"`
}

// parseXMLResponse parses XML response and converts it to SDK response format
func (cw *clientWrapper) parseXMLResponse(xmlData []byte) (*client.GetDockerFileStoreCredentialResponse, error) {
	log.Debugf("[DEBUG] Parsing XML data: %s", string(xmlData))

	var xmlResp XMLGetDockerFileStoreCredentialResponse
	if err := xml.Unmarshal(xmlData, &xmlResp); err != nil {
		log.Debugf("[DEBUG] XML unmarshal failed: %v", err)
		return nil, fmt.Errorf("failed to parse XML response: %w", err)
	}

	log.Debugf("[DEBUG] Parsed XML data:")
	log.Debugf("[DEBUG] - RequestId: %s", xmlResp.RequestId)
	log.Debugf("[DEBUG] - HttpStatusCode: %d", xmlResp.HttpStatusCode)
	log.Debugf("[DEBUG] - Code: %s", xmlResp.Code)
	log.Debugf("[DEBUG] - Success: %t", xmlResp.Success)
	log.Debugf("[DEBUG] - TaskId: %s", xmlResp.Data.TaskId)
	log.Debugf("[DEBUG] - OssUrl: %s", xmlResp.Data.OssUrl)

	// Convert to SDK response format
	response := &client.GetDockerFileStoreCredentialResponse{
		Headers:    make(map[string]*string),
		StatusCode: dara.Int32(int32(xmlResp.HttpStatusCode)),
		Body: &client.GetDockerFileStoreCredentialResponseBody{
			RequestId:      dara.String(xmlResp.RequestId),
			HttpStatusCode: dara.Int32(int32(xmlResp.HttpStatusCode)),
			Code:           dara.String(xmlResp.Code),
			Success:        dara.Bool(xmlResp.Success),
			Data: &client.GetDockerFileStoreCredentialResponseBodyData{
				TaskId: dara.String(xmlResp.Data.TaskId),
				OssUrl: dara.String(xmlResp.Data.OssUrl),
			},
		},
	}

	log.Debugf("[DEBUG] Created SDK response with OssUrl: %s", xmlResp.Data.OssUrl)
	return response, nil
}

// XMLCreateDockerImageTaskResponse represents the XML response structure for CreateDockerImageTask
type XMLCreateDockerImageTaskResponse struct {
	XMLName        xml.Name `xml:"CreateDockerImageTaskResponse"`
	RequestId      string   `xml:"RequestId"`
	HttpStatusCode int      `xml:"HttpStatusCode"`
	Data           struct {
		TaskId string `xml:"TaskId"`
	} `xml:"Data"`
	Code    string `xml:"Code"`
	Success bool   `xml:"Success"`
}

// parseCreateDockerImageTaskXMLResponse parses XML response and converts it to SDK response format
func (cw *clientWrapper) parseCreateDockerImageTaskXMLResponse(xmlData []byte) (*client.CreateDockerImageTaskResponse, error) {
	log.Debugf("[DEBUG] Parsing CreateDockerImageTask XML data: %s", string(xmlData))

	var xmlResp XMLCreateDockerImageTaskResponse
	if err := xml.Unmarshal(xmlData, &xmlResp); err != nil {
		log.Debugf("[DEBUG] CreateDockerImageTask XML unmarshal failed: %v", err)
		return nil, fmt.Errorf("failed to parse CreateDockerImageTask XML response: %w", err)
	}

	log.Debugf("[DEBUG] Parsed CreateDockerImageTask XML data:")
	log.Debugf("[DEBUG] - RequestId: %s", xmlResp.RequestId)
	log.Debugf("[DEBUG] - HttpStatusCode: %d", xmlResp.HttpStatusCode)
	log.Debugf("[DEBUG] - Code: %s", xmlResp.Code)
	log.Debugf("[DEBUG] - Success: %t", xmlResp.Success)
	log.Debugf("[DEBUG] - TaskId: %s", xmlResp.Data.TaskId)

	// Convert to SDK response format
	response := &client.CreateDockerImageTaskResponse{
		Headers:    make(map[string]*string),
		StatusCode: dara.Int32(int32(xmlResp.HttpStatusCode)),
		Body: &client.CreateDockerImageTaskResponseBody{
			RequestId:      dara.String(xmlResp.RequestId),
			HttpStatusCode: dara.Int32(int32(xmlResp.HttpStatusCode)),
			Code:           dara.String(xmlResp.Code),
			Success:        dara.Bool(xmlResp.Success),
			Data: &client.CreateDockerImageTaskResponseBodyData{
				TaskId: dara.String(xmlResp.Data.TaskId),
			},
		},
	}

	log.Debugf("[DEBUG] Created CreateDockerImageTask SDK response with TaskId: %s", xmlResp.Data.TaskId)
	return response, nil
}

// XMLGetDockerImageTaskResponse represents the XML response structure for GetDockerImageTask
type XMLGetDockerImageTaskResponse struct {
	XMLName        xml.Name `xml:"GetDockerImageTaskResponse"`
	RequestId      string   `xml:"RequestId"`
	HttpStatusCode int      `xml:"HttpStatusCode"`
	Data           struct {
		Status  string `xml:"Status"`
		ImageId string `xml:"ImageId"`
		TaskMsg string `xml:"TaskMsg"`
	} `xml:"Data"`
	Code    string `xml:"Code"`
	Success bool   `xml:"Success"`
}

// XMLListMcpImagesResponse represents the XML response structure for ListMcpImages
type XMLListMcpImagesResponse struct {
	XMLName        xml.Name `xml:"ListMcpImagesResponse"`
	RequestId      string   `xml:"RequestId"`
	HttpStatusCode int      `xml:"HttpStatusCode"`
	Data           struct {
		Images []struct {
			ImageId             string `xml:"ImageId"`
			ImageName           string `xml:"ImageName"`
			ImageBuildType      string `xml:"ImageBuildType"`
			ImageIntro          string `xml:"ImageIntro"`
			ImageApplyScene     string `xml:"ImageApplyScene"`
			ImageResourceStatus string `xml:"ImageResourceStatus"`
			ImageInfo           struct {
				OsName         string `xml:"OsName"`
				OsVersion      string `xml:"OsVersion"`
				PlatformName   string `xml:"PlatformName"`
				Status         string `xml:"Status"`
				DataDiskSize   int32  `xml:"DataDiskSize"`
				SystemDiskSize int32  `xml:"SystemDiskSize"`
				FotaVersion    string `xml:"FotaVersion"`
				UpdateTime     string `xml:"UpdateTime"`
			} `xml:"ImageInfo"`
			ToolInfo []struct {
				McpServerId   string `xml:"McpServerId"`
				McpServerName string `xml:"McpServerName"`
				ToolList      []struct {
					Tool        string `xml:"Tool"`
					Description string `xml:"Description"`
				} `xml:"ToolList"`
			} `xml:"ToolInfo"`
		} `xml:"data"`
	} `xml:"Data"`
	Code       string `xml:"Code"`
	Success    bool   `xml:"Success"`
	TotalCount int32  `xml:"TotalCount"`
	PageSize   int32  `xml:"PageSize"`
	PageStart  int32  `xml:"PageStart"`
	NextToken  string `xml:"NextToken"`
}

// XMLCreateResourceGroupResponse represents the XML response structure for CreateResourceGroup
type XMLCreateResourceGroupResponse struct {
	XMLName        xml.Name `xml:"CreateResourceGroupResponse"`
	RequestId      string   `xml:"RequestId"`
	HttpStatusCode int      `xml:"HttpStatusCode"`
	Code           string   `xml:"Code"`
	Success        bool     `xml:"Success"`
	Message        string   `xml:"Message"`
}

// XMLDeleteResourceGroupResponse represents the XML response structure for DeleteResourceGroup
type XMLDeleteResourceGroupResponse struct {
	XMLName        xml.Name `xml:"DeleteResourceGroupResponse"`
	RequestId      string   `xml:"RequestId"`
	HttpStatusCode int      `xml:"HttpStatusCode"`
	Code           string   `xml:"Code"`
	Success        bool     `xml:"Success"`
	Message        string   `xml:"Message"`
}

// XMLGetMcpImageInfoResponse represents the XML response structure for GetMcpImageInfo
type XMLGetMcpImageInfoResponse struct {
	XMLName        xml.Name `xml:"GetMcpImageInfoResponse"`
	RequestId      string   `xml:"RequestId"`
	HttpStatusCode int      `xml:"HttpStatusCode"`
	Code           string   `xml:"Code"`
	Success        bool     `xml:"Success"`
	Message        string   `xml:"Message"`
	Data           struct {
		ImageId             string `xml:"ImageId"`
		ImageName           string `xml:"ImageName"`
		ImageBuildType      string `xml:"ImageBuildType"`
		ImageApplyScene     string `xml:"ImageApplyScene"`
		ImageResourceStatus string `xml:"ImageResourceStatus"`
		ImageInfo           struct {
			OsName         string `xml:"OsName"`
			OsVersion      string `xml:"OsVersion"`
			PlatformName   string `xml:"PlatformName"`
			Status         string `xml:"Status"`
			DataDiskSize   int32  `xml:"DataDiskSize"`
			SystemDiskSize int32  `xml:"SystemDiskSize"`
			UpdateTime     string `xml:"UpdateTime"`
			ImageType      string `xml:"ImageType"`
		} `xml:"ImageInfo"`
		ImageBuildInfo struct {
			TaskId                  string `xml:"TaskId"`
			VersionId               string `xml:"VersionId"`
			ApiKeyId                string `xml:"ApiKeyId"`
			InstanceReady           bool   `xml:"InstanceReady"`
			AndroidMobileGroupId    string `xml:"AndroidMobileGroupId"`
			AndroidMobileInstanceId string `xml:"AndroidMobileInstanceId"`
		} `xml:"ImageBuildInfo"`
	} `xml:"Data"`
}

// parseGetDockerImageTaskXMLResponse parses XML response and converts it to SDK response format
func (cw *clientWrapper) parseGetDockerImageTaskXMLResponse(xmlData []byte) (*client.GetDockerImageTaskResponse, error) {
	log.Debugf("[DEBUG] Parsing GetDockerImageTask XML data: %s", string(xmlData))

	var xmlResp XMLGetDockerImageTaskResponse
	if err := xml.Unmarshal(xmlData, &xmlResp); err != nil {
		log.Debugf("[DEBUG] GetDockerImageTask XML unmarshal failed: %v", err)
		return nil, fmt.Errorf("failed to parse GetDockerImageTask XML response: %w", err)
	}

	log.Debugf("[DEBUG] Parsed GetDockerImageTask XML data:")
	log.Debugf("[DEBUG] - RequestId: %s", xmlResp.RequestId)
	log.Debugf("[DEBUG] - HttpStatusCode: %d", xmlResp.HttpStatusCode)
	log.Debugf("[DEBUG] - Code: %s", xmlResp.Code)
	log.Debugf("[DEBUG] - Success: %t", xmlResp.Success)
	log.Debugf("[DEBUG] - Status: %s", xmlResp.Data.Status)
	log.Debugf("[DEBUG] - ImageId: %s", xmlResp.Data.ImageId)
	log.Debugf("[DEBUG] - TaskMsg: %s", xmlResp.Data.TaskMsg)

	// Convert to SDK response format
	response := &client.GetDockerImageTaskResponse{
		Headers:    make(map[string]*string),
		StatusCode: dara.Int32(int32(xmlResp.HttpStatusCode)),
		Body: &client.GetDockerImageTaskResponseBody{
			RequestId:      dara.String(xmlResp.RequestId),
			HttpStatusCode: dara.Int32(int32(xmlResp.HttpStatusCode)),
			Code:           dara.String(xmlResp.Code),
			Success:        dara.Bool(xmlResp.Success),
			Data: &client.GetDockerImageTaskResponseBodyData{
				Status:  dara.String(xmlResp.Data.Status),
				ImageId: dara.String(xmlResp.Data.ImageId),
				TaskMsg: dara.String(xmlResp.Data.TaskMsg),
			},
		},
	}

	log.Debugf("[DEBUG] Created GetDockerImageTask SDK response with Status: %s", xmlResp.Data.Status)
	return response, nil
}

// parseListMcpImagesXMLResponse parses XML response and converts it to SDK response format
func (cw *clientWrapper) parseListMcpImagesXMLResponse(xmlData []byte) (*client.ListMcpImagesResponse, error) {
	log.Debugf("[DEBUG] Parsing ListMcpImages XML data (size: %d bytes)", len(xmlData))

	var xmlResp XMLListMcpImagesResponse
	if err := xml.Unmarshal(xmlData, &xmlResp); err != nil {
		log.Debugf("[DEBUG] ListMcpImages XML unmarshal failed: %v", err)
		return nil, fmt.Errorf("failed to parse ListMcpImages XML response: %w", err)
	}

	log.Debugf("[DEBUG] Parsed ListMcpImages XML data:")
	log.Debugf("[DEBUG] - RequestId: %s", xmlResp.RequestId)
	log.Debugf("[DEBUG] - HttpStatusCode: %d", xmlResp.HttpStatusCode)
	log.Debugf("[DEBUG] - Code: %s", xmlResp.Code)
	log.Debugf("[DEBUG] - Success: %t", xmlResp.Success)
	log.Debugf("[DEBUG] - TotalCount: %d", xmlResp.TotalCount)
	log.Debugf("[DEBUG] - Data length: %d", len(xmlResp.Data.Images))

	// Convert XML data to SDK format
	var sdkData []*client.ListMcpImagesResponseBodyData
	for _, xmlImage := range xmlResp.Data.Images {
		// Convert ImageInfo
		var imageInfo *client.ListMcpImagesResponseBodyDataImageInfo
		if xmlImage.ImageInfo.OsName != "" || xmlImage.ImageInfo.OsVersion != "" {
			imageInfo = &client.ListMcpImagesResponseBodyDataImageInfo{
				OsName:         dara.String(xmlImage.ImageInfo.OsName),
				OsVersion:      dara.String(xmlImage.ImageInfo.OsVersion),
				PlatformName:   dara.String(xmlImage.ImageInfo.PlatformName),
				Status:         dara.String(xmlImage.ImageInfo.Status),
				DataDiskSize:   dara.Int32(xmlImage.ImageInfo.DataDiskSize),
				SystemDiskSize: dara.Int32(xmlImage.ImageInfo.SystemDiskSize),
				FotaVersion:    dara.String(xmlImage.ImageInfo.FotaVersion),
				UpdateTime:     dara.String(xmlImage.ImageInfo.UpdateTime),
			}
		}

		// Convert ToolInfo
		var toolInfo []*client.ListMcpImagesResponseBodyDataToolInfo
		for _, xmlTool := range xmlImage.ToolInfo {
			var toolList []*client.ListMcpImagesResponseBodyDataToolInfoToolList
			for _, xmlToolItem := range xmlTool.ToolList {
				toolList = append(toolList, &client.ListMcpImagesResponseBodyDataToolInfoToolList{
					Tool:        dara.String(xmlToolItem.Tool),
					Description: dara.String(xmlToolItem.Description),
				})
			}

			toolInfo = append(toolInfo, &client.ListMcpImagesResponseBodyDataToolInfo{
				McpServerId:   dara.String(xmlTool.McpServerId),
				McpServerName: dara.String(xmlTool.McpServerName),
				ToolList:      toolList,
			})
		}

		sdkImage := &client.ListMcpImagesResponseBodyData{
			ImageId:             dara.String(xmlImage.ImageId),
			ImageName:           dara.String(xmlImage.ImageName),
			ImageBuildType:      dara.String(xmlImage.ImageBuildType),
			ImageIntro:          dara.String(xmlImage.ImageIntro),
			ImageApplyScene:     dara.String(xmlImage.ImageApplyScene),
			ImageResourceStatus: dara.String(xmlImage.ImageResourceStatus),
			ImageInfo:           imageInfo,
			ToolInfo:            toolInfo,
		}
		sdkData = append(sdkData, sdkImage)
	}

	// Convert to SDK response format
	response := &client.ListMcpImagesResponse{
		Headers:    make(map[string]*string),
		StatusCode: dara.Int32(int32(xmlResp.HttpStatusCode)),
		Body: &client.ListMcpImagesResponseBody{
			RequestId:      dara.String(xmlResp.RequestId),
			HttpStatusCode: dara.Int32(int32(xmlResp.HttpStatusCode)),
			Code:           dara.String(xmlResp.Code),
			Success:        dara.Bool(xmlResp.Success),
			Data:           sdkData,
			TotalCount:     dara.Int32(xmlResp.TotalCount),
			PageSize:       dara.Int32(xmlResp.PageSize),
			PageStart:      dara.Int32(xmlResp.PageStart),
			NextToken:      dara.String(xmlResp.NextToken),
		},
	}

	log.Debugf("[DEBUG] Created ListMcpImages SDK response with %d images", len(sdkData))
	return response, nil
}

// parseCreateResourceGroupXMLResponse parses XML response and converts it to SDK response format
func (cw *clientWrapper) parseCreateResourceGroupXMLResponse(xmlData []byte) (*client.CreateResourceGroupResponse, error) {
	log.Debugf("[DEBUG] Parsing CreateResourceGroup XML data: %s", string(xmlData))

	var xmlResp XMLCreateResourceGroupResponse
	if err := xml.Unmarshal(xmlData, &xmlResp); err != nil {
		log.Debugf("[DEBUG] CreateResourceGroup XML unmarshal failed: %v", err)
		return nil, fmt.Errorf("failed to parse CreateResourceGroup XML response: %w", err)
	}

	log.Debugf("[DEBUG] Parsed CreateResourceGroup XML data:")
	log.Debugf("[DEBUG] - RequestId: %s", xmlResp.RequestId)
	log.Debugf("[DEBUG] - HttpStatusCode: %d", xmlResp.HttpStatusCode)
	log.Debugf("[DEBUG] - Code: %s", xmlResp.Code)
	log.Debugf("[DEBUG] - Success: %t", xmlResp.Success)
	log.Debugf("[DEBUG] - Message: %s", xmlResp.Message)

	// Convert to SDK response format
	response := &client.CreateResourceGroupResponse{
		Headers:    make(map[string]*string),
		StatusCode: dara.Int32(int32(xmlResp.HttpStatusCode)),
		Body: &client.CreateResourceGroupResponseBody{
			RequestId:      dara.String(xmlResp.RequestId),
			HttpStatusCode: dara.Int32(int32(xmlResp.HttpStatusCode)),
			Code:           dara.String(xmlResp.Code),
			Success:        dara.Bool(xmlResp.Success),
			Message:        dara.String(xmlResp.Message),
		},
	}

	log.Debugf("[DEBUG] Created CreateResourceGroup SDK response")
	return response, nil
}

// parseDeleteResourceGroupXMLResponse parses XML response and converts it to SDK response format
func (cw *clientWrapper) parseDeleteResourceGroupXMLResponse(xmlData []byte) (*client.DeleteResourceGroupResponse, error) {
	log.Debugf("[DEBUG] Parsing DeleteResourceGroup XML data: %s", string(xmlData))

	var xmlResp XMLDeleteResourceGroupResponse
	if err := xml.Unmarshal(xmlData, &xmlResp); err != nil {
		log.Debugf("[DEBUG] DeleteResourceGroup XML unmarshal failed: %v", err)
		return nil, fmt.Errorf("failed to parse DeleteResourceGroup XML response: %w", err)
	}

	log.Debugf("[DEBUG] Parsed DeleteResourceGroup XML data:")
	log.Debugf("[DEBUG] - RequestId: %s", xmlResp.RequestId)
	log.Debugf("[DEBUG] - HttpStatusCode: %d", xmlResp.HttpStatusCode)
	log.Debugf("[DEBUG] - Code: %s", xmlResp.Code)
	log.Debugf("[DEBUG] - Success: %t", xmlResp.Success)
	log.Debugf("[DEBUG] - Message: %s", xmlResp.Message)

	// Convert to SDK response format
	response := &client.DeleteResourceGroupResponse{
		Headers:    make(map[string]*string),
		StatusCode: dara.Int32(int32(xmlResp.HttpStatusCode)),
		Body: &client.DeleteResourceGroupResponseBody{
			RequestId:      dara.String(xmlResp.RequestId),
			HttpStatusCode: dara.Int32(int32(xmlResp.HttpStatusCode)),
			Code:           dara.String(xmlResp.Code),
			Success:        dara.Bool(xmlResp.Success),
			Message:        dara.String(xmlResp.Message),
		},
	}

	log.Debugf("[DEBUG] Created DeleteResourceGroup SDK response")
	return response, nil
}

// GetDockerFileStoreCredential wraps the SDK client method
func (cw *clientWrapper) GetDockerFileStoreCredential(ctx context.Context, request *client.GetDockerFileStoreCredentialRequest) (*client.GetDockerFileStoreCredentialResponse, error) {
	log.Debugf("[DEBUG] ClientWrapper: Getting SDK client...")
	sdkClient, err := cw.getClient()
	if err != nil {
		log.Debugf("[DEBUG] ClientWrapper: Failed to get SDK client: %v", err)
		return nil, err
	}
	log.Debugf("[DEBUG] ClientWrapper: SDK client obtained, making API call...")

	// Get runtime options with debug enabled if verbose
	runtimeOptions := cw.getRuntimeOptions()

	// Log basic request information in verbose mode
	if log.GetLevel() >= log.DebugLevel {
		log.Debugf("[DEBUG] Making GetDockerFileStoreCredential request...")
	}

	resp, err := sdkClient.GetDockerFileStoreCredentialWithOptions(request, runtimeOptions)

	// Log detailed response information in verbose mode
	if log.GetLevel() >= log.DebugLevel {
		if err != nil {
			// Check if this is a known XML parsing error
			errStr := err.Error()
			if bytes.Contains([]byte(errStr), []byte("readObjectStart: expect { or n, but found")) {
				log.Debugf("[DEBUG] HTTP Response: XML format detected, will use custom parser")
			} else {
				log.Debugf("[DEBUG] HTTP Response Error: %v", err)
				// Try to extract more details from the error
				log.Debugf("[DEBUG] Error type: %T", err)
				log.Debugf("[DEBUG] Error string: %s", err.Error())
			}
		} else {
			log.Debugf("[DEBUG] GetDockerFileStoreCredential request completed successfully")
		}
	}

	if err != nil {
		// Check if this is an XML parsing error
		errStr := err.Error()
		if bytes.Contains([]byte(errStr), []byte("readObjectStart: expect { or n, but found")) {
			log.Debugf("[DEBUG] SDK returned XML response, using custom XML parser...")

			// Use cached XML response if available
			if xmlResponseCache != "" {
				log.Debugf("[DEBUG] Parsing cached XML response...")

				// Parse the cached XML directly
				customResponse, parseErr := cw.parseXMLResponse([]byte(xmlResponseCache))
				if parseErr != nil {
					log.Debugf("[DEBUG] Custom XML parsing failed: %v", parseErr)
					return nil, fmt.Errorf("XML parsing failed: %w", parseErr)
				}

				log.Debugf("[DEBUG] XML response parsed successfully")

				return customResponse, nil
			} else {
				log.Debugf("[DEBUG] No cached XML response available")
				return nil, fmt.Errorf("XML parsing failed and no cached response available: %w", err)
			}
		}

		log.Debugf("[DEBUG] ClientWrapper: SDK API call failed: %v", err)
		return nil, err
	}
	log.Debugf("[DEBUG] ClientWrapper: SDK API call completed successfully")

	return resp, nil
}

// CreateDockerImageTask wraps the SDK client method
func (cw *clientWrapper) CreateDockerImageTask(ctx context.Context, request *client.CreateDockerImageTaskRequest) (*client.CreateDockerImageTaskResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}

	// Get runtime options with debug enabled if verbose
	runtimeOptions := cw.getRuntimeOptions()

	// Log basic request information in verbose mode
	if log.GetLevel() >= log.DebugLevel {
		log.Debugf("[DEBUG] Making CreateDockerImageTask request...")
	}

	resp, err := sdkClient.CreateDockerImageTaskWithOptions(request, runtimeOptions)

	// Log detailed response information in verbose mode
	if log.GetLevel() >= log.DebugLevel {
		if err != nil {
			// Check if this is a known XML parsing error
			errStr := err.Error()
			if bytes.Contains([]byte(errStr), []byte("readObjectStart: expect { or n, but found")) {
				log.Debugf("[DEBUG] CreateDockerImageTask HTTP Response: XML format detected, will use custom parser")
			} else {
				log.Debugf("[DEBUG] CreateDockerImageTask HTTP Response Error: %v", err)
				// Try to extract more details from the error
				log.Debugf("[DEBUG] Error type: %T", err)
				log.Debugf("[DEBUG] Error string: %s", err.Error())
			}
		} else {
			log.Debugf("[DEBUG] CreateDockerImageTask request completed successfully")
		}
	}

	if err != nil {
		// Check if this is an XML parsing error
		errStr := err.Error()
		if bytes.Contains([]byte(errStr), []byte("readObjectStart: expect { or n, but found")) {
			log.Debugf("[DEBUG] SDK returned XML response, using custom XML parser...")

			// Use cached XML response if available
			if xmlResponseCache != "" {
				log.Debugf("[DEBUG] Parsing cached XML response...")

				// Parse the cached XML directly
				customResponse, parseErr := cw.parseCreateDockerImageTaskXMLResponse([]byte(xmlResponseCache))
				if parseErr != nil {
					log.Debugf("[DEBUG] Custom XML parsing failed: %v", parseErr)
					return nil, fmt.Errorf("XML parsing failed: %w", parseErr)
				}

				log.Debugf("[DEBUG] XML response parsed successfully")
				return customResponse, nil
			} else {
				log.Debugf("[DEBUG] No cached XML response available")
				return nil, fmt.Errorf("XML parsing failed and no cached response available: %w", err)
			}
		}

		log.Debugf("[DEBUG] ClientWrapper: SDK API call failed: %v", err)
		return nil, err
	}
	log.Debugf("[DEBUG] ClientWrapper: SDK API call completed successfully")
	return resp, nil
}

// GetDockerImageTask wraps the SDK client method
func (cw *clientWrapper) GetDockerImageTask(ctx context.Context, request *client.GetDockerImageTaskRequest) (*client.GetDockerImageTaskResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}

	// Get runtime options with debug enabled if verbose
	runtimeOptions := cw.getRuntimeOptions()

	// Log basic request information in verbose mode
	if log.GetLevel() >= log.DebugLevel {
		log.Debugf("[DEBUG] Making GetDockerImageTask request...")
	}

	resp, err := sdkClient.GetDockerImageTaskWithOptions(request, runtimeOptions)

	// Log detailed response information in verbose mode
	if log.GetLevel() >= log.DebugLevel {
		if err != nil {
			// Check if this is a known XML parsing error
			errStr := err.Error()
			if bytes.Contains([]byte(errStr), []byte("readObjectStart: expect { or n, but found")) {
				log.Debugf("[DEBUG] GetDockerImageTask HTTP Response: XML format detected, will use custom parser")
			} else {
				log.Debugf("[DEBUG] GetDockerImageTask HTTP Response Error: %v", err)
				// Try to extract more details from the error
				log.Debugf("[DEBUG] Error type: %T", err)
				log.Debugf("[DEBUG] Error string: %s", err.Error())
			}
		} else {
			log.Debugf("[DEBUG] GetDockerImageTask request completed successfully")
		}
	}

	if err != nil {
		// Check if this is an XML parsing error
		errStr := err.Error()
		if bytes.Contains([]byte(errStr), []byte("readObjectStart: expect { or n, but found")) {
			log.Debugf("[DEBUG] SDK returned XML response, using custom XML parser...")

			// Use cached XML response if available
			if xmlResponseCache != "" {
				log.Debugf("[DEBUG] Parsing cached XML response...")

				// Parse the cached XML directly
				customResponse, parseErr := cw.parseGetDockerImageTaskXMLResponse([]byte(xmlResponseCache))
				if parseErr != nil {
					log.Debugf("[DEBUG] Custom XML parsing failed: %v", parseErr)
					return nil, fmt.Errorf("XML parsing failed: %w", parseErr)
				}

				log.Debugf("[DEBUG] XML response parsed successfully")
				return customResponse, nil
			} else {
				log.Debugf("[DEBUG] No cached XML response available")
				return nil, fmt.Errorf("XML parsing failed and no cached response available: %w", err)
			}
		}

		log.Debugf("[DEBUG] ClientWrapper: SDK API call failed: %v", err)
		return nil, err
	}
	log.Debugf("[DEBUG] ClientWrapper: SDK API call completed successfully")
	return resp, nil
}

// ListMcpImages wraps the SDK client method
func (cw *clientWrapper) ListMcpImages(ctx context.Context, request *client.ListMcpImagesRequest) (*client.ListMcpImagesResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}

	// Get runtime options with debug enabled if verbose
	runtimeOptions := cw.getRuntimeOptions()

	// Log basic request information in verbose mode
	if log.GetLevel() >= log.DebugLevel {
		log.Debugf("[DEBUG] Making ListMcpImages request...")
	}

	resp, err := sdkClient.ListMcpImagesWithOptions(request, runtimeOptions)

	// Log detailed response information in verbose mode
	if log.GetLevel() >= log.DebugLevel {
		if err != nil {
			// Check if this is a known XML parsing error
			errStr := err.Error()
			if bytes.Contains([]byte(errStr), []byte("readObjectStart: expect { or n, but found")) || bytes.Contains([]byte(errStr), []byte("invalid character '<' looking for beginning of value")) {
				log.Debugf("[DEBUG] ListMcpImages HTTP Response: XML format detected, will use custom parser")
			} else {
				log.Debugf("[DEBUG] ListMcpImages HTTP Response Error: %v", err)
				// Try to extract more details from the error
				log.Debugf("[DEBUG] Error type: %T", err)
				log.Debugf("[DEBUG] Error string: %s", err.Error())
			}
		} else {
			log.Debugf("[DEBUG] ListMcpImages request completed successfully")
		}
	}

	if err != nil {
		// Check if this is an XML parsing error
		errStr := err.Error()
		if bytes.Contains([]byte(errStr), []byte("readObjectStart: expect { or n, but found")) || bytes.Contains([]byte(errStr), []byte("invalid character '<' looking for beginning of value")) {
			log.Debugf("[DEBUG] SDK returned XML response, using custom XML parser...")

			// Use cached XML response if available
			if xmlResponseCache != "" {
				log.Debugf("[DEBUG] Parsing cached XML response...")

				// Parse the cached XML directly
				customResponse, parseErr := cw.parseListMcpImagesXMLResponse([]byte(xmlResponseCache))
				if parseErr != nil {
					log.Debugf("[DEBUG] Custom XML parsing failed: %v", parseErr)
					return nil, fmt.Errorf("XML parsing failed: %w", parseErr)
				}

				log.Debugf("[DEBUG] XML response parsed successfully")
				return customResponse, nil
			} else {
				log.Debugf("[DEBUG] No cached XML response available")
				return nil, fmt.Errorf("XML parsing failed and no cached response available: %w", err)
			}
		}

		log.Debugf("[DEBUG] ClientWrapper: SDK API call failed: %v", err)
		return nil, err
	}
	log.Debugf("[DEBUG] ClientWrapper: SDK API call completed successfully")
	return resp, nil
}

func (cw *clientWrapper) CreateResourceGroup(ctx context.Context, request *client.CreateResourceGroupRequest) (*client.CreateResourceGroupResponse, error) {
	log.Debugf("[DEBUG] ClientWrapper: CreateResourceGroup called")

	// Log request details in verbose mode
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

	// Get SDK client
	sdkClient, err := cw.getClient()
	if err != nil {
		log.Debugf("[DEBUG] ClientWrapper: Failed to get SDK client: %v", err)
		return nil, err
	}

	// Get runtime options
	runtimeOptions := cw.getRuntimeOptions()

	// Log basic request information in verbose mode
	if log.GetLevel() >= log.DebugLevel {
		log.Debugf("[DEBUG] Making CreateResourceGroup request...")
	}

	// Call the underlying SDK method
	resp, err := sdkClient.CreateResourceGroupWithContext(ctx, request, runtimeOptions)

	// Log detailed response information in verbose mode
	if log.GetLevel() >= log.DebugLevel {
		if err != nil {
			// Check if this is a known XML parsing error
			errStr := err.Error()
			if bytes.Contains([]byte(errStr), []byte("invalid character '<' looking for beginning of value")) || bytes.Contains([]byte(errStr), []byte("readObjectStart: expect { or n, but found")) {
				log.Debugf("[DEBUG] CreateResourceGroup HTTP Response: XML format detected, will use custom parser")
			} else {
				log.Debugf("[DEBUG] CreateResourceGroup HTTP Response Error: %v", err)
				log.Debugf("[DEBUG] Error type: %T", err)
				log.Debugf("[DEBUG] Error string: %s", err.Error())
			}
		} else {
			log.Debugf("[DEBUG] CreateResourceGroup request completed successfully")
		}
	}

	if err != nil {
		// Check if this is an XML parsing error
		errStr := err.Error()
		if bytes.Contains([]byte(errStr), []byte("invalid character '<' looking for beginning of value")) || bytes.Contains([]byte(errStr), []byte("readObjectStart: expect { or n, but found")) {
			log.Debugf("[DEBUG] SDK returned XML response, using custom XML parser...")

			// Use cached XML response if available
			if xmlResponseCache != "" {
				log.Debugf("[DEBUG] Parsing cached XML response...")

				// Parse the cached XML directly
				customResponse, parseErr := cw.parseCreateResourceGroupXMLResponse([]byte(xmlResponseCache))
				if parseErr != nil {
					log.Debugf("[DEBUG] Custom XML parsing failed: %v", parseErr)
					return nil, fmt.Errorf("XML parsing failed: %w", parseErr)
				}

				log.Debugf("[DEBUG] XML response parsed successfully")
				return customResponse, nil
			} else {
				log.Debugf("[DEBUG] No cached XML response available")
				return nil, fmt.Errorf("XML parsing failed and no cached response available: %w", err)
			}
		}

		log.Debugf("[DEBUG] ClientWrapper: CreateResourceGroup SDK call failed: %v", err)
		return nil, err
	}

	log.Debugf("[DEBUG] ClientWrapper: CreateResourceGroup completed successfully")
	return resp, nil
}

func (cw *clientWrapper) DeleteResourceGroup(ctx context.Context, request *client.DeleteResourceGroupRequest) (*client.DeleteResourceGroupResponse, error) {
	log.Debugf("[DEBUG] ClientWrapper: DeleteResourceGroup called")
	if log.GetLevel() >= log.DebugLevel {
		log.Debugf("[DEBUG] ClientWrapper: Request ImageId = %v", request.GetImageId())
		if request.GetImageId() != nil {
			log.Debugf("[DEBUG] ClientWrapper: ImageId value = %s", *request.GetImageId())
		}
	}

	// Get SDK client
	sdkClient, err := cw.getClient()
	if err != nil {
		log.Debugf("[DEBUG] ClientWrapper: Failed to get SDK client: %v", err)
		return nil, err
	}

	// Get runtime options
	runtimeOptions := cw.getRuntimeOptions()

	// Log basic request information in verbose mode
	if log.GetLevel() >= log.DebugLevel {
		log.Debugf("[DEBUG] Making DeleteResourceGroup request...")
	}

	// Call the underlying SDK method
	resp, err := sdkClient.DeleteResourceGroupWithContext(ctx, request, runtimeOptions)

	// Log detailed response information in verbose mode
	if log.GetLevel() >= log.DebugLevel {
		if err != nil {
			// Check if this is a known XML parsing error
			errStr := err.Error()
			if bytes.Contains([]byte(errStr), []byte("invalid character '<' looking for beginning of value")) || bytes.Contains([]byte(errStr), []byte("readObjectStart: expect { or n, but found")) {
				log.Debugf("[DEBUG] DeleteResourceGroup HTTP Response: XML format detected, will use custom parser")
			} else {
				log.Debugf("[DEBUG] DeleteResourceGroup HTTP Response Error: %v", err)
				log.Debugf("[DEBUG] Error type: %T", err)
				log.Debugf("[DEBUG] Error string: %s", err.Error())
			}
		} else {
			log.Debugf("[DEBUG] DeleteResourceGroup request completed successfully")
		}
	}

	if err != nil {
		// Check if this is an XML parsing error
		errStr := err.Error()
		if bytes.Contains([]byte(errStr), []byte("invalid character '<' looking for beginning of value")) || bytes.Contains([]byte(errStr), []byte("readObjectStart: expect { or n, but found")) {
			log.Debugf("[DEBUG] SDK returned XML response, using custom XML parser...")

			// Use cached XML response if available
			if xmlResponseCache != "" {
				log.Debugf("[DEBUG] Parsing cached XML response...")

				// Parse the cached XML directly
				customResponse, parseErr := cw.parseDeleteResourceGroupXMLResponse([]byte(xmlResponseCache))
				if parseErr != nil {
					log.Debugf("[DEBUG] Custom XML parsing failed: %v", parseErr)
					return nil, fmt.Errorf("XML parsing failed: %w", parseErr)
				}

				log.Debugf("[DEBUG] XML response parsed successfully")
				return customResponse, nil
			} else {
				log.Debugf("[DEBUG] No cached XML response available")
				return nil, fmt.Errorf("XML parsing failed and no cached response available: %w", err)
			}
		}

		log.Debugf("[DEBUG] ClientWrapper: DeleteResourceGroup SDK call failed: %v", err)
		return nil, err
	}

	log.Debugf("[DEBUG] ClientWrapper: DeleteResourceGroup completed successfully")
	return resp, nil
}

// parseGetMcpImageInfoXMLResponse parses XML response and converts it to SDK response format
func (cw *clientWrapper) parseGetMcpImageInfoXMLResponse(xmlData []byte) (*client.GetMcpImageInfoResponse, error) {
	log.Debugf("[DEBUG] Parsing GetMcpImageInfo XML data: %s", string(xmlData))

	var xmlResp XMLGetMcpImageInfoResponse
	if err := xml.Unmarshal(xmlData, &xmlResp); err != nil {
		log.Debugf("[DEBUG] GetMcpImageInfo XML unmarshal failed: %v", err)
		return nil, fmt.Errorf("failed to parse GetMcpImageInfo XML response: %w", err)
	}

	log.Debugf("[DEBUG] Parsed GetMcpImageInfo XML data:")
	log.Debugf("[DEBUG] - RequestId: %s", xmlResp.RequestId)
	log.Debugf("[DEBUG] - HttpStatusCode: %d", xmlResp.HttpStatusCode)
	log.Debugf("[DEBUG] - Code: %s", xmlResp.Code)
	log.Debugf("[DEBUG] - Success: %t", xmlResp.Success)
	log.Debugf("[DEBUG] - Message: %s", xmlResp.Message)
	log.Debugf("[DEBUG] - ImageId: %s", xmlResp.Data.ImageId)
	log.Debugf("[DEBUG] - ImageBuildType: %s", xmlResp.Data.ImageBuildType)
	log.Debugf("[DEBUG] - ImageResourceStatus: %s", xmlResp.Data.ImageResourceStatus)

	// Convert ImageInfo
	var imageInfo *client.GetMcpImageInfoResponseBodyDataImageInfo
	if xmlResp.Data.ImageInfo.OsName != "" || xmlResp.Data.ImageInfo.Status != "" {
		imageInfo = &client.GetMcpImageInfoResponseBodyDataImageInfo{
			OsName:         dara.String(xmlResp.Data.ImageInfo.OsName),
			OsVersion:      dara.String(xmlResp.Data.ImageInfo.OsVersion),
			PlatformName:   dara.String(xmlResp.Data.ImageInfo.PlatformName),
			Status:         dara.String(xmlResp.Data.ImageInfo.Status),
			DataDiskSize:   dara.Int32(xmlResp.Data.ImageInfo.DataDiskSize),
			SystemDiskSize: dara.Int32(xmlResp.Data.ImageInfo.SystemDiskSize),
			UpdateTime:     dara.String(xmlResp.Data.ImageInfo.UpdateTime),
		}
	}

	// Convert ImageBuildInfo
	var imageBuildInfo *client.GetMcpImageInfoResponseBodyDataImageBuildInfo
	if xmlResp.Data.ImageBuildInfo.TaskId != "" || xmlResp.Data.ImageBuildInfo.VersionId != "" {
		imageBuildInfo = &client.GetMcpImageInfoResponseBodyDataImageBuildInfo{
			TaskId:                  dara.String(xmlResp.Data.ImageBuildInfo.TaskId),
			VersionId:               dara.String(xmlResp.Data.ImageBuildInfo.VersionId),
			ApiKeyId:                dara.String(xmlResp.Data.ImageBuildInfo.ApiKeyId),
			InstanceReady:           dara.Bool(xmlResp.Data.ImageBuildInfo.InstanceReady),
			AndroidMobileGroupId:    dara.String(xmlResp.Data.ImageBuildInfo.AndroidMobileGroupId),
			AndroidMobileInstanceId: dara.String(xmlResp.Data.ImageBuildInfo.AndroidMobileInstanceId),
		}
	}

	// Convert to SDK response format
	response := &client.GetMcpImageInfoResponse{
		Headers:    make(map[string]*string),
		StatusCode: dara.Int32(int32(xmlResp.HttpStatusCode)),
		Body: &client.GetMcpImageInfoResponseBody{
			RequestId:      dara.String(xmlResp.RequestId),
			HttpStatusCode: dara.Int32(int32(xmlResp.HttpStatusCode)),
			Code:           dara.String(xmlResp.Code),
			Success:        dara.Bool(xmlResp.Success),
			Message:        dara.String(xmlResp.Message),
			Data: &client.GetMcpImageInfoResponseBodyData{
				ImageId:         dara.String(xmlResp.Data.ImageId),
				ImageName:       dara.String(xmlResp.Data.ImageName),
				ImageBuildType:  dara.String(xmlResp.Data.ImageBuildType),
				ImageApplyScene: dara.String(xmlResp.Data.ImageApplyScene),
				ImageInfo:       imageInfo,
				ImageBuildInfo:  imageBuildInfo,
			},
		},
	}

	// Store ImageResourceStatus and ImageType in response headers as SDK model doesn't have these fields
	// This allows us to retrieve them from the response when needed
	if xmlResp.Data.ImageResourceStatus != "" {
		response.Headers["X-Image-Resource-Status"] = dara.String(xmlResp.Data.ImageResourceStatus)
		log.Debugf("[DEBUG] Stored ImageResourceStatus in header: %s", xmlResp.Data.ImageResourceStatus)
	}

	if xmlResp.Data.ImageInfo.ImageType != "" {
		response.Headers["X-Image-Type"] = dara.String(xmlResp.Data.ImageInfo.ImageType)
		log.Debugf("[DEBUG] Stored ImageType in header: %s", xmlResp.Data.ImageInfo.ImageType)
	}

	log.Debugf("[DEBUG] Created GetMcpImageInfo SDK response")
	return response, nil
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

	// Get SDK client
	sdkClient, err := cw.getClient()
	if err != nil {
		log.Debugf("[DEBUG] ClientWrapper: Failed to get SDK client: %v", err)
		return nil, err
	}

	// Get runtime options
	runtimeOptions := cw.getRuntimeOptions()

	// Log basic request information in verbose mode
	if log.GetLevel() >= log.DebugLevel {
		log.Debugf("[DEBUG] Making GetMcpImageInfo request...")
	}

	// Call the underlying SDK method
	resp, err := sdkClient.GetMcpImageInfoWithContext(ctx, request, runtimeOptions)

	// Log detailed response information in verbose mode
	if log.GetLevel() >= log.DebugLevel {
		if err != nil {
			// Check if this is a known XML parsing error
			errStr := err.Error()
			if bytes.Contains([]byte(errStr), []byte("invalid character '<' looking for beginning of value")) || bytes.Contains([]byte(errStr), []byte("readObjectStart: expect { or n, but found")) {
				log.Debugf("[DEBUG] GetMcpImageInfo HTTP Response: XML format detected, will use custom parser")
			} else {
				log.Debugf("[DEBUG] GetMcpImageInfo HTTP Response Error: %v", err)
				log.Debugf("[DEBUG] Error type: %T", err)
				log.Debugf("[DEBUG] Error string: %s", err.Error())
			}
		} else {
			log.Debugf("[DEBUG] GetMcpImageInfo request completed successfully")
		}
	}

	if err != nil {
		// Check if this is an XML parsing error
		errStr := err.Error()
		if bytes.Contains([]byte(errStr), []byte("invalid character '<' looking for beginning of value")) || bytes.Contains([]byte(errStr), []byte("readObjectStart: expect { or n, but found")) {
			log.Debugf("[DEBUG] SDK returned XML response, using custom XML parser...")

			// Use cached XML response if available
			if xmlResponseCache != "" {
				log.Debugf("[DEBUG] Parsing cached XML response...")

				// Parse the cached XML directly
				customResponse, parseErr := cw.parseGetMcpImageInfoXMLResponse([]byte(xmlResponseCache))
				if parseErr != nil {
					log.Debugf("[DEBUG] Custom XML parsing failed: %v", parseErr)
					return nil, fmt.Errorf("XML parsing failed: %w", parseErr)
				}

				log.Debugf("[DEBUG] XML response parsed successfully")
				return customResponse, nil
			} else {
				log.Debugf("[DEBUG] No cached XML response available")
				return nil, fmt.Errorf("XML parsing failed and no cached response available: %w", err)
			}
		}

		log.Debugf("[DEBUG] ClientWrapper: GetMcpImageInfo SDK call failed: %v", err)
		return nil, err
	}

	log.Debugf("[DEBUG] ClientWrapper: GetMcpImageInfo completed successfully")
	return resp, nil
}

// XMLGetDockerfileTemplateResponse represents the XML response structure for GetDockerfileTemplate
type XMLGetDockerfileTemplateResponse struct {
	XMLName        xml.Name `xml:"GetDockerfileTemplateResponse"`
	RequestId      string   `xml:"RequestId"`
	HttpStatusCode int      `xml:"HttpStatusCode"`
	Data           struct {
		OssDownloadUrl    string `xml:"OssDownloadUrl"`
		NonEditLineNum    int32  `xml:"NonEditLineNum"`
		DockerfileContent string `xml:"DockerfileContent"`
	} `xml:"Data"`
	Code    string `xml:"Code"`
	Success bool   `xml:"Success"`
	Message string `xml:"Message"`
}

// parseGetDockerfileTemplateXMLResponse parses XML response and converts it to SDK response format
func (cw *clientWrapper) parseGetDockerfileTemplateXMLResponse(xmlData []byte) (*client.GetDockerfileTemplateResponse, error) {
	log.Debugf("[DEBUG] Parsing GetDockerfileTemplate XML data: %s", string(xmlData))

	var xmlResp XMLGetDockerfileTemplateResponse
	if err := xml.Unmarshal(xmlData, &xmlResp); err != nil {
		log.Debugf("[DEBUG] GetDockerfileTemplate XML unmarshal failed: %v", err)
		return nil, fmt.Errorf("failed to parse GetDockerfileTemplate XML response: %w", err)
	}

	log.Debugf("[DEBUG] Parsed GetDockerfileTemplate XML data:")
	log.Debugf("[DEBUG] - RequestId: %s", xmlResp.RequestId)
	log.Debugf("[DEBUG] - HttpStatusCode: %d", xmlResp.HttpStatusCode)
	log.Debugf("[DEBUG] - Code: %s", xmlResp.Code)
	log.Debugf("[DEBUG] - Success: %t", xmlResp.Success)
	log.Debugf("[DEBUG] - Message: %s", xmlResp.Message)
	log.Debugf("[DEBUG] - OssDownloadUrl: %s", xmlResp.Data.OssDownloadUrl)
	log.Debugf("[DEBUG] - NonEditLineNum: %d", xmlResp.Data.NonEditLineNum)
	log.Debugf("[DEBUG] - DockerfileContent length: %d", len(xmlResp.Data.DockerfileContent))

	// Convert to SDK response format
	response := &client.GetDockerfileTemplateResponse{
		Headers:    make(map[string]*string),
		StatusCode: dara.Int32(int32(xmlResp.HttpStatusCode)),
		Body: &client.GetDockerfileTemplateResponseBody{
			RequestId:      dara.String(xmlResp.RequestId),
			HttpStatusCode: dara.Int32(int32(xmlResp.HttpStatusCode)),
			Code:           dara.String(xmlResp.Code),
			Success:        dara.Bool(xmlResp.Success),
			Message:        dara.String(xmlResp.Message),
			Data: &client.GetDockerfileTemplateResponseBodyData{
				OssDownloadUrl:    dara.String(xmlResp.Data.OssDownloadUrl),
				NonEditLineNum:    dara.Int32(xmlResp.Data.NonEditLineNum),
				DockerfileContent: dara.String(xmlResp.Data.DockerfileContent),
			},
		},
	}

	log.Debugf("[DEBUG] Created GetDockerfileTemplate SDK response with OssDownloadUrl: %s, NonEditLineNum: %d", xmlResp.Data.OssDownloadUrl, xmlResp.Data.NonEditLineNum)
	return response, nil
}

// GetDockerfileTemplate wraps the SDK client method
func (cw *clientWrapper) GetDockerfileTemplate(ctx context.Context, request *client.GetDockerfileTemplateRequest) (*client.GetDockerfileTemplateResponse, error) {
	sdkClient, err := cw.getClient()
	if err != nil {
		return nil, err
	}

	// Get runtime options with debug enabled if verbose
	runtimeOptions := cw.getRuntimeOptions()

	// Log basic request information in verbose mode
	if log.GetLevel() >= log.DebugLevel {
		log.Debugf("[DEBUG] Making GetDockerfileTemplate request...")
	}

	// Validate request
	if err := request.Validate(); err != nil {
		return nil, err
	}

	// Prepare query parameters
	query := map[string]interface{}{}
	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}
	if !dara.IsNil(request.SourceImageId) {
		query["SourceImageId"] = request.SourceImageId
	}
	// Template is reserved for future extension
	if !dara.IsNil(request.Template) {
		query["Template"] = request.Template
	}

	// Prepare API request
	apiRequest := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Headers: map[string]*string{
			"Accept": dara.String("application/json"),
		},
	}

	// Prepare API parameters
	params := &openapiutil.Params{
		Action:      dara.String("GetDockerfileTemplate"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}

	// Call API
	result := &client.GetDockerfileTemplateResponse{}
	body, err := sdkClient.CallApi(params, apiRequest, runtimeOptions)
	if err != nil {
		log.Debugf("[DEBUG] GetDockerfileTemplate API call failed: %v", err)

		// Check if this is an XML parsing error (in case API returns XML instead of JSON)
		errStr := err.Error()
		if bytes.Contains([]byte(errStr), []byte("readObjectStart: expect { or n, but found")) || bytes.Contains([]byte(errStr), []byte("invalid character '<' looking for beginning of value")) {
			log.Debugf("[DEBUG] SDK returned XML response, using custom XML parser...")

			// Use cached XML response if available
			if xmlResponseCache != "" {
				log.Debugf("[DEBUG] Parsing cached XML response...")

				// Parse the cached XML directly
				customResponse, parseErr := cw.parseGetDockerfileTemplateXMLResponse([]byte(xmlResponseCache))
				if parseErr != nil {
					log.Debugf("[DEBUG] Custom XML parsing failed: %v", parseErr)
					return nil, fmt.Errorf("XML parsing failed: %w", parseErr)
				}

				log.Debugf("[DEBUG] XML response parsed successfully")
				return customResponse, nil
			} else {
				log.Debugf("[DEBUG] No cached XML response available")
				return nil, fmt.Errorf("XML parsing failed and no cached response available: %w", err)
			}
		}

		return nil, err
	}

	// Convert response body to result
	if err := dara.Convert(body, &result); err != nil {
		log.Debugf("[DEBUG] Failed to convert response: %v", err)
		return nil, fmt.Errorf("failed to convert response: %w", err)
	}

	log.Debugf("[DEBUG] ClientWrapper: GetDockerfileTemplate completed successfully")
	return result, nil
}
