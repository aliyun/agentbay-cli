// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcpResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMcpResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMcpResourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListMcpResourcesResponseBody) *ListMcpResourcesResponse
	GetBody() *ListMcpResourcesResponseBody
}

type ListMcpResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMcpResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMcpResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMcpResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListMcpResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMcpResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMcpResourcesResponse) GetBody() *ListMcpResourcesResponseBody {
	return s.Body
}

func (s *ListMcpResourcesResponse) SetHeaders(v map[string]*string) *ListMcpResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListMcpResourcesResponse) SetStatusCode(v int32) *ListMcpResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMcpResourcesResponse) SetBody(v *ListMcpResourcesResponseBody) *ListMcpResourcesResponse {
	s.Body = v
	return s
}

func (s *ListMcpResourcesResponse) Validate() error {
	return dara.Validate(s)
}
