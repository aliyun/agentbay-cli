// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package config_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/agentbay/agentbay-cli/internal/config"
)

func TestAccessKeyFromEnv(t *testing.T) {
	origID := os.Getenv(config.EnvAccessKeyID)
	origSec := os.Getenv(config.EnvAccessKeySecret)
	origSess := os.Getenv(config.EnvAccessKeySessionToken)
	t.Cleanup(func() {
		restoreEnv(config.EnvAccessKeyID, origID)
		restoreEnv(config.EnvAccessKeySecret, origSec)
		restoreEnv(config.EnvAccessKeySessionToken, origSess)
	})

	require.NoError(t, os.Unsetenv(config.EnvAccessKeyID))
	require.NoError(t, os.Unsetenv(config.EnvAccessKeySecret))
	require.NoError(t, os.Unsetenv(config.EnvAccessKeySessionToken))
	assert.False(t, config.HasAccessKeyFromEnv())

	t.Setenv(config.EnvAccessKeyID, "ak-id")
	t.Setenv(config.EnvAccessKeySecret, "ak-secret")
	assert.True(t, config.HasAccessKeyFromEnv())
	id, sec, sess, ok := config.AccessKeyFromEnv()
	assert.True(t, ok)
	assert.Equal(t, "ak-id", id)
	assert.Equal(t, "ak-secret", sec)
	assert.Equal(t, "", sess)

	t.Setenv(config.EnvAccessKeySessionToken, "sts")
	id, sec, sess, ok = config.AccessKeyFromEnv()
	assert.True(t, ok)
	assert.Equal(t, "sts", sess)
}

func restoreEnv(key, val string) {
	if val == "" {
		_ = os.Unsetenv(key)
	} else {
		_ = os.Setenv(key, val)
	}
}
