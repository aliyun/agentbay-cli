// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// DescribeInstanceTypes 查询实例规格
func (client *Client) DescribeInstanceTypesWithOptions(request *DescribeInstanceTypesRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceTypesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	req := &openapiutil.OpenApiRequest{
		Query:   openapiutil.Query(query),
		Headers: map[string]*string{"Accept": dara.String("application/json")},
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceTypes"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("string"),
	}
	_result = &DescribeInstanceTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		reqID := ""
		if _body != nil {
			reqID = extractRequestIDFromResponse(_body)
		}
		return _result, &ErrWithRequestID{Err: _err, RequestID: reqID}
	}
	_result, _err = parseDescribeInstanceTypesResponse(_body)
	return _result, _err
}

func (client *Client) DescribeInstanceTypes(request *DescribeInstanceTypesRequest) (_result *DescribeInstanceTypesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceTypesResponse{}
	_body, _err := client.DescribeInstanceTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceTypesWithContext(ctx context.Context, request *DescribeInstanceTypesRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceTypesResponse, _err error) {
	return client.DescribeInstanceTypesWithOptions(request, runtime)
}

// parseDescribeInstanceTypesResponse builds DescribeInstanceTypesResponse from CallApi map (bodyType "string").
func parseDescribeInstanceTypesResponse(res map[string]interface{}) (*DescribeInstanceTypesResponse, error) {
	out := &DescribeInstanceTypesResponse{}
	bodyStr := ""
	switch v := res["body"].(type) {
	case string:
		bodyStr = v
	case []byte:
		bodyStr = string(v)
	default:
		return nil, &ErrWithRequestID{Err: errors.New("missing or invalid body in response"), RequestID: extractRequestIDFromResponse(res)}
	}
	out.RawBody = bodyStr
	parsed := &DescribeInstanceTypesResponseBody{}
	if bodyStr != "" {
		if err := json.Unmarshal([]byte(bodyStr), parsed); err != nil {
			return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
		}
	}
	out.Body = parsed
	if h, ok := res["headers"].(map[string]*string); ok {
		out.Headers = h
	} else if h, ok := res["headers"].(map[string]interface{}); ok {
		out.Headers = make(map[string]*string)
		for k, v := range h {
			if s, ok := v.(string); ok {
				out.Headers[k] = dara.String(s)
			} else if p, ok := v.(*string); ok && p != nil {
				out.Headers[k] = p
			}
		}
	}
	if sc, ok := res["statusCode"].(int); ok {
		out.StatusCode = dara.Int32(int32(sc))
	}
	if sc, ok := res["statusCode"].(int32); ok {
		out.StatusCode = &sc
	}
	if out.StatusCode == nil && res["statusCode"] != nil {
		if n, err := strconv.Atoi(dara.ToString(res["statusCode"])); err == nil {
			out.StatusCode = dara.Int32(int32(n))
		}
	}
	return out, nil
}

// DescribeMcpPolicyData 查询策略配置数据
func (client *Client) DescribeMcpPolicyDataWithOptions(request *DescribeMcpPolicyDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeMcpPolicyDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	req := &openapiutil.OpenApiRequest{
		Query:   openapiutil.Query(query),
		Headers: map[string]*string{"Accept": dara.String("application/json")},
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMcpPolicyData"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("string"),
	}
	_result = &DescribeMcpPolicyDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		reqID := ""
		if _body != nil {
			reqID = extractRequestIDFromResponse(_body)
		}
		return _result, &ErrWithRequestID{Err: _err, RequestID: reqID}
	}
	_result, _err = parseDescribeMcpPolicyDataResponse(_body)
	return _result, _err
}

func (client *Client) DescribeMcpPolicyData(request *DescribeMcpPolicyDataRequest) (_result *DescribeMcpPolicyDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMcpPolicyDataResponse{}
	_body, _err := client.DescribeMcpPolicyDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMcpPolicyDataWithContext(ctx context.Context, request *DescribeMcpPolicyDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeMcpPolicyDataResponse, _err error) {
	return client.DescribeMcpPolicyDataWithOptions(request, runtime)
}

// parseDescribeMcpPolicyDataResponse builds DescribeMcpPolicyDataResponse from CallApi map (bodyType "string").
func parseDescribeMcpPolicyDataResponse(res map[string]interface{}) (*DescribeMcpPolicyDataResponse, error) {
	out := &DescribeMcpPolicyDataResponse{}
	bodyStr := ""
	switch v := res["body"].(type) {
	case string:
		bodyStr = v
	case []byte:
		bodyStr = string(v)
	default:
		return nil, &ErrWithRequestID{Err: errors.New("missing or invalid body in response"), RequestID: extractRequestIDFromResponse(res)}
	}
	out.RawBody = bodyStr
	parsed := &DescribeMcpPolicyDataResponseBody{}
	if bodyStr != "" {
		if err := json.Unmarshal([]byte(bodyStr), parsed); err != nil {
			return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
		}
	}
	out.Body = parsed
	if h, ok := res["headers"].(map[string]*string); ok {
		out.Headers = h
	} else if h, ok := res["headers"].(map[string]interface{}); ok {
		out.Headers = make(map[string]*string)
		for k, v := range h {
			if s, ok := v.(string); ok {
				out.Headers[k] = dara.String(s)
			} else if p, ok := v.(*string); ok && p != nil {
				out.Headers[k] = p
			}
		}
	}
	if sc, ok := res["statusCode"].(int); ok {
		out.StatusCode = dara.Int32(int32(sc))
	}
	if sc, ok := res["statusCode"].(int32); ok {
		out.StatusCode = &sc
	}
	if out.StatusCode == nil && res["statusCode"] != nil {
		if n, err := strconv.Atoi(dara.ToString(res["statusCode"])); err == nil {
			out.StatusCode = dara.Int32(int32(n))
		}
	}
	return out, nil
}

// SaveMcpPolicyData 保存策略配置数据
func (client *Client) SaveMcpPolicyDataWithOptions(request *SaveMcpPolicyDataRequest, runtime *dara.RuntimeOptions) (_result *SaveMcpPolicyDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}
	if !dara.IsNil(request.PolicyId) {
		body["PolicyId"] = request.PolicyId
	}
	// Helper function to marshal nested objects to JSON string
	marshalNested := func(v interface{}) string {
		b, _ := json.Marshal(v)
		return string(b)
	}
	if request.GroupSpec != nil {
		body["GroupSpec"] = marshalNested(map[string]interface{}{
			"AppInstanceType": request.GroupSpec.AppInstanceType,
			"RegionName":      request.GroupSpec.RegionName,
			"Memory":          request.GroupSpec.Memory,
			"Cpu":             request.GroupSpec.Cpu,
			"RegionId":        request.GroupSpec.RegionId,
		})
	}
	if request.SandboxLifeCycle != nil {
		body["SandboxLifeCycle"] = marshalNested(map[string]interface{}{
			"IdleTimeoutSwitch": request.SandboxLifeCycle.IdleTimeoutSwitch,
			"HibernateTimeout":  request.SandboxLifeCycle.HibernateTimeout,
			"DesktopMaxRuntime": request.SandboxLifeCycle.DesktopMaxRuntime,
			"UserIdleTimeout":   request.SandboxLifeCycle.UserIdleTimeout,
		})
	}
	if request.NetworkData != nil {
		body["NetworkData"] = marshalNested(map[string]interface{}{
			"VpcId":            request.NetworkData.VpcId,
			"OfficeSiteType":   request.NetworkData.OfficeSiteType,
			"DnsAddress":       request.NetworkData.DnsAddress,
			"VpcName":          request.NetworkData.VpcName,
			"SessionBandwidth": request.NetworkData.SessionBandwidth,
		})
	}
	if request.ScreenSettings != nil {
		body["ScreenSettings"] = marshalNested(map[string]interface{}{
			"ClientControlMenu": request.ScreenSettings.ClientControlMenu,
			"ScreenDisplayMode": request.ScreenSettings.ScreenDisplayMode,
			"Taskbar":           request.ScreenSettings.Taskbar,
			"KioskModeEnabled":  request.ScreenSettings.KioskModeEnabled,
		})
	}
	if request.NetworkConfig != nil {
		body["NetworkConfig"] = marshalNested(map[string]interface{}{
			"Enabled": request.NetworkConfig.Enabled,
		})
	}
	if request.DisplayConfig != nil {
		body["DisplayConfig"] = marshalNested(map[string]interface{}{
			"DisplayMode": request.DisplayConfig.DisplayMode,
		})
	}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body:    openapiutil.ParseToMap(body),
		Headers: map[string]*string{"Accept": dara.String("application/json")},
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveMcpPolicyData"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("string"),
	}
	_result = &SaveMcpPolicyDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		reqID := ""
		if _body != nil {
			reqID = extractRequestIDFromResponse(_body)
		}
		return _result, &ErrWithRequestID{Err: _err, RequestID: reqID}
	}
	_result, _err = parseSaveMcpPolicyDataResponse(_body)
	return _result, _err
}

func (client *Client) SaveMcpPolicyData(request *SaveMcpPolicyDataRequest) (_result *SaveMcpPolicyDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveMcpPolicyDataResponse{}
	_body, _err := client.SaveMcpPolicyDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveMcpPolicyDataWithContext(ctx context.Context, request *SaveMcpPolicyDataRequest, runtime *dara.RuntimeOptions) (_result *SaveMcpPolicyDataResponse, _err error) {
	return client.SaveMcpPolicyDataWithOptions(request, runtime)
}

// parseSaveMcpPolicyDataResponse builds SaveMcpPolicyDataResponse from CallApi map (bodyType "string").
func parseSaveMcpPolicyDataResponse(res map[string]interface{}) (*SaveMcpPolicyDataResponse, error) {
	out := &SaveMcpPolicyDataResponse{}
	bodyStr := ""
	switch v := res["body"].(type) {
	case string:
		bodyStr = v
	case []byte:
		bodyStr = string(v)
	default:
		return nil, &ErrWithRequestID{Err: errors.New("missing or invalid body in response"), RequestID: extractRequestIDFromResponse(res)}
	}
	out.RawBody = bodyStr
	parsed := &SaveMcpPolicyDataResponseBody{}
	if bodyStr != "" {
		trimmed := strings.TrimSpace(bodyStr)
		if len(trimmed) > 0 && trimmed[0] == '<' {
			// XML response - not expected for this API, but handle gracefully
			return nil, &ErrWithRequestID{Err: errors.New("unexpected XML response for SaveMcpPolicyData"), RequestID: extractRequestIDFromResponse(res)}
		}
		if err := json.Unmarshal([]byte(bodyStr), parsed); err != nil {
			return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
		}
	}
	out.Body = parsed
	if h, ok := res["headers"].(map[string]*string); ok {
		out.Headers = h
	} else if h, ok := res["headers"].(map[string]interface{}); ok {
		out.Headers = make(map[string]*string)
		for k, v := range h {
			if s, ok := v.(string); ok {
				out.Headers[k] = dara.String(s)
			} else if p, ok := v.(*string); ok && p != nil {
				out.Headers[k] = p
			}
		}
	}
	if sc, ok := res["statusCode"].(int); ok {
		out.StatusCode = dara.Int32(int32(sc))
	}
	if sc, ok := res["statusCode"].(int32); ok {
		out.StatusCode = &sc
	}
	if out.StatusCode == nil && res["statusCode"] != nil {
		if n, err := strconv.Atoi(dara.ToString(res["statusCode"])); err == nil {
			out.StatusCode = dara.Int32(int32(n))
		}
	}
	return out, nil
}

// DescribeOfficeSites 查询办公网络
func (client *Client) DescribeOfficeSitesWithOptions(request *DescribeOfficeSitesRequest, runtime *dara.RuntimeOptions) (_result *DescribeOfficeSitesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OfficeSiteType) {
		query["OfficeSiteType"] = request.OfficeSiteType
	}
	if !dara.IsNil(request.RegionName) {
		query["RegionName"] = request.RegionName
	}

	req := &openapiutil.OpenApiRequest{
		Query:   openapiutil.Query(query),
		Headers: map[string]*string{"Accept": dara.String("application/json")},
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOfficeSites"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("string"),
	}
	_result = &DescribeOfficeSitesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		reqID := ""
		if _body != nil {
			reqID = extractRequestIDFromResponse(_body)
		}
		return _result, &ErrWithRequestID{Err: _err, RequestID: reqID}
	}
	_result, _err = parseDescribeOfficeSitesResponse(_body)
	return _result, _err
}

func (client *Client) DescribeOfficeSites(request *DescribeOfficeSitesRequest) (_result *DescribeOfficeSitesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeOfficeSitesResponse{}
	_body, _err := client.DescribeOfficeSitesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOfficeSitesWithContext(ctx context.Context, request *DescribeOfficeSitesRequest, runtime *dara.RuntimeOptions) (_result *DescribeOfficeSitesResponse, _err error) {
	return client.DescribeOfficeSitesWithOptions(request, runtime)
}

// parseDescribeOfficeSitesResponse builds DescribeOfficeSitesResponse from CallApi map (bodyType "string").
func parseDescribeOfficeSitesResponse(res map[string]interface{}) (*DescribeOfficeSitesResponse, error) {
	out := &DescribeOfficeSitesResponse{}
	bodyStr := ""
	switch v := res["body"].(type) {
	case string:
		bodyStr = v
	case []byte:
		bodyStr = string(v)
	default:
		return nil, &ErrWithRequestID{Err: errors.New("missing or invalid body in response"), RequestID: extractRequestIDFromResponse(res)}
	}
	out.RawBody = bodyStr
	parsed := &DescribeOfficeSitesResponseBody{}
	if bodyStr != "" {
		if err := json.Unmarshal([]byte(bodyStr), parsed); err != nil {
			return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
		}
	}
	out.Body = parsed
	if h, ok := res["headers"].(map[string]*string); ok {
		out.Headers = h
	} else if h, ok := res["headers"].(map[string]interface{}); ok {
		out.Headers = make(map[string]*string)
		for k, v := range h {
			if s, ok := v.(string); ok {
				out.Headers[k] = dara.String(s)
			} else if p, ok := v.(*string); ok && p != nil {
				out.Headers[k] = p
			}
		}
	}
	if sc, ok := res["statusCode"].(int); ok {
		out.StatusCode = dara.Int32(int32(sc))
	}
	if sc, ok := res["statusCode"].(int32); ok {
		out.StatusCode = &sc
	}
	if out.StatusCode == nil && res["statusCode"] != nil {
		if n, err := strconv.Atoi(dara.ToString(res["statusCode"])); err == nil {
			out.StatusCode = dara.Int32(int32(n))
		}
	}
	return out, nil
}
