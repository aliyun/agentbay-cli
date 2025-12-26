// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"os"
	"testing"

	"github.com/agentbay/agentbay-cli/internal/config"
)

func TestGetEnvironment(t *testing.T) {
	tests := []struct {
		name     string
		envValue string
		want     config.Environment
	}{
		{
			name:     "production when env is production",
			envValue: "production",
			want:     config.EnvProduction,
		},
		{
			name:     "production when env is prod",
			envValue: "prod",
			want:     config.EnvProduction,
		},
		{
			name:     "production when env is empty",
			envValue: "",
			want:     config.EnvProduction,
		},
		{
			name:     "prerelease when env is prerelease",
			envValue: "prerelease",
			want:     config.EnvPreRelease,
		},
		{
			name:     "prerelease when env is pre",
			envValue: "pre",
			want:     config.EnvPreRelease,
		},
		{
			name:     "prerelease when env is staging",
			envValue: "staging",
			want:     config.EnvPreRelease,
		},
		{
			name:     "production when env is unknown",
			envValue: "unknown",
			want:     config.EnvProduction,
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

			got := config.GetEnvironment()
			if got != tt.want {
				t.Errorf("GetEnvironment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEnvironmentConfig(t *testing.T) {
	tests := []struct {
		name         string
		envValue     string
		wantEnv      config.Environment
		wantEndpoint string
		wantClientID string
	}{
		{
			name:         "production config",
			envValue:     "production",
			wantEnv:      config.EnvProduction,
			wantEndpoint: "xiaoying.cn-shanghai.aliyuncs.com",
			wantClientID: "4032653160518150541",
		},
		{
			name:         "prerelease config",
			envValue:     "prerelease",
			wantEnv:      config.EnvPreRelease,
			wantEndpoint: "xiaoying-pre.cn-hangzhou.aliyuncs.com",
			wantClientID: "4019057658592127596",
		},
		{
			name:         "default to production",
			envValue:     "",
			wantEnv:      config.EnvProduction,
			wantEndpoint: "xiaoying.cn-shanghai.aliyuncs.com",
			wantClientID: "4032653160518150541",
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

			got := config.GetEnvironmentConfig()
			if got.Name != tt.wantEnv {
				t.Errorf("GetEnvironmentConfig().Name = %v, want %v", got.Name, tt.wantEnv)
			}
			if got.Endpoint != tt.wantEndpoint {
				t.Errorf("GetEnvironmentConfig().Endpoint = %v, want %v", got.Endpoint, tt.wantEndpoint)
			}
			if got.ClientID != tt.wantClientID {
				t.Errorf("GetEnvironmentConfig().ClientID = %v, want %v", got.ClientID, tt.wantClientID)
			}
		})
	}
}

func TestGetClientID(t *testing.T) {
	tests := []struct {
		name     string
		envValue string
		want     string
	}{
		{
			name:     "production client ID",
			envValue: "production",
			want:     "4032653160518150541",
		},
		{
			name:     "prerelease client ID",
			envValue: "prerelease",
			want:     "4019057658592127596",
		},
		{
			name:     "default client ID",
			envValue: "",
			want:     "4032653160518150541",
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

			got := config.GetClientID()
			if got != tt.want {
				t.Errorf("GetClientID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDefaultEndpoint(t *testing.T) {
	tests := []struct {
		name     string
		envValue string
		want     string
	}{
		{
			name:     "production endpoint",
			envValue: "production",
			want:     "xiaoying.cn-shanghai.aliyuncs.com",
		},
		{
			name:     "prerelease endpoint",
			envValue: "prerelease",
			want:     "xiaoying-pre.cn-hangzhou.aliyuncs.com",
		},
		{
			name:     "default endpoint",
			envValue: "",
			want:     "xiaoying.cn-shanghai.aliyuncs.com",
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

			got := config.GetDefaultEndpoint()
			if got != tt.want {
				t.Errorf("GetDefaultEndpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
