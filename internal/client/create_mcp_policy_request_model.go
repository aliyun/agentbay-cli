// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcpPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConcurrency(v int32) *CreateMcpPolicyRequest
	GetConcurrency() *int32
	SetDebugMode(v bool) *CreateMcpPolicyRequest
	GetDebugMode() *bool
	SetDesktopMaxRuntime(v int32) *CreateMcpPolicyRequest
	GetDesktopMaxRuntime() *int32
	SetDisplayConfig(v *CreateMcpPolicyRequestDisplayConfig) *CreateMcpPolicyRequest
	GetDisplayConfig() *CreateMcpPolicyRequestDisplayConfig
	SetIdleTimeoutSwitch(v bool) *CreateMcpPolicyRequest
	GetIdleTimeoutSwitch() *bool
	SetMcpIdleTimeout(v int32) *CreateMcpPolicyRequest
	GetMcpIdleTimeout() *int32
	SetNetworkConfig(v *CreateMcpPolicyRequestNetworkConfig) *CreateMcpPolicyRequest
	GetNetworkConfig() *CreateMcpPolicyRequestNetworkConfig
	SetPolicyName(v string) *CreateMcpPolicyRequest
	GetPolicyName() *string
	SetUserIdleTimeout(v int32) *CreateMcpPolicyRequest
	GetUserIdleTimeout() *int32
}

type CreateMcpPolicyRequest struct {
	Concurrency       *int32                               `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	DebugMode         *bool                                `json:"DebugMode,omitempty" xml:"DebugMode,omitempty"`
	DesktopMaxRuntime *int32                               `json:"DesktopMaxRuntime,omitempty" xml:"DesktopMaxRuntime,omitempty"`
	DisplayConfig     *CreateMcpPolicyRequestDisplayConfig `json:"DisplayConfig,omitempty" xml:"DisplayConfig,omitempty" type:"Struct"`
	IdleTimeoutSwitch *bool                                `json:"IdleTimeoutSwitch,omitempty" xml:"IdleTimeoutSwitch,omitempty"`
	McpIdleTimeout    *int32                               `json:"McpIdleTimeout,omitempty" xml:"McpIdleTimeout,omitempty"`
	NetworkConfig     *CreateMcpPolicyRequestNetworkConfig `json:"NetworkConfig,omitempty" xml:"NetworkConfig,omitempty" type:"Struct"`
	PolicyName        *string                              `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	UserIdleTimeout   *int32                               `json:"UserIdleTimeout,omitempty" xml:"UserIdleTimeout,omitempty"`
}

func (s CreateMcpPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateMcpPolicyRequest) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *CreateMcpPolicyRequest) GetDebugMode() *bool {
	return s.DebugMode
}

func (s *CreateMcpPolicyRequest) GetDesktopMaxRuntime() *int32 {
	return s.DesktopMaxRuntime
}

func (s *CreateMcpPolicyRequest) GetDisplayConfig() *CreateMcpPolicyRequestDisplayConfig {
	return s.DisplayConfig
}

func (s *CreateMcpPolicyRequest) GetIdleTimeoutSwitch() *bool {
	return s.IdleTimeoutSwitch
}

func (s *CreateMcpPolicyRequest) GetMcpIdleTimeout() *int32 {
	return s.McpIdleTimeout
}

func (s *CreateMcpPolicyRequest) GetNetworkConfig() *CreateMcpPolicyRequestNetworkConfig {
	return s.NetworkConfig
}

func (s *CreateMcpPolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *CreateMcpPolicyRequest) GetUserIdleTimeout() *int32 {
	return s.UserIdleTimeout
}

func (s *CreateMcpPolicyRequest) SetConcurrency(v int32) *CreateMcpPolicyRequest {
	s.Concurrency = &v
	return s
}

func (s *CreateMcpPolicyRequest) SetDebugMode(v bool) *CreateMcpPolicyRequest {
	s.DebugMode = &v
	return s
}

func (s *CreateMcpPolicyRequest) SetDesktopMaxRuntime(v int32) *CreateMcpPolicyRequest {
	s.DesktopMaxRuntime = &v
	return s
}

func (s *CreateMcpPolicyRequest) SetDisplayConfig(v *CreateMcpPolicyRequestDisplayConfig) *CreateMcpPolicyRequest {
	s.DisplayConfig = v
	return s
}

func (s *CreateMcpPolicyRequest) SetIdleTimeoutSwitch(v bool) *CreateMcpPolicyRequest {
	s.IdleTimeoutSwitch = &v
	return s
}

func (s *CreateMcpPolicyRequest) SetMcpIdleTimeout(v int32) *CreateMcpPolicyRequest {
	s.McpIdleTimeout = &v
	return s
}

func (s *CreateMcpPolicyRequest) SetNetworkConfig(v *CreateMcpPolicyRequestNetworkConfig) *CreateMcpPolicyRequest {
	s.NetworkConfig = v
	return s
}

func (s *CreateMcpPolicyRequest) SetPolicyName(v string) *CreateMcpPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *CreateMcpPolicyRequest) SetUserIdleTimeout(v int32) *CreateMcpPolicyRequest {
	s.UserIdleTimeout = &v
	return s
}

func (s *CreateMcpPolicyRequest) Validate() error {
	return dara.Validate(s)
}

type CreateMcpPolicyRequestDisplayConfig struct {
	DisplayMode *string `json:"DisplayMode,omitempty" xml:"DisplayMode,omitempty"`
}

func (s CreateMcpPolicyRequestDisplayConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpPolicyRequestDisplayConfig) GoString() string {
	return s.String()
}

func (s *CreateMcpPolicyRequestDisplayConfig) GetDisplayMode() *string {
	return s.DisplayMode
}

func (s *CreateMcpPolicyRequestDisplayConfig) SetDisplayMode(v string) *CreateMcpPolicyRequestDisplayConfig {
	s.DisplayMode = &v
	return s
}

func (s *CreateMcpPolicyRequestDisplayConfig) Validate() error {
	return dara.Validate(s)
}

type CreateMcpPolicyRequestNetworkConfig struct {
	Enabled         *bool                                            `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	GlobalConfig    *CreateMcpPolicyRequestNetworkConfigGlobalConfig `json:"GlobalConfig,omitempty" xml:"GlobalConfig,omitempty" type:"Struct"`
	Mode            *string                                          `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Rules           []*CreateMcpPolicyRequestNetworkConfigRules      `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	NetworkPolicyId *string                                          `json:"networkPolicyId,omitempty" xml:"networkPolicyId,omitempty"`
}

func (s CreateMcpPolicyRequestNetworkConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpPolicyRequestNetworkConfig) GoString() string {
	return s.String()
}

func (s *CreateMcpPolicyRequestNetworkConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateMcpPolicyRequestNetworkConfig) GetGlobalConfig() *CreateMcpPolicyRequestNetworkConfigGlobalConfig {
	return s.GlobalConfig
}

func (s *CreateMcpPolicyRequestNetworkConfig) GetMode() *string {
	return s.Mode
}

func (s *CreateMcpPolicyRequestNetworkConfig) GetRules() []*CreateMcpPolicyRequestNetworkConfigRules {
	return s.Rules
}

func (s *CreateMcpPolicyRequestNetworkConfig) GetNetworkPolicyId() *string {
	return s.NetworkPolicyId
}

func (s *CreateMcpPolicyRequestNetworkConfig) SetEnabled(v bool) *CreateMcpPolicyRequestNetworkConfig {
	s.Enabled = &v
	return s
}

func (s *CreateMcpPolicyRequestNetworkConfig) SetGlobalConfig(v *CreateMcpPolicyRequestNetworkConfigGlobalConfig) *CreateMcpPolicyRequestNetworkConfig {
	s.GlobalConfig = v
	return s
}

func (s *CreateMcpPolicyRequestNetworkConfig) SetMode(v string) *CreateMcpPolicyRequestNetworkConfig {
	s.Mode = &v
	return s
}

func (s *CreateMcpPolicyRequestNetworkConfig) SetRules(v []*CreateMcpPolicyRequestNetworkConfigRules) *CreateMcpPolicyRequestNetworkConfig {
	s.Rules = v
	return s
}

func (s *CreateMcpPolicyRequestNetworkConfig) SetNetworkPolicyId(v string) *CreateMcpPolicyRequestNetworkConfig {
	s.NetworkPolicyId = &v
	return s
}

func (s *CreateMcpPolicyRequestNetworkConfig) Validate() error {
	return dara.Validate(s)
}

type CreateMcpPolicyRequestNetworkConfigGlobalConfig struct {
	ProxyConfig *CreateMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig `json:"ProxyConfig,omitempty" xml:"ProxyConfig,omitempty" type:"Struct"`
}

func (s CreateMcpPolicyRequestNetworkConfigGlobalConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpPolicyRequestNetworkConfigGlobalConfig) GoString() string {
	return s.String()
}

func (s *CreateMcpPolicyRequestNetworkConfigGlobalConfig) GetProxyConfig() *CreateMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig {
	return s.ProxyConfig
}

func (s *CreateMcpPolicyRequestNetworkConfigGlobalConfig) SetProxyConfig(v *CreateMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig) *CreateMcpPolicyRequestNetworkConfigGlobalConfig {
	s.ProxyConfig = v
	return s
}

func (s *CreateMcpPolicyRequestNetworkConfigGlobalConfig) Validate() error {
	return dara.Validate(s)
}

type CreateMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig struct {
	Address         *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Password        *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Port            *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	ProxyType       *string `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
	RedirectChannel *string `json:"RedirectChannel,omitempty" xml:"RedirectChannel,omitempty"`
	Username        *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s CreateMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig) GoString() string {
	return s.String()
}

func (s *CreateMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig) GetAddress() *string {
	return s.Address
}

func (s *CreateMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig) GetPassword() *string {
	return s.Password
}

func (s *CreateMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig) GetPort() *int32 {
	return s.Port
}

func (s *CreateMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig) GetProxyType() *string {
	return s.ProxyType
}

func (s *CreateMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig) GetRedirectChannel() *string {
	return s.RedirectChannel
}

func (s *CreateMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig) GetUsername() *string {
	return s.Username
}

func (s *CreateMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig) SetAddress(v string) *CreateMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig {
	s.Address = &v
	return s
}

func (s *CreateMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig) SetPassword(v string) *CreateMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig {
	s.Password = &v
	return s
}

func (s *CreateMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig) SetPort(v int32) *CreateMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig {
	s.Port = &v
	return s
}

func (s *CreateMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig) SetProxyType(v string) *CreateMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig {
	s.ProxyType = &v
	return s
}

func (s *CreateMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig) SetRedirectChannel(v string) *CreateMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig {
	s.RedirectChannel = &v
	return s
}

func (s *CreateMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig) SetUsername(v string) *CreateMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig {
	s.Username = &v
	return s
}

func (s *CreateMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig) Validate() error {
	return dara.Validate(s)
}

type CreateMcpPolicyRequestNetworkConfigRules struct {
	Domain      *string                                              `json:"Domain,omitempty" xml:"Domain,omitempty"`
	ProxyConfig *CreateMcpPolicyRequestNetworkConfigRulesProxyConfig `json:"ProxyConfig,omitempty" xml:"ProxyConfig,omitempty" type:"Struct"`
	RuleId      *string                                              `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s CreateMcpPolicyRequestNetworkConfigRules) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpPolicyRequestNetworkConfigRules) GoString() string {
	return s.String()
}

func (s *CreateMcpPolicyRequestNetworkConfigRules) GetDomain() *string {
	return s.Domain
}

func (s *CreateMcpPolicyRequestNetworkConfigRules) GetProxyConfig() *CreateMcpPolicyRequestNetworkConfigRulesProxyConfig {
	return s.ProxyConfig
}

func (s *CreateMcpPolicyRequestNetworkConfigRules) GetRuleId() *string {
	return s.RuleId
}

func (s *CreateMcpPolicyRequestNetworkConfigRules) SetDomain(v string) *CreateMcpPolicyRequestNetworkConfigRules {
	s.Domain = &v
	return s
}

func (s *CreateMcpPolicyRequestNetworkConfigRules) SetProxyConfig(v *CreateMcpPolicyRequestNetworkConfigRulesProxyConfig) *CreateMcpPolicyRequestNetworkConfigRules {
	s.ProxyConfig = v
	return s
}

func (s *CreateMcpPolicyRequestNetworkConfigRules) SetRuleId(v string) *CreateMcpPolicyRequestNetworkConfigRules {
	s.RuleId = &v
	return s
}

func (s *CreateMcpPolicyRequestNetworkConfigRules) Validate() error {
	return dara.Validate(s)
}

type CreateMcpPolicyRequestNetworkConfigRulesProxyConfig struct {
	Address         *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Password        *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Port            *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	ProxyType       *string `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
	RedirectChannel *string `json:"RedirectChannel,omitempty" xml:"RedirectChannel,omitempty"`
	Username        *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s CreateMcpPolicyRequestNetworkConfigRulesProxyConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpPolicyRequestNetworkConfigRulesProxyConfig) GoString() string {
	return s.String()
}

func (s *CreateMcpPolicyRequestNetworkConfigRulesProxyConfig) GetAddress() *string {
	return s.Address
}

func (s *CreateMcpPolicyRequestNetworkConfigRulesProxyConfig) GetPassword() *string {
	return s.Password
}

func (s *CreateMcpPolicyRequestNetworkConfigRulesProxyConfig) GetPort() *int32 {
	return s.Port
}

func (s *CreateMcpPolicyRequestNetworkConfigRulesProxyConfig) GetProxyType() *string {
	return s.ProxyType
}

func (s *CreateMcpPolicyRequestNetworkConfigRulesProxyConfig) GetRedirectChannel() *string {
	return s.RedirectChannel
}

func (s *CreateMcpPolicyRequestNetworkConfigRulesProxyConfig) GetUsername() *string {
	return s.Username
}

func (s *CreateMcpPolicyRequestNetworkConfigRulesProxyConfig) SetAddress(v string) *CreateMcpPolicyRequestNetworkConfigRulesProxyConfig {
	s.Address = &v
	return s
}

func (s *CreateMcpPolicyRequestNetworkConfigRulesProxyConfig) SetPassword(v string) *CreateMcpPolicyRequestNetworkConfigRulesProxyConfig {
	s.Password = &v
	return s
}

func (s *CreateMcpPolicyRequestNetworkConfigRulesProxyConfig) SetPort(v int32) *CreateMcpPolicyRequestNetworkConfigRulesProxyConfig {
	s.Port = &v
	return s
}

func (s *CreateMcpPolicyRequestNetworkConfigRulesProxyConfig) SetProxyType(v string) *CreateMcpPolicyRequestNetworkConfigRulesProxyConfig {
	s.ProxyType = &v
	return s
}

func (s *CreateMcpPolicyRequestNetworkConfigRulesProxyConfig) SetRedirectChannel(v string) *CreateMcpPolicyRequestNetworkConfigRulesProxyConfig {
	s.RedirectChannel = &v
	return s
}

func (s *CreateMcpPolicyRequestNetworkConfigRulesProxyConfig) SetUsername(v string) *CreateMcpPolicyRequestNetworkConfigRulesProxyConfig {
	s.Username = &v
	return s
}

func (s *CreateMcpPolicyRequestNetworkConfigRulesProxyConfig) Validate() error {
	return dara.Validate(s)
}
