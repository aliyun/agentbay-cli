// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDesktopReleaseTimingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v string) *UpdateDesktopReleaseTimingsRequest
	GetApiKeyId() *string
	SetDesktopMaxRuntime(v int32) *UpdateDesktopReleaseTimingsRequest
	GetDesktopMaxRuntime() *int32
	SetMcpIdleTimeout(v int32) *UpdateDesktopReleaseTimingsRequest
	GetMcpIdleTimeout() *int32
	SetUserIdleTimeout(v int32) *UpdateDesktopReleaseTimingsRequest
	GetUserIdleTimeout() *int32
}

type UpdateDesktopReleaseTimingsRequest struct {
	ApiKeyId          *string `json:"ApiKeyId,omitempty" xml:"ApiKeyId,omitempty"`
	DesktopMaxRuntime *int32  `json:"DesktopMaxRuntime,omitempty" xml:"DesktopMaxRuntime,omitempty"`
	McpIdleTimeout    *int32  `json:"McpIdleTimeout,omitempty" xml:"McpIdleTimeout,omitempty"`
	UserIdleTimeout   *int32  `json:"UserIdleTimeout,omitempty" xml:"UserIdleTimeout,omitempty"`
}

func (s UpdateDesktopReleaseTimingsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDesktopReleaseTimingsRequest) GoString() string {
	return s.String()
}

func (s *UpdateDesktopReleaseTimingsRequest) GetApiKeyId() *string {
	return s.ApiKeyId
}

func (s *UpdateDesktopReleaseTimingsRequest) GetDesktopMaxRuntime() *int32 {
	return s.DesktopMaxRuntime
}

func (s *UpdateDesktopReleaseTimingsRequest) GetMcpIdleTimeout() *int32 {
	return s.McpIdleTimeout
}

func (s *UpdateDesktopReleaseTimingsRequest) GetUserIdleTimeout() *int32 {
	return s.UserIdleTimeout
}

func (s *UpdateDesktopReleaseTimingsRequest) SetApiKeyId(v string) *UpdateDesktopReleaseTimingsRequest {
	s.ApiKeyId = &v
	return s
}

func (s *UpdateDesktopReleaseTimingsRequest) SetDesktopMaxRuntime(v int32) *UpdateDesktopReleaseTimingsRequest {
	s.DesktopMaxRuntime = &v
	return s
}

func (s *UpdateDesktopReleaseTimingsRequest) SetMcpIdleTimeout(v int32) *UpdateDesktopReleaseTimingsRequest {
	s.McpIdleTimeout = &v
	return s
}

func (s *UpdateDesktopReleaseTimingsRequest) SetUserIdleTimeout(v int32) *UpdateDesktopReleaseTimingsRequest {
	s.UserIdleTimeout = &v
	return s
}

func (s *UpdateDesktopReleaseTimingsRequest) Validate() error {
	return dara.Validate(s)
}
