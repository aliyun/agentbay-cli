// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"errors"
)

// ShareDockerRepoRequest is the request struct for ShareDockerRepo
type ShareDockerRepoRequest struct {
	TargetAliUid *int64 `json:"TargetAliUid,omitempty" xml:"TargetAliUid,omitempty"`
}

// Validate validates the ShareDockerRepoRequest
func (s *ShareDockerRepoRequest) Validate() error {
	if s.TargetAliUid == nil || *s.TargetAliUid == 0 {
		return errors.New("TargetAliUid is required")
	}
	return nil
}
