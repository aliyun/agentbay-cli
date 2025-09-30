// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopPlaygroundTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopPlaygroundTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopPlaygroundTaskResponse
	GetStatusCode() *int32
	SetBody(v *StopPlaygroundTaskResponseBody) *StopPlaygroundTaskResponse
	GetBody() *StopPlaygroundTaskResponseBody
}

type StopPlaygroundTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopPlaygroundTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopPlaygroundTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StopPlaygroundTaskResponse) GoString() string {
	return s.String()
}

func (s *StopPlaygroundTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopPlaygroundTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopPlaygroundTaskResponse) GetBody() *StopPlaygroundTaskResponseBody {
	return s.Body
}

func (s *StopPlaygroundTaskResponse) SetHeaders(v map[string]*string) *StopPlaygroundTaskResponse {
	s.Headers = v
	return s
}

func (s *StopPlaygroundTaskResponse) SetStatusCode(v int32) *StopPlaygroundTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StopPlaygroundTaskResponse) SetBody(v *StopPlaygroundTaskResponseBody) *StopPlaygroundTaskResponse {
	s.Body = v
	return s
}

func (s *StopPlaygroundTaskResponse) Validate() error {
	return dara.Validate(s)
}
