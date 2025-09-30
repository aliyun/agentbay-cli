// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMcpImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateMcpImageResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateMcpImageResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateMcpImageResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateMcpImageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateMcpImageResponseBody
	GetSuccess() *bool
}

type UpdateMcpImageResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMcpImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpImageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMcpImageResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateMcpImageResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateMcpImageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateMcpImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMcpImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMcpImageResponseBody) SetCode(v string) *UpdateMcpImageResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateMcpImageResponseBody) SetHttpStatusCode(v int32) *UpdateMcpImageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateMcpImageResponseBody) SetMessage(v string) *UpdateMcpImageResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateMcpImageResponseBody) SetRequestId(v string) *UpdateMcpImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMcpImageResponseBody) SetSuccess(v bool) *UpdateMcpImageResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMcpImageResponseBody) Validate() error {
	return dara.Validate(s)
}
