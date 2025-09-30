// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcpImagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMcpImagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMcpImagesResponse
	GetStatusCode() *int32
	SetBody(v *ListMcpImagesResponseBody) *ListMcpImagesResponse
	GetBody() *ListMcpImagesResponseBody
}

type ListMcpImagesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMcpImagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMcpImagesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMcpImagesResponse) GoString() string {
	return s.String()
}

func (s *ListMcpImagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMcpImagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMcpImagesResponse) GetBody() *ListMcpImagesResponseBody {
	return s.Body
}

func (s *ListMcpImagesResponse) SetHeaders(v map[string]*string) *ListMcpImagesResponse {
	s.Headers = v
	return s
}

func (s *ListMcpImagesResponse) SetStatusCode(v int32) *ListMcpImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMcpImagesResponse) SetBody(v *ListMcpImagesResponseBody) *ListMcpImagesResponse {
	s.Body = v
	return s
}

func (s *ListMcpImagesResponse) Validate() error {
	return dara.Validate(s)
}
