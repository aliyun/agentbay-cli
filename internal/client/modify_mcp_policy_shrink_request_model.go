// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMcpPolicyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConcurrency(v int32) *ModifyMcpPolicyShrinkRequest
	GetConcurrency() *int32
	SetDebugMode(v bool) *ModifyMcpPolicyShrinkRequest
	GetDebugMode() *bool
	SetDesktopMaxRuntime(v int32) *ModifyMcpPolicyShrinkRequest
	GetDesktopMaxRuntime() *int32
	SetDisplayConfigShrink(v string) *ModifyMcpPolicyShrinkRequest
	GetDisplayConfigShrink() *string
	SetIdleTimeoutSwitch(v bool) *ModifyMcpPolicyShrinkRequest
	GetIdleTimeoutSwitch() *bool
	SetMcpIdleTimeout(v int32) *ModifyMcpPolicyShrinkRequest
	GetMcpIdleTimeout() *int32
	SetNetworkConfigShrink(v string) *ModifyMcpPolicyShrinkRequest
	GetNetworkConfigShrink() *string
	SetPolicyId(v string) *ModifyMcpPolicyShrinkRequest
	GetPolicyId() *string
	SetPolicyName(v string) *ModifyMcpPolicyShrinkRequest
	GetPolicyName() *string
	SetUserIdleTimeout(v int32) *ModifyMcpPolicyShrinkRequest
	GetUserIdleTimeout() *int32
}

type ModifyMcpPolicyShrinkRequest struct {
	Concurrency         *int32  `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	DebugMode           *bool   `json:"DebugMode,omitempty" xml:"DebugMode,omitempty"`
	DesktopMaxRuntime   *int32  `json:"DesktopMaxRuntime,omitempty" xml:"DesktopMaxRuntime,omitempty"`
	DisplayConfigShrink *string `json:"DisplayConfig,omitempty" xml:"DisplayConfig,omitempty"`
	IdleTimeoutSwitch   *bool   `json:"IdleTimeoutSwitch,omitempty" xml:"IdleTimeoutSwitch,omitempty"`
	McpIdleTimeout      *int32  `json:"McpIdleTimeout,omitempty" xml:"McpIdleTimeout,omitempty"`
	NetworkConfigShrink *string `json:"NetworkConfig,omitempty" xml:"NetworkConfig,omitempty"`
	PolicyId            *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	PolicyName          *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	UserIdleTimeout     *int32  `json:"UserIdleTimeout,omitempty" xml:"UserIdleTimeout,omitempty"`
}

func (s ModifyMcpPolicyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyMcpPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyMcpPolicyShrinkRequest) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *ModifyMcpPolicyShrinkRequest) GetDebugMode() *bool {
	return s.DebugMode
}

func (s *ModifyMcpPolicyShrinkRequest) GetDesktopMaxRuntime() *int32 {
	return s.DesktopMaxRuntime
}

func (s *ModifyMcpPolicyShrinkRequest) GetDisplayConfigShrink() *string {
	return s.DisplayConfigShrink
}

func (s *ModifyMcpPolicyShrinkRequest) GetIdleTimeoutSwitch() *bool {
	return s.IdleTimeoutSwitch
}

func (s *ModifyMcpPolicyShrinkRequest) GetMcpIdleTimeout() *int32 {
	return s.McpIdleTimeout
}

func (s *ModifyMcpPolicyShrinkRequest) GetNetworkConfigShrink() *string {
	return s.NetworkConfigShrink
}

func (s *ModifyMcpPolicyShrinkRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ModifyMcpPolicyShrinkRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ModifyMcpPolicyShrinkRequest) GetUserIdleTimeout() *int32 {
	return s.UserIdleTimeout
}

func (s *ModifyMcpPolicyShrinkRequest) SetConcurrency(v int32) *ModifyMcpPolicyShrinkRequest {
	s.Concurrency = &v
	return s
}

func (s *ModifyMcpPolicyShrinkRequest) SetDebugMode(v bool) *ModifyMcpPolicyShrinkRequest {
	s.DebugMode = &v
	return s
}

func (s *ModifyMcpPolicyShrinkRequest) SetDesktopMaxRuntime(v int32) *ModifyMcpPolicyShrinkRequest {
	s.DesktopMaxRuntime = &v
	return s
}

func (s *ModifyMcpPolicyShrinkRequest) SetDisplayConfigShrink(v string) *ModifyMcpPolicyShrinkRequest {
	s.DisplayConfigShrink = &v
	return s
}

func (s *ModifyMcpPolicyShrinkRequest) SetIdleTimeoutSwitch(v bool) *ModifyMcpPolicyShrinkRequest {
	s.IdleTimeoutSwitch = &v
	return s
}

func (s *ModifyMcpPolicyShrinkRequest) SetMcpIdleTimeout(v int32) *ModifyMcpPolicyShrinkRequest {
	s.McpIdleTimeout = &v
	return s
}

func (s *ModifyMcpPolicyShrinkRequest) SetNetworkConfigShrink(v string) *ModifyMcpPolicyShrinkRequest {
	s.NetworkConfigShrink = &v
	return s
}

func (s *ModifyMcpPolicyShrinkRequest) SetPolicyId(v string) *ModifyMcpPolicyShrinkRequest {
	s.PolicyId = &v
	return s
}

func (s *ModifyMcpPolicyShrinkRequest) SetPolicyName(v string) *ModifyMcpPolicyShrinkRequest {
	s.PolicyName = &v
	return s
}

func (s *ModifyMcpPolicyShrinkRequest) SetUserIdleTimeout(v int32) *ModifyMcpPolicyShrinkRequest {
	s.UserIdleTimeout = &v
	return s
}

func (s *ModifyMcpPolicyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
