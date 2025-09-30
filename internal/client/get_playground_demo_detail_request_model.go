// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPlaygroundDemoDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExampleId(v string) *GetPlaygroundDemoDetailRequest
	GetExampleId() *string
	SetSceneType(v string) *GetPlaygroundDemoDetailRequest
	GetSceneType() *string
	SetTemplateId(v string) *GetPlaygroundDemoDetailRequest
	GetTemplateId() *string
}

type GetPlaygroundDemoDetailRequest struct {
	ExampleId  *string `json:"ExampleId,omitempty" xml:"ExampleId,omitempty"`
	SceneType  *string `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetPlaygroundDemoDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPlaygroundDemoDetailRequest) GoString() string {
	return s.String()
}

func (s *GetPlaygroundDemoDetailRequest) GetExampleId() *string {
	return s.ExampleId
}

func (s *GetPlaygroundDemoDetailRequest) GetSceneType() *string {
	return s.SceneType
}

func (s *GetPlaygroundDemoDetailRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetPlaygroundDemoDetailRequest) SetExampleId(v string) *GetPlaygroundDemoDetailRequest {
	s.ExampleId = &v
	return s
}

func (s *GetPlaygroundDemoDetailRequest) SetSceneType(v string) *GetPlaygroundDemoDetailRequest {
	s.SceneType = &v
	return s
}

func (s *GetPlaygroundDemoDetailRequest) SetTemplateId(v string) *GetPlaygroundDemoDetailRequest {
	s.TemplateId = &v
	return s
}

func (s *GetPlaygroundDemoDetailRequest) Validate() error {
	return dara.Validate(s)
}
