// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client_test

import (
	"testing"

	"github.com/agentbay/agentbay-cli/internal/client"
)

func TestDeleteResourceGroupRequest(t *testing.T) {
	t.Run("SetImageId should set ImageId correctly", func(t *testing.T) {
		req := &client.DeleteResourceGroupRequest{}
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
		req := &client.DeleteResourceGroupRequest{}

		if req.GetImageId() != nil {
			t.Errorf("Expected ImageId to be nil, got %v", req.GetImageId())
		}
	})

	t.Run("Validate should not return error for valid request", func(t *testing.T) {
		req := &client.DeleteResourceGroupRequest{}
		req.SetImageId("imgc-test123456")

		err := req.Validate()
		if err != nil {
			t.Errorf("Expected no validation error, got %v", err)
		}
	})
}
