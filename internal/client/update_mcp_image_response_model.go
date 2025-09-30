// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMcpImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMcpImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMcpImageResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMcpImageResponseBody) *UpdateMcpImageResponse
	GetBody() *UpdateMcpImageResponseBody
}

type UpdateMcpImageResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMcpImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMcpImageResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpImageResponse) GoString() string {
	return s.String()
}

func (s *UpdateMcpImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMcpImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMcpImageResponse) GetBody() *UpdateMcpImageResponseBody {
	return s.Body
}

func (s *UpdateMcpImageResponse) SetHeaders(v map[string]*string) *UpdateMcpImageResponse {
	s.Headers = v
	return s
}

func (s *UpdateMcpImageResponse) SetStatusCode(v int32) *UpdateMcpImageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMcpImageResponse) SetBody(v *UpdateMcpImageResponseBody) *UpdateMcpImageResponse {
	s.Body = v
	return s
}

func (s *UpdateMcpImageResponse) Validate() error {
	return dara.Validate(s)
}
