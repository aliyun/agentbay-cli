// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"errors"
)

// UpdateImageReserveMinAmountRequest is the request struct for UpdateImageReserveMinAmount
type UpdateImageReserveMinAmountRequest struct {
	ImageId          *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ReserveMinAmount *int32  `json:"ReserveMinAmount,omitempty" xml:"ReserveMinAmount,omitempty"`
}

// Validate validates the UpdateImageReserveMinAmountRequest
func (s *UpdateImageReserveMinAmountRequest) Validate() error {
	if s.ImageId == nil || *s.ImageId == "" {
		return errors.New("ImageId is required")
	}
	if s.ReserveMinAmount == nil {
		return errors.New("ReserveMinAmount is required")
	}
	if *s.ReserveMinAmount < 0 {
		return errors.New("ReserveMinAmount must be greater than or equal to 0")
	}
	return nil
}

// SetImageId sets the ImageId field
func (s *UpdateImageReserveMinAmountRequest) SetImageId(v string) *UpdateImageReserveMinAmountRequest {
	s.ImageId = &v
	return s
}

// GetImageId returns the ImageId value or empty string if nil
func (s *UpdateImageReserveMinAmountRequest) GetImageId() string {
	if s == nil || s.ImageId == nil {
		return ""
	}
	return *s.ImageId
}

// SetReserveMinAmount sets the ReserveMinAmount field
func (s *UpdateImageReserveMinAmountRequest) SetReserveMinAmount(v int32) *UpdateImageReserveMinAmountRequest {
	s.ReserveMinAmount = &v
	return s
}

// GetReserveMinAmount returns the ReserveMinAmount value or 0 if nil
func (s *UpdateImageReserveMinAmountRequest) GetReserveMinAmount() int32 {
	if s == nil || s.ReserveMinAmount == nil {
		return 0
	}
	return *s.ReserveMinAmount
}
