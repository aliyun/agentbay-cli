// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

// SaveMcpPolicyDataRequest - 保存策略配置数据请求
type SaveMcpPolicyDataRequest struct {
	ImageId          *string           `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	PolicyId         *string           `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	GroupSpec        *GroupSpec        `json:"GroupSpec,omitempty" xml:"GroupSpec,omitempty"`
	SandboxLifeCycle *SandboxLifeCycle `json:"SandboxLifeCycle,omitempty" xml:"SandboxLifeCycle,omitempty"`
	NetworkData      *NetworkData      `json:"NetworkData,omitempty" xml:"NetworkData,omitempty"`
	ScreenSettings   *ScreenSettings   `json:"ScreenSettings,omitempty" xml:"ScreenSettings,omitempty"`
	NetworkConfig    *NetworkConfig    `json:"NetworkConfig,omitempty" xml:"NetworkConfig,omitempty"`
	DisplayConfig    *DisplayConfig    `json:"DisplayConfig,omitempty" xml:"DisplayConfig,omitempty"`
	RegionId         *string           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s *SaveMcpPolicyDataRequest) String() string {
	return dara.Prettify(s)
}

func (s *SaveMcpPolicyDataRequest) GoString() string {
	return s.String()
}

func (s *SaveMcpPolicyDataRequest) GetImageId() *string {
	return s.ImageId
}

func (s *SaveMcpPolicyDataRequest) SetImageId(v string) *SaveMcpPolicyDataRequest {
	s.ImageId = &v
	return s
}

func (s *SaveMcpPolicyDataRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *SaveMcpPolicyDataRequest) SetPolicyId(v string) *SaveMcpPolicyDataRequest {
	s.PolicyId = &v
	return s
}

func (s *SaveMcpPolicyDataRequest) GetGroupSpec() *GroupSpec {
	return s.GroupSpec
}

func (s *SaveMcpPolicyDataRequest) SetGroupSpec(v *GroupSpec) *SaveMcpPolicyDataRequest {
	s.GroupSpec = v
	return s
}

func (s *SaveMcpPolicyDataRequest) GetSandboxLifeCycle() *SandboxLifeCycle {
	return s.SandboxLifeCycle
}

func (s *SaveMcpPolicyDataRequest) SetSandboxLifeCycle(v *SandboxLifeCycle) *SaveMcpPolicyDataRequest {
	s.SandboxLifeCycle = v
	return s
}

func (s *SaveMcpPolicyDataRequest) GetNetworkData() *NetworkData {
	return s.NetworkData
}

func (s *SaveMcpPolicyDataRequest) SetNetworkData(v *NetworkData) *SaveMcpPolicyDataRequest {
	s.NetworkData = v
	return s
}

func (s *SaveMcpPolicyDataRequest) GetScreenSettings() *ScreenSettings {
	return s.ScreenSettings
}

func (s *SaveMcpPolicyDataRequest) SetScreenSettings(v *ScreenSettings) *SaveMcpPolicyDataRequest {
	s.ScreenSettings = v
	return s
}

func (s *SaveMcpPolicyDataRequest) GetNetworkConfig() *NetworkConfig {
	return s.NetworkConfig
}

func (s *SaveMcpPolicyDataRequest) SetNetworkConfig(v *NetworkConfig) *SaveMcpPolicyDataRequest {
	s.NetworkConfig = v
	return s
}

func (s *SaveMcpPolicyDataRequest) GetDisplayConfig() *DisplayConfig {
	return s.DisplayConfig
}

func (s *SaveMcpPolicyDataRequest) SetDisplayConfig(v *DisplayConfig) *SaveMcpPolicyDataRequest {
	s.DisplayConfig = v
	return s
}

func (s *SaveMcpPolicyDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SaveMcpPolicyDataRequest) SetRegionId(v string) *SaveMcpPolicyDataRequest {
	s.RegionId = &v
	return s
}

func (s *SaveMcpPolicyDataRequest) Validate() error {
	return dara.Validate(s)
}

// SaveMcpPolicyDataResponseBody - 响应体
// Note: Data is a boolean indicating success
type SaveMcpPolicyDataResponseBody struct {
	RequestId      string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Data           *bool    `json:"Data,omitempty" xml:"Data,omitempty"`
	Code           *string  `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool    `json:"Success,omitempty" xml:"Success,omitempty"`
	Message        *string  `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s *SaveMcpPolicyDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s *SaveMcpPolicyDataResponseBody) GetRequestId() *string {
	return &s.RequestId
}

func (s *SaveMcpPolicyDataResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SaveMcpPolicyDataResponseBody) GetData() *bool {
	return s.Data
}

func (s *SaveMcpPolicyDataResponseBody) GetCode() *string {
	return s.Code
}

func (s *SaveMcpPolicyDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveMcpPolicyDataResponseBody) GetMessage() *string {
	return s.Message
}

// SaveMcpPolicyDataResponse - 响应
type SaveMcpPolicyDataResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveMcpPolicyDataResponseBody     `json:"body,omitempty" xml:"body,omitempty"`
	RawBody    string                             `json:"-"` // Store raw body for debugging
}

func (s SaveMcpPolicyDataResponse) String() string {
	return dara.Prettify(s)
}

func (s *SaveMcpPolicyDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveMcpPolicyDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveMcpPolicyDataResponse) GetBody() *SaveMcpPolicyDataResponseBody {
	return s.Body
}
