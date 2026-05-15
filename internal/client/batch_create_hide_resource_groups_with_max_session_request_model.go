// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"errors"
)

// BatchCreateHideResourceGroupsWithMaxSessionRequest is the request struct for BatchCreateHideResourceGroupsWithMaxSession
type BatchCreateHideResourceGroupsWithMaxSessionRequest struct {
	ImageId       *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	MaxSessionNum *int32  `json:"MaxSessionNum,omitempty" xml:"MaxSessionNum,omitempty"`
}

// Validate validates the BatchCreateHideResourceGroupsWithMaxSessionRequest
func (s *BatchCreateHideResourceGroupsWithMaxSessionRequest) Validate() error {
	if s.ImageId == nil || *s.ImageId == "" {
		return errors.New("ImageId is required")
	}
	if s.MaxSessionNum == nil {
		return errors.New("MaxSessionNum is required")
	}
	if *s.MaxSessionNum < 1 {
		return errors.New("MaxSessionNum must be greater than or equal to 1")
	}
	return nil
}

// SetImageId sets the ImageId field
func (s *BatchCreateHideResourceGroupsWithMaxSessionRequest) SetImageId(v string) *BatchCreateHideResourceGroupsWithMaxSessionRequest {
	s.ImageId = &v
	return s
}

// SetMaxSessionNum sets the MaxSessionNum field
func (s *BatchCreateHideResourceGroupsWithMaxSessionRequest) SetMaxSessionNum(v int32) *BatchCreateHideResourceGroupsWithMaxSessionRequest {
	s.MaxSessionNum = &v
	return s
}
