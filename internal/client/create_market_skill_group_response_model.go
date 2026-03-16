// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMarketSkillGroupResponseBodyData interface {
	dara.Model
	String() string
	GoString() string
	GetGroupId() *string
	SetGroupId(v string) *CreateMarketSkillGroupResponseBodyData
}

type CreateMarketSkillGroupResponseBodyData struct {
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s CreateMarketSkillGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateMarketSkillGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateMarketSkillGroupResponseBodyData) GetGroupId() *string { return s.GroupId }
func (s *CreateMarketSkillGroupResponseBodyData) SetGroupId(v string) *CreateMarketSkillGroupResponseBodyData {
	s.GroupId = &v
	return s
}
func (s *CreateMarketSkillGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type iCreateMarketSkillGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateMarketSkillGroupResponseBody
	GetCode() *string
	SetData(v *CreateMarketSkillGroupResponseBodyData) *CreateMarketSkillGroupResponseBody
	GetData() *CreateMarketSkillGroupResponseBodyData
	SetRequestId(v string) *CreateMarketSkillGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateMarketSkillGroupResponseBody
	GetSuccess() *bool
}

type CreateMarketSkillGroupResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateMarketSkillGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMarketSkillGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMarketSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMarketSkillGroupResponseBody) GetCode() *string  { return s.Code }
func (s *CreateMarketSkillGroupResponseBody) SetCode(v string) *CreateMarketSkillGroupResponseBody {
	s.Code = &v
	return s
}
func (s *CreateMarketSkillGroupResponseBody) GetData() *CreateMarketSkillGroupResponseBodyData {
	return s.Data
}
func (s *CreateMarketSkillGroupResponseBody) SetData(v *CreateMarketSkillGroupResponseBodyData) *CreateMarketSkillGroupResponseBody {
	s.Data = v
	return s
}
func (s *CreateMarketSkillGroupResponseBody) GetRequestId() *string { return s.RequestId }
func (s *CreateMarketSkillGroupResponseBody) SetRequestId(v string) *CreateMarketSkillGroupResponseBody {
	s.RequestId = &v
	return s
}
func (s *CreateMarketSkillGroupResponseBody) GetSuccess() *bool { return s.Success }
func (s *CreateMarketSkillGroupResponseBody) SetSuccess(v bool) *CreateMarketSkillGroupResponseBody {
	s.Success = &v
	return s
}
func (s *CreateMarketSkillGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type iCreateMarketSkillGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMarketSkillGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMarketSkillGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateMarketSkillGroupResponseBody) *CreateMarketSkillGroupResponse
	GetBody() *CreateMarketSkillGroupResponseBody
}

type CreateMarketSkillGroupResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMarketSkillGroupResponseBody   `json:"body,omitempty" xml:"body,omitempty"`
	// RawBody is the raw response body (set when BodyType is "string"); used for -v debug output.
	RawBody string `json:"-"`
}

func (s CreateMarketSkillGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMarketSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateMarketSkillGroupResponse) GetHeaders() map[string]*string    { return s.Headers }
func (s *CreateMarketSkillGroupResponse) SetHeaders(v map[string]*string) *CreateMarketSkillGroupResponse {
	s.Headers = v
	return s
}
func (s *CreateMarketSkillGroupResponse) GetStatusCode() *int32 { return s.StatusCode }
func (s *CreateMarketSkillGroupResponse) SetStatusCode(v int32) *CreateMarketSkillGroupResponse {
	s.StatusCode = &v
	return s
}
func (s *CreateMarketSkillGroupResponse) GetBody() *CreateMarketSkillGroupResponseBody { return s.Body }
func (s *CreateMarketSkillGroupResponse) SetBody(v *CreateMarketSkillGroupResponseBody) *CreateMarketSkillGroupResponse {
	s.Body = v
	return s
}
func (s *CreateMarketSkillGroupResponse) Validate() error {
	return dara.Validate(s)
}
