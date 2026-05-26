// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMarketSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSkillId(v string) *UpdateMarketSkillRequest
	GetSkillId() *string
	SetOssBucket(v string) *UpdateMarketSkillRequest
	GetOssBucket() *string
	SetOssFilePath(v string) *UpdateMarketSkillRequest
	GetOssFilePath() *string
	SetTagList(v []string) *UpdateMarketSkillRequest
	GetTagList() []string
	SetIcon(v string) *UpdateMarketSkillRequest
	GetIcon() *string
}

type UpdateMarketSkillRequest struct {
	SkillId     *string  `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	OssBucket   *string  `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssFilePath *string  `json:"OssFilePath,omitempty" xml:"OssFilePath,omitempty"`
	TagList     []string `json:"TagList,omitempty" xml:"TagList,omitempty"`
	Icon        *string  `json:"Icon,omitempty" xml:"Icon,omitempty"`
}

func (s UpdateMarketSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMarketSkillRequest) GoString() string {
	return s.String()
}

func (s *UpdateMarketSkillRequest) GetSkillId() *string { return s.SkillId }
func (s *UpdateMarketSkillRequest) SetSkillId(v string) *UpdateMarketSkillRequest {
	s.SkillId = &v
	return s
}
func (s *UpdateMarketSkillRequest) GetOssBucket() *string { return s.OssBucket }
func (s *UpdateMarketSkillRequest) SetOssBucket(v string) *UpdateMarketSkillRequest {
	s.OssBucket = &v
	return s
}
func (s *UpdateMarketSkillRequest) GetOssFilePath() *string { return s.OssFilePath }
func (s *UpdateMarketSkillRequest) SetOssFilePath(v string) *UpdateMarketSkillRequest {
	s.OssFilePath = &v
	return s
}
func (s *UpdateMarketSkillRequest) GetTagList() []string { return s.TagList }
func (s *UpdateMarketSkillRequest) SetTagList(v []string) *UpdateMarketSkillRequest {
	s.TagList = v
	return s
}
func (s *UpdateMarketSkillRequest) GetIcon() *string { return s.Icon }
func (s *UpdateMarketSkillRequest) SetIcon(v string) *UpdateMarketSkillRequest {
	s.Icon = &v
	return s
}
func (s *UpdateMarketSkillRequest) Validate() error {
	return dara.Validate(s)
}
