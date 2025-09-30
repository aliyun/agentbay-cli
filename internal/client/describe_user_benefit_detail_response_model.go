// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserBenefitDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserBenefitDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserBenefitDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserBenefitDetailResponseBody) *DescribeUserBenefitDetailResponse
	GetBody() *DescribeUserBenefitDetailResponseBody
}

type DescribeUserBenefitDetailResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserBenefitDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserBenefitDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserBenefitDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserBenefitDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserBenefitDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserBenefitDetailResponse) GetBody() *DescribeUserBenefitDetailResponseBody {
	return s.Body
}

func (s *DescribeUserBenefitDetailResponse) SetHeaders(v map[string]*string) *DescribeUserBenefitDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserBenefitDetailResponse) SetStatusCode(v int32) *DescribeUserBenefitDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserBenefitDetailResponse) SetBody(v *DescribeUserBenefitDetailResponseBody) *DescribeUserBenefitDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeUserBenefitDetailResponse) Validate() error {
	return dara.Validate(s)
}
