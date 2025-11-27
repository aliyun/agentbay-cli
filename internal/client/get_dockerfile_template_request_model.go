// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDockerfileTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplate(v string) *GetDockerfileTemplateRequest
	GetTemplate() *string
}

type GetDockerfileTemplateRequest struct {
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s GetDockerfileTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDockerfileTemplateRequest) GoString() string {
	return s.String()
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

