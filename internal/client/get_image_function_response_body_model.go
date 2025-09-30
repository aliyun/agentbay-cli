// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageFunctionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetImageFunctionResponseBody
	GetCode() *string
	SetData(v []*GetImageFunctionResponseBodyData) *GetImageFunctionResponseBody
	GetData() []*GetImageFunctionResponseBodyData
	SetHttpStatusCode(v int32) *GetImageFunctionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetImageFunctionResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetImageFunctionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetImageFunctionResponseBody
	GetSuccess() *bool
}

type GetImageFunctionResponseBody struct {
	Code           *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*GetImageFunctionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetImageFunctionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetImageFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageFunctionResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetImageFunctionResponseBody) GetData() []*GetImageFunctionResponseBodyData {
	return s.Data
}

func (s *GetImageFunctionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetImageFunctionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetImageFunctionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetImageFunctionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetImageFunctionResponseBody) SetCode(v string) *GetImageFunctionResponseBody {
	s.Code = &v
	return s
}

func (s *GetImageFunctionResponseBody) SetData(v []*GetImageFunctionResponseBodyData) *GetImageFunctionResponseBody {
	s.Data = v
	return s
}

func (s *GetImageFunctionResponseBody) SetHttpStatusCode(v int32) *GetImageFunctionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetImageFunctionResponseBody) SetMessage(v string) *GetImageFunctionResponseBody {
	s.Message = &v
	return s
}

func (s *GetImageFunctionResponseBody) SetRequestId(v string) *GetImageFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageFunctionResponseBody) SetSuccess(v bool) *GetImageFunctionResponseBody {
	s.Success = &v
	return s
}

func (s *GetImageFunctionResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetImageFunctionResponseBodyData struct {
	ChoiceType   *string `json:"ChoiceType,omitempty" xml:"ChoiceType,omitempty"`
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FunctionId   *string `json:"FunctionId,omitempty" xml:"FunctionId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetImageFunctionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetImageFunctionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetImageFunctionResponseBodyData) GetChoiceType() *string {
	return s.ChoiceType
}

func (s *GetImageFunctionResponseBodyData) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *GetImageFunctionResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetImageFunctionResponseBodyData) GetFunctionId() *string {
	return s.FunctionId
}

func (s *GetImageFunctionResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetImageFunctionResponseBodyData) SetChoiceType(v string) *GetImageFunctionResponseBodyData {
	s.ChoiceType = &v
	return s
}

func (s *GetImageFunctionResponseBodyData) SetDefaultValue(v string) *GetImageFunctionResponseBodyData {
	s.DefaultValue = &v
	return s
}

func (s *GetImageFunctionResponseBodyData) SetDescription(v string) *GetImageFunctionResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetImageFunctionResponseBodyData) SetFunctionId(v string) *GetImageFunctionResponseBodyData {
	s.FunctionId = &v
	return s
}

func (s *GetImageFunctionResponseBodyData) SetName(v string) *GetImageFunctionResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetImageFunctionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
