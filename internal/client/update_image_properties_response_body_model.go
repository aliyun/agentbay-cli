// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateImagePropertiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateImagePropertiesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateImagePropertiesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateImagePropertiesResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateImagePropertiesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateImagePropertiesResponseBody
	GetSuccess() *bool
}

type UpdateImagePropertiesResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateImagePropertiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateImagePropertiesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateImagePropertiesResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateImagePropertiesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateImagePropertiesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateImagePropertiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateImagePropertiesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateImagePropertiesResponseBody) SetCode(v string) *UpdateImagePropertiesResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateImagePropertiesResponseBody) SetHttpStatusCode(v int32) *UpdateImagePropertiesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateImagePropertiesResponseBody) SetMessage(v string) *UpdateImagePropertiesResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateImagePropertiesResponseBody) SetRequestId(v string) *UpdateImagePropertiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateImagePropertiesResponseBody) SetSuccess(v bool) *UpdateImagePropertiesResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateImagePropertiesResponseBody) Validate() error {
	return dara.Validate(s)
}
