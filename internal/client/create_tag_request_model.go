// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTagNameList(v []string) *CreateTagRequest
	GetTagNameList() []string
}

type CreateTagRequest struct {
	TagNameList []string `json:"TagNameList,omitempty" xml:"TagNameList,omitempty"`
}

func (s CreateTagRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTagRequest) GoString() string {
	return s.String()
}

func (s *CreateTagRequest) GetTagNameList() []string { return s.TagNameList }
func (s *CreateTagRequest) SetTagNameList(v []string) *CreateTagRequest {
	s.TagNameList = v
	return s
}
func (s *CreateTagRequest) Validate() error {
	return dara.Validate(s)
}
