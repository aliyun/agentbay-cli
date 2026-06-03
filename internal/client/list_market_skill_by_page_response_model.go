// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

// ListMarketSkillByPageResponse is the top-level response for ListMarketSkillByPage.
type ListMarketSkillByPageResponse struct {
	Headers    map[string]*string
	StatusCode *int32
	Body       *ListMarketSkillByPageResponseBody
	RawBody    string
}

// ListMarketSkillByPageResponseBody is the body of the ListMarketSkillByPage response.
type ListMarketSkillByPageResponseBody struct {
	Code           *string
	Data           *ListMarketSkillByPageResponseBodyData
	HttpStatusCode *int32
	Message        *string
	RequestId      *string
	Success        *bool
}

// ListMarketSkillByPageResponseBodyData holds the paginated result data.
type ListMarketSkillByPageResponseBodyData struct {
	TotalCount *int32
	TotalPage  *int32
	PageSize   *int32
	PageNumber *int32
	Result     []*ListMarketSkillByPageResponseBodyDataResult
}

// ListMarketSkillByPageResponseBodyDataResult is a single skill entry in the result list.
type ListMarketSkillByPageResponseBodyDataResult struct {
	SkillName   *string
	SkillId     *string
	TenantTags  []string
	SkillStatus *string
	GmtModified *string
	GmtCreate   *string
	Description *string
	Icon        *string
}

// --- nil-safe getters for ListMarketSkillByPageResponseBody ---

// GetCode returns Code safely.
func (b *ListMarketSkillByPageResponseBody) GetCode() *string {
	if b == nil {
		return nil
	}
	return b.Code
}

// GetData returns Data safely.
func (b *ListMarketSkillByPageResponseBody) GetData() *ListMarketSkillByPageResponseBodyData {
	if b == nil {
		return nil
	}
	return b.Data
}

// GetHttpStatusCode returns HttpStatusCode safely.
func (b *ListMarketSkillByPageResponseBody) GetHttpStatusCode() *int32 {
	if b == nil {
		return nil
	}
	return b.HttpStatusCode
}

// GetMessage returns Message safely.
func (b *ListMarketSkillByPageResponseBody) GetMessage() *string {
	if b == nil {
		return nil
	}
	return b.Message
}

// GetRequestId returns RequestId safely.
func (b *ListMarketSkillByPageResponseBody) GetRequestId() *string {
	if b == nil {
		return nil
	}
	return b.RequestId
}

// GetSuccess returns Success safely.
func (b *ListMarketSkillByPageResponseBody) GetSuccess() *bool {
	if b == nil {
		return nil
	}
	return b.Success
}

// --- nil-safe getters for ListMarketSkillByPageResponseBodyData ---

// GetTotalCount returns TotalCount safely.
func (d *ListMarketSkillByPageResponseBodyData) GetTotalCount() *int32 {
	if d == nil {
		return nil
	}
	return d.TotalCount
}

// GetTotalPage returns TotalPage safely.
func (d *ListMarketSkillByPageResponseBodyData) GetTotalPage() *int32 {
	if d == nil {
		return nil
	}
	return d.TotalPage
}

// GetPageSize returns PageSize safely.
func (d *ListMarketSkillByPageResponseBodyData) GetPageSize() *int32 {
	if d == nil {
		return nil
	}
	return d.PageSize
}

// GetPageNumber returns PageNumber safely.
func (d *ListMarketSkillByPageResponseBodyData) GetPageNumber() *int32 {
	if d == nil {
		return nil
	}
	return d.PageNumber
}

// GetResult returns Result safely.
func (d *ListMarketSkillByPageResponseBodyData) GetResult() []*ListMarketSkillByPageResponseBodyDataResult {
	if d == nil {
		return nil
	}
	return d.Result
}

// --- nil-safe getters for ListMarketSkillByPageResponseBodyDataResult ---

// GetSkillName returns SkillName safely.
func (r *ListMarketSkillByPageResponseBodyDataResult) GetSkillName() *string {
	if r == nil {
		return nil
	}
	return r.SkillName
}

// GetSkillId returns SkillId safely.
func (r *ListMarketSkillByPageResponseBodyDataResult) GetSkillId() *string {
	if r == nil {
		return nil
	}
	return r.SkillId
}

// GetSkillStatus returns SkillStatus safely.
func (r *ListMarketSkillByPageResponseBodyDataResult) GetSkillStatus() *string {
	if r == nil {
		return nil
	}
	return r.SkillStatus
}

// GetGmtModified returns GmtModified safely.
func (r *ListMarketSkillByPageResponseBodyDataResult) GetGmtModified() *string {
	if r == nil {
		return nil
	}
	return r.GmtModified
}
