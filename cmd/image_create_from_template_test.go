package cmd

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/agentbay/agentbay-cli/internal/client"
)

func TestParseSourceImageRef(t *testing.T) {
	tests := []struct {
		name            string
		sourceImage     string
		wantRegistry    string
		wantPhysical    string
		wantAliUID      int64
		wantFullPath    bool
		wantErrContains string
	}{
		{
			name:         "short path",
			sourceImage:  "/customer_cli/1730408327554214:cli-test-0.0.1",
			wantPhysical: "/customer_cli/1730408327554214:cli-test-0.0.1",
			wantAliUID:   1730408327554214,
		},
		{
			name:         "full registry path",
			sourceImage:  "ai-container-registry.cn-hangzhou.cr.aliyuncs.com/customer_cli/1730408327554214:cli-test-0.0.1",
			wantRegistry: "ai-container-registry.cn-hangzhou.cr.aliyuncs.com",
			wantPhysical: "/customer_cli/1730408327554214:cli-test-0.0.1",
			wantAliUID:   1730408327554214,
			wantFullPath: true,
		},
		{
			name:            "missing tag",
			sourceImage:     "/customer_cli/1730408327554214",
			wantErrContains: "must include a tag",
		},
		{
			name:            "non numeric uid",
			sourceImage:     "/customer_cli/not-a-uid:v1",
			wantErrContains: "must be a positive integer",
		},
		{
			name:            "unsupported namespace",
			sourceImage:     "/other_ns/1730408327554214:v1",
			wantErrContains: "unsupported namespace",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ref, err := parseSourceImageRef(tt.sourceImage)
			if tt.wantErrContains != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.wantErrContains)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.wantRegistry, ref.Registry)
			require.Equal(t, tt.wantPhysical, ref.PhysicalImageID)
			require.Equal(t, tt.wantAliUID, ref.RepoAliUID)
			require.Equal(t, tt.wantFullPath, ref.IsFullPath)
		})
	}
}

func TestAuthorizeSourceImageOwnRepositorySkipsSharedQuery(t *testing.T) {
	ref, err := parseSourceImageRef("/customer_cli/1730408327554214:cli-test-0.0.1")
	require.NoError(t, err)
	cache := &acrCredentialCache{
		RegistryURL: "ai-container-registry.cn-hangzhou.cr.aliyuncs.com",
		Namespace:   "customer_cli",
		RepoName:    "1730408327554214",
	}
	called := false
	auth, err := authorizeSourceImage(context.Background(), ref, cache, func(context.Context, *client.ListSharedDockerReposRequest) (*client.ListSharedDockerReposResponse, error) {
		called = true
		return nil, fmt.Errorf("should not be called")
	})
	require.NoError(t, err)
	require.False(t, called)
	require.True(t, auth.OwnRepository)
	require.Equal(t, "Own repository", auth.SourceType)
	require.Equal(t, "ai-container-registry.cn-hangzhou.cr.aliyuncs.com/customer_cli/1730408327554214:cli-test-0.0.1", auth.DisplaySourceImage)
}

func TestAuthorizeSourceImageSharedRepositoryPrintsRequestID(t *testing.T) {
	ref, err := parseSourceImageRef("/customer_cli/1160165251879674:cli-test-0.0.1")
	require.NoError(t, err)
	cache := &acrCredentialCache{
		RegistryURL: "ai-container-registry.cn-hangzhou.cr.aliyuncs.com",
		Namespace:   "customer_cli",
		RepoName:    "1730408327554214",
	}

	output := captureImageCreateFromTemplateStdout(t, func() {
		auth, err := authorizeSourceImage(context.Background(), ref, cache, func(ctx context.Context, req *client.ListSharedDockerReposRequest) (*client.ListSharedDockerReposResponse, error) {
			require.Equal(t, "Incoming", *req.Direction)
			require.Equal(t, int64(1160165251879674), *req.QueryAliUid)
			require.Equal(t, int32(1), *req.PageStart)
			require.Equal(t, int32(10), *req.PageSize)
			return &client.ListSharedDockerReposResponse{
				Body: &client.ListSharedDockerReposResponseBody{
					Code:      stringPtr("ok"),
					RequestId: stringPtr("req-shared-1"),
					Success:   boolPtr(true),
					Data: []*client.ListSharedDockerReposResponseBodyDataItem{
						{PeerAliUid: int64Pointer(1160165251879674), Status: stringPtr("ACTIVE")},
					},
				},
			}, nil
		})
		require.NoError(t, err)
		require.False(t, auth.OwnRepository)
		require.Equal(t, "Shared repository (owner AliUID: ****9674)", auth.SourceType)
		require.Equal(t, "/customer_cli/1160165251879674:cli-test-0.0.1", auth.DisplaySourceImage)
	})
	require.Contains(t, output, "[INFO] ListSharedDockerRepos Request ID: req-shared-1")
}

func TestAuthorizeSourceImageSharedRepositoryRejectsEmptyData(t *testing.T) {
	ref, err := parseSourceImageRef("/customer_cli/1160165251879674:cli-test-0.0.1")
	require.NoError(t, err)
	cache := &acrCredentialCache{
		RegistryURL: "ai-container-registry.cn-hangzhou.cr.aliyuncs.com",
		Namespace:   "customer_cli",
		RepoName:    "1730408327554214",
	}

	err = func() error {
		_, err := authorizeSourceImage(context.Background(), ref, cache, func(context.Context, *client.ListSharedDockerReposRequest) (*client.ListSharedDockerReposResponse, error) {
			return &client.ListSharedDockerReposResponse{
				Body: &client.ListSharedDockerReposResponseBody{
					Code:      stringPtr("ok"),
					RequestId: stringPtr("req-empty"),
					Success:   boolPtr(true),
					Data:      []*client.ListSharedDockerReposResponseBodyDataItem{},
				},
			}, nil
		})
		return err
	}()
	require.Error(t, err)
	require.Contains(t, err.Error(), "no incoming Docker repo sharing authorization")
}

func TestExtractCreateFromTemplateRequestID(t *testing.T) {
	require.Equal(t, "req-create-1", extractCreateFromTemplateRequestID([]byte(`{"RequestId":"req-create-1","Code":"ok"}`)))
	require.Equal(t, "req-create-2", extractCreateFromTemplateRequestID([]byte(`{"RequestID":"req-create-2","Code":"ok"}`)))
	require.Empty(t, extractCreateFromTemplateRequestID([]byte(`not-json`)))
}

func captureImageCreateFromTemplateStdout(t *testing.T, fn func()) string {
	t.Helper()
	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	require.NoError(t, err)
	os.Stdout = w
	fn()
	require.NoError(t, w.Close())
	os.Stdout = oldStdout
	var buf bytes.Buffer
	_, err = buf.ReadFrom(r)
	require.NoError(t, err)
	return strings.TrimSpace(buf.String())
}

func int64Pointer(v int64) *int64 {
	return &v
}
