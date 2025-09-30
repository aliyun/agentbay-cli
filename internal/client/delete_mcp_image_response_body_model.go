// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcpImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteMcpImageResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteMcpImageResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteMcpImageResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteMcpImageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteMcpImageResponseBody
	GetSuccess() *bool
}

type DeleteMcpImageResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMcpImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcpImageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMcpImageResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteMcpImageResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteMcpImageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteMcpImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMcpImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMcpImageResponseBody) SetCode(v string) *DeleteMcpImageResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMcpImageResponseBody) SetHttpStatusCode(v int32) *DeleteMcpImageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteMcpImageResponseBody) SetMessage(v string) *DeleteMcpImageResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMcpImageResponseBody) SetRequestId(v string) *DeleteMcpImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMcpImageResponseBody) SetSuccess(v bool) *DeleteMcpImageResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMcpImageResponseBody) Validate() error {
	return dara.Validate(s)
}
