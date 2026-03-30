// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/agentbay/agentbay-cli/internal/agentbay"
	"github.com/agentbay/agentbay-cli/internal/client"
)

// mockGetMcpImageInfoClient implements agentbay.Client for GetImageInfo tests.
type mockGetMcpImageInfoClient struct {
	resp *client.GetMcpImageInfoResponse
	err  error
}

func (m *mockGetMcpImageInfoClient) GetMcpImageInfo(ctx context.Context, request *client.GetMcpImageInfoRequest) (*client.GetMcpImageInfoResponse, error) {
	return m.resp, m.err
}

func (m *mockGetMcpImageInfoClient) GetDockerFileStoreCredential(ctx context.Context, request *client.GetDockerFileStoreCredentialRequest) (*client.GetDockerFileStoreCredentialResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (m *mockGetMcpImageInfoClient) CreateDockerImageTask(ctx context.Context, request *client.CreateDockerImageTaskRequest) (*client.CreateDockerImageTaskResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (m *mockGetMcpImageInfoClient) GetDockerImageTask(ctx context.Context, request *client.GetDockerImageTaskRequest) (*client.GetDockerImageTaskResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (m *mockGetMcpImageInfoClient) ListMcpImages(ctx context.Context, request *client.ListMcpImagesRequest) (*client.ListMcpImagesResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (m *mockGetMcpImageInfoClient) CreateResourceGroup(ctx context.Context, request *client.CreateResourceGroupRequest) (*client.CreateResourceGroupResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (m *mockGetMcpImageInfoClient) DeleteResourceGroup(ctx context.Context, request *client.DeleteResourceGroupRequest) (*client.DeleteResourceGroupResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (m *mockGetMcpImageInfoClient) GetDockerfileTemplate(ctx context.Context, request *client.GetDockerfileTemplateRequest) (*client.GetDockerfileTemplateResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (m *mockGetMcpImageInfoClient) GetMarketSkillCredential(ctx context.Context, request *client.GetMarketSkillCredentialRequest) (*client.GetMarketSkillCredentialResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (m *mockGetMcpImageInfoClient) CreateMarketSkill(ctx context.Context, request *client.CreateMarketSkillRequest) (*client.CreateMarketSkillResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (m *mockGetMcpImageInfoClient) DescribeMarketSkillDetail(ctx context.Context, request *client.DescribeMarketSkillDetailRequest) (*client.DescribeMarketSkillDetailResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (m *mockGetMcpImageInfoClient) CreateMarketSkillGroup(ctx context.Context, request *client.CreateMarketSkillGroupRequest) (*client.CreateMarketSkillGroupResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (m *mockGetMcpImageInfoClient) ListMarketGroupSkill(ctx context.Context, request *client.ListMarketGroupSkillRequest) (*client.ListMarketGroupSkillResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (m *mockGetMcpImageInfoClient) AddMarketGroupSkill(ctx context.Context, request *client.AddMarketGroupSkillRequest) (*client.AddMarketGroupSkillResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (m *mockGetMcpImageInfoClient) RemoveMarketGroupSkill(ctx context.Context, request *client.RemoveMarketGroupSkillRequest) (*client.RemoveMarketGroupSkillResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

var _ agentbay.Client = (*mockGetMcpImageInfoClient)(nil)

func TestInferImageTypeFromImageID(t *testing.T) {
	t.Parallel()
	require.Equal(t, "User", inferImageTypeFromImageID("imgc-0a9mg1h2l5dwec9vs"))
	require.Equal(t, "System", inferImageTypeFromImageID("aliyun-mcp-ubuntu"))
}

func TestGetImageInfo_JSONBodyWithoutSyntheticHeaders(t *testing.T) {
	t.Parallel()
	success := true
	resp := &client.GetMcpImageInfoResponse{
		Body: &client.GetMcpImageInfoResponseBody{
			Success: &success,
			Data: &client.GetMcpImageInfoResponseBodyData{
				ImageResourceStatus: stringPtr("IMAGE_AVAILABLE"),
				ImageInfo: &client.GetMcpImageInfoResponseBodyDataImageInfo{
					ImageType: stringPtr("User"),
					Status:    stringPtr("IMAGE_AVAILABLE"),
				},
			},
		},
	}
	m := &mockGetMcpImageInfoClient{resp: resp}
	info, err := GetImageInfo(context.Background(), m, "imgc-testid")
	require.NoError(t, err)
	require.Equal(t, "IMAGE_AVAILABLE", info.ResourceStatus)
	require.Equal(t, "User", info.ImageType)
}

func TestGetImageInfo_InfersUserWhenImageTypeMissing(t *testing.T) {
	t.Parallel()
	success := true
	resp := &client.GetMcpImageInfoResponse{
		Body: &client.GetMcpImageInfoResponseBody{
			Success: &success,
			Data: &client.GetMcpImageInfoResponseBodyData{
				ImageResourceStatus: stringPtr("IMAGE_AVAILABLE"),
				ImageInfo: &client.GetMcpImageInfoResponseBodyDataImageInfo{
					Status: stringPtr("IMAGE_AVAILABLE"),
				},
			},
		},
	}
	m := &mockGetMcpImageInfoClient{resp: resp}
	info, err := GetImageInfo(context.Background(), m, "imgc-abc")
	require.NoError(t, err)
	require.Equal(t, "User", info.ImageType)
}
