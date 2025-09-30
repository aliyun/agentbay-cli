// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserBenefitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyUserBenefitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyUserBenefitResponse
	GetStatusCode() *int32
	SetBody(v *ModifyUserBenefitResponseBody) *ModifyUserBenefitResponse
	GetBody() *ModifyUserBenefitResponseBody
}

type ModifyUserBenefitResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyUserBenefitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyUserBenefitResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserBenefitResponse) GoString() string {
	return s.String()
}

func (s *ModifyUserBenefitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyUserBenefitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyUserBenefitResponse) GetBody() *ModifyUserBenefitResponseBody {
	return s.Body
}

func (s *ModifyUserBenefitResponse) SetHeaders(v map[string]*string) *ModifyUserBenefitResponse {
	s.Headers = v
	return s
}

func (s *ModifyUserBenefitResponse) SetStatusCode(v int32) *ModifyUserBenefitResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyUserBenefitResponse) SetBody(v *ModifyUserBenefitResponseBody) *ModifyUserBenefitResponse {
	s.Body = v
	return s
}

func (s *ModifyUserBenefitResponse) Validate() error {
	return dara.Validate(s)
}
