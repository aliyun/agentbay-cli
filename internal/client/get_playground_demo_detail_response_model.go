// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPlaygroundDemoDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPlaygroundDemoDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPlaygroundDemoDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetPlaygroundDemoDetailResponseBody) *GetPlaygroundDemoDetailResponse
	GetBody() *GetPlaygroundDemoDetailResponseBody
}

type GetPlaygroundDemoDetailResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPlaygroundDemoDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPlaygroundDemoDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPlaygroundDemoDetailResponse) GoString() string {
	return s.String()
}

func (s *GetPlaygroundDemoDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPlaygroundDemoDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPlaygroundDemoDetailResponse) GetBody() *GetPlaygroundDemoDetailResponseBody {
	return s.Body
}

func (s *GetPlaygroundDemoDetailResponse) SetHeaders(v map[string]*string) *GetPlaygroundDemoDetailResponse {
	s.Headers = v
	return s
}

func (s *GetPlaygroundDemoDetailResponse) SetStatusCode(v int32) *GetPlaygroundDemoDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPlaygroundDemoDetailResponse) SetBody(v *GetPlaygroundDemoDetailResponseBody) *GetPlaygroundDemoDetailResponse {
	s.Body = v
	return s
}

func (s *GetPlaygroundDemoDetailResponse) Validate() error {
	return dara.Validate(s)
}
