// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMcpPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExcludedPolicyIds(v []*string) *DescribeMcpPoliciesRequest
	GetExcludedPolicyIds() []*string
	SetMaxResults(v int32) *DescribeMcpPoliciesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeMcpPoliciesRequest
	GetNextToken() *string
	SetPolicyIds(v []*string) *DescribeMcpPoliciesRequest
	GetPolicyIds() []*string
	SetPolicyName(v string) *DescribeMcpPoliciesRequest
	GetPolicyName() *string
}

type DescribeMcpPoliciesRequest struct {
	ExcludedPolicyIds []*string `json:"ExcludedPolicyIds,omitempty" xml:"ExcludedPolicyIds,omitempty" type:"Repeated"`
	MaxResults        *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken         *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PolicyIds         []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	PolicyName        *string   `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
}

func (s DescribeMcpPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMcpPoliciesRequest) GoString() string {
	return s.String()
}

func (s *DescribeMcpPoliciesRequest) GetExcludedPolicyIds() []*string {
	return s.ExcludedPolicyIds
}

func (s *DescribeMcpPoliciesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeMcpPoliciesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeMcpPoliciesRequest) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *DescribeMcpPoliciesRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *DescribeMcpPoliciesRequest) SetExcludedPolicyIds(v []*string) *DescribeMcpPoliciesRequest {
	s.ExcludedPolicyIds = v
	return s
}

func (s *DescribeMcpPoliciesRequest) SetMaxResults(v int32) *DescribeMcpPoliciesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeMcpPoliciesRequest) SetNextToken(v string) *DescribeMcpPoliciesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeMcpPoliciesRequest) SetPolicyIds(v []*string) *DescribeMcpPoliciesRequest {
	s.PolicyIds = v
	return s
}

func (s *DescribeMcpPoliciesRequest) SetPolicyName(v string) *DescribeMcpPoliciesRequest {
	s.PolicyName = &v
	return s
}

func (s *DescribeMcpPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
