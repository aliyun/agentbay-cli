// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeFunctionExistedResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InvokeFunctionExistedResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InvokeFunctionExistedResponse
	GetStatusCode() *int32
	SetBody(v *InvokeFunctionExistedResponseBody) *InvokeFunctionExistedResponse
	GetBody() *InvokeFunctionExistedResponseBody
}

type InvokeFunctionExistedResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InvokeFunctionExistedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvokeFunctionExistedResponse) String() string {
	return dara.Prettify(s)
}

func (s InvokeFunctionExistedResponse) GoString() string {
	return s.String()
}

func (s *InvokeFunctionExistedResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InvokeFunctionExistedResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InvokeFunctionExistedResponse) GetBody() *InvokeFunctionExistedResponseBody {
	return s.Body
}

func (s *InvokeFunctionExistedResponse) SetHeaders(v map[string]*string) *InvokeFunctionExistedResponse {
	s.Headers = v
	return s
}

func (s *InvokeFunctionExistedResponse) SetStatusCode(v int32) *InvokeFunctionExistedResponse {
	s.StatusCode = &v
	return s
}

func (s *InvokeFunctionExistedResponse) SetBody(v *InvokeFunctionExistedResponseBody) *InvokeFunctionExistedResponse {
	s.Body = v
	return s
}

func (s *InvokeFunctionExistedResponse) Validate() error {
	return dara.Validate(s)
}
