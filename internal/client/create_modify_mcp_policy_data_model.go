// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

// CreateModifyMcpPolicyDataRequest - Create/Modify 策略配置数据的公共请求结构
// Used by both CreateMcpPolicyData and ModifyMcpPolicyData APIs (same parameter structure, different Action).
type CreateModifyMcpPolicyDataRequest struct {
	ImageId                       *string           `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	SandboxLifeCycle              *SandboxLifeCycle `json:"SandboxLifeCycle,omitempty" xml:"SandboxLifeCycle,omitempty"`
	NetworkConfig                 *NetworkConfig    `json:"NetworkConfig,omitempty" xml:"NetworkConfig,omitempty"`
	DisplayConfig                 *DisplayConfig    `json:"DisplayConfig,omitempty" xml:"DisplayConfig,omitempty"`
	Taskbar                       *string           `json:"Taskbar,omitempty" xml:"Taskbar,omitempty"`
	ScreenDisplayMode             *string           `json:"ScreenDisplayMode,omitempty" xml:"ScreenDisplayMode,omitempty"`
	ClientControlMenu             *string           `json:"ClientControlMenu,omitempty" xml:"ClientControlMenu,omitempty"`
	BusinessType                  *int32            `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	ResourceType                  *string           `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	DisconnectKeepSession         *string           `json:"DisconnectKeepSession,omitempty" xml:"DisconnectKeepSession,omitempty"`
	Name                          *string           `json:"Name,omitempty" xml:"Name,omitempty"`
	InternetCommunicationProtocol *string           `json:"InternetCommunicationProtocol,omitempty" xml:"InternetCommunicationProtocol,omitempty"`
	ResolutionWidth               *int32            `json:"ResolutionWidth,omitempty" xml:"ResolutionWidth,omitempty"`
	ResolutionHeight              *int32            `json:"ResolutionHeight,omitempty" xml:"ResolutionHeight,omitempty"`
	RegionName                    *string           `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s *CreateModifyMcpPolicyDataRequest) String() string {
	return dara.Prettify(s)
}

func (s *CreateModifyMcpPolicyDataRequest) GoString() string {
	return s.String()
}

func (s *CreateModifyMcpPolicyDataRequest) Validate() error {
	return dara.Validate(s)
}

// ---------- CreateMcpPolicyData Response ----------

// CreateMcpPolicyDataResponseBody - 响应体
// PolicyId is the edsPolicyId string returned by the server.
type CreateMcpPolicyDataResponseBody struct {
	RequestId      string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	PolicyId       *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s *CreateMcpPolicyDataResponseBody) GetPolicyId() *string {
	return s.PolicyId
}

func (s *CreateMcpPolicyDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s *CreateMcpPolicyDataResponseBody) GetRequestId() *string {
	return &s.RequestId
}

func (s *CreateMcpPolicyDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMcpPolicyDataResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateMcpPolicyDataResponseBody) GetMessage() *string {
	return s.Message
}

// CreateMcpPolicyDataResponse - 响应
type CreateMcpPolicyDataResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMcpPolicyDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
	RawBody    string                           `json:"-"`
}

func (s CreateMcpPolicyDataResponse) String() string {
	return dara.Prettify(s)
}

func (s *CreateMcpPolicyDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMcpPolicyDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMcpPolicyDataResponse) GetBody() *CreateMcpPolicyDataResponseBody {
	return s.Body
}

// ---------- ModifyMcpPolicyData Response ----------

// ModifyMcpPolicyDataResponseBody - 响应体 (same structure as SaveMcpPolicyData)
type ModifyMcpPolicyDataResponseBody struct {
	RequestId      string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Data           *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s *ModifyMcpPolicyDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s *ModifyMcpPolicyDataResponseBody) GetRequestId() *string {
	return &s.RequestId
}

func (s *ModifyMcpPolicyDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyMcpPolicyDataResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyMcpPolicyDataResponseBody) GetMessage() *string {
	return s.Message
}

// ModifyMcpPolicyDataResponse - 响应
type ModifyMcpPolicyDataResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyMcpPolicyDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
	RawBody    string                           `json:"-"`
}

func (s ModifyMcpPolicyDataResponse) String() string {
	return dara.Prettify(s)
}

func (s *ModifyMcpPolicyDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyMcpPolicyDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyMcpPolicyDataResponse) GetBody() *ModifyMcpPolicyDataResponseBody {
	return s.Body
}
