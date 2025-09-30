// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopPlaygroundTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *StopPlaygroundTaskRequest
	GetTaskId() *string
}

type StopPlaygroundTaskRequest struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopPlaygroundTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StopPlaygroundTaskRequest) GoString() string {
	return s.String()
}

func (s *StopPlaygroundTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *StopPlaygroundTaskRequest) SetTaskId(v string) *StopPlaygroundTaskRequest {
	s.TaskId = &v
	return s
}

func (s *StopPlaygroundTaskRequest) Validate() error {
	return dara.Validate(s)
}
