// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMarketSkillCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *GetMarketSkillCredentialRequest
	GetFileName() *string
}

type GetMarketSkillCredentialRequest struct {
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s GetMarketSkillCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMarketSkillCredentialRequest) GoString() string {
	return s.String()
}

func (s *GetMarketSkillCredentialRequest) GetFileName() *string {
	return s.FileName
}

func (s *GetMarketSkillCredentialRequest) SetFileName(v string) *GetMarketSkillCredentialRequest {
	s.FileName = &v
	return s
}

func (s *GetMarketSkillCredentialRequest) Validate() error {
	return dara.Validate(s)
}
