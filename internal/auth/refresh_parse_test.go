// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package auth

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseOAuthExpiresIn(t *testing.T) {
	t.Run("string number", func(t *testing.T) {
		n, err := parseOAuthExpiresIn(json.RawMessage(`"3599"`))
		require.NoError(t, err)
		assert.Equal(t, 3599, n)
	})

	t.Run("JSON number", func(t *testing.T) {
		n, err := parseOAuthExpiresIn(json.RawMessage(`3600`))
		require.NoError(t, err)
		assert.Equal(t, 3600, n)
	})

	t.Run("missing", func(t *testing.T) {
		_, err := parseOAuthExpiresIn(json.RawMessage(``))
		assert.Error(t, err)
	})

	t.Run("null", func(t *testing.T) {
		_, err := parseOAuthExpiresIn(json.RawMessage(`null`))
		assert.Error(t, err)
	})
}
