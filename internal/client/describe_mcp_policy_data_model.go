// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

// DescribeMcpPolicyDataRequest - 查询策略配置数据请求
type DescribeMcpPolicyDataRequest struct {
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
}

func (s *DescribeMcpPolicyDataRequest) String() string {
	return dara.Prettify(s)
}

func (s *DescribeMcpPolicyDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeMcpPolicyDataRequest) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeMcpPolicyDataRequest) SetImageId(v string) *DescribeMcpPolicyDataRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeMcpPolicyDataRequest) Validate() error {
	return dara.Validate(s)
}

// GroupSpec - 组规格
type GroupSpec struct {
	AppInstanceType *string `json:"AppInstanceType,omitempty" xml:"AppInstanceType,omitempty"`
	RegionName      *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	Memory          *int32  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Cpu             *int32  `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s *GroupSpec) String() string {
	return dara.Prettify(s)
}

// ScreenSettings - 屏幕设置
type ScreenSettings struct {
	ClientControlMenu *string `json:"ClientControlMenu,omitempty" xml:"ClientControlMenu,omitempty"`
	ScreenDisplayMode *string `json:"ScreenDisplayMode,omitempty" xml:"ScreenDisplayMode,omitempty"`
	Taskbar           *string `json:"Taskbar,omitempty" xml:"Taskbar,omitempty"`
	KioskModeEnabled  *bool   `json:"KioskModeEnabled,omitempty" xml:"KioskModeEnabled,omitempty"`
}

func (s *ScreenSettings) String() string {
	return dara.Prettify(s)
}

// SandboxLifeCycle - 沙箱生命周期
type SandboxLifeCycle struct {
	IdleTimeoutSwitch *bool   `json:"IdleTimeoutSwitch,omitempty" xml:"IdleTimeoutSwitch,omitempty"`
	HibernateTimeout  *float64 `json:"HibernateTimeout,omitempty" xml:"HibernateTimeout,omitempty"`
	DesktopMaxRuntime *float64 `json:"DesktopMaxRuntime,omitempty" xml:"DesktopMaxRuntime,omitempty"`
	UserIdleTimeout   *float64 `json:"UserIdleTimeout,omitempty" xml:"UserIdleTimeout,omitempty"`
}

func (s *SandboxLifeCycle) String() string {
	return dara.Prettify(s)
}

// NetworkConfig - 网络配置
type NetworkConfig struct {
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s *NetworkConfig) String() string {
	return dara.Prettify(s)
}

// DisplayConfig - 显示配置
type DisplayConfig struct {
	DisplayMode *string `json:"DisplayMode,omitempty" xml:"DisplayMode,omitempty"`
}

func (s *DisplayConfig) String() string {
	return dara.Prettify(s)
}

// NetworkData - 网络数据
type NetworkData struct {
	VpcId            *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	OfficeSiteType   *string `json:"OfficeSiteType,omitempty" xml:"OfficeSiteType,omitempty"`
	DnsAddress       *string `json:"DnsAddress,omitempty" xml:"DnsAddress,omitempty"`
	VpcName          *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
	SessionBandwidth *int32  `json:"SessionBandwidth,omitempty" xml:"SessionBandwidth,omitempty"`
}

func (s *NetworkData) String() string {
	return dara.Prettify(s)
}

// DescribeMcpPolicyDataResponseBodyData - 响应数据
type DescribeMcpPolicyDataResponseBodyData struct {
	IsDefaultData    *bool             `json:"IsDefaultData,omitempty" xml:"IsDefaultData,omitempty"`
	GroupSpec        *GroupSpec        `json:"GroupSpec,omitempty" xml:"GroupSpec,omitempty"`
	ScreenSettings   *ScreenSettings   `json:"ScreenSettings,omitempty" xml:"ScreenSettings,omitempty"`
	ImageId          *string           `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	SandboxLifeCycle *SandboxLifeCycle `json:"SandboxLifeCycle,omitempty" xml:"SandboxLifeCycle,omitempty"`
	NetworkConfig    *NetworkConfig    `json:"NetworkConfig,omitempty" xml:"NetworkConfig,omitempty"`
	DisplayConfig    *DisplayConfig    `json:"DisplayConfig,omitempty" xml:"DisplayConfig,omitempty"`
	NetworkData      *NetworkData      `json:"NetworkData,omitempty" xml:"NetworkData,omitempty"`
	AliUid           *int64            `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	PolicyId         *string           `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s *DescribeMcpPolicyDataResponseBodyData) String() string {
	return dara.Prettify(s)
}

// DescribeMcpPolicyDataResponseBody - 响应体
type DescribeMcpPolicyDataResponseBody struct {
	RequestId      string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Data           *DescribeMcpPolicyDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty"`
	Code           *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s *DescribeMcpPolicyDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s *DescribeMcpPolicyDataResponseBody) GetRequestId() *string {
	return &s.RequestId
}

func (s *DescribeMcpPolicyDataResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeMcpPolicyDataResponseBody) GetData() *DescribeMcpPolicyDataResponseBodyData {
	return s.Data
}

func (s *DescribeMcpPolicyDataResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeMcpPolicyDataResponseBody) GetSuccess() *bool {
	return s.Success
}

// DescribeMcpPolicyDataResponse - 响应
type DescribeMcpPolicyDataResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMcpPolicyDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
	RawBody    string                           `json:"-"` // Store raw body for debugging
}

func (s DescribeMcpPolicyDataResponse) String() string {
	return dara.Prettify(s)
}

func (s *DescribeMcpPolicyDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMcpPolicyDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMcpPolicyDataResponse) GetBody() *DescribeMcpPolicyDataResponseBody {
	return s.Body
}
