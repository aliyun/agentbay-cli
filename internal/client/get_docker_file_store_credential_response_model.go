// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDockerFileStoreCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDockerFileStoreCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDockerFileStoreCredentialResponse
	GetStatusCode() *int32
	SetBody(v *GetDockerFileStoreCredentialResponseBody) *GetDockerFileStoreCredentialResponse
	GetBody() *GetDockerFileStoreCredentialResponseBody
}

type GetDockerFileStoreCredentialResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDockerFileStoreCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDockerFileStoreCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDockerFileStoreCredentialResponse) GoString() string {
	return s.String()
}

func (s *GetDockerFileStoreCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDockerFileStoreCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDockerFileStoreCredentialResponse) GetBody() *GetDockerFileStoreCredentialResponseBody {
	return s.Body
}

func (s *GetDockerFileStoreCredentialResponse) SetHeaders(v map[string]*string) *GetDockerFileStoreCredentialResponse {
	s.Headers = v
	return s
}

func (s *GetDockerFileStoreCredentialResponse) SetStatusCode(v int32) *GetDockerFileStoreCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDockerFileStoreCredentialResponse) SetBody(v *GetDockerFileStoreCredentialResponseBody) *GetDockerFileStoreCredentialResponse {
	s.Body = v
	return s
}

func (s *GetDockerFileStoreCredentialResponse) Validate() error {
	return dara.Validate(s)
}
