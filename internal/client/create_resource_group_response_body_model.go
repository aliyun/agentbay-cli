// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateResourceGroupResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateResourceGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateResourceGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateResourceGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateResourceGroupResponseBody
	GetSuccess() *bool
}

type CreateResourceGroupResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateResourceGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateResourceGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateResourceGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateResourceGroupResponseBody) SetCode(v string) *CreateResourceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateResourceGroupResponseBody) SetHttpStatusCode(v int32) *CreateResourceGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateResourceGroupResponseBody) SetMessage(v string) *CreateResourceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateResourceGroupResponseBody) SetRequestId(v string) *CreateResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResourceGroupResponseBody) SetSuccess(v bool) *CreateResourceGroupResponseBody {
	s.Success = &v
	return s
}

func (s *CreateResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
