// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// Environment represents the deployment environment
type Environment string

const (
	// EnvProduction is the production environment (China)
	EnvProduction Environment = "production"
	// EnvPreRelease is the pre-release environment (China)
	EnvPreRelease Environment = "prerelease"
	// EnvInternationalProduction is the international production environment (e.g. ap-southeast-1)
	EnvInternationalProduction Environment = "international"
	// EnvInternationalPreRelease is the international pre-release environment (to be configured)
	EnvInternationalPreRelease Environment = "international-pre"
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
		Endpoint: "xiaoying.cn-shanghai.aliyuncs.com",
		ClientID: "4032653160518150541",
	}

	// Pre-release environment configuration
	prereleaseConfig = EnvironmentConfig{
		Name:     EnvPreRelease,
		Endpoint: "xiaoying-pre.cn-hangzhou.aliyuncs.com",
		ClientID: "4019057658592127596",
	}

	// International production: default endpoint and OAuth client for alibabacloud.com
	productionInternationalConfig = EnvironmentConfig{
		Name:     EnvInternationalProduction,
		Endpoint: "xiaoying.ap-southeast-1.aliyuncs.com",
		ClientID: "4192690673476752832",
	}

	// International pre-release: placeholder, to be configured later (预发)
	preReleaseInternationalConfig = EnvironmentConfig{
		Name:     EnvInternationalPreRelease,
		Endpoint: "xiaoying-pre.ap-southeast-1.aliyuncs.com", // TODO: replace when international pre is available
		ClientID: "4192690673476752832",                      // TODO: replace with international pre OAuth client ID
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
	case "international", "prod-international", "intl", "international-prod":
		log.Debugf("[DEBUG] Using international production environment")
		return EnvInternationalProduction
	case "international-pre", "pre-international", "intl-pre", "staging-international":
		log.Debugf("[DEBUG] Using international pre-release environment")
		return EnvInternationalPreRelease
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
	case EnvInternationalProduction:
		return productionInternationalConfig
	case EnvInternationalPreRelease:
		return preReleaseInternationalConfig
	default:
		return productionConfig
	}
}

// GetClientID returns the OAuth client ID for the current environment.
// If AGENTBAY_OAUTH_CLIENT_ID is set, it overrides the environment default.
// Use this for international login: the client ID from domestic (aliyun.com) is not
// valid on international (alibabacloud.com); set AGENTBAY_OAUTH_CLIENT_ID to the
// client ID of an app registered on Alibaba Cloud International.
func GetClientID() string {
	if id := os.Getenv("AGENTBAY_OAUTH_CLIENT_ID"); id != "" {
		log.Debugf("[DEBUG] Using OAuth client ID from AGENTBAY_OAUTH_CLIENT_ID")
		return id
	}
	return GetEnvironmentConfig().ClientID
}

// GetDefaultEndpoint returns the default API endpoint for the current environment
func GetDefaultEndpoint() string {
	return GetEnvironmentConfig().Endpoint
}
