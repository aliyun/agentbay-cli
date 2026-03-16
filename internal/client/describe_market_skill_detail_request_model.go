// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMarketSkillDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSkillId(v string) *DescribeMarketSkillDetailRequest
	GetSkillId() *string
}

type DescribeMarketSkillDetailRequest struct {
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
}

func (s DescribeMarketSkillDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMarketSkillDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeMarketSkillDetailRequest) GetSkillId() *string { return s.SkillId }
func (s *DescribeMarketSkillDetailRequest) SetSkillId(v string) *DescribeMarketSkillDetailRequest {
	s.SkillId = &v
	return s
}
func (s *DescribeMarketSkillDetailRequest) Validate() error {
	return dara.Validate(s)
}
