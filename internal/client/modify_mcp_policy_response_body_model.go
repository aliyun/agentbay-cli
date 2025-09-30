// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMcpPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyMcpPolicyResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ModifyMcpPolicyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyMcpPolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyMcpPolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyMcpPolicyResponseBody
	GetSuccess() *bool
}

type ModifyMcpPolicyResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyMcpPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyMcpPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMcpPolicyResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyMcpPolicyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyMcpPolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyMcpPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyMcpPolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyMcpPolicyResponseBody) SetCode(v string) *ModifyMcpPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyMcpPolicyResponseBody) SetHttpStatusCode(v int32) *ModifyMcpPolicyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyMcpPolicyResponseBody) SetMessage(v string) *ModifyMcpPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyMcpPolicyResponseBody) SetRequestId(v string) *ModifyMcpPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMcpPolicyResponseBody) SetSuccess(v bool) *ModifyMcpPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyMcpPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
