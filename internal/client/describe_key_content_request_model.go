// This file is auto-generated, don't edit it. Thanks.
package client

import "errors"

// DescribeKeyContentRequest is the request struct for DescribeKeyContent
type DescribeKeyContentRequest struct {
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
}

// Validate checks that required fields are present
func (s *DescribeKeyContentRequest) Validate() error {
	if s.KeyId == nil || *s.KeyId == "" {
		return errors.New("KeyId is required")
	}
	return nil
}
