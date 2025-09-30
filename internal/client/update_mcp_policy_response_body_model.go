// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMcpPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateMcpPolicyResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateMcpPolicyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateMcpPolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateMcpPolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateMcpPolicyResponseBody
	GetSuccess() *bool
}

type UpdateMcpPolicyResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMcpPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMcpPolicyResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateMcpPolicyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateMcpPolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateMcpPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMcpPolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMcpPolicyResponseBody) SetCode(v string) *UpdateMcpPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateMcpPolicyResponseBody) SetHttpStatusCode(v int32) *UpdateMcpPolicyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateMcpPolicyResponseBody) SetMessage(v string) *UpdateMcpPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateMcpPolicyResponseBody) SetRequestId(v string) *UpdateMcpPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMcpPolicyResponseBody) SetSuccess(v bool) *UpdateMcpPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMcpPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
