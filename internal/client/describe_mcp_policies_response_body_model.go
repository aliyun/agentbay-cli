// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMcpPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeMcpPoliciesResponseBody
	GetCode() *string
	SetData(v []*DescribeMcpPoliciesResponseBodyData) *DescribeMcpPoliciesResponseBody
	GetData() []*DescribeMcpPoliciesResponseBodyData
	SetHttpStatusCode(v int32) *DescribeMcpPoliciesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeMcpPoliciesResponseBody
	GetMessage() *string
	SetNextToken(v string) *DescribeMcpPoliciesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeMcpPoliciesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeMcpPoliciesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *DescribeMcpPoliciesResponseBody
	GetTotalCount() *int32
}

type DescribeMcpPoliciesResponseBody struct {
	Code           *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*DescribeMcpPoliciesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                                 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken      *string                                `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId      *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeMcpPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMcpPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMcpPoliciesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeMcpPoliciesResponseBody) GetData() []*DescribeMcpPoliciesResponseBodyData {
	return s.Data
}

func (s *DescribeMcpPoliciesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeMcpPoliciesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeMcpPoliciesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeMcpPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMcpPoliciesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeMcpPoliciesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeMcpPoliciesResponseBody) SetCode(v string) *DescribeMcpPoliciesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMcpPoliciesResponseBody) SetData(v []*DescribeMcpPoliciesResponseBodyData) *DescribeMcpPoliciesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMcpPoliciesResponseBody) SetHttpStatusCode(v int32) *DescribeMcpPoliciesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeMcpPoliciesResponseBody) SetMessage(v string) *DescribeMcpPoliciesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMcpPoliciesResponseBody) SetNextToken(v string) *DescribeMcpPoliciesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeMcpPoliciesResponseBody) SetRequestId(v string) *DescribeMcpPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMcpPoliciesResponseBody) SetSuccess(v bool) *DescribeMcpPoliciesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMcpPoliciesResponseBody) SetTotalCount(v int32) *DescribeMcpPoliciesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeMcpPoliciesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeMcpPoliciesResponseBodyData struct {
	ApiKeyIds         []*string                                         `json:"ApiKeyIds,omitempty" xml:"ApiKeyIds,omitempty" type:"Repeated"`
	Concurrency       *int32                                            `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	CreateTime        *string                                           `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DesktopMaxRuntime *int32                                            `json:"DesktopMaxRuntime,omitempty" xml:"DesktopMaxRuntime,omitempty"`
	DisplayConfig     *DescribeMcpPoliciesResponseBodyDataDisplayConfig `json:"DisplayConfig,omitempty" xml:"DisplayConfig,omitempty" type:"Struct"`
	IdleTimeoutSwitch *bool                                             `json:"IdleTimeoutSwitch,omitempty" xml:"IdleTimeoutSwitch,omitempty"`
	McpIdleTimeout    *int32                                            `json:"McpIdleTimeout,omitempty" xml:"McpIdleTimeout,omitempty"`
	NetworkConfig     *DescribeMcpPoliciesResponseBodyDataNetworkConfig `json:"NetworkConfig,omitempty" xml:"NetworkConfig,omitempty" type:"Struct"`
	PolicyId          *string                                           `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	PolicyName        *string                                           `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	UserIdleTimeout   *int32                                            `json:"UserIdleTimeout,omitempty" xml:"UserIdleTimeout,omitempty"`
}

func (s DescribeMcpPoliciesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeMcpPoliciesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMcpPoliciesResponseBodyData) GetApiKeyIds() []*string {
	return s.ApiKeyIds
}

func (s *DescribeMcpPoliciesResponseBodyData) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *DescribeMcpPoliciesResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeMcpPoliciesResponseBodyData) GetDesktopMaxRuntime() *int32 {
	return s.DesktopMaxRuntime
}

func (s *DescribeMcpPoliciesResponseBodyData) GetDisplayConfig() *DescribeMcpPoliciesResponseBodyDataDisplayConfig {
	return s.DisplayConfig
}

func (s *DescribeMcpPoliciesResponseBodyData) GetIdleTimeoutSwitch() *bool {
	return s.IdleTimeoutSwitch
}

func (s *DescribeMcpPoliciesResponseBodyData) GetMcpIdleTimeout() *int32 {
	return s.McpIdleTimeout
}

func (s *DescribeMcpPoliciesResponseBodyData) GetNetworkConfig() *DescribeMcpPoliciesResponseBodyDataNetworkConfig {
	return s.NetworkConfig
}

func (s *DescribeMcpPoliciesResponseBodyData) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DescribeMcpPoliciesResponseBodyData) GetPolicyName() *string {
	return s.PolicyName
}

func (s *DescribeMcpPoliciesResponseBodyData) GetUserIdleTimeout() *int32 {
	return s.UserIdleTimeout
}

func (s *DescribeMcpPoliciesResponseBodyData) SetApiKeyIds(v []*string) *DescribeMcpPoliciesResponseBodyData {
	s.ApiKeyIds = v
	return s
}

func (s *DescribeMcpPoliciesResponseBodyData) SetConcurrency(v int32) *DescribeMcpPoliciesResponseBodyData {
	s.Concurrency = &v
	return s
}

func (s *DescribeMcpPoliciesResponseBodyData) SetCreateTime(v string) *DescribeMcpPoliciesResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeMcpPoliciesResponseBodyData) SetDesktopMaxRuntime(v int32) *DescribeMcpPoliciesResponseBodyData {
	s.DesktopMaxRuntime = &v
	return s
}

func (s *DescribeMcpPoliciesResponseBodyData) SetDisplayConfig(v *DescribeMcpPoliciesResponseBodyDataDisplayConfig) *DescribeMcpPoliciesResponseBodyData {
	s.DisplayConfig = v
	return s
}

func (s *DescribeMcpPoliciesResponseBodyData) SetIdleTimeoutSwitch(v bool) *DescribeMcpPoliciesResponseBodyData {
	s.IdleTimeoutSwitch = &v
	return s
}

func (s *DescribeMcpPoliciesResponseBodyData) SetMcpIdleTimeout(v int32) *DescribeMcpPoliciesResponseBodyData {
	s.McpIdleTimeout = &v
	return s
}

func (s *DescribeMcpPoliciesResponseBodyData) SetNetworkConfig(v *DescribeMcpPoliciesResponseBodyDataNetworkConfig) *DescribeMcpPoliciesResponseBodyData {
	s.NetworkConfig = v
	return s
}

func (s *DescribeMcpPoliciesResponseBodyData) SetPolicyId(v string) *DescribeMcpPoliciesResponseBodyData {
	s.PolicyId = &v
	return s
}

func (s *DescribeMcpPoliciesResponseBodyData) SetPolicyName(v string) *DescribeMcpPoliciesResponseBodyData {
	s.PolicyName = &v
	return s
}

func (s *DescribeMcpPoliciesResponseBodyData) SetUserIdleTimeout(v int32) *DescribeMcpPoliciesResponseBodyData {
	s.UserIdleTimeout = &v
	return s
}

func (s *DescribeMcpPoliciesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeMcpPoliciesResponseBodyDataDisplayConfig struct {
	DisplayMode *string `json:"DisplayMode,omitempty" xml:"DisplayMode,omitempty"`
}

func (s DescribeMcpPoliciesResponseBodyDataDisplayConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeMcpPoliciesResponseBodyDataDisplayConfig) GoString() string {
	return s.String()
}

func (s *DescribeMcpPoliciesResponseBodyDataDisplayConfig) GetDisplayMode() *string {
	return s.DisplayMode
}

func (s *DescribeMcpPoliciesResponseBodyDataDisplayConfig) SetDisplayMode(v string) *DescribeMcpPoliciesResponseBodyDataDisplayConfig {
	s.DisplayMode = &v
	return s
}

func (s *DescribeMcpPoliciesResponseBodyDataDisplayConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeMcpPoliciesResponseBodyDataNetworkConfig struct {
	Enabled         *bool                                                         `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	GlobalConfig    *DescribeMcpPoliciesResponseBodyDataNetworkConfigGlobalConfig `json:"GlobalConfig,omitempty" xml:"GlobalConfig,omitempty" type:"Struct"`
	Mode            *string                                                       `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Name            *string                                                       `json:"Name,omitempty" xml:"Name,omitempty"`
	NetworkPolicyId *string                                                       `json:"NetworkPolicyId,omitempty" xml:"NetworkPolicyId,omitempty"`
}

func (s DescribeMcpPoliciesResponseBodyDataNetworkConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeMcpPoliciesResponseBodyDataNetworkConfig) GoString() string {
	return s.String()
}

func (s *DescribeMcpPoliciesResponseBodyDataNetworkConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeMcpPoliciesResponseBodyDataNetworkConfig) GetGlobalConfig() *DescribeMcpPoliciesResponseBodyDataNetworkConfigGlobalConfig {
	return s.GlobalConfig
}

func (s *DescribeMcpPoliciesResponseBodyDataNetworkConfig) GetMode() *string {
	return s.Mode
}

func (s *DescribeMcpPoliciesResponseBodyDataNetworkConfig) GetName() *string {
	return s.Name
}

func (s *DescribeMcpPoliciesResponseBodyDataNetworkConfig) GetNetworkPolicyId() *string {
	return s.NetworkPolicyId
}

func (s *DescribeMcpPoliciesResponseBodyDataNetworkConfig) SetEnabled(v bool) *DescribeMcpPoliciesResponseBodyDataNetworkConfig {
	s.Enabled = &v
	return s
}

func (s *DescribeMcpPoliciesResponseBodyDataNetworkConfig) SetGlobalConfig(v *DescribeMcpPoliciesResponseBodyDataNetworkConfigGlobalConfig) *DescribeMcpPoliciesResponseBodyDataNetworkConfig {
	s.GlobalConfig = v
	return s
}

func (s *DescribeMcpPoliciesResponseBodyDataNetworkConfig) SetMode(v string) *DescribeMcpPoliciesResponseBodyDataNetworkConfig {
	s.Mode = &v
	return s
}

func (s *DescribeMcpPoliciesResponseBodyDataNetworkConfig) SetName(v string) *DescribeMcpPoliciesResponseBodyDataNetworkConfig {
	s.Name = &v
	return s
}

func (s *DescribeMcpPoliciesResponseBodyDataNetworkConfig) SetNetworkPolicyId(v string) *DescribeMcpPoliciesResponseBodyDataNetworkConfig {
	s.NetworkPolicyId = &v
	return s
}

func (s *DescribeMcpPoliciesResponseBodyDataNetworkConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeMcpPoliciesResponseBodyDataNetworkConfigGlobalConfig struct {
	ProxyConfig *DescribeMcpPoliciesResponseBodyDataNetworkConfigGlobalConfigProxyConfig `json:"ProxyConfig,omitempty" xml:"ProxyConfig,omitempty" type:"Struct"`
}

func (s DescribeMcpPoliciesResponseBodyDataNetworkConfigGlobalConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeMcpPoliciesResponseBodyDataNetworkConfigGlobalConfig) GoString() string {
	return s.String()
}

func (s *DescribeMcpPoliciesResponseBodyDataNetworkConfigGlobalConfig) GetProxyConfig() *DescribeMcpPoliciesResponseBodyDataNetworkConfigGlobalConfigProxyConfig {
	return s.ProxyConfig
}

func (s *DescribeMcpPoliciesResponseBodyDataNetworkConfigGlobalConfig) SetProxyConfig(v *DescribeMcpPoliciesResponseBodyDataNetworkConfigGlobalConfigProxyConfig) *DescribeMcpPoliciesResponseBodyDataNetworkConfigGlobalConfig {
	s.ProxyConfig = v
	return s
}

func (s *DescribeMcpPoliciesResponseBodyDataNetworkConfigGlobalConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeMcpPoliciesResponseBodyDataNetworkConfigGlobalConfigProxyConfig struct {
	Address         *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Password        *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Port            *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	ProxyType       *string `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
	RedirectChannel *string `json:"RedirectChannel,omitempty" xml:"RedirectChannel,omitempty"`
	Username        *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s DescribeMcpPoliciesResponseBodyDataNetworkConfigGlobalConfigProxyConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeMcpPoliciesResponseBodyDataNetworkConfigGlobalConfigProxyConfig) GoString() string {
	return s.String()
}

func (s *DescribeMcpPoliciesResponseBodyDataNetworkConfigGlobalConfigProxyConfig) GetAddress() *string {
	return s.Address
}

func (s *DescribeMcpPoliciesResponseBodyDataNetworkConfigGlobalConfigProxyConfig) GetPassword() *string {
	return s.Password
}

func (s *DescribeMcpPoliciesResponseBodyDataNetworkConfigGlobalConfigProxyConfig) GetPort() *int32 {
	return s.Port
}

func (s *DescribeMcpPoliciesResponseBodyDataNetworkConfigGlobalConfigProxyConfig) GetProxyType() *string {
	return s.ProxyType
}

func (s *DescribeMcpPoliciesResponseBodyDataNetworkConfigGlobalConfigProxyConfig) GetRedirectChannel() *string {
	return s.RedirectChannel
}

func (s *DescribeMcpPoliciesResponseBodyDataNetworkConfigGlobalConfigProxyConfig) GetUsername() *string {
	return s.Username
}

func (s *DescribeMcpPoliciesResponseBodyDataNetworkConfigGlobalConfigProxyConfig) SetAddress(v string) *DescribeMcpPoliciesResponseBodyDataNetworkConfigGlobalConfigProxyConfig {
	s.Address = &v
	return s
}

func (s *DescribeMcpPoliciesResponseBodyDataNetworkConfigGlobalConfigProxyConfig) SetPassword(v string) *DescribeMcpPoliciesResponseBodyDataNetworkConfigGlobalConfigProxyConfig {
	s.Password = &v
	return s
}

func (s *DescribeMcpPoliciesResponseBodyDataNetworkConfigGlobalConfigProxyConfig) SetPort(v int32) *DescribeMcpPoliciesResponseBodyDataNetworkConfigGlobalConfigProxyConfig {
	s.Port = &v
	return s
}

func (s *DescribeMcpPoliciesResponseBodyDataNetworkConfigGlobalConfigProxyConfig) SetProxyType(v string) *DescribeMcpPoliciesResponseBodyDataNetworkConfigGlobalConfigProxyConfig {
	s.ProxyType = &v
	return s
}

func (s *DescribeMcpPoliciesResponseBodyDataNetworkConfigGlobalConfigProxyConfig) SetRedirectChannel(v string) *DescribeMcpPoliciesResponseBodyDataNetworkConfigGlobalConfigProxyConfig {
	s.RedirectChannel = &v
	return s
}

func (s *DescribeMcpPoliciesResponseBodyDataNetworkConfigGlobalConfigProxyConfig) SetUsername(v string) *DescribeMcpPoliciesResponseBodyDataNetworkConfigGlobalConfigProxyConfig {
	s.Username = &v
	return s
}

func (s *DescribeMcpPoliciesResponseBodyDataNetworkConfigGlobalConfigProxyConfig) Validate() error {
	return dara.Validate(s)
}
