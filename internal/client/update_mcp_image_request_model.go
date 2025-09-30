// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMcpImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v string) *UpdateMcpImageRequest
	GetApiKeyId() *string
	SetImageId(v string) *UpdateMcpImageRequest
	GetImageId() *string
}

type UpdateMcpImageRequest struct {
	ApiKeyId *string `json:"ApiKeyId,omitempty" xml:"ApiKeyId,omitempty"`
	ImageId  *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
}

func (s UpdateMcpImageRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpImageRequest) GoString() string {
	return s.String()
}

func (s *UpdateMcpImageRequest) GetApiKeyId() *string {
	return s.ApiKeyId
}

func (s *UpdateMcpImageRequest) GetImageId() *string {
	return s.ImageId
}

func (s *UpdateMcpImageRequest) SetApiKeyId(v string) *UpdateMcpImageRequest {
	s.ApiKeyId = &v
	return s
}

func (s *UpdateMcpImageRequest) SetImageId(v string) *UpdateMcpImageRequest {
	s.ImageId = &v
	return s
}

func (s *UpdateMcpImageRequest) Validate() error {
	return dara.Validate(s)
}
