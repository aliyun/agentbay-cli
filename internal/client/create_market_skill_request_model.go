// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMarketSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOssBucket(v string) *CreateMarketSkillRequest
	GetOssBucket() *string
	SetOssFilePath(v string) *CreateMarketSkillRequest
	GetOssFilePath() *string
	SetTagList(v []string) *CreateMarketSkillRequest
	GetTagList() []string
	SetIcon(v string) *CreateMarketSkillRequest
	GetIcon() *string
}

type CreateMarketSkillRequest struct {
	OssBucket   *string  `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssFilePath *string  `json:"OssFilePath,omitempty" xml:"OssFilePath,omitempty"`
	TagList     []string `json:"TagList,omitempty" xml:"TagList,omitempty"`
	Icon        *string  `json:"Icon,omitempty" xml:"Icon,omitempty"`
}

func (s CreateMarketSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMarketSkillRequest) GoString() string {
	return s.String()
}

func (s *CreateMarketSkillRequest) GetOssBucket() *string { return s.OssBucket }
func (s *CreateMarketSkillRequest) SetOssBucket(v string) *CreateMarketSkillRequest {
	s.OssBucket = &v
	return s
}
func (s *CreateMarketSkillRequest) GetOssFilePath() *string { return s.OssFilePath }
func (s *CreateMarketSkillRequest) SetOssFilePath(v string) *CreateMarketSkillRequest {
	s.OssFilePath = &v
	return s
}
func (s *CreateMarketSkillRequest) GetTagList() []string { return s.TagList }
func (s *CreateMarketSkillRequest) SetTagList(v []string) *CreateMarketSkillRequest {
	s.TagList = v
	return s
}
func (s *CreateMarketSkillRequest) GetIcon() *string { return s.Icon }
func (s *CreateMarketSkillRequest) SetIcon(v string) *CreateMarketSkillRequest {
	s.Icon = &v
	return s
}
func (s *CreateMarketSkillRequest) Validate() error {
	return dara.Validate(s)
}
