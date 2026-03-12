// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"encoding/xml"

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

// ListMarketGroupSkillResponseBodyWrapper unmarshals Body when gateway returns JSON
// with "body" as either an object or a string (XML content).
type ListMarketGroupSkillResponseBodyWrapper struct {
	*ListMarketGroupSkillResponseBody
}

// UnmarshalJSON supports gateway returning body as JSON object or as string (XML).
func (s *ListMarketGroupSkillResponseBodyWrapper) UnmarshalJSON(data []byte) error {
	s.ListMarketGroupSkillResponseBody = &ListMarketGroupSkillResponseBody{}
	if len(data) >= 2 && data[0] == '"' {
		var str string
		if err := json.Unmarshal(data, &str); err != nil {
			return err
		}
		return xml.Unmarshal([]byte(str), s.ListMarketGroupSkillResponseBody)
	}
	return json.Unmarshal(data, s.ListMarketGroupSkillResponseBody)
}

type ListMarketGroupSkillResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMarketGroupSkillResponseBodyWrapper `json:"body,omitempty" xml:"body,omitempty"`
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
