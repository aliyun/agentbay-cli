// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceGroupSessionBandwidthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateResourceGroupSessionBandwidthResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateResourceGroupSessionBandwidthResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateResourceGroupSessionBandwidthResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateResourceGroupSessionBandwidthResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateResourceGroupSessionBandwidthResponseBody
	GetSuccess() *bool
}

type UpdateResourceGroupSessionBandwidthResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateResourceGroupSessionBandwidthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceGroupSessionBandwidthResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceGroupSessionBandwidthResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateResourceGroupSessionBandwidthResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateResourceGroupSessionBandwidthResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateResourceGroupSessionBandwidthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateResourceGroupSessionBandwidthResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateResourceGroupSessionBandwidthResponseBody) SetCode(v string) *UpdateResourceGroupSessionBandwidthResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateResourceGroupSessionBandwidthResponseBody) SetHttpStatusCode(v int32) *UpdateResourceGroupSessionBandwidthResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateResourceGroupSessionBandwidthResponseBody) SetMessage(v string) *UpdateResourceGroupSessionBandwidthResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateResourceGroupSessionBandwidthResponseBody) SetRequestId(v string) *UpdateResourceGroupSessionBandwidthResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateResourceGroupSessionBandwidthResponseBody) SetSuccess(v bool) *UpdateResourceGroupSessionBandwidthResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateResourceGroupSessionBandwidthResponseBody) Validate() error {
	return dara.Validate(s)
}
