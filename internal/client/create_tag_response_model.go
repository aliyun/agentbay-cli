// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTagResponse
	GetStatusCode() *int32
	SetBody(v *CreateTagResponseBody) *CreateTagResponse
	GetBody() *CreateTagResponseBody
}

type CreateTagResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
	// RawBody is the raw response body string; printed on error for debugging.
	RawBody string `json:"-"`
}

func (s CreateTagResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTagResponse) GoString() string {
	return s.String()
}

func (s *CreateTagResponse) GetHeaders() map[string]*string { return s.Headers }
func (s *CreateTagResponse) SetHeaders(v map[string]*string) *CreateTagResponse {
	s.Headers = v
	return s
}
func (s *CreateTagResponse) GetStatusCode() *int32 { return s.StatusCode }
func (s *CreateTagResponse) SetStatusCode(v int32) *CreateTagResponse {
	s.StatusCode = &v
	return s
}
func (s *CreateTagResponse) GetBody() *CreateTagResponseBody { return s.Body }
func (s *CreateTagResponse) SetBody(v *CreateTagResponseBody) *CreateTagResponse {
	s.Body = v
	return s
}
func (s *CreateTagResponse) Validate() error {
	return dara.Validate(s)
}
