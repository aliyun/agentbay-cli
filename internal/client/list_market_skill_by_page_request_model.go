// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

// ListMarketSkillByPageRequest holds the parameters for listing market skills by page.
type ListMarketSkillByPageRequest struct {
	// PageNo is the page number (default: 1).
	PageNo *int32
	// PageSize is the number of results per page (default: 10).
	PageSize *int32
	// SkillName filters by skill name (optional).
	SkillName *string
	// TagList filters by tag names (optional). Will be serialized as JSON array string.
	TagList []string
}

// GetPageNo returns the PageNo value safely.
func (r *ListMarketSkillByPageRequest) GetPageNo() *int32 {
	if r == nil {
		return nil
	}
	return r.PageNo
}

// GetPageSize returns the PageSize value safely.
func (r *ListMarketSkillByPageRequest) GetPageSize() *int32 {
	if r == nil {
		return nil
	}
	return r.PageSize
}

// GetSkillName returns the SkillName value safely.
func (r *ListMarketSkillByPageRequest) GetSkillName() *string {
	if r == nil {
		return nil
	}
	return r.SkillName
}

// GetTagList returns the TagList value safely.
func (r *ListMarketSkillByPageRequest) GetTagList() []string {
	if r == nil {
		return nil
	}
	return r.TagList
}

// Validate validates the request parameters.
func (r *ListMarketSkillByPageRequest) Validate() error {
	// No required fields
	return nil
}
