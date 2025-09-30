// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcpToolInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *GetMcpToolInfoRequest
	GetImageId() *string
	SetMaxResults(v int32) *GetMcpToolInfoRequest
	GetMaxResults() *int32
	SetNextToken(v string) *GetMcpToolInfoRequest
	GetNextToken() *string
}

type GetMcpToolInfoRequest struct {
	ImageId    *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s GetMcpToolInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMcpToolInfoRequest) GoString() string {
	return s.String()
}

func (s *GetMcpToolInfoRequest) GetImageId() *string {
	return s.ImageId
}

func (s *GetMcpToolInfoRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetMcpToolInfoRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetMcpToolInfoRequest) SetImageId(v string) *GetMcpToolInfoRequest {
	s.ImageId = &v
	return s
}

func (s *GetMcpToolInfoRequest) SetMaxResults(v int32) *GetMcpToolInfoRequest {
	s.MaxResults = &v
	return s
}

func (s *GetMcpToolInfoRequest) SetNextToken(v string) *GetMcpToolInfoRequest {
	s.NextToken = &v
	return s
}

func (s *GetMcpToolInfoRequest) Validate() error {
	return dara.Validate(s)
}
