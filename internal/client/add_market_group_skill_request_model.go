// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMarketGroupSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *AddMarketGroupSkillRequest
	GetGroupId() *string
	SetSkillId(v string) *AddMarketGroupSkillRequest
	GetSkillId() *string
}

type AddMarketGroupSkillRequest struct {
	GroupId  *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SkillId  *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
}

func (s AddMarketGroupSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s AddMarketGroupSkillRequest) GoString() string {
	return s.String()
}

func (s *AddMarketGroupSkillRequest) GetGroupId() *string  { return s.GroupId }
func (s *AddMarketGroupSkillRequest) SetGroupId(v string) *AddMarketGroupSkillRequest {
	s.GroupId = &v
	return s
}
func (s *AddMarketGroupSkillRequest) GetSkillId() *string  { return s.SkillId }
func (s *AddMarketGroupSkillRequest) SetSkillId(v string) *AddMarketGroupSkillRequest {
	s.SkillId = &v
	return s
}
func (s *AddMarketGroupSkillRequest) Validate() error {
	return dara.Validate(s)
}
