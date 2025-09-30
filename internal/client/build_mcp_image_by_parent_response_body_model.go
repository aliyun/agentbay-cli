// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBuildMcpImageByParentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BuildMcpImageByParentResponseBody
	GetCode() *string
	SetData(v *BuildMcpImageByParentResponseBodyData) *BuildMcpImageByParentResponseBody
	GetData() *BuildMcpImageByParentResponseBodyData
	SetHttpStatusCode(v int32) *BuildMcpImageByParentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *BuildMcpImageByParentResponseBody
	GetMessage() *string
	SetRequestId(v string) *BuildMcpImageByParentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BuildMcpImageByParentResponseBody
	GetSuccess() *bool
}

type BuildMcpImageByParentResponseBody struct {
	Code           *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *BuildMcpImageByParentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BuildMcpImageByParentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BuildMcpImageByParentResponseBody) GoString() string {
	return s.String()
}

func (s *BuildMcpImageByParentResponseBody) GetCode() *string {
	return s.Code
}

func (s *BuildMcpImageByParentResponseBody) GetData() *BuildMcpImageByParentResponseBodyData {
	return s.Data
}

func (s *BuildMcpImageByParentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *BuildMcpImageByParentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BuildMcpImageByParentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BuildMcpImageByParentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BuildMcpImageByParentResponseBody) SetCode(v string) *BuildMcpImageByParentResponseBody {
	s.Code = &v
	return s
}

func (s *BuildMcpImageByParentResponseBody) SetData(v *BuildMcpImageByParentResponseBodyData) *BuildMcpImageByParentResponseBody {
	s.Data = v
	return s
}

func (s *BuildMcpImageByParentResponseBody) SetHttpStatusCode(v int32) *BuildMcpImageByParentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BuildMcpImageByParentResponseBody) SetMessage(v string) *BuildMcpImageByParentResponseBody {
	s.Message = &v
	return s
}

func (s *BuildMcpImageByParentResponseBody) SetRequestId(v string) *BuildMcpImageByParentResponseBody {
	s.RequestId = &v
	return s
}

func (s *BuildMcpImageByParentResponseBody) SetSuccess(v bool) *BuildMcpImageByParentResponseBody {
	s.Success = &v
	return s
}

func (s *BuildMcpImageByParentResponseBody) Validate() error {
	return dara.Validate(s)
}

type BuildMcpImageByParentResponseBodyData struct {
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s BuildMcpImageByParentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BuildMcpImageByParentResponseBodyData) GoString() string {
	return s.String()
}

func (s *BuildMcpImageByParentResponseBodyData) GetImageId() *string {
	return s.ImageId
}

func (s *BuildMcpImageByParentResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *BuildMcpImageByParentResponseBodyData) GetVersionId() *string {
	return s.VersionId
}

func (s *BuildMcpImageByParentResponseBodyData) SetImageId(v string) *BuildMcpImageByParentResponseBodyData {
	s.ImageId = &v
	return s
}

func (s *BuildMcpImageByParentResponseBodyData) SetTaskId(v string) *BuildMcpImageByParentResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *BuildMcpImageByParentResponseBodyData) SetVersionId(v string) *BuildMcpImageByParentResponseBodyData {
	s.VersionId = &v
	return s
}

func (s *BuildMcpImageByParentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
