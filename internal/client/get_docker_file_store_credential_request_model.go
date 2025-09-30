// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDockerFileStoreCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSource(v string) *GetDockerFileStoreCredentialRequest
	GetSource() *string
}

type GetDockerFileStoreCredentialRequest struct {
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s GetDockerFileStoreCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDockerFileStoreCredentialRequest) GoString() string {
	return s.String()
}

func (s *GetDockerFileStoreCredentialRequest) GetSource() *string {
	return s.Source
}

func (s *GetDockerFileStoreCredentialRequest) SetSource(v string) *GetDockerFileStoreCredentialRequest {
	s.Source = &v
	return s
}

func (s *GetDockerFileStoreCredentialRequest) Validate() error {
	return dara.Validate(s)
}
