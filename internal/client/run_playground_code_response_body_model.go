// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunPlaygroundCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RunPlaygroundCodeResponseBody
	GetCode() *string
	SetData(v *RunPlaygroundCodeResponseBodyData) *RunPlaygroundCodeResponseBody
	GetData() *RunPlaygroundCodeResponseBodyData
	SetHttpStatusCode(v int32) *RunPlaygroundCodeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RunPlaygroundCodeResponseBody
	GetMessage() *string
	SetRequestId(v string) *RunPlaygroundCodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RunPlaygroundCodeResponseBody
	GetSuccess() *bool
}

type RunPlaygroundCodeResponseBody struct {
	Code           *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *RunPlaygroundCodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RunPlaygroundCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunPlaygroundCodeResponseBody) GoString() string {
	return s.String()
}

func (s *RunPlaygroundCodeResponseBody) GetCode() *string {
	return s.Code
}

func (s *RunPlaygroundCodeResponseBody) GetData() *RunPlaygroundCodeResponseBodyData {
	return s.Data
}

func (s *RunPlaygroundCodeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RunPlaygroundCodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RunPlaygroundCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunPlaygroundCodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RunPlaygroundCodeResponseBody) SetCode(v string) *RunPlaygroundCodeResponseBody {
	s.Code = &v
	return s
}

func (s *RunPlaygroundCodeResponseBody) SetData(v *RunPlaygroundCodeResponseBodyData) *RunPlaygroundCodeResponseBody {
	s.Data = v
	return s
}

func (s *RunPlaygroundCodeResponseBody) SetHttpStatusCode(v int32) *RunPlaygroundCodeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RunPlaygroundCodeResponseBody) SetMessage(v string) *RunPlaygroundCodeResponseBody {
	s.Message = &v
	return s
}

func (s *RunPlaygroundCodeResponseBody) SetRequestId(v string) *RunPlaygroundCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunPlaygroundCodeResponseBody) SetSuccess(v bool) *RunPlaygroundCodeResponseBody {
	s.Success = &v
	return s
}

func (s *RunPlaygroundCodeResponseBody) Validate() error {
	return dara.Validate(s)
}

type RunPlaygroundCodeResponseBodyData struct {
	ErrMsg      *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	ResourceId  *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceUrl *string `json:"ResourceUrl,omitempty" xml:"ResourceUrl,omitempty"`
	SessionId   *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RunPlaygroundCodeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RunPlaygroundCodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *RunPlaygroundCodeResponseBodyData) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *RunPlaygroundCodeResponseBodyData) GetResourceId() *string {
	return s.ResourceId
}

func (s *RunPlaygroundCodeResponseBodyData) GetResourceUrl() *string {
	return s.ResourceUrl
}

func (s *RunPlaygroundCodeResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *RunPlaygroundCodeResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *RunPlaygroundCodeResponseBodyData) SetErrMsg(v string) *RunPlaygroundCodeResponseBodyData {
	s.ErrMsg = &v
	return s
}

func (s *RunPlaygroundCodeResponseBodyData) SetResourceId(v string) *RunPlaygroundCodeResponseBodyData {
	s.ResourceId = &v
	return s
}

func (s *RunPlaygroundCodeResponseBodyData) SetResourceUrl(v string) *RunPlaygroundCodeResponseBodyData {
	s.ResourceUrl = &v
	return s
}

func (s *RunPlaygroundCodeResponseBodyData) SetSessionId(v string) *RunPlaygroundCodeResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *RunPlaygroundCodeResponseBodyData) SetTaskId(v string) *RunPlaygroundCodeResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *RunPlaygroundCodeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
