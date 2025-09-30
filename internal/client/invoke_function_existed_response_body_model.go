// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeFunctionExistedResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InvokeFunctionExistedResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *InvokeFunctionExistedResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *InvokeFunctionExistedResponseBody
	GetMessage() *string
	SetRequestId(v string) *InvokeFunctionExistedResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InvokeFunctionExistedResponseBody
	GetSuccess() *bool
}

type InvokeFunctionExistedResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InvokeFunctionExistedResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InvokeFunctionExistedResponseBody) GoString() string {
	return s.String()
}

func (s *InvokeFunctionExistedResponseBody) GetCode() *string {
	return s.Code
}

func (s *InvokeFunctionExistedResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *InvokeFunctionExistedResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InvokeFunctionExistedResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InvokeFunctionExistedResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InvokeFunctionExistedResponseBody) SetCode(v string) *InvokeFunctionExistedResponseBody {
	s.Code = &v
	return s
}

func (s *InvokeFunctionExistedResponseBody) SetHttpStatusCode(v int32) *InvokeFunctionExistedResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *InvokeFunctionExistedResponseBody) SetMessage(v string) *InvokeFunctionExistedResponseBody {
	s.Message = &v
	return s
}

func (s *InvokeFunctionExistedResponseBody) SetRequestId(v string) *InvokeFunctionExistedResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvokeFunctionExistedResponseBody) SetSuccess(v bool) *InvokeFunctionExistedResponseBody {
	s.Success = &v
	return s
}

func (s *InvokeFunctionExistedResponseBody) Validate() error {
	return dara.Validate(s)
}
