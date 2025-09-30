// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcpImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMcpImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMcpImageResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMcpImageResponseBody) *DeleteMcpImageResponse
	GetBody() *DeleteMcpImageResponseBody
}

type DeleteMcpImageResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMcpImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMcpImageResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcpImageResponse) GoString() string {
	return s.String()
}

func (s *DeleteMcpImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMcpImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMcpImageResponse) GetBody() *DeleteMcpImageResponseBody {
	return s.Body
}

func (s *DeleteMcpImageResponse) SetHeaders(v map[string]*string) *DeleteMcpImageResponse {
	s.Headers = v
	return s
}

func (s *DeleteMcpImageResponse) SetStatusCode(v int32) *DeleteMcpImageResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMcpImageResponse) SetBody(v *DeleteMcpImageResponseBody) *DeleteMcpImageResponse {
	s.Body = v
	return s
}

func (s *DeleteMcpImageResponse) Validate() error {
	return dara.Validate(s)
}
