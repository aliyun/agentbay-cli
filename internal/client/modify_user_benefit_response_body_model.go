// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserBenefitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyUserBenefitResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ModifyUserBenefitResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyUserBenefitResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyUserBenefitResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyUserBenefitResponseBody
	GetSuccess() *bool
}

type ModifyUserBenefitResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyUserBenefitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserBenefitResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUserBenefitResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyUserBenefitResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyUserBenefitResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyUserBenefitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyUserBenefitResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyUserBenefitResponseBody) SetCode(v string) *ModifyUserBenefitResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyUserBenefitResponseBody) SetHttpStatusCode(v int32) *ModifyUserBenefitResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyUserBenefitResponseBody) SetMessage(v string) *ModifyUserBenefitResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyUserBenefitResponseBody) SetRequestId(v string) *ModifyUserBenefitResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyUserBenefitResponseBody) SetSuccess(v bool) *ModifyUserBenefitResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyUserBenefitResponseBody) Validate() error {
	return dara.Validate(s)
}
