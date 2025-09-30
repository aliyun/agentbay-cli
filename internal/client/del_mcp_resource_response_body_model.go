// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDelMcpResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DelMcpResourceResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DelMcpResourceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DelMcpResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DelMcpResourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DelMcpResourceResponseBody
	GetSuccess() *bool
}

type DelMcpResourceResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DelMcpResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DelMcpResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DelMcpResourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DelMcpResourceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DelMcpResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DelMcpResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DelMcpResourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DelMcpResourceResponseBody) SetCode(v string) *DelMcpResourceResponseBody {
	s.Code = &v
	return s
}

func (s *DelMcpResourceResponseBody) SetHttpStatusCode(v int32) *DelMcpResourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DelMcpResourceResponseBody) SetMessage(v string) *DelMcpResourceResponseBody {
	s.Message = &v
	return s
}

func (s *DelMcpResourceResponseBody) SetRequestId(v string) *DelMcpResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DelMcpResourceResponseBody) SetSuccess(v bool) *DelMcpResourceResponseBody {
	s.Success = &v
	return s
}

func (s *DelMcpResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
