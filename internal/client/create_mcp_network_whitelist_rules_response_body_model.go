// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcpNetworkWhitelistRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateMcpNetworkWhitelistRulesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateMcpNetworkWhitelistRulesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateMcpNetworkWhitelistRulesResponseBody
	GetMessage() *string
	SetPolicyId(v string) *CreateMcpNetworkWhitelistRulesResponseBody
	GetPolicyId() *string
	SetRequestId(v string) *CreateMcpNetworkWhitelistRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateMcpNetworkWhitelistRulesResponseBody
	GetSuccess() *bool
}

type CreateMcpNetworkWhitelistRulesResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	PolicyId       *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMcpNetworkWhitelistRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpNetworkWhitelistRulesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMcpNetworkWhitelistRulesResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateMcpNetworkWhitelistRulesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateMcpNetworkWhitelistRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateMcpNetworkWhitelistRulesResponseBody) GetPolicyId() *string {
	return s.PolicyId
}

func (s *CreateMcpNetworkWhitelistRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMcpNetworkWhitelistRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMcpNetworkWhitelistRulesResponseBody) SetCode(v string) *CreateMcpNetworkWhitelistRulesResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMcpNetworkWhitelistRulesResponseBody) SetHttpStatusCode(v int32) *CreateMcpNetworkWhitelistRulesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateMcpNetworkWhitelistRulesResponseBody) SetMessage(v string) *CreateMcpNetworkWhitelistRulesResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMcpNetworkWhitelistRulesResponseBody) SetPolicyId(v string) *CreateMcpNetworkWhitelistRulesResponseBody {
	s.PolicyId = &v
	return s
}

func (s *CreateMcpNetworkWhitelistRulesResponseBody) SetRequestId(v string) *CreateMcpNetworkWhitelistRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMcpNetworkWhitelistRulesResponseBody) SetSuccess(v bool) *CreateMcpNetworkWhitelistRulesResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMcpNetworkWhitelistRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
