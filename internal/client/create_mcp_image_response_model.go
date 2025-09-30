// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcpImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMcpImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMcpImageResponse
	GetStatusCode() *int32
	SetBody(v *CreateMcpImageResponseBody) *CreateMcpImageResponse
	GetBody() *CreateMcpImageResponseBody
}

type CreateMcpImageResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMcpImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMcpImageResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpImageResponse) GoString() string {
	return s.String()
}

func (s *CreateMcpImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMcpImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMcpImageResponse) GetBody() *CreateMcpImageResponseBody {
	return s.Body
}

func (s *CreateMcpImageResponse) SetHeaders(v map[string]*string) *CreateMcpImageResponse {
	s.Headers = v
	return s
}

func (s *CreateMcpImageResponse) SetStatusCode(v int32) *CreateMcpImageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMcpImageResponse) SetBody(v *CreateMcpImageResponseBody) *CreateMcpImageResponse {
	s.Body = v
	return s
}

func (s *CreateMcpImageResponse) Validate() error {
	return dara.Validate(s)
}
