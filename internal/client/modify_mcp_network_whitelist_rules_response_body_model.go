// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMcpNetworkWhitelistRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyMcpNetworkWhitelistRulesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ModifyMcpNetworkWhitelistRulesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyMcpNetworkWhitelistRulesResponseBody
	GetMessage() *string
	SetPolicyId(v string) *ModifyMcpNetworkWhitelistRulesResponseBody
	GetPolicyId() *string
	SetRequestId(v string) *ModifyMcpNetworkWhitelistRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyMcpNetworkWhitelistRulesResponseBody
	GetSuccess() *bool
}

type ModifyMcpNetworkWhitelistRulesResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	PolicyId       *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyMcpNetworkWhitelistRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyMcpNetworkWhitelistRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMcpNetworkWhitelistRulesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyMcpNetworkWhitelistRulesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyMcpNetworkWhitelistRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyMcpNetworkWhitelistRulesResponseBody) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ModifyMcpNetworkWhitelistRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyMcpNetworkWhitelistRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyMcpNetworkWhitelistRulesResponseBody) SetCode(v string) *ModifyMcpNetworkWhitelistRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyMcpNetworkWhitelistRulesResponseBody) SetHttpStatusCode(v int32) *ModifyMcpNetworkWhitelistRulesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyMcpNetworkWhitelistRulesResponseBody) SetMessage(v string) *ModifyMcpNetworkWhitelistRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyMcpNetworkWhitelistRulesResponseBody) SetPolicyId(v string) *ModifyMcpNetworkWhitelistRulesResponseBody {
	s.PolicyId = &v
	return s
}

func (s *ModifyMcpNetworkWhitelistRulesResponseBody) SetRequestId(v string) *ModifyMcpNetworkWhitelistRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMcpNetworkWhitelistRulesResponseBody) SetSuccess(v bool) *ModifyMcpNetworkWhitelistRulesResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyMcpNetworkWhitelistRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
