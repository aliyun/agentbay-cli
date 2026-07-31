// This file is auto-generated, don't edit it. Thanks.
package client

// UpdateImageReserveMinAmountResponseBody is the response body struct for UpdateImageReserveMinAmount
type UpdateImageReserveMinAmountResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

// UpdateImageReserveMinAmountResponse is the response struct for UpdateImageReserveMinAmount
type UpdateImageReserveMinAmountResponse struct {
	Headers    map[string]*string
	StatusCode *int32
	Body       *UpdateImageReserveMinAmountResponseBody
}

// GetCode returns the Code value or empty string if nil
func (s *UpdateImageReserveMinAmountResponseBody) GetCode() string {
	if s == nil || s.Code == nil {
		return ""
	}
	return *s.Code
}

// GetSuccess returns whether the request was successful
func (s *UpdateImageReserveMinAmountResponseBody) GetSuccess() bool {
	if s == nil || s.Success == nil {
		return false
	}
	return *s.Success
}

// GetRequestId returns the RequestId value or empty string if nil
func (s *UpdateImageReserveMinAmountResponseBody) GetRequestId() string {
	if s == nil || s.RequestId == nil {
		return ""
	}
	return *s.RequestId
}
