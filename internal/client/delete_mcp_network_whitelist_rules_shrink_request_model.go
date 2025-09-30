// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcpNetworkWhitelistRulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMcpPolicyId(v string) *DeleteMcpNetworkWhitelistRulesShrinkRequest
	GetMcpPolicyId() *string
	SetRuleIdsShrink(v string) *DeleteMcpNetworkWhitelistRulesShrinkRequest
	GetRuleIdsShrink() *string
}

type DeleteMcpNetworkWhitelistRulesShrinkRequest struct {
	// This parameter is required.
	McpPolicyId *string `json:"McpPolicyId,omitempty" xml:"McpPolicyId,omitempty"`
	// This parameter is required.
	RuleIdsShrink *string `json:"RuleIds,omitempty" xml:"RuleIds,omitempty"`
}

func (s DeleteMcpNetworkWhitelistRulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcpNetworkWhitelistRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteMcpNetworkWhitelistRulesShrinkRequest) GetMcpPolicyId() *string {
	return s.McpPolicyId
}

func (s *DeleteMcpNetworkWhitelistRulesShrinkRequest) GetRuleIdsShrink() *string {
	return s.RuleIdsShrink
}

func (s *DeleteMcpNetworkWhitelistRulesShrinkRequest) SetMcpPolicyId(v string) *DeleteMcpNetworkWhitelistRulesShrinkRequest {
	s.McpPolicyId = &v
	return s
}

func (s *DeleteMcpNetworkWhitelistRulesShrinkRequest) SetRuleIdsShrink(v string) *DeleteMcpNetworkWhitelistRulesShrinkRequest {
	s.RuleIdsShrink = &v
	return s
}

func (s *DeleteMcpNetworkWhitelistRulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
