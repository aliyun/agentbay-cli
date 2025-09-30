// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMcpPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v string) *UpdateMcpPolicyRequest
	GetApiKeyId() *string
	SetPolicyId(v string) *UpdateMcpPolicyRequest
	GetPolicyId() *string
}

type UpdateMcpPolicyRequest struct {
	ApiKeyId *string `json:"ApiKeyId,omitempty" xml:"ApiKeyId,omitempty"`
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s UpdateMcpPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateMcpPolicyRequest) GetApiKeyId() *string {
	return s.ApiKeyId
}

func (s *UpdateMcpPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *UpdateMcpPolicyRequest) SetApiKeyId(v string) *UpdateMcpPolicyRequest {
	s.ApiKeyId = &v
	return s
}

func (s *UpdateMcpPolicyRequest) SetPolicyId(v string) *UpdateMcpPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *UpdateMcpPolicyRequest) Validate() error {
	return dara.Validate(s)
}
