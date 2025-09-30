// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMcpApiKeyConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyMcpApiKeyConfigResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ModifyMcpApiKeyConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyMcpApiKeyConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyMcpApiKeyConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyMcpApiKeyConfigResponseBody
	GetSuccess() *bool
}

type ModifyMcpApiKeyConfigResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyMcpApiKeyConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyMcpApiKeyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMcpApiKeyConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyMcpApiKeyConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyMcpApiKeyConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyMcpApiKeyConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyMcpApiKeyConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyMcpApiKeyConfigResponseBody) SetCode(v string) *ModifyMcpApiKeyConfigResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyMcpApiKeyConfigResponseBody) SetHttpStatusCode(v int32) *ModifyMcpApiKeyConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyMcpApiKeyConfigResponseBody) SetMessage(v string) *ModifyMcpApiKeyConfigResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyMcpApiKeyConfigResponseBody) SetRequestId(v string) *ModifyMcpApiKeyConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMcpApiKeyConfigResponseBody) SetSuccess(v bool) *ModifyMcpApiKeyConfigResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyMcpApiKeyConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
