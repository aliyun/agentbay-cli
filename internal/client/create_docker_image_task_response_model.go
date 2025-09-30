// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDockerImageTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDockerImageTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDockerImageTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateDockerImageTaskResponseBody) *CreateDockerImageTaskResponse
	GetBody() *CreateDockerImageTaskResponseBody
}

type CreateDockerImageTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDockerImageTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDockerImageTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDockerImageTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDockerImageTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDockerImageTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDockerImageTaskResponse) GetBody() *CreateDockerImageTaskResponseBody {
	return s.Body
}

func (s *CreateDockerImageTaskResponse) SetHeaders(v map[string]*string) *CreateDockerImageTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateDockerImageTaskResponse) SetStatusCode(v int32) *CreateDockerImageTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDockerImageTaskResponse) SetBody(v *CreateDockerImageTaskResponseBody) *CreateDockerImageTaskResponse {
	s.Body = v
	return s
}

func (s *CreateDockerImageTaskResponse) Validate() error {
	return dara.Validate(s)
}
