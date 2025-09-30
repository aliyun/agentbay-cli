// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMcpNetworkWhitelistRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMcpNetworkWhitelistRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMcpNetworkWhitelistRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMcpNetworkWhitelistRulesResponseBody) *DescribeMcpNetworkWhitelistRulesResponse
	GetBody() *DescribeMcpNetworkWhitelistRulesResponseBody
}

type DescribeMcpNetworkWhitelistRulesResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMcpNetworkWhitelistRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMcpNetworkWhitelistRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMcpNetworkWhitelistRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeMcpNetworkWhitelistRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMcpNetworkWhitelistRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMcpNetworkWhitelistRulesResponse) GetBody() *DescribeMcpNetworkWhitelistRulesResponseBody {
	return s.Body
}

func (s *DescribeMcpNetworkWhitelistRulesResponse) SetHeaders(v map[string]*string) *DescribeMcpNetworkWhitelistRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeMcpNetworkWhitelistRulesResponse) SetStatusCode(v int32) *DescribeMcpNetworkWhitelistRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMcpNetworkWhitelistRulesResponse) SetBody(v *DescribeMcpNetworkWhitelistRulesResponseBody) *DescribeMcpNetworkWhitelistRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeMcpNetworkWhitelistRulesResponse) Validate() error {
	return dara.Validate(s)
}
