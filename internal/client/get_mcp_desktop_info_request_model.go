// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcpDesktopInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v string) *GetMcpDesktopInfoRequest
	GetApiKeyId() *string
}

type GetMcpDesktopInfoRequest struct {
	ApiKeyId *string `json:"ApiKeyId,omitempty" xml:"ApiKeyId,omitempty"`
}

func (s GetMcpDesktopInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMcpDesktopInfoRequest) GoString() string {
	return s.String()
}

func (s *GetMcpDesktopInfoRequest) GetApiKeyId() *string {
	return s.ApiKeyId
}

func (s *GetMcpDesktopInfoRequest) SetApiKeyId(v string) *GetMcpDesktopInfoRequest {
	s.ApiKeyId = &v
	return s
}

func (s *GetMcpDesktopInfoRequest) Validate() error {
	return dara.Validate(s)
}
