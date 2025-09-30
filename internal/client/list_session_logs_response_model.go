// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSessionLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSessionLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSessionLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListSessionLogsResponseBody) *ListSessionLogsResponse
	GetBody() *ListSessionLogsResponseBody
}

type ListSessionLogsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSessionLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSessionLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSessionLogsResponse) GoString() string {
	return s.String()
}

func (s *ListSessionLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSessionLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSessionLogsResponse) GetBody() *ListSessionLogsResponseBody {
	return s.Body
}

func (s *ListSessionLogsResponse) SetHeaders(v map[string]*string) *ListSessionLogsResponse {
	s.Headers = v
	return s
}

func (s *ListSessionLogsResponse) SetStatusCode(v int32) *ListSessionLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSessionLogsResponse) SetBody(v *ListSessionLogsResponseBody) *ListSessionLogsResponse {
	s.Body = v
	return s
}

func (s *ListSessionLogsResponse) Validate() error {
	return dara.Validate(s)
}
