// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcpNetworkWhitelistRulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMcpPolicyId(v string) *CreateMcpNetworkWhitelistRulesShrinkRequest
	GetMcpPolicyId() *string
	SetRulesShrink(v string) *CreateMcpNetworkWhitelistRulesShrinkRequest
	GetRulesShrink() *string
}

type CreateMcpNetworkWhitelistRulesShrinkRequest struct {
	// This parameter is required.
	McpPolicyId *string `json:"McpPolicyId,omitempty" xml:"McpPolicyId,omitempty"`
	// This parameter is required.
	RulesShrink *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
}

func (s CreateMcpNetworkWhitelistRulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpNetworkWhitelistRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateMcpNetworkWhitelistRulesShrinkRequest) GetMcpPolicyId() *string {
	return s.McpPolicyId
}

func (s *CreateMcpNetworkWhitelistRulesShrinkRequest) GetRulesShrink() *string {
	return s.RulesShrink
}

func (s *CreateMcpNetworkWhitelistRulesShrinkRequest) SetMcpPolicyId(v string) *CreateMcpNetworkWhitelistRulesShrinkRequest {
	s.McpPolicyId = &v
	return s
}

func (s *CreateMcpNetworkWhitelistRulesShrinkRequest) SetRulesShrink(v string) *CreateMcpNetworkWhitelistRulesShrinkRequest {
	s.RulesShrink = &v
	return s
}

func (s *CreateMcpNetworkWhitelistRulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
