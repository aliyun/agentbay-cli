// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDockerImageTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageName(v string) *CreateDockerImageTaskRequest
	GetImageName() *string
	SetSource(v string) *CreateDockerImageTaskRequest
	GetSource() *string
	SetSourceImageId(v string) *CreateDockerImageTaskRequest
	GetSourceImageId() *string
	SetTaskId(v string) *CreateDockerImageTaskRequest
	GetTaskId() *string
}

type CreateDockerImageTaskRequest struct {
	ImageName     *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceImageId *string `json:"SourceImageId,omitempty" xml:"SourceImageId,omitempty"`
	TaskId        *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateDockerImageTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDockerImageTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDockerImageTaskRequest) GetImageName() *string {
	return s.ImageName
}

func (s *CreateDockerImageTaskRequest) GetSource() *string {
	return s.Source
}

func (s *CreateDockerImageTaskRequest) GetSourceImageId() *string {
	return s.SourceImageId
}

func (s *CreateDockerImageTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateDockerImageTaskRequest) SetImageName(v string) *CreateDockerImageTaskRequest {
	s.ImageName = &v
	return s
}

func (s *CreateDockerImageTaskRequest) SetSource(v string) *CreateDockerImageTaskRequest {
	s.Source = &v
	return s
}

func (s *CreateDockerImageTaskRequest) SetSourceImageId(v string) *CreateDockerImageTaskRequest {
	s.SourceImageId = &v
	return s
}

func (s *CreateDockerImageTaskRequest) SetTaskId(v string) *CreateDockerImageTaskRequest {
	s.TaskId = &v
	return s
}

func (s *CreateDockerImageTaskRequest) Validate() error {
	return dara.Validate(s)
}
