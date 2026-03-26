// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMarketSkillCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMarketSkillCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMarketSkillCredentialResponse
	GetStatusCode() *int32
	SetBody(v *GetMarketSkillCredentialResponseBody) *GetMarketSkillCredentialResponse
	GetBody() *GetMarketSkillCredentialResponseBody
}

type GetMarketSkillCredentialResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMarketSkillCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMarketSkillCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMarketSkillCredentialResponse) GoString() string {
	return s.String()
}

func (s *GetMarketSkillCredentialResponse) GetHeaders() map[string]*string    { return s.Headers }
func (s *GetMarketSkillCredentialResponse) SetHeaders(v map[string]*string) *GetMarketSkillCredentialResponse {
	s.Headers = v
	return s
}
func (s *GetMarketSkillCredentialResponse) GetStatusCode() *int32 { return s.StatusCode }
func (s *GetMarketSkillCredentialResponse) SetStatusCode(v int32) *GetMarketSkillCredentialResponse {
	s.StatusCode = &v
	return s
}
func (s *GetMarketSkillCredentialResponse) GetBody() *GetMarketSkillCredentialResponseBody {
	return s.Body
}
func (s *GetMarketSkillCredentialResponse) SetBody(v *GetMarketSkillCredentialResponseBody) *GetMarketSkillCredentialResponse {
	s.Body = v
	return s
}
func (s *GetMarketSkillCredentialResponse) Validate() error {
	return dara.Validate(s)
}
