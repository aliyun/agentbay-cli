// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPlaygroundDemoDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPlaygroundDemoDetailResponseBody
	GetCode() *string
	SetData(v *GetPlaygroundDemoDetailResponseBodyData) *GetPlaygroundDemoDetailResponseBody
	GetData() *GetPlaygroundDemoDetailResponseBodyData
	SetHttpStatusCode(v int32) *GetPlaygroundDemoDetailResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetPlaygroundDemoDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPlaygroundDemoDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPlaygroundDemoDetailResponseBody
	GetSuccess() *bool
}

type GetPlaygroundDemoDetailResponseBody struct {
	Code           *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetPlaygroundDemoDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPlaygroundDemoDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPlaygroundDemoDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetPlaygroundDemoDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPlaygroundDemoDetailResponseBody) GetData() *GetPlaygroundDemoDetailResponseBodyData {
	return s.Data
}

func (s *GetPlaygroundDemoDetailResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetPlaygroundDemoDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPlaygroundDemoDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPlaygroundDemoDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPlaygroundDemoDetailResponseBody) SetCode(v string) *GetPlaygroundDemoDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetPlaygroundDemoDetailResponseBody) SetData(v *GetPlaygroundDemoDetailResponseBodyData) *GetPlaygroundDemoDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetPlaygroundDemoDetailResponseBody) SetHttpStatusCode(v int32) *GetPlaygroundDemoDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPlaygroundDemoDetailResponseBody) SetMessage(v string) *GetPlaygroundDemoDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetPlaygroundDemoDetailResponseBody) SetRequestId(v string) *GetPlaygroundDemoDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPlaygroundDemoDetailResponseBody) SetSuccess(v bool) *GetPlaygroundDemoDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetPlaygroundDemoDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetPlaygroundDemoDetailResponseBodyData struct {
	CodeContent *string `json:"CodeContent,omitempty" xml:"CodeContent,omitempty"`
}

func (s GetPlaygroundDemoDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPlaygroundDemoDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPlaygroundDemoDetailResponseBodyData) GetCodeContent() *string {
	return s.CodeContent
}

func (s *GetPlaygroundDemoDetailResponseBodyData) SetCodeContent(v string) *GetPlaygroundDemoDetailResponseBodyData {
	s.CodeContent = &v
	return s
}

func (s *GetPlaygroundDemoDetailResponseBodyData) Validate() error {
	return dara.Validate(s)
}
