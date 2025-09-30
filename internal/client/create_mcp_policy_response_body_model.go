// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcpPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateMcpPolicyResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateMcpPolicyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateMcpPolicyResponseBody
	GetMessage() *string
	SetPolicyId(v string) *CreateMcpPolicyResponseBody
	GetPolicyId() *string
	SetRequestId(v string) *CreateMcpPolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateMcpPolicyResponseBody
	GetSuccess() *bool
}

type CreateMcpPolicyResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	PolicyId       *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMcpPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMcpPolicyResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateMcpPolicyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateMcpPolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateMcpPolicyResponseBody) GetPolicyId() *string {
	return s.PolicyId
}

func (s *CreateMcpPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMcpPolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMcpPolicyResponseBody) SetCode(v string) *CreateMcpPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMcpPolicyResponseBody) SetHttpStatusCode(v int32) *CreateMcpPolicyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateMcpPolicyResponseBody) SetMessage(v string) *CreateMcpPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMcpPolicyResponseBody) SetPolicyId(v string) *CreateMcpPolicyResponseBody {
	s.PolicyId = &v
	return s
}

func (s *CreateMcpPolicyResponseBody) SetRequestId(v string) *CreateMcpPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMcpPolicyResponseBody) SetSuccess(v bool) *CreateMcpPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMcpPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
