// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDockerfileTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDockerfileTemplateResponseBody
	GetCode() *string
	SetData(v *GetDockerfileTemplateResponseBodyData) *GetDockerfileTemplateResponseBody
	GetData() *GetDockerfileTemplateResponseBodyData
	SetHttpStatusCode(v int32) *GetDockerfileTemplateResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDockerfileTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDockerfileTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDockerfileTemplateResponseBody
	GetSuccess() *bool
}

type GetDockerfileTemplateResponseBody struct {
	Code           *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetDockerfileTemplateResponseBodyData   `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDockerfileTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDockerfileTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetDockerfileTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDockerfileTemplateResponseBody) GetData() *GetDockerfileTemplateResponseBodyData {
	return s.Data
}

func (s *GetDockerfileTemplateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDockerfileTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDockerfileTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDockerfileTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDockerfileTemplateResponseBody) SetCode(v string) *GetDockerfileTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *GetDockerfileTemplateResponseBody) SetData(v *GetDockerfileTemplateResponseBodyData) *GetDockerfileTemplateResponseBody {
	s.Data = v
	return s
}

func (s *GetDockerfileTemplateResponseBody) SetHttpStatusCode(v int32) *GetDockerfileTemplateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDockerfileTemplateResponseBody) SetMessage(v string) *GetDockerfileTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *GetDockerfileTemplateResponseBody) SetRequestId(v string) *GetDockerfileTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDockerfileTemplateResponseBody) SetSuccess(v bool) *GetDockerfileTemplateResponseBody {
	s.Success = &v
	return s
}

type GetDockerfileTemplateResponseBodyData struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s *GetDockerfileTemplateResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *GetDockerfileTemplateResponseBodyData) SetContent(v string) *GetDockerfileTemplateResponseBodyData {
	s.Content = &v
	return s
}

