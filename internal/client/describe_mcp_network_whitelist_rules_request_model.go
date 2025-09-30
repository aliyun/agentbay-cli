// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMcpNetworkWhitelistRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeMcpNetworkWhitelistRulesRequest
	GetMaxResults() *int32
	SetMcpPolicyId(v string) *DescribeMcpNetworkWhitelistRulesRequest
	GetMcpPolicyId() *string
	SetNextToken(v string) *DescribeMcpNetworkWhitelistRulesRequest
	GetNextToken() *string
}

type DescribeMcpNetworkWhitelistRulesRequest struct {
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// This parameter is required.
	McpPolicyId *string `json:"McpPolicyId,omitempty" xml:"McpPolicyId,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeMcpNetworkWhitelistRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMcpNetworkWhitelistRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeMcpNetworkWhitelistRulesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeMcpNetworkWhitelistRulesRequest) GetMcpPolicyId() *string {
	return s.McpPolicyId
}

func (s *DescribeMcpNetworkWhitelistRulesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeMcpNetworkWhitelistRulesRequest) SetMaxResults(v int32) *DescribeMcpNetworkWhitelistRulesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeMcpNetworkWhitelistRulesRequest) SetMcpPolicyId(v string) *DescribeMcpNetworkWhitelistRulesRequest {
	s.McpPolicyId = &v
	return s
}

func (s *DescribeMcpNetworkWhitelistRulesRequest) SetNextToken(v string) *DescribeMcpNetworkWhitelistRulesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeMcpNetworkWhitelistRulesRequest) Validate() error {
	return dara.Validate(s)
}
