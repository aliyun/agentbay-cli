// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"errors"
)

// DescribeMcpApiKeyRequest is the request struct for DescribeMcpApiKey
type DescribeMcpApiKeyRequest struct {
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
}

// Validate validates the DescribeMcpApiKeyRequest
func (s *DescribeMcpApiKeyRequest) Validate() error {
	if s.ApiKey == nil || *s.ApiKey == "" {
		return errors.New("ApiKey is required")
	}
	return nil
}
