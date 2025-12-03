// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/browser"
	"github.com/spf13/cobra"

	"github.com/agentbay/agentbay-cli/internal/auth"
	"github.com/agentbay/agentbay-cli/internal/config"
)

// OAuth constants are now defined in constants.go

var LoginCmd = &cobra.Command{
	Use:     "login",
	Short:   "Log in to AgentBay",
	Long:    "Authenticate with AgentBay using OAuth in your browser",
	Args:    cobra.NoArgs,
	GroupID: "core",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runLogin(cmd)
	},
}

func runLogin(cmd *cobra.Command) error {
	fmt.Println("Starting AgentBay authentication...")

	// Check if already authenticated
	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	if cfg.IsAuthenticated() && !cfg.IsTokenExpired() {
		fmt.Println("You are already logged in to AgentBay!")
		return nil
	}

	// Generate random state for OAuth security
	state, err := auth.GenerateState()
	if err != nil {
		return fmt.Errorf("failed to generate OAuth state: %w", err)
	}

	// Build authorization URL
	authURL := auth.BuildAuthURL(GetClientID(), RedirectURI, state)

	// Start local callback server first
	fmt.Printf("Starting local callback server on port %s...\n", CallbackPort)

	// Create context with timeout for callback server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// Start callback server in background
	codeChan := make(chan string, 1)
	errChan := make(chan error, 1)

	go func() {
		code, err := auth.StartCallbackServer(ctx, CallbackPort)
		if err != nil {
			errChan <- err
			return
		}
		codeChan <- code
	}()

	// Wait for server to start or fail (with timeout)
	select {
	case err := <-errChan:
		// Server failed to start (e.g., port occupied)
		errStr := err.Error()
		if contains(errStr, "port") && contains(errStr, "occupied") {
			fmt.Fprintf(os.Stderr, "\n[ERROR] Port %s is occupied.\n", CallbackPort)
			fmt.Fprintf(os.Stderr, "Please close the program using this port and try again.\n")
			fmt.Fprintf(os.Stderr, "You can check which process is using the port with:\n")
			fmt.Fprintf(os.Stderr, "  - macOS/Linux: lsof -i :%s\n", CallbackPort)
			fmt.Fprintf(os.Stderr, "  - Windows: netstat -ano | findstr :%s\n", CallbackPort)
			return fmt.Errorf("port %s is occupied", CallbackPort)
		}
		return fmt.Errorf("failed to start callback server: %v", err)
	case <-time.After(500 * time.Millisecond):
		// Server should be ready by now (port check and startup completed)
	}

	// Server is ready, now open browser
	fmt.Println("Opening browser for authentication...")
	fmt.Printf("If the browser doesn't open automatically, please visit:\n%s\n\n", authURL)

	err = browser.OpenURL(authURL)
	if err != nil {
		fmt.Printf("Warning: Failed to open browser automatically: %v\n", err)
		fmt.Println("Please copy the URL above and paste it into your browser to complete authentication.")
	} else {
		fmt.Println("Browser opened successfully!")
	}

	fmt.Printf("Waiting for callback on http://localhost:%s/callback...\n", CallbackPort)

	// Wait for callback
	select {
	case code := <-codeChan:
		fmt.Println("Authentication successful!")
		fmt.Printf("Received authorization code: %s...\n", code[:min(len(code), 20)])

		// Exchange code for token
		fmt.Println("Exchanging authorization code for access token...")

		tokenResponse, err := auth.ExchangeCodeForToken(GetClientID(), RedirectURI, code)
		if err != nil {
			fmt.Printf("Debug: Token exchange failed with error: %v\n", err)
			return fmt.Errorf("failed to exchange code for token: %w", err)
		}
		fmt.Printf("Debug: Token exchange successful, access token length: %d\n", len(tokenResponse.AccessToken))

		// Convert ExpiresIn from string to int
		expiresIn, err := strconv.Atoi(tokenResponse.ExpiresIn)
		if err != nil {
			fmt.Printf("Warning: Invalid expires_in value '%s', using default 3600 seconds\n", tokenResponse.ExpiresIn)
			expiresIn = 3600
		}

		// Save tokens to configuration
		fmt.Println("Saving authentication tokens...")

		err = cfg.SaveTokens(
			tokenResponse.AccessToken,
			tokenResponse.TokenType,
			expiresIn,
			tokenResponse.RefreshToken,
			tokenResponse.IDToken,
		)
		if err != nil {
			fmt.Printf("Warning: Failed to save tokens: %v\n", err)
			fmt.Println("You are logged in, but tokens were not saved to config file.")
			return nil
		}

		fmt.Println("Authentication tokens saved successfully!")
		fmt.Println("You are now logged in to AgentBay!")

		return nil
	case err := <-errChan:
		// Check if error is related to port occupancy
		errStr := err.Error()
		if contains(errStr, "port") && contains(errStr, "occupied") {
			fmt.Fprintf(os.Stderr, "\n[ERROR] Port %s is occupied.\n", CallbackPort)
			fmt.Fprintf(os.Stderr, "Please close the program using this port and try again.\n")
			fmt.Fprintf(os.Stderr, "You can check which process is using the port with:\n")
			fmt.Fprintf(os.Stderr, "  - macOS/Linux: lsof -i :%s\n", CallbackPort)
			fmt.Fprintf(os.Stderr, "  - Windows: netstat -ano | findstr :%s\n", CallbackPort)
			return fmt.Errorf("port %s is occupied", CallbackPort)
		}
		return fmt.Errorf("authentication failed: %v", err)
	case <-ctx.Done():
		return fmt.Errorf("authentication timeout: please try again")
	}
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// contains checks if a string contains a substring (case-insensitive)
func contains(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}
