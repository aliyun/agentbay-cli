// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type ListMarketGroupSkillResponseBodyDataItem struct {
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	GroupId   *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s ListMarketGroupSkillResponseBodyDataItem) String() string {
	return dara.Prettify(s)
}

func (s ListMarketGroupSkillResponseBodyDataItem) GoString() string {
	return s.String()
}

// ListMarketGroupSkill response: Data is array of {GroupName, GroupId} per user's JSON sample
type ListMarketGroupSkillResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []ListMarketGroupSkillResponseBodyDataItem `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListMarketGroupSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMarketGroupSkillResponseBody) GoString() string {
	return s.String()
}

func (s *ListMarketGroupSkillResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMarketGroupSkillResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMarketGroupSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMarketGroupSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMarketGroupSkillResponse) GoString() string {
	return s.String()
}

func (s *ListMarketGroupSkillResponse) Validate() error {
	return dara.Validate(s)
}
