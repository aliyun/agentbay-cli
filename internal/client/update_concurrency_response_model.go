// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConcurrencyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateConcurrencyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateConcurrencyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateConcurrencyResponseBody) *UpdateConcurrencyResponse
	GetBody() *UpdateConcurrencyResponseBody
}

type UpdateConcurrencyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateConcurrencyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateConcurrencyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateConcurrencyResponse) GoString() string {
	return s.String()
}

func (s *UpdateConcurrencyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateConcurrencyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateConcurrencyResponse) GetBody() *UpdateConcurrencyResponseBody {
	return s.Body
}

func (s *UpdateConcurrencyResponse) SetHeaders(v map[string]*string) *UpdateConcurrencyResponse {
	s.Headers = v
	return s
}

func (s *UpdateConcurrencyResponse) SetStatusCode(v int32) *UpdateConcurrencyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateConcurrencyResponse) SetBody(v *UpdateConcurrencyResponseBody) *UpdateConcurrencyResponse {
	s.Body = v
	return s
}

func (s *UpdateConcurrencyResponse) Validate() error {
	return dara.Validate(s)
}
