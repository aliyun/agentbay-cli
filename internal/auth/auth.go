// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/agentbay/agentbay-cli/internal/client"
)

// OAuth endpoints and constants
var (
	authEndpoint   = "https://signin.aliyun.com/oauth2/v1/auth"
	tokenEndpoint  = "https://oauth.aliyun.com/v1/token"
	revokeEndpoint = "https://oauth.aliyun.com/v1/revoke"
)

// OAuth client configuration
const (
	DefaultClientID = "4019057658592127596"
)

// TokenResponse represents the OAuth token response
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    string `json:"expires_in"` // Changed to string to handle server response
	RefreshToken string `json:"refresh_token"`
	IDToken      string `json:"id_token"`
}

// RefreshResponse represents the OAuth refresh token response
type RefreshResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// BuildAuthURL constructs the OAuth authorization URL
func BuildAuthURL(clientID, redirectURI, state string) string {
	params := url.Values{}
	params.Set("client_id", clientID)
	params.Set("redirect_uri", redirectURI)
	params.Set("response_type", "code")
	params.Set("state", state)
	params.Set("scope", "/acs/xiaoying")

	return authEndpoint + "?" + params.Encode()
}

// ExchangeCodeForToken exchanges authorization code for access token
func ExchangeCodeForToken(clientID, redirectURI, code string) (*TokenResponse, error) {
	data := url.Values{}
	data.Set("code", code)
	data.Set("client_id", clientID)
	data.Set("redirect_uri", redirectURI)
	data.Set("grant_type", "authorization_code")

	resp, err := http.PostForm(tokenEndpoint, data)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange code for token: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("token exchange failed with status: %d", resp.StatusCode)
	}

	var tokenResponse TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return nil, fmt.Errorf("failed to decode token response: %w", err)
	}

	return &tokenResponse, nil
}

// RefreshAccessToken refreshes the access token using refresh token
func RefreshAccessToken(clientID, refreshToken string) (*RefreshResponse, error) {
	data := url.Values{}
	data.Set("refresh_token", refreshToken)
	data.Set("client_id", clientID)
	data.Set("grant_type", "refresh_token")

	resp, err := http.PostForm(tokenEndpoint, data)
	if err != nil {
		return nil, fmt.Errorf("failed to refresh token: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("token refresh failed with status: %d", resp.StatusCode)
	}

	var refreshResponse RefreshResponse
	if err := json.NewDecoder(resp.Body).Decode(&refreshResponse); err != nil {
		return nil, fmt.Errorf("failed to decode refresh response: %w", err)
	}

	return &refreshResponse, nil
}

// RevokeToken revokes the given token with an optional token type hint
func RevokeToken(clientID, token string) error {
	return RevokeTokenWithHint(clientID, token, "")
}

// RevokeTokenWithHint revokes the given token with a specific token type hint
func RevokeTokenWithHint(clientID, token, tokenTypeHint string) error {
	data := url.Values{}
	data.Set("token", token)
	data.Set("client_id", clientID)

	// Add token_type_hint if provided (as per RFC 7009)
	if tokenTypeHint != "" {
		data.Set("token_type_hint", tokenTypeHint)
	}

	resp, err := http.PostForm(revokeEndpoint, data)
	if err != nil {
		return fmt.Errorf("failed to revoke token: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Read response body for error details
		bodyBytes, readErr := io.ReadAll(resp.Body)
		if readErr == nil && len(bodyBytes) > 0 {
			var result map[string]interface{}
			if json.Unmarshal(bodyBytes, &result) == nil {
				return fmt.Errorf("token revocation failed with status: %d, error: %v", resp.StatusCode, result)
			}
			return fmt.Errorf("token revocation failed with status: %d, body: %s", resp.StatusCode, string(bodyBytes))
		}
		return fmt.Errorf("token revocation failed with status: %d", resp.StatusCode)
	}

	return nil
}

// PortRetryConfig returns retry configuration for port availability checks
func PortRetryConfig() *client.RetryConfig {
	return &client.RetryConfig{
		MaxRetries:    2,                      // Retry 2 times (3 total attempts)
		InitialDelay:  500 * time.Millisecond, // Start with 500ms
		MaxDelay:      2 * time.Second,        // Max 2 seconds
		BackoffFactor: 2.0,                    // Double each time
	}
}

// checkPortAvailabilityWithRetry checks if a port is available with exponential backoff retry
// Returns true if port becomes available, false if still occupied after retries
func checkPortAvailabilityWithRetry(port string, retryConfig *client.RetryConfig) (bool, error) {
	delay := retryConfig.InitialDelay

	for attempt := 0; attempt <= retryConfig.MaxRetries; attempt++ {
		if !IsPortOccupied(port) {
			if attempt > 0 {
				log.Debugf("[RETRY] Port %s is now available after %d attempt(s)", port, attempt)
			}
			return true, nil
		}

		// Port is still occupied
		if attempt < retryConfig.MaxRetries {
			log.Debugf("[RETRY] Port %s is occupied (attempt %d/%d), retrying in %v...",
				port, attempt+1, retryConfig.MaxRetries+1, delay)
			time.Sleep(delay)

			// Calculate next delay with exponential backoff
			delay = time.Duration(float64(delay) * retryConfig.BackoffFactor)
			if delay > retryConfig.MaxDelay {
				delay = retryConfig.MaxDelay
			}
		}
	}

	return false, fmt.Errorf("port %s is still occupied after %d attempts", port, retryConfig.MaxRetries+1)
}

// StartCallbackServer starts a local HTTP server to handle OAuth callbacks
// It binds the port first to ensure atomic port acquisition
func StartCallbackServer(ctx context.Context, port string) (string, error) {
	// Bind port first (atomic operation - binding means we own the port)
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return "", fmt.Errorf("port %s is occupied. Please close the program using this port and try again", port)
	}

	var code string
	var serverErr error
	var wg sync.WaitGroup
	wg.Add(1)

	// Create a new ServeMux to avoid conflicts with global handlers
	mux := http.NewServeMux()
	server := &http.Server{
		Handler: mux,
	}

	mux.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		defer wg.Done()

		// Get authorization code
		code = r.URL.Query().Get("code")
		if code == "" {
			serverErr = fmt.Errorf("no code in callback")
			http.Error(w, "No code", http.StatusBadRequest)
			return
		}

		// Return success page
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(GetSuccessHTML()))

		// Delay server close to ensure browser receives the success page
		go func() {
			time.Sleep(500 * time.Millisecond)
			server.Close()
			listener.Close()
		}()
	})

	// Start server using the already-bound listener
	go func() {
		if err := server.Serve(listener); err != http.ErrServerClosed && err != nil {
			log.Debugf("[DEBUG] Server error: %v", err)
		}
	}()

	// Wait for callback or timeout
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		// Callback received
		if serverErr != nil {
			server.Close()
			listener.Close()
			return "", serverErr
		}
		return code, nil
	case <-ctx.Done():
		server.Close()
		listener.Close()
		return "", fmt.Errorf("callback timeout: %v", ctx.Err())
	}
}

// GenerateState generates a random state parameter for OAuth
func GenerateState() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to generate random state: %w", err)
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

// GetSuccessHTML returns the HTML page shown after successful authentication
func GetSuccessHTML() string {
	return `<!DOCTYPE html>
<html>
<head>
    <title>Authentication Successful</title>
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            margin: 0;
            background-color: #f5f5f5;
        }
        .container {
            text-align: center;
            background: white;
            padding: 2rem;
            border-radius: 8px;
            box-shadow: 0 2px 10px rgba(0,0,0,0.1);
        }
        .success {
            color: #28a745;
            font-size: 1.5rem;
            margin-bottom: 1rem;
        }
        .message {
            color: #666;
            margin-bottom: 1rem;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="success">Authentication Successful</div>
        <div class="message">You have successfully authenticated with AgentBay.</div>
        <div class="message">You can now close this window and return to the terminal.</div>
    </div>
</body>
</html>`
}

// IsPortOccupied checks if a port is already in use
func IsPortOccupied(port string) bool {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return true
	}
	ln.Close()
	return false
}

// TokenConfig interface for accessing config methods
// Note: We use a struct with time.Time directly to avoid circular dependency
type TokenConfig interface {
	GetTokens() (accessToken string, refreshToken string, expiresAt time.Time, err error)
	RefreshTokens(accessToken, tokenType string, expiresIn int) error
	IsTokenExpired() bool
	ClearTokens() error
}

// RefreshTokenIfNeeded checks and refreshes token if it's about to expire (within 5 minutes)
// This provides automatic token management for seamless API access
func RefreshTokenIfNeeded(cfg TokenConfig, clientID string) error {
	_, refreshToken, expiresAt, err := cfg.GetTokens()
	if err != nil {
		return fmt.Errorf("no valid token found: %w", err)
	}

	// Check if token is about to expire (within 5 minutes)
	if !cfg.IsTokenExpired() && time.Until(expiresAt) > 5*time.Minute {
		log.Debug("Token is still valid, no refresh needed")
		return nil
	}

	log.Info("Token is approaching expiry or expired, refreshing...")

	// Perform token refresh
	refreshResp, err := RefreshAccessToken(clientID, refreshToken)
	if err != nil {
		// If refresh fails, clear the tokens
		cfg.ClearTokens()
		return fmt.Errorf("token refresh failed, please run 'agentbay login' to reauthenticate: %w", err)
	}

	// Save new access token
	err = cfg.RefreshTokens(
		refreshResp.AccessToken,
		refreshResp.TokenType,
		refreshResp.ExpiresIn,
	)
	if err != nil {
		return fmt.Errorf("failed to save refreshed tokens: %w", err)
	}

	log.Info("Token refreshed successfully")
	return nil
}
