// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveMarketGroupSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *RemoveMarketGroupSkillRequest
	GetGroupId() *string
	SetSkillId(v string) *RemoveMarketGroupSkillRequest
	GetSkillId() *string
}

type RemoveMarketGroupSkillRequest struct {
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
}

func (s RemoveMarketGroupSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveMarketGroupSkillRequest) GoString() string {
	return s.String()
}

func (s *RemoveMarketGroupSkillRequest) GetGroupId() *string { return s.GroupId }
func (s *RemoveMarketGroupSkillRequest) SetGroupId(v string) *RemoveMarketGroupSkillRequest {
	s.GroupId = &v
	return s
}
func (s *RemoveMarketGroupSkillRequest) GetSkillId() *string { return s.SkillId }
func (s *RemoveMarketGroupSkillRequest) SetSkillId(v string) *RemoveMarketGroupSkillRequest {
	s.SkillId = &v
	return s
}
func (s *RemoveMarketGroupSkillRequest) Validate() error {
	return dara.Validate(s)
}
