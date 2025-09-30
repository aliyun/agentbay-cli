// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageFunctionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetImageFunctionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetImageFunctionResponse
	GetStatusCode() *int32
	SetBody(v *GetImageFunctionResponseBody) *GetImageFunctionResponse
	GetBody() *GetImageFunctionResponseBody
}

type GetImageFunctionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetImageFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetImageFunctionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetImageFunctionResponse) GoString() string {
	return s.String()
}

func (s *GetImageFunctionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetImageFunctionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetImageFunctionResponse) GetBody() *GetImageFunctionResponseBody {
	return s.Body
}

func (s *GetImageFunctionResponse) SetHeaders(v map[string]*string) *GetImageFunctionResponse {
	s.Headers = v
	return s
}

func (s *GetImageFunctionResponse) SetStatusCode(v int32) *GetImageFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetImageFunctionResponse) SetBody(v *GetImageFunctionResponseBody) *GetImageFunctionResponse {
	s.Body = v
	return s
}

func (s *GetImageFunctionResponse) Validate() error {
	return dara.Validate(s)
}
