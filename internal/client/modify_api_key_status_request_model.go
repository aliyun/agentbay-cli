// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"errors"
)

// ModifyApiKeyStatusRequest is the request struct for ModifyApiKeyStatus
type ModifyApiKeyStatusRequest struct {
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

// Validate validates the ModifyApiKeyStatusRequest
func (s *ModifyApiKeyStatusRequest) Validate() error {
	if s.ApiKey == nil || *s.ApiKey == "" {
		return errors.New("ApiKey is required")
	}
	if s.Status == nil || *s.Status == "" {
		return errors.New("Status is required")
	}
	if *s.Status != "ENABLED" && *s.Status != "DISABLED" {
		return errors.New("Status must be ENABLED or DISABLED")
	}
	return nil
}
