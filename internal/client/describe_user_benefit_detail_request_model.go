// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserBenefitDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceGroupId(v string) *DescribeUserBenefitDetailRequest
	GetResourceGroupId() *string
	SetSessionBandwidth(v int32) *DescribeUserBenefitDetailRequest
	GetSessionBandwidth() *int32
}

type DescribeUserBenefitDetailRequest struct {
	ResourceGroupId  *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SessionBandwidth *int32  `json:"SessionBandwidth,omitempty" xml:"SessionBandwidth,omitempty"`
}

func (s DescribeUserBenefitDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserBenefitDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserBenefitDetailRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeUserBenefitDetailRequest) GetSessionBandwidth() *int32 {
	return s.SessionBandwidth
}

func (s *DescribeUserBenefitDetailRequest) SetResourceGroupId(v string) *DescribeUserBenefitDetailRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeUserBenefitDetailRequest) SetSessionBandwidth(v int32) *DescribeUserBenefitDetailRequest {
	s.SessionBandwidth = &v
	return s
}

func (s *DescribeUserBenefitDetailRequest) Validate() error {
	return dara.Validate(s)
}
