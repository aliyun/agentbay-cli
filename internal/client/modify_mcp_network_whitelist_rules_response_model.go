// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMcpNetworkWhitelistRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyMcpNetworkWhitelistRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyMcpNetworkWhitelistRulesResponse
	GetStatusCode() *int32
	SetBody(v *ModifyMcpNetworkWhitelistRulesResponseBody) *ModifyMcpNetworkWhitelistRulesResponse
	GetBody() *ModifyMcpNetworkWhitelistRulesResponseBody
}

type ModifyMcpNetworkWhitelistRulesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyMcpNetworkWhitelistRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyMcpNetworkWhitelistRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyMcpNetworkWhitelistRulesResponse) GoString() string {
	return s.String()
}

func (s *ModifyMcpNetworkWhitelistRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyMcpNetworkWhitelistRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyMcpNetworkWhitelistRulesResponse) GetBody() *ModifyMcpNetworkWhitelistRulesResponseBody {
	return s.Body
}

func (s *ModifyMcpNetworkWhitelistRulesResponse) SetHeaders(v map[string]*string) *ModifyMcpNetworkWhitelistRulesResponse {
	s.Headers = v
	return s
}

func (s *ModifyMcpNetworkWhitelistRulesResponse) SetStatusCode(v int32) *ModifyMcpNetworkWhitelistRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMcpNetworkWhitelistRulesResponse) SetBody(v *ModifyMcpNetworkWhitelistRulesResponseBody) *ModifyMcpNetworkWhitelistRulesResponse {
	s.Body = v
	return s
}

func (s *ModifyMcpNetworkWhitelistRulesResponse) Validate() error {
	return dara.Validate(s)
}
