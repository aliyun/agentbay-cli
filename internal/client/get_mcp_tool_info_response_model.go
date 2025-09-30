// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcpToolInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMcpToolInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMcpToolInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetMcpToolInfoResponseBody) *GetMcpToolInfoResponse
	GetBody() *GetMcpToolInfoResponseBody
}

type GetMcpToolInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMcpToolInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMcpToolInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMcpToolInfoResponse) GoString() string {
	return s.String()
}

func (s *GetMcpToolInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMcpToolInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMcpToolInfoResponse) GetBody() *GetMcpToolInfoResponseBody {
	return s.Body
}

func (s *GetMcpToolInfoResponse) SetHeaders(v map[string]*string) *GetMcpToolInfoResponse {
	s.Headers = v
	return s
}

func (s *GetMcpToolInfoResponse) SetStatusCode(v int32) *GetMcpToolInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMcpToolInfoResponse) SetBody(v *GetMcpToolInfoResponseBody) *GetMcpToolInfoResponse {
	s.Body = v
	return s
}

func (s *GetMcpToolInfoResponse) Validate() error {
	return dara.Validate(s)
}
