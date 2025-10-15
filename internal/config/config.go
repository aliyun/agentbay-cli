// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"time"
)

// Config represents the CLI configuration
type Config struct {
	Token *Token `json:"token,omitempty"` // OAuth token authentication
}

// Token represents OAuth authentication tokens
type Token struct {
	AccessToken  string    `json:"access_token"`
	TokenType    string    `json:"token_type"`
	ExpiresIn    int       `json:"expires_in"`
	RefreshToken string    `json:"refresh_token"`
	IDToken      string    `json:"id_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}

var (
	ErrNoTokenFound = errors.New("no authentication token found. Run 'agentbay login' to authenticate")
)

// GetConfig loads the configuration from file or creates a new one
func GetConfig() (*Config, error) {
	configFilePath, err := getConfigPath()
	if err != nil {
		return nil, err
	}

	var c Config

	// Try to load existing config file
	_, err = os.Stat(configFilePath)
	if os.IsNotExist(err) {
		// No config file exists, create new config
		c = Config{}
	} else if err != nil {
		return nil, err
	} else {
		// Config file exists, load it
		configContent, err := os.ReadFile(configFilePath)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(configContent, &c)
		if err != nil {
			return nil, err
		}
	}

	return &c, nil
}

// Save writes the configuration to file
func (c *Config) Save() error {
	configFilePath, err := getConfigPath()
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(configFilePath), 0755)
	if err != nil {
		return err
	}

	configContent, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configFilePath, configContent, 0600) // More secure permissions for auth data
}

// GetToken retrieves the authentication token object
func (c *Config) GetToken() (*Token, error) {
	if c.Token == nil {
		return nil, ErrNoTokenFound
	}
	return c.Token, nil
}

// GetTokens returns token information for the refresh mechanism
// This method implements the auth.TokenConfig interface
func (c *Config) GetTokens() (accessToken string, refreshToken string, expiresAt time.Time, err error) {
	if c.Token == nil {
		return "", "", time.Time{}, ErrNoTokenFound
	}
	return c.Token.AccessToken, c.Token.RefreshToken, c.Token.ExpiresAt, nil
}

// SaveTokens saves OAuth authentication tokens to the configuration
func (c *Config) SaveTokens(accessToken, tokenType string, expiresIn int, refreshToken, idToken string) error {
	// Calculate expiration time
	expiresAt := time.Now().Add(time.Duration(expiresIn) * time.Second)

	// Update config with tokens
	c.Token = &Token{
		AccessToken:  accessToken,
		TokenType:    tokenType,
		ExpiresIn:    expiresIn,
		RefreshToken: refreshToken,
		IDToken:      idToken,
		ExpiresAt:    expiresAt,
	}

	return c.Save()
}

// RefreshTokens updates the access token with new values from refresh
func (c *Config) RefreshTokens(accessToken, tokenType string, expiresIn int) error {
	if c.Token == nil {
		return ErrNoTokenFound
	}

	// Update access token and expiration, keep refresh token unchanged
	c.Token.AccessToken = accessToken
	c.Token.TokenType = tokenType
	c.Token.ExpiresIn = expiresIn
	c.Token.ExpiresAt = time.Now().Add(time.Duration(expiresIn) * time.Second)

	return c.Save()
}

// ClearTokens removes authentication tokens from the configuration
func (c *Config) ClearTokens() error {
	c.Token = nil
	return c.Save()
}

// IsAuthenticated checks if the user is authenticated (has tokens)
func (c *Config) IsAuthenticated() bool {
	return c.Token != nil && c.Token.AccessToken != ""
}

// IsTokenExpired checks if the access token is expired
func (c *Config) IsTokenExpired() bool {
	if c.Token == nil {
		return true
	}
	return time.Now().After(c.Token.ExpiresAt)
}

// getConfigPath returns the path to the configuration file
func getConfigPath() (string, error) {
	configDir, err := ConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(configDir, "config.json"), nil
}

// ConfigDir returns the configuration directory path
func ConfigDir() (string, error) {
	// Check for environment variable override first
	agentbayConfigDir := os.Getenv("AGENTBAY_CLI_CONFIG_DIR")
	if agentbayConfigDir != "" {
		return agentbayConfigDir, nil
	}

	// Use OS-specific standard config directory
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(userConfigDir, "agentbay"), nil
}

// ConfigFile returns the configuration file path
func ConfigFile() (string, error) {
	configDir, err := ConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, "config.json"), nil
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	return &Config{}
}
