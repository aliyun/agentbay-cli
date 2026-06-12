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

// --- UpdateMarketSkill ---

type xmlUpdateMarketSkillResponse struct {
	XMLName        xml.Name `xml:"UpdateMarketSkillResponse"`
	HttpStatusCode *int32   `xml:"HttpStatusCode"`
	Data           string   `xml:"Data"`
	RequestId      *string  `xml:"RequestId"`
	Code           *string  `xml:"Code"`
	Success        *bool    `xml:"Success"`
}

// parseUpdateMarketSkillResponse builds UpdateMarketSkillResponse from CallApi map (bodyType "string").
// Backend may return XML or JSON; JSON may use Data as either a string (skill id) or an object {SkillId}.
func parseUpdateMarketSkillResponse(res map[string]interface{}) (*UpdateMarketSkillResponse, error) {
	bodyStr, err := rawBodyStringFromMap(res)
	if err != nil {
		return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
	}
	out := &UpdateMarketSkillResponse{RawBody: bodyStr}
	parsed := &CreateMarketSkillResponseBody{}
	trimmed := strings.TrimSpace(bodyStr)
	if bodyStr != "" {
		if len(trimmed) > 0 && trimmed[0] == '<' {
			var xmlResp xmlUpdateMarketSkillResponse
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

// --- DescribeWarmUpStatusOpen ---

type describeWarmUpStatusOpenJSONWireDataImage struct {
	ImageId               *string         `json:"ImageId"`
	TotalMaxSize          json.RawMessage `json:"TotalMaxSize"`
	GroupCount            json.RawMessage `json:"GroupCount"`
	AvailableInstanceSize json.RawMessage `json:"AvailableInstanceSize"`
}

type describeWarmUpStatusOpenJSONWire struct {
	Code           *string         `json:"Code"`
	Message        *string         `json:"Message"`
	RequestId      *string         `json:"RequestId"`
	HttpStatusCode json.RawMessage `json:"HttpStatusCode"`
	Success        *bool           `json:"Success"`
	Data           *struct {
		MaxSessionNumLimit    json.RawMessage                             `json:"MaxSessionNumLimit"`
		TotalUsedSessionQuota json.RawMessage                             `json:"TotalUsedSessionQuota"`
		AvailableSessionQuota json.RawMessage                             `json:"AvailableSessionQuota"`
		MaxImageCount         json.RawMessage                             `json:"MaxImageCount"`
		CurrentImageCount     json.RawMessage                             `json:"CurrentImageCount"`
		Images                []describeWarmUpStatusOpenJSONWireDataImage `json:"Images"`
	} `json:"Data"`
}

type xmlDescribeWarmUpStatusOpenDataImage struct {
	ImageId               string `xml:"ImageId"`
	TotalMaxSize          string `xml:"TotalMaxSize"`
	GroupCount            string `xml:"GroupCount"`
	AvailableInstanceSize string `xml:"AvailableInstanceSize"`
}

type xmlDescribeWarmUpStatusOpenData struct {
	MaxSessionNumLimit    string                                 `xml:"MaxSessionNumLimit"`
	TotalUsedSessionQuota string                                 `xml:"TotalUsedSessionQuota"`
	AvailableSessionQuota string                                 `xml:"AvailableSessionQuota"`
	MaxImageCount         string                                 `xml:"MaxImageCount"`
	CurrentImageCount     string                                 `xml:"CurrentImageCount"`
	Images                []xmlDescribeWarmUpStatusOpenDataImage `xml:"Images>Image"`
}

type xmlDescribeWarmUpStatusOpenResponse struct {
	XMLName        xml.Name                        `xml:"DescribeWarmUpStatusOpenResponse"`
	RequestId      string                          `xml:"RequestId"`
	HttpStatusCode string                          `xml:"HttpStatusCode"`
	Code           string                          `xml:"Code"`
	Success        bool                            `xml:"Success"`
	Message        string                          `xml:"Message"`
	Data           xmlDescribeWarmUpStatusOpenData `xml:"Data"`
}

func parseDescribeWarmUpStatusOpenResponse(res map[string]interface{}) (*DescribeWarmUpStatusOpenResponse, error) {
	bodyStr, err := rawBodyStringFromMap(res)
	if err != nil {
		return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
	}
	out := &DescribeWarmUpStatusOpenResponse{Headers: make(map[string]*string)}
	parsed := &DescribeWarmUpStatusOpenResponseBody{}
	trimmed := strings.TrimSpace(bodyStr)
	if bodyStr != "" {
		if len(trimmed) > 0 && trimmed[0] == '<' {
			var xr xmlDescribeWarmUpStatusOpenResponse
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
			data := &DescribeWarmUpStatusOpenResponseBodyData{}
			if s := strings.TrimSpace(xr.Data.MaxSessionNumLimit); s != "" {
				if n, perr := strconv.ParseInt(s, 10, 32); perr == nil {
					data.MaxSessionNumLimit = dara.Int32(int32(n))
				}
			}
			if s := strings.TrimSpace(xr.Data.TotalUsedSessionQuota); s != "" {
				if n, perr := strconv.ParseInt(s, 10, 32); perr == nil {
					data.TotalUsedSessionQuota = dara.Int32(int32(n))
				}
			}
			if s := strings.TrimSpace(xr.Data.AvailableSessionQuota); s != "" {
				if n, perr := strconv.ParseInt(s, 10, 32); perr == nil {
					data.AvailableSessionQuota = dara.Int32(int32(n))
				}
			}
			if s := strings.TrimSpace(xr.Data.MaxImageCount); s != "" {
				if n, perr := strconv.ParseInt(s, 10, 32); perr == nil {
					data.MaxImageCount = dara.Int32(int32(n))
				}
			}
			if s := strings.TrimSpace(xr.Data.CurrentImageCount); s != "" {
				if n, perr := strconv.ParseInt(s, 10, 32); perr == nil {
					data.CurrentImageCount = dara.Int32(int32(n))
				}
			}
			for _, img := range xr.Data.Images {
				item := &DescribeWarmUpStatusOpenResponseBodyDataImage{
					ImageId: dara.String(img.ImageId),
				}
				if s := strings.TrimSpace(img.TotalMaxSize); s != "" {
					if n, perr := strconv.ParseInt(s, 10, 32); perr == nil {
						item.TotalMaxSize = dara.Int32(int32(n))
					}
				}
				if s := strings.TrimSpace(img.GroupCount); s != "" {
					if n, perr := strconv.ParseInt(s, 10, 32); perr == nil {
						item.GroupCount = dara.Int32(int32(n))
					}
				}
				if s := strings.TrimSpace(img.AvailableInstanceSize); s != "" {
					if n, perr := strconv.ParseInt(s, 10, 32); perr == nil {
						item.AvailableInstanceSize = dara.Int32(int32(n))
					}
				}
				data.Images = append(data.Images, item)
			}
			parsed.Data = data
		} else {
			var wire describeWarmUpStatusOpenJSONWire
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
			if wire.Data != nil {
				data := &DescribeWarmUpStatusOpenResponseBodyData{}
				v, derr := int32FromFlexibleJSON(wire.Data.MaxSessionNumLimit)
				if derr != nil {
					return nil, &ErrWithRequestID{Err: fmt.Errorf("Data.MaxSessionNumLimit: %w", derr), RequestID: extractRequestIDFromResponse(res)}
				}
				data.MaxSessionNumLimit = v
				v, derr = int32FromFlexibleJSON(wire.Data.TotalUsedSessionQuota)
				if derr != nil {
					return nil, &ErrWithRequestID{Err: fmt.Errorf("Data.TotalUsedSessionQuota: %w", derr), RequestID: extractRequestIDFromResponse(res)}
				}
				data.TotalUsedSessionQuota = v
				v, derr = int32FromFlexibleJSON(wire.Data.AvailableSessionQuota)
				if derr != nil {
					return nil, &ErrWithRequestID{Err: fmt.Errorf("Data.AvailableSessionQuota: %w", derr), RequestID: extractRequestIDFromResponse(res)}
				}
				data.AvailableSessionQuota = v
				v, derr = int32FromFlexibleJSON(wire.Data.MaxImageCount)
				if derr != nil {
					return nil, &ErrWithRequestID{Err: fmt.Errorf("Data.MaxImageCount: %w", derr), RequestID: extractRequestIDFromResponse(res)}
				}
				data.MaxImageCount = v
				v, derr = int32FromFlexibleJSON(wire.Data.CurrentImageCount)
				if derr != nil {
					return nil, &ErrWithRequestID{Err: fmt.Errorf("Data.CurrentImageCount: %w", derr), RequestID: extractRequestIDFromResponse(res)}
				}
				data.CurrentImageCount = v
				for _, img := range wire.Data.Images {
					item := &DescribeWarmUpStatusOpenResponseBodyDataImage{
						ImageId: img.ImageId,
					}
					sz, derr := int32FromFlexibleJSON(img.TotalMaxSize)
					if derr != nil {
						return nil, &ErrWithRequestID{Err: fmt.Errorf("Images.TotalMaxSize: %w", derr), RequestID: extractRequestIDFromResponse(res)}
					}
					item.TotalMaxSize = sz
					gc, derr := int32FromFlexibleJSON(img.GroupCount)
					if derr != nil {
						return nil, &ErrWithRequestID{Err: fmt.Errorf("Images.GroupCount: %w", derr), RequestID: extractRequestIDFromResponse(res)}
					}
					item.GroupCount = gc
					ais, derr := int32FromFlexibleJSON(img.AvailableInstanceSize)
					if derr != nil {
						return nil, &ErrWithRequestID{Err: fmt.Errorf("Images.AvailableInstanceSize: %w", derr), RequestID: extractRequestIDFromResponse(res)}
					}
					item.AvailableInstanceSize = ais
					data.Images = append(data.Images, item)
				}
				parsed.Data = data
			}
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
					Status:      keyWire.Status,
					GmtCreate:   keyWire.GmtCreate,
					LastUseDate: keyWire.LastUseDate,
					ApiKey:      keyWire.ApiKey,
					Concurrency: c,
					KeyId:       keyWire.KeyId,
					Name:        keyWire.Name,
					BoundPolicy: boundPolicy,
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

// --- DescribeKeyContent ---

// describeKeyContentJSONWire handles the outer-wrapped response format:
// {"code":"200","data":{"ApiKey":"akm-xxx","RequestId":"..."},"httpStatusCode":"200","requestId":"...","successResponse":true}
type describeKeyContentJSONWire struct {
	Code           *string         `json:"code"`
	Data           json.RawMessage `json:"data"`
	HttpStatusCode json.RawMessage `json:"httpStatusCode"`
	Message        *string         `json:"message"`
	RequestId      *string         `json:"requestId"`
	Success        *bool           `json:"successResponse"`
}

// describeKeyContentDataJSONWire is the inner data payload
type describeKeyContentDataJSONWire struct {
	ApiKey    *string `json:"ApiKey"`
	RequestId *string `json:"RequestId"`
}

type xmlDescribeKeyContentResponse struct {
	XMLName        xml.Name `xml:"DescribeKeyContentResponse"`
	RequestId      string   `xml:"RequestId"`
	HttpStatusCode string   `xml:"HttpStatusCode"`
	Code           string   `xml:"Code"`
	Success        bool     `xml:"Success"`
	Message        string   `xml:"Message"`
	Data           struct {
		ApiKey    string `xml:"ApiKey"`
		RequestId string `xml:"RequestId"`
	} `xml:"Data"`
}

func parseDescribeKeyContentResponse(res map[string]interface{}) (*DescribeKeyContentResponse, error) {
	bodyStr, err := rawBodyStringFromMap(res)
	if err != nil {
		return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
	}
	out := &DescribeKeyContentResponse{Headers: make(map[string]*string)}
	parsed := &DescribeKeyContentResponseBody{}
	trimmed := strings.TrimSpace(bodyStr)
	if bodyStr != "" {
		if len(trimmed) > 0 && trimmed[0] == '<' {
			// XML branch
			var xr xmlDescribeKeyContentResponse
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
			parsed.Data = &DescribeKeyContentResponseBodyData{
				ApiKey:    dara.String(xr.Data.ApiKey),
				RequestId: dara.String(xr.Data.RequestId),
			}
		} else {
			// JSON branch: try outer-wrapped format first
			var wire describeKeyContentJSONWire
			if err := json.Unmarshal([]byte(bodyStr), &wire); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}
			if len(wire.Data) > 0 && string(wire.Data) != "null" {
				// Outer-wrapped format: body has code, data, httpStatusCode, etc.
				parsed.Code = wire.Code
				parsed.Message = wire.Message
				parsed.RequestId = wire.RequestId
				parsed.Success = wire.Success
				n, derr := int32FromFlexibleJSON(wire.HttpStatusCode)
				if derr != nil {
					return nil, &ErrWithRequestID{Err: fmt.Errorf("HttpStatusCode: %w", derr), RequestID: extractRequestIDFromResponse(res)}
				}
				parsed.HttpStatusCode = n
				var dataWire describeKeyContentDataJSONWire
				if err := json.Unmarshal(wire.Data, &dataWire); err != nil {
					return nil, &ErrWithRequestID{Err: fmt.Errorf("Data: %w", err), RequestID: extractRequestIDFromResponse(res)}
				}
				parsed.Data = &DescribeKeyContentResponseBodyData{
					ApiKey:    dataWire.ApiKey,
					RequestId: dataWire.RequestId,
				}
			} else {
				// Direct data payload: body IS the data object
				var dataWire describeKeyContentDataJSONWire
				if err := json.Unmarshal([]byte(bodyStr), &dataWire); err != nil {
					return nil, &ErrWithRequestID{Err: fmt.Errorf("Data: %w", err), RequestID: extractRequestIDFromResponse(res)}
				}
				parsed.Data = &DescribeKeyContentResponseBodyData{
					ApiKey:    dataWire.ApiKey,
					RequestId: dataWire.RequestId,
				}
				parsed.RequestId = dataWire.RequestId
			}
		}
	}
	out.Body = parsed
	applyMapHeadersAndStatus(&out.Headers, &out.StatusCode, res)
	return out, nil
}

// --- ListTag ---

type listTagJSONWire struct {
	Code           *string         `json:"Code"`
	Message        *string         `json:"Message"`
	RequestId      *string         `json:"RequestId"`
	HttpStatusCode json.RawMessage `json:"HttpStatusCode"`
	Success        *bool           `json:"Success"`
	Data           json.RawMessage `json:"Data"`
}

type xmlListTagResponse struct {
	XMLName        xml.Name `xml:"ListTagResponse"`
	RequestId      string   `xml:"RequestId"`
	HttpStatusCode string   `xml:"HttpStatusCode"`
	Code           string   `xml:"Code"`
	Success        bool     `xml:"Success"`
	Message        string   `xml:"Message"`
	Data           []struct {
		TagName string `xml:"TagName"`
		TagId   string `xml:"TagId"`
	} `xml:"Data"`
}

func parseListTagResponse(res map[string]interface{}) (*ListTagResponse, error) {
	bodyStr, err := rawBodyStringFromMap(res)
	if err != nil {
		return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
	}
	out := &ListTagResponse{Headers: make(map[string]*string)}
	parsed := &ListTagResponseBody{}
	trimmed := strings.TrimSpace(bodyStr)
	if bodyStr != "" {
		if len(trimmed) > 0 && trimmed[0] == '<' {
			var xr xmlListTagResponse
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
			for _, item := range xr.Data {
				parsed.Data = append(parsed.Data, ListTagResponseBodyDataItem{
					TagName: dara.String(item.TagName),
					TagId:   dara.String(item.TagId),
				})
			}
		} else {
			var wire listTagJSONWire
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
			if len(wire.Data) > 0 {
				var items []ListTagResponseBodyDataItem
				if err := json.Unmarshal(wire.Data, &items); err == nil {
					parsed.Data = items
				}
			}
		}
	}
	out.Body = parsed
	applyMapHeadersAndStatus(&out.Headers, &out.StatusCode, res)
	return out, nil
}

// --- ListMarketSkillByPage ---

// listMarketSkillByPageOuterJSONWire handles the outer-wrapped response format:
// {"code":"200","data":{"RequestId":"...","HttpStatusCode":200,"Data":{...},"Code":"ok"},...}
type listMarketSkillByPageOuterJSONWire struct {
	Code           *string         `json:"code"`
	Data           json.RawMessage `json:"data"`
	HttpStatusCode json.RawMessage `json:"httpStatusCode"`
	Message        *string         `json:"message"`
	RequestId      *string         `json:"requestId"`
	Success        *bool           `json:"successResponse"`
}

// listMarketSkillByPageInnerJSONWire is the inner "data" payload with pagination info.
type listMarketSkillByPageInnerJSONWire struct {
	RequestId      *string                            `json:"RequestId"`
	HttpStatusCode json.RawMessage                    `json:"HttpStatusCode"`
	Code           *string                            `json:"Code"`
	Data           *listMarketSkillByPageDataJSONWire `json:"Data"`
}

// listMarketSkillByPageDataJSONWire holds the actual paginated data.
type listMarketSkillByPageDataJSONWire struct {
	TotalCount json.RawMessage                       `json:"TotalCount"`
	TotalPage  json.RawMessage                       `json:"TotalPage"`
	PageSize   json.RawMessage                       `json:"PageSize"`
	PageNumber json.RawMessage                       `json:"PageNumber"`
	Result     []listMarketSkillByPageResultJSONWire `json:"Result"`
}

// listMarketSkillByPageResultJSONWire is a single skill entry.
type listMarketSkillByPageResultJSONWire struct {
	SkillName   *string  `json:"SkillName"`
	SkillId     *string  `json:"SkillId"`
	TenantTags  []string `json:"TenantTags"`
	SkillStatus *string  `json:"SkillStatus"`
	GmtModified *string  `json:"GmtModified"`
	GmtCreate   *string  `json:"GmtCreate"`
	Description *string  `json:"Description"`
	Icon        *string  `json:"Icon"`
}

type xmlListMarketSkillByPageResult struct {
	SkillName   string   `xml:"SkillName"`
	SkillId     string   `xml:"SkillId"`
	TenantTags  []string `xml:"TenantTags>Tag"`
	SkillStatus string   `xml:"SkillStatus"`
	GmtModified string   `xml:"GmtModified"`
	GmtCreate   string   `xml:"GmtCreate"`
	Description string   `xml:"Description"`
	Icon        string   `xml:"Icon"`
}

type xmlListMarketSkillByPageResponse struct {
	XMLName        xml.Name `xml:"ListMarketSkillByPageResponse"`
	RequestId      string   `xml:"RequestId"`
	HttpStatusCode string   `xml:"HttpStatusCode"`
	Code           string   `xml:"Code"`
	Success        bool     `xml:"Success"`
	Message        string   `xml:"Message"`
	Data           struct {
		TotalCount string                           `xml:"TotalCount"`
		TotalPage  string                           `xml:"TotalPage"`
		PageSize   string                           `xml:"PageSize"`
		PageNumber string                           `xml:"PageNumber"`
		Result     []xmlListMarketSkillByPageResult `xml:"Result>Item"`
	} `xml:"Data"`
}

func parseListMarketSkillByPageResponse(res map[string]interface{}) (*ListMarketSkillByPageResponse, error) {
	bodyStr, err := rawBodyStringFromMap(res)
	if err != nil {
		return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
	}
	out := &ListMarketSkillByPageResponse{Headers: make(map[string]*string)}
	parsed := &ListMarketSkillByPageResponseBody{}
	trimmed := strings.TrimSpace(bodyStr)
	if bodyStr != "" {
		if len(trimmed) > 0 && trimmed[0] == '<' {
			// XML branch
			var xr xmlListMarketSkillByPageResponse
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
			data := &ListMarketSkillByPageResponseBodyData{}
			if s := strings.TrimSpace(xr.Data.TotalCount); s != "" {
				if n, perr := strconv.ParseInt(s, 10, 32); perr == nil {
					data.TotalCount = dara.Int32(int32(n))
				}
			}
			if s := strings.TrimSpace(xr.Data.TotalPage); s != "" {
				if n, perr := strconv.ParseInt(s, 10, 32); perr == nil {
					data.TotalPage = dara.Int32(int32(n))
				}
			}
			if s := strings.TrimSpace(xr.Data.PageSize); s != "" {
				if n, perr := strconv.ParseInt(s, 10, 32); perr == nil {
					data.PageSize = dara.Int32(int32(n))
				}
			}
			if s := strings.TrimSpace(xr.Data.PageNumber); s != "" {
				if n, perr := strconv.ParseInt(s, 10, 32); perr == nil {
					data.PageNumber = dara.Int32(int32(n))
				}
			}
			for _, item := range xr.Data.Result {
				data.Result = append(data.Result, &ListMarketSkillByPageResponseBodyDataResult{
					SkillName:   dara.String(item.SkillName),
					SkillId:     dara.String(item.SkillId),
					TenantTags:  item.TenantTags,
					SkillStatus: dara.String(item.SkillStatus),
					GmtModified: dara.String(item.GmtModified),
					GmtCreate:   dara.String(item.GmtCreate),
					Description: dara.String(item.Description),
					Icon:        dara.String(item.Icon),
				})
			}
			parsed.Data = data
		} else {
			// JSON branch: outer-wrapped format
			// {"code":"200","data":{"RequestId":"...","HttpStatusCode":200,"Data":{...},"Code":"ok"}}
			var outer listMarketSkillByPageOuterJSONWire
			if err := json.Unmarshal([]byte(bodyStr), &outer); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}
			parsed.RequestId = outer.RequestId
			n, derr := int32FromFlexibleJSON(outer.HttpStatusCode)
			if derr != nil {
				return nil, &ErrWithRequestID{Err: fmt.Errorf("HttpStatusCode: %w", derr), RequestID: extractRequestIDFromResponse(res)}
			}
			parsed.HttpStatusCode = n
			parsed.Success = outer.Success

			if len(outer.Data) > 0 && string(outer.Data) != "null" {
				var inner listMarketSkillByPageInnerJSONWire
				if err := json.Unmarshal(outer.Data, &inner); err != nil {
					return nil, &ErrWithRequestID{Err: fmt.Errorf("data: %w", err), RequestID: extractRequestIDFromResponse(res)}
				}
				// Use inner Code and RequestId if available
				if inner.Code != nil {
					parsed.Code = inner.Code
				} else {
					parsed.Code = outer.Code
				}
				if inner.RequestId != nil {
					parsed.RequestId = inner.RequestId
				}
				if inner.Data != nil {
					// Double-wrapped ACS format: outer.Data is {"RequestId":...,"Data":{pagination}}
					data := &ListMarketSkillByPageResponseBodyData{}
					tc, derr := int32FromFlexibleJSON(inner.Data.TotalCount)
					if derr != nil {
						return nil, &ErrWithRequestID{Err: fmt.Errorf("Data.TotalCount: %w", derr), RequestID: extractRequestIDFromResponse(res)}
					}
					data.TotalCount = tc
					tp, derr := int32FromFlexibleJSON(inner.Data.TotalPage)
					if derr != nil {
						return nil, &ErrWithRequestID{Err: fmt.Errorf("Data.TotalPage: %w", derr), RequestID: extractRequestIDFromResponse(res)}
					}
					data.TotalPage = tp
					ps, derr := int32FromFlexibleJSON(inner.Data.PageSize)
					if derr != nil {
						return nil, &ErrWithRequestID{Err: fmt.Errorf("Data.PageSize: %w", derr), RequestID: extractRequestIDFromResponse(res)}
					}
					data.PageSize = ps
					pn, derr := int32FromFlexibleJSON(inner.Data.PageNumber)
					if derr != nil {
						return nil, &ErrWithRequestID{Err: fmt.Errorf("Data.PageNumber: %w", derr), RequestID: extractRequestIDFromResponse(res)}
					}
					data.PageNumber = pn
					for _, item := range inner.Data.Result {
						data.Result = append(data.Result, &ListMarketSkillByPageResponseBodyDataResult{
							SkillName:   item.SkillName,
							SkillId:     item.SkillId,
							TenantTags:  item.TenantTags,
							SkillStatus: item.SkillStatus,
							GmtModified: item.GmtModified,
							GmtCreate:   item.GmtCreate,
							Description: item.Description,
							Icon:        item.Icon,
						})
					}
					parsed.Data = data
				} else {
					// Single-wrapped format: outer.Data is the pagination data directly
					// e.g. outer.Data = {"TotalCount":6,"TotalPage":1,...,"Result":[...]}
					// This happens when Go JSON case-insensitive matching maps "Data" -> outer.Data.
					var directData listMarketSkillByPageDataJSONWire
					if jerr := json.Unmarshal(outer.Data, &directData); jerr == nil && len(directData.Result) > 0 {
						data := &ListMarketSkillByPageResponseBodyData{}
						tc, _ := int32FromFlexibleJSON(directData.TotalCount)
						data.TotalCount = tc
						tp, _ := int32FromFlexibleJSON(directData.TotalPage)
						data.TotalPage = tp
						ps, _ := int32FromFlexibleJSON(directData.PageSize)
						data.PageSize = ps
						pn, _ := int32FromFlexibleJSON(directData.PageNumber)
						data.PageNumber = pn
						for _, item := range directData.Result {
							data.Result = append(data.Result, &ListMarketSkillByPageResponseBodyDataResult{
								SkillName:   item.SkillName,
								SkillId:     item.SkillId,
								TenantTags:  item.TenantTags,
								SkillStatus: item.SkillStatus,
								GmtModified: item.GmtModified,
								GmtCreate:   item.GmtCreate,
								Description: item.Description,
								Icon:        item.Icon,
							})
						}
						parsed.Data = data
					} else if jerr == nil {
						// Empty result set
						data := &ListMarketSkillByPageResponseBodyData{}
						tc, _ := int32FromFlexibleJSON(directData.TotalCount)
						data.TotalCount = tc
						tp, _ := int32FromFlexibleJSON(directData.TotalPage)
						data.TotalPage = tp
						ps, _ := int32FromFlexibleJSON(directData.PageSize)
						data.PageSize = ps
						pn, _ := int32FromFlexibleJSON(directData.PageNumber)
						data.PageNumber = pn
						parsed.Data = data
					}
				}
			} else {
				// outer.Data is nil: body has no data field at all
				parsed.Code = outer.Code
			}
		}
	}
	out.Body = parsed
	applyMapHeadersAndStatus(&out.Headers, &out.StatusCode, res)
	return out, nil
}

// --- CreateTag ---

type createTagJSONWire struct {
	Code           *string         `json:"Code"`
	Message        *string         `json:"Message"`
	RequestId      *string         `json:"RequestId"`
	HttpStatusCode json.RawMessage `json:"HttpStatusCode"`
	Success        *bool           `json:"Success"`
	Data           json.RawMessage `json:"Data"`
}

type xmlCreateTagResponse struct {
	XMLName        xml.Name `xml:"CreateTagResponse"`
	RequestId      string   `xml:"RequestId"`
	HttpStatusCode string   `xml:"HttpStatusCode"`
	Code           string   `xml:"Code"`
	Success        bool     `xml:"Success"`
	Message        string   `xml:"Message"`
	Data           []struct {
		TagName string `xml:"TagName"`
		TagId   string `xml:"TagId"`
	} `xml:"Data"`
}

func parseCreateTagResponse(res map[string]interface{}) (*CreateTagResponse, error) {
	bodyStr, err := rawBodyStringFromMap(res)
	if err != nil {
		return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
	}
	out := &CreateTagResponse{Headers: make(map[string]*string)}
	parsed := &CreateTagResponseBody{}
	trimmed := strings.TrimSpace(bodyStr)
	if bodyStr != "" {
		if len(trimmed) > 0 && trimmed[0] == '<' {
			var xr xmlCreateTagResponse
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
			for _, item := range xr.Data {
				parsed.Data = append(parsed.Data, CreateTagResponseBodyDataItem{
					TagName: dara.String(item.TagName),
					TagId:   dara.String(item.TagId),
				})
			}
		} else {
			var wire createTagJSONWire
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
			if len(wire.Data) > 0 {
				var items []CreateTagResponseBodyDataItem
				if err := json.Unmarshal(wire.Data, &items); err == nil {
					parsed.Data = items
				}
			}
		}
	}
	out.Body = parsed
	applyMapHeadersAndStatus(&out.Headers, &out.StatusCode, res)
	return out, nil
}

// --- DeleteMarketSkill ---

type deleteMarketSkillJSONWire struct {
	Code           *string         `json:"Code"`
	Data           *bool           `json:"Data"`
	HttpStatusCode json.RawMessage `json:"HttpStatusCode"`
	Message        *string         `json:"Message"`
	RequestId      *string         `json:"RequestId"`
	Success        *bool           `json:"Success"`
}

type xmlDeleteMarketSkillResponse struct {
	XMLName        struct{} `xml:"DeleteMarketSkillResponse"`
	Code           string   `xml:"Code"`
	Data           string   `xml:"Data"`
	HttpStatusCode string   `xml:"HttpStatusCode"`
	Message        string   `xml:"Message"`
	RequestId      string   `xml:"RequestId"`
	Success        string   `xml:"Success"`
}

func parseDeleteMarketSkillResponse(res map[string]interface{}) (*DeleteMarketSkillResponse, error) {
	bodyStr, err := rawBodyStringFromMap(res)
	if err != nil {
		return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
	}
	out := &DeleteMarketSkillResponse{Headers: make(map[string]*string)}
	parsed := &DeleteMarketSkillResponseBody{}
	trimmed := strings.TrimSpace(bodyStr)
	if bodyStr != "" {
		if len(trimmed) > 0 && trimmed[0] == '<' {
			var xr xmlDeleteMarketSkillResponse
			if err := xml.Unmarshal([]byte(bodyStr), &xr); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}
			parsed.Code = dara.String(xr.Code)
			parsed.RequestId = dara.String(xr.RequestId)
			parsed.Message = dara.String(xr.Message)
			if s := strings.TrimSpace(xr.HttpStatusCode); s != "" {
				if n, perr := strconv.ParseInt(s, 10, 32); perr == nil {
					parsed.HttpStatusCode = dara.Int32(int32(n))
				}
			}
			if s := strings.TrimSpace(xr.Success); s == "true" {
				parsed.Success = dara.Bool(true)
			} else if s == "false" {
				parsed.Success = dara.Bool(false)
			}
			if s := strings.TrimSpace(xr.Data); s == "true" {
				parsed.Data = dara.Bool(true)
			} else if s == "false" {
				parsed.Data = dara.Bool(false)
			}
		} else {
			var wire deleteMarketSkillJSONWire
			if err := json.Unmarshal([]byte(bodyStr), &wire); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}
			parsed.Code = wire.Code
			parsed.Message = wire.Message
			parsed.RequestId = wire.RequestId
			parsed.Success = wire.Success
			parsed.Data = wire.Data
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

// --- ShareDockerRepo ---

type shareDockerRepoJSONWireData struct {
	TargetAliUid *int64  `json:"TargetAliUid"`
	OwnerAliUid  *int64  `json:"OwnerAliUid"`
	AcrRepoName  *string `json:"AcrRepoName"`
	Status       *string `json:"Status"`
}

type shareDockerRepoJSONWire struct {
	Code           *string                      `json:"Code"`
	Message        *string                      `json:"Message"`
	RequestId      *string                      `json:"RequestId"`
	HttpStatusCode json.RawMessage              `json:"HttpStatusCode"`
	Success        *bool                        `json:"Success"`
	Data           *shareDockerRepoJSONWireData `json:"Data"`
}

type xmlShareDockerRepoResponseData struct {
	TargetAliUid int64  `xml:"TargetAliUid"`
	OwnerAliUid  int64  `xml:"OwnerAliUid"`
	AcrRepoName  string `xml:"AcrRepoName"`
	Status       string `xml:"Status"`
}

type xmlShareDockerRepoResponse struct {
	XMLName        xml.Name                        `xml:"ShareDockerRepoResponse"`
	RequestId      string                          `xml:"RequestId"`
	HttpStatusCode string                          `xml:"HttpStatusCode"`
	Code           string                          `xml:"Code"`
	Success        bool                            `xml:"Success"`
	Message        string                          `xml:"Message"`
	Data           *xmlShareDockerRepoResponseData `xml:"Data"`
}

func parseShareDockerRepoResponse(res map[string]interface{}) (*ShareDockerRepoResponse, error) {
	bodyStr, err := rawBodyStringFromMap(res)
	if err != nil {
		return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
	}
	out := &ShareDockerRepoResponse{Headers: make(map[string]*string)}
	parsed := &ShareDockerRepoResponseBody{}
	trimmed := strings.TrimSpace(bodyStr)
	if bodyStr != "" {
		if len(trimmed) > 0 && trimmed[0] == '<' {
			var xr xmlShareDockerRepoResponse
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
			if xr.Data != nil {
				parsed.Data = &ShareDockerRepoResponseBodyData{
					TargetAliUid: dara.Int64(xr.Data.TargetAliUid),
					OwnerAliUid:  dara.Int64(xr.Data.OwnerAliUid),
					AcrRepoName:  dara.String(xr.Data.AcrRepoName),
					Status:       dara.String(xr.Data.Status),
				}
			}
		} else {
			var wire shareDockerRepoJSONWire
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
			if wire.Data != nil {
				parsed.Data = &ShareDockerRepoResponseBodyData{
					TargetAliUid: wire.Data.TargetAliUid,
					OwnerAliUid:  wire.Data.OwnerAliUid,
					AcrRepoName:  wire.Data.AcrRepoName,
					Status:       wire.Data.Status,
				}
			}
		}
	}
	out.Body = parsed
	applyMapHeadersAndStatus(&out.Headers, &out.StatusCode, res)
	return out, nil
}

// --- UnshareDockerRepo ---

type unshareDockerRepoJSONWireData struct {
	Revoked *bool `json:"Revoked"`
}

type unshareDockerRepoJSONWire struct {
	Code           *string                        `json:"Code"`
	Message        *string                        `json:"Message"`
	RequestId      *string                        `json:"RequestId"`
	HttpStatusCode json.RawMessage                `json:"HttpStatusCode"`
	Success        *bool                          `json:"Success"`
	Data           *unshareDockerRepoJSONWireData `json:"Data"`
}

type xmlUnshareDockerRepoResponseData struct {
	Revoked string `xml:"Revoked"`
}

type xmlUnshareDockerRepoResponse struct {
	XMLName        xml.Name                          `xml:"UnshareDockerRepoResponse"`
	RequestId      string                            `xml:"RequestId"`
	HttpStatusCode string                            `xml:"HttpStatusCode"`
	Code           string                            `xml:"Code"`
	Success        bool                              `xml:"Success"`
	Message        string                            `xml:"Message"`
	Data           *xmlUnshareDockerRepoResponseData `xml:"Data"`
}

func parseUnshareDockerRepoResponse(res map[string]interface{}) (*UnshareDockerRepoResponse, error) {
	bodyStr, err := rawBodyStringFromMap(res)
	if err != nil {
		return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
	}
	out := &UnshareDockerRepoResponse{Headers: make(map[string]*string)}
	parsed := &UnshareDockerRepoResponseBody{}
	trimmed := strings.TrimSpace(bodyStr)
	if bodyStr != "" {
		if len(trimmed) > 0 && trimmed[0] == '<' {
			var xr xmlUnshareDockerRepoResponse
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
			if xr.Data != nil {
				if s := strings.TrimSpace(xr.Data.Revoked); s == "true" {
					parsed.Data = &UnshareDockerRepoResponseBodyData{Revoked: dara.Bool(true)}
				} else if s == "false" {
					parsed.Data = &UnshareDockerRepoResponseBodyData{Revoked: dara.Bool(false)}
				}
			}
		} else {
			var wire unshareDockerRepoJSONWire
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
			if wire.Data != nil {
				parsed.Data = &UnshareDockerRepoResponseBodyData{
					Revoked: wire.Data.Revoked,
				}
			}
		}
	}
	out.Body = parsed
	applyMapHeadersAndStatus(&out.Headers, &out.StatusCode, res)
	return out, nil
}

// --- ListSharedDockerRepos ---

type listSharedDockerReposJSONWireDataItem struct {
	PeerAliUid *int64  `json:"PeerAliUid"`
	Status     *string `json:"Status"`
}

type listSharedDockerReposJSONWire struct {
	Code           *string                                  `json:"Code"`
	Message        *string                                  `json:"Message"`
	RequestId      *string                                  `json:"RequestId"`
	HttpStatusCode json.RawMessage                          `json:"HttpStatusCode"`
	Success        *bool                                    `json:"Success"`
	Data           []*listSharedDockerReposJSONWireDataItem `json:"Data"`
}

type xmlListSharedDockerReposResponseDataItem struct {
	PeerAliUid int64  `xml:"PeerAliUid"`
	Status     string `xml:"Status"`
}

type xmlListSharedDockerReposResponse struct {
	XMLName        xml.Name                                    `xml:"ListSharedDockerReposResponse"`
	RequestId      string                                      `xml:"RequestId"`
	HttpStatusCode string                                      `xml:"HttpStatusCode"`
	Code           string                                      `xml:"Code"`
	Success        bool                                        `xml:"Success"`
	Message        string                                      `xml:"Message"`
	Data           []*xmlListSharedDockerReposResponseDataItem `xml:"Data>object"`
}

func parseListSharedDockerReposResponse(res map[string]interface{}) (*ListSharedDockerReposResponse, error) {
	bodyStr, err := rawBodyStringFromMap(res)
	if err != nil {
		return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
	}
	out := &ListSharedDockerReposResponse{Headers: make(map[string]*string)}
	parsed := &ListSharedDockerReposResponseBody{}
	trimmed := strings.TrimSpace(bodyStr)
	if bodyStr != "" {
		if len(trimmed) > 0 && trimmed[0] == '<' {
			var xr xmlListSharedDockerReposResponse
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
			for _, item := range xr.Data {
				if item != nil {
					parsed.Data = append(parsed.Data, &ListSharedDockerReposResponseBodyDataItem{
						PeerAliUid: dara.Int64(item.PeerAliUid),
						Status:     dara.String(item.Status),
					})
				}
			}
		} else {
			var wire listSharedDockerReposJSONWire
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
			for _, item := range wire.Data {
				if item != nil {
					parsed.Data = append(parsed.Data, &ListSharedDockerReposResponseBodyDataItem{
						PeerAliUid: item.PeerAliUid,
						Status:     item.Status,
					})
				}
			}
		}
	}
	if parsed.Data == nil {
		parsed.Data = []*ListSharedDockerReposResponseBodyDataItem{}
	}
	out.Body = parsed
	applyMapHeadersAndStatus(&out.Headers, &out.StatusCode, res)
	return out, nil
}

// parseCreateSimpleOfficeSiteResponse builds CreateSimpleOfficeSiteResponse from CallApi map.
// Data field is a string (OfficeSiteId). HttpStatusCode uses int32FromFlexibleJSON for flexible parsing.
func parseCreateSimpleOfficeSiteResponse(res map[string]interface{}) (*CreateSimpleOfficeSiteResponse, error) {
	out := &CreateSimpleOfficeSiteResponse{}
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
	parsed := &CreateSimpleOfficeSiteResponseBody{}
	if bodyStr != "" {
		trimmed := strings.TrimSpace(bodyStr)
		if len(trimmed) > 0 && trimmed[0] == '<' {
			// XML response
			var xmlResp xmlCreateSimpleOfficeSiteResponse
			if err := xml.Unmarshal([]byte(bodyStr), &xmlResp); err != nil {
				return nil, &ErrWithRequestID{Err: fmt.Errorf("XML parse error: %w", err), RequestID: extractRequestIDFromResponse(res)}
			}
			if xmlResp.RequestId != "" {
				parsed.RequestId = xmlResp.RequestId
			}
			if xmlResp.Data != nil {
				parsed.Data = xmlResp.Data
			}
			if xmlResp.Code != nil {
				parsed.Code = xmlResp.Code
			}
			if xmlResp.Success != nil {
				parsed.Success = xmlResp.Success
			}
			if xmlResp.Message != nil {
				parsed.Message = xmlResp.Message
			}
			if xmlResp.HttpStatusCode != nil {
				parsed.HttpStatusCode = xmlResp.HttpStatusCode
			}
		} else {
			// JSON response - use flexible parsing for HttpStatusCode
			type wireCreateSimpleOfficeSiteResponse struct {
				RequestId      string          `json:"RequestId"`
				HttpStatusCode json.RawMessage `json:"HttpStatusCode"`
				Data           *string         `json:"Data"`
				Code           *string         `json:"Code"`
				Success        *bool           `json:"Success"`
				Message        *string         `json:"Message"`
			}
			var wire wireCreateSimpleOfficeSiteResponse
			if err := json.Unmarshal([]byte(bodyStr), &wire); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}
			parsed.RequestId = wire.RequestId
			parsed.Data = wire.Data
			parsed.Code = wire.Code
			parsed.Success = wire.Success
			parsed.Message = wire.Message
			if len(wire.HttpStatusCode) > 0 {
				ht, derr := int32FromFlexibleJSON(wire.HttpStatusCode)
				if derr == nil {
					parsed.HttpStatusCode = ht
				}
			}
		}
	}
	out.Body = parsed
	applyMapHeadersAndStatus(&out.Headers, &out.StatusCode, res)
	return out, nil
}

// xmlCreateSimpleOfficeSiteResponse for XML response parsing
type xmlCreateSimpleOfficeSiteResponse struct {
	XMLName        xml.Name `xml:"CreateSimpleOfficeSiteResponse"`
	RequestId      string   `xml:"RequestId"`
	HttpStatusCode *int32   `xml:"HttpStatusCode"`
	Data           *string  `xml:"Data"`
	Code           *string  `xml:"Code"`
	Success        *bool    `xml:"Success"`
	Message        *string  `xml:"Message"`
}
