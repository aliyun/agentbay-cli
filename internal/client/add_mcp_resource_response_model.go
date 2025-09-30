// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMcpResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddMcpResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddMcpResourceResponse
	GetStatusCode() *int32
	SetBody(v *AddMcpResourceResponseBody) *AddMcpResourceResponse
	GetBody() *AddMcpResourceResponseBody
}

type AddMcpResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddMcpResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddMcpResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s AddMcpResourceResponse) GoString() string {
	return s.String()
}

func (s *AddMcpResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddMcpResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddMcpResourceResponse) GetBody() *AddMcpResourceResponseBody {
	return s.Body
}

func (s *AddMcpResourceResponse) SetHeaders(v map[string]*string) *AddMcpResourceResponse {
	s.Headers = v
	return s
}

func (s *AddMcpResourceResponse) SetStatusCode(v int32) *AddMcpResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *AddMcpResourceResponse) SetBody(v *AddMcpResourceResponseBody) *AddMcpResourceResponse {
	s.Body = v
	return s
}

func (s *AddMcpResourceResponse) Validate() error {
	return dara.Validate(s)
}
