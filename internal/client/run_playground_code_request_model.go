// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunPlaygroundCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCodeContent(v string) *RunPlaygroundCodeRequest
	GetCodeContent() *string
	SetSceneType(v string) *RunPlaygroundCodeRequest
	GetSceneType() *string
}

type RunPlaygroundCodeRequest struct {
	CodeContent *string `json:"CodeContent,omitempty" xml:"CodeContent,omitempty"`
	SceneType   *string `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
}

func (s RunPlaygroundCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s RunPlaygroundCodeRequest) GoString() string {
	return s.String()
}

func (s *RunPlaygroundCodeRequest) GetCodeContent() *string {
	return s.CodeContent
}

func (s *RunPlaygroundCodeRequest) GetSceneType() *string {
	return s.SceneType
}

func (s *RunPlaygroundCodeRequest) SetCodeContent(v string) *RunPlaygroundCodeRequest {
	s.CodeContent = &v
	return s
}

func (s *RunPlaygroundCodeRequest) SetSceneType(v string) *RunPlaygroundCodeRequest {
	s.SceneType = &v
	return s
}

func (s *RunPlaygroundCodeRequest) Validate() error {
	return dara.Validate(s)
}
