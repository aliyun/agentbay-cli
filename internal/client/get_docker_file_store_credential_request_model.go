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
	SetFilePath(v string) *GetDockerFileStoreCredentialRequest
	GetFilePath() *string
	SetIsDockerfile(v string) *GetDockerFileStoreCredentialRequest
	GetIsDockerfile() *string
	SetTaskId(v string) *GetDockerFileStoreCredentialRequest
	GetTaskId() *string
}

type GetDockerFileStoreCredentialRequest struct {
	Source       *string `json:"Source,omitempty" xml:"Source,omitempty"`
	FilePath     *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	IsDockerfile *string `json:"IsDockerfile,omitempty" xml:"IsDockerfile,omitempty"`
	TaskId       *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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

func (s *GetDockerFileStoreCredentialRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *GetDockerFileStoreCredentialRequest) SetFilePath(v string) *GetDockerFileStoreCredentialRequest {
	s.FilePath = &v
	return s
}

func (s *GetDockerFileStoreCredentialRequest) GetIsDockerfile() *string {
	return s.IsDockerfile
}

func (s *GetDockerFileStoreCredentialRequest) SetIsDockerfile(v string) *GetDockerFileStoreCredentialRequest {
	s.IsDockerfile = &v
	return s
}

func (s *GetDockerFileStoreCredentialRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetDockerFileStoreCredentialRequest) SetTaskId(v string) *GetDockerFileStoreCredentialRequest {
	s.TaskId = &v
	return s
}

func (s *GetDockerFileStoreCredentialRequest) Validate() error {
	return dara.Validate(s)
}
