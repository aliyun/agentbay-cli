// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMcpNetworkWhitelistRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDebugMode(v bool) *ModifyMcpNetworkWhitelistRulesRequest
	GetDebugMode() *bool
	SetMcpPolicyId(v string) *ModifyMcpNetworkWhitelistRulesRequest
	GetMcpPolicyId() *string
	SetRules(v []*ModifyMcpNetworkWhitelistRulesRequestRules) *ModifyMcpNetworkWhitelistRulesRequest
	GetRules() []*ModifyMcpNetworkWhitelistRulesRequestRules
}

type ModifyMcpNetworkWhitelistRulesRequest struct {
	DebugMode *bool `json:"DebugMode,omitempty" xml:"DebugMode,omitempty"`
	// This parameter is required.
	McpPolicyId *string `json:"McpPolicyId,omitempty" xml:"McpPolicyId,omitempty"`
	// This parameter is required.
	Rules []*ModifyMcpNetworkWhitelistRulesRequestRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s ModifyMcpNetworkWhitelistRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyMcpNetworkWhitelistRulesRequest) GoString() string {
	return s.String()
}

func (s *ModifyMcpNetworkWhitelistRulesRequest) GetDebugMode() *bool {
	return s.DebugMode
}

func (s *ModifyMcpNetworkWhitelistRulesRequest) GetMcpPolicyId() *string {
	return s.McpPolicyId
}

func (s *ModifyMcpNetworkWhitelistRulesRequest) GetRules() []*ModifyMcpNetworkWhitelistRulesRequestRules {
	return s.Rules
}

func (s *ModifyMcpNetworkWhitelistRulesRequest) SetDebugMode(v bool) *ModifyMcpNetworkWhitelistRulesRequest {
	s.DebugMode = &v
	return s
}

func (s *ModifyMcpNetworkWhitelistRulesRequest) SetMcpPolicyId(v string) *ModifyMcpNetworkWhitelistRulesRequest {
	s.McpPolicyId = &v
	return s
}

func (s *ModifyMcpNetworkWhitelistRulesRequest) SetRules(v []*ModifyMcpNetworkWhitelistRulesRequestRules) *ModifyMcpNetworkWhitelistRulesRequest {
	s.Rules = v
	return s
}

func (s *ModifyMcpNetworkWhitelistRulesRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyMcpNetworkWhitelistRulesRequestRules struct {
	// This parameter is required.
	Domain      *string                                                `json:"Domain,omitempty" xml:"Domain,omitempty"`
	ProxyConfig *ModifyMcpNetworkWhitelistRulesRequestRulesProxyConfig `json:"ProxyConfig,omitempty" xml:"ProxyConfig,omitempty" type:"Struct"`
	// This parameter is required.
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s ModifyMcpNetworkWhitelistRulesRequestRules) String() string {
	return dara.Prettify(s)
}

func (s ModifyMcpNetworkWhitelistRulesRequestRules) GoString() string {
	return s.String()
}

func (s *ModifyMcpNetworkWhitelistRulesRequestRules) GetDomain() *string {
	return s.Domain
}

func (s *ModifyMcpNetworkWhitelistRulesRequestRules) GetProxyConfig() *ModifyMcpNetworkWhitelistRulesRequestRulesProxyConfig {
	return s.ProxyConfig
}

func (s *ModifyMcpNetworkWhitelistRulesRequestRules) GetRuleId() *string {
	return s.RuleId
}

func (s *ModifyMcpNetworkWhitelistRulesRequestRules) SetDomain(v string) *ModifyMcpNetworkWhitelistRulesRequestRules {
	s.Domain = &v
	return s
}

func (s *ModifyMcpNetworkWhitelistRulesRequestRules) SetProxyConfig(v *ModifyMcpNetworkWhitelistRulesRequestRulesProxyConfig) *ModifyMcpNetworkWhitelistRulesRequestRules {
	s.ProxyConfig = v
	return s
}

func (s *ModifyMcpNetworkWhitelistRulesRequestRules) SetRuleId(v string) *ModifyMcpNetworkWhitelistRulesRequestRules {
	s.RuleId = &v
	return s
}

func (s *ModifyMcpNetworkWhitelistRulesRequestRules) Validate() error {
	return dara.Validate(s)
}

type ModifyMcpNetworkWhitelistRulesRequestRulesProxyConfig struct {
	Address  *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	// This parameter is required.
	ProxyType *string `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
	// This parameter is required.
	RedirectChannel *string `json:"RedirectChannel,omitempty" xml:"RedirectChannel,omitempty"`
	Username        *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ModifyMcpNetworkWhitelistRulesRequestRulesProxyConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyMcpNetworkWhitelistRulesRequestRulesProxyConfig) GoString() string {
	return s.String()
}

func (s *ModifyMcpNetworkWhitelistRulesRequestRulesProxyConfig) GetAddress() *string {
	return s.Address
}

func (s *ModifyMcpNetworkWhitelistRulesRequestRulesProxyConfig) GetPassword() *string {
	return s.Password
}

func (s *ModifyMcpNetworkWhitelistRulesRequestRulesProxyConfig) GetPort() *int32 {
	return s.Port
}

func (s *ModifyMcpNetworkWhitelistRulesRequestRulesProxyConfig) GetProxyType() *string {
	return s.ProxyType
}

func (s *ModifyMcpNetworkWhitelistRulesRequestRulesProxyConfig) GetRedirectChannel() *string {
	return s.RedirectChannel
}

func (s *ModifyMcpNetworkWhitelistRulesRequestRulesProxyConfig) GetUsername() *string {
	return s.Username
}

func (s *ModifyMcpNetworkWhitelistRulesRequestRulesProxyConfig) SetAddress(v string) *ModifyMcpNetworkWhitelistRulesRequestRulesProxyConfig {
	s.Address = &v
	return s
}

func (s *ModifyMcpNetworkWhitelistRulesRequestRulesProxyConfig) SetPassword(v string) *ModifyMcpNetworkWhitelistRulesRequestRulesProxyConfig {
	s.Password = &v
	return s
}

func (s *ModifyMcpNetworkWhitelistRulesRequestRulesProxyConfig) SetPort(v int32) *ModifyMcpNetworkWhitelistRulesRequestRulesProxyConfig {
	s.Port = &v
	return s
}

func (s *ModifyMcpNetworkWhitelistRulesRequestRulesProxyConfig) SetProxyType(v string) *ModifyMcpNetworkWhitelistRulesRequestRulesProxyConfig {
	s.ProxyType = &v
	return s
}

func (s *ModifyMcpNetworkWhitelistRulesRequestRulesProxyConfig) SetRedirectChannel(v string) *ModifyMcpNetworkWhitelistRulesRequestRulesProxyConfig {
	s.RedirectChannel = &v
	return s
}

func (s *ModifyMcpNetworkWhitelistRulesRequestRulesProxyConfig) SetUsername(v string) *ModifyMcpNetworkWhitelistRulesRequestRulesProxyConfig {
	s.Username = &v
	return s
}

func (s *ModifyMcpNetworkWhitelistRulesRequestRulesProxyConfig) Validate() error {
	return dara.Validate(s)
}
