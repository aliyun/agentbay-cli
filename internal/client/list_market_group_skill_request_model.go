// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMarketGroupSkillRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ListMarketGroupSkillRequest struct {
}

func (s ListMarketGroupSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMarketGroupSkillRequest) GoString() string {
	return s.String()
}

func (s *ListMarketGroupSkillRequest) Validate() error {
	return dara.Validate(s)
}
