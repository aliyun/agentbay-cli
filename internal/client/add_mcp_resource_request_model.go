// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMcpResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v string) *AddMcpResourceRequest
	GetApiKeyId() *string
	SetConcurrency(v string) *AddMcpResourceRequest
	GetConcurrency() *string
	SetDesktopMaxRuntime(v int32) *AddMcpResourceRequest
	GetDesktopMaxRuntime() *int32
	SetImageId(v string) *AddMcpResourceRequest
	GetImageId() *string
	SetMcpIdleTimeout(v int32) *AddMcpResourceRequest
	GetMcpIdleTimeout() *int32
	SetResourceType(v string) *AddMcpResourceRequest
	GetResourceType() *string
	SetUserIdleTimeout(v int32) *AddMcpResourceRequest
	GetUserIdleTimeout() *int32
	SetUserName(v string) *AddMcpResourceRequest
	GetUserName() *string
}

type AddMcpResourceRequest struct {
	ApiKeyId          *string `json:"ApiKeyId,omitempty" xml:"ApiKeyId,omitempty"`
	Concurrency       *string `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	DesktopMaxRuntime *int32  `json:"DesktopMaxRuntime,omitempty" xml:"DesktopMaxRuntime,omitempty"`
	ImageId           *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	McpIdleTimeout    *int32  `json:"McpIdleTimeout,omitempty" xml:"McpIdleTimeout,omitempty"`
	ResourceType      *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	UserIdleTimeout   *int32  `json:"UserIdleTimeout,omitempty" xml:"UserIdleTimeout,omitempty"`
	UserName          *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s AddMcpResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s AddMcpResourceRequest) GoString() string {
	return s.String()
}

func (s *AddMcpResourceRequest) GetApiKeyId() *string {
	return s.ApiKeyId
}

func (s *AddMcpResourceRequest) GetConcurrency() *string {
	return s.Concurrency
}

func (s *AddMcpResourceRequest) GetDesktopMaxRuntime() *int32 {
	return s.DesktopMaxRuntime
}

func (s *AddMcpResourceRequest) GetImageId() *string {
	return s.ImageId
}

func (s *AddMcpResourceRequest) GetMcpIdleTimeout() *int32 {
	return s.McpIdleTimeout
}

func (s *AddMcpResourceRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *AddMcpResourceRequest) GetUserIdleTimeout() *int32 {
	return s.UserIdleTimeout
}

func (s *AddMcpResourceRequest) GetUserName() *string {
	return s.UserName
}

func (s *AddMcpResourceRequest) SetApiKeyId(v string) *AddMcpResourceRequest {
	s.ApiKeyId = &v
	return s
}

func (s *AddMcpResourceRequest) SetConcurrency(v string) *AddMcpResourceRequest {
	s.Concurrency = &v
	return s
}

func (s *AddMcpResourceRequest) SetDesktopMaxRuntime(v int32) *AddMcpResourceRequest {
	s.DesktopMaxRuntime = &v
	return s
}

func (s *AddMcpResourceRequest) SetImageId(v string) *AddMcpResourceRequest {
	s.ImageId = &v
	return s
}

func (s *AddMcpResourceRequest) SetMcpIdleTimeout(v int32) *AddMcpResourceRequest {
	s.McpIdleTimeout = &v
	return s
}

func (s *AddMcpResourceRequest) SetResourceType(v string) *AddMcpResourceRequest {
	s.ResourceType = &v
	return s
}

func (s *AddMcpResourceRequest) SetUserIdleTimeout(v int32) *AddMcpResourceRequest {
	s.UserIdleTimeout = &v
	return s
}

func (s *AddMcpResourceRequest) SetUserName(v string) *AddMcpResourceRequest {
	s.UserName = &v
	return s
}

func (s *AddMcpResourceRequest) Validate() error {
	return dara.Validate(s)
}
