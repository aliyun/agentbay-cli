// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMcpApiKeyConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v string) *ModifyMcpApiKeyConfigRequest
	GetApiKeyId() *string
	SetConcurrency(v int32) *ModifyMcpApiKeyConfigRequest
	GetConcurrency() *int32
}

type ModifyMcpApiKeyConfigRequest struct {
	ApiKeyId    *string `json:"ApiKeyId,omitempty" xml:"ApiKeyId,omitempty"`
	Concurrency *int32  `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
}

func (s ModifyMcpApiKeyConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyMcpApiKeyConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyMcpApiKeyConfigRequest) GetApiKeyId() *string {
	return s.ApiKeyId
}

func (s *ModifyMcpApiKeyConfigRequest) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *ModifyMcpApiKeyConfigRequest) SetApiKeyId(v string) *ModifyMcpApiKeyConfigRequest {
	s.ApiKeyId = &v
	return s
}

func (s *ModifyMcpApiKeyConfigRequest) SetConcurrency(v int32) *ModifyMcpApiKeyConfigRequest {
	s.Concurrency = &v
	return s
}

func (s *ModifyMcpApiKeyConfigRequest) Validate() error {
	return dara.Validate(s)
}
