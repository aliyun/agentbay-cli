// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/agentbay/agentbay-cli/internal/client"
)

// mockImageListClient implements agentbay.Client for testing image list functions
type mockImageListClient struct {
	userImages   []*client.ListMcpImagesResponseBodyData
	systemImages []*client.ListMcpImagesResponseBodyData
	userError    error
	systemError  error
	userTotal    int32
	systemTotal  int32
}

func (m *mockImageListClient) ListMcpImages(ctx context.Context, req *client.ListMcpImagesRequest) (*client.ListMcpImagesResponse, error) {
	if req.ImageType == nil {
		return nil, fmt.Errorf("ImageType is required")
	}

	imageType := *req.ImageType
	if imageType == "User" {
		if m.userError != nil {
			return nil, m.userError
		}
		return &client.ListMcpImagesResponse{
			Body: &client.ListMcpImagesResponseBody{
				Data:       m.userImages,
				TotalCount: &m.userTotal,
				Success:    boolPtr(true),
			},
		}, nil
	} else if imageType == "System" {
		if m.systemError != nil {
			return nil, m.systemError
		}
		return &client.ListMcpImagesResponse{
			Body: &client.ListMcpImagesResponseBody{
				Data:       m.systemImages,
				TotalCount: &m.systemTotal,
				Success:    boolPtr(true),
			},
		}, nil
	}

	return nil, fmt.Errorf("unknown image type: %s", imageType)
}

// Implement other required methods (stubs for testing)
func (m *mockImageListClient) GetDockerFileStoreCredential(ctx context.Context, request *client.GetDockerFileStoreCredentialRequest) (*client.GetDockerFileStoreCredentialResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (m *mockImageListClient) CreateDockerImageTask(ctx context.Context, request *client.CreateDockerImageTaskRequest) (*client.CreateDockerImageTaskResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (m *mockImageListClient) GetDockerImageTask(ctx context.Context, request *client.GetDockerImageTaskRequest) (*client.GetDockerImageTaskResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (m *mockImageListClient) GetMcpImageInfo(ctx context.Context, request *client.GetMcpImageInfoRequest) (*client.GetMcpImageInfoResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (m *mockImageListClient) CreateResourceGroup(ctx context.Context, request *client.CreateResourceGroupRequest) (*client.CreateResourceGroupResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (m *mockImageListClient) DeleteResourceGroup(ctx context.Context, request *client.DeleteResourceGroupRequest) (*client.DeleteResourceGroupResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

// Helper functions for testing
func boolPtr(b bool) *bool {
	return &b
}

func stringPtr(s string) *string {
	return &s
}

// createMockImage creates a mock image data
func createMockImage(imageId, imageName, imageType, status string) *client.ListMcpImagesResponseBodyData {
	buildType := imageType
	return &client.ListMcpImagesResponseBodyData{
		ImageId:             stringPtr(imageId),
		ImageName:           stringPtr(imageName),
		ImageBuildType:      stringPtr(buildType),
		ImageResourceStatus: stringPtr(status),
		ImageInfo: &client.ListMcpImagesResponseBodyDataImageInfo{
			OsName:    stringPtr("Linux"),
			OsVersion: stringPtr("Debian 12"),
		},
		ImageApplyScene: stringPtr("CodeSpace"),
	}
}

func TestRunImageListWithBothTypes(t *testing.T) {
	ctx := context.Background()

	t.Run("should successfully list both user and system images", func(t *testing.T) {
		// Capture stdout
		var buf bytes.Buffer
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		mockClient := &mockImageListClient{
			userImages: []*client.ListMcpImagesResponseBodyData{
				createMockImage("imgc-1234567890", "my-user-image", "User", "IMAGE_AVAILABLE"),
				createMockImage("imgc-0987654321", "another-user-image", "User", "RESOURCE_PUBLISHED"),
			},
			systemImages: []*client.ListMcpImagesResponseBodyData{
				createMockImage("code-space-debian-12", "Debian 12", "System", "IMAGE_AVAILABLE"),
				createMockImage("code-space-ubuntu-22", "Ubuntu 22.04", "System", "IMAGE_AVAILABLE"),
			},
			userTotal:   2,
			systemTotal: 2,
		}

		err := runImageListWithBothTypes(ctx, mockClient, "", 1, 10)

		// Restore stdout
		w.Close()
		os.Stdout = oldStdout
		_, _ = buf.ReadFrom(r)

		require.NoError(t, err)
	})

	t.Run("should handle empty results", func(t *testing.T) {
		// Capture stdout
		var buf bytes.Buffer
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		mockClient := &mockImageListClient{
			userImages:   []*client.ListMcpImagesResponseBodyData{},
			systemImages: []*client.ListMcpImagesResponseBodyData{},
			userTotal:    0,
			systemTotal:  0,
		}

		err := runImageListWithBothTypes(ctx, mockClient, "", 1, 10)

		// Restore stdout
		w.Close()
		os.Stdout = oldStdout
		_, _ = buf.ReadFrom(r)

		require.NoError(t, err)
	})

	t.Run("should handle user images error gracefully", func(t *testing.T) {
		// Capture stdout
		var buf bytes.Buffer
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		mockClient := &mockImageListClient{
			userError:    fmt.Errorf("API error"),
			systemImages: []*client.ListMcpImagesResponseBodyData{},
			systemTotal:  0,
		}

		err := runImageListWithBothTypes(ctx, mockClient, "", 1, 10)

		// Restore stdout
		w.Close()
		os.Stdout = oldStdout
		_, _ = buf.ReadFrom(r)

		require.Error(t, err)
		assert.Contains(t, err.Error(), "failed to get user images")
	})

	t.Run("should continue with user images when system images fail", func(t *testing.T) {
		// Capture stdout
		var buf bytes.Buffer
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		mockClient := &mockImageListClient{
			userImages: []*client.ListMcpImagesResponseBodyData{
				createMockImage("imgc-1234567890", "my-user-image", "User", "IMAGE_AVAILABLE"),
			},
			systemError:  fmt.Errorf("system API error"),
			userTotal:    1,
			systemTotal:  0,
		}

		err := runImageListWithBothTypes(ctx, mockClient, "", 1, 10)

		// Restore stdout
		w.Close()
		os.Stdout = oldStdout
		io.Copy(&buf, r)
		outputStr := buf.String()

		// Should not return error, just show warning
		require.NoError(t, err)
		assert.Contains(t, outputStr, "WARN", "should show warning when system images fail")
	})

	t.Run("should correctly separate user and system images", func(t *testing.T) {
		// Capture stdout
		var buf bytes.Buffer
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		mockClient := &mockImageListClient{
			userImages: []*client.ListMcpImagesResponseBodyData{
				createMockImage("imgc-1234567890", "user-image-1", "User", "IMAGE_AVAILABLE"),
				createMockImage("imgc-0987654321", "user-image-2", "User", "IMAGE_AVAILABLE"),
			},
			systemImages: []*client.ListMcpImagesResponseBodyData{
				createMockImage("code-space-debian-12", "Debian 12", "System", "IMAGE_AVAILABLE"),
			},
			userTotal:   2,
			systemTotal: 1,
		}

		err := runImageListWithBothTypes(ctx, mockClient, "", 1, 10)

		// Restore stdout
		w.Close()
		os.Stdout = oldStdout
		io.Copy(&buf, r)
		outputStr := buf.String()

		require.NoError(t, err)
		assert.Contains(t, outputStr, "USER IMAGES", "should show user images section")
		assert.Contains(t, outputStr, "SYSTEM IMAGES", "should show system images section")
		assert.Contains(t, outputStr, "imgc-1234567890", "should contain user image")
		assert.Contains(t, outputStr, "code-space-debian-12", "should contain system image")
	})

	t.Run("should handle OS type filter", func(t *testing.T) {
		// Capture stdout
		var buf bytes.Buffer
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		mockClient := &mockImageListClient{
			userImages: []*client.ListMcpImagesResponseBodyData{
				createMockImage("imgc-1234567890", "linux-image", "User", "IMAGE_AVAILABLE"),
			},
			systemImages: []*client.ListMcpImagesResponseBodyData{
				createMockImage("code-space-debian-12", "Debian 12", "System", "IMAGE_AVAILABLE"),
			},
			userTotal:   1,
			systemTotal: 1,
		}

		err := runImageListWithBothTypes(ctx, mockClient, "Linux", 1, 10)

		// Restore stdout
		w.Close()
		os.Stdout = oldStdout
		_, _ = buf.ReadFrom(r)

		require.NoError(t, err)
	})
}

func TestPrintImageTable(t *testing.T) {
	t.Run("should print table header correctly", func(t *testing.T) {
		// Capture stdout
		var buf bytes.Buffer
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		images := []*client.ListMcpImagesResponseBodyData{
			createMockImage("imgc-1234567890", "test-image", "User", "IMAGE_AVAILABLE"),
		}

		printImageTable(images)

		// Restore stdout and read output
		w.Close()
		os.Stdout = oldStdout
		io.Copy(&buf, r)
		outputStr := buf.String()

		// Check that header is present
		assert.Contains(t, outputStr, "IMAGE ID")
		assert.Contains(t, outputStr, "IMAGE NAME")
		assert.Contains(t, outputStr, "TYPE")
		assert.Contains(t, outputStr, "STATUS")
		assert.Contains(t, outputStr, "OS")
		assert.Contains(t, outputStr, "APPLY SCENE")
	})

	t.Run("should handle empty image list", func(t *testing.T) {
		// Capture stdout
		var buf bytes.Buffer
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		images := []*client.ListMcpImagesResponseBodyData{}

		printImageTable(images)

		// Restore stdout
		w.Close()
		os.Stdout = oldStdout
		_, _ = buf.ReadFrom(r)

		// Should not crash, just print header
	})

	t.Run("should handle nil images", func(t *testing.T) {
		// Capture stdout
		var buf bytes.Buffer
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		images := []*client.ListMcpImagesResponseBodyData{
			nil,
			createMockImage("imgc-1234567890", "test-image", "User", "IMAGE_AVAILABLE"),
			nil,
		}

		printImageTable(images)

		// Restore stdout
		w.Close()
		os.Stdout = oldStdout
		io.Copy(&buf, r)
		outputStr := buf.String()

		// Should skip nil images and only print valid ones
		assert.Contains(t, outputStr, "imgc-1234567890", "should contain valid image")
		assert.NotContains(t, outputStr, "<nil>", "should not contain nil")
	})

	t.Run("should format image data correctly", func(t *testing.T) {
		// Capture stdout
		var buf bytes.Buffer
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		images := []*client.ListMcpImagesResponseBodyData{
			createMockImage("imgc-1234567890", "my-custom-image", "User", "IMAGE_AVAILABLE"),
		}

		printImageTable(images)

		// Restore stdout and read output
		w.Close()
		os.Stdout = oldStdout
		io.Copy(&buf, r)
		outputStr := buf.String()

		// Check that image data is present
		assert.Contains(t, outputStr, "imgc-1234567890")
		assert.Contains(t, outputStr, "my-custom-image")
	})

	t.Run("should correctly identify user images by imgc- prefix", func(t *testing.T) {
		ctx := context.Background()
		// Capture stdout
		var buf bytes.Buffer
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		// Test with mixed images - some with imgc- prefix (user), some without (system)
		userImage := createMockImage("imgc-1234567890", "user-image", "User", "IMAGE_AVAILABLE")
		systemImage := createMockImage("code-space-debian-12", "system-image", "System", "IMAGE_AVAILABLE")

		mockClient := &mockImageListClient{
			userImages: []*client.ListMcpImagesResponseBodyData{userImage},
			systemImages: []*client.ListMcpImagesResponseBodyData{systemImage},
			userTotal:    1,
			systemTotal:  1,
		}

		err := runImageListWithBothTypes(ctx, mockClient, "", 1, 10)

		// Restore stdout
		w.Close()
		os.Stdout = oldStdout
		io.Copy(&buf, r)
		outputStr := buf.String()

		require.NoError(t, err)
		// Check that images are separated correctly
		userSectionIndex := strings.Index(outputStr, "USER IMAGES")
		systemSectionIndex := strings.Index(outputStr, "SYSTEM IMAGES")
		assert.Greater(t, userSectionIndex, -1, "should have USER IMAGES section")
		assert.Greater(t, systemSectionIndex, -1, "should have SYSTEM IMAGES section")
		assert.Greater(t, systemSectionIndex, userSectionIndex, "system section should come after user section")
	})
}

