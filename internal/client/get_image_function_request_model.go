// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageFunctionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *GetImageFunctionRequest
	GetImageId() *string
}

type GetImageFunctionRequest struct {
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
}

func (s GetImageFunctionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetImageFunctionRequest) GoString() string {
	return s.String()
}

func (s *GetImageFunctionRequest) GetImageId() *string {
	return s.ImageId
}

func (s *GetImageFunctionRequest) SetImageId(v string) *GetImageFunctionRequest {
	s.ImageId = &v
	return s
}

func (s *GetImageFunctionRequest) Validate() error {
	return dara.Validate(s)
}
