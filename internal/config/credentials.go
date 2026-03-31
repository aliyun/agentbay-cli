// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"fmt"
	"os"
	"strings"
)

// Environment variable names for AccessKey-style authentication (alternative to OAuth).
const (
	EnvAccessKeyID           = "AGENTBAY_ACCESS_KEY_ID"
	EnvAccessKeySecret       = "AGENTBAY_ACCESS_KEY_SECRET"
	EnvAccessKeySessionToken = "AGENTBAY_ACCESS_KEY_SESSION_TOKEN"
)

// HasAccessKeyFromEnv reports whether both AccessKey ID and secret are set (non-empty after trim).
func HasAccessKeyFromEnv() bool {
	_, _, _, ok := AccessKeyFromEnv()
	return ok
}

// AccessKeyFromEnv reads AccessKey credentials from the environment.
// ok is true when both ID and secret are non-empty.
func AccessKeyFromEnv() (accessKeyID, accessKeySecret, securityToken string, ok bool) {
	accessKeyID = strings.TrimSpace(os.Getenv(EnvAccessKeyID))
	accessKeySecret = strings.TrimSpace(os.Getenv(EnvAccessKeySecret))
	securityToken = strings.TrimSpace(os.Getenv(EnvAccessKeySessionToken))
	if accessKeyID == "" || accessKeySecret == "" {
		return "", "", "", false
	}
	return accessKeyID, accessKeySecret, securityToken, true
}

// ErrNotAuthenticated is returned when neither OAuth tokens nor AccessKey env credentials are available.
func ErrNotAuthenticated() error {
	return fmt.Errorf("not authenticated. Please run 'agentbay login' or set %s and %s", EnvAccessKeyID, EnvAccessKeySecret)
}
