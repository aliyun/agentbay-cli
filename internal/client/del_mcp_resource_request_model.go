// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDelMcpResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v string) *DelMcpResourceRequest
	GetApiKeyId() *string
}

type DelMcpResourceRequest struct {
	ApiKeyId *string `json:"ApiKeyId,omitempty" xml:"ApiKeyId,omitempty"`
}

func (s DelMcpResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DelMcpResourceRequest) GoString() string {
	return s.String()
}

func (s *DelMcpResourceRequest) GetApiKeyId() *string {
	return s.ApiKeyId
}

func (s *DelMcpResourceRequest) SetApiKeyId(v string) *DelMcpResourceRequest {
	s.ApiKeyId = &v
	return s
}

func (s *DelMcpResourceRequest) Validate() error {
	return dara.Validate(s)
}
