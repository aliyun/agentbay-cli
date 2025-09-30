// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBuildMcpImageByParentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BuildMcpImageByParentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BuildMcpImageByParentResponse
	GetStatusCode() *int32
	SetBody(v *BuildMcpImageByParentResponseBody) *BuildMcpImageByParentResponse
	GetBody() *BuildMcpImageByParentResponseBody
}

type BuildMcpImageByParentResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BuildMcpImageByParentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BuildMcpImageByParentResponse) String() string {
	return dara.Prettify(s)
}

func (s BuildMcpImageByParentResponse) GoString() string {
	return s.String()
}

func (s *BuildMcpImageByParentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BuildMcpImageByParentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BuildMcpImageByParentResponse) GetBody() *BuildMcpImageByParentResponseBody {
	return s.Body
}

func (s *BuildMcpImageByParentResponse) SetHeaders(v map[string]*string) *BuildMcpImageByParentResponse {
	s.Headers = v
	return s
}

func (s *BuildMcpImageByParentResponse) SetStatusCode(v int32) *BuildMcpImageByParentResponse {
	s.StatusCode = &v
	return s
}

func (s *BuildMcpImageByParentResponse) SetBody(v *BuildMcpImageByParentResponseBody) *BuildMcpImageByParentResponse {
	s.Body = v
	return s
}

func (s *BuildMcpImageByParentResponse) Validate() error {
	return dara.Validate(s)
}
