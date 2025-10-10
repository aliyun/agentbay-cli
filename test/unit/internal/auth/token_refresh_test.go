// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package auth_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/agentbay/agentbay-cli/internal/auth"
)

// mockTokenConfig implements auth.TokenConfig for testing
type mockTokenConfig struct {
	accessToken   string
	refreshToken  string
	expiresAt     time.Time
	isExpired     bool
	refreshCalled bool
	clearCalled   bool
	refreshErr    error
	getTokensErr  error
}

func (m *mockTokenConfig) GetTokens() (accessToken string, refreshToken string, expiresAt time.Time, err error) {
	if m.getTokensErr != nil {
		return "", "", time.Time{}, m.getTokensErr
	}
	return m.accessToken, m.refreshToken, m.expiresAt, nil
}

func (m *mockTokenConfig) RefreshTokens(accessToken, tokenType string, expiresIn int) error {
	m.refreshCalled = true
	if m.refreshErr != nil {
		return m.refreshErr
	}
	// Update mock token
	m.accessToken = accessToken
	m.expiresAt = time.Now().Add(time.Duration(expiresIn) * time.Second)
	m.isExpired = false
	return nil
}

func (m *mockTokenConfig) IsTokenExpired() bool {
	return m.isExpired
}

func (m *mockTokenConfig) ClearTokens() error {
	m.clearCalled = true
	m.accessToken = ""
	m.refreshToken = ""
	return nil
}

func TestRefreshTokenIfNeeded(t *testing.T) {
	t.Run("should not refresh when token is valid and not expiring soon", func(t *testing.T) {
		cfg := &mockTokenConfig{
			accessToken:  "valid-token",
			refreshToken: "refresh-token",
			expiresAt:    time.Now().Add(10 * time.Minute), // Expires in 10 minutes
			isExpired:    false,
		}

		err := auth.RefreshTokenIfNeeded(cfg, "test-client-id")
		assert.NoError(t, err)
		assert.False(t, cfg.refreshCalled, "should not call refresh when token is valid")
	})

	t.Run("should refresh when token is about to expire", func(t *testing.T) {
		cfg := &mockTokenConfig{
			accessToken:  "expiring-token",
			refreshToken: "refresh-token",
			expiresAt:    time.Now().Add(3 * time.Minute), // Expires in 3 minutes (< 5 min threshold)
			isExpired:    false,
		}

		// Note: This test will fail in real scenario because it tries to call the real API
		// In a real implementation, we would need to mock the HTTP client
		// For now, we just test the logic flow
		err := auth.RefreshTokenIfNeeded(cfg, "test-client-id")

		// We expect an error because the mock doesn't have a real refresh token
		// But we can verify the logic was attempted
		assert.Error(t, err, "expected error when refresh token is invalid")
		assert.True(t, cfg.clearCalled, "should clear tokens when refresh fails")
	})

	t.Run("should refresh when token is expired", func(t *testing.T) {
		cfg := &mockTokenConfig{
			accessToken:  "expired-token",
			refreshToken: "refresh-token",
			expiresAt:    time.Now().Add(-1 * time.Hour), // Already expired
			isExpired:    true,
		}

		err := auth.RefreshTokenIfNeeded(cfg, "test-client-id")

		// We expect an error because the mock doesn't have a real refresh token
		assert.Error(t, err, "expected error when refresh token is invalid")
		assert.True(t, cfg.clearCalled, "should clear tokens when refresh fails")
	})

	t.Run("should return error when no token found", func(t *testing.T) {
		cfg := &mockTokenConfig{
			getTokensErr: assert.AnError,
		}

		err := auth.RefreshTokenIfNeeded(cfg, "test-client-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "no valid token found")
	})

	t.Run("should handle token at just over 5 minute threshold", func(t *testing.T) {
		cfg := &mockTokenConfig{
			accessToken:  "threshold-token",
			refreshToken: "refresh-token",
			expiresAt:    time.Now().Add(5*time.Minute + 1*time.Second), // Just over 5 minutes
			isExpired:    false,
		}

		err := auth.RefreshTokenIfNeeded(cfg, "test-client-id")

		// Just over 5 minutes should not refresh (> 5 minutes condition)
		assert.NoError(t, err)
		assert.False(t, cfg.refreshCalled, "should not refresh when token is just over 5 minute threshold")
	})

	t.Run("should refresh when token is just under 5 minute threshold", func(t *testing.T) {
		cfg := &mockTokenConfig{
			accessToken:  "almost-threshold-token",
			refreshToken: "refresh-token",
			expiresAt:    time.Now().Add(4*time.Minute + 59*time.Second), // Just under 5 minutes
			isExpired:    false,
		}

		err := auth.RefreshTokenIfNeeded(cfg, "test-client-id")

		// Just under 5 minutes should trigger refresh
		assert.Error(t, err, "expected error when refresh token is invalid")
		assert.True(t, cfg.clearCalled, "should attempt refresh just under 5 minute threshold")
	})
}

func TestTokenConfig(t *testing.T) {
	t.Run("mockTokenConfig should implement auth.TokenConfig", func(t *testing.T) {
		var _ auth.TokenConfig = (*mockTokenConfig)(nil)
	})
}
