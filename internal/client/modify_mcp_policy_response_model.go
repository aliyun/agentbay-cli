// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMcpPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyMcpPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyMcpPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyMcpPolicyResponseBody) *ModifyMcpPolicyResponse
	GetBody() *ModifyMcpPolicyResponseBody
}

type ModifyMcpPolicyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyMcpPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyMcpPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyMcpPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyMcpPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyMcpPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyMcpPolicyResponse) GetBody() *ModifyMcpPolicyResponseBody {
	return s.Body
}

func (s *ModifyMcpPolicyResponse) SetHeaders(v map[string]*string) *ModifyMcpPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyMcpPolicyResponse) SetStatusCode(v int32) *ModifyMcpPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMcpPolicyResponse) SetBody(v *ModifyMcpPolicyResponseBody) *ModifyMcpPolicyResponse {
	s.Body = v
	return s
}

func (s *ModifyMcpPolicyResponse) Validate() error {
	return dara.Validate(s)
}
