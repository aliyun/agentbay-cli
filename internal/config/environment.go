// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Environment represents the deployment environment
type Environment string

const (
	// EnvProduction is the production environment
	EnvProduction Environment = "production"
	// EnvPreRelease is the pre-release environment
	EnvPreRelease Environment = "prerelease"
)

// EnvironmentConfig holds environment-specific configuration
type EnvironmentConfig struct {
	Name     Environment
	Endpoint string
	ClientID string
}

var (
	// Production environment configuration
	productionConfig = EnvironmentConfig{
		Name:     EnvProduction,
		Endpoint: "xiaoying-share.cn-shanghai.aliyuncs.com",
		ClientID: "4032653160518150541",
	}

	// Pre-release environment configuration
	prereleaseConfig = EnvironmentConfig{
		Name:     EnvPreRelease,
		Endpoint: "xiaoying-pre.cn-hangzhou.aliyuncs.com",
		ClientID: "4019057658592127596",
	}
)

// GetEnvironment returns the current environment based on AGENTBAY_ENV
// Defaults to production if not set or invalid
func GetEnvironment() Environment {
	env := os.Getenv("AGENTBAY_ENV")

	switch env {
	case "prerelease", "pre", "staging":
		log.Debugf("[DEBUG] Using pre-release environment")
		return EnvPreRelease
	case "production", "prod", "":
		log.Debugf("[DEBUG] Using production environment")
		return EnvProduction
	default:
		log.Warnf("[WARN] Unknown environment '%s', defaulting to production", env)
		return EnvProduction
	}
}

// GetEnvironmentConfig returns the configuration for the current environment
func GetEnvironmentConfig() EnvironmentConfig {
	env := GetEnvironment()

	switch env {
	case EnvPreRelease:
		return prereleaseConfig
	default:
		return productionConfig
	}
}

// GetClientID returns the OAuth client ID for the current environment
func GetClientID() string {
	return GetEnvironmentConfig().ClientID
}

// GetDefaultEndpoint returns the default API endpoint for the current environment
func GetDefaultEndpoint() string {
	return GetEnvironmentConfig().Endpoint
}

// IsDomesticEndpoint checks if the endpoint is domestic (China) based on domain
func IsDomesticEndpoint(endpoint string) bool {
	// Domestic endpoints typically contain .cn or cn- prefix
	// e.g., xiaoying-share.cn-shanghai.aliyuncs.com (domestic)
	return strings.Contains(endpoint, ".cn") || strings.Contains(endpoint, "cn-")
}

// GetDefaultSourceImageId returns the default SourceImageId based on environment and endpoint
func GetDefaultSourceImageId(env Environment, endpoint string) string {
	isDomestic := IsDomesticEndpoint(endpoint)

	if env == EnvPreRelease {
		if isDomestic {
			return "imgc-07if81rziujpkp72y"
		}
		return "code-space-debian-12"
	}

	// Production environment: both domestic and international use the same
	return "code-space-debian-12"
}
