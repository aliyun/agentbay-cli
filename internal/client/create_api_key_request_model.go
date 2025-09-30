// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateApiKeyRequest
	GetDescription() *string
	SetName(v string) *CreateApiKeyRequest
	GetName() *string
}

type CreateApiKeyRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApiKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateApiKeyRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateApiKeyRequest) GetName() *string {
	return s.Name
}

func (s *CreateApiKeyRequest) SetDescription(v string) *CreateApiKeyRequest {
	s.Description = &v
	return s
}

func (s *CreateApiKeyRequest) SetName(v string) *CreateApiKeyRequest {
	s.Name = &v
	return s
}

func (s *CreateApiKeyRequest) Validate() error {
	return dara.Validate(s)
}
