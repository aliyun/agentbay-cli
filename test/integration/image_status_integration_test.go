// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

//go:build integration
// +build integration

package integration_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/spf13/cobra"

	"github.com/agentbay/agentbay-cli/cmd"
	"github.com/agentbay/agentbay-cli/internal/config"
)

func TestImageStatusCommand_NoAuthentication(t *testing.T) {
	tempDir := t.TempDir()
	t.Setenv("AGENTBAY_CLI_CONFIG_DIR", tempDir)
	_ = os.Unsetenv(config.EnvAccessKeyID)
	_ = os.Unsetenv(config.EnvAccessKeySecret)
	_ = os.Unsetenv(config.EnvAccessKeySessionToken)

	rootCmd := &cobra.Command{Use: "agentbay"}
	rootCmd.AddGroup(&cobra.Group{ID: "management", Title: "Management Commands"})
	rootCmd.AddCommand(cmd.ImageCmd)
	rootCmd.SetArgs([]string{"image", "status", "imgc-test"})

	var buf bytes.Buffer
	rootCmd.SetOut(&buf)
	rootCmd.SetErr(&buf)

	err := rootCmd.Execute()
	if err == nil {
		t.Fatal("expected error when not authenticated")
	}
	combined := buf.String()
	if !strings.Contains(strings.ToLower(combined), "not authenticated") && !strings.Contains(strings.ToLower(err.Error()), "not authenticated") {
		t.Fatalf("expected not authenticated in output or error; err=%v out=%q", err, combined)
	}
}

func TestImageStatusCommand_WithAuthenticatedConfig(t *testing.T) {
	if os.Getenv("RUN_INTEGRATION_TESTS") == "" {
		t.Skip("Skipping: set RUN_INTEGRATION_TESTS=1 to run against real API")
	}

	cfg, err := config.GetConfig()
	if err != nil {
		t.Fatalf("load config: %v", err)
	}
	if !cfg.IsAuthenticated() {
		t.Skip("Skipping: need OAuth tokens or AGENTBAY_ACCESS_KEY_* for real API")
	}

	imageID := os.Getenv("AGENTBAY_TEST_IMAGE_ID")
	if imageID == "" {
		imageID = "imgc-0a9mg1h2l5dwec9vs"
	}

	rootCmd := &cobra.Command{Use: "agentbay"}
	rootCmd.AddGroup(&cobra.Group{ID: "management", Title: "Management Commands"})
	rootCmd.AddCommand(cmd.ImageCmd)
	rootCmd.SetArgs([]string{"image", "status", imageID})

	var stdout, stderr bytes.Buffer
	rootCmd.SetOut(&stdout)
	rootCmd.SetErr(&stderr)

	err = rootCmd.Execute()
	out := stdout.String() + stderr.String()
	if err != nil {
		t.Logf("stdout+stderr: %s", out)
		t.Fatalf("image status failed: %v", err)
	}
	if !strings.Contains(out, "[INFO] Resource status (API):") {
		t.Errorf("expected status output; got:\n%s", out)
	}
	if !strings.Contains(out, imageID) {
		t.Errorf("expected image id in output; got:\n%s", out)
	}
}
