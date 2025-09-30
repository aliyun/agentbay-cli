// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcpPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyIds(v []*string) *DeleteMcpPoliciesRequest
	GetPolicyIds() []*string
}

type DeleteMcpPoliciesRequest struct {
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
}

func (s DeleteMcpPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcpPoliciesRequest) GoString() string {
	return s.String()
}

func (s *DeleteMcpPoliciesRequest) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *DeleteMcpPoliciesRequest) SetPolicyIds(v []*string) *DeleteMcpPoliciesRequest {
	s.PolicyIds = v
	return s
}

func (s *DeleteMcpPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
