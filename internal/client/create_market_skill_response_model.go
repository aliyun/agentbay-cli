// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMarketSkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMarketSkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMarketSkillResponse
	GetStatusCode() *int32
	SetBody(v *CreateMarketSkillResponseBody) *CreateMarketSkillResponse
	GetBody() *CreateMarketSkillResponseBody
}

type CreateMarketSkillResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMarketSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
	// RawBody is the raw response body string; printed on error for debugging.
	RawBody string `json:"-"`
}

func (s CreateMarketSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMarketSkillResponse) GoString() string {
	return s.String()
}

func (s *CreateMarketSkillResponse) GetHeaders() map[string]*string { return s.Headers }
func (s *CreateMarketSkillResponse) SetHeaders(v map[string]*string) *CreateMarketSkillResponse {
	s.Headers = v
	return s
}
func (s *CreateMarketSkillResponse) GetStatusCode() *int32 { return s.StatusCode }
func (s *CreateMarketSkillResponse) SetStatusCode(v int32) *CreateMarketSkillResponse {
	s.StatusCode = &v
	return s
}
func (s *CreateMarketSkillResponse) GetBody() *CreateMarketSkillResponseBody { return s.Body }
func (s *CreateMarketSkillResponse) SetBody(v *CreateMarketSkillResponseBody) *CreateMarketSkillResponse {
	s.Body = v
	return s
}
func (s *CreateMarketSkillResponse) Validate() error {
	return dara.Validate(s)
}
