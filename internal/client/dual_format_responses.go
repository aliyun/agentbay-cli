// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/alibabacloud-go/tea/dara"
)

// --- shared helpers ---

func rawBodyStringFromMap(res map[string]interface{}) (string, error) {
	switch v := res["body"].(type) {
	case string:
		return v, nil
	case []byte:
		return string(v), nil
	default:
		return "", errors.New("missing or invalid body in response")
	}
}

func applyMapHeadersAndStatus(outHeaders *map[string]*string, outStatus **int32, res map[string]interface{}) {
	if *outHeaders == nil {
		*outHeaders = make(map[string]*string)
	}
	if h, ok := res["headers"].(map[string]*string); ok {
		for k, v := range h {
			(*outHeaders)[k] = v
		}
	} else if h, ok := res["headers"].(map[string]interface{}); ok {
		for k, v := range h {
			if s, ok := v.(string); ok {
				(*outHeaders)[k] = dara.String(s)
			} else if p, ok := v.(*string); ok && p != nil {
				(*outHeaders)[k] = p
			}
		}
	}
	if sc, ok := res["statusCode"].(int); ok {
		*outStatus = dara.Int32(int32(sc))
	}
	if sc, ok := res["statusCode"].(int32); ok {
		*outStatus = &sc
	}
	if *outStatus == nil && res["statusCode"] != nil {
		if n, err := strconv.Atoi(dara.ToString(res["statusCode"])); err == nil {
			*outStatus = dara.Int32(int32(n))
		}
	}
}

func mergeDerivedGetMcpImageInfoHeaders(out *GetMcpImageInfoResponse) {
	if out == nil || out.Body == nil || out.Body.Data == nil {
		return
	}
	if out.Headers == nil {
		out.Headers = make(map[string]*string)
	}
	d := out.Body.Data
	if d.ImageResourceStatus != nil && *d.ImageResourceStatus != "" {
		out.Headers["X-Image-Resource-Status"] = d.ImageResourceStatus
	}
	if d.ImageInfo != nil && d.ImageInfo.ImageType != nil && *d.ImageInfo.ImageType != "" {
		out.Headers["X-Image-Type"] = d.ImageInfo.ImageType
	}
}

// --- CreateDockerImageTask ---

type xmlCreateDockerImageTaskResponse struct {
	XMLName        xml.Name `xml:"CreateDockerImageTaskResponse"`
	RequestId      string   `xml:"RequestId"`
	HttpStatusCode int      `xml:"HttpStatusCode"`
	Data           struct {
		TaskId string `xml:"TaskId"`
	} `xml:"Data"`
	Code    string `xml:"Code"`
	Success bool   `xml:"Success"`
}

func parseCreateDockerImageTaskResponse(res map[string]interface{}) (*CreateDockerImageTaskResponse, error) {
	bodyStr, err := rawBodyStringFromMap(res)
	if err != nil {
		return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
	}
	out := &CreateDockerImageTaskResponse{Headers: make(map[string]*string)}
	trimmed := strings.TrimSpace(bodyStr)
	if len(trimmed) > 0 && trimmed[0] == '<' {
		var xr xmlCreateDockerImageTaskResponse
		if err := xml.Unmarshal([]byte(bodyStr), &xr); err != nil {
			return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
		}
		out.StatusCode = dara.Int32(int32(xr.HttpStatusCode))
		out.Body = &CreateDockerImageTaskResponseBody{
			RequestId:      dara.String(xr.RequestId),
			HttpStatusCode: dara.Int32(int32(xr.HttpStatusCode)),
			Code:           dara.String(xr.Code),
			Success:        dara.Bool(xr.Success),
			Data: &CreateDockerImageTaskResponseBodyData{
				TaskId: dara.String(xr.Data.TaskId),
			},
		}
	} else {
		parsed := &CreateDockerImageTaskResponseBody{}
		if bodyStr != "" {
			if err := json.Unmarshal([]byte(bodyStr), parsed); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}
		}
		out.Body = parsed
	}
	applyMapHeadersAndStatus(&out.Headers, &out.StatusCode, res)
	return out, nil
}

// --- GetDockerImageTask ---

type xmlGetDockerImageTaskResponse struct {
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

func parseGetDockerImageTaskResponse(res map[string]interface{}) (*GetDockerImageTaskResponse, error) {
	bodyStr, err := rawBodyStringFromMap(res)
	if err != nil {
		return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
	}
	out := &GetDockerImageTaskResponse{Headers: make(map[string]*string)}
	trimmed := strings.TrimSpace(bodyStr)
	if len(trimmed) > 0 && trimmed[0] == '<' {
		var xr xmlGetDockerImageTaskResponse
		if err := xml.Unmarshal([]byte(bodyStr), &xr); err != nil {
			return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
		}
		out.StatusCode = dara.Int32(int32(xr.HttpStatusCode))
		out.Body = &GetDockerImageTaskResponseBody{
			RequestId:      dara.String(xr.RequestId),
			HttpStatusCode: dara.Int32(int32(xr.HttpStatusCode)),
			Code:           dara.String(xr.Code),
			Success:        dara.Bool(xr.Success),
			Data: &GetDockerImageTaskResponseBodyData{
				Status:  dara.String(xr.Data.Status),
				ImageId: dara.String(xr.Data.ImageId),
				TaskMsg: dara.String(xr.Data.TaskMsg),
			},
		}
	} else {
		parsed := &GetDockerImageTaskResponseBody{}
		if bodyStr != "" {
			if err := json.Unmarshal([]byte(bodyStr), parsed); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}
		}
		out.Body = parsed
	}
	applyMapHeadersAndStatus(&out.Headers, &out.StatusCode, res)
	return out, nil
}

// --- CreateMarketSkill ---

type xmlCreateMarketSkillResponse struct {
	XMLName        xml.Name `xml:"CreateMarketSkillResponse"`
	HttpStatusCode *int32   `xml:"HttpStatusCode"`
	Data           string   `xml:"Data"`
	RequestId      *string  `xml:"RequestId"`
	Code           *string  `xml:"Code"`
	Success        *bool    `xml:"Success"`
}

// createMarketSkillJSONWire decodes JSON without forcing Data to be an object; some gateways return Data as the skill id string.
type createMarketSkillJSONWire struct {
	Code           *string         `json:"Code"`
	Data           json.RawMessage `json:"Data"`
	HttpStatusCode *int32          `json:"HttpStatusCode"`
	Message        *string         `json:"Message"`
	RequestId      *string         `json:"RequestId"`
	Success        *bool           `json:"Success"`
}

func parseCreateMarketSkillDataField(raw json.RawMessage) (*CreateMarketSkillResponseBodyData, error) {
	if len(raw) == 0 {
		return nil, nil
	}
	s := bytes.TrimSpace(raw)
	if len(s) == 0 || string(s) == "null" {
		return nil, nil
	}
	switch s[0] {
	case '"':
		var id string
		if err := json.Unmarshal(s, &id); err != nil {
			return nil, err
		}
		id = strings.TrimSpace(id)
		if id == "" {
			return nil, nil
		}
		return &CreateMarketSkillResponseBodyData{SkillId: &id}, nil
	case '{':
		var d CreateMarketSkillResponseBodyData
		if err := json.Unmarshal(s, &d); err != nil {
			return nil, err
		}
		return &d, nil
	default:
		return nil, fmt.Errorf("CreateMarketSkill response Data: unsupported JSON (expected string skill id or object)")
	}
}

// parseCreateMarketSkillResponse builds CreateMarketSkillResponse from CallApi map (bodyType "string").
// Backend may return XML or JSON; JSON may use Data as either a string (skill id) or an object {SkillId}.
func parseCreateMarketSkillResponse(res map[string]interface{}) (*CreateMarketSkillResponse, error) {
	bodyStr, err := rawBodyStringFromMap(res)
	if err != nil {
		return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
	}
	out := &CreateMarketSkillResponse{RawBody: bodyStr}
	parsed := &CreateMarketSkillResponseBody{}
	trimmed := strings.TrimSpace(bodyStr)
	if bodyStr != "" {
		if len(trimmed) > 0 && trimmed[0] == '<' {
			var xmlResp xmlCreateMarketSkillResponse
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
			var wire createMarketSkillJSONWire
			if err := json.Unmarshal([]byte(bodyStr), &wire); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}
			parsed.Code = wire.Code
			parsed.HttpStatusCode = wire.HttpStatusCode
			parsed.Message = wire.Message
			parsed.RequestId = wire.RequestId
			parsed.Success = wire.Success
			data, derr := parseCreateMarketSkillDataField(wire.Data)
			if derr != nil {
				return nil, &ErrWithRequestID{Err: derr, RequestID: extractRequestIDFromResponse(res)}
			}
			parsed.Data = data
		}
	}
	out.Body = parsed
	applyMapHeadersAndStatus(&out.Headers, &out.StatusCode, res)
	return out, nil
}

// --- ListMcpImages ---

type xmlListMcpImagesResponse struct {
	XMLName        xml.Name `xml:"ListMcpImagesResponse"`
	RequestId      string   `xml:"RequestId"`
	HttpStatusCode int      `xml:"HttpStatusCode"`
	Data           struct {
		Images []struct {
			ImageId                string `xml:"ImageId"`
			ImageName              string `xml:"ImageName"`
			ImageBuildType         string `xml:"ImageBuildType"`
			ImageIntro             string `xml:"ImageIntro"`
			ImageApplyScene        string `xml:"ImageApplyScene"`
			ImageResourceStatus    string `xml:"ImageResourceStatus"`
			ImageResourceGroupInfo struct {
				ResourceGroupId string `xml:"ResourceGroupId"`
			} `xml:"ImageResourceGroupInfo"`
			ImageInfo struct {
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

func parseListMcpImagesResponse(res map[string]interface{}) (*ListMcpImagesResponse, error) {
	bodyStr, err := rawBodyStringFromMap(res)
	if err != nil {
		return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
	}
	out := &ListMcpImagesResponse{Headers: make(map[string]*string)}
	trimmed := strings.TrimSpace(bodyStr)
	if len(trimmed) > 0 && trimmed[0] == '<' {
		var xr xmlListMcpImagesResponse
		if err := xml.Unmarshal([]byte(bodyStr), &xr); err != nil {
			return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
		}
		var sdkData []*ListMcpImagesResponseBodyData
		for _, xmlImage := range xr.Data.Images {
			var imageInfo *ListMcpImagesResponseBodyDataImageInfo
			if xmlImage.ImageInfo.OsName != "" || xmlImage.ImageInfo.OsVersion != "" {
				imageInfo = &ListMcpImagesResponseBodyDataImageInfo{
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
			var toolInfo []*ListMcpImagesResponseBodyDataToolInfo
			for _, xmlTool := range xmlImage.ToolInfo {
				var toolList []*ListMcpImagesResponseBodyDataToolInfoToolList
				for _, xmlToolItem := range xmlTool.ToolList {
					toolList = append(toolList, &ListMcpImagesResponseBodyDataToolInfoToolList{
						Tool:        dara.String(xmlToolItem.Tool),
						Description: dara.String(xmlToolItem.Description),
					})
				}
				toolInfo = append(toolInfo, &ListMcpImagesResponseBodyDataToolInfo{
					McpServerId:   dara.String(xmlTool.McpServerId),
					McpServerName: dara.String(xmlTool.McpServerName),
					ToolList:      toolList,
				})
			}
			var imageResourceGroupInfo *ListMcpImagesResponseBodyDataImageResourceGroupInfo
			if xmlImage.ImageResourceGroupInfo.ResourceGroupId != "" {
				imageResourceGroupInfo = &ListMcpImagesResponseBodyDataImageResourceGroupInfo{
					ResourceGroupId: dara.String(xmlImage.ImageResourceGroupInfo.ResourceGroupId),
				}
			}
			sdkData = append(sdkData, &ListMcpImagesResponseBodyData{
				ImageId:                dara.String(xmlImage.ImageId),
				ImageName:              dara.String(xmlImage.ImageName),
				ImageBuildType:         dara.String(xmlImage.ImageBuildType),
				ImageIntro:             dara.String(xmlImage.ImageIntro),
				ImageApplyScene:        dara.String(xmlImage.ImageApplyScene),
				ImageResourceStatus:    dara.String(xmlImage.ImageResourceStatus),
				ImageResourceGroupInfo: imageResourceGroupInfo,
				ImageInfo:              imageInfo,
				ToolInfo:               toolInfo,
			})
		}
		out.StatusCode = dara.Int32(int32(xr.HttpStatusCode))
		out.Body = &ListMcpImagesResponseBody{
			RequestId:      dara.String(xr.RequestId),
			HttpStatusCode: dara.Int32(int32(xr.HttpStatusCode)),
			Code:           dara.String(xr.Code),
			Success:        dara.Bool(xr.Success),
			Data:           sdkData,
			TotalCount:     dara.Int32(xr.TotalCount),
			PageSize:       dara.Int32(xr.PageSize),
			PageStart:      dara.Int32(xr.PageStart),
			NextToken:      dara.String(xr.NextToken),
		}
	} else {
		parsed := &ListMcpImagesResponseBody{}
		if bodyStr != "" {
			if err := json.Unmarshal([]byte(bodyStr), parsed); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}
		}
		out.Body = parsed
	}
	applyMapHeadersAndStatus(&out.Headers, &out.StatusCode, res)
	return out, nil
}

// --- GetMcpImageInfo ---

type xmlGetMcpImageInfoResponse struct {
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

func parseGetMcpImageInfoResponse(res map[string]interface{}) (*GetMcpImageInfoResponse, error) {
	bodyStr, err := rawBodyStringFromMap(res)
	if err != nil {
		return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
	}
	out := &GetMcpImageInfoResponse{Headers: make(map[string]*string)}
	trimmed := strings.TrimSpace(bodyStr)
	if len(trimmed) > 0 && trimmed[0] == '<' {
		var xr xmlGetMcpImageInfoResponse
		if err := xml.Unmarshal([]byte(bodyStr), &xr); err != nil {
			return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
		}
		var imageInfo *GetMcpImageInfoResponseBodyDataImageInfo
		if xr.Data.ImageInfo.OsName != "" || xr.Data.ImageInfo.Status != "" || xr.Data.ImageInfo.ImageType != "" {
			imageInfo = &GetMcpImageInfoResponseBodyDataImageInfo{
				OsName:         dara.String(xr.Data.ImageInfo.OsName),
				OsVersion:      dara.String(xr.Data.ImageInfo.OsVersion),
				PlatformName:   dara.String(xr.Data.ImageInfo.PlatformName),
				Status:         dara.String(xr.Data.ImageInfo.Status),
				DataDiskSize:   dara.Int32(xr.Data.ImageInfo.DataDiskSize),
				SystemDiskSize: dara.Int32(xr.Data.ImageInfo.SystemDiskSize),
				UpdateTime:     dara.String(xr.Data.ImageInfo.UpdateTime),
				ImageType:      dara.String(xr.Data.ImageInfo.ImageType),
			}
		}
		var imageBuildInfo *GetMcpImageInfoResponseBodyDataImageBuildInfo
		if xr.Data.ImageBuildInfo.TaskId != "" || xr.Data.ImageBuildInfo.VersionId != "" {
			imageBuildInfo = &GetMcpImageInfoResponseBodyDataImageBuildInfo{
				TaskId:                  dara.String(xr.Data.ImageBuildInfo.TaskId),
				VersionId:               dara.String(xr.Data.ImageBuildInfo.VersionId),
				ApiKeyId:                dara.String(xr.Data.ImageBuildInfo.ApiKeyId),
				InstanceReady:           dara.Bool(xr.Data.ImageBuildInfo.InstanceReady),
				AndroidMobileGroupId:    dara.String(xr.Data.ImageBuildInfo.AndroidMobileGroupId),
				AndroidMobileInstanceId: dara.String(xr.Data.ImageBuildInfo.AndroidMobileInstanceId),
			}
		}
		out.StatusCode = dara.Int32(int32(xr.HttpStatusCode))
		out.Body = &GetMcpImageInfoResponseBody{
			RequestId:      dara.String(xr.RequestId),
			HttpStatusCode: dara.Int32(int32(xr.HttpStatusCode)),
			Code:           dara.String(xr.Code),
			Success:        dara.Bool(xr.Success),
			Message:        dara.String(xr.Message),
			Data: &GetMcpImageInfoResponseBodyData{
				ImageId:             dara.String(xr.Data.ImageId),
				ImageName:           dara.String(xr.Data.ImageName),
				ImageBuildType:      dara.String(xr.Data.ImageBuildType),
				ImageApplyScene:     dara.String(xr.Data.ImageApplyScene),
				ImageResourceStatus: dara.String(xr.Data.ImageResourceStatus),
				ImageInfo:           imageInfo,
				ImageBuildInfo:      imageBuildInfo,
			},
		}
	} else {
		parsed := &GetMcpImageInfoResponseBody{}
		if bodyStr != "" {
			if err := json.Unmarshal([]byte(bodyStr), parsed); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}
		}
		out.Body = parsed
	}
	applyMapHeadersAndStatus(&out.Headers, &out.StatusCode, res)
	mergeDerivedGetMcpImageInfoHeaders(out)
	return out, nil
}

// --- CreateResourceGroup ---

type xmlCreateResourceGroupResponse struct {
	XMLName        xml.Name `xml:"CreateResourceGroupResponse"`
	RequestId      string   `xml:"RequestId"`
	HttpStatusCode int      `xml:"HttpStatusCode"`
	Code           string   `xml:"Code"`
	Success        bool     `xml:"Success"`
	Message        string   `xml:"Message"`
}

func parseCreateResourceGroupResponse(res map[string]interface{}) (*CreateResourceGroupResponse, error) {
	bodyStr, err := rawBodyStringFromMap(res)
	if err != nil {
		return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
	}
	out := &CreateResourceGroupResponse{Headers: make(map[string]*string)}
	trimmed := strings.TrimSpace(bodyStr)
	if len(trimmed) > 0 && trimmed[0] == '<' {
		var xr xmlCreateResourceGroupResponse
		if err := xml.Unmarshal([]byte(bodyStr), &xr); err != nil {
			return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
		}
		out.StatusCode = dara.Int32(int32(xr.HttpStatusCode))
		out.Body = &CreateResourceGroupResponseBody{
			RequestId:      dara.String(xr.RequestId),
			HttpStatusCode: dara.Int32(int32(xr.HttpStatusCode)),
			Code:           dara.String(xr.Code),
			Success:        dara.Bool(xr.Success),
			Message:        dara.String(xr.Message),
		}
	} else {
		parsed := &CreateResourceGroupResponseBody{}
		if bodyStr != "" {
			if err := json.Unmarshal([]byte(bodyStr), parsed); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}
		}
		out.Body = parsed
	}
	applyMapHeadersAndStatus(&out.Headers, &out.StatusCode, res)
	return out, nil
}

// --- DeleteResourceGroup ---

type xmlDeleteResourceGroupResponse struct {
	XMLName        xml.Name `xml:"DeleteResourceGroupResponse"`
	RequestId      string   `xml:"RequestId"`
	HttpStatusCode int      `xml:"HttpStatusCode"`
	Code           string   `xml:"Code"`
	Success        bool     `xml:"Success"`
	Message        string   `xml:"Message"`
}

func parseDeleteResourceGroupResponse(res map[string]interface{}) (*DeleteResourceGroupResponse, error) {
	bodyStr, err := rawBodyStringFromMap(res)
	if err != nil {
		return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
	}
	out := &DeleteResourceGroupResponse{Headers: make(map[string]*string)}
	trimmed := strings.TrimSpace(bodyStr)
	if len(trimmed) > 0 && trimmed[0] == '<' {
		var xr xmlDeleteResourceGroupResponse
		if err := xml.Unmarshal([]byte(bodyStr), &xr); err != nil {
			return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
		}
		out.StatusCode = dara.Int32(int32(xr.HttpStatusCode))
		out.Body = &DeleteResourceGroupResponseBody{
			RequestId:      dara.String(xr.RequestId),
			HttpStatusCode: dara.Int32(int32(xr.HttpStatusCode)),
			Code:           dara.String(xr.Code),
			Success:        dara.Bool(xr.Success),
			Message:        dara.String(xr.Message),
		}
	} else {
		parsed := &DeleteResourceGroupResponseBody{}
		if bodyStr != "" {
			if err := json.Unmarshal([]byte(bodyStr), parsed); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}
		}
		out.Body = parsed
	}
	applyMapHeadersAndStatus(&out.Headers, &out.StatusCode, res)
	return out, nil
}

// --- DeleteMcpImage ---

type xmlDeleteMcpImageResponse struct {
	XMLName        xml.Name `xml:"DeleteMcpImageResponse"`
	RequestId      string   `xml:"RequestId"`
	HttpStatusCode int      `xml:"HttpStatusCode"`
	Code           string   `xml:"Code"`
	Success        bool     `xml:"Success"`
	Message        string   `xml:"Message"`
}

func parseDeleteMcpImageResponse(res map[string]interface{}) (*DeleteMcpImageResponse, error) {
	bodyStr, err := rawBodyStringFromMap(res)
	if err != nil {
		return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
	}
	out := &DeleteMcpImageResponse{Headers: make(map[string]*string)}
	trimmed := strings.TrimSpace(bodyStr)
	if len(trimmed) > 0 && trimmed[0] == '<' {
		var xr xmlDeleteMcpImageResponse
		if err := xml.Unmarshal([]byte(bodyStr), &xr); err != nil {
			return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
		}
		out.StatusCode = dara.Int32(int32(xr.HttpStatusCode))
		out.Body = &DeleteMcpImageResponseBody{
			RequestId:      dara.String(xr.RequestId),
			HttpStatusCode: dara.Int32(int32(xr.HttpStatusCode)),
			Code:           dara.String(xr.Code),
			Success:        dara.Bool(xr.Success),
			Message:        dara.String(xr.Message),
		}
	} else {
		parsed := &DeleteMcpImageResponseBody{}
		if bodyStr != "" {
			if err := json.Unmarshal([]byte(bodyStr), parsed); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}
		}
		out.Body = parsed
	}
	applyMapHeadersAndStatus(&out.Headers, &out.StatusCode, res)
	return out, nil
}

// int32FromFlexibleJSON parses a JSON value that may be a number or a decimal string (some gateways stringify ints).
func int32FromFlexibleJSON(raw json.RawMessage) (*int32, error) {
	if len(raw) == 0 {
		return nil, nil
	}
	var v interface{}
	if err := json.Unmarshal(raw, &v); err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	switch x := v.(type) {
	case float64:
		return dara.Int32(int32(x)), nil
	case string:
		s := strings.TrimSpace(x)
		if s == "" {
			return nil, nil
		}
		n, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return nil, err
		}
		return dara.Int32(int32(n)), nil
	default:
		return nil, fmt.Errorf("expected number or string, got %T", v)
	}
}

// --- BatchCreateHideResourceGroupsWithMaxSession ---

// batchCreateHideResourceGroupsWithMaxSessionJSONWire decodes JSON tolerantly,
// because some gateways stringify HttpStatusCode (e.g. "200") which breaks *int32 unmarshal.
type batchCreateHideResourceGroupsWithMaxSessionJSONWire struct {
	Code           *string         `json:"Code"`
	Message        *string         `json:"Message"`
	RequestId      *string         `json:"RequestId"`
	HttpStatusCode json.RawMessage `json:"HttpStatusCode"`
	Success        *bool           `json:"Success"`
}

type xmlBatchCreateHideResourceGroupsWithMaxSessionResponse struct {
	XMLName        xml.Name `xml:"BatchCreateHideResourceGroupsWithMaxSessionResponse"`
	RequestId      string   `xml:"RequestId"`
	HttpStatusCode string   `xml:"HttpStatusCode"`
	Code           string   `xml:"Code"`
	Success        bool     `xml:"Success"`
	Message        string   `xml:"Message"`
}

func parseBatchCreateHideResourceGroupsWithMaxSessionResponse(res map[string]interface{}) (*BatchCreateHideResourceGroupsWithMaxSessionResponse, error) {
	bodyStr, err := rawBodyStringFromMap(res)
	if err != nil {
		return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
	}
	out := &BatchCreateHideResourceGroupsWithMaxSessionResponse{Headers: make(map[string]*string)}
	parsed := &BatchCreateHideResourceGroupsWithMaxSessionResponseBody{}
	trimmed := strings.TrimSpace(bodyStr)
	if bodyStr != "" {
		if len(trimmed) > 0 && trimmed[0] == '<' {
			var xr xmlBatchCreateHideResourceGroupsWithMaxSessionResponse
			if err := xml.Unmarshal([]byte(bodyStr), &xr); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}
			parsed.Code = dara.String(xr.Code)
			parsed.RequestId = dara.String(xr.RequestId)
			parsed.Success = dara.Bool(xr.Success)
			parsed.Message = dara.String(xr.Message)
			if s := strings.TrimSpace(xr.HttpStatusCode); s != "" {
				if n, perr := strconv.ParseInt(s, 10, 32); perr == nil {
					parsed.HttpStatusCode = dara.Int32(int32(n))
				}
			}
		} else {
			var wire batchCreateHideResourceGroupsWithMaxSessionJSONWire
			if err := json.Unmarshal([]byte(bodyStr), &wire); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}
			parsed.Code = wire.Code
			parsed.Message = wire.Message
			parsed.RequestId = wire.RequestId
			parsed.Success = wire.Success
			n, derr := int32FromFlexibleJSON(wire.HttpStatusCode)
			if derr != nil {
				return nil, &ErrWithRequestID{Err: fmt.Errorf("HttpStatusCode: %w", derr), RequestID: extractRequestIDFromResponse(res)}
			}
			parsed.HttpStatusCode = n
		}
	}
	out.Body = parsed
	applyMapHeadersAndStatus(&out.Headers, &out.StatusCode, res)
	return out, nil
}

// --- GetDockerfileTemplate ---

type getDockerfileTemplateJSONWire struct {
	Code *string `json:"Code"`
	Data *struct {
		OssDownloadUrl    *string         `json:"OssDownloadUrl"`
		NonEditLineNum    json.RawMessage `json:"NonEditLineNum"`
		DockerfileContent *string         `json:"DockerfileContent"`
	} `json:"Data"`
	HttpStatusCode *int32  `json:"HttpStatusCode"`
	Message        *string `json:"Message"`
	RequestId      *string `json:"RequestId"`
	Success        *bool   `json:"Success"`
}

type xmlGetDockerfileTemplateResponse struct {
	XMLName        xml.Name `xml:"GetDockerfileTemplateResponse"`
	RequestId      string   `xml:"RequestId"`
	HttpStatusCode int      `xml:"HttpStatusCode"`
	Data           struct {
		OssDownloadUrl    string `xml:"OssDownloadUrl"`
		NonEditLineNum    string `xml:"NonEditLineNum"`
		DockerfileContent string `xml:"DockerfileContent"`
	} `xml:"Data"`
	Code    string `xml:"Code"`
	Success bool   `xml:"Success"`
	Message string `xml:"Message"`
}

func parseGetDockerfileTemplateResponse(res map[string]interface{}) (*GetDockerfileTemplateResponse, error) {
	bodyStr, err := rawBodyStringFromMap(res)
	if err != nil {
		return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
	}
	out := &GetDockerfileTemplateResponse{Headers: make(map[string]*string)}
	trimmed := strings.TrimSpace(bodyStr)
	if len(trimmed) > 0 && trimmed[0] == '<' {
		var xr xmlGetDockerfileTemplateResponse
		if err := xml.Unmarshal([]byte(bodyStr), &xr); err != nil {
			return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
		}
		var nonEdit *int32
		if xr.Data.NonEditLineNum != "" {
			if num, err := strconv.ParseInt(xr.Data.NonEditLineNum, 10, 32); err == nil {
				n := int32(num)
				nonEdit = &n
			}
		}
		data := &GetDockerfileTemplateResponseBodyData{
			OssDownloadUrl: dara.String(xr.Data.OssDownloadUrl),
		}
		if nonEdit != nil {
			data.NonEditLineNum = nonEdit
		}
		if xr.Data.DockerfileContent != "" {
			data.DockerfileContent = dara.String(xr.Data.DockerfileContent)
		}
		out.StatusCode = dara.Int32(int32(xr.HttpStatusCode))
		out.Body = &GetDockerfileTemplateResponseBody{
			RequestId:      dara.String(xr.RequestId),
			HttpStatusCode: dara.Int32(int32(xr.HttpStatusCode)),
			Code:           dara.String(xr.Code),
			Success:        dara.Bool(xr.Success),
			Message:        dara.String(xr.Message),
			Data:           data,
		}
	} else {
		parsed := &GetDockerfileTemplateResponseBody{}
		if bodyStr != "" {
			var wire getDockerfileTemplateJSONWire
			if err := json.Unmarshal([]byte(bodyStr), &wire); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}
			parsed.Code = wire.Code
			parsed.HttpStatusCode = wire.HttpStatusCode
			parsed.Message = wire.Message
			parsed.RequestId = wire.RequestId
			parsed.Success = wire.Success
			if wire.Data != nil {
				data := &GetDockerfileTemplateResponseBodyData{
					OssDownloadUrl:    wire.Data.OssDownloadUrl,
					DockerfileContent: wire.Data.DockerfileContent,
				}
				n, err := int32FromFlexibleJSON(wire.Data.NonEditLineNum)
				if err != nil {
					return nil, &ErrWithRequestID{Err: fmt.Errorf("Data.NonEditLineNum: %w", err), RequestID: extractRequestIDFromResponse(res)}
				}
				data.NonEditLineNum = n
				parsed.Data = data
			}
		}
		out.Body = parsed
	}
	applyMapHeadersAndStatus(&out.Headers, &out.StatusCode, res)
	return out, nil
}

// --- DescribeMcpApiKey ---

type describeMcpApiKeyJSONWire struct {
	Code           *string         `json:"Code"`
	Data           json.RawMessage `json:"Data"`
	HttpStatusCode json.RawMessage `json:"HttpStatusCode"`
	Message        *string         `json:"Message"`
	RequestId      *string         `json:"RequestId"`
	Success        *bool           `json:"Success"`
}

type xmlDescribeMcpApiKeyResponse struct {
	XMLName        xml.Name `xml:"DescribeMcpApiKeyResponse"`
	RequestId      string   `xml:"RequestId"`
	HttpStatusCode string   `xml:"HttpStatusCode"`
	Code           string   `xml:"Code"`
	Success        bool     `xml:"Success"`
	Message        string   `xml:"Message"`
	Data           struct {
		Status   string `xml:"Status"`
		ApiKeyId string `xml:"ApiKeyId"`
		Name     string `xml:"Name"`
		AliUid   string `xml:"AliUid"`
	} `xml:"Data"`
}

func parseDescribeMcpApiKeyResponse(res map[string]interface{}) (*DescribeMcpApiKeyResponse, error) {
	bodyStr, err := rawBodyStringFromMap(res)
	if err != nil {
		return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
	}
	out := &DescribeMcpApiKeyResponse{Headers: make(map[string]*string)}
	parsed := &DescribeMcpApiKeyResponseBody{}
	trimmed := strings.TrimSpace(bodyStr)
	if bodyStr != "" {
		if len(trimmed) > 0 && trimmed[0] == '<' {
			var xr xmlDescribeMcpApiKeyResponse
			if err := xml.Unmarshal([]byte(bodyStr), &xr); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}
			parsed.Code = dara.String(xr.Code)
			parsed.RequestId = dara.String(xr.RequestId)
			parsed.Success = dara.Bool(xr.Success)
			parsed.Message = dara.String(xr.Message)
			if s := strings.TrimSpace(xr.HttpStatusCode); s != "" {
				if n, perr := strconv.ParseInt(s, 10, 32); perr == nil {
					parsed.HttpStatusCode = dara.Int32(int32(n))
				}
			}
			parsed.Data = &DescribeMcpApiKeyResponseBodyData{
				Status:   dara.String(xr.Data.Status),
				ApiKeyId: dara.String(xr.Data.ApiKeyId),
				Name:     dara.String(xr.Data.Name),
				AliUid:   dara.String(xr.Data.AliUid),
			}
		} else {
			var wire describeMcpApiKeyJSONWire
			if err := json.Unmarshal([]byte(bodyStr), &wire); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}
			parsed.Code = wire.Code
			parsed.Message = wire.Message
			parsed.RequestId = wire.RequestId
			parsed.Success = wire.Success
			n, derr := int32FromFlexibleJSON(wire.HttpStatusCode)
			if derr != nil {
				return nil, &ErrWithRequestID{Err: fmt.Errorf("HttpStatusCode: %w", derr), RequestID: extractRequestIDFromResponse(res)}
			}
			parsed.HttpStatusCode = n
			if len(wire.Data) > 0 && string(wire.Data) != "null" {
				var data DescribeMcpApiKeyResponseBodyData
				if err := json.Unmarshal(wire.Data, &data); err != nil {
					return nil, &ErrWithRequestID{Err: fmt.Errorf("Data: %w", err), RequestID: extractRequestIDFromResponse(res)}
				}
				parsed.Data = &data
			}
		}
	}
	out.Body = parsed
	applyMapHeadersAndStatus(&out.Headers, &out.StatusCode, res)
	return out, nil
}

// --- ModifyApiKeyStatus ---

type modifyApiKeyStatusJSONWire struct {
	Code           *string         `json:"Code"`
	Message        *string         `json:"Message"`
	RequestId      *string         `json:"RequestId"`
	HttpStatusCode json.RawMessage `json:"HttpStatusCode"`
	Success        *bool           `json:"Success"`
}

type xmlModifyApiKeyStatusResponse struct {
	XMLName        xml.Name `xml:"ModifyApiKeyStatusResponse"`
	RequestId      string   `xml:"RequestId"`
	HttpStatusCode string   `xml:"HttpStatusCode"`
	Code           string   `xml:"Code"`
	Success        bool     `xml:"Success"`
	Message        string   `xml:"Message"`
}

func parseModifyApiKeyStatusResponse(res map[string]interface{}) (*ModifyApiKeyStatusResponse, error) {
	bodyStr, err := rawBodyStringFromMap(res)
	if err != nil {
		return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
	}
	out := &ModifyApiKeyStatusResponse{Headers: make(map[string]*string)}
	parsed := &ModifyApiKeyStatusResponseBody{}
	trimmed := strings.TrimSpace(bodyStr)
	if bodyStr != "" {
		if len(trimmed) > 0 && trimmed[0] == '<' {
			var xr xmlModifyApiKeyStatusResponse
			if err := xml.Unmarshal([]byte(bodyStr), &xr); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}
			parsed.Code = dara.String(xr.Code)
			parsed.RequestId = dara.String(xr.RequestId)
			parsed.Success = dara.Bool(xr.Success)
			parsed.Message = dara.String(xr.Message)
			if s := strings.TrimSpace(xr.HttpStatusCode); s != "" {
				if n, perr := strconv.ParseInt(s, 10, 32); perr == nil {
					parsed.HttpStatusCode = dara.Int32(int32(n))
				}
			}
		} else {
			var wire modifyApiKeyStatusJSONWire
			if err := json.Unmarshal([]byte(bodyStr), &wire); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}
			parsed.Code = wire.Code
			parsed.Message = wire.Message
			parsed.RequestId = wire.RequestId
			parsed.Success = wire.Success
			n, derr := int32FromFlexibleJSON(wire.HttpStatusCode)
			if derr != nil {
				return nil, &ErrWithRequestID{Err: fmt.Errorf("HttpStatusCode: %w", derr), RequestID: extractRequestIDFromResponse(res)}
			}
			parsed.HttpStatusCode = n
		}
	}
	out.Body = parsed
	applyMapHeadersAndStatus(&out.Headers, &out.StatusCode, res)
	return out, nil
}

type deleteApiKeyJSONWire struct {
	Code           *string         `json:"Code"`
	Message        *string         `json:"Message"`
	RequestId      *string         `json:"RequestId"`
	HttpStatusCode json.RawMessage `json:"HttpStatusCode"`
	Success        *bool           `json:"Success"`
}

type xmlDeleteApiKeyResponse struct {
	XMLName        xml.Name `xml:"DeleteApiKeyResponse"`
	RequestId      string   `xml:"RequestId"`
	HttpStatusCode string   `xml:"HttpStatusCode"`
	Code           string   `xml:"Code"`
	Success        bool     `xml:"Success"`
	Message        string   `xml:"Message"`
}

func parseDeleteApiKeyResponse(res map[string]interface{}) (*DeleteApiKeyResponse, error) {
	bodyStr, err := rawBodyStringFromMap(res)
	if err != nil {
		return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
	}
	out := &DeleteApiKeyResponse{Headers: make(map[string]*string)}
	parsed := &DeleteApiKeyResponseBody{}
	trimmed := strings.TrimSpace(bodyStr)
	if bodyStr != "" {
		if len(trimmed) > 0 && trimmed[0] == '<' {
			var xr xmlDeleteApiKeyResponse
			if err := xml.Unmarshal([]byte(bodyStr), &xr); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}
			parsed.Code = dara.String(xr.Code)
			parsed.RequestId = dara.String(xr.RequestId)
			parsed.Success = dara.Bool(xr.Success)
			parsed.Message = dara.String(xr.Message)
			if s := strings.TrimSpace(xr.HttpStatusCode); s != "" {
				if n, perr := strconv.ParseInt(s, 10, 32); perr == nil {
					parsed.HttpStatusCode = dara.Int32(int32(n))
				}
			}
		} else {
			var wire deleteApiKeyJSONWire
			if err := json.Unmarshal([]byte(bodyStr), &wire); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}
			parsed.Code = wire.Code
			parsed.Message = wire.Message
			parsed.RequestId = wire.RequestId
			parsed.Success = wire.Success
			n, derr := int32FromFlexibleJSON(wire.HttpStatusCode)
			if derr != nil {
				return nil, &ErrWithRequestID{Err: fmt.Errorf("HttpStatusCode: %w", derr), RequestID: extractRequestIDFromResponse(res)}
			}
			parsed.HttpStatusCode = n
		}
	}
	out.Body = parsed
	applyMapHeadersAndStatus(&out.Headers, &out.StatusCode, res)
	return out, nil
}

// --- DescribeApiKeys ---

// describeApiKeysJSONWire handles the case where the SDK returns a wrapped response
// ({"code":"200","data":{...},"successResponse":true,...}). This may happen in some
// environments, but the standard Alibaba Cloud SDK with BodyType "string" typically
// strips the outer wrapper and returns only the data payload as the body string.
type describeApiKeysJSONWire struct {
	Code           *string         `json:"code"`
	Data           json.RawMessage `json:"data"`
	HttpStatusCode json.RawMessage `json:"httpStatusCode"`
	Message        *string         `json:"message"`
	RequestId      *string         `json:"requestId"`
	Success        *bool           `json:"successResponse"`
}

// describeApiKeysDataJSONWire represents the data payload returned by the SDK.
// The SDK with BodyType "string" returns the data payload directly as the body string,
// so this struct is used to parse the body when there is no outer wrapper.
// Field names match the actual server response (mixed case).
type describeApiKeysDataJSONWire struct {
	ApiKeys   []describeApiKeysApiKeyJSONWire `json:"ApiKeys"`
	RequestId *string                         `json:"requestId"`
	Count     *string                         `json:"Count"`
	NextToken *string                         `json:"NextToken"`
}

type describeApiKeysApiKeyJSONWire struct {
	Status        *string         `json:"Status"`
	GmtCreate     *string         `json:"GmtCreate"`
	LastUseDate   *string         `json:"LastUseDate"`
	ApiKey        *string         `json:"ApiKey"`
	Concurrency   json.RawMessage `json:"Concurrency"`
	KeyId         *string         `json:"KeyId"`
	Name          *string         `json:"Name"`
	BoundPolicy   json.RawMessage `json:"BoundPolicy"`
	BoundResource json.RawMessage `json:"BoundResource"`
}

type xmlDescribeApiKeysApiKeyEntry struct {
	Status      string `xml:"Status"`
	GmtCreate   string `xml:"GmtCreate"`
	LastUseDate string `xml:"LastUseDate"`
	ApiKey      string `xml:"ApiKey"`
	Concurrency string `xml:"Concurrency"`
	KeyId       string `xml:"KeyId"`
	Name        string `xml:"Name"`
}

type xmlDescribeApiKeysResponse struct {
	XMLName        xml.Name `xml:"DescribeApiKeysResponse"`
	RequestId      string   `xml:"RequestId"`
	HttpStatusCode string   `xml:"HttpStatusCode"`
	Code           string   `xml:"Code"`
	Success        bool     `xml:"Success"`
	Message        string   `xml:"Message"`
	Data           struct {
		ApiKeys   []xmlDescribeApiKeysApiKeyEntry `xml:"ApiKeys>ApiKey"`
		RequestId string                          `xml:"RequestId"`
		Count     string                          `xml:"Count"`
		NextToken string                          `xml:"NextToken"`
	} `xml:"Data"`
}

func parseDescribeApiKeysResponse(res map[string]interface{}) (*DescribeApiKeysResponse, error) {
	bodyStr, err := rawBodyStringFromMap(res)
	if err != nil {
		return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
	}
	out := &DescribeApiKeysResponse{Headers: make(map[string]*string)}
	parsed := &DescribeApiKeysResponseBody{}
	trimmed := strings.TrimSpace(bodyStr)
	if bodyStr != "" {
		if len(trimmed) > 0 && trimmed[0] == '<' {
			var xr xmlDescribeApiKeysResponse
			if err := xml.Unmarshal([]byte(bodyStr), &xr); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}
			parsed.Code = dara.String(xr.Code)
			parsed.RequestId = dara.String(xr.RequestId)
			parsed.Success = dara.Bool(xr.Success)
			parsed.Message = dara.String(xr.Message)
			if s := strings.TrimSpace(xr.HttpStatusCode); s != "" {
				if n, perr := strconv.ParseInt(s, 10, 32); perr == nil {
					parsed.HttpStatusCode = dara.Int32(int32(n))
				}
			}
			var apiKeys []*DescribeApiKeysResponseBodyDataApiKey
			for _, xmlKey := range xr.Data.ApiKeys {
				var concurrency *int32
				if s := strings.TrimSpace(xmlKey.Concurrency); s != "" {
					if n, perr := strconv.ParseInt(s, 10, 32); perr == nil {
						concurrency = dara.Int32(int32(n))
					}
				}
				apiKeys = append(apiKeys, &DescribeApiKeysResponseBodyDataApiKey{
					Status:      dara.String(xmlKey.Status),
					GmtCreate:   dara.String(xmlKey.GmtCreate),
					LastUseDate: dara.String(xmlKey.LastUseDate),
					ApiKey:      dara.String(xmlKey.ApiKey),
					Concurrency: concurrency,
					KeyId:       dara.String(xmlKey.KeyId),
					Name:        dara.String(xmlKey.Name),
				})
			}
			parsed.Data = &DescribeApiKeysResponseBodyData{
				ApiKeys:   apiKeys,
				RequestId: dara.String(xr.Data.RequestId),
				Count:     dara.String(xr.Data.Count),
				NextToken: dara.String(xr.Data.NextToken),
			}
		} else {
			// The Alibaba Cloud SDK with BodyType "string" typically returns only
			// the data payload as the body string, NOT the full wrapped response.
			// So the body is directly: {"ApiKeys":[...], "requestId":"...", ...}
			// rather than: {"code":"200","data":{"ApiKeys":[...]}, ...}
			// We handle both formats: first try the outer wrapper; if Data is empty,
			// parse the body directly as the data payload.

			var wire describeApiKeysJSONWire
			if err := json.Unmarshal([]byte(bodyStr), &wire); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}

			var rawApiKeys []describeApiKeysApiKeyJSONWire
			var dataRequestId, dataCount, dataNextToken *string

			if len(wire.Data) > 0 && string(wire.Data) != "null" {
				// Wrapped response format: body has outer-level code, data, etc.
				parsed.Code = wire.Code
				parsed.Message = wire.Message
				parsed.RequestId = wire.RequestId
				parsed.Success = wire.Success
				n, derr := int32FromFlexibleJSON(wire.HttpStatusCode)
				if derr != nil {
					return nil, &ErrWithRequestID{Err: fmt.Errorf("HttpStatusCode: %w", derr), RequestID: extractRequestIDFromResponse(res)}
				}
				parsed.HttpStatusCode = n

				var dataWire describeApiKeysDataJSONWire
				if err := json.Unmarshal(wire.Data, &dataWire); err != nil {
					return nil, &ErrWithRequestID{Err: fmt.Errorf("Data: %w", err), RequestID: extractRequestIDFromResponse(res)}
				}
				rawApiKeys = dataWire.ApiKeys
				dataRequestId = dataWire.RequestId
				dataCount = dataWire.Count
				dataNextToken = dataWire.NextToken
			} else {
				// Direct data payload: body IS the data (SDK strips outer wrapper).
				// Outer-level Code/Success/HttpStatusCode are not available in the body;
				// they will remain nil and the command-level SOP handles that correctly.
				var dataWire describeApiKeysDataJSONWire
				if err := json.Unmarshal([]byte(bodyStr), &dataWire); err != nil {
					return nil, &ErrWithRequestID{Err: fmt.Errorf("Data: %w", err), RequestID: extractRequestIDFromResponse(res)}
				}
				parsed.RequestId = dataWire.RequestId
				rawApiKeys = dataWire.ApiKeys
				dataRequestId = dataWire.RequestId
				dataCount = dataWire.Count
				dataNextToken = dataWire.NextToken
			}

			// Parse ApiKeys into SDK model structs
			var apiKeys []*DescribeApiKeysResponseBodyDataApiKey
			for _, keyWire := range rawApiKeys {
				c, cerr := int32FromFlexibleJSON(keyWire.Concurrency)
				if cerr != nil {
					return nil, &ErrWithRequestID{Err: fmt.Errorf("ApiKeys.Concurrency: %w", cerr), RequestID: extractRequestIDFromResponse(res)}
				}
				var boundPolicy *DescribeApiKeysResponseBodyDataApiKeyBoundPolicy
				if len(keyWire.BoundPolicy) > 0 && string(keyWire.BoundPolicy) != "null" {
					json.Unmarshal(keyWire.BoundPolicy, &boundPolicy)
				}
				apiKeys = append(apiKeys, &DescribeApiKeysResponseBodyDataApiKey{
					Status:        keyWire.Status,
					GmtCreate:     keyWire.GmtCreate,
					LastUseDate:   keyWire.LastUseDate,
					ApiKey:        keyWire.ApiKey,
					Concurrency:   c,
					KeyId:         keyWire.KeyId,
					Name:          keyWire.Name,
					BoundPolicy:   boundPolicy,
				})
			}
			parsed.Data = &DescribeApiKeysResponseBodyData{
				ApiKeys:   apiKeys,
				RequestId: dataRequestId,
				Count:     dataCount,
				NextToken: dataNextToken,
			}
		}
	}
	out.Body = parsed
	applyMapHeadersAndStatus(&out.Headers, &out.StatusCode, res)
	return out, nil
}
