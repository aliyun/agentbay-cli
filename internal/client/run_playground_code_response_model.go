// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunPlaygroundCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunPlaygroundCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunPlaygroundCodeResponse
	GetStatusCode() *int32
	SetBody(v *RunPlaygroundCodeResponseBody) *RunPlaygroundCodeResponse
	GetBody() *RunPlaygroundCodeResponseBody
}

type RunPlaygroundCodeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunPlaygroundCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunPlaygroundCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s RunPlaygroundCodeResponse) GoString() string {
	return s.String()
}

func (s *RunPlaygroundCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunPlaygroundCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunPlaygroundCodeResponse) GetBody() *RunPlaygroundCodeResponseBody {
	return s.Body
}

func (s *RunPlaygroundCodeResponse) SetHeaders(v map[string]*string) *RunPlaygroundCodeResponse {
	s.Headers = v
	return s
}

func (s *RunPlaygroundCodeResponse) SetStatusCode(v int32) *RunPlaygroundCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *RunPlaygroundCodeResponse) SetBody(v *RunPlaygroundCodeResponseBody) *RunPlaygroundCodeResponse {
	s.Body = v
	return s
}

func (s *RunPlaygroundCodeResponse) Validate() error {
	return dara.Validate(s)
}
