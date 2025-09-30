// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDelMcpResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DelMcpResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DelMcpResourceResponse
	GetStatusCode() *int32
	SetBody(v *DelMcpResourceResponseBody) *DelMcpResourceResponse
	GetBody() *DelMcpResourceResponseBody
}

type DelMcpResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DelMcpResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DelMcpResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DelMcpResourceResponse) GoString() string {
	return s.String()
}

func (s *DelMcpResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DelMcpResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DelMcpResourceResponse) GetBody() *DelMcpResourceResponseBody {
	return s.Body
}

func (s *DelMcpResourceResponse) SetHeaders(v map[string]*string) *DelMcpResourceResponse {
	s.Headers = v
	return s
}

func (s *DelMcpResourceResponse) SetStatusCode(v int32) *DelMcpResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DelMcpResourceResponse) SetBody(v *DelMcpResourceResponseBody) *DelMcpResourceResponse {
	s.Body = v
	return s
}

func (s *DelMcpResourceResponse) Validate() error {
	return dara.Validate(s)
}
