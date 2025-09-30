// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMcpApiKeyConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyMcpApiKeyConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyMcpApiKeyConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyMcpApiKeyConfigResponseBody) *ModifyMcpApiKeyConfigResponse
	GetBody() *ModifyMcpApiKeyConfigResponseBody
}

type ModifyMcpApiKeyConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyMcpApiKeyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyMcpApiKeyConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyMcpApiKeyConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyMcpApiKeyConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyMcpApiKeyConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyMcpApiKeyConfigResponse) GetBody() *ModifyMcpApiKeyConfigResponseBody {
	return s.Body
}

func (s *ModifyMcpApiKeyConfigResponse) SetHeaders(v map[string]*string) *ModifyMcpApiKeyConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyMcpApiKeyConfigResponse) SetStatusCode(v int32) *ModifyMcpApiKeyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMcpApiKeyConfigResponse) SetBody(v *ModifyMcpApiKeyConfigResponseBody) *ModifyMcpApiKeyConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyMcpApiKeyConfigResponse) Validate() error {
	return dara.Validate(s)
}
