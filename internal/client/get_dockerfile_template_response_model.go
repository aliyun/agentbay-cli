// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDockerfileTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDockerfileTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDockerfileTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetDockerfileTemplateResponseBody) *GetDockerfileTemplateResponse
	GetBody() *GetDockerfileTemplateResponseBody
}

type GetDockerfileTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDockerfileTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDockerfileTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDockerfileTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetDockerfileTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDockerfileTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDockerfileTemplateResponse) GetBody() *GetDockerfileTemplateResponseBody {
	return s.Body
}

func (s *GetDockerfileTemplateResponse) SetHeaders(v map[string]*string) *GetDockerfileTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetDockerfileTemplateResponse) SetStatusCode(v int32) *GetDockerfileTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDockerfileTemplateResponse) SetBody(v *GetDockerfileTemplateResponseBody) *GetDockerfileTemplateResponse {
	s.Body = v
	return s
}

func (s *GetDockerfileTemplateResponse) Validate() error {
	return dara.Validate(s)
}
