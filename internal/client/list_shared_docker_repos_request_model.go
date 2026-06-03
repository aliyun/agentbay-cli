// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"errors"
	"strings"
)

// ListSharedDockerReposRequest is the request struct for ListSharedDockerRepos
type ListSharedDockerReposRequest struct {
	Direction   *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageStart   *int32  `json:"PageStart,omitempty" xml:"PageStart,omitempty"`
	QueryAliUid *int64  `json:"QueryAliUid,omitempty" xml:"QueryAliUid,omitempty"`
}

// Validate validates the ListSharedDockerReposRequest
func (s *ListSharedDockerReposRequest) Validate() error {
	if s.Direction != nil && *s.Direction != "" {
		d := strings.ToLower(*s.Direction)
		if d != "outgoing" && d != "incoming" {
			return errors.New("Direction must be 'Outgoing' or 'Incoming'")
		}
	}
	return nil
}
