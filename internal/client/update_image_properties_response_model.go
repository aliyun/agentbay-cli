// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateImagePropertiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateImagePropertiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateImagePropertiesResponse
	GetStatusCode() *int32
	SetBody(v *UpdateImagePropertiesResponseBody) *UpdateImagePropertiesResponse
	GetBody() *UpdateImagePropertiesResponseBody
}

type UpdateImagePropertiesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateImagePropertiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateImagePropertiesResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateImagePropertiesResponse) GoString() string {
	return s.String()
}

func (s *UpdateImagePropertiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateImagePropertiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateImagePropertiesResponse) GetBody() *UpdateImagePropertiesResponseBody {
	return s.Body
}

func (s *UpdateImagePropertiesResponse) SetHeaders(v map[string]*string) *UpdateImagePropertiesResponse {
	s.Headers = v
	return s
}

func (s *UpdateImagePropertiesResponse) SetStatusCode(v int32) *UpdateImagePropertiesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateImagePropertiesResponse) SetBody(v *UpdateImagePropertiesResponseBody) *UpdateImagePropertiesResponse {
	s.Body = v
	return s
}

func (s *UpdateImagePropertiesResponse) Validate() error {
	return dara.Validate(s)
}
