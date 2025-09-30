// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcpResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v string) *GetMcpResourceRequest
	GetApiKeyId() *string
}

type GetMcpResourceRequest struct {
	ApiKeyId *string `json:"ApiKeyId,omitempty" xml:"ApiKeyId,omitempty"`
}

func (s GetMcpResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResourceRequest) GoString() string {
	return s.String()
}

func (s *GetMcpResourceRequest) GetApiKeyId() *string {
	return s.ApiKeyId
}

func (s *GetMcpResourceRequest) SetApiKeyId(v string) *GetMcpResourceRequest {
	s.ApiKeyId = &v
	return s
}

func (s *GetMcpResourceRequest) Validate() error {
	return dara.Validate(s)
}
