// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

//go:build integration
// +build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/agentbay/agentbay-cli/internal/agentbay"
	"github.com/agentbay/agentbay-cli/internal/client"
	"github.com/agentbay/agentbay-cli/internal/config"
	"github.com/alibabacloud-go/tea/dara"
	log "github.com/sirupsen/logrus"
)

func TestGetMcpImageInfo(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	cfg, err := config.GetConfig()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	agentbayClient := agentbay.NewClientFromConfig(cfg)

	request := &client.GetMcpImageInfoRequest{}
	request.SetImageId("imgc-0a9mg1h2l5dwec9vs")

	t.Logf("Making GetMcpImageInfo request with ImageId: %s", *request.GetImageId())

	resp, err := agentbayClient.GetMcpImageInfo(context.Background(), request)
	if err != nil {
		t.Fatalf("GetMcpImageInfo failed: %v", err)
	}

	if resp == nil {
		t.Fatal("Response is nil")
	}

	if resp.Body == nil {
		t.Fatal("Response body is nil")
	}

	t.Logf("Response received:")
	t.Logf("  StatusCode: %d", dara.Int32Value(resp.StatusCode))
	t.Logf("  HttpStatusCode: %d", dara.Int32Value(resp.Body.HttpStatusCode))
	t.Logf("  Code: %s", dara.StringValue(resp.Body.Code))
	t.Logf("  Success: %t", dara.BoolValue(resp.Body.Success))
	t.Logf("  Message: %s", dara.StringValue(resp.Body.Message))
	t.Logf("  RequestId: %s", dara.StringValue(resp.Body.RequestId))

	if resp.Body.Data != nil {
		t.Logf("  Data:")
		t.Logf("    ImageId: %s", dara.StringValue(resp.Body.Data.ImageId))
		t.Logf("    ImageName: %s", dara.StringValue(resp.Body.Data.ImageName))
		t.Logf("    ImageBuildType: %s", dara.StringValue(resp.Body.Data.ImageBuildType))
		t.Logf("    ImageApplyScene: %s", dara.StringValue(resp.Body.Data.ImageApplyScene))

		if resp.Body.Data.ImageInfo != nil {
			t.Logf("    ImageInfo:")
			t.Logf("      OsName: %s", dara.StringValue(resp.Body.Data.ImageInfo.OsName))
			t.Logf("      OsVersion: %s", dara.StringValue(resp.Body.Data.ImageInfo.OsVersion))
			t.Logf("      PlatformName: %s", dara.StringValue(resp.Body.Data.ImageInfo.PlatformName))
			t.Logf("      Status: %s", dara.StringValue(resp.Body.Data.ImageInfo.Status))
			t.Logf("      DataDiskSize: %d", dara.Int32Value(resp.Body.Data.ImageInfo.DataDiskSize))
			t.Logf("      SystemDiskSize: %d", dara.Int32Value(resp.Body.Data.ImageInfo.SystemDiskSize))
			t.Logf("      UpdateTime: %s", dara.StringValue(resp.Body.Data.ImageInfo.UpdateTime))
		}

		// Check ImageType from response headers
		if resp.Headers != nil {
			if imageType, ok := resp.Headers["X-Image-Type"]; ok && imageType != nil {
				t.Logf("    ImageType (from header): %s", dara.StringValue(imageType))
			}
		}

		if resp.Body.Data.ImageBuildInfo != nil {
			t.Logf("    ImageBuildInfo:")
			t.Logf("      TaskId: %s", dara.StringValue(resp.Body.Data.ImageBuildInfo.TaskId))
			t.Logf("      VersionId: %s", dara.StringValue(resp.Body.Data.ImageBuildInfo.VersionId))
			t.Logf("      ApiKeyId: %s", dara.StringValue(resp.Body.Data.ImageBuildInfo.ApiKeyId))
			t.Logf("      InstanceReady: %t", dara.BoolValue(resp.Body.Data.ImageBuildInfo.InstanceReady))
			t.Logf("      AndroidMobileGroupId: %s", dara.StringValue(resp.Body.Data.ImageBuildInfo.AndroidMobileGroupId))
			t.Logf("      AndroidMobileInstanceId: %s", dara.StringValue(resp.Body.Data.ImageBuildInfo.AndroidMobileInstanceId))
		}
	}

	if resp.Body.Data == nil {
		t.Fatal("Data is nil")
	}

	// Verify that we got image type
	if resp.Body.Data.ImageBuildType == nil || *resp.Body.Data.ImageBuildType == "" {
		t.Error("ImageBuildType is empty")
	} else {
		t.Logf("[SUCCESS] ImageBuildType obtained: %s", *resp.Body.Data.ImageBuildType)
	}

	// Verify that we got resource status (ImageInfo.Status serves as resource status)
	if resp.Body.Data.ImageInfo != nil && resp.Body.Data.ImageInfo.Status != nil && *resp.Body.Data.ImageInfo.Status != "" {
		t.Logf("[SUCCESS] Image Status obtained: %s", *resp.Body.Data.ImageInfo.Status)
	} else {
		t.Error("Image Status is empty")
	}

	// Note: ImageResourceStatus exists in XML response but not in SDK model
	// It can be seen in debug logs as: IMAGE_AVAILABLE
	t.Logf("[INFO] Check debug logs for ImageResourceStatus from XML response")
}
