// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSessionDurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSessionDurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSessionDurationResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSessionDurationResponseBody) *DescribeSessionDurationResponse
	GetBody() *DescribeSessionDurationResponseBody
}

type DescribeSessionDurationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSessionDurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSessionDurationResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSessionDurationResponse) GoString() string {
	return s.String()
}

func (s *DescribeSessionDurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSessionDurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSessionDurationResponse) GetBody() *DescribeSessionDurationResponseBody {
	return s.Body
}

func (s *DescribeSessionDurationResponse) SetHeaders(v map[string]*string) *DescribeSessionDurationResponse {
	s.Headers = v
	return s
}

func (s *DescribeSessionDurationResponse) SetStatusCode(v int32) *DescribeSessionDurationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSessionDurationResponse) SetBody(v *DescribeSessionDurationResponseBody) *DescribeSessionDurationResponse {
	s.Body = v
	return s
}

func (s *DescribeSessionDurationResponse) Validate() error {
	return dara.Validate(s)
}
