// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcpPolicyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConcurrency(v int32) *CreateMcpPolicyShrinkRequest
	GetConcurrency() *int32
	SetDebugMode(v bool) *CreateMcpPolicyShrinkRequest
	GetDebugMode() *bool
	SetDesktopMaxRuntime(v int32) *CreateMcpPolicyShrinkRequest
	GetDesktopMaxRuntime() *int32
	SetDisplayConfigShrink(v string) *CreateMcpPolicyShrinkRequest
	GetDisplayConfigShrink() *string
	SetIdleTimeoutSwitch(v bool) *CreateMcpPolicyShrinkRequest
	GetIdleTimeoutSwitch() *bool
	SetMcpIdleTimeout(v int32) *CreateMcpPolicyShrinkRequest
	GetMcpIdleTimeout() *int32
	SetNetworkConfigShrink(v string) *CreateMcpPolicyShrinkRequest
	GetNetworkConfigShrink() *string
	SetPolicyName(v string) *CreateMcpPolicyShrinkRequest
	GetPolicyName() *string
	SetUserIdleTimeout(v int32) *CreateMcpPolicyShrinkRequest
	GetUserIdleTimeout() *int32
}

type CreateMcpPolicyShrinkRequest struct {
	Concurrency         *int32  `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	DebugMode           *bool   `json:"DebugMode,omitempty" xml:"DebugMode,omitempty"`
	DesktopMaxRuntime   *int32  `json:"DesktopMaxRuntime,omitempty" xml:"DesktopMaxRuntime,omitempty"`
	DisplayConfigShrink *string `json:"DisplayConfig,omitempty" xml:"DisplayConfig,omitempty"`
	IdleTimeoutSwitch   *bool   `json:"IdleTimeoutSwitch,omitempty" xml:"IdleTimeoutSwitch,omitempty"`
	McpIdleTimeout      *int32  `json:"McpIdleTimeout,omitempty" xml:"McpIdleTimeout,omitempty"`
	NetworkConfigShrink *string `json:"NetworkConfig,omitempty" xml:"NetworkConfig,omitempty"`
	PolicyName          *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	UserIdleTimeout     *int32  `json:"UserIdleTimeout,omitempty" xml:"UserIdleTimeout,omitempty"`
}

func (s CreateMcpPolicyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateMcpPolicyShrinkRequest) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *CreateMcpPolicyShrinkRequest) GetDebugMode() *bool {
	return s.DebugMode
}

func (s *CreateMcpPolicyShrinkRequest) GetDesktopMaxRuntime() *int32 {
	return s.DesktopMaxRuntime
}

func (s *CreateMcpPolicyShrinkRequest) GetDisplayConfigShrink() *string {
	return s.DisplayConfigShrink
}

func (s *CreateMcpPolicyShrinkRequest) GetIdleTimeoutSwitch() *bool {
	return s.IdleTimeoutSwitch
}

func (s *CreateMcpPolicyShrinkRequest) GetMcpIdleTimeout() *int32 {
	return s.McpIdleTimeout
}

func (s *CreateMcpPolicyShrinkRequest) GetNetworkConfigShrink() *string {
	return s.NetworkConfigShrink
}

func (s *CreateMcpPolicyShrinkRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *CreateMcpPolicyShrinkRequest) GetUserIdleTimeout() *int32 {
	return s.UserIdleTimeout
}

func (s *CreateMcpPolicyShrinkRequest) SetConcurrency(v int32) *CreateMcpPolicyShrinkRequest {
	s.Concurrency = &v
	return s
}

func (s *CreateMcpPolicyShrinkRequest) SetDebugMode(v bool) *CreateMcpPolicyShrinkRequest {
	s.DebugMode = &v
	return s
}

func (s *CreateMcpPolicyShrinkRequest) SetDesktopMaxRuntime(v int32) *CreateMcpPolicyShrinkRequest {
	s.DesktopMaxRuntime = &v
	return s
}

func (s *CreateMcpPolicyShrinkRequest) SetDisplayConfigShrink(v string) *CreateMcpPolicyShrinkRequest {
	s.DisplayConfigShrink = &v
	return s
}

func (s *CreateMcpPolicyShrinkRequest) SetIdleTimeoutSwitch(v bool) *CreateMcpPolicyShrinkRequest {
	s.IdleTimeoutSwitch = &v
	return s
}

func (s *CreateMcpPolicyShrinkRequest) SetMcpIdleTimeout(v int32) *CreateMcpPolicyShrinkRequest {
	s.McpIdleTimeout = &v
	return s
}

func (s *CreateMcpPolicyShrinkRequest) SetNetworkConfigShrink(v string) *CreateMcpPolicyShrinkRequest {
	s.NetworkConfigShrink = &v
	return s
}

func (s *CreateMcpPolicyShrinkRequest) SetPolicyName(v string) *CreateMcpPolicyShrinkRequest {
	s.PolicyName = &v
	return s
}

func (s *CreateMcpPolicyShrinkRequest) SetUserIdleTimeout(v int32) *CreateMcpPolicyShrinkRequest {
	s.UserIdleTimeout = &v
	return s
}

func (s *CreateMcpPolicyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
