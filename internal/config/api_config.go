// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// APIConfig stores API configuration
type APIConfig struct {
	RegionID  string `json:"region_id"`
	Endpoint  string `json:"endpoint"`
	TimeoutMs int    `json:"timeout_ms"`
}

// DefaultAPIConfig returns the default API configuration
func DefaultAPIConfig() APIConfig {
	return APIConfig{
		RegionID:  "cn-shanghai",
		Endpoint:  "xiaoying-share.cn-shanghai.aliyuncs.com",
		TimeoutMs: 60000,
	}
}

// LoadAPIConfig loads the API configuration from environment variables or uses defaults
func LoadAPIConfig(cfg *APIConfig) APIConfig {
	if cfg != nil {
		// If config is explicitly provided, use it directly
		return APIConfig{
			RegionID:  cfg.RegionID,
			Endpoint:  cfg.Endpoint,
			TimeoutMs: cfg.TimeoutMs,
		}
	}

	// Use environment variables if set, otherwise use defaults
	config := DefaultAPIConfig()

	if regionID := os.Getenv("AGENTBAY_CLI_REGION_ID"); regionID != "" {
		config.RegionID = regionID
		log.Debugf("[DEBUG] Using region ID from environment: %s", regionID)
	} else {
		log.Debugf("[DEBUG] Using default region ID: %s", config.RegionID)
	}

	if endpoint := os.Getenv("AGENTBAY_CLI_ENDPOINT"); endpoint != "" {
		config.Endpoint = endpoint
		log.Debugf("[DEBUG] Using endpoint from environment: %s", endpoint)
	} else {
		log.Debugf("[DEBUG] Using default endpoint: %s", config.Endpoint)
	}

	// Also check the legacy AGENTBAY_API_URL for backward compatibility
	if apiURL := os.Getenv("AGENTBAY_API_URL"); apiURL != "" {
		config.Endpoint = apiURL
		log.Debugf("[DEBUG] Using endpoint from AGENTBAY_API_URL (legacy): %s", apiURL)
	}

	if timeoutMS := os.Getenv("AGENTBAY_CLI_TIMEOUT_MS"); timeoutMS != "" {
		if timeout, err := strconv.Atoi(timeoutMS); err == nil {
			config.TimeoutMs = timeout
			log.Debugf("[DEBUG] Using timeout from environment: %d ms", timeout)
		} else {
			log.Warnf("Warning: Failed to parse AGENTBAY_CLI_TIMEOUT_MS as integer: %v, using default value %d", err, config.TimeoutMs)
		}
	} else {
		log.Debugf("[DEBUG] Using default timeout: %d ms", config.TimeoutMs)
	}

	return config
}

// GetFullEndpoint returns the full endpoint URL with https:// prefix if needed
func (c *APIConfig) GetFullEndpoint() string {
	endpoint := c.Endpoint

	// Add https:// prefix if not present
	if endpoint != "" && !hasProtocol(endpoint) {
		endpoint = "https://" + endpoint
	}

	return endpoint
}

// hasProtocol checks if the endpoint already has a protocol prefix
func hasProtocol(endpoint string) bool {
	return (len(endpoint) >= 7 && endpoint[:7] == "http://") ||
		(len(endpoint) >= 8 && endpoint[:8] == "https://")
}
