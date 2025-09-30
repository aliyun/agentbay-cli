// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceGroupSessionBandwidthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceGroupId(v string) *UpdateResourceGroupSessionBandwidthRequest
	GetResourceGroupId() *string
	SetSessionBandwidth(v int32) *UpdateResourceGroupSessionBandwidthRequest
	GetSessionBandwidth() *int32
}

type UpdateResourceGroupSessionBandwidthRequest struct {
	ResourceGroupId  *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SessionBandwidth *int32  `json:"SessionBandwidth,omitempty" xml:"SessionBandwidth,omitempty"`
}

func (s UpdateResourceGroupSessionBandwidthRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceGroupSessionBandwidthRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceGroupSessionBandwidthRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateResourceGroupSessionBandwidthRequest) GetSessionBandwidth() *int32 {
	return s.SessionBandwidth
}

func (s *UpdateResourceGroupSessionBandwidthRequest) SetResourceGroupId(v string) *UpdateResourceGroupSessionBandwidthRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateResourceGroupSessionBandwidthRequest) SetSessionBandwidth(v int32) *UpdateResourceGroupSessionBandwidthRequest {
	s.SessionBandwidth = &v
	return s
}

func (s *UpdateResourceGroupSessionBandwidthRequest) Validate() error {
	return dara.Validate(s)
}
