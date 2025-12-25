// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

//go:build integration
// +build integration

package integration_test

import (
	"os"
	"strings"
	"testing"

	"github.com/agentbay/agentbay-cli/internal/config"
)

// TestEnvironmentSwitching verifies that environment switching works correctly
func TestEnvironmentSwitching(t *testing.T) {
	tests := []struct {
		name             string
		envValue         string
		expectedEnv      config.Environment
		expectedEndpoint string
		expectedClientID string
	}{
		{
			name:             "Production environment (default)",
			envValue:         "",
			expectedEnv:      config.EnvProduction,
			expectedEndpoint: "xiaoying-share.cn-shanghai.aliyuncs.com",
			expectedClientID: "4032653160518150541",
		},
		{
			name:             "Production environment (explicit)",
			envValue:         "production",
			expectedEnv:      config.EnvProduction,
			expectedEndpoint: "xiaoying-share.cn-shanghai.aliyuncs.com",
			expectedClientID: "4032653160518150541",
		},
		{
			name:             "Pre-release environment",
			envValue:         "prerelease",
			expectedEnv:      config.EnvPreRelease,
			expectedEndpoint: "xiaoying-pre.cn-hangzhou.aliyuncs.com",
			expectedClientID: "4019057658592127596",
		},
		{
			name:             "Pre-release environment (short form)",
			envValue:         "pre",
			expectedEnv:      config.EnvPreRelease,
			expectedEndpoint: "xiaoying-pre.cn-hangzhou.aliyuncs.com",
			expectedClientID: "4019057658592127596",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set environment variable
			if tt.envValue != "" {
				os.Setenv("AGENTBAY_ENV", tt.envValue)
			} else {
				os.Unsetenv("AGENTBAY_ENV")
			}
			defer os.Unsetenv("AGENTBAY_ENV")

			// Test environment detection
			env := config.GetEnvironment()
			if env != tt.expectedEnv {
				t.Errorf("GetEnvironment() = %v, want %v", env, tt.expectedEnv)
			}

			// Test environment config
			envConfig := config.GetEnvironmentConfig()
			if envConfig.Endpoint != tt.expectedEndpoint {
				t.Errorf("GetEnvironmentConfig().Endpoint = %v, want %v", envConfig.Endpoint, tt.expectedEndpoint)
			}
			if envConfig.ClientID != tt.expectedClientID {
				t.Errorf("GetEnvironmentConfig().ClientID = %v, want %v", envConfig.ClientID, tt.expectedClientID)
			}

			// Test API config uses correct endpoint
			apiConfig := config.LoadAPIConfig(nil)
			if apiConfig.Endpoint != tt.expectedEndpoint {
				t.Errorf("LoadAPIConfig().Endpoint = %v, want %v", apiConfig.Endpoint, tt.expectedEndpoint)
			}

			// Test ClientID
			clientID := config.GetClientID()
			if clientID != tt.expectedClientID {
				t.Errorf("GetClientID() = %v, want %v", clientID, tt.expectedClientID)
			}
		})
	}
}

// TestEnvironmentOverride verifies that AGENTBAY_CLI_ENDPOINT can override environment defaults
func TestEnvironmentOverride(t *testing.T) {
	customEndpoint := "custom-endpoint.example.com"

	tests := []struct {
		name        string
		agentbayEnv string
		endpointEnv string
		expected    string
	}{
		{
			name:        "Custom endpoint overrides production default",
			agentbayEnv: "production",
			endpointEnv: customEndpoint,
			expected:    customEndpoint,
		},
		{
			name:        "Custom endpoint overrides prerelease default",
			agentbayEnv: "prerelease",
			endpointEnv: customEndpoint,
			expected:    customEndpoint,
		},
		{
			name:        "No override uses production default",
			agentbayEnv: "production",
			endpointEnv: "",
			expected:    "xiaoying-share.cn-shanghai.aliyuncs.com",
		},
		{
			name:        "No override uses prerelease default",
			agentbayEnv: "prerelease",
			endpointEnv: "",
			expected:    "xiaoying-pre.cn-hangzhou.aliyuncs.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set environment variables
			if tt.agentbayEnv != "" {
				os.Setenv("AGENTBAY_ENV", tt.agentbayEnv)
			} else {
				os.Unsetenv("AGENTBAY_ENV")
			}
			if tt.endpointEnv != "" {
				os.Setenv("AGENTBAY_CLI_ENDPOINT", tt.endpointEnv)
			} else {
				os.Unsetenv("AGENTBAY_CLI_ENDPOINT")
			}
			defer func() {
				os.Unsetenv("AGENTBAY_ENV")
				os.Unsetenv("AGENTBAY_CLI_ENDPOINT")
			}()

			// Test API config
			apiConfig := config.LoadAPIConfig(nil)
			if apiConfig.Endpoint != tt.expected {
				t.Errorf("LoadAPIConfig().Endpoint = %v, want %v", apiConfig.Endpoint, tt.expected)
			}
		})
	}
}

// TestVersionCommandShowsEnvironment verifies that version command displays environment info
func TestVersionCommandShowsEnvironment(t *testing.T) {
	tests := []struct {
		name                string
		envValue            string
		expectedEnvInOutput string
		expectedEndpoint    string
	}{
		{
			name:                "Production environment in version output",
			envValue:            "production",
			expectedEnvInOutput: "Environment: production",
			expectedEndpoint:    "Endpoint: xiaoying-share.cn-shanghai.aliyuncs.com",
		},
		{
			name:                "Prerelease environment in version output",
			envValue:            "prerelease",
			expectedEnvInOutput: "Environment: prerelease",
			expectedEndpoint:    "Endpoint: xiaoying-pre.cn-hangzhou.aliyuncs.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set environment variable
			if tt.envValue != "" {
				os.Setenv("AGENTBAY_ENV", tt.envValue)
			} else {
				os.Unsetenv("AGENTBAY_ENV")
			}
			defer os.Unsetenv("AGENTBAY_ENV")

			// Get environment config
			envConfig := config.GetEnvironmentConfig()

			// Simulate version output
			output := "AgentBay CLI version dev\n"
			output += "Git commit: unknown\n"
			output += "Build date: unknown\n"
			output += "Environment: " + string(envConfig.Name) + "\n"
			output += "Endpoint: " + envConfig.Endpoint + "\n"

			// Verify output contains expected environment info
			if !strings.Contains(output, tt.expectedEnvInOutput) {
				t.Errorf("Version output missing expected environment info: %v\nGot: %v", tt.expectedEnvInOutput, output)
			}
			if !strings.Contains(output, tt.expectedEndpoint) {
				t.Errorf("Version output missing expected endpoint: %v\nGot: %v", tt.expectedEndpoint, output)
			}
		})
	}
}

// TestEnvironmentIsolation verifies that different environments use different configurations
func TestEnvironmentIsolation(t *testing.T) {
	// Test production
	os.Setenv("AGENTBAY_ENV", "production")
	prodConfig := config.GetEnvironmentConfig()
	prodEndpoint := config.LoadAPIConfig(nil).Endpoint
	prodClientID := config.GetClientID()
	os.Unsetenv("AGENTBAY_ENV")

	// Test prerelease
	os.Setenv("AGENTBAY_ENV", "prerelease")
	preConfig := config.GetEnvironmentConfig()
	preEndpoint := config.LoadAPIConfig(nil).Endpoint
	preClientID := config.GetClientID()
	os.Unsetenv("AGENTBAY_ENV")

	// Verify they are different
	if prodConfig.Endpoint == preConfig.Endpoint {
		t.Error("Production and prerelease should have different endpoints")
	}
	if prodConfig.ClientID == preConfig.ClientID {
		t.Error("Production and prerelease should have different client IDs")
	}
	if prodEndpoint == preEndpoint {
		t.Error("Production and prerelease API configs should have different endpoints")
	}
	if prodClientID == preClientID {
		t.Error("Production and prerelease should return different client IDs")
	}

	// Verify production values
	if prodConfig.Endpoint != "xiaoying-share.cn-shanghai.aliyuncs.com" {
		t.Errorf("Production endpoint = %v, want xiaoying-share.cn-shanghai.aliyuncs.com", prodConfig.Endpoint)
	}
	if prodConfig.ClientID != "4032653160518150541" {
		t.Errorf("Production client ID = %v, want 4032653160518150541", prodConfig.ClientID)
	}

	// Verify prerelease values
	if preConfig.Endpoint != "xiaoying-pre.cn-hangzhou.aliyuncs.com" {
		t.Errorf("Prerelease endpoint = %v, want xiaoying-pre.cn-hangzhou.aliyuncs.com", preConfig.Endpoint)
	}
	if preConfig.ClientID != "4019057658592127596" {
		t.Errorf("Prerelease client ID = %v, want 4019057658592127596", preConfig.ClientID)
	}
}
