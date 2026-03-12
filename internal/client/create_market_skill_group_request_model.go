// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMarketSkillGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *CreateMarketSkillGroupRequest
	GetGroupName() *string
}

type CreateMarketSkillGroupRequest struct {
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s CreateMarketSkillGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMarketSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateMarketSkillGroupRequest) GetGroupName() *string { return s.GroupName }
func (s *CreateMarketSkillGroupRequest) SetGroupName(v string) *CreateMarketSkillGroupRequest {
	s.GroupName = &v
	return s
}
func (s *CreateMarketSkillGroupRequest) Validate() error {
	return dara.Validate(s)
}
