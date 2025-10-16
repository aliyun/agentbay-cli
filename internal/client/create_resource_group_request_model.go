// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizRegionId(v string) *CreateResourceGroupRequest
	GetBizRegionId() *string
	SetCpu(v int32) *CreateResourceGroupRequest
	GetCpu() *int32
	SetImageId(v string) *CreateResourceGroupRequest
	GetImageId() *string
	SetMemory(v int32) *CreateResourceGroupRequest
	GetMemory() *int32
	SetOfficeSiteId(v string) *CreateResourceGroupRequest
	GetOfficeSiteId() *string
	SetOfficeSiteType(v string) *CreateResourceGroupRequest
	GetOfficeSiteType() *string
	SetPolicyId(v string) *CreateResourceGroupRequest
	GetPolicyId() *string
	SetRegionId(v string) *CreateResourceGroupRequest
	GetRegionId() *string
	SetSessionBandwidth(v int32) *CreateResourceGroupRequest
	GetSessionBandwidth() *int32
	SetVSwitchId(v string) *CreateResourceGroupRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateResourceGroupRequest
	GetVpcId() *string
}

type CreateResourceGroupRequest struct {
	BizRegionId      *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	Cpu              *int32  `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	ImageId          *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	Memory           *int32  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	OfficeSiteId     *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	OfficeSiteType   *string `json:"OfficeSiteType,omitempty" xml:"OfficeSiteType,omitempty"`
	PolicyId         *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SessionBandwidth *int32  `json:"SessionBandwidth,omitempty" xml:"SessionBandwidth,omitempty"`
	VSwitchId        *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId            *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *CreateResourceGroupRequest) GetCpu() *int32 {
	return s.Cpu
}

func (s *CreateResourceGroupRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CreateResourceGroupRequest) GetMemory() *int32 {
	return s.Memory
}

func (s *CreateResourceGroupRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *CreateResourceGroupRequest) GetOfficeSiteType() *string {
	return s.OfficeSiteType
}

func (s *CreateResourceGroupRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *CreateResourceGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateResourceGroupRequest) GetSessionBandwidth() *int32 {
	return s.SessionBandwidth
}

func (s *CreateResourceGroupRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateResourceGroupRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateResourceGroupRequest) SetBizRegionId(v string) *CreateResourceGroupRequest {
	s.BizRegionId = &v
	return s
}

func (s *CreateResourceGroupRequest) SetCpu(v int32) *CreateResourceGroupRequest {
	s.Cpu = &v
	return s
}

func (s *CreateResourceGroupRequest) SetImageId(v string) *CreateResourceGroupRequest {
	s.ImageId = &v
	return s
}

func (s *CreateResourceGroupRequest) SetMemory(v int32) *CreateResourceGroupRequest {
	s.Memory = &v
	return s
}

func (s *CreateResourceGroupRequest) SetOfficeSiteId(v string) *CreateResourceGroupRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateResourceGroupRequest) SetOfficeSiteType(v string) *CreateResourceGroupRequest {
	s.OfficeSiteType = &v
	return s
}

func (s *CreateResourceGroupRequest) SetPolicyId(v string) *CreateResourceGroupRequest {
	s.PolicyId = &v
	return s
}

func (s *CreateResourceGroupRequest) SetRegionId(v string) *CreateResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateResourceGroupRequest) SetSessionBandwidth(v int32) *CreateResourceGroupRequest {
	s.SessionBandwidth = &v
	return s
}

func (s *CreateResourceGroupRequest) SetVSwitchId(v string) *CreateResourceGroupRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateResourceGroupRequest) SetVpcId(v string) *CreateResourceGroupRequest {
	s.VpcId = &v
	return s
}

func (s *CreateResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
