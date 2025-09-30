// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcpPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMcpPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMcpPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMcpPoliciesResponseBody) *DeleteMcpPoliciesResponse
	GetBody() *DeleteMcpPoliciesResponseBody
}

type DeleteMcpPoliciesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMcpPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMcpPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcpPoliciesResponse) GoString() string {
	return s.String()
}

func (s *DeleteMcpPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMcpPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMcpPoliciesResponse) GetBody() *DeleteMcpPoliciesResponseBody {
	return s.Body
}

func (s *DeleteMcpPoliciesResponse) SetHeaders(v map[string]*string) *DeleteMcpPoliciesResponse {
	s.Headers = v
	return s
}

func (s *DeleteMcpPoliciesResponse) SetStatusCode(v int32) *DeleteMcpPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMcpPoliciesResponse) SetBody(v *DeleteMcpPoliciesResponseBody) *DeleteMcpPoliciesResponse {
	s.Body = v
	return s
}

func (s *DeleteMcpPoliciesResponse) Validate() error {
	return dara.Validate(s)
}
