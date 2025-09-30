// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcpNetworkWhitelistRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMcpPolicyId(v string) *DeleteMcpNetworkWhitelistRulesRequest
	GetMcpPolicyId() *string
	SetRuleIds(v []*string) *DeleteMcpNetworkWhitelistRulesRequest
	GetRuleIds() []*string
}

type DeleteMcpNetworkWhitelistRulesRequest struct {
	// This parameter is required.
	McpPolicyId *string `json:"McpPolicyId,omitempty" xml:"McpPolicyId,omitempty"`
	// This parameter is required.
	RuleIds []*string `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Repeated"`
}

func (s DeleteMcpNetworkWhitelistRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcpNetworkWhitelistRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteMcpNetworkWhitelistRulesRequest) GetMcpPolicyId() *string {
	return s.McpPolicyId
}

func (s *DeleteMcpNetworkWhitelistRulesRequest) GetRuleIds() []*string {
	return s.RuleIds
}

func (s *DeleteMcpNetworkWhitelistRulesRequest) SetMcpPolicyId(v string) *DeleteMcpNetworkWhitelistRulesRequest {
	s.McpPolicyId = &v
	return s
}

func (s *DeleteMcpNetworkWhitelistRulesRequest) SetRuleIds(v []*string) *DeleteMcpNetworkWhitelistRulesRequest {
	s.RuleIds = v
	return s
}

func (s *DeleteMcpNetworkWhitelistRulesRequest) Validate() error {
	return dara.Validate(s)
}
