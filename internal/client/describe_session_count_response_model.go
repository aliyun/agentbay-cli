// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSessionCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSessionCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSessionCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSessionCountResponseBody) *DescribeSessionCountResponse
	GetBody() *DescribeSessionCountResponseBody
}

type DescribeSessionCountResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSessionCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSessionCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSessionCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeSessionCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSessionCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSessionCountResponse) GetBody() *DescribeSessionCountResponseBody {
	return s.Body
}

func (s *DescribeSessionCountResponse) SetHeaders(v map[string]*string) *DescribeSessionCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeSessionCountResponse) SetStatusCode(v int32) *DescribeSessionCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSessionCountResponse) SetBody(v *DescribeSessionCountResponseBody) *DescribeSessionCountResponse {
	s.Body = v
	return s
}

func (s *DescribeSessionCountResponse) Validate() error {
	return dara.Validate(s)
}
