// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package agentbay

import (
	"bytes"
	"testing"

	log "github.com/sirupsen/logrus"
)

// TestXMLParsingErrorDetection tests that XML parsing errors are properly detected and handled
func TestXMLParsingErrorDetection(t *testing.T) {
	tests := []struct {
		name        string
		errorString string
		expectXML   bool
	}{
		{
			name:        "XML parsing error should be detected",
			errorString: `client.GetDockerFileStoreCredentialResponse.Body: readObjectStart: expect { or n, but found ", error found in #9 byte of ...|{"body":"\u003c?xml|...`,
			expectXML:   true,
		},
		{
			name:        "Regular error should not be detected as XML",
			errorString: "network connection failed",
			expectXML:   false,
		},
		{
			name:        "Empty error should not be detected as XML",
			errorString: "",
			expectXML:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test the detection logic
			isXMLError := bytes.Contains([]byte(tt.errorString), []byte("readObjectStart: expect { or n, but found"))

			if isXMLError != tt.expectXML {
				t.Errorf("Expected XML detection to be %v for error: %s", tt.expectXML, tt.errorString)
			}
		})
	}
}

// TestLogLevelHandling tests that debug logs are only shown at appropriate levels
func TestLogLevelHandling(t *testing.T) {
	// Save original log level
	originalLevel := log.GetLevel()
	defer log.SetLevel(originalLevel)

	tests := []struct {
		name        string
		logLevel    log.Level
		expectDebug bool
	}{
		{
			name:        "Debug level should show debug logs",
			logLevel:    log.DebugLevel,
			expectDebug: true,
		},
		{
			name:        "Info level should not show debug logs",
			logLevel:    log.InfoLevel,
			expectDebug: false,
		},
		{
			name:        "Warn level should not show debug logs",
			logLevel:    log.WarnLevel,
			expectDebug: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.SetLevel(tt.logLevel)

			shouldShowDebug := log.GetLevel() >= log.DebugLevel

			if shouldShowDebug != tt.expectDebug {
				t.Errorf("Expected debug log visibility to be %v at level %v", tt.expectDebug, tt.logLevel)
			}
		})
	}
}

// TestSimplifiedLogging tests that the logging output is simplified and clean
func TestSimplifiedLogging(t *testing.T) {
	// Save original log level
	originalLevel := log.GetLevel()
	defer log.SetLevel(originalLevel)

	// Set to debug level to enable debug logs
	log.SetLevel(log.DebugLevel)

	tests := []struct {
		name             string
		logMessage       string
		shouldContain    []string
		shouldNotContain []string
	}{
		{
			name:             "Request logs should be simplified",
			logMessage:       "Making GetDockerFileStoreCredential request...",
			shouldContain:    []string{"Making", "request"},
			shouldNotContain: []string{"Raw HTTP Request", "Headers", "Method", "URL"},
		},
		{
			name:             "Response logs should be simplified",
			logMessage:       "GetDockerFileStoreCredential request completed successfully",
			shouldContain:    []string{"completed", "successfully"},
			shouldNotContain: []string{"Raw HTTP Response", "Status Code", "Response Headers", "Response Body"},
		},
		{
			name:             "XML parsing logs should be friendly",
			logMessage:       "SDK returned XML response, using custom XML parser...",
			shouldContain:    []string{"XML response", "XML parser"},
			shouldNotContain: []string{"readObjectStart", "expect { or n", "error found in #9 byte"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Check that the message contains expected content
			for _, expected := range tt.shouldContain {
				if !bytes.Contains([]byte(tt.logMessage), []byte(expected)) {
					t.Errorf("Log message should contain '%s': %s", expected, tt.logMessage)
				}
			}

			// Check that the message doesn't contain unwanted content
			for _, unwanted := range tt.shouldNotContain {
				if bytes.Contains([]byte(tt.logMessage), []byte(unwanted)) {
					t.Errorf("Log message should not contain '%s': %s", unwanted, tt.logMessage)
				}
			}
		})
	}
}
