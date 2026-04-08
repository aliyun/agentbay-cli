// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"errors"
)

// CreateApiKeyRequest is the request struct for CreateApiKey
type CreateApiKeyRequest struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

// Validate validates the CreateApiKeyRequest
func (s *CreateApiKeyRequest) Validate() error {
	if s.Name == nil || *s.Name == "" {
		return errors.New("Name is required")
	}
	return nil
}
