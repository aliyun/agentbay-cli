// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDockerImageTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSource(v string) *GetDockerImageTaskRequest
	GetSource() *string
	SetTaskId(v string) *GetDockerImageTaskRequest
	GetTaskId() *string
}

type GetDockerImageTaskRequest struct {
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetDockerImageTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDockerImageTaskRequest) GoString() string {
	return s.String()
}

func (s *GetDockerImageTaskRequest) GetSource() *string {
	return s.Source
}

func (s *GetDockerImageTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetDockerImageTaskRequest) SetSource(v string) *GetDockerImageTaskRequest {
	s.Source = &v
	return s
}

func (s *GetDockerImageTaskRequest) SetTaskId(v string) *GetDockerImageTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetDockerImageTaskRequest) Validate() error {
	return dara.Validate(s)
}
