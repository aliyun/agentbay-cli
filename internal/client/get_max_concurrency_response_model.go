// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMaxConcurrencyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMaxConcurrencyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMaxConcurrencyResponse
	GetStatusCode() *int32
	SetBody(v *GetMaxConcurrencyResponseBody) *GetMaxConcurrencyResponse
	GetBody() *GetMaxConcurrencyResponseBody
}

type GetMaxConcurrencyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMaxConcurrencyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMaxConcurrencyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMaxConcurrencyResponse) GoString() string {
	return s.String()
}

func (s *GetMaxConcurrencyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMaxConcurrencyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMaxConcurrencyResponse) GetBody() *GetMaxConcurrencyResponseBody {
	return s.Body
}

func (s *GetMaxConcurrencyResponse) SetHeaders(v map[string]*string) *GetMaxConcurrencyResponse {
	s.Headers = v
	return s
}

func (s *GetMaxConcurrencyResponse) SetStatusCode(v int32) *GetMaxConcurrencyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMaxConcurrencyResponse) SetBody(v *GetMaxConcurrencyResponseBody) *GetMaxConcurrencyResponse {
	s.Body = v
	return s
}

func (s *GetMaxConcurrencyResponse) Validate() error {
	return dara.Validate(s)
}
