// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"errors"
)

// UnshareDockerRepoRequest is the request struct for UnshareDockerRepo
type UnshareDockerRepoRequest struct {
	TargetAliUid *int64 `json:"TargetAliUid,omitempty" xml:"TargetAliUid,omitempty"`
}

// Validate validates the UnshareDockerRepoRequest
func (s *UnshareDockerRepoRequest) Validate() error {
	if s.TargetAliUid == nil || *s.TargetAliUid == 0 {
		return errors.New("TargetAliUid is required")
	}
	return nil
}
