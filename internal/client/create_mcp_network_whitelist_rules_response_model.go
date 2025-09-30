// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcpNetworkWhitelistRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMcpNetworkWhitelistRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMcpNetworkWhitelistRulesResponse
	GetStatusCode() *int32
	SetBody(v *CreateMcpNetworkWhitelistRulesResponseBody) *CreateMcpNetworkWhitelistRulesResponse
	GetBody() *CreateMcpNetworkWhitelistRulesResponseBody
}

type CreateMcpNetworkWhitelistRulesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMcpNetworkWhitelistRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMcpNetworkWhitelistRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpNetworkWhitelistRulesResponse) GoString() string {
	return s.String()
}

func (s *CreateMcpNetworkWhitelistRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMcpNetworkWhitelistRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMcpNetworkWhitelistRulesResponse) GetBody() *CreateMcpNetworkWhitelistRulesResponseBody {
	return s.Body
}

func (s *CreateMcpNetworkWhitelistRulesResponse) SetHeaders(v map[string]*string) *CreateMcpNetworkWhitelistRulesResponse {
	s.Headers = v
	return s
}

func (s *CreateMcpNetworkWhitelistRulesResponse) SetStatusCode(v int32) *CreateMcpNetworkWhitelistRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMcpNetworkWhitelistRulesResponse) SetBody(v *CreateMcpNetworkWhitelistRulesResponseBody) *CreateMcpNetworkWhitelistRulesResponse {
	s.Body = v
	return s
}

func (s *CreateMcpNetworkWhitelistRulesResponse) Validate() error {
	return dara.Validate(s)
}
