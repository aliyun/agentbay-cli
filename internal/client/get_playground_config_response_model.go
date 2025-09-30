// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPlaygroundConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPlaygroundConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPlaygroundConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetPlaygroundConfigResponseBody) *GetPlaygroundConfigResponse
	GetBody() *GetPlaygroundConfigResponseBody
}

type GetPlaygroundConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPlaygroundConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPlaygroundConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPlaygroundConfigResponse) GoString() string {
	return s.String()
}

func (s *GetPlaygroundConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPlaygroundConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPlaygroundConfigResponse) GetBody() *GetPlaygroundConfigResponseBody {
	return s.Body
}

func (s *GetPlaygroundConfigResponse) SetHeaders(v map[string]*string) *GetPlaygroundConfigResponse {
	s.Headers = v
	return s
}

func (s *GetPlaygroundConfigResponse) SetStatusCode(v int32) *GetPlaygroundConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPlaygroundConfigResponse) SetBody(v *GetPlaygroundConfigResponseBody) *GetPlaygroundConfigResponse {
	s.Body = v
	return s
}

func (s *GetPlaygroundConfigResponse) Validate() error {
	return dara.Validate(s)
}
