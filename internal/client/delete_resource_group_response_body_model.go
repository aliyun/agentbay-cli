// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteResourceGroupResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteResourceGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteResourceGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteResourceGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteResourceGroupResponseBody
	GetSuccess() *bool
}

type DeleteResourceGroupResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteResourceGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteResourceGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteResourceGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteResourceGroupResponseBody) SetCode(v string) *DeleteResourceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteResourceGroupResponseBody) SetHttpStatusCode(v int32) *DeleteResourceGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteResourceGroupResponseBody) SetMessage(v string) *DeleteResourceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteResourceGroupResponseBody) SetRequestId(v string) *DeleteResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteResourceGroupResponseBody) SetSuccess(v bool) *DeleteResourceGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
