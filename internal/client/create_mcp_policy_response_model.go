// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcpPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMcpPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMcpPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateMcpPolicyResponseBody) *CreateMcpPolicyResponse
	GetBody() *CreateMcpPolicyResponseBody
}

type CreateMcpPolicyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMcpPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMcpPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateMcpPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMcpPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMcpPolicyResponse) GetBody() *CreateMcpPolicyResponseBody {
	return s.Body
}

func (s *CreateMcpPolicyResponse) SetHeaders(v map[string]*string) *CreateMcpPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateMcpPolicyResponse) SetStatusCode(v int32) *CreateMcpPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMcpPolicyResponse) SetBody(v *CreateMcpPolicyResponseBody) *CreateMcpPolicyResponse {
	s.Body = v
	return s
}

func (s *CreateMcpPolicyResponse) Validate() error {
	return dara.Validate(s)
}
