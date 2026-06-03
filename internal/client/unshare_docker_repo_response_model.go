// This file is auto-generated, don't edit it. Thanks.
package client

// UnshareDockerRepoResponseBodyData is the data field of UnshareDockerRepo response
type UnshareDockerRepoResponseBodyData struct {
	Revoked *bool `json:"Revoked,omitempty" xml:"Revoked,omitempty"`
}

// UnshareDockerRepoResponseBody is the response body struct for UnshareDockerRepo
type UnshareDockerRepoResponseBody struct {
	Code           *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	HttpStatusCode *int32                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
	Data           *UnshareDockerRepoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty"`
}

// UnshareDockerRepoResponse is the response struct for UnshareDockerRepo
type UnshareDockerRepoResponse struct {
	Headers    map[string]*string
	StatusCode *int32
	Body       *UnshareDockerRepoResponseBody
}

// GetCode returns the Code value or empty string if nil
func (s *UnshareDockerRepoResponseBody) GetCode() string {
	if s == nil || s.Code == nil {
		return ""
	}
	return *s.Code
}

// GetMessage returns the Message value or empty string if nil
func (s *UnshareDockerRepoResponseBody) GetMessage() string {
	if s == nil || s.Message == nil {
		return ""
	}
	return *s.Message
}

// GetRequestId returns the RequestId value or empty string if nil
func (s *UnshareDockerRepoResponseBody) GetRequestId() string {
	if s == nil || s.RequestId == nil {
		return ""
	}
	return *s.RequestId
}
