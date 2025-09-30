// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindMcpPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindMcpPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindMcpPolicyResponse
	GetStatusCode() *int32
	SetBody(v *UnbindMcpPolicyResponseBody) *UnbindMcpPolicyResponse
	GetBody() *UnbindMcpPolicyResponseBody
}

type UnbindMcpPolicyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindMcpPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindMcpPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindMcpPolicyResponse) GoString() string {
	return s.String()
}

func (s *UnbindMcpPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindMcpPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindMcpPolicyResponse) GetBody() *UnbindMcpPolicyResponseBody {
	return s.Body
}

func (s *UnbindMcpPolicyResponse) SetHeaders(v map[string]*string) *UnbindMcpPolicyResponse {
	s.Headers = v
	return s
}

func (s *UnbindMcpPolicyResponse) SetStatusCode(v int32) *UnbindMcpPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindMcpPolicyResponse) SetBody(v *UnbindMcpPolicyResponseBody) *UnbindMcpPolicyResponse {
	s.Body = v
	return s
}

func (s *UnbindMcpPolicyResponse) Validate() error {
	return dara.Validate(s)
}
