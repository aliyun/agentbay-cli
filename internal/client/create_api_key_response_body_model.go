// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateApiKeyResponseBody
	GetCode() *string
	SetData(v string) *CreateApiKeyResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateApiKeyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateApiKeyResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateApiKeyResponseBody
	GetRequestId() *string
}

type CreateApiKeyResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApiKeyResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateApiKeyResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateApiKeyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateApiKeyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApiKeyResponseBody) SetCode(v string) *CreateApiKeyResponseBody {
	s.Code = &v
	return s
}

func (s *CreateApiKeyResponseBody) SetData(v string) *CreateApiKeyResponseBody {
	s.Data = &v
	return s
}

func (s *CreateApiKeyResponseBody) SetHttpStatusCode(v int32) *CreateApiKeyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateApiKeyResponseBody) SetMessage(v string) *CreateApiKeyResponseBody {
	s.Message = &v
	return s
}

func (s *CreateApiKeyResponseBody) SetRequestId(v string) *CreateApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApiKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
