// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"

	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// CreateSimpleOfficeSite 创建简单办公网络
func (client *Client) CreateSimpleOfficeSiteWithOptions(request *CreateSimpleOfficeSiteRequest, runtime *dara.RuntimeOptions) (_result *CreateSimpleOfficeSiteResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.VpcType) {
		body["VpcType"] = request.VpcType
	}
	if !dara.IsNil(request.OfficeSiteName) {
		body["OfficeSiteName"] = request.OfficeSiteName
	}
	if !dara.IsNil(request.VpcId) {
		body["VpcId"] = request.VpcId
	}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}
	if !dara.IsNil(request.RegionName) {
		body["RegionName"] = request.RegionName
	}
	if !dara.IsNil(request.DesktopAccessType) {
		body["DesktopAccessType"] = request.DesktopAccessType
	}

	req := &openapiutil.OpenApiRequest{
		Body:    openapiutil.ParseToMap(body),
		Headers: map[string]*string{"Accept": dara.String("application/json")},
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSimpleOfficeSite"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("string"),
	}
	_result = &CreateSimpleOfficeSiteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		reqID := ""
		if _body != nil {
			reqID = extractRequestIDFromResponse(_body)
		}
		return _result, &ErrWithRequestID{Err: _err, RequestID: reqID}
	}
	_result, _err = parseCreateSimpleOfficeSiteResponse(_body)
	return _result, _err
}

func (client *Client) CreateSimpleOfficeSite(request *CreateSimpleOfficeSiteRequest) (_result *CreateSimpleOfficeSiteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSimpleOfficeSiteResponse{}
	_body, _err := client.CreateSimpleOfficeSiteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSimpleOfficeSiteWithContext(ctx context.Context, request *CreateSimpleOfficeSiteRequest, runtime *dara.RuntimeOptions) (_result *CreateSimpleOfficeSiteResponse, _err error) {
	return client.CreateSimpleOfficeSiteWithOptions(request, runtime)
}
