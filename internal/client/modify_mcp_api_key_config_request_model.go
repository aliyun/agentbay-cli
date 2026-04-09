// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"errors"
)

// ModifyMcpApiKeyConfigRequest is the request struct for ModifyMcpApiKeyConfig
type ModifyMcpApiKeyConfigRequest struct {
	ApiKeyId    *string `json:"ApiKeyId,omitempty" xml:"ApiKeyId,omitempty"`
	Concurrency *int32  `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
}

// Validate validates the ModifyMcpApiKeyConfigRequest
func (s *ModifyMcpApiKeyConfigRequest) Validate() error {
	if s.ApiKeyId == nil || *s.ApiKeyId == "" {
		return errors.New("ApiKeyId is required")
	}
	if s.Concurrency != nil && (*s.Concurrency < 1) {
		return errors.New("Concurrency must be greater than or equal to 1")
	}
	return nil
}
