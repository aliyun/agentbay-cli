// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"

	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 获取dockerfile文件放置位置
//
// @param request - GetDockerFileStoreCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDockerFileStoreCredentialResponse
func (client *Client) GetDockerFileStoreCredentialWithContext(ctx context.Context, request *GetDockerFileStoreCredentialRequest, runtime *dara.RuntimeOptions) (_result *GetDockerFileStoreCredentialResponse, _err error) {
	return client.GetDockerFileStoreCredentialWithOptions(request, runtime)
}

// GetMarketSkillCredentialWithContext 获取 Skill 上传凭证（OSS）
func (client *Client) GetMarketSkillCredentialWithContext(ctx context.Context, request *GetMarketSkillCredentialRequest, runtime *dara.RuntimeOptions) (_result *GetMarketSkillCredentialResponse, _err error) {
	return client.GetMarketSkillCredentialWithOptions(request, runtime)
}

// CreateMarketSkillWithContext 通过 OSS 创建 Skill
// Uses BodyType "string" and parseCreateMarketSkillResponse (backend may return XML).
func (client *Client) CreateMarketSkillWithContext(ctx context.Context, request *CreateMarketSkillRequest, runtime *dara.RuntimeOptions) (_result *CreateMarketSkillResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}
	if !dara.IsNil(request.OssFilePath) {
		query["OssFilePath"] = request.OssFilePath
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Headers: map[string]*string{
			"Accept": dara.String("application/xml"),
		},
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMarketSkill"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("string"),
	}
	_result = &CreateMarketSkillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_result, _err = parseCreateMarketSkillResponse(_body)
	return _result, _err
}

// DescribeMarketSkillDetailWithContext 查询 Skill 详情
// Uses BodyType "string" and parseDescribeMarketSkillDetailResponse (same as DescribeMarketSkillDetailWithOptions).
func (client *Client) DescribeMarketSkillDetailWithContext(ctx context.Context, request *DescribeMarketSkillDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeMarketSkillDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SkillId) {
		query["SkillId"] = request.SkillId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Headers: map[string]*string{
			"Accept": dara.String("application/xml"),
		},
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMarketSkillDetail"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("string"),
	}
	_result = &DescribeMarketSkillDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		reqID := ""
		if _body != nil {
			reqID = extractRequestIDFromResponse(_body)
		}
		return _result, &ErrWithRequestID{Err: _err, RequestID: reqID}
	}
	_result, _err = parseDescribeMarketSkillDetailResponse(_body)
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
func (client *Client) CreateDockerImageTaskWithContext(ctx context.Context, request *CreateDockerImageTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateDockerImageTaskResponse, _err error) {
	return client.CreateDockerImageTaskWithOptions(request, runtime)
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
func (client *Client) GetDockerImageTaskWithContext(ctx context.Context, request *GetDockerImageTaskRequest, runtime *dara.RuntimeOptions) (_result *GetDockerImageTaskResponse, _err error) {
	return client.GetDockerImageTaskWithOptions(request, runtime)
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
func (client *Client) ListMcpImagesWithContext(ctx context.Context, request *ListMcpImagesRequest, runtime *dara.RuntimeOptions) (_result *ListMcpImagesResponse, _err error) {
	return client.ListMcpImagesWithOptions(request, runtime)
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
func (client *Client) GetMcpImageInfoWithContext(ctx context.Context, request *GetMcpImageInfoRequest, runtime *dara.RuntimeOptions) (_result *GetMcpImageInfoResponse, _err error) {
	return client.GetMcpImageInfoWithOptions(request, runtime)
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
func (client *Client) CreateResourceGroupWithContext(ctx context.Context, request *CreateResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateResourceGroupResponse, _err error) {
	return client.CreateResourceGroupWithOptions(request, runtime)
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
func (client *Client) DeleteResourceGroupWithContext(ctx context.Context, request *DeleteResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteResourceGroupResponse, _err error) {
	return client.DeleteResourceGroupWithOptions(request, runtime)
}
