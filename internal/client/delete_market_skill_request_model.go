// This file is auto-generated, don't edit it. Thanks.
package client

import "errors"

// DeleteMarketSkillRequest is the request struct for DeleteMarketSkill
type DeleteMarketSkillRequest struct {
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
}

// Validate validates the DeleteMarketSkillRequest
func (s *DeleteMarketSkillRequest) Validate() error {
	if s.SkillId == nil || *s.SkillId == "" {
		return errors.New("SkillId is required")
	}
	return nil
}

// GetSkillId returns the SkillId value or empty string if nil
func (s *DeleteMarketSkillRequest) GetSkillId() string {
	if s == nil || s.SkillId == nil {
		return ""
	}
	return *s.SkillId
}

// SetSkillId sets the SkillId value
func (s *DeleteMarketSkillRequest) SetSkillId(v string) *DeleteMarketSkillRequest {
	s.SkillId = &v
	return s
}
