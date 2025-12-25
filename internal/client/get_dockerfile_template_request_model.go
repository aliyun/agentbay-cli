// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDockerfileTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSource(v string) *GetDockerfileTemplateRequest
	GetSource() *string
	SetSourceImageId(v string) *GetDockerfileTemplateRequest
	GetSourceImageId() *string
	SetTemplate(v string) *GetDockerfileTemplateRequest
	GetTemplate() *string
}

type GetDockerfileTemplateRequest struct {
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceImageId *string `json:"SourceImageId,omitempty" xml:"SourceImageId,omitempty"`
	Template      *string `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s GetDockerfileTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDockerfileTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetDockerfileTemplateRequest) GetSource() *string {
	return s.Source
}

func (s *GetDockerfileTemplateRequest) SetSource(v string) *GetDockerfileTemplateRequest {
	s.Source = &v
	return s
}

func (s *GetDockerfileTemplateRequest) GetSourceImageId() *string {
	return s.SourceImageId
}

func (s *GetDockerfileTemplateRequest) SetSourceImageId(v string) *GetDockerfileTemplateRequest {
	s.SourceImageId = &v
	return s
}

func (s *GetDockerfileTemplateRequest) GetTemplate() *string {
	return s.Template
}

func (s *GetDockerfileTemplateRequest) SetTemplate(v string) *GetDockerfileTemplateRequest {
	s.Template = &v
	return s
}

func (s *GetDockerfileTemplateRequest) Validate() error {
	return dara.Validate(s)
}
