// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSessionContextResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSessionContextResponseBody
	GetCode() *string
	SetData(v *DescribeSessionContextResponseBodyData) *DescribeSessionContextResponseBody
	GetData() *DescribeSessionContextResponseBodyData
	SetHttpStatusCode(v int32) *DescribeSessionContextResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeSessionContextResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSessionContextResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeSessionContextResponseBody
	GetSuccess() *bool
}

type DescribeSessionContextResponseBody struct {
	Code           *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *DescribeSessionContextResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSessionContextResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSessionContextResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSessionContextResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSessionContextResponseBody) GetData() *DescribeSessionContextResponseBodyData {
	return s.Data
}

func (s *DescribeSessionContextResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeSessionContextResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSessionContextResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSessionContextResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeSessionContextResponseBody) SetCode(v string) *DescribeSessionContextResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSessionContextResponseBody) SetData(v *DescribeSessionContextResponseBodyData) *DescribeSessionContextResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSessionContextResponseBody) SetHttpStatusCode(v int32) *DescribeSessionContextResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeSessionContextResponseBody) SetMessage(v string) *DescribeSessionContextResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSessionContextResponseBody) SetRequestId(v string) *DescribeSessionContextResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSessionContextResponseBody) SetSuccess(v bool) *DescribeSessionContextResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSessionContextResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSessionContextResponseBodyData struct {
	ContextOverUsed *int64  `json:"ContextOverUsed,omitempty" xml:"ContextOverUsed,omitempty"`
	ContextSpec     *int64  `json:"ContextSpec,omitempty" xml:"ContextSpec,omitempty"`
	ContextUsage    *int64  `json:"ContextUsage,omitempty" xml:"ContextUsage,omitempty"`
	UpdateTime      *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeSessionContextResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSessionContextResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSessionContextResponseBodyData) GetContextOverUsed() *int64 {
	return s.ContextOverUsed
}

func (s *DescribeSessionContextResponseBodyData) GetContextSpec() *int64 {
	return s.ContextSpec
}

func (s *DescribeSessionContextResponseBodyData) GetContextUsage() *int64 {
	return s.ContextUsage
}

func (s *DescribeSessionContextResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeSessionContextResponseBodyData) SetContextOverUsed(v int64) *DescribeSessionContextResponseBodyData {
	s.ContextOverUsed = &v
	return s
}

func (s *DescribeSessionContextResponseBodyData) SetContextSpec(v int64) *DescribeSessionContextResponseBodyData {
	s.ContextSpec = &v
	return s
}

func (s *DescribeSessionContextResponseBodyData) SetContextUsage(v int64) *DescribeSessionContextResponseBodyData {
	s.ContextUsage = &v
	return s
}

func (s *DescribeSessionContextResponseBodyData) SetUpdateTime(v string) *DescribeSessionContextResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *DescribeSessionContextResponseBodyData) Validate() error {
	return dara.Validate(s)
}
