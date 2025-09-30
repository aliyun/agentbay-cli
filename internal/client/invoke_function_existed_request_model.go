// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeFunctionExistedRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InvokeFunctionExistedRequest
	GetCode() *string
	SetSceneType(v string) *InvokeFunctionExistedRequest
	GetSceneType() *string
}

type InvokeFunctionExistedRequest struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	SceneType *string `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
}

func (s InvokeFunctionExistedRequest) String() string {
	return dara.Prettify(s)
}

func (s InvokeFunctionExistedRequest) GoString() string {
	return s.String()
}

func (s *InvokeFunctionExistedRequest) GetCode() *string {
	return s.Code
}

func (s *InvokeFunctionExistedRequest) GetSceneType() *string {
	return s.SceneType
}

func (s *InvokeFunctionExistedRequest) SetCode(v string) *InvokeFunctionExistedRequest {
	s.Code = &v
	return s
}

func (s *InvokeFunctionExistedRequest) SetSceneType(v string) *InvokeFunctionExistedRequest {
	s.SceneType = &v
	return s
}

func (s *InvokeFunctionExistedRequest) Validate() error {
	return dara.Validate(s)
}
