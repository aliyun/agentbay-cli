// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListResourceGroupsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListResourceGroupsRequest
	GetNextToken() *string
}

type ListResourceGroupsRequest struct {
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListResourceGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListResourceGroupsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResourceGroupsRequest) SetMaxResults(v int32) *ListResourceGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListResourceGroupsRequest) SetNextToken(v string) *ListResourceGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListResourceGroupsRequest) Validate() error {
	return dara.Validate(s)
}
