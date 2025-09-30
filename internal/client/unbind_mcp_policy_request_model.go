// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindMcpPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v string) *UnbindMcpPolicyRequest
	GetApiKeyId() *string
}

type UnbindMcpPolicyRequest struct {
	ApiKeyId *string `json:"ApiKeyId,omitempty" xml:"ApiKeyId,omitempty"`
}

func (s UnbindMcpPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindMcpPolicyRequest) GoString() string {
	return s.String()
}

func (s *UnbindMcpPolicyRequest) GetApiKeyId() *string {
	return s.ApiKeyId
}

func (s *UnbindMcpPolicyRequest) SetApiKeyId(v string) *UnbindMcpPolicyRequest {
	s.ApiKeyId = &v
	return s
}

func (s *UnbindMcpPolicyRequest) Validate() error {
	return dara.Validate(s)
}
