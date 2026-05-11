// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client_test

import (
	"testing"

	"github.com/agentbay/agentbay-cli/internal/client"
)

func TestDeleteMcpImageRequest(t *testing.T) {
	t.Run("SetImageId should set ImageId correctly", func(t *testing.T) {
		req := &client.DeleteMcpImageRequest{}
		imageId := "imgc-test123456"

		req.SetImageId(imageId)

		if req.GetImageId() == nil {
			t.Fatal("Expected ImageId to be set, got nil")
		}

		if *req.GetImageId() != imageId {
			t.Errorf("Expected ImageId to be %s, got %s", imageId, *req.GetImageId())
		}
	})

	t.Run("GetImageId should return nil for unset ImageId", func(t *testing.T) {
		req := &client.DeleteMcpImageRequest{}

		if req.GetImageId() != nil {
			t.Errorf("Expected ImageId to be nil, got %v", req.GetImageId())
		}
	})

	t.Run("Validate should not return error for valid request", func(t *testing.T) {
		req := &client.DeleteMcpImageRequest{}
		req.SetImageId("imgc-test123456")

		err := req.Validate()
		if err != nil {
			t.Errorf("Expected no validation error, got %v", err)
		}
	})

	t.Run("String should return non-empty string", func(t *testing.T) {
		req := &client.DeleteMcpImageRequest{}
		req.SetImageId("imgc-test123456")

		s := req.String()
		if s == "" {
			t.Error("Expected non-empty string representation")
		}
	})
}

func TestDeleteMcpImageResponseBody(t *testing.T) {
	t.Run("SetRequestId should set RequestId correctly", func(t *testing.T) {
		body := &client.DeleteMcpImageResponseBody{}
		reqId := "req-abc123"

		body.SetRequestId(reqId)

		if body.GetRequestId() == nil {
			t.Fatal("Expected RequestId to be set, got nil")
		}

		if *body.GetRequestId() != reqId {
			t.Errorf("Expected RequestId to be %s, got %s", reqId, *body.GetRequestId())
		}
	})

	t.Run("SetSuccess should set Success correctly", func(t *testing.T) {
		body := &client.DeleteMcpImageResponseBody{}

		body.SetSuccess(true)

		if body.GetSuccess() == nil {
			t.Fatal("Expected Success to be set, got nil")
		}

		if !*body.GetSuccess() {
			t.Error("Expected Success to be true")
		}
	})

	t.Run("SetCode and SetMessage should work correctly", func(t *testing.T) {
		body := &client.DeleteMcpImageResponseBody{}
		body.SetCode("200")
		body.SetMessage("OK")
		body.SetHttpStatusCode(200)

		if body.GetCode() == nil || *body.GetCode() != "200" {
			t.Errorf("Expected Code to be '200', got %v", body.GetCode())
		}
		if body.GetMessage() == nil || *body.GetMessage() != "OK" {
			t.Errorf("Expected Message to be 'OK', got %v", body.GetMessage())
		}
		if body.GetHttpStatusCode() == nil || *body.GetHttpStatusCode() != 200 {
			t.Errorf("Expected HttpStatusCode to be 200, got %v", body.GetHttpStatusCode())
		}
	})
}

func TestDeleteMcpImageResponse(t *testing.T) {
	t.Run("SetBody should set Body correctly", func(t *testing.T) {
		resp := &client.DeleteMcpImageResponse{}
		body := &client.DeleteMcpImageResponseBody{}
		body.SetSuccess(true)

		resp.SetBody(body)

		if resp.GetBody() == nil {
			t.Fatal("Expected Body to be set, got nil")
		}

		if !*resp.GetBody().GetSuccess() {
			t.Error("Expected Body.Success to be true")
		}
	})

	t.Run("SetStatusCode should set StatusCode correctly", func(t *testing.T) {
		resp := &client.DeleteMcpImageResponse{}
		resp.SetStatusCode(200)

		if resp.GetStatusCode() == nil {
			t.Fatal("Expected StatusCode to be set, got nil")
		}

		if *resp.GetStatusCode() != 200 {
			t.Errorf("Expected StatusCode to be 200, got %d", *resp.GetStatusCode())
		}
	})

	t.Run("SetHeaders should set Headers correctly", func(t *testing.T) {
		resp := &client.DeleteMcpImageResponse{}
		val := "test-value"
		headers := map[string]*string{"X-Test": &val}
		resp.SetHeaders(headers)

		if resp.GetHeaders() == nil {
			t.Fatal("Expected Headers to be set, got nil")
		}

		if *resp.GetHeaders()["X-Test"] != "test-value" {
			t.Errorf("Expected header X-Test to be 'test-value', got %v", resp.GetHeaders()["X-Test"])
		}
	})
}
