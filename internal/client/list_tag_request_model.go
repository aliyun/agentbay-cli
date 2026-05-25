// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ListTagRequest struct {
}

func (s ListTagRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagRequest) GoString() string {
	return s.String()
}

func (s *ListTagRequest) Validate() error {
	return dara.Validate(s)
}
