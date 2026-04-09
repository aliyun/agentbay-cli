// This file is auto-generated, don't edit it. Thanks.
package client

// {Action}ResponseBody is the response body struct for {Action}
type {Action}ResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	// Data field type depends on actual API response (string or object)
	Data           *string `{DataField}`
}

// {Action}Response is the response struct for {Action}
type {Action}Response struct {
	Headers    map[string]*string
	StatusCode *int32
	Body       *{Action}ResponseBody
}

// GetCode returns the Code value or empty string if nil
func (s *{Action}ResponseBody) GetCode() string {
	if s == nil || s.Code == nil {
		return ""
	}
	return *s.Code
}

// GetData returns the Data value or empty string if nil
func (s *{Action}ResponseBody) GetData() string {
	if s == nil || s.Data == nil {
		return ""
	}
	return *s.Data
}

// GetSuccess returns whether the request was successful
func (s *{Action}ResponseBody) GetSuccess() bool {
	if s == nil || s.Success == nil {
		return false
	}
	return *s.Success
}
