// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package auth

import (
	"time"
)

// ConfigAdapter provides a type that implements the TokenConfig interface
// while wrapping a config object that has different method signatures
type ConfigAdapter struct {
	getTokens     func() (accessToken string, refreshToken string, expiresAt time.Time, err error)
	refreshTokens func(accessToken, tokenType string, expiresIn int) error
	isExpired     func() bool
	clearTokens   func() error
}

// NewConfigAdapter creates a new ConfigAdapter
func NewConfigAdapter(
	getTokens func() (string, string, time.Time, error),
	refreshTokens func(string, string, int) error,
	isExpired func() bool,
	clearTokens func() error,
) TokenConfig {
	return &ConfigAdapter{
		getTokens:     getTokens,
		refreshTokens: refreshTokens,
		isExpired:     isExpired,
		clearTokens:   clearTokens,
	}
}

func (ca *ConfigAdapter) GetTokens() (accessToken string, refreshToken string, expiresAt time.Time, err error) {
	return ca.getTokens()
}

func (ca *ConfigAdapter) RefreshTokens(accessToken, tokenType string, expiresIn int) error {
	return ca.refreshTokens(accessToken, tokenType, expiresIn)
}

func (ca *ConfigAdapter) IsTokenExpired() bool {
	return ca.isExpired()
}

func (ca *ConfigAdapter) ClearTokens() error {
	return ca.clearTokens()
}
