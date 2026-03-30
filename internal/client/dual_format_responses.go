// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"encoding/xml"
	"errors"
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

// --- GetDockerfileTemplate ---

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
			if err := json.Unmarshal([]byte(bodyStr), parsed); err != nil {
				return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
			}
		}
		out.Body = parsed
	}
	applyMapHeadersAndStatus(&out.Headers, &out.StatusCode, res)
	return out, nil
}
