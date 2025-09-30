// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcpImageInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMcpImageInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMcpImageInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetMcpImageInfoResponseBody) *GetMcpImageInfoResponse
	GetBody() *GetMcpImageInfoResponseBody
}

type GetMcpImageInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMcpImageInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMcpImageInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMcpImageInfoResponse) GoString() string {
	return s.String()
}

func (s *GetMcpImageInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMcpImageInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMcpImageInfoResponse) GetBody() *GetMcpImageInfoResponseBody {
	return s.Body
}

func (s *GetMcpImageInfoResponse) SetHeaders(v map[string]*string) *GetMcpImageInfoResponse {
	s.Headers = v
	return s
}

func (s *GetMcpImageInfoResponse) SetStatusCode(v int32) *GetMcpImageInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMcpImageInfoResponse) SetBody(v *GetMcpImageInfoResponseBody) *GetMcpImageInfoResponse {
	s.Body = v
	return s
}

func (s *GetMcpImageInfoResponse) Validate() error {
	return dara.Validate(s)
}
