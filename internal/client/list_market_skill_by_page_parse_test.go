// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"testing"
)

func TestParseListMarketSkillByPageResponse_JSONNumberFields(t *testing.T) {
	// JSON with numeric fields as numbers
	body := `{"code":"200","data":{"RequestId":"A4E9C0A5-7BD3-1B1C-A3C5-D54F9472F3AE","HttpStatusCode":200,"Data":{"TotalCount":2,"TotalPage":1,"PageSize":10,"PageNumber":1,"Result":[{"SkillName":"lxy-find-skills","SkillId":"skill-04p87enx9u4moq5fi","TenantTags":["哈哈","阿里云"],"SkillStatus":"VERIFY_PASSED","GmtModified":"2026-05-26T02:37:59.000+00:00"},{"SkillName":"stock-watcher","SkillId":"skill-04p87lvcjt9o1o9uj","TenantTags":[],"SkillStatus":"INIT","GmtModified":"2026-04-04T08:42:11.000+00:00"}]},"Code":"ok"},"httpStatusCode":"200","requestId":"A4E9C0A5-7BD3-1B1C-A3C5-D54F9472F3AE","successResponse":true}`
	res := map[string]interface{}{
		"body":       body,
		"statusCode": 200,
	}
	resp, err := parseListMarketSkillByPageResponse(res)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Body == nil {
		t.Fatal("body is nil")
	}
	if resp.Body.Data == nil {
		t.Fatal("data is nil")
	}
	if resp.Body.Data.TotalCount == nil || *resp.Body.Data.TotalCount != 2 {
		t.Errorf("expected TotalCount=2, got %v", resp.Body.Data.TotalCount)
	}
	if resp.Body.Data.TotalPage == nil || *resp.Body.Data.TotalPage != 1 {
		t.Errorf("expected TotalPage=1, got %v", resp.Body.Data.TotalPage)
	}
	if resp.Body.Data.PageNumber == nil || *resp.Body.Data.PageNumber != 1 {
		t.Errorf("expected PageNumber=1, got %v", resp.Body.Data.PageNumber)
	}
	if len(resp.Body.Data.Result) != 2 {
		t.Errorf("expected 2 results, got %d", len(resp.Body.Data.Result))
	}
	first := resp.Body.Data.Result[0]
	if first.SkillName == nil || *first.SkillName != "lxy-find-skills" {
		t.Errorf("expected SkillName=lxy-find-skills, got %v", first.SkillName)
	}
	if len(first.TenantTags) != 2 {
		t.Errorf("expected 2 TenantTags, got %d", len(first.TenantTags))
	}
	if resp.Body.Code == nil || *resp.Body.Code != "ok" {
		t.Errorf("expected Code=ok, got %v", resp.Body.Code)
	}
}

func TestParseListMarketSkillByPageResponse_JSONStringFields(t *testing.T) {
	// JSON with numeric fields as strings (some gateways stringify integers)
	body := `{"code":"200","data":{"RequestId":"A4E9C0A5-7BD3-1B1C-A3C5-D54F9472F3AE","HttpStatusCode":"200","Data":{"TotalCount":"6","TotalPage":"1","PageSize":"10","PageNumber":"1","Result":[{"SkillName":"test-skill","SkillId":"skill-abc123","TenantTags":["test"],"SkillStatus":"INIT","GmtModified":"2026-01-01T00:00:00.000+00:00"}]},"Code":"ok"},"httpStatusCode":"200","requestId":"A4E9C0A5-7BD3-1B1C-A3C5-D54F9472F3AE","successResponse":true}`
	res := map[string]interface{}{
		"body":       body,
		"statusCode": 200,
	}
	resp, err := parseListMarketSkillByPageResponse(res)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Body == nil || resp.Body.Data == nil {
		t.Fatal("body or data is nil")
	}
	if resp.Body.Data.TotalCount == nil || *resp.Body.Data.TotalCount != 6 {
		t.Errorf("expected TotalCount=6, got %v", resp.Body.Data.TotalCount)
	}
	if resp.Body.Data.PageSize == nil || *resp.Body.Data.PageSize != 10 {
		t.Errorf("expected PageSize=10, got %v", resp.Body.Data.PageSize)
	}
	if len(resp.Body.Data.Result) != 1 {
		t.Errorf("expected 1 result, got %d", len(resp.Body.Data.Result))
	}
}

func TestParseListMarketSkillByPageResponse_XML(t *testing.T) {
	body := `<?xml version="1.0" encoding="UTF-8"?>
<ListMarketSkillByPageResponse>
  <RequestId>A4E9C0A5-7BD3-1B1C</RequestId>
  <HttpStatusCode>200</HttpStatusCode>
  <Code>ok</Code>
  <Success>true</Success>
  <Message></Message>
  <Data>
    <TotalCount>1</TotalCount>
    <TotalPage>1</TotalPage>
    <PageSize>10</PageSize>
    <PageNumber>1</PageNumber>
    <Result>
      <Item>
        <SkillName>xml-skill</SkillName>
        <SkillId>skill-xmltest</SkillId>
        <SkillStatus>INIT</SkillStatus>
        <GmtModified>2026-01-01T00:00:00.000+00:00</GmtModified>
      </Item>
    </Result>
  </Data>
</ListMarketSkillByPageResponse>`
	res := map[string]interface{}{
		"body":       body,
		"statusCode": 200,
	}
	resp, err := parseListMarketSkillByPageResponse(res)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Body == nil || resp.Body.Data == nil {
		t.Fatal("body or data is nil")
	}
	if resp.Body.Data.TotalCount == nil || *resp.Body.Data.TotalCount != 1 {
		t.Errorf("expected TotalCount=1, got %v", resp.Body.Data.TotalCount)
	}
	if len(resp.Body.Data.Result) != 1 {
		t.Errorf("expected 1 result, got %d", len(resp.Body.Data.Result))
	}
	if resp.Body.Data.Result[0].SkillName == nil || *resp.Body.Data.Result[0].SkillName != "xml-skill" {
		t.Errorf("expected SkillName=xml-skill, got %v", resp.Body.Data.Result[0].SkillName)
	}
}

func TestParseListMarketSkillByPageResponse_FlatFormat(t *testing.T) {
	// Flat format: top-level keys in PascalCase, no outer wrapper.
	// This is the actual API response format observed in production.
	body := `{"HttpStatusCode":200,"Data":{"TotalCount":6,"TotalPage":1,"PageSize":10,"PageNumber":1,"Result":[{"SkillName":"moltbook-hot-posts","SkillId":"skill-04p87enxa6ijpcgql","TenantTags":[],"SkillStatus":"INIT","GmtModified":"2026-05-26T03:28:25.000+00:00","Description":"test desc","Icon":"https://example.com/icon.png"},{"SkillName":"lxy-find-skills","SkillId":"skill-04p87enx9u4moq5fi","TenantTags":["\u54c8\u54c8","\u963f\u91cc\u4e91"],"SkillStatus":"VERIFY_PASSED","GmtModified":"2026-05-26T02:37:59.000+00:00"}]},"RequestId":"2D73B946-83B3-1F0E-A126-10805A8D109B","Code":"ok"}`
	res := map[string]interface{}{
		"body":       body,
		"statusCode": 200,
	}
	resp, err := parseListMarketSkillByPageResponse(res)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Body == nil {
		t.Fatal("body is nil")
	}
	if resp.Body.Code == nil || *resp.Body.Code != "ok" {
		t.Errorf("expected Code=ok, got %v", resp.Body.Code)
	}
	if resp.Body.RequestId == nil || *resp.Body.RequestId != "2D73B946-83B3-1F0E-A126-10805A8D109B" {
		t.Errorf("unexpected RequestId: %v", resp.Body.RequestId)
	}
	if resp.Body.Data == nil {
		t.Fatal("data is nil")
	}
	if resp.Body.Data.TotalCount == nil || *resp.Body.Data.TotalCount != 6 {
		t.Errorf("expected TotalCount=6, got %v", resp.Body.Data.TotalCount)
	}
	if resp.Body.Data.TotalPage == nil || *resp.Body.Data.TotalPage != 1 {
		t.Errorf("expected TotalPage=1, got %v", resp.Body.Data.TotalPage)
	}
	if len(resp.Body.Data.Result) != 2 {
		t.Fatalf("expected 2 results, got %d", len(resp.Body.Data.Result))
	}
	first := resp.Body.Data.Result[0]
	if first.SkillName == nil || *first.SkillName != "moltbook-hot-posts" {
		t.Errorf("expected first SkillName=moltbook-hot-posts, got %v", first.SkillName)
	}
	second := resp.Body.Data.Result[1]
	if len(second.TenantTags) != 2 {
		t.Errorf("expected 2 TenantTags on second skill, got %d", len(second.TenantTags))
	}
}

func TestParseListMarketSkillByPageResponse_EmptyResult(t *testing.T) {
	body := `{"code":"200","data":{"RequestId":"test-id","HttpStatusCode":200,"Data":{"TotalCount":0,"TotalPage":0,"PageSize":10,"PageNumber":1,"Result":[]},"Code":"ok"},"httpStatusCode":"200","requestId":"test-id","successResponse":true}`
	res := map[string]interface{}{
		"body":       body,
		"statusCode": 200,
	}
	resp, err := parseListMarketSkillByPageResponse(res)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Body == nil || resp.Body.Data == nil {
		t.Fatal("body or data is nil")
	}
	if resp.Body.Data.TotalCount == nil || *resp.Body.Data.TotalCount != 0 {
		t.Errorf("expected TotalCount=0, got %v", resp.Body.Data.TotalCount)
	}
	if len(resp.Body.Data.Result) != 0 {
		t.Errorf("expected 0 results, got %d", len(resp.Body.Data.Result))
	}
}
