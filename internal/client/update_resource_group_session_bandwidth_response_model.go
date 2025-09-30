// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceGroupSessionBandwidthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateResourceGroupSessionBandwidthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateResourceGroupSessionBandwidthResponse
	GetStatusCode() *int32
	SetBody(v *UpdateResourceGroupSessionBandwidthResponseBody) *UpdateResourceGroupSessionBandwidthResponse
	GetBody() *UpdateResourceGroupSessionBandwidthResponseBody
}

type UpdateResourceGroupSessionBandwidthResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResourceGroupSessionBandwidthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResourceGroupSessionBandwidthResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceGroupSessionBandwidthResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceGroupSessionBandwidthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateResourceGroupSessionBandwidthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateResourceGroupSessionBandwidthResponse) GetBody() *UpdateResourceGroupSessionBandwidthResponseBody {
	return s.Body
}

func (s *UpdateResourceGroupSessionBandwidthResponse) SetHeaders(v map[string]*string) *UpdateResourceGroupSessionBandwidthResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceGroupSessionBandwidthResponse) SetStatusCode(v int32) *UpdateResourceGroupSessionBandwidthResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourceGroupSessionBandwidthResponse) SetBody(v *UpdateResourceGroupSessionBandwidthResponseBody) *UpdateResourceGroupSessionBandwidthResponse {
	s.Body = v
	return s
}

func (s *UpdateResourceGroupSessionBandwidthResponse) Validate() error {
	return dara.Validate(s)
}
