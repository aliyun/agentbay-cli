// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateImagePropertiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v string) *UpdateImagePropertiesRequest
	GetApiKeyId() *string
	SetImageProperties(v string) *UpdateImagePropertiesRequest
	GetImageProperties() *string
}

type UpdateImagePropertiesRequest struct {
	ApiKeyId        *string `json:"ApiKeyId,omitempty" xml:"ApiKeyId,omitempty"`
	ImageProperties *string `json:"ImageProperties,omitempty" xml:"ImageProperties,omitempty"`
}

func (s UpdateImagePropertiesRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateImagePropertiesRequest) GoString() string {
	return s.String()
}

func (s *UpdateImagePropertiesRequest) GetApiKeyId() *string {
	return s.ApiKeyId
}

func (s *UpdateImagePropertiesRequest) GetImageProperties() *string {
	return s.ImageProperties
}

func (s *UpdateImagePropertiesRequest) SetApiKeyId(v string) *UpdateImagePropertiesRequest {
	s.ApiKeyId = &v
	return s
}

func (s *UpdateImagePropertiesRequest) SetImageProperties(v string) *UpdateImagePropertiesRequest {
	s.ImageProperties = &v
	return s
}

func (s *UpdateImagePropertiesRequest) Validate() error {
	return dara.Validate(s)
}
