// This file is auto-generated, don't edit it. Thanks.
package client

// ShareDockerRepoResponseBodyData is the data field of ShareDockerRepo response
type ShareDockerRepoResponseBodyData struct {
	TargetAliUid *int64  `json:"TargetAliUid,omitempty" xml:"TargetAliUid,omitempty"`
	OwnerAliUid  *int64  `json:"OwnerAliUid,omitempty" xml:"OwnerAliUid,omitempty"`
	AcrRepoName  *string `json:"AcrRepoName,omitempty" xml:"AcrRepoName,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

// ShareDockerRepoResponseBody is the response body struct for ShareDockerRepo
type ShareDockerRepoResponseBody struct {
	Code           *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	HttpStatusCode *int32                           `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
	Data           *ShareDockerRepoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty"`
}

// ShareDockerRepoResponse is the response struct for ShareDockerRepo
type ShareDockerRepoResponse struct {
	Headers    map[string]*string
	StatusCode *int32
	Body       *ShareDockerRepoResponseBody
}

// GetCode returns the Code value or empty string if nil
func (s *ShareDockerRepoResponseBody) GetCode() string {
	if s == nil || s.Code == nil {
		return ""
	}
	return *s.Code
}

// GetMessage returns the Message value or empty string if nil
func (s *ShareDockerRepoResponseBody) GetMessage() string {
	if s == nil || s.Message == nil {
		return ""
	}
	return *s.Message
}

// GetRequestId returns the RequestId value or empty string if nil
func (s *ShareDockerRepoResponseBody) GetRequestId() string {
	if s == nil || s.RequestId == nil {
		return ""
	}
	return *s.RequestId
}
