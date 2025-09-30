// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMcpResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddMcpResourceResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *AddMcpResourceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddMcpResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddMcpResourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddMcpResourceResponseBody
	GetSuccess() *bool
}

type AddMcpResourceResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddMcpResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddMcpResourceResponseBody) GoString() string {
	return s.String()
}

func (s *AddMcpResourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddMcpResourceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddMcpResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddMcpResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddMcpResourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddMcpResourceResponseBody) SetCode(v string) *AddMcpResourceResponseBody {
	s.Code = &v
	return s
}

func (s *AddMcpResourceResponseBody) SetHttpStatusCode(v int32) *AddMcpResourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddMcpResourceResponseBody) SetMessage(v string) *AddMcpResourceResponseBody {
	s.Message = &v
	return s
}

func (s *AddMcpResourceResponseBody) SetRequestId(v string) *AddMcpResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddMcpResourceResponseBody) SetSuccess(v bool) *AddMcpResourceResponseBody {
	s.Success = &v
	return s
}

func (s *AddMcpResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
