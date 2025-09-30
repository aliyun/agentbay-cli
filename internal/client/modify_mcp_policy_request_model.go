// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMcpPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConcurrency(v int32) *ModifyMcpPolicyRequest
	GetConcurrency() *int32
	SetDebugMode(v bool) *ModifyMcpPolicyRequest
	GetDebugMode() *bool
	SetDesktopMaxRuntime(v int32) *ModifyMcpPolicyRequest
	GetDesktopMaxRuntime() *int32
	SetDisplayConfig(v *ModifyMcpPolicyRequestDisplayConfig) *ModifyMcpPolicyRequest
	GetDisplayConfig() *ModifyMcpPolicyRequestDisplayConfig
	SetIdleTimeoutSwitch(v bool) *ModifyMcpPolicyRequest
	GetIdleTimeoutSwitch() *bool
	SetMcpIdleTimeout(v int32) *ModifyMcpPolicyRequest
	GetMcpIdleTimeout() *int32
	SetNetworkConfig(v *ModifyMcpPolicyRequestNetworkConfig) *ModifyMcpPolicyRequest
	GetNetworkConfig() *ModifyMcpPolicyRequestNetworkConfig
	SetPolicyId(v string) *ModifyMcpPolicyRequest
	GetPolicyId() *string
	SetPolicyName(v string) *ModifyMcpPolicyRequest
	GetPolicyName() *string
	SetUserIdleTimeout(v int32) *ModifyMcpPolicyRequest
	GetUserIdleTimeout() *int32
}

type ModifyMcpPolicyRequest struct {
	Concurrency       *int32                               `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	DebugMode         *bool                                `json:"DebugMode,omitempty" xml:"DebugMode,omitempty"`
	DesktopMaxRuntime *int32                               `json:"DesktopMaxRuntime,omitempty" xml:"DesktopMaxRuntime,omitempty"`
	DisplayConfig     *ModifyMcpPolicyRequestDisplayConfig `json:"DisplayConfig,omitempty" xml:"DisplayConfig,omitempty" type:"Struct"`
	IdleTimeoutSwitch *bool                                `json:"IdleTimeoutSwitch,omitempty" xml:"IdleTimeoutSwitch,omitempty"`
	McpIdleTimeout    *int32                               `json:"McpIdleTimeout,omitempty" xml:"McpIdleTimeout,omitempty"`
	NetworkConfig     *ModifyMcpPolicyRequestNetworkConfig `json:"NetworkConfig,omitempty" xml:"NetworkConfig,omitempty" type:"Struct"`
	PolicyId          *string                              `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	PolicyName        *string                              `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	UserIdleTimeout   *int32                               `json:"UserIdleTimeout,omitempty" xml:"UserIdleTimeout,omitempty"`
}

func (s ModifyMcpPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyMcpPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyMcpPolicyRequest) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *ModifyMcpPolicyRequest) GetDebugMode() *bool {
	return s.DebugMode
}

func (s *ModifyMcpPolicyRequest) GetDesktopMaxRuntime() *int32 {
	return s.DesktopMaxRuntime
}

func (s *ModifyMcpPolicyRequest) GetDisplayConfig() *ModifyMcpPolicyRequestDisplayConfig {
	return s.DisplayConfig
}

func (s *ModifyMcpPolicyRequest) GetIdleTimeoutSwitch() *bool {
	return s.IdleTimeoutSwitch
}

func (s *ModifyMcpPolicyRequest) GetMcpIdleTimeout() *int32 {
	return s.McpIdleTimeout
}

func (s *ModifyMcpPolicyRequest) GetNetworkConfig() *ModifyMcpPolicyRequestNetworkConfig {
	return s.NetworkConfig
}

func (s *ModifyMcpPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ModifyMcpPolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ModifyMcpPolicyRequest) GetUserIdleTimeout() *int32 {
	return s.UserIdleTimeout
}

func (s *ModifyMcpPolicyRequest) SetConcurrency(v int32) *ModifyMcpPolicyRequest {
	s.Concurrency = &v
	return s
}

func (s *ModifyMcpPolicyRequest) SetDebugMode(v bool) *ModifyMcpPolicyRequest {
	s.DebugMode = &v
	return s
}

func (s *ModifyMcpPolicyRequest) SetDesktopMaxRuntime(v int32) *ModifyMcpPolicyRequest {
	s.DesktopMaxRuntime = &v
	return s
}

func (s *ModifyMcpPolicyRequest) SetDisplayConfig(v *ModifyMcpPolicyRequestDisplayConfig) *ModifyMcpPolicyRequest {
	s.DisplayConfig = v
	return s
}

func (s *ModifyMcpPolicyRequest) SetIdleTimeoutSwitch(v bool) *ModifyMcpPolicyRequest {
	s.IdleTimeoutSwitch = &v
	return s
}

func (s *ModifyMcpPolicyRequest) SetMcpIdleTimeout(v int32) *ModifyMcpPolicyRequest {
	s.McpIdleTimeout = &v
	return s
}

func (s *ModifyMcpPolicyRequest) SetNetworkConfig(v *ModifyMcpPolicyRequestNetworkConfig) *ModifyMcpPolicyRequest {
	s.NetworkConfig = v
	return s
}

func (s *ModifyMcpPolicyRequest) SetPolicyId(v string) *ModifyMcpPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *ModifyMcpPolicyRequest) SetPolicyName(v string) *ModifyMcpPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *ModifyMcpPolicyRequest) SetUserIdleTimeout(v int32) *ModifyMcpPolicyRequest {
	s.UserIdleTimeout = &v
	return s
}

func (s *ModifyMcpPolicyRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyMcpPolicyRequestDisplayConfig struct {
	DisplayMode *string `json:"DisplayMode,omitempty" xml:"DisplayMode,omitempty"`
}

func (s ModifyMcpPolicyRequestDisplayConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyMcpPolicyRequestDisplayConfig) GoString() string {
	return s.String()
}

func (s *ModifyMcpPolicyRequestDisplayConfig) GetDisplayMode() *string {
	return s.DisplayMode
}

func (s *ModifyMcpPolicyRequestDisplayConfig) SetDisplayMode(v string) *ModifyMcpPolicyRequestDisplayConfig {
	s.DisplayMode = &v
	return s
}

func (s *ModifyMcpPolicyRequestDisplayConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyMcpPolicyRequestNetworkConfig struct {
	Enabled         *bool                                            `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	GlobalConfig    *ModifyMcpPolicyRequestNetworkConfigGlobalConfig `json:"GlobalConfig,omitempty" xml:"GlobalConfig,omitempty" type:"Struct"`
	Mode            *string                                          `json:"Mode,omitempty" xml:"Mode,omitempty"`
	NetworkPolicyId *string                                          `json:"NetworkPolicyId,omitempty" xml:"NetworkPolicyId,omitempty"`
	Rules           []*ModifyMcpPolicyRequestNetworkConfigRules      `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s ModifyMcpPolicyRequestNetworkConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyMcpPolicyRequestNetworkConfig) GoString() string {
	return s.String()
}

func (s *ModifyMcpPolicyRequestNetworkConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *ModifyMcpPolicyRequestNetworkConfig) GetGlobalConfig() *ModifyMcpPolicyRequestNetworkConfigGlobalConfig {
	return s.GlobalConfig
}

func (s *ModifyMcpPolicyRequestNetworkConfig) GetMode() *string {
	return s.Mode
}

func (s *ModifyMcpPolicyRequestNetworkConfig) GetNetworkPolicyId() *string {
	return s.NetworkPolicyId
}

func (s *ModifyMcpPolicyRequestNetworkConfig) GetRules() []*ModifyMcpPolicyRequestNetworkConfigRules {
	return s.Rules
}

func (s *ModifyMcpPolicyRequestNetworkConfig) SetEnabled(v bool) *ModifyMcpPolicyRequestNetworkConfig {
	s.Enabled = &v
	return s
}

func (s *ModifyMcpPolicyRequestNetworkConfig) SetGlobalConfig(v *ModifyMcpPolicyRequestNetworkConfigGlobalConfig) *ModifyMcpPolicyRequestNetworkConfig {
	s.GlobalConfig = v
	return s
}

func (s *ModifyMcpPolicyRequestNetworkConfig) SetMode(v string) *ModifyMcpPolicyRequestNetworkConfig {
	s.Mode = &v
	return s
}

func (s *ModifyMcpPolicyRequestNetworkConfig) SetNetworkPolicyId(v string) *ModifyMcpPolicyRequestNetworkConfig {
	s.NetworkPolicyId = &v
	return s
}

func (s *ModifyMcpPolicyRequestNetworkConfig) SetRules(v []*ModifyMcpPolicyRequestNetworkConfigRules) *ModifyMcpPolicyRequestNetworkConfig {
	s.Rules = v
	return s
}

func (s *ModifyMcpPolicyRequestNetworkConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyMcpPolicyRequestNetworkConfigGlobalConfig struct {
	ProxyConfig *ModifyMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig `json:"ProxyConfig,omitempty" xml:"ProxyConfig,omitempty" type:"Struct"`
}

func (s ModifyMcpPolicyRequestNetworkConfigGlobalConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyMcpPolicyRequestNetworkConfigGlobalConfig) GoString() string {
	return s.String()
}

func (s *ModifyMcpPolicyRequestNetworkConfigGlobalConfig) GetProxyConfig() *ModifyMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig {
	return s.ProxyConfig
}

func (s *ModifyMcpPolicyRequestNetworkConfigGlobalConfig) SetProxyConfig(v *ModifyMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig) *ModifyMcpPolicyRequestNetworkConfigGlobalConfig {
	s.ProxyConfig = v
	return s
}

func (s *ModifyMcpPolicyRequestNetworkConfigGlobalConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig struct {
	Address         *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Password        *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Port            *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	ProxyType       *string `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
	RedirectChannel *string `json:"RedirectChannel,omitempty" xml:"RedirectChannel,omitempty"`
	Username        *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ModifyMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig) GoString() string {
	return s.String()
}

func (s *ModifyMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig) GetAddress() *string {
	return s.Address
}

func (s *ModifyMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig) GetPassword() *string {
	return s.Password
}

func (s *ModifyMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig) GetPort() *int32 {
	return s.Port
}

func (s *ModifyMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig) GetProxyType() *string {
	return s.ProxyType
}

func (s *ModifyMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig) GetRedirectChannel() *string {
	return s.RedirectChannel
}

func (s *ModifyMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig) GetUsername() *string {
	return s.Username
}

func (s *ModifyMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig) SetAddress(v string) *ModifyMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig {
	s.Address = &v
	return s
}

func (s *ModifyMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig) SetPassword(v string) *ModifyMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig {
	s.Password = &v
	return s
}

func (s *ModifyMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig) SetPort(v int32) *ModifyMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig {
	s.Port = &v
	return s
}

func (s *ModifyMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig) SetProxyType(v string) *ModifyMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig {
	s.ProxyType = &v
	return s
}

func (s *ModifyMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig) SetRedirectChannel(v string) *ModifyMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig {
	s.RedirectChannel = &v
	return s
}

func (s *ModifyMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig) SetUsername(v string) *ModifyMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig {
	s.Username = &v
	return s
}

func (s *ModifyMcpPolicyRequestNetworkConfigGlobalConfigProxyConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyMcpPolicyRequestNetworkConfigRules struct {
	Domain      *string                                              `json:"Domain,omitempty" xml:"Domain,omitempty"`
	ProxyConfig *ModifyMcpPolicyRequestNetworkConfigRulesProxyConfig `json:"ProxyConfig,omitempty" xml:"ProxyConfig,omitempty" type:"Struct"`
	RuleId      *string                                              `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s ModifyMcpPolicyRequestNetworkConfigRules) String() string {
	return dara.Prettify(s)
}

func (s ModifyMcpPolicyRequestNetworkConfigRules) GoString() string {
	return s.String()
}

func (s *ModifyMcpPolicyRequestNetworkConfigRules) GetDomain() *string {
	return s.Domain
}

func (s *ModifyMcpPolicyRequestNetworkConfigRules) GetProxyConfig() *ModifyMcpPolicyRequestNetworkConfigRulesProxyConfig {
	return s.ProxyConfig
}

func (s *ModifyMcpPolicyRequestNetworkConfigRules) GetRuleId() *string {
	return s.RuleId
}

func (s *ModifyMcpPolicyRequestNetworkConfigRules) SetDomain(v string) *ModifyMcpPolicyRequestNetworkConfigRules {
	s.Domain = &v
	return s
}

func (s *ModifyMcpPolicyRequestNetworkConfigRules) SetProxyConfig(v *ModifyMcpPolicyRequestNetworkConfigRulesProxyConfig) *ModifyMcpPolicyRequestNetworkConfigRules {
	s.ProxyConfig = v
	return s
}

func (s *ModifyMcpPolicyRequestNetworkConfigRules) SetRuleId(v string) *ModifyMcpPolicyRequestNetworkConfigRules {
	s.RuleId = &v
	return s
}

func (s *ModifyMcpPolicyRequestNetworkConfigRules) Validate() error {
	return dara.Validate(s)
}

type ModifyMcpPolicyRequestNetworkConfigRulesProxyConfig struct {
	Address         *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Password        *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Port            *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	ProxyType       *string `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
	RedirectChannel *string `json:"RedirectChannel,omitempty" xml:"RedirectChannel,omitempty"`
	Username        *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ModifyMcpPolicyRequestNetworkConfigRulesProxyConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyMcpPolicyRequestNetworkConfigRulesProxyConfig) GoString() string {
	return s.String()
}

func (s *ModifyMcpPolicyRequestNetworkConfigRulesProxyConfig) GetAddress() *string {
	return s.Address
}

func (s *ModifyMcpPolicyRequestNetworkConfigRulesProxyConfig) GetPassword() *string {
	return s.Password
}

func (s *ModifyMcpPolicyRequestNetworkConfigRulesProxyConfig) GetPort() *int32 {
	return s.Port
}

func (s *ModifyMcpPolicyRequestNetworkConfigRulesProxyConfig) GetProxyType() *string {
	return s.ProxyType
}

func (s *ModifyMcpPolicyRequestNetworkConfigRulesProxyConfig) GetRedirectChannel() *string {
	return s.RedirectChannel
}

func (s *ModifyMcpPolicyRequestNetworkConfigRulesProxyConfig) GetUsername() *string {
	return s.Username
}

func (s *ModifyMcpPolicyRequestNetworkConfigRulesProxyConfig) SetAddress(v string) *ModifyMcpPolicyRequestNetworkConfigRulesProxyConfig {
	s.Address = &v
	return s
}

func (s *ModifyMcpPolicyRequestNetworkConfigRulesProxyConfig) SetPassword(v string) *ModifyMcpPolicyRequestNetworkConfigRulesProxyConfig {
	s.Password = &v
	return s
}

func (s *ModifyMcpPolicyRequestNetworkConfigRulesProxyConfig) SetPort(v int32) *ModifyMcpPolicyRequestNetworkConfigRulesProxyConfig {
	s.Port = &v
	return s
}

func (s *ModifyMcpPolicyRequestNetworkConfigRulesProxyConfig) SetProxyType(v string) *ModifyMcpPolicyRequestNetworkConfigRulesProxyConfig {
	s.ProxyType = &v
	return s
}

func (s *ModifyMcpPolicyRequestNetworkConfigRulesProxyConfig) SetRedirectChannel(v string) *ModifyMcpPolicyRequestNetworkConfigRulesProxyConfig {
	s.RedirectChannel = &v
	return s
}

func (s *ModifyMcpPolicyRequestNetworkConfigRulesProxyConfig) SetUsername(v string) *ModifyMcpPolicyRequestNetworkConfigRulesProxyConfig {
	s.Username = &v
	return s
}

func (s *ModifyMcpPolicyRequestNetworkConfigRulesProxyConfig) Validate() error {
	return dara.Validate(s)
}
