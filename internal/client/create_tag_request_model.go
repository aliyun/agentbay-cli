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
	SetTagList(v []string) *CreateTagRequest
	GetTagList() []string
}

type CreateTagRequest struct {
	TagList []string `json:"TagList,omitempty" xml:"TagList,omitempty"`
}

func (s CreateTagRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTagRequest) GoString() string {
	return s.String()
}

func (s *CreateTagRequest) GetTagList() []string { return s.TagList }
func (s *CreateTagRequest) SetTagList(v []string) *CreateTagRequest {
	s.TagList = v
	return s
}
func (s *CreateTagRequest) Validate() error {
	return dara.Validate(s)
}
