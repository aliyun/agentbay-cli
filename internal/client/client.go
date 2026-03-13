// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strconv"
	"strings"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// ErrWithRequestID wraps an error and attaches a backend request ID for debugging (e.g. when -v is used).
type ErrWithRequestID struct {
	Err       error
	RequestID string
}

func (e *ErrWithRequestID) Error() string { return e.Err.Error() }
func (e *ErrWithRequestID) Unwrap() error { return e.Err }

// extractRequestIDFromResponse gets RequestId from CallApi response map (headers or XML body).
func extractRequestIDFromResponse(res map[string]interface{}) string {
	// Prefer header x-acs-request-id
	if h, ok := res["headers"].(map[string]interface{}); ok {
		for k, v := range h {
			if strings.EqualFold(k, "x-acs-request-id") {
				if s, ok := v.(string); ok && s != "" {
					return s
				}
			}
		}
	}
	// Else try to parse body as XML for RequestId
	if b, ok := res["body"]; ok && b != nil {
		var bodyStr string
		switch v := b.(type) {
		case string:
			bodyStr = v
		case []byte:
			bodyStr = string(v)
		default:
			return ""
		}
		if bodyStr == "" {
			return ""
		}
		var v struct {
			RequestId string `xml:"RequestId"`
		}
		if err := xml.Unmarshal([]byte(bodyStr), &v); err == nil && v.RequestId != "" {
			return v.RequestId
		}
	}
	return ""
}

// getMarketSkillCredentialResponseXML is used only for XML unmarshaling; Aliyun uses root "GetMarketSkillCredentialResponse".
type getMarketSkillCredentialResponseXML struct {
	XMLName        xml.Name                         `xml:"GetMarketSkillCredentialResponse"`
	Code           *string                          `xml:"Code"`
	Data           *getMarketSkillCredentialDataXML `xml:"Data"`
	HttpStatusCode *int32                           `xml:"HttpStatusCode"`
	Message        *string                          `xml:"Message"`
	RequestId      *string                          `xml:"RequestId"`
	Success        *bool                            `xml:"Success"`
}
type getMarketSkillCredentialDataXML struct {
	OssUrl      *string `xml:"OssUrl"`
	Url         *string `xml:"Url"`
	OssBucket   *string `xml:"OssBucket"`
	OssFilePath *string `xml:"OssFilePath"`
}

// parseGetMarketSkillCredentialResponse builds GetMarketSkillCredentialResponse from CallApi map (bodyType "string").
// Backend may return XML (pre-release) or JSON; we parse body manually like ListMarketGroupSkill.
func parseGetMarketSkillCredentialResponse(res map[string]interface{}) (*GetMarketSkillCredentialResponse, error) {
	out := &GetMarketSkillCredentialResponse{}
	bodyStr := ""
	switch v := res["body"].(type) {
	case string:
		bodyStr = v
	case []byte:
		bodyStr = string(v)
	default:
		return nil, &ErrWithRequestID{Err: errors.New("missing or invalid body in response"), RequestID: extractRequestIDFromResponse(res)}
	}
	parsed := &GetMarketSkillCredentialResponseBody{}
	if bodyStr != "" {
		trimmed := strings.TrimSpace(bodyStr)
		if len(trimmed) > 0 && trimmed[0] == '<' {
			var xmlResp getMarketSkillCredentialResponseXML
			if err := xml.Unmarshal([]byte(bodyStr), &xmlResp); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}
			parsed.Code = xmlResp.Code
			parsed.HttpStatusCode = xmlResp.HttpStatusCode
			parsed.Message = xmlResp.Message
			parsed.RequestId = xmlResp.RequestId
			parsed.Success = xmlResp.Success
			if xmlResp.Data != nil {
				d := xmlResp.Data
				parsed.Data = &GetMarketSkillCredentialResponseBodyData{
					OssUrl:      d.OssUrl,
					Url:         d.Url,
					OssBucket:   d.OssBucket,
					OssFilePath: d.OssFilePath,
				}
			}
		} else {
			if err := json.Unmarshal([]byte(bodyStr), parsed); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}
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

// listMarketGroupSkillResponseXML is used only for XML unmarshaling; backend returns root ListMarketGroupSkillResponse with <Data><data>...</data><data>...</data></Data>.
type listMarketGroupSkillResponseXML struct {
	XMLName        xml.Name                                  `xml:"ListMarketGroupSkillResponse"`
	HttpStatusCode *int32                                    `xml:"HttpStatusCode"`
	Data           []listMarketGroupSkillResponseDataItemXML `xml:"Data>data"`
	RequestId      *string                                   `xml:"RequestId"`
	Code           *string                                   `xml:"Code"`
	Success        *bool                                     `xml:"Success"`
}

type listMarketGroupSkillResponseDataItemXML struct {
	GroupName *string `xml:"GroupName"`
	GroupId   *string `xml:"GroupId"`
}

// parseListMarketGroupSkillResponse builds ListMarketGroupSkillResponse from CallApi map (bodyType "string").
func parseListMarketGroupSkillResponse(res map[string]interface{}) (*ListMarketGroupSkillResponse, error) {
	out := &ListMarketGroupSkillResponse{}
	bodyStr := ""
	switch v := res["body"].(type) {
	case string:
		bodyStr = v
	case []byte:
		bodyStr = string(v)
	default:
		return nil, &ErrWithRequestID{Err: errors.New("missing or invalid body in response"), RequestID: extractRequestIDFromResponse(res)}
	}
	parsed := &ListMarketGroupSkillResponseBody{}
	if bodyStr != "" {
		trimmed := strings.TrimSpace(bodyStr)
		if len(trimmed) > 0 && trimmed[0] == '<' {
			var xmlResp listMarketGroupSkillResponseXML
			if err := xml.Unmarshal([]byte(bodyStr), &xmlResp); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}
			parsed.Code = xmlResp.Code
			parsed.RequestId = xmlResp.RequestId
			parsed.Success = xmlResp.Success
			if len(xmlResp.Data) > 0 {
				parsed.Data = make([]ListMarketGroupSkillResponseBodyDataItem, len(xmlResp.Data))
				for i, d := range xmlResp.Data {
					parsed.Data[i] = ListMarketGroupSkillResponseBodyDataItem{GroupName: d.GroupName, GroupId: d.GroupId}
				}
			}
		} else {
			if err := json.Unmarshal([]byte(bodyStr), parsed); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}
		}
	}
	out.Body = &ListMarketGroupSkillResponseBodyWrapper{ListMarketGroupSkillResponseBody: parsed}
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

// createMarketSkillGroupResponseXML is used only for XML unmarshaling; backend returns <Data>group-id</Data> (chardata) or <Data><GroupId>...</GroupId></Data>.
type createMarketSkillGroupResponseXML struct {
	XMLName        xml.Name                      `xml:"CreateMarketSkillGroupResponse"`
	HttpStatusCode *int32                        `xml:"HttpStatusCode"`
	Data           createMarketSkillGroupDataXML `xml:"Data"`
	RequestId      *string                       `xml:"RequestId"`
	Code           *string                       `xml:"Code"`
	Success        *bool                         `xml:"Success"`
}
type createMarketSkillGroupDataXML struct {
	Value   string  `xml:",chardata"`
	GroupId *string `xml:"GroupId"`
}

// parseCreateMarketSkillGroupResponse builds CreateMarketSkillGroupResponse from CallApi map (bodyType "string").
// Backend may return XML (pre-release) or JSON; we parse body manually like ListMarketGroupSkill.
// JSON response may have Data as a string (group-id) or as an object {"GroupId": "..."}; we support both.
func parseCreateMarketSkillGroupResponse(res map[string]interface{}) (*CreateMarketSkillGroupResponse, error) {
	out := &CreateMarketSkillGroupResponse{}
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
	parsed := &CreateMarketSkillGroupResponseBody{}
	if bodyStr != "" {
		trimmed := strings.TrimSpace(bodyStr)
		if len(trimmed) > 0 && trimmed[0] == '<' {
			var xmlResp createMarketSkillGroupResponseXML
			if err := xml.Unmarshal([]byte(bodyStr), &xmlResp); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}
			parsed.Code = xmlResp.Code
			parsed.RequestId = xmlResp.RequestId
			parsed.Success = xmlResp.Success
			if xmlResp.Data.GroupId != nil && *xmlResp.Data.GroupId != "" {
				parsed.Data = &CreateMarketSkillGroupResponseBodyData{GroupId: xmlResp.Data.GroupId}
			} else if s := strings.TrimSpace(xmlResp.Data.Value); s != "" {
				parsed.Data = &CreateMarketSkillGroupResponseBodyData{GroupId: &s}
			}
		} else {
			// JSON: backend may return "Data": "group-id-string" or "Data": {"GroupId": "..."}
			var raw struct {
				Code      *string     `json:"Code,omitempty"`
				Data      interface{} `json:"Data,omitempty"`
				RequestId *string     `json:"RequestId,omitempty"`
				Success   *bool       `json:"Success,omitempty"`
			}
			if err := json.Unmarshal([]byte(bodyStr), &raw); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}
			parsed.Code = raw.Code
			parsed.RequestId = raw.RequestId
			parsed.Success = raw.Success
			if raw.Data != nil {
				switch v := raw.Data.(type) {
				case string:
					parsed.Data = &CreateMarketSkillGroupResponseBodyData{GroupId: &v}
				default:
					var dataObj CreateMarketSkillGroupResponseBodyData
					b, _ := json.Marshal(raw.Data)
					if err := json.Unmarshal(b, &dataObj); err == nil {
						parsed.Data = &dataObj
					}
				}
			}
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
	if !dara.IsNil(request.FilePath) {
		query["FilePath"] = request.FilePath
	}
	if !dara.IsNil(request.IsDockerfile) {
		query["IsDockerfile"] = request.IsDockerfile
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

// GetMarketSkillCredential 获取 Skill 上传凭证（OSS）
// Uses BodyType "string" so the SDK returns raw body; backend may return XML, so we parse body manually (same as ListMarketGroupSkill).
func (client *Client) GetMarketSkillCredentialWithOptions(request *GetMarketSkillCredentialRequest, runtime *dara.RuntimeOptions) (_result *GetMarketSkillCredentialResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}
	req := &openapiutil.OpenApiRequest{
		Query:   openapiutil.Query(query),
		Headers: map[string]*string{"Accept": dara.String("application/json")},
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMarketSkillCredential"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("string"),
	}
	_result = &GetMarketSkillCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		reqID := ""
		if _body != nil {
			reqID = extractRequestIDFromResponse(_body)
		}
		return _result, &ErrWithRequestID{Err: _err, RequestID: reqID}
	}
	_result, _err = parseGetMarketSkillCredentialResponse(_body)
	return _result, _err
}

func (client *Client) GetMarketSkillCredential(request *GetMarketSkillCredentialRequest) (_result *GetMarketSkillCredentialResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMarketSkillCredentialResponse{}
	_body, _err := client.GetMarketSkillCredentialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// createMarketSkillResponseXML is used only for XML unmarshaling; backend returns <CreateMarketSkillResponse><Data>skill-id</Data>...</CreateMarketSkillResponse>.
type createMarketSkillResponseXML struct {
	XMLName        xml.Name `xml:"CreateMarketSkillResponse"`
	HttpStatusCode *int32   `xml:"HttpStatusCode"`
	Data           string   `xml:"Data"`
	RequestId      *string  `xml:"RequestId"`
	Code           *string  `xml:"Code"`
	Success        *bool    `xml:"Success"`
}

// parseCreateMarketSkillResponse builds CreateMarketSkillResponse from CallApi map (bodyType "string").
// Backend may return XML (pre-release) or JSON; we parse body manually like CreateMarketSkillGroup.
func parseCreateMarketSkillResponse(res map[string]interface{}) (*CreateMarketSkillResponse, error) {
	out := &CreateMarketSkillResponse{}
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
	parsed := &CreateMarketSkillResponseBody{}
	if bodyStr != "" {
		trimmed := strings.TrimSpace(bodyStr)
		if len(trimmed) > 0 && trimmed[0] == '<' {
			var xmlResp createMarketSkillResponseXML
			if err := xml.Unmarshal([]byte(bodyStr), &xmlResp); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}
			parsed.Code = xmlResp.Code
			parsed.HttpStatusCode = xmlResp.HttpStatusCode
			parsed.RequestId = xmlResp.RequestId
			parsed.Success = xmlResp.Success
			if s := strings.TrimSpace(xmlResp.Data); s != "" {
				parsed.Data = &CreateMarketSkillResponseBodyData{SkillId: &s}
			}
		} else {
			if err := json.Unmarshal([]byte(bodyStr), parsed); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}
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

// CreateMarketSkill 通过 OSS 创建 Skill
// Uses BodyType "string" so we parse XML/JSON manually (backend pre-release returns XML).
func (client *Client) CreateMarketSkillWithOptions(request *CreateMarketSkillRequest, runtime *dara.RuntimeOptions) (_result *CreateMarketSkillResponse, _err error) {
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
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("string"),
	}
	_result = &CreateMarketSkillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		reqID := ""
		if _body != nil {
			reqID = extractRequestIDFromResponse(_body)
		}
		return _result, &ErrWithRequestID{Err: _err, RequestID: reqID}
	}
	_result, _err = parseCreateMarketSkillResponse(_body)
	return _result, _err
}

func (client *Client) CreateMarketSkill(request *CreateMarketSkillRequest) (_result *CreateMarketSkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMarketSkillResponse{}
	_body, _err := client.CreateMarketSkillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// describeMarketSkillDetailResponseXML is used only for XML unmarshaling; backend may return XML.
// Backend may use SkillId or SkillID in Data.
type describeMarketSkillDetailResponseXML struct {
	XMLName        xml.Name `xml:"DescribeMarketSkillDetailResponse"`
	HttpStatusCode *int32   `xml:"HttpStatusCode"`
	Data           *struct {
		SkillId     *string `xml:"SkillId"`
		SkillID     *string `xml:"SkillID"`
		Name        *string `xml:"Name"`
		Description *string `xml:"Description"`
	} `xml:"Data"`
	RequestId *string `xml:"RequestId"`
	Code      *string `xml:"Code"`
	Success   *bool   `xml:"Success"`
}

// parseDescribeMarketSkillDetailResponse builds DescribeMarketSkillDetailResponse from CallApi map (bodyType "string").
func parseDescribeMarketSkillDetailResponse(res map[string]interface{}) (*DescribeMarketSkillDetailResponse, error) {
	out := &DescribeMarketSkillDetailResponse{}
	bodyStr := ""
	switch v := res["body"].(type) {
	case string:
		bodyStr = v
	case []byte:
		bodyStr = string(v)
	default:
		return nil, &ErrWithRequestID{Err: errors.New("missing or invalid body in response"), RequestID: extractRequestIDFromResponse(res)}
	}
	parsed := &DescribeMarketSkillDetailResponseBody{}
	if bodyStr != "" {
		trimmed := strings.TrimSpace(bodyStr)
		if len(trimmed) > 0 && trimmed[0] == '<' {
			var xmlResp describeMarketSkillDetailResponseXML
			if err := xml.Unmarshal([]byte(bodyStr), &xmlResp); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}
			parsed.Code = xmlResp.Code
			parsed.HttpStatusCode = xmlResp.HttpStatusCode
			parsed.RequestId = xmlResp.RequestId
			parsed.Success = xmlResp.Success
			if xmlResp.Data != nil {
				skillIdVal := xmlResp.Data.SkillId
				if skillIdVal == nil {
					skillIdVal = xmlResp.Data.SkillID
				}
				parsed.Data = &DescribeMarketSkillDetailResponseBodyData{
					SkillId:     skillIdVal,
					Name:        xmlResp.Data.Name,
					Description: xmlResp.Data.Description,
				}
			}
		} else {
			if err := json.Unmarshal([]byte(bodyStr), parsed); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}
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

// DescribeMarketSkillDetail 查询 Skill 详情
// Uses BodyType "string" so we parse XML/JSON manually (backend pre-release returns XML).
func (client *Client) DescribeMarketSkillDetailWithOptions(request *DescribeMarketSkillDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeMarketSkillDetailResponse, _err error) {
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

func (client *Client) DescribeMarketSkillDetail(request *DescribeMarketSkillDetailRequest) (_result *DescribeMarketSkillDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMarketSkillDetailResponse{}
	_body, _err := client.DescribeMarketSkillDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// CreateMarketSkillGroup 创建技能组
// Uses BodyType "string" so the SDK returns raw body; backend may return XML, so we parse body manually (same as ListMarketGroupSkill).
func (client *Client) CreateMarketSkillGroupWithOptions(request *CreateMarketSkillGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateMarketSkillGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}
	req := &openapiutil.OpenApiRequest{
		Query:   openapiutil.Query(query),
		Headers: map[string]*string{"Accept": dara.String("application/json")},
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMarketSkillGroup"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("string"),
	}
	_result = &CreateMarketSkillGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		reqID := ""
		if _body != nil {
			reqID = extractRequestIDFromResponse(_body)
		}
		return _result, &ErrWithRequestID{Err: _err, RequestID: reqID}
	}
	_result, _err = parseCreateMarketSkillGroupResponse(_body)
	return _result, _err
}

func (client *Client) CreateMarketSkillGroup(request *CreateMarketSkillGroupRequest) (_result *CreateMarketSkillGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMarketSkillGroupResponse{}
	_body, _err := client.CreateMarketSkillGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// ListMarketGroupSkill 列出技能组
// Uses BodyType "string" so the SDK returns raw body; backend may return XML, so we parse body manually.
func (client *Client) ListMarketGroupSkillWithOptions(request *ListMarketGroupSkillRequest, runtime *dara.RuntimeOptions) (_result *ListMarketGroupSkillResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	req := &openapiutil.OpenApiRequest{
		Query:   openapiutil.Query(query),
		Headers: map[string]*string{"Accept": dara.String("application/json")},
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMarketGroupSkill"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("string"),
	}
	_result = &ListMarketGroupSkillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		reqID := ""
		if _body != nil {
			reqID = extractRequestIDFromResponse(_body)
		}
		return _result, &ErrWithRequestID{Err: _err, RequestID: reqID}
	}
	_result, _err = parseListMarketGroupSkillResponse(_body)
	return _result, _err
}

func (client *Client) ListMarketGroupSkill(request *ListMarketGroupSkillRequest) (_result *ListMarketGroupSkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListMarketGroupSkillResponse{}
	_body, _err := client.ListMarketGroupSkillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// addMarketGroupSkillResponseXML is used only for XML unmarshaling; backend may return XML.
type addMarketGroupSkillResponseXML struct {
	XMLName   xml.Name `xml:"AddMarketGroupSkillResponse"`
	Code      *string  `xml:"Code"`
	Data      *bool    `xml:"Data"`
	RequestId *string  `xml:"RequestId"`
	Success   *bool    `xml:"Success"`
}

// parseAddMarketGroupSkillResponse builds AddMarketGroupSkillResponse from CallApi map (bodyType "string").
func parseAddMarketGroupSkillResponse(res map[string]interface{}) (*AddMarketGroupSkillResponse, error) {
	out := &AddMarketGroupSkillResponse{}
	bodyStr := ""
	switch v := res["body"].(type) {
	case string:
		bodyStr = v
	case []byte:
		bodyStr = string(v)
	default:
		return nil, &ErrWithRequestID{Err: errors.New("missing or invalid body in response"), RequestID: extractRequestIDFromResponse(res)}
	}
	parsed := &AddMarketGroupSkillResponseBody{}
	if bodyStr != "" {
		trimmed := strings.TrimSpace(bodyStr)
		if len(trimmed) > 0 && trimmed[0] == '<' {
			var xmlResp addMarketGroupSkillResponseXML
			if err := xml.Unmarshal([]byte(bodyStr), &xmlResp); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}
			parsed.Code = xmlResp.Code
			parsed.Data = xmlResp.Data
			parsed.RequestId = xmlResp.RequestId
			parsed.Success = xmlResp.Success
		} else {
			if err := json.Unmarshal([]byte(bodyStr), parsed); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}
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

// AddMarketGroupSkill 组内添加技能
// Uses BodyType "string" so we parse XML/JSON manually (backend pre-release returns XML).
func (client *Client) AddMarketGroupSkillWithOptions(request *AddMarketGroupSkillRequest, runtime *dara.RuntimeOptions) (_result *AddMarketGroupSkillResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}
	if !dara.IsNil(request.SkillId) {
		query["SkillId"] = request.SkillId
	}
	req := &openapiutil.OpenApiRequest{
		Query:   openapiutil.Query(query),
		Headers: map[string]*string{"Accept": dara.String("application/json")},
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddMarketGroupSkill"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("string"),
	}
	_result = &AddMarketGroupSkillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		reqID := ""
		if _body != nil {
			reqID = extractRequestIDFromResponse(_body)
		}
		return _result, &ErrWithRequestID{Err: _err, RequestID: reqID}
	}
	_result, _err = parseAddMarketGroupSkillResponse(_body)
	return _result, _err
}

func (client *Client) AddMarketGroupSkill(request *AddMarketGroupSkillRequest) (_result *AddMarketGroupSkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddMarketGroupSkillResponse{}
	_body, _err := client.AddMarketGroupSkillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// removeMarketGroupSkillResponseXML is used only for XML unmarshaling; backend may return XML.
type removeMarketGroupSkillResponseXML struct {
	XMLName   xml.Name `xml:"RemoveMarketGroupSkillResponse"`
	Code      *string  `xml:"Code"`
	Data      *bool    `xml:"Data"`
	RequestId *string  `xml:"RequestId"`
	Success   *bool    `xml:"Success"`
}

// parseRemoveMarketGroupSkillResponse builds RemoveMarketGroupSkillResponse from CallApi map (bodyType "string").
func parseRemoveMarketGroupSkillResponse(res map[string]interface{}) (*RemoveMarketGroupSkillResponse, error) {
	out := &RemoveMarketGroupSkillResponse{}
	bodyStr := ""
	switch v := res["body"].(type) {
	case string:
		bodyStr = v
	case []byte:
		bodyStr = string(v)
	default:
		return nil, &ErrWithRequestID{Err: errors.New("missing or invalid body in response"), RequestID: extractRequestIDFromResponse(res)}
	}
	parsed := &RemoveMarketGroupSkillResponseBody{}
	if bodyStr != "" {
		trimmed := strings.TrimSpace(bodyStr)
		if len(trimmed) > 0 && trimmed[0] == '<' {
			var xmlResp removeMarketGroupSkillResponseXML
			if err := xml.Unmarshal([]byte(bodyStr), &xmlResp); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}
			parsed.Code = xmlResp.Code
			parsed.Data = xmlResp.Data
			parsed.RequestId = xmlResp.RequestId
			parsed.Success = xmlResp.Success
		} else {
			if err := json.Unmarshal([]byte(bodyStr), parsed); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}
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

// RemoveMarketGroupSkill 组内移除技能
// Uses BodyType "string" so we parse XML/JSON manually (backend pre-release returns XML).
func (client *Client) RemoveMarketGroupSkillWithOptions(request *RemoveMarketGroupSkillRequest, runtime *dara.RuntimeOptions) (_result *RemoveMarketGroupSkillResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}
	if !dara.IsNil(request.SkillId) {
		query["SkillId"] = request.SkillId
	}
	req := &openapiutil.OpenApiRequest{
		Query:   openapiutil.Query(query),
		Headers: map[string]*string{"Accept": dara.String("application/json")},
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveMarketGroupSkill"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("string"),
	}
	_result = &RemoveMarketGroupSkillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		reqID := ""
		if _body != nil {
			reqID = extractRequestIDFromResponse(_body)
		}
		return _result, &ErrWithRequestID{Err: _err, RequestID: reqID}
	}
	_result, _err = parseRemoveMarketGroupSkillResponse(_body)
	return _result, _err
}

func (client *Client) RemoveMarketGroupSkill(request *RemoveMarketGroupSkillRequest) (_result *RemoveMarketGroupSkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveMarketGroupSkillResponse{}
	_body, _err := client.RemoveMarketGroupSkillWithOptions(request, runtime)
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
	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
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
// 下载dockerfile模版
//
// @param request - GetDockerfileTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDockerfileTemplateResponse
func (client *Client) GetDockerfileTemplateWithOptions(request *GetDockerfileTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetDockerfileTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.SourceImageId) {
		query["SourceImageId"] = request.SourceImageId
	}

	if !dara.IsNil(request.Template) {
		query["Template"] = request.Template
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDockerfileTemplate"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDockerfileTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 下载dockerfile模版
//
// @param request - GetDockerfileTemplateRequest
//
// @return GetDockerfileTemplateResponse
func (client *Client) GetDockerfileTemplate(request *GetDockerfileTemplateRequest) (_result *GetDockerfileTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDockerfileTemplateResponse{}
	_body, _err := client.GetDockerfileTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
