// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSessionToolLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSessionToolLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSessionToolLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListSessionToolLogsResponseBody) *ListSessionToolLogsResponse
	GetBody() *ListSessionToolLogsResponseBody
}

type ListSessionToolLogsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSessionToolLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSessionToolLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSessionToolLogsResponse) GoString() string {
	return s.String()
}

func (s *ListSessionToolLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSessionToolLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSessionToolLogsResponse) GetBody() *ListSessionToolLogsResponseBody {
	return s.Body
}

func (s *ListSessionToolLogsResponse) SetHeaders(v map[string]*string) *ListSessionToolLogsResponse {
	s.Headers = v
	return s
}

func (s *ListSessionToolLogsResponse) SetStatusCode(v int32) *ListSessionToolLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSessionToolLogsResponse) SetBody(v *ListSessionToolLogsResponseBody) *ListSessionToolLogsResponse {
	s.Body = v
	return s
}

func (s *ListSessionToolLogsResponse) Validate() error {
	return dara.Validate(s)
}
