// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcpPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteMcpPoliciesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteMcpPoliciesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteMcpPoliciesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteMcpPoliciesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteMcpPoliciesResponseBody
	GetSuccess() *bool
}

type DeleteMcpPoliciesResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMcpPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcpPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMcpPoliciesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteMcpPoliciesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteMcpPoliciesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteMcpPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMcpPoliciesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMcpPoliciesResponseBody) SetCode(v string) *DeleteMcpPoliciesResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMcpPoliciesResponseBody) SetHttpStatusCode(v int32) *DeleteMcpPoliciesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteMcpPoliciesResponseBody) SetMessage(v string) *DeleteMcpPoliciesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMcpPoliciesResponseBody) SetRequestId(v string) *DeleteMcpPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMcpPoliciesResponseBody) SetSuccess(v bool) *DeleteMcpPoliciesResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMcpPoliciesResponseBody) Validate() error {
	return dara.Validate(s)
}
