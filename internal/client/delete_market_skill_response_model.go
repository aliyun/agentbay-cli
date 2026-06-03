// This file is auto-generated, don't edit it. Thanks.
package client

// DeleteMarketSkillResponseBody is the response body struct for DeleteMarketSkill
type DeleteMarketSkillResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

// DeleteMarketSkillResponse is the response struct for DeleteMarketSkill
type DeleteMarketSkillResponse struct {
	Headers    map[string]*string
	StatusCode *int32
	Body       *DeleteMarketSkillResponseBody
}

// GetCode returns the Code value or empty string if nil
func (s *DeleteMarketSkillResponseBody) GetCode() string {
	if s == nil || s.Code == nil {
		return ""
	}
	return *s.Code
}

// GetRequestId returns the RequestId value or empty string if nil
func (s *DeleteMarketSkillResponseBody) GetRequestId() string {
	if s == nil || s.RequestId == nil {
		return ""
	}
	return *s.RequestId
}

// GetSuccess returns the Success value or false if nil
func (s *DeleteMarketSkillResponseBody) GetSuccess() bool {
	if s == nil || s.Success == nil {
		return false
	}
	return *s.Success
}

// GetMessage returns the Message value or empty string if nil
func (s *DeleteMarketSkillResponseBody) GetMessage() string {
	if s == nil || s.Message == nil {
		return ""
	}
	return *s.Message
}
