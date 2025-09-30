// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConcurrencyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateConcurrencyResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateConcurrencyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateConcurrencyResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateConcurrencyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateConcurrencyResponseBody
	GetSuccess() *bool
}

type UpdateConcurrencyResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateConcurrencyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateConcurrencyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConcurrencyResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateConcurrencyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateConcurrencyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateConcurrencyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateConcurrencyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateConcurrencyResponseBody) SetCode(v string) *UpdateConcurrencyResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateConcurrencyResponseBody) SetHttpStatusCode(v int32) *UpdateConcurrencyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateConcurrencyResponseBody) SetMessage(v string) *UpdateConcurrencyResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateConcurrencyResponseBody) SetRequestId(v string) *UpdateConcurrencyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConcurrencyResponseBody) SetSuccess(v bool) *UpdateConcurrencyResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateConcurrencyResponseBody) Validate() error {
	return dara.Validate(s)
}
