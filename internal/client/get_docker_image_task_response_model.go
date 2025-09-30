// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDockerImageTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDockerImageTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDockerImageTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetDockerImageTaskResponseBody) *GetDockerImageTaskResponse
	GetBody() *GetDockerImageTaskResponseBody
}

type GetDockerImageTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDockerImageTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDockerImageTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDockerImageTaskResponse) GoString() string {
	return s.String()
}

func (s *GetDockerImageTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDockerImageTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDockerImageTaskResponse) GetBody() *GetDockerImageTaskResponseBody {
	return s.Body
}

func (s *GetDockerImageTaskResponse) SetHeaders(v map[string]*string) *GetDockerImageTaskResponse {
	s.Headers = v
	return s
}

func (s *GetDockerImageTaskResponse) SetStatusCode(v int32) *GetDockerImageTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDockerImageTaskResponse) SetBody(v *GetDockerImageTaskResponseBody) *GetDockerImageTaskResponse {
	s.Body = v
	return s
}

func (s *GetDockerImageTaskResponse) Validate() error {
	return dara.Validate(s)
}
