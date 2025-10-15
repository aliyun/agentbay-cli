// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package config_test

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/agentbay/agentbay-cli/internal/config"
)

func TestConfig(t *testing.T) {
	// Create temporary directory for test config
	tempDir, err := os.MkdirTemp("", "agentbay-config-test")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Set environment variable to use temp directory
	originalConfigDir := os.Getenv("AGENTBAY_CLI_CONFIG_DIR")
	os.Setenv("AGENTBAY_CLI_CONFIG_DIR", tempDir)
	defer func() {
		if originalConfigDir == "" {
			os.Unsetenv("AGENTBAY_CLI_CONFIG_DIR")
		} else {
			os.Setenv("AGENTBAY_CLI_CONFIG_DIR", originalConfigDir)
		}
	}()

	t.Run("GetConfig creates new config when file doesn't exist", func(t *testing.T) {
		cfg, err := config.GetConfig()
		assert.NoError(t, err)
		assert.NotNil(t, cfg)
		assert.Nil(t, cfg.Token)
		assert.False(t, cfg.IsAuthenticated())
	})

	t.Run("SaveTokens should save OAuth tokens correctly", func(t *testing.T) {
		cfg, err := config.GetConfig()
		require.NoError(t, err)

		accessToken := "eyJraWQiOiJrMTIzNCIsImVuYyI6test"
		tokenType := "Bearer"
		expiresIn := 3600
		refreshToken := "Ccx63VVeTn2dxV7ovXXfLtAqLLERAtest"
		idToken := "eyJhbGciOiJIUzI1test"

		err = cfg.SaveTokens(accessToken, tokenType, expiresIn, refreshToken, idToken)
		assert.NoError(t, err)
		assert.True(t, cfg.IsAuthenticated())

		// Verify tokens are saved correctly
		token, err := cfg.GetToken()
		assert.NoError(t, err)
		assert.Equal(t, accessToken, token.AccessToken)
		assert.Equal(t, tokenType, token.TokenType)
		assert.Equal(t, refreshToken, token.RefreshToken)
		assert.Equal(t, idToken, token.IDToken)
		assert.True(t, token.ExpiresAt.After(time.Now()))
	})

	t.Run("GetConfig should load existing config from file", func(t *testing.T) {
		// First save a config
		config1, err := config.GetConfig()
		require.NoError(t, err)

		accessToken := "test-access-token"
		err = config1.SaveTokens(accessToken, "Bearer", 3600, "test-refresh", "test-id")
		require.NoError(t, err)

		// Load config again
		config2, err := config.GetConfig()
		assert.NoError(t, err)
		assert.True(t, config2.IsAuthenticated())

		token, err := config2.GetToken()
		assert.NoError(t, err)
		assert.Equal(t, accessToken, token.AccessToken)
	})

	t.Run("ClearTokens should remove authentication data", func(t *testing.T) {
		cfg, err := config.GetConfig()
		require.NoError(t, err)

		// First save tokens
		err = cfg.SaveTokens("test-token", "Bearer", 3600, "refresh", "id")
		require.NoError(t, err)
		assert.True(t, cfg.IsAuthenticated())

		// Clear tokens
		err = cfg.ClearTokens()
		assert.NoError(t, err)
		assert.False(t, cfg.IsAuthenticated())

		// Verify tokens are cleared
		_, err = cfg.GetToken()
		assert.Error(t, err)
		assert.Equal(t, config.ErrNoTokenFound, err)
	})

	t.Run("RefreshTokens should update access token", func(t *testing.T) {
		cfg, err := config.GetConfig()
		require.NoError(t, err)

		// Save initial tokens
		err = cfg.SaveTokens("old-token", "Bearer", 3600, "refresh-token", "id-token")
		require.NoError(t, err)

		// Refresh tokens
		newAccessToken := "new-access-token"
		err = cfg.RefreshTokens(newAccessToken, "Bearer", 7200)
		assert.NoError(t, err)

		// Verify new token is saved
		token, err := cfg.GetToken()
		assert.NoError(t, err)
		assert.Equal(t, newAccessToken, token.AccessToken)
		assert.Equal(t, "refresh-token", token.RefreshToken) // Should remain unchanged
	})

	t.Run("IsTokenExpired should check token expiration correctly", func(t *testing.T) {
		cfg, err := config.GetConfig()
		require.NoError(t, err)

		// Save token that expires in 1 second
		err = cfg.SaveTokens("test-token", "Bearer", 1, "refresh", "id")
		require.NoError(t, err)

		// Should not be expired immediately
		assert.False(t, cfg.IsTokenExpired())

		// Wait for expiration
		time.Sleep(2 * time.Second)
		assert.True(t, cfg.IsTokenExpired())
	})
}

func TestConfigDir(t *testing.T) {
	t.Run("ConfigDir should use environment variable when set", func(t *testing.T) {
		testDir := "/tmp/test-agentbay-config"
		originalConfigDir := os.Getenv("AGENTBAY_CLI_CONFIG_DIR")
		os.Setenv("AGENTBAY_CLI_CONFIG_DIR", testDir)
		defer func() {
			if originalConfigDir == "" {
				os.Unsetenv("AGENTBAY_CLI_CONFIG_DIR")
			} else {
				os.Setenv("AGENTBAY_CLI_CONFIG_DIR", originalConfigDir)
			}
		}()

		configDir, err := config.ConfigDir()
		assert.NoError(t, err)
		assert.Equal(t, testDir, configDir)
	})

	t.Run("ConfigDir should use default when environment variable not set", func(t *testing.T) {
		originalConfigDir := os.Getenv("AGENTBAY_CLI_CONFIG_DIR")
		os.Unsetenv("AGENTBAY_CLI_CONFIG_DIR")
		defer func() {
			if originalConfigDir != "" {
				os.Setenv("AGENTBAY_CLI_CONFIG_DIR", originalConfigDir)
			}
		}()

		configDir, err := config.ConfigDir()
		assert.NoError(t, err)
		assert.Contains(t, configDir, "agentbay")
	})
}

func TestConfigFile(t *testing.T) {
	t.Run("ConfigFile should return correct path", func(t *testing.T) {
		tempDir, err := os.MkdirTemp("", "agentbay-config-test")
		require.NoError(t, err)
		defer os.RemoveAll(tempDir)

		originalConfigDir := os.Getenv("AGENTBAY_CLI_CONFIG_DIR")
		os.Setenv("AGENTBAY_CLI_CONFIG_DIR", tempDir)
		defer func() {
			if originalConfigDir == "" {
				os.Unsetenv("AGENTBAY_CLI_CONFIG_DIR")
			} else {
				os.Setenv("AGENTBAY_CLI_CONFIG_DIR", originalConfigDir)
			}
		}()

		configFile, err := config.ConfigFile()
		assert.NoError(t, err)
		assert.Equal(t, filepath.Join(tempDir, "config.json"), configFile)
	})
}
