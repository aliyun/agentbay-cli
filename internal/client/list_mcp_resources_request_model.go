// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcpResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListMcpResourcesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListMcpResourcesRequest
	GetNextToken() *string
}

type ListMcpResourcesRequest struct {
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListMcpResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMcpResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListMcpResourcesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMcpResourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMcpResourcesRequest) SetMaxResults(v int32) *ListMcpResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMcpResourcesRequest) SetNextToken(v string) *ListMcpResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListMcpResourcesRequest) Validate() error {
	return dara.Validate(s)
}
