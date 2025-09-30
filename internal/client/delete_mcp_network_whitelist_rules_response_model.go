// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcpNetworkWhitelistRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMcpNetworkWhitelistRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMcpNetworkWhitelistRulesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMcpNetworkWhitelistRulesResponseBody) *DeleteMcpNetworkWhitelistRulesResponse
	GetBody() *DeleteMcpNetworkWhitelistRulesResponseBody
}

type DeleteMcpNetworkWhitelistRulesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMcpNetworkWhitelistRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMcpNetworkWhitelistRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcpNetworkWhitelistRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteMcpNetworkWhitelistRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMcpNetworkWhitelistRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMcpNetworkWhitelistRulesResponse) GetBody() *DeleteMcpNetworkWhitelistRulesResponseBody {
	return s.Body
}

func (s *DeleteMcpNetworkWhitelistRulesResponse) SetHeaders(v map[string]*string) *DeleteMcpNetworkWhitelistRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteMcpNetworkWhitelistRulesResponse) SetStatusCode(v int32) *DeleteMcpNetworkWhitelistRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMcpNetworkWhitelistRulesResponse) SetBody(v *DeleteMcpNetworkWhitelistRulesResponseBody) *DeleteMcpNetworkWhitelistRulesResponse {
	s.Body = v
	return s
}

func (s *DeleteMcpNetworkWhitelistRulesResponse) Validate() error {
	return dara.Validate(s)
}
