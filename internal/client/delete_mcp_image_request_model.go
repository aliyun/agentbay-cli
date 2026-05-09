// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcpImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *DeleteMcpImageRequest
	GetImageId() *string
}

type DeleteMcpImageRequest struct {
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
}

func (s DeleteMcpImageRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcpImageRequest) GoString() string {
	return s.String()
}

func (s *DeleteMcpImageRequest) GetImageId() *string {
	return s.ImageId
}

func (s *DeleteMcpImageRequest) SetImageId(v string) *DeleteMcpImageRequest {
	s.ImageId = &v
	return s
}

func (s *DeleteMcpImageRequest) Validate() error {
	return dara.Validate(s)
}
