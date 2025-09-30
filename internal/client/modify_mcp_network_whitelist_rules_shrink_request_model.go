// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMcpNetworkWhitelistRulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDebugMode(v bool) *ModifyMcpNetworkWhitelistRulesShrinkRequest
	GetDebugMode() *bool
	SetMcpPolicyId(v string) *ModifyMcpNetworkWhitelistRulesShrinkRequest
	GetMcpPolicyId() *string
	SetRulesShrink(v string) *ModifyMcpNetworkWhitelistRulesShrinkRequest
	GetRulesShrink() *string
}

type ModifyMcpNetworkWhitelistRulesShrinkRequest struct {
	DebugMode *bool `json:"DebugMode,omitempty" xml:"DebugMode,omitempty"`
	// This parameter is required.
	McpPolicyId *string `json:"McpPolicyId,omitempty" xml:"McpPolicyId,omitempty"`
	// This parameter is required.
	RulesShrink *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
}

func (s ModifyMcpNetworkWhitelistRulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyMcpNetworkWhitelistRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyMcpNetworkWhitelistRulesShrinkRequest) GetDebugMode() *bool {
	return s.DebugMode
}

func (s *ModifyMcpNetworkWhitelistRulesShrinkRequest) GetMcpPolicyId() *string {
	return s.McpPolicyId
}

func (s *ModifyMcpNetworkWhitelistRulesShrinkRequest) GetRulesShrink() *string {
	return s.RulesShrink
}

func (s *ModifyMcpNetworkWhitelistRulesShrinkRequest) SetDebugMode(v bool) *ModifyMcpNetworkWhitelistRulesShrinkRequest {
	s.DebugMode = &v
	return s
}

func (s *ModifyMcpNetworkWhitelistRulesShrinkRequest) SetMcpPolicyId(v string) *ModifyMcpNetworkWhitelistRulesShrinkRequest {
	s.McpPolicyId = &v
	return s
}

func (s *ModifyMcpNetworkWhitelistRulesShrinkRequest) SetRulesShrink(v string) *ModifyMcpNetworkWhitelistRulesShrinkRequest {
	s.RulesShrink = &v
	return s
}

func (s *ModifyMcpNetworkWhitelistRulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
