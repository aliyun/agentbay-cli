// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcpConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigType(v string) *GetMcpConfigRequest
	GetConfigType() *string
	SetImageKeyword(v string) *GetMcpConfigRequest
	GetImageKeyword() *string
}

type GetMcpConfigRequest struct {
	ConfigType   *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	ImageKeyword *string `json:"ImageKeyword,omitempty" xml:"ImageKeyword,omitempty"`
}

func (s GetMcpConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMcpConfigRequest) GoString() string {
	return s.String()
}

func (s *GetMcpConfigRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *GetMcpConfigRequest) GetImageKeyword() *string {
	return s.ImageKeyword
}

func (s *GetMcpConfigRequest) SetConfigType(v string) *GetMcpConfigRequest {
	s.ConfigType = &v
	return s
}

func (s *GetMcpConfigRequest) SetImageKeyword(v string) *GetMcpConfigRequest {
	s.ImageKeyword = &v
	return s
}

func (s *GetMcpConfigRequest) Validate() error {
	return dara.Validate(s)
}
