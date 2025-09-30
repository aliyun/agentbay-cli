// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDockerFileStoreCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDockerFileStoreCredentialResponseBody
	GetCode() *string
	SetData(v *GetDockerFileStoreCredentialResponseBodyData) *GetDockerFileStoreCredentialResponseBody
	GetData() *GetDockerFileStoreCredentialResponseBodyData
	SetHttpStatusCode(v int32) *GetDockerFileStoreCredentialResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDockerFileStoreCredentialResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDockerFileStoreCredentialResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDockerFileStoreCredentialResponseBody
	GetSuccess() *bool
}

type GetDockerFileStoreCredentialResponseBody struct {
	Code           *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetDockerFileStoreCredentialResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDockerFileStoreCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDockerFileStoreCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *GetDockerFileStoreCredentialResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDockerFileStoreCredentialResponseBody) GetData() *GetDockerFileStoreCredentialResponseBodyData {
	return s.Data
}

func (s *GetDockerFileStoreCredentialResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDockerFileStoreCredentialResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDockerFileStoreCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDockerFileStoreCredentialResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDockerFileStoreCredentialResponseBody) SetCode(v string) *GetDockerFileStoreCredentialResponseBody {
	s.Code = &v
	return s
}

func (s *GetDockerFileStoreCredentialResponseBody) SetData(v *GetDockerFileStoreCredentialResponseBodyData) *GetDockerFileStoreCredentialResponseBody {
	s.Data = v
	return s
}

func (s *GetDockerFileStoreCredentialResponseBody) SetHttpStatusCode(v int32) *GetDockerFileStoreCredentialResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDockerFileStoreCredentialResponseBody) SetMessage(v string) *GetDockerFileStoreCredentialResponseBody {
	s.Message = &v
	return s
}

func (s *GetDockerFileStoreCredentialResponseBody) SetRequestId(v string) *GetDockerFileStoreCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDockerFileStoreCredentialResponseBody) SetSuccess(v bool) *GetDockerFileStoreCredentialResponseBody {
	s.Success = &v
	return s
}

func (s *GetDockerFileStoreCredentialResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDockerFileStoreCredentialResponseBodyData struct {
	OssUrl *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetDockerFileStoreCredentialResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDockerFileStoreCredentialResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDockerFileStoreCredentialResponseBodyData) GetOssUrl() *string {
	return s.OssUrl
}

func (s *GetDockerFileStoreCredentialResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetDockerFileStoreCredentialResponseBodyData) SetOssUrl(v string) *GetDockerFileStoreCredentialResponseBodyData {
	s.OssUrl = &v
	return s
}

func (s *GetDockerFileStoreCredentialResponseBodyData) SetTaskId(v string) *GetDockerFileStoreCredentialResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetDockerFileStoreCredentialResponseBodyData) Validate() error {
	return dara.Validate(s)
}
