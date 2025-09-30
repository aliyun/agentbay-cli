// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcpDesktopInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMcpDesktopInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMcpDesktopInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetMcpDesktopInfoResponseBody) *GetMcpDesktopInfoResponse
	GetBody() *GetMcpDesktopInfoResponseBody
}

type GetMcpDesktopInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMcpDesktopInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMcpDesktopInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMcpDesktopInfoResponse) GoString() string {
	return s.String()
}

func (s *GetMcpDesktopInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMcpDesktopInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMcpDesktopInfoResponse) GetBody() *GetMcpDesktopInfoResponseBody {
	return s.Body
}

func (s *GetMcpDesktopInfoResponse) SetHeaders(v map[string]*string) *GetMcpDesktopInfoResponse {
	s.Headers = v
	return s
}

func (s *GetMcpDesktopInfoResponse) SetStatusCode(v int32) *GetMcpDesktopInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMcpDesktopInfoResponse) SetBody(v *GetMcpDesktopInfoResponseBody) *GetMcpDesktopInfoResponse {
	s.Body = v
	return s
}

func (s *GetMcpDesktopInfoResponse) Validate() error {
	return dara.Validate(s)
}
