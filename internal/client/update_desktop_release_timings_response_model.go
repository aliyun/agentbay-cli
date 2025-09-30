// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDesktopReleaseTimingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDesktopReleaseTimingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDesktopReleaseTimingsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDesktopReleaseTimingsResponseBody) *UpdateDesktopReleaseTimingsResponse
	GetBody() *UpdateDesktopReleaseTimingsResponseBody
}

type UpdateDesktopReleaseTimingsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDesktopReleaseTimingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDesktopReleaseTimingsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDesktopReleaseTimingsResponse) GoString() string {
	return s.String()
}

func (s *UpdateDesktopReleaseTimingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDesktopReleaseTimingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDesktopReleaseTimingsResponse) GetBody() *UpdateDesktopReleaseTimingsResponseBody {
	return s.Body
}

func (s *UpdateDesktopReleaseTimingsResponse) SetHeaders(v map[string]*string) *UpdateDesktopReleaseTimingsResponse {
	s.Headers = v
	return s
}

func (s *UpdateDesktopReleaseTimingsResponse) SetStatusCode(v int32) *UpdateDesktopReleaseTimingsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDesktopReleaseTimingsResponse) SetBody(v *UpdateDesktopReleaseTimingsResponseBody) *UpdateDesktopReleaseTimingsResponse {
	s.Body = v
	return s
}

func (s *UpdateDesktopReleaseTimingsResponse) Validate() error {
	return dara.Validate(s)
}
