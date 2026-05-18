// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"encoding/json"
	"errors"
)

// DeleteApiKeyRequest is the request struct for DeleteApiKey
type DeleteApiKeyRequest struct {
	// KeyIdListJson is a JSON array string of API key IDs to delete, e.g. `["ak-xxx"]`
	KeyIdListJson *string `json:"KeyIdListJson,omitempty" xml:"KeyIdListJson,omitempty"`
}

// Validate validates the DeleteApiKeyRequest
func (s *DeleteApiKeyRequest) Validate() error {
	if s.KeyIdListJson == nil || *s.KeyIdListJson == "" {
		return errors.New("KeyIdListJson is required")
	}
	// Validate that it is a valid JSON array
	var ids []string
	if err := json.Unmarshal([]byte(*s.KeyIdListJson), &ids); err != nil {
		return errors.New("KeyIdListJson must be a valid JSON array of strings")
	}
	if len(ids) == 0 {
		return errors.New("KeyIdListJson must contain at least one key ID")
	}
	return nil
}

// GetKeyIdListJson returns the KeyIdListJson value or empty string if nil
func (s *DeleteApiKeyRequest) GetKeyIdListJson() string {
	if s == nil || s.KeyIdListJson == nil {
		return ""
	}
	return *s.KeyIdListJson
}

// SetKeyIdListJson sets the KeyIdListJson value
func (s *DeleteApiKeyRequest) SetKeyIdListJson(v string) *DeleteApiKeyRequest {
	s.KeyIdListJson = &v
	return s
}
