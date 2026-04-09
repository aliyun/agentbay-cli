// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"errors"
)

// {Action}Request is the request struct for {Action}
type {Action}Request struct {
	{Field1} *string `{Field1Tag}`
	{Field2} *int32  `{Field2Tag}`
}

// Validate validates the {Action}Request
func (s *{Action}Request) Validate() error {
	if s.{RequiredField} == nil || *s.{RequiredField} == "" {
		return errors.New("{RequiredField} is required")
	}
	return nil
}
