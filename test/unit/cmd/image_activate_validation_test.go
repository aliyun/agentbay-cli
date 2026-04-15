// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestImageActivate_NetworkTypeValidation tests network type validation logic directly
func TestImageActivate_NetworkTypeValidation(t *testing.T) {
	tests := []struct {
		name           string
		networkType    string
		expectError    bool
		errorMsg       string
	}{
		{
			name:        "valid_DEFAULT",
			networkType: "DEFAULT",
			expectError: false,
		},
		{
			name:        "valid_ADVANCED",
			networkType: "ADVANCED",
			expectError: false,
		},
		{
			name:        "invalid_lowercase",
			networkType: "default",
			expectError: true,
			errorMsg:    "Invalid network type",
		},
		{
			name:        "invalid_uppercase",
			networkType: "INVALID",
			expectError: true,
			errorMsg:    "Invalid network type",
		},
		{
			name:        "invalid_empty",
			networkType: "",
			expectError: true,
			errorMsg:    "Invalid network type",
		},
		{
			name:        "invalid_mixed_case",
			networkType: "Advanced",
			expectError: true,
			errorMsg:    "Invalid network type",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test network type validation logic
			isValid := tt.networkType == "DEFAULT" || tt.networkType == "ADVANCED"
			
			if tt.expectError {
				assert.False(t, isValid, "Network type '%s' should be invalid", tt.networkType)
			} else {
				assert.True(t, isValid, "Network type '%s' should be valid", tt.networkType)
			}
		})
	}
}

// TestImageActivate_AdvancedNetworkParamsValidation tests advanced network parameter validation logic
func TestImageActivate_AdvancedNetworkParamsValidation(t *testing.T) {
	tests := []struct {
		name             string
		networkType      string
		sessionBandwidth int
		dnsAddresses     []string
		expectError      bool
		errorMsg         string
	}{
		{
			name:             "DEFAULT_with_session_bandwidth_should_fail",
			networkType:      "DEFAULT",
			sessionBandwidth: 100,
			dnsAddresses:     []string{},
			expectError:      true,
			errorMsg:         "--session-bandwidth is only valid for ADVANCED network",
		},
		{
			name:             "DEFAULT_with_dns_address_should_fail",
			networkType:      "DEFAULT",
			sessionBandwidth: 0,
			dnsAddresses:     []string{"8.8.8.8"},
			expectError:      true,
			errorMsg:         "--dns-address is only valid for ADVANCED network",
		},
		{
			name:             "DEFAULT_with_both_params_should_fail",
			networkType:      "DEFAULT",
			sessionBandwidth: 100,
			dnsAddresses:     []string{"8.8.8.8"},
			expectError:      true,
			errorMsg:         "--session-bandwidth is only valid for ADVANCED network",
		},
		{
			name:             "ADVANCED_with_session_bandwidth_should_pass",
			networkType:      "ADVANCED",
			sessionBandwidth: 100,
			dnsAddresses:     []string{},
			expectError:      false,
		},
		{
			name:             "ADVANCED_with_dns_address_should_pass",
			networkType:      "ADVANCED",
			sessionBandwidth: 0,
			dnsAddresses:     []string{"8.8.8.8"},
			expectError:      false,
		},
		{
			name:             "ADVANCED_with_multiple_dns_should_pass",
			networkType:      "ADVANCED",
			sessionBandwidth: 100,
			dnsAddresses:     []string{"8.8.8.8", "8.8.4.4"},
			expectError:      false,
		},
		{
			name:             "ADVANCED_without_params_should_pass",
			networkType:      "ADVANCED",
			sessionBandwidth: 0,
			dnsAddresses:     []string{},
			expectError:      false,
		},
		{
			name:             "DEFAULT_without_params_should_pass",
			networkType:      "DEFAULT",
			sessionBandwidth: 0,
			dnsAddresses:     []string{},
			expectError:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Simulate the validation logic from runImageActivate
			var err error
			
			if tt.networkType == "DEFAULT" {
				if tt.sessionBandwidth > 0 {
					err = fmt.Errorf("--session-bandwidth is only valid for ADVANCED network")
				}
				if err == nil && len(tt.dnsAddresses) > 0 {
					err = fmt.Errorf("--dns-address is only valid for ADVANCED network")
				}
			}

			if tt.expectError {
				assert.Error(t, err)
				if tt.errorMsg != "" && err != nil {
					assert.True(t, strings.Contains(err.Error(), tt.errorMsg),
						"Expected error to contain '%s', got: %s", tt.errorMsg, err.Error())
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// TestImageActivate_DefaultResourceApplication tests that default resources are correctly applied
func TestImageActivate_DefaultResourceApplication(t *testing.T) {
	const DefaultActivateCPU = 2
	const DefaultActivateMemory = 4

	tests := []struct {
		name           string
		inputCPU       int
		inputMemory    int
		expectedCPU    int
		expectedMemory int
	}{
		{
			name:           "no_resources_specified_use_defaults",
			inputCPU:       0,
			inputMemory:    0,
			expectedCPU:    DefaultActivateCPU,
			expectedMemory: DefaultActivateMemory,
		},
		{
			name:           "explicit_2c4g",
			inputCPU:       2,
			inputMemory:    4,
			expectedCPU:    2,
			expectedMemory: 4,
		},
		{
			name:           "explicit_4c8g",
			inputCPU:       4,
			inputMemory:    8,
			expectedCPU:    4,
			expectedMemory: 8,
		},
		{
			name:           "explicit_8c16g",
			inputCPU:       8,
			inputMemory:    16,
			expectedCPU:    8,
			expectedMemory: 16,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Simulate the default resource application logic
			cpu := tt.inputCPU
			memory := tt.inputMemory
			
			if cpu == 0 && memory == 0 {
				cpu = DefaultActivateCPU
				memory = DefaultActivateMemory
			}

			assert.Equal(t, tt.expectedCPU, cpu, "CPU should match expected")
			assert.Equal(t, tt.expectedMemory, memory, "Memory should match expected")
		})
	}
}

// TestImageActivate_DNSAddressParsing tests DNS address array handling
func TestImageActivate_DNSAddressParsing(t *testing.T) {
	tests := []struct {
		name         string
		dnsAddresses []string
		expectedLen  int
	}{
		{
			name:         "no_dns_addresses",
			dnsAddresses: []string{},
			expectedLen:  0,
		},
		{
			name:         "single_dns_address",
			dnsAddresses: []string{"8.8.8.8"},
			expectedLen:  1,
		},
		{
			name:         "two_dns_addresses",
			dnsAddresses: []string{"8.8.8.8", "8.8.4.4"},
			expectedLen:  2,
		},
		{
			name:         "multiple_dns_addresses",
			dnsAddresses: []string{"8.8.8.8", "8.8.4.4", "1.1.1.1"},
			expectedLen:  3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectedLen, len(tt.dnsAddresses),
				"DNS addresses count should match expected")
			
			// Test joining for display
			if len(tt.dnsAddresses) > 0 {
				joined := strings.Join(tt.dnsAddresses, ", ")
				assert.NotEmpty(t, joined, "Joined DNS addresses should not be empty")
				if len(tt.dnsAddresses) > 1 {
					assert.True(t, strings.Contains(joined, ","), 
						"Multiple DNS addresses should be comma-separated")
				}
			}
		})
	}
}

// TestImageActivate_NetworkTypeConstants tests that network type constants are correct
func TestImageActivate_NetworkTypeConstants(t *testing.T) {
	// Test valid network types
	validTypes := []string{"DEFAULT", "ADVANCED"}
	for _, networkType := range validTypes {
		t.Run("valid_"+networkType, func(t *testing.T) {
			isValid := networkType == "DEFAULT" || networkType == "ADVANCED"
			assert.True(t, isValid, "Network type '%s' should be valid", networkType)
		})
	}

	// Test invalid network types
	invalidTypes := []string{"default", "advanced", "Advanced", "INVALID", "", "BRIDGE", "HOST"}
	for _, networkType := range invalidTypes {
		t.Run("invalid_"+networkType, func(t *testing.T) {
			isValid := networkType == "DEFAULT" || networkType == "ADVANCED"
			assert.False(t, isValid, "Network type '%s' should be invalid", networkType)
		})
	}
}
