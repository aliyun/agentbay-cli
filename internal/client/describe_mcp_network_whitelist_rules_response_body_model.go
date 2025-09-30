// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMcpNetworkWhitelistRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeMcpNetworkWhitelistRulesResponseBody
	GetCode() *string
	SetData(v *DescribeMcpNetworkWhitelistRulesResponseBodyData) *DescribeMcpNetworkWhitelistRulesResponseBody
	GetData() *DescribeMcpNetworkWhitelistRulesResponseBodyData
	SetHttpStatusCode(v string) *DescribeMcpNetworkWhitelistRulesResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DescribeMcpNetworkWhitelistRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeMcpNetworkWhitelistRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeMcpNetworkWhitelistRulesResponseBody
	GetSuccess() *bool
}

type DescribeMcpNetworkWhitelistRulesResponseBody struct {
	Code           *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *DescribeMcpNetworkWhitelistRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *string                                           `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMcpNetworkWhitelistRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMcpNetworkWhitelistRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBody) GetData() *DescribeMcpNetworkWhitelistRulesResponseBodyData {
	return s.Data
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBody) SetCode(v string) *DescribeMcpNetworkWhitelistRulesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBody) SetData(v *DescribeMcpNetworkWhitelistRulesResponseBodyData) *DescribeMcpNetworkWhitelistRulesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBody) SetHttpStatusCode(v string) *DescribeMcpNetworkWhitelistRulesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBody) SetMessage(v string) *DescribeMcpNetworkWhitelistRulesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBody) SetRequestId(v string) *DescribeMcpNetworkWhitelistRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBody) SetSuccess(v bool) *DescribeMcpNetworkWhitelistRulesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeMcpNetworkWhitelistRulesResponseBodyData struct {
	NextToken  *string                                                  `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Rules      []*DescribeMcpNetworkWhitelistRulesResponseBodyDataRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	TotalCount *int32                                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeMcpNetworkWhitelistRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeMcpNetworkWhitelistRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBodyData) GetRules() []*DescribeMcpNetworkWhitelistRulesResponseBodyDataRules {
	return s.Rules
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBodyData) SetNextToken(v string) *DescribeMcpNetworkWhitelistRulesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBodyData) SetRules(v []*DescribeMcpNetworkWhitelistRulesResponseBodyDataRules) *DescribeMcpNetworkWhitelistRulesResponseBodyData {
	s.Rules = v
	return s
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBodyData) SetTotalCount(v int32) *DescribeMcpNetworkWhitelistRulesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeMcpNetworkWhitelistRulesResponseBodyDataRules struct {
	Domain      *string                                                           `json:"Domain,omitempty" xml:"Domain,omitempty"`
	ProxyConfig *DescribeMcpNetworkWhitelistRulesResponseBodyDataRulesProxyConfig `json:"ProxyConfig,omitempty" xml:"ProxyConfig,omitempty" type:"Struct"`
	RuleId      *string                                                           `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeMcpNetworkWhitelistRulesResponseBodyDataRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeMcpNetworkWhitelistRulesResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBodyDataRules) GetDomain() *string {
	return s.Domain
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBodyDataRules) GetProxyConfig() *DescribeMcpNetworkWhitelistRulesResponseBodyDataRulesProxyConfig {
	return s.ProxyConfig
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBodyDataRules) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBodyDataRules) SetDomain(v string) *DescribeMcpNetworkWhitelistRulesResponseBodyDataRules {
	s.Domain = &v
	return s
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBodyDataRules) SetProxyConfig(v *DescribeMcpNetworkWhitelistRulesResponseBodyDataRulesProxyConfig) *DescribeMcpNetworkWhitelistRulesResponseBodyDataRules {
	s.ProxyConfig = v
	return s
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBodyDataRules) SetRuleId(v string) *DescribeMcpNetworkWhitelistRulesResponseBodyDataRules {
	s.RuleId = &v
	return s
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBodyDataRules) Validate() error {
	return dara.Validate(s)
}

type DescribeMcpNetworkWhitelistRulesResponseBodyDataRulesProxyConfig struct {
	Address         *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Password        *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Port            *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	ProxyType       *string `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
	RedirectChannel *string `json:"RedirectChannel,omitempty" xml:"RedirectChannel,omitempty"`
	Username        *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s DescribeMcpNetworkWhitelistRulesResponseBodyDataRulesProxyConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeMcpNetworkWhitelistRulesResponseBodyDataRulesProxyConfig) GoString() string {
	return s.String()
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBodyDataRulesProxyConfig) GetAddress() *string {
	return s.Address
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBodyDataRulesProxyConfig) GetPassword() *string {
	return s.Password
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBodyDataRulesProxyConfig) GetPort() *int32 {
	return s.Port
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBodyDataRulesProxyConfig) GetProxyType() *string {
	return s.ProxyType
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBodyDataRulesProxyConfig) GetRedirectChannel() *string {
	return s.RedirectChannel
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBodyDataRulesProxyConfig) GetUsername() *string {
	return s.Username
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBodyDataRulesProxyConfig) SetAddress(v string) *DescribeMcpNetworkWhitelistRulesResponseBodyDataRulesProxyConfig {
	s.Address = &v
	return s
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBodyDataRulesProxyConfig) SetPassword(v string) *DescribeMcpNetworkWhitelistRulesResponseBodyDataRulesProxyConfig {
	s.Password = &v
	return s
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBodyDataRulesProxyConfig) SetPort(v int32) *DescribeMcpNetworkWhitelistRulesResponseBodyDataRulesProxyConfig {
	s.Port = &v
	return s
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBodyDataRulesProxyConfig) SetProxyType(v string) *DescribeMcpNetworkWhitelistRulesResponseBodyDataRulesProxyConfig {
	s.ProxyType = &v
	return s
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBodyDataRulesProxyConfig) SetRedirectChannel(v string) *DescribeMcpNetworkWhitelistRulesResponseBodyDataRulesProxyConfig {
	s.RedirectChannel = &v
	return s
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBodyDataRulesProxyConfig) SetUsername(v string) *DescribeMcpNetworkWhitelistRulesResponseBodyDataRulesProxyConfig {
	s.Username = &v
	return s
}

func (s *DescribeMcpNetworkWhitelistRulesResponseBodyDataRulesProxyConfig) Validate() error {
	return dara.Validate(s)
}
