// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type RemoveMarketGroupSkillResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveMarketGroupSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveMarketGroupSkillResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveMarketGroupSkillResponseBody) Validate() error {
	return dara.Validate(s)
}

type RemoveMarketGroupSkillResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveMarketGroupSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveMarketGroupSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveMarketGroupSkillResponse) GoString() string {
	return s.String()
}

func (s *RemoveMarketGroupSkillResponse) Validate() error {
	return dara.Validate(s)
}
