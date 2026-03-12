// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type AddMarketGroupSkillResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddMarketGroupSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddMarketGroupSkillResponseBody) GoString() string {
	return s.String()
}

func (s *AddMarketGroupSkillResponseBody) Validate() error {
	return dara.Validate(s)
}

type AddMarketGroupSkillResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddMarketGroupSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddMarketGroupSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s AddMarketGroupSkillResponse) GoString() string {
	return s.String()
}

func (s *AddMarketGroupSkillResponse) Validate() error {
	return dara.Validate(s)
}
