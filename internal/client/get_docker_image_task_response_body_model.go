// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDockerImageTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDockerImageTaskResponseBody
	GetCode() *string
	SetData(v *GetDockerImageTaskResponseBodyData) *GetDockerImageTaskResponseBody
	GetData() *GetDockerImageTaskResponseBodyData
	SetHttpStatusCode(v int32) *GetDockerImageTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDockerImageTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDockerImageTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDockerImageTaskResponseBody
	GetSuccess() *bool
}

type GetDockerImageTaskResponseBody struct {
	Code           *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetDockerImageTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDockerImageTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDockerImageTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetDockerImageTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDockerImageTaskResponseBody) GetData() *GetDockerImageTaskResponseBodyData {
	return s.Data
}

func (s *GetDockerImageTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDockerImageTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDockerImageTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDockerImageTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDockerImageTaskResponseBody) SetCode(v string) *GetDockerImageTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetDockerImageTaskResponseBody) SetData(v *GetDockerImageTaskResponseBodyData) *GetDockerImageTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetDockerImageTaskResponseBody) SetHttpStatusCode(v int32) *GetDockerImageTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDockerImageTaskResponseBody) SetMessage(v string) *GetDockerImageTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetDockerImageTaskResponseBody) SetRequestId(v string) *GetDockerImageTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDockerImageTaskResponseBody) SetSuccess(v bool) *GetDockerImageTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetDockerImageTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDockerImageTaskResponseBodyData struct {
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskMsg *string `json:"TaskMsg,omitempty" xml:"TaskMsg,omitempty"`
}

func (s GetDockerImageTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDockerImageTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDockerImageTaskResponseBodyData) GetImageId() *string {
	return s.ImageId
}

func (s *GetDockerImageTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetDockerImageTaskResponseBodyData) GetTaskMsg() *string {
	return s.TaskMsg
}

func (s *GetDockerImageTaskResponseBodyData) SetImageId(v string) *GetDockerImageTaskResponseBodyData {
	s.ImageId = &v
	return s
}

func (s *GetDockerImageTaskResponseBodyData) SetStatus(v string) *GetDockerImageTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetDockerImageTaskResponseBodyData) SetTaskMsg(v string) *GetDockerImageTaskResponseBodyData {
	s.TaskMsg = &v
	return s
}

func (s *GetDockerImageTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
