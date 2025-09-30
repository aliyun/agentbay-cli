// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMcpPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMcpPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMcpPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMcpPoliciesResponseBody) *DescribeMcpPoliciesResponse
	GetBody() *DescribeMcpPoliciesResponseBody
}

type DescribeMcpPoliciesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMcpPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMcpPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMcpPoliciesResponse) GoString() string {
	return s.String()
}

func (s *DescribeMcpPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMcpPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMcpPoliciesResponse) GetBody() *DescribeMcpPoliciesResponseBody {
	return s.Body
}

func (s *DescribeMcpPoliciesResponse) SetHeaders(v map[string]*string) *DescribeMcpPoliciesResponse {
	s.Headers = v
	return s
}

func (s *DescribeMcpPoliciesResponse) SetStatusCode(v int32) *DescribeMcpPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMcpPoliciesResponse) SetBody(v *DescribeMcpPoliciesResponseBody) *DescribeMcpPoliciesResponse {
	s.Body = v
	return s
}

func (s *DescribeMcpPoliciesResponse) Validate() error {
	return dara.Validate(s)
}
