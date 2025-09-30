// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcpNetworkWhitelistRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMcpPolicyId(v string) *CreateMcpNetworkWhitelistRulesRequest
	GetMcpPolicyId() *string
	SetRules(v []*CreateMcpNetworkWhitelistRulesRequestRules) *CreateMcpNetworkWhitelistRulesRequest
	GetRules() []*CreateMcpNetworkWhitelistRulesRequestRules
}

type CreateMcpNetworkWhitelistRulesRequest struct {
	// This parameter is required.
	McpPolicyId *string `json:"McpPolicyId,omitempty" xml:"McpPolicyId,omitempty"`
	// This parameter is required.
	Rules []*CreateMcpNetworkWhitelistRulesRequestRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s CreateMcpNetworkWhitelistRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpNetworkWhitelistRulesRequest) GoString() string {
	return s.String()
}

func (s *CreateMcpNetworkWhitelistRulesRequest) GetMcpPolicyId() *string {
	return s.McpPolicyId
}

func (s *CreateMcpNetworkWhitelistRulesRequest) GetRules() []*CreateMcpNetworkWhitelistRulesRequestRules {
	return s.Rules
}

func (s *CreateMcpNetworkWhitelistRulesRequest) SetMcpPolicyId(v string) *CreateMcpNetworkWhitelistRulesRequest {
	s.McpPolicyId = &v
	return s
}

func (s *CreateMcpNetworkWhitelistRulesRequest) SetRules(v []*CreateMcpNetworkWhitelistRulesRequestRules) *CreateMcpNetworkWhitelistRulesRequest {
	s.Rules = v
	return s
}

func (s *CreateMcpNetworkWhitelistRulesRequest) Validate() error {
	return dara.Validate(s)
}

type CreateMcpNetworkWhitelistRulesRequestRules struct {
	// This parameter is required.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	ProxyConfig *CreateMcpNetworkWhitelistRulesRequestRulesProxyConfig `json:"ProxyConfig,omitempty" xml:"ProxyConfig,omitempty" type:"Struct"`
}

func (s CreateMcpNetworkWhitelistRulesRequestRules) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpNetworkWhitelistRulesRequestRules) GoString() string {
	return s.String()
}

func (s *CreateMcpNetworkWhitelistRulesRequestRules) GetDomain() *string {
	return s.Domain
}

func (s *CreateMcpNetworkWhitelistRulesRequestRules) GetProxyConfig() *CreateMcpNetworkWhitelistRulesRequestRulesProxyConfig {
	return s.ProxyConfig
}

func (s *CreateMcpNetworkWhitelistRulesRequestRules) SetDomain(v string) *CreateMcpNetworkWhitelistRulesRequestRules {
	s.Domain = &v
	return s
}

func (s *CreateMcpNetworkWhitelistRulesRequestRules) SetProxyConfig(v *CreateMcpNetworkWhitelistRulesRequestRulesProxyConfig) *CreateMcpNetworkWhitelistRulesRequestRules {
	s.ProxyConfig = v
	return s
}

func (s *CreateMcpNetworkWhitelistRulesRequestRules) Validate() error {
	return dara.Validate(s)
}

type CreateMcpNetworkWhitelistRulesRequestRulesProxyConfig struct {
	Address  *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	// This parameter is required.
	ProxyType *string `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
	// This parameter is required.
	RedirectChannel *string `json:"RedirectChannel,omitempty" xml:"RedirectChannel,omitempty"`
	Username        *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s CreateMcpNetworkWhitelistRulesRequestRulesProxyConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpNetworkWhitelistRulesRequestRulesProxyConfig) GoString() string {
	return s.String()
}

func (s *CreateMcpNetworkWhitelistRulesRequestRulesProxyConfig) GetAddress() *string {
	return s.Address
}

func (s *CreateMcpNetworkWhitelistRulesRequestRulesProxyConfig) GetPassword() *string {
	return s.Password
}

func (s *CreateMcpNetworkWhitelistRulesRequestRulesProxyConfig) GetPort() *int32 {
	return s.Port
}

func (s *CreateMcpNetworkWhitelistRulesRequestRulesProxyConfig) GetProxyType() *string {
	return s.ProxyType
}

func (s *CreateMcpNetworkWhitelistRulesRequestRulesProxyConfig) GetRedirectChannel() *string {
	return s.RedirectChannel
}

func (s *CreateMcpNetworkWhitelistRulesRequestRulesProxyConfig) GetUsername() *string {
	return s.Username
}

func (s *CreateMcpNetworkWhitelistRulesRequestRulesProxyConfig) SetAddress(v string) *CreateMcpNetworkWhitelistRulesRequestRulesProxyConfig {
	s.Address = &v
	return s
}

func (s *CreateMcpNetworkWhitelistRulesRequestRulesProxyConfig) SetPassword(v string) *CreateMcpNetworkWhitelistRulesRequestRulesProxyConfig {
	s.Password = &v
	return s
}

func (s *CreateMcpNetworkWhitelistRulesRequestRulesProxyConfig) SetPort(v int32) *CreateMcpNetworkWhitelistRulesRequestRulesProxyConfig {
	s.Port = &v
	return s
}

func (s *CreateMcpNetworkWhitelistRulesRequestRulesProxyConfig) SetProxyType(v string) *CreateMcpNetworkWhitelistRulesRequestRulesProxyConfig {
	s.ProxyType = &v
	return s
}

func (s *CreateMcpNetworkWhitelistRulesRequestRulesProxyConfig) SetRedirectChannel(v string) *CreateMcpNetworkWhitelistRulesRequestRulesProxyConfig {
	s.RedirectChannel = &v
	return s
}

func (s *CreateMcpNetworkWhitelistRulesRequestRulesProxyConfig) SetUsername(v string) *CreateMcpNetworkWhitelistRulesRequestRulesProxyConfig {
	s.Username = &v
	return s
}

func (s *CreateMcpNetworkWhitelistRulesRequestRulesProxyConfig) Validate() error {
	return dara.Validate(s)
}
