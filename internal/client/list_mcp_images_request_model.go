// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcpImagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureList(v []*string) *ListMcpImagesRequest
	GetFeatureList() []*string
	SetImageType(v string) *ListMcpImagesRequest
	GetImageType() *string
	SetMaxResults(v int32) *ListMcpImagesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListMcpImagesRequest
	GetNextToken() *string
	SetOsType(v string) *ListMcpImagesRequest
	GetOsType() *string
	SetPageSize(v int32) *ListMcpImagesRequest
	GetPageSize() *int32
	SetPageStart(v int32) *ListMcpImagesRequest
	GetPageStart() *int32
}

type ListMcpImagesRequest struct {
	FeatureList []*string `json:"FeatureList,omitempty" xml:"FeatureList,omitempty" type:"Repeated"`
	ImageType   *string   `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	MaxResults  *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OsType      *string   `json:"OsType,omitempty" xml:"OsType,omitempty"`
	PageSize    *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageStart   *int32    `json:"PageStart,omitempty" xml:"PageStart,omitempty"`
}

func (s ListMcpImagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMcpImagesRequest) GoString() string {
	return s.String()
}

func (s *ListMcpImagesRequest) GetFeatureList() []*string {
	return s.FeatureList
}

func (s *ListMcpImagesRequest) GetImageType() *string {
	return s.ImageType
}

func (s *ListMcpImagesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMcpImagesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMcpImagesRequest) GetOsType() *string {
	return s.OsType
}

func (s *ListMcpImagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMcpImagesRequest) GetPageStart() *int32 {
	return s.PageStart
}

func (s *ListMcpImagesRequest) SetFeatureList(v []*string) *ListMcpImagesRequest {
	s.FeatureList = v
	return s
}

func (s *ListMcpImagesRequest) SetImageType(v string) *ListMcpImagesRequest {
	s.ImageType = &v
	return s
}

func (s *ListMcpImagesRequest) SetMaxResults(v int32) *ListMcpImagesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMcpImagesRequest) SetNextToken(v string) *ListMcpImagesRequest {
	s.NextToken = &v
	return s
}

func (s *ListMcpImagesRequest) SetOsType(v string) *ListMcpImagesRequest {
	s.OsType = &v
	return s
}

func (s *ListMcpImagesRequest) SetPageSize(v int32) *ListMcpImagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListMcpImagesRequest) SetPageStart(v int32) *ListMcpImagesRequest {
	s.PageStart = &v
	return s
}

func (s *ListMcpImagesRequest) Validate() error {
	return dara.Validate(s)
}
