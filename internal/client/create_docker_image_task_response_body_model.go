// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDockerImageTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateDockerImageTaskResponseBody
	GetCode() *string
	SetData(v *CreateDockerImageTaskResponseBodyData) *CreateDockerImageTaskResponseBody
	GetData() *CreateDockerImageTaskResponseBodyData
	SetHttpStatusCode(v int32) *CreateDockerImageTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateDockerImageTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateDockerImageTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDockerImageTaskResponseBody
	GetSuccess() *bool
}

type CreateDockerImageTaskResponseBody struct {
	Code           *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *CreateDockerImageTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDockerImageTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDockerImageTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDockerImageTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateDockerImageTaskResponseBody) GetData() *CreateDockerImageTaskResponseBodyData {
	return s.Data
}

func (s *CreateDockerImageTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateDockerImageTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDockerImageTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDockerImageTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDockerImageTaskResponseBody) SetCode(v string) *CreateDockerImageTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDockerImageTaskResponseBody) SetData(v *CreateDockerImageTaskResponseBodyData) *CreateDockerImageTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateDockerImageTaskResponseBody) SetHttpStatusCode(v int32) *CreateDockerImageTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateDockerImageTaskResponseBody) SetMessage(v string) *CreateDockerImageTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDockerImageTaskResponseBody) SetRequestId(v string) *CreateDockerImageTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDockerImageTaskResponseBody) SetSuccess(v bool) *CreateDockerImageTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDockerImageTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateDockerImageTaskResponseBodyData struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateDockerImageTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateDockerImageTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDockerImageTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateDockerImageTaskResponseBodyData) SetTaskId(v string) *CreateDockerImageTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateDockerImageTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
