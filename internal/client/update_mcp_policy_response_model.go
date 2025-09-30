// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMcpPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMcpPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMcpPolicyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMcpPolicyResponseBody) *UpdateMcpPolicyResponse
	GetBody() *UpdateMcpPolicyResponseBody
}

type UpdateMcpPolicyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMcpPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMcpPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateMcpPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMcpPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMcpPolicyResponse) GetBody() *UpdateMcpPolicyResponseBody {
	return s.Body
}

func (s *UpdateMcpPolicyResponse) SetHeaders(v map[string]*string) *UpdateMcpPolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateMcpPolicyResponse) SetStatusCode(v int32) *UpdateMcpPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMcpPolicyResponse) SetBody(v *UpdateMcpPolicyResponseBody) *UpdateMcpPolicyResponse {
	s.Body = v
	return s
}

func (s *UpdateMcpPolicyResponse) Validate() error {
	return dara.Validate(s)
}
