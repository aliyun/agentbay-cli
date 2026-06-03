// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMarketSkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMarketSkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMarketSkillResponse
	GetStatusCode() *int32
	SetBody(v *CreateMarketSkillResponseBody) *UpdateMarketSkillResponse
	GetBody() *CreateMarketSkillResponseBody
}

// UpdateMarketSkillResponse wraps the response for UpdateMarketSkill.
// Body reuses CreateMarketSkillResponseBody since Update returns the same structure as Create.
type UpdateMarketSkillResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMarketSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
	// RawBody is the raw response body string; printed on error for debugging.
	RawBody string `json:"-"`
}

func (s UpdateMarketSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMarketSkillResponse) GoString() string {
	return s.String()
}

func (s *UpdateMarketSkillResponse) GetHeaders() map[string]*string { return s.Headers }
func (s *UpdateMarketSkillResponse) SetHeaders(v map[string]*string) *UpdateMarketSkillResponse {
	s.Headers = v
	return s
}
func (s *UpdateMarketSkillResponse) GetStatusCode() *int32 { return s.StatusCode }
func (s *UpdateMarketSkillResponse) SetStatusCode(v int32) *UpdateMarketSkillResponse {
	s.StatusCode = &v
	return s
}
func (s *UpdateMarketSkillResponse) GetBody() *CreateMarketSkillResponseBody { return s.Body }
func (s *UpdateMarketSkillResponse) SetBody(v *CreateMarketSkillResponseBody) *UpdateMarketSkillResponse {
	s.Body = v
	return s
}
func (s *UpdateMarketSkillResponse) Validate() error {
	return dara.Validate(s)
}
