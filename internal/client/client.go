// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

type Client struct {
	openapi.Client
	DisableSDKError *bool
}

func NewClient(config *openapiutil.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapiutil.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("xiaoying"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !dara.IsNil(endpoint) {
		_result = endpoint
		return _result, _err
	}

	if !dara.IsNil(endpointMap) && !dara.IsNil(endpointMap[dara.StringValue(regionId)]) {
		_result = endpointMap[dara.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := openapiutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建Mcp资源
//
// @param request - AddMcpResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddMcpResourceResponse
func (client *Client) AddMcpResourceWithOptions(request *AddMcpResourceRequest, runtime *dara.RuntimeOptions) (_result *AddMcpResourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKeyId) {
		query["ApiKeyId"] = request.ApiKeyId
	}

	if !dara.IsNil(request.Concurrency) {
		query["Concurrency"] = request.Concurrency
	}

	if !dara.IsNil(request.DesktopMaxRuntime) {
		query["DesktopMaxRuntime"] = request.DesktopMaxRuntime
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.McpIdleTimeout) {
		query["McpIdleTimeout"] = request.McpIdleTimeout
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.UserIdleTimeout) {
		query["UserIdleTimeout"] = request.UserIdleTimeout
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddMcpResource"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddMcpResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建Mcp资源
//
// @param request - AddMcpResourceRequest
//
// @return AddMcpResourceResponse
func (client *Client) AddMcpResource(request *AddMcpResourceRequest) (_result *AddMcpResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddMcpResourceResponse{}
	_body, _err := client.AddMcpResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建镜像变更
//
// @param request - BuildMcpImageByParentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BuildMcpImageByParentResponse
func (client *Client) BuildMcpImageByParentWithOptions(request *BuildMcpImageByParentRequest, runtime *dara.RuntimeOptions) (_result *BuildMcpImageByParentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataDiskSize) {
		query["DataDiskSize"] = request.DataDiskSize
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ImageName) {
		query["ImageName"] = request.ImageName
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.OsType) {
		query["OsType"] = request.OsType
	}

	if !dara.IsNil(request.ParentImageId) {
		query["ParentImageId"] = request.ParentImageId
	}

	if !dara.IsNil(request.PerformanceType) {
		query["PerformanceType"] = request.PerformanceType
	}

	if !dara.IsNil(request.SystemDiskSize) {
		query["SystemDiskSize"] = request.SystemDiskSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BuildMcpImageByParent"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BuildMcpImageByParentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建镜像变更
//
// @param request - BuildMcpImageByParentRequest
//
// @return BuildMcpImageByParentResponse
func (client *Client) BuildMcpImageByParent(request *BuildMcpImageByParentRequest) (_result *BuildMcpImageByParentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BuildMcpImageByParentResponse{}
	_body, _err := client.BuildMcpImageByParentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建apiKey
//
// @param request - CreateApiKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApiKeyResponse
func (client *Client) CreateApiKeyWithOptions(request *CreateApiKeyRequest, runtime *dara.RuntimeOptions) (_result *CreateApiKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApiKey"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApiKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建apiKey
//
// @param request - CreateApiKeyRequest
//
// @return CreateApiKeyResponse
func (client *Client) CreateApiKey(request *CreateApiKeyRequest) (_result *CreateApiKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateApiKeyResponse{}
	_body, _err := client.CreateApiKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建docker镜像任务
//
// @param request - CreateDockerImageTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDockerImageTaskResponse
func (client *Client) CreateDockerImageTaskWithOptions(request *CreateDockerImageTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateDockerImageTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageName) {
		query["ImageName"] = request.ImageName
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.SourceImageId) {
		query["SourceImageId"] = request.SourceImageId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Headers: map[string]*string{
			"Accept": dara.String("application/xml"),
		},
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDockerImageTask"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("xml"),
	}
	_result = &CreateDockerImageTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建docker镜像任务
//
// @param request - CreateDockerImageTaskRequest
//
// @return CreateDockerImageTaskResponse
func (client *Client) CreateDockerImageTask(request *CreateDockerImageTaskRequest) (_result *CreateDockerImageTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDockerImageTaskResponse{}
	_body, _err := client.CreateDockerImageTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建MCP镜像
//
// @param request - CreateMcpImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMcpImageResponse
func (client *Client) CreateMcpImageWithOptions(request *CreateMcpImageRequest, runtime *dara.RuntimeOptions) (_result *CreateMcpImageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKeyId) {
		query["ApiKeyId"] = request.ApiKeyId
	}

	if !dara.IsNil(request.CustomerAliUid) {
		query["CustomerAliUid"] = request.CustomerAliUid
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.ImageName) {
		query["ImageName"] = request.ImageName
	}

	if !dara.IsNil(request.ParentImageId) {
		query["ParentImageId"] = request.ParentImageId
	}

	if !dara.IsNil(request.PersistentId) {
		query["PersistentId"] = request.PersistentId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMcpImage"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMcpImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建MCP镜像
//
// @param request - CreateMcpImageRequest
//
// @return CreateMcpImageResponse
func (client *Client) CreateMcpImage(request *CreateMcpImageRequest) (_result *CreateMcpImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMcpImageResponse{}
	_body, _err := client.CreateMcpImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建策略
//
// @param tmpReq - CreateMcpNetworkWhitelistRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMcpNetworkWhitelistRulesResponse
func (client *Client) CreateMcpNetworkWhitelistRulesWithOptions(tmpReq *CreateMcpNetworkWhitelistRulesRequest, runtime *dara.RuntimeOptions) (_result *CreateMcpNetworkWhitelistRulesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateMcpNetworkWhitelistRulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Rules) {
		request.RulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Rules, dara.String("Rules"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.McpPolicyId) {
		body["McpPolicyId"] = request.McpPolicyId
	}

	if !dara.IsNil(request.RulesShrink) {
		body["Rules"] = request.RulesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMcpNetworkWhitelistRules"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMcpNetworkWhitelistRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建策略
//
// @param request - CreateMcpNetworkWhitelistRulesRequest
//
// @return CreateMcpNetworkWhitelistRulesResponse
func (client *Client) CreateMcpNetworkWhitelistRules(request *CreateMcpNetworkWhitelistRulesRequest) (_result *CreateMcpNetworkWhitelistRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMcpNetworkWhitelistRulesResponse{}
	_body, _err := client.CreateMcpNetworkWhitelistRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建策略
//
// @param tmpReq - CreateMcpPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMcpPolicyResponse
func (client *Client) CreateMcpPolicyWithOptions(tmpReq *CreateMcpPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateMcpPolicyResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateMcpPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DisplayConfig) {
		request.DisplayConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DisplayConfig, dara.String("DisplayConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.NetworkConfig) {
		request.NetworkConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NetworkConfig, dara.String("NetworkConfig"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Concurrency) {
		body["Concurrency"] = request.Concurrency
	}

	if !dara.IsNil(request.DebugMode) {
		body["DebugMode"] = request.DebugMode
	}

	if !dara.IsNil(request.DesktopMaxRuntime) {
		body["DesktopMaxRuntime"] = request.DesktopMaxRuntime
	}

	if !dara.IsNil(request.DisplayConfigShrink) {
		body["DisplayConfig"] = request.DisplayConfigShrink
	}

	if !dara.IsNil(request.IdleTimeoutSwitch) {
		body["IdleTimeoutSwitch"] = request.IdleTimeoutSwitch
	}

	if !dara.IsNil(request.McpIdleTimeout) {
		body["McpIdleTimeout"] = request.McpIdleTimeout
	}

	if !dara.IsNil(request.NetworkConfigShrink) {
		body["NetworkConfig"] = request.NetworkConfigShrink
	}

	if !dara.IsNil(request.PolicyName) {
		body["PolicyName"] = request.PolicyName
	}

	if !dara.IsNil(request.UserIdleTimeout) {
		body["UserIdleTimeout"] = request.UserIdleTimeout
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMcpPolicy"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMcpPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建策略
//
// @param request - CreateMcpPolicyRequest
//
// @return CreateMcpPolicyResponse
func (client *Client) CreateMcpPolicy(request *CreateMcpPolicyRequest) (_result *CreateMcpPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMcpPolicyResponse{}
	_body, _err := client.CreateMcpPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建交付组
//
// @param request - CreateResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceGroupResponse
func (client *Client) CreateResourceGroupWithOptions(request *CreateResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateResourceGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BizRegionId) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.Cpu) {
		body["Cpu"] = request.Cpu
	}

	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.Memory) {
		body["Memory"] = request.Memory
	}

	if !dara.IsNil(request.OfficeSiteId) {
		body["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.OfficeSiteType) {
		body["OfficeSiteType"] = request.OfficeSiteType
	}

	if !dara.IsNil(request.PolicyId) {
		body["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SessionBandwidth) {
		body["SessionBandwidth"] = request.SessionBandwidth
	}

	if !dara.IsNil(request.VSwitchId) {
		body["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VpcId) {
		body["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateResourceGroup"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建交付组
//
// @param request - CreateResourceGroupRequest
//
// @return CreateResourceGroupResponse
func (client *Client) CreateResourceGroup(request *CreateResourceGroupRequest) (_result *CreateResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateResourceGroupResponse{}
	_body, _err := client.CreateResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除Mcp资源
//
// @param request - DelMcpResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DelMcpResourceResponse
func (client *Client) DelMcpResourceWithOptions(request *DelMcpResourceRequest, runtime *dara.RuntimeOptions) (_result *DelMcpResourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKeyId) {
		query["ApiKeyId"] = request.ApiKeyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DelMcpResource"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DelMcpResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除Mcp资源
//
// @param request - DelMcpResourceRequest
//
// @return DelMcpResourceResponse
func (client *Client) DelMcpResource(request *DelMcpResourceRequest) (_result *DelMcpResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DelMcpResourceResponse{}
	_body, _err := client.DelMcpResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除Mcp镜像
//
// @param request - DeleteMcpImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMcpImageResponse
func (client *Client) DeleteMcpImageWithOptions(request *DeleteMcpImageRequest, runtime *dara.RuntimeOptions) (_result *DeleteMcpImageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMcpImage"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMcpImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除Mcp镜像
//
// @param request - DeleteMcpImageRequest
//
// @return DeleteMcpImageResponse
func (client *Client) DeleteMcpImage(request *DeleteMcpImageRequest) (_result *DeleteMcpImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMcpImageResponse{}
	_body, _err := client.DeleteMcpImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除白名单策略
//
// @param tmpReq - DeleteMcpNetworkWhitelistRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMcpNetworkWhitelistRulesResponse
func (client *Client) DeleteMcpNetworkWhitelistRulesWithOptions(tmpReq *DeleteMcpNetworkWhitelistRulesRequest, runtime *dara.RuntimeOptions) (_result *DeleteMcpNetworkWhitelistRulesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DeleteMcpNetworkWhitelistRulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RuleIds) {
		request.RuleIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RuleIds, dara.String("RuleIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.McpPolicyId) {
		body["McpPolicyId"] = request.McpPolicyId
	}

	if !dara.IsNil(request.RuleIdsShrink) {
		body["RuleIds"] = request.RuleIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMcpNetworkWhitelistRules"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMcpNetworkWhitelistRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除白名单策略
//
// @param request - DeleteMcpNetworkWhitelistRulesRequest
//
// @return DeleteMcpNetworkWhitelistRulesResponse
func (client *Client) DeleteMcpNetworkWhitelistRules(request *DeleteMcpNetworkWhitelistRulesRequest) (_result *DeleteMcpNetworkWhitelistRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMcpNetworkWhitelistRulesResponse{}
	_body, _err := client.DeleteMcpNetworkWhitelistRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除策略
//
// @param request - DeleteMcpPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMcpPoliciesResponse
func (client *Client) DeleteMcpPoliciesWithOptions(request *DeleteMcpPoliciesRequest, runtime *dara.RuntimeOptions) (_result *DeleteMcpPoliciesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PolicyIds) {
		body["PolicyIds"] = request.PolicyIds
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMcpPolicies"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMcpPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除策略
//
// @param request - DeleteMcpPoliciesRequest
//
// @return DeleteMcpPoliciesResponse
func (client *Client) DeleteMcpPolicies(request *DeleteMcpPoliciesRequest) (_result *DeleteMcpPoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMcpPoliciesResponse{}
	_body, _err := client.DeleteMcpPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除交付组
//
// @param request - DeleteResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceGroupResponse
func (client *Client) DeleteResourceGroupWithOptions(request *DeleteResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteResourceGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteResourceGroup"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除交付组
//
// @param request - DeleteResourceGroupRequest
//
// @return DeleteResourceGroupResponse
func (client *Client) DeleteResourceGroup(request *DeleteResourceGroupRequest) (_result *DeleteResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteResourceGroupResponse{}
	_body, _err := client.DeleteResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询context特定目录文件
//
// @param request - DescribeContextFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeContextFilesResponse
func (client *Client) DescribeContextFilesWithOptions(request *DescribeContextFilesRequest, runtime *dara.RuntimeOptions) (_result *DescribeContextFilesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ContextId) {
		body["ContextId"] = request.ContextId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ParentFolderPath) {
		body["ParentFolderPath"] = request.ParentFolderPath
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeContextFiles"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeContextFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询context特定目录文件
//
// @param request - DescribeContextFilesRequest
//
// @return DescribeContextFilesResponse
func (client *Client) DescribeContextFiles(request *DescribeContextFilesRequest) (_result *DescribeContextFilesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeContextFilesResponse{}
	_body, _err := client.DescribeContextFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除白名单策略
//
// @param request - DescribeMcpNetworkWhitelistRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMcpNetworkWhitelistRulesResponse
func (client *Client) DescribeMcpNetworkWhitelistRulesWithOptions(request *DescribeMcpNetworkWhitelistRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeMcpNetworkWhitelistRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.McpPolicyId) {
		body["McpPolicyId"] = request.McpPolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMcpNetworkWhitelistRules"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMcpNetworkWhitelistRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除白名单策略
//
// @param request - DescribeMcpNetworkWhitelistRulesRequest
//
// @return DescribeMcpNetworkWhitelistRulesResponse
func (client *Client) DescribeMcpNetworkWhitelistRules(request *DescribeMcpNetworkWhitelistRulesRequest) (_result *DescribeMcpNetworkWhitelistRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMcpNetworkWhitelistRulesResponse{}
	_body, _err := client.DescribeMcpNetworkWhitelistRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询策略
//
// @param request - DescribeMcpPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMcpPoliciesResponse
func (client *Client) DescribeMcpPoliciesWithOptions(request *DescribeMcpPoliciesRequest, runtime *dara.RuntimeOptions) (_result *DescribeMcpPoliciesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ExcludedPolicyIds) {
		body["ExcludedPolicyIds"] = request.ExcludedPolicyIds
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PolicyIds) {
		body["PolicyIds"] = request.PolicyIds
	}

	if !dara.IsNil(request.PolicyName) {
		body["PolicyName"] = request.PolicyName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMcpPolicies"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMcpPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询策略
//
// @param request - DescribeMcpPoliciesRequest
//
// @return DescribeMcpPoliciesResponse
func (client *Client) DescribeMcpPolicies(request *DescribeMcpPoliciesRequest) (_result *DescribeMcpPoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMcpPoliciesResponse{}
	_body, _err := client.DescribeMcpPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 统计context用量
//
// @param request - DescribeSessionContextRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSessionContextResponse
func (client *Client) DescribeSessionContextWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeSessionContextResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSessionContext"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSessionContextResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 统计context用量
//
// @return DescribeSessionContextResponse
func (client *Client) DescribeSessionContext() (_result *DescribeSessionContextResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSessionContextResponse{}
	_body, _err := client.DescribeSessionContextWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 统计会话数
//
// @param request - DescribeSessionCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSessionCountResponse
func (client *Client) DescribeSessionCountWithOptions(request *DescribeSessionCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeSessionCountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKeyId) {
		query["ApiKeyId"] = request.ApiKeyId
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSessionCount"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSessionCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 统计会话数
//
// @param request - DescribeSessionCountRequest
//
// @return DescribeSessionCountResponse
func (client *Client) DescribeSessionCount(request *DescribeSessionCountRequest) (_result *DescribeSessionCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSessionCountResponse{}
	_body, _err := client.DescribeSessionCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 统计会话时长
//
// @param request - DescribeSessionDurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSessionDurationResponse
func (client *Client) DescribeSessionDurationWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeSessionDurationResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSessionDuration"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSessionDurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 统计会话时长
//
// @return DescribeSessionDurationResponse
func (client *Client) DescribeSessionDuration() (_result *DescribeSessionDurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSessionDurationResponse{}
	_body, _err := client.DescribeSessionDurationWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户权益
//
// @param request - DescribeUserBenefitDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserBenefitDetailResponse
func (client *Client) DescribeUserBenefitDetailWithOptions(request *DescribeUserBenefitDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserBenefitDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SessionBandwidth) {
		body["SessionBandwidth"] = request.SessionBandwidth
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserBenefitDetail"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserBenefitDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户权益
//
// @param request - DescribeUserBenefitDetailRequest
//
// @return DescribeUserBenefitDetailResponse
func (client *Client) DescribeUserBenefitDetail(request *DescribeUserBenefitDetailRequest) (_result *DescribeUserBenefitDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserBenefitDetailResponse{}
	_body, _err := client.DescribeUserBenefitDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取上传context文件url
//
// @param request - GetContextFileDownloadUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetContextFileDownloadUrlResponse
func (client *Client) GetContextFileDownloadUrlWithOptions(request *GetContextFileDownloadUrlRequest, runtime *dara.RuntimeOptions) (_result *GetContextFileDownloadUrlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Authorization) {
		body["Authorization"] = request.Authorization
	}

	if !dara.IsNil(request.ContextId) {
		body["ContextId"] = request.ContextId
	}

	if !dara.IsNil(request.FilePath) {
		body["FilePath"] = request.FilePath
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetContextFileDownloadUrl"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetContextFileDownloadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取上传context文件url
//
// @param request - GetContextFileDownloadUrlRequest
//
// @return GetContextFileDownloadUrlResponse
func (client *Client) GetContextFileDownloadUrl(request *GetContextFileDownloadUrlRequest) (_result *GetContextFileDownloadUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetContextFileDownloadUrlResponse{}
	_body, _err := client.GetContextFileDownloadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取dockerfile文件放置位置
//
// @param request - GetDockerFileStoreCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDockerFileStoreCredentialResponse
func (client *Client) GetDockerFileStoreCredentialWithOptions(request *GetDockerFileStoreCredentialRequest, runtime *dara.RuntimeOptions) (_result *GetDockerFileStoreCredentialResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Headers: map[string]*string{
			"Accept": dara.String("application/xml"),
		},
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDockerFileStoreCredential"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("xml"),
	}
	_result = &GetDockerFileStoreCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取dockerfile文件放置位置
//
// @param request - GetDockerFileStoreCredentialRequest
//
// @return GetDockerFileStoreCredentialResponse
func (client *Client) GetDockerFileStoreCredential(request *GetDockerFileStoreCredentialRequest) (_result *GetDockerFileStoreCredentialResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDockerFileStoreCredentialResponse{}
	_body, _err := client.GetDockerFileStoreCredentialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取docker镜像任务详情
//
// @param request - GetDockerImageTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDockerImageTaskResponse
func (client *Client) GetDockerImageTaskWithOptions(request *GetDockerImageTaskRequest, runtime *dara.RuntimeOptions) (_result *GetDockerImageTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Headers: map[string]*string{
			"Accept": dara.String("application/xml"),
		},
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDockerImageTask"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("xml"),
	}
	_result = &GetDockerImageTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取docker镜像任务详情
//
// @param request - GetDockerImageTaskRequest
//
// @return GetDockerImageTaskResponse
func (client *Client) GetDockerImageTask(request *GetDockerImageTaskRequest) (_result *GetDockerImageTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDockerImageTaskResponse{}
	_body, _err := client.GetDockerImageTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取镜像能力配置项
//
// @param request - GetImageFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageFunctionResponse
func (client *Client) GetImageFunctionWithOptions(request *GetImageFunctionRequest, runtime *dara.RuntimeOptions) (_result *GetImageFunctionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetImageFunction"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetImageFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取镜像能力配置项
//
// @param request - GetImageFunctionRequest
//
// @return GetImageFunctionResponse
func (client *Client) GetImageFunction(request *GetImageFunctionRequest) (_result *GetImageFunctionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetImageFunctionResponse{}
	_body, _err := client.GetImageFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取最大并发数上限
//
// @param request - GetMaxConcurrencyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMaxConcurrencyResponse
func (client *Client) GetMaxConcurrencyWithOptions(runtime *dara.RuntimeOptions) (_result *GetMaxConcurrencyResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetMaxConcurrency"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMaxConcurrencyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取最大并发数上限
//
// @return GetMaxConcurrencyResponse
func (client *Client) GetMaxConcurrency() (_result *GetMaxConcurrencyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMaxConcurrencyResponse{}
	_body, _err := client.GetMaxConcurrencyWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询知识详情
//
// @param request - GetMcpConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMcpConfigResponse
func (client *Client) GetMcpConfigWithOptions(request *GetMcpConfigRequest, runtime *dara.RuntimeOptions) (_result *GetMcpConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigType) {
		query["ConfigType"] = request.ConfigType
	}

	if !dara.IsNil(request.ImageKeyword) {
		query["ImageKeyword"] = request.ImageKeyword
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMcpConfig"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMcpConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询知识详情
//
// @param request - GetMcpConfigRequest
//
// @return GetMcpConfigResponse
func (client *Client) GetMcpConfig(request *GetMcpConfigRequest) (_result *GetMcpConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMcpConfigResponse{}
	_body, _err := client.GetMcpConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # GetMcpDesktopInfo
//
// @param request - GetMcpDesktopInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMcpDesktopInfoResponse
func (client *Client) GetMcpDesktopInfoWithOptions(request *GetMcpDesktopInfoRequest, runtime *dara.RuntimeOptions) (_result *GetMcpDesktopInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKeyId) {
		query["ApiKeyId"] = request.ApiKeyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMcpDesktopInfo"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMcpDesktopInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetMcpDesktopInfo
//
// @param request - GetMcpDesktopInfoRequest
//
// @return GetMcpDesktopInfoResponse
func (client *Client) GetMcpDesktopInfo(request *GetMcpDesktopInfoRequest) (_result *GetMcpDesktopInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMcpDesktopInfoResponse{}
	_body, _err := client.GetMcpDesktopInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取mcp镜像信息
//
// @param request - GetMcpImageInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMcpImageInfoResponse
func (client *Client) GetMcpImageInfoWithOptions(request *GetMcpImageInfoRequest, runtime *dara.RuntimeOptions) (_result *GetMcpImageInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMcpImageInfo"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMcpImageInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取mcp镜像信息
//
// @param request - GetMcpImageInfoRequest
//
// @return GetMcpImageInfoResponse
func (client *Client) GetMcpImageInfo(request *GetMcpImageInfoRequest) (_result *GetMcpImageInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMcpImageInfoResponse{}
	_body, _err := client.GetMcpImageInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取镜像总览
//
// @param request - GetMcpImageOverviewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMcpImageOverviewResponse
func (client *Client) GetMcpImageOverviewWithOptions(runtime *dara.RuntimeOptions) (_result *GetMcpImageOverviewResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetMcpImageOverview"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMcpImageOverviewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取镜像总览
//
// @return GetMcpImageOverviewResponse
func (client *Client) GetMcpImageOverview() (_result *GetMcpImageOverviewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMcpImageOverviewResponse{}
	_body, _err := client.GetMcpImageOverviewWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看Mcp资源详情
//
// @param request - GetMcpResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMcpResourceResponse
func (client *Client) GetMcpResourceWithOptions(request *GetMcpResourceRequest, runtime *dara.RuntimeOptions) (_result *GetMcpResourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKeyId) {
		query["ApiKeyId"] = request.ApiKeyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMcpResource"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMcpResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看Mcp资源详情
//
// @param request - GetMcpResourceRequest
//
// @return GetMcpResourceResponse
func (client *Client) GetMcpResource(request *GetMcpResourceRequest) (_result *GetMcpResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMcpResourceResponse{}
	_body, _err := client.GetMcpResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询mcp工具信息
//
// @param request - GetMcpToolInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMcpToolInfoResponse
func (client *Client) GetMcpToolInfoWithOptions(request *GetMcpToolInfoRequest, runtime *dara.RuntimeOptions) (_result *GetMcpToolInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMcpToolInfo"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMcpToolInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询mcp工具信息
//
// @param request - GetMcpToolInfoRequest
//
// @return GetMcpToolInfoResponse
func (client *Client) GetMcpToolInfo(request *GetMcpToolInfoRequest) (_result *GetMcpToolInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMcpToolInfoResponse{}
	_body, _err := client.GetMcpToolInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取playground的配置
//
// @param request - GetPlaygroundConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPlaygroundConfigResponse
func (client *Client) GetPlaygroundConfigWithOptions(runtime *dara.RuntimeOptions) (_result *GetPlaygroundConfigResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetPlaygroundConfig"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPlaygroundConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取playground的配置
//
// @return GetPlaygroundConfigResponse
func (client *Client) GetPlaygroundConfig() (_result *GetPlaygroundConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPlaygroundConfigResponse{}
	_body, _err := client.GetPlaygroundConfigWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询playground模版详情
//
// @param request - GetPlaygroundDemoDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPlaygroundDemoDetailResponse
func (client *Client) GetPlaygroundDemoDetailWithOptions(request *GetPlaygroundDemoDetailRequest, runtime *dara.RuntimeOptions) (_result *GetPlaygroundDemoDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExampleId) {
		query["ExampleId"] = request.ExampleId
	}

	if !dara.IsNil(request.SceneType) {
		query["SceneType"] = request.SceneType
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPlaygroundDemoDetail"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPlaygroundDemoDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询playground模版详情
//
// @param request - GetPlaygroundDemoDetailRequest
//
// @return GetPlaygroundDemoDetailResponse
func (client *Client) GetPlaygroundDemoDetail(request *GetPlaygroundDemoDetailRequest) (_result *GetPlaygroundDemoDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPlaygroundDemoDetailResponse{}
	_body, _err := client.GetPlaygroundDemoDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 如果是示例代码复用已创建的函数
//
// @param request - InvokeFunctionExistedRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvokeFunctionExistedResponse
func (client *Client) InvokeFunctionExistedWithOptions(request *InvokeFunctionExistedRequest, runtime *dara.RuntimeOptions) (_result *InvokeFunctionExistedResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		query["Code"] = request.Code
	}

	if !dara.IsNil(request.SceneType) {
		query["SceneType"] = request.SceneType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InvokeFunctionExisted"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InvokeFunctionExistedResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 如果是示例代码复用已创建的函数
//
// @param request - InvokeFunctionExistedRequest
//
// @return InvokeFunctionExistedResponse
func (client *Client) InvokeFunctionExisted(request *InvokeFunctionExistedRequest) (_result *InvokeFunctionExistedResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InvokeFunctionExistedResponse{}
	_body, _err := client.InvokeFunctionExistedWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询支持mcp镜像列表
//
// @param request - ListMcpImagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMcpImagesResponse
func (client *Client) ListMcpImagesWithOptions(request *ListMcpImagesRequest, runtime *dara.RuntimeOptions) (_result *ListMcpImagesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FeatureList) {
		query["FeatureList"] = request.FeatureList
	}

	if !dara.IsNil(request.ImageType) {
		query["ImageType"] = request.ImageType
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OsType) {
		query["OsType"] = request.OsType
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PageStart) {
		query["PageStart"] = request.PageStart
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMcpImages"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMcpImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询支持mcp镜像列表
//
// @param request - ListMcpImagesRequest
//
// @return ListMcpImagesResponse
func (client *Client) ListMcpImages(request *ListMcpImagesRequest) (_result *ListMcpImagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListMcpImagesResponse{}
	_body, _err := client.ListMcpImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Mcp资源列表
//
// @param request - ListMcpResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMcpResourcesResponse
func (client *Client) ListMcpResourcesWithOptions(request *ListMcpResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListMcpResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMcpResources"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMcpResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Mcp资源列表
//
// @param request - ListMcpResourcesRequest
//
// @return ListMcpResourcesResponse
func (client *Client) ListMcpResources(request *ListMcpResourcesRequest) (_result *ListMcpResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListMcpResourcesResponse{}
	_body, _err := client.ListMcpResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询独立资源池
//
// @param request - ListResourceGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceGroupsResponse
func (client *Client) ListResourceGroupsWithOptions(request *ListResourceGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListResourceGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceGroups"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询独立资源池
//
// @param request - ListResourceGroupsRequest
//
// @return ListResourceGroupsResponse
func (client *Client) ListResourceGroups(request *ListResourceGroupsRequest) (_result *ListResourceGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListResourceGroupsResponse{}
	_body, _err := client.ListResourceGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询会话日志列表
//
// @param request - ListSessionLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSessionLogsResponse
func (client *Client) ListSessionLogsWithOptions(request *ListSessionLogsRequest, runtime *dara.RuntimeOptions) (_result *ListSessionLogsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.ImageType) {
		query["ImageType"] = request.ImageType
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSessionLogs"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSessionLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询会话日志列表
//
// @param request - ListSessionLogsRequest
//
// @return ListSessionLogsResponse
func (client *Client) ListSessionLogs(request *ListSessionLogsRequest) (_result *ListSessionLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSessionLogsResponse{}
	_body, _err := client.ListSessionLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询会话中工具调用记录
//
// @param request - ListSessionToolLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSessionToolLogsResponse
func (client *Client) ListSessionToolLogsWithOptions(request *ListSessionToolLogsRequest, runtime *dara.RuntimeOptions) (_result *ListSessionToolLogsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSessionToolLogs"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSessionToolLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询会话中工具调用记录
//
// @param request - ListSessionToolLogsRequest
//
// @return ListSessionToolLogsResponse
func (client *Client) ListSessionToolLogs(request *ListSessionToolLogsRequest) (_result *ListSessionToolLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSessionToolLogsResponse{}
	_body, _err := client.ListSessionToolLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 解除策略关联
//
// @param request - ModifyMcpApiKeyConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyMcpApiKeyConfigResponse
func (client *Client) ModifyMcpApiKeyConfigWithOptions(request *ModifyMcpApiKeyConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyMcpApiKeyConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiKeyId) {
		body["ApiKeyId"] = request.ApiKeyId
	}

	if !dara.IsNil(request.Concurrency) {
		body["Concurrency"] = request.Concurrency
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyMcpApiKeyConfig"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyMcpApiKeyConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解除策略关联
//
// @param request - ModifyMcpApiKeyConfigRequest
//
// @return ModifyMcpApiKeyConfigResponse
func (client *Client) ModifyMcpApiKeyConfig(request *ModifyMcpApiKeyConfigRequest) (_result *ModifyMcpApiKeyConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyMcpApiKeyConfigResponse{}
	_body, _err := client.ModifyMcpApiKeyConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建策略
//
// @param tmpReq - ModifyMcpNetworkWhitelistRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyMcpNetworkWhitelistRulesResponse
func (client *Client) ModifyMcpNetworkWhitelistRulesWithOptions(tmpReq *ModifyMcpNetworkWhitelistRulesRequest, runtime *dara.RuntimeOptions) (_result *ModifyMcpNetworkWhitelistRulesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ModifyMcpNetworkWhitelistRulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Rules) {
		request.RulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Rules, dara.String("Rules"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DebugMode) {
		query["DebugMode"] = request.DebugMode
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.McpPolicyId) {
		body["McpPolicyId"] = request.McpPolicyId
	}

	if !dara.IsNil(request.RulesShrink) {
		body["Rules"] = request.RulesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyMcpNetworkWhitelistRules"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyMcpNetworkWhitelistRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建策略
//
// @param request - ModifyMcpNetworkWhitelistRulesRequest
//
// @return ModifyMcpNetworkWhitelistRulesResponse
func (client *Client) ModifyMcpNetworkWhitelistRules(request *ModifyMcpNetworkWhitelistRulesRequest) (_result *ModifyMcpNetworkWhitelistRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyMcpNetworkWhitelistRulesResponse{}
	_body, _err := client.ModifyMcpNetworkWhitelistRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改策略
//
// @param tmpReq - ModifyMcpPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyMcpPolicyResponse
func (client *Client) ModifyMcpPolicyWithOptions(tmpReq *ModifyMcpPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyMcpPolicyResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ModifyMcpPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DisplayConfig) {
		request.DisplayConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DisplayConfig, dara.String("DisplayConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.NetworkConfig) {
		request.NetworkConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NetworkConfig, dara.String("NetworkConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.NetworkConfigShrink) {
		query["NetworkConfig"] = request.NetworkConfigShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Concurrency) {
		body["Concurrency"] = request.Concurrency
	}

	if !dara.IsNil(request.DebugMode) {
		body["DebugMode"] = request.DebugMode
	}

	if !dara.IsNil(request.DesktopMaxRuntime) {
		body["DesktopMaxRuntime"] = request.DesktopMaxRuntime
	}

	if !dara.IsNil(request.DisplayConfigShrink) {
		body["DisplayConfig"] = request.DisplayConfigShrink
	}

	if !dara.IsNil(request.IdleTimeoutSwitch) {
		body["IdleTimeoutSwitch"] = request.IdleTimeoutSwitch
	}

	if !dara.IsNil(request.McpIdleTimeout) {
		body["McpIdleTimeout"] = request.McpIdleTimeout
	}

	if !dara.IsNil(request.PolicyId) {
		body["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.PolicyName) {
		body["PolicyName"] = request.PolicyName
	}

	if !dara.IsNil(request.UserIdleTimeout) {
		body["UserIdleTimeout"] = request.UserIdleTimeout
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyMcpPolicy"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyMcpPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改策略
//
// @param request - ModifyMcpPolicyRequest
//
// @return ModifyMcpPolicyResponse
func (client *Client) ModifyMcpPolicy(request *ModifyMcpPolicyRequest) (_result *ModifyMcpPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyMcpPolicyResponse{}
	_body, _err := client.ModifyMcpPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改用户权益
//
// @param request - ModifyUserBenefitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyUserBenefitResponse
func (client *Client) ModifyUserBenefitWithOptions(request *ModifyUserBenefitRequest, runtime *dara.RuntimeOptions) (_result *ModifyUserBenefitResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChangeScene) {
		query["ChangeScene"] = request.ChangeScene
	}

	if !dara.IsNil(request.NewBenefitLevel) {
		query["NewBenefitLevel"] = request.NewBenefitLevel
	}

	if !dara.IsNil(request.OrderBenefitParam) {
		query["OrderBenefitParam"] = request.OrderBenefitParam
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SessionBandwidth) {
		body["SessionBandwidth"] = request.SessionBandwidth
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyUserBenefit"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyUserBenefitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改用户权益
//
// @param request - ModifyUserBenefitRequest
//
// @return ModifyUserBenefitResponse
func (client *Client) ModifyUserBenefit(request *ModifyUserBenefitRequest) (_result *ModifyUserBenefitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyUserBenefitResponse{}
	_body, _err := client.ModifyUserBenefitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 执行playground代码
//
// @param request - RunPlaygroundCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunPlaygroundCodeResponse
func (client *Client) RunPlaygroundCodeWithOptions(request *RunPlaygroundCodeRequest, runtime *dara.RuntimeOptions) (_result *RunPlaygroundCodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CodeContent) {
		query["CodeContent"] = request.CodeContent
	}

	if !dara.IsNil(request.SceneType) {
		query["SceneType"] = request.SceneType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunPlaygroundCode"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunPlaygroundCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 执行playground代码
//
// @param request - RunPlaygroundCodeRequest
//
// @return RunPlaygroundCodeResponse
func (client *Client) RunPlaygroundCode(request *RunPlaygroundCodeRequest) (_result *RunPlaygroundCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunPlaygroundCodeResponse{}
	_body, _err := client.RunPlaygroundCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 终止playground任务
//
// @param request - StopPlaygroundTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopPlaygroundTaskResponse
func (client *Client) StopPlaygroundTaskWithOptions(request *StopPlaygroundTaskRequest, runtime *dara.RuntimeOptions) (_result *StopPlaygroundTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopPlaygroundTask"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopPlaygroundTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 终止playground任务
//
// @param request - StopPlaygroundTaskRequest
//
// @return StopPlaygroundTaskResponse
func (client *Client) StopPlaygroundTask(request *StopPlaygroundTaskRequest) (_result *StopPlaygroundTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopPlaygroundTaskResponse{}
	_body, _err := client.StopPlaygroundTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 解除策略关联
//
// @param request - UnbindMcpPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindMcpPolicyResponse
func (client *Client) UnbindMcpPolicyWithOptions(request *UnbindMcpPolicyRequest, runtime *dara.RuntimeOptions) (_result *UnbindMcpPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiKeyId) {
		body["ApiKeyId"] = request.ApiKeyId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindMcpPolicy"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindMcpPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解除策略关联
//
// @param request - UnbindMcpPolicyRequest
//
// @return UnbindMcpPolicyResponse
func (client *Client) UnbindMcpPolicy(request *UnbindMcpPolicyRequest) (_result *UnbindMcpPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnbindMcpPolicyResponse{}
	_body, _err := client.UnbindMcpPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改池内云电脑最大创建数量
//
// @param request - UpdateConcurrencyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConcurrencyResponse
func (client *Client) UpdateConcurrencyWithOptions(request *UpdateConcurrencyRequest, runtime *dara.RuntimeOptions) (_result *UpdateConcurrencyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKeyId) {
		query["ApiKeyId"] = request.ApiKeyId
	}

	if !dara.IsNil(request.Concurrency) {
		query["Concurrency"] = request.Concurrency
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateConcurrency"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateConcurrencyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改池内云电脑最大创建数量
//
// @param request - UpdateConcurrencyRequest
//
// @return UpdateConcurrencyResponse
func (client *Client) UpdateConcurrency(request *UpdateConcurrencyRequest) (_result *UpdateConcurrencyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateConcurrencyResponse{}
	_body, _err := client.UpdateConcurrencyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改池内云电脑最大创建数量
//
// @param request - UpdateDesktopReleaseTimingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDesktopReleaseTimingsResponse
func (client *Client) UpdateDesktopReleaseTimingsWithOptions(request *UpdateDesktopReleaseTimingsRequest, runtime *dara.RuntimeOptions) (_result *UpdateDesktopReleaseTimingsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKeyId) {
		query["ApiKeyId"] = request.ApiKeyId
	}

	if !dara.IsNil(request.DesktopMaxRuntime) {
		query["DesktopMaxRuntime"] = request.DesktopMaxRuntime
	}

	if !dara.IsNil(request.McpIdleTimeout) {
		query["McpIdleTimeout"] = request.McpIdleTimeout
	}

	if !dara.IsNil(request.UserIdleTimeout) {
		query["UserIdleTimeout"] = request.UserIdleTimeout
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDesktopReleaseTimings"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDesktopReleaseTimingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改池内云电脑最大创建数量
//
// @param request - UpdateDesktopReleaseTimingsRequest
//
// @return UpdateDesktopReleaseTimingsResponse
func (client *Client) UpdateDesktopReleaseTimings(request *UpdateDesktopReleaseTimingsRequest) (_result *UpdateDesktopReleaseTimingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDesktopReleaseTimingsResponse{}
	_body, _err := client.UpdateDesktopReleaseTimingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新Mcp镜像
//
// @param request - UpdateImagePropertiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateImagePropertiesResponse
func (client *Client) UpdateImagePropertiesWithOptions(request *UpdateImagePropertiesRequest, runtime *dara.RuntimeOptions) (_result *UpdateImagePropertiesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKeyId) {
		query["ApiKeyId"] = request.ApiKeyId
	}

	if !dara.IsNil(request.ImageProperties) {
		query["ImageProperties"] = request.ImageProperties
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateImageProperties"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateImagePropertiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新Mcp镜像
//
// @param request - UpdateImagePropertiesRequest
//
// @return UpdateImagePropertiesResponse
func (client *Client) UpdateImageProperties(request *UpdateImagePropertiesRequest) (_result *UpdateImagePropertiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateImagePropertiesResponse{}
	_body, _err := client.UpdateImagePropertiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新Mcp镜像
//
// @param request - UpdateMcpImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMcpImageResponse
func (client *Client) UpdateMcpImageWithOptions(request *UpdateMcpImageRequest, runtime *dara.RuntimeOptions) (_result *UpdateMcpImageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKeyId) {
		query["ApiKeyId"] = request.ApiKeyId
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMcpImage"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMcpImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新Mcp镜像
//
// @param request - UpdateMcpImageRequest
//
// @return UpdateMcpImageResponse
func (client *Client) UpdateMcpImage(request *UpdateMcpImageRequest) (_result *UpdateMcpImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateMcpImageResponse{}
	_body, _err := client.UpdateMcpImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更改关联策略
//
// @param request - UpdateMcpPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMcpPolicyResponse
func (client *Client) UpdateMcpPolicyWithOptions(request *UpdateMcpPolicyRequest, runtime *dara.RuntimeOptions) (_result *UpdateMcpPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiKeyId) {
		body["ApiKeyId"] = request.ApiKeyId
	}

	if !dara.IsNil(request.PolicyId) {
		body["PolicyId"] = request.PolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMcpPolicy"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMcpPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更改关联策略
//
// @param request - UpdateMcpPolicyRequest
//
// @return UpdateMcpPolicyResponse
func (client *Client) UpdateMcpPolicy(request *UpdateMcpPolicyRequest) (_result *UpdateMcpPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateMcpPolicyResponse{}
	_body, _err := client.UpdateMcpPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建交付组
//
// @param request - UpdateResourceGroupSessionBandwidthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceGroupSessionBandwidthResponse
func (client *Client) UpdateResourceGroupSessionBandwidthWithOptions(request *UpdateResourceGroupSessionBandwidthRequest, runtime *dara.RuntimeOptions) (_result *UpdateResourceGroupSessionBandwidthResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SessionBandwidth) {
		body["SessionBandwidth"] = request.SessionBandwidth
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateResourceGroupSessionBandwidth"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateResourceGroupSessionBandwidthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建交付组
//
// @param request - UpdateResourceGroupSessionBandwidthRequest
//
// @return UpdateResourceGroupSessionBandwidthResponse
func (client *Client) UpdateResourceGroupSessionBandwidth(request *UpdateResourceGroupSessionBandwidthRequest) (_result *UpdateResourceGroupSessionBandwidthResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateResourceGroupSessionBandwidthResponse{}
	_body, _err := client.UpdateResourceGroupSessionBandwidthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
