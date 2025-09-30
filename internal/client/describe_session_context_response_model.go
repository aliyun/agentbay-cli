// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSessionContextResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSessionContextResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSessionContextResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSessionContextResponseBody) *DescribeSessionContextResponse
	GetBody() *DescribeSessionContextResponseBody
}

type DescribeSessionContextResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSessionContextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSessionContextResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSessionContextResponse) GoString() string {
	return s.String()
}

func (s *DescribeSessionContextResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSessionContextResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSessionContextResponse) GetBody() *DescribeSessionContextResponseBody {
	return s.Body
}

func (s *DescribeSessionContextResponse) SetHeaders(v map[string]*string) *DescribeSessionContextResponse {
	s.Headers = v
	return s
}

func (s *DescribeSessionContextResponse) SetStatusCode(v int32) *DescribeSessionContextResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSessionContextResponse) SetBody(v *DescribeSessionContextResponseBody) *DescribeSessionContextResponse {
	s.Body = v
	return s
}

func (s *DescribeSessionContextResponse) Validate() error {
	return dara.Validate(s)
}
