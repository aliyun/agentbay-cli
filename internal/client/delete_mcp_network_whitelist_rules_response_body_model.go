// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcpNetworkWhitelistRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteMcpNetworkWhitelistRulesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteMcpNetworkWhitelistRulesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteMcpNetworkWhitelistRulesResponseBody
	GetMessage() *string
	SetPolicyId(v string) *DeleteMcpNetworkWhitelistRulesResponseBody
	GetPolicyId() *string
	SetRequestId(v string) *DeleteMcpNetworkWhitelistRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteMcpNetworkWhitelistRulesResponseBody
	GetSuccess() *bool
}

type DeleteMcpNetworkWhitelistRulesResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	PolicyId       *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMcpNetworkWhitelistRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcpNetworkWhitelistRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMcpNetworkWhitelistRulesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteMcpNetworkWhitelistRulesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteMcpNetworkWhitelistRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteMcpNetworkWhitelistRulesResponseBody) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DeleteMcpNetworkWhitelistRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMcpNetworkWhitelistRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMcpNetworkWhitelistRulesResponseBody) SetCode(v string) *DeleteMcpNetworkWhitelistRulesResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMcpNetworkWhitelistRulesResponseBody) SetHttpStatusCode(v int32) *DeleteMcpNetworkWhitelistRulesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteMcpNetworkWhitelistRulesResponseBody) SetMessage(v string) *DeleteMcpNetworkWhitelistRulesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMcpNetworkWhitelistRulesResponseBody) SetPolicyId(v string) *DeleteMcpNetworkWhitelistRulesResponseBody {
	s.PolicyId = &v
	return s
}

func (s *DeleteMcpNetworkWhitelistRulesResponseBody) SetRequestId(v string) *DeleteMcpNetworkWhitelistRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMcpNetworkWhitelistRulesResponseBody) SetSuccess(v bool) *DeleteMcpNetworkWhitelistRulesResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMcpNetworkWhitelistRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
