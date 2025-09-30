// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcpImageInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *GetMcpImageInfoRequest
	GetImageId() *string
}

type GetMcpImageInfoRequest struct {
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
}

func (s GetMcpImageInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMcpImageInfoRequest) GoString() string {
	return s.String()
}

func (s *GetMcpImageInfoRequest) GetImageId() *string {
	return s.ImageId
}

func (s *GetMcpImageInfoRequest) SetImageId(v string) *GetMcpImageInfoRequest {
	s.ImageId = &v
	return s
}

func (s *GetMcpImageInfoRequest) Validate() error {
	return dara.Validate(s)
}
