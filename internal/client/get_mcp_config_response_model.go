// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcpConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMcpConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMcpConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetMcpConfigResponseBody) *GetMcpConfigResponse
	GetBody() *GetMcpConfigResponseBody
}

type GetMcpConfigResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMcpConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMcpConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMcpConfigResponse) GoString() string {
	return s.String()
}

func (s *GetMcpConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMcpConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMcpConfigResponse) GetBody() *GetMcpConfigResponseBody {
	return s.Body
}

func (s *GetMcpConfigResponse) SetHeaders(v map[string]*string) *GetMcpConfigResponse {
	s.Headers = v
	return s
}

func (s *GetMcpConfigResponse) SetStatusCode(v int32) *GetMcpConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMcpConfigResponse) SetBody(v *GetMcpConfigResponseBody) *GetMcpConfigResponse {
	s.Body = v
	return s
}

func (s *GetMcpConfigResponse) Validate() error {
	return dara.Validate(s)
}
