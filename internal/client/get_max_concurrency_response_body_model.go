// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMaxConcurrencyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMaxConcurrencyResponseBody
	GetCode() *string
	SetData(v *GetMaxConcurrencyResponseBodyData) *GetMaxConcurrencyResponseBody
	GetData() *GetMaxConcurrencyResponseBodyData
	SetHttpStatusCode(v int32) *GetMaxConcurrencyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetMaxConcurrencyResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMaxConcurrencyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMaxConcurrencyResponseBody
	GetSuccess() *bool
}

type GetMaxConcurrencyResponseBody struct {
	Code           *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetMaxConcurrencyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMaxConcurrencyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMaxConcurrencyResponseBody) GoString() string {
	return s.String()
}

func (s *GetMaxConcurrencyResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMaxConcurrencyResponseBody) GetData() *GetMaxConcurrencyResponseBodyData {
	return s.Data
}

func (s *GetMaxConcurrencyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMaxConcurrencyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMaxConcurrencyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMaxConcurrencyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMaxConcurrencyResponseBody) SetCode(v string) *GetMaxConcurrencyResponseBody {
	s.Code = &v
	return s
}

func (s *GetMaxConcurrencyResponseBody) SetData(v *GetMaxConcurrencyResponseBodyData) *GetMaxConcurrencyResponseBody {
	s.Data = v
	return s
}

func (s *GetMaxConcurrencyResponseBody) SetHttpStatusCode(v int32) *GetMaxConcurrencyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMaxConcurrencyResponseBody) SetMessage(v string) *GetMaxConcurrencyResponseBody {
	s.Message = &v
	return s
}

func (s *GetMaxConcurrencyResponseBody) SetRequestId(v string) *GetMaxConcurrencyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMaxConcurrencyResponseBody) SetSuccess(v bool) *GetMaxConcurrencyResponseBody {
	s.Success = &v
	return s
}

func (s *GetMaxConcurrencyResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMaxConcurrencyResponseBodyData struct {
	Concurrency *int32 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
}

func (s GetMaxConcurrencyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMaxConcurrencyResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMaxConcurrencyResponseBodyData) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *GetMaxConcurrencyResponseBodyData) SetConcurrency(v int32) *GetMaxConcurrencyResponseBodyData {
	s.Concurrency = &v
	return s
}

func (s *GetMaxConcurrencyResponseBodyData) Validate() error {
	return dara.Validate(s)
}
