// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindMcpPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UnbindMcpPolicyResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UnbindMcpPolicyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UnbindMcpPolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *UnbindMcpPolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UnbindMcpPolicyResponseBody
	GetSuccess() *bool
}

type UnbindMcpPolicyResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnbindMcpPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindMcpPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindMcpPolicyResponseBody) GetCode() *string {
	return s.Code
}

func (s *UnbindMcpPolicyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UnbindMcpPolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UnbindMcpPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindMcpPolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UnbindMcpPolicyResponseBody) SetCode(v string) *UnbindMcpPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *UnbindMcpPolicyResponseBody) SetHttpStatusCode(v int32) *UnbindMcpPolicyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UnbindMcpPolicyResponseBody) SetMessage(v string) *UnbindMcpPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *UnbindMcpPolicyResponseBody) SetRequestId(v string) *UnbindMcpPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindMcpPolicyResponseBody) SetSuccess(v bool) *UnbindMcpPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *UnbindMcpPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
