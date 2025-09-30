// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSessionCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v string) *DescribeSessionCountRequest
	GetApiKeyId() *string
	SetImageId(v string) *DescribeSessionCountRequest
	GetImageId() *string
}

type DescribeSessionCountRequest struct {
	ApiKeyId *string `json:"ApiKeyId,omitempty" xml:"ApiKeyId,omitempty"`
	ImageId  *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
}

func (s DescribeSessionCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSessionCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeSessionCountRequest) GetApiKeyId() *string {
	return s.ApiKeyId
}

func (s *DescribeSessionCountRequest) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeSessionCountRequest) SetApiKeyId(v string) *DescribeSessionCountRequest {
	s.ApiKeyId = &v
	return s
}

func (s *DescribeSessionCountRequest) SetImageId(v string) *DescribeSessionCountRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeSessionCountRequest) Validate() error {
	return dara.Validate(s)
}
