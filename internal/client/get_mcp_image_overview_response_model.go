// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcpImageOverviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMcpImageOverviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMcpImageOverviewResponse
	GetStatusCode() *int32
	SetBody(v *GetMcpImageOverviewResponseBody) *GetMcpImageOverviewResponse
	GetBody() *GetMcpImageOverviewResponseBody
}

type GetMcpImageOverviewResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMcpImageOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMcpImageOverviewResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMcpImageOverviewResponse) GoString() string {
	return s.String()
}

func (s *GetMcpImageOverviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMcpImageOverviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMcpImageOverviewResponse) GetBody() *GetMcpImageOverviewResponseBody {
	return s.Body
}

func (s *GetMcpImageOverviewResponse) SetHeaders(v map[string]*string) *GetMcpImageOverviewResponse {
	s.Headers = v
	return s
}

func (s *GetMcpImageOverviewResponse) SetStatusCode(v int32) *GetMcpImageOverviewResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMcpImageOverviewResponse) SetBody(v *GetMcpImageOverviewResponseBody) *GetMcpImageOverviewResponse {
	s.Body = v
	return s
}

func (s *GetMcpImageOverviewResponse) Validate() error {
	return dara.Validate(s)
}
