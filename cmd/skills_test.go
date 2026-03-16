// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"archive/zip"
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseSkillFrontmatter(t *testing.T) {
	tests := []struct {
		name        string
		content     []byte
		wantName    string
		wantDesc    string
		wantErr     bool
		errContains string
	}{
		{
			name:     "name only",
			content:  []byte("name: my-skill\n"),
			wantName: "my-skill",
			wantDesc: "",
		},
		{
			name:     "name and description",
			content:  []byte("name: foo\ndescription: A test skill\n"),
			wantName: "foo",
			wantDesc: "A test skill",
		},
		{
			name:     "name with spaces trimmed",
			content:  []byte("name:   spaced-name  \n"),
			wantName: "spaced-name",
		},
		{
			name:        "missing name",
			content:     []byte("description: only desc\n"),
			wantErr:     true,
			errContains: "name:",
		},
		{
			name:        "empty content",
			content:     []byte(""),
			wantErr:     true,
			errContains: "name:",
		},
		{
			name:     "description with colon inside",
			content:  []byte("name: x\ndescription: Use when: doing something\n"),
			wantName: "x",
			wantDesc: "Use when: doing something",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotName, gotDesc, err := parseSkillFrontmatter(tt.content)
			if tt.wantErr {
				require.Error(t, err)
				if tt.errContains != "" {
					assert.Contains(t, err.Error(), tt.errContains)
				}
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.wantName, gotName)
			assert.Equal(t, tt.wantDesc, gotDesc)
		})
	}
}

func TestParseOSSBucketAndPath(t *testing.T) {
	tests := []struct {
		name        string
		ossURL      string
		wantBucket  string
		wantPath    string
		wantErr     bool
		errContains string
	}{
		{
			name:       "valid OSS URL",
			ossURL:     "https://agentbay-market-skill.oss-cn-hangzhou.aliyuncs.com/prefix/key?Expires=123",
			wantBucket: "agentbay-market-skill",
			wantPath:   "prefix/key",
		},
		{
			name:       "root path",
			ossURL:     "https://mybucket.oss-cn-shanghai.aliyuncs.com/file.zip",
			wantBucket: "mybucket",
			wantPath:   "file.zip",
		},
		{
			name:        "missing host",
			ossURL:      "https:///path",
			wantErr:     true,
			errContains: "host",
		},
		{
			name:        "invalid URL",
			ossURL:      ":not-a-url",
			wantErr:     true,
		},
		{
			name:        "host not in bucket.oss form",
			ossURL:      "https://singlelabel/path",
			wantErr:     true,
			errContains: "bucket",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBucket, gotPath, err := parseOSSBucketAndPath(tt.ossURL)
			if tt.wantErr {
				require.Error(t, err)
				if tt.errContains != "" {
					assert.Contains(t, err.Error(), tt.errContains)
				}
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.wantBucket, gotBucket)
			assert.Equal(t, tt.wantPath, gotPath)
		})
	}
}

func TestZipSkillDir(t *testing.T) {
	tempDir := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(tempDir, "SKILL.md"), []byte("name: test\n"), 0644))
	require.NoError(t, os.WriteFile(filepath.Join(tempDir, "extra.txt"), []byte("data"), 0644))
	subDir := filepath.Join(tempDir, "sub")
	require.NoError(t, os.MkdirAll(subDir, 0755))
	require.NoError(t, os.WriteFile(filepath.Join(subDir, "nested.txt"), []byte("nested"), 0644))

	buf, err := zipSkillDir(tempDir)
	require.NoError(t, err)
	require.NotNil(t, buf)

	r, err := zip.NewReader(bytes.NewReader(buf.Bytes()), int64(buf.Len()))
	require.NoError(t, err)

	names := make([]string, 0, len(r.File))
	for _, f := range r.File {
		names = append(names, f.Name)
	}
	assert.Contains(t, names, "SKILL.md")
	assert.Contains(t, names, "extra.txt")
	assert.Contains(t, names, "sub/nested.txt")
	assert.Len(t, names, 3)
}

func TestZipSkillDir_nonexistent(t *testing.T) {
	_, err := zipSkillDir(filepath.Join(t.TempDir(), "nonexistent"))
	require.Error(t, err)
}

func TestStrPtr(t *testing.T) {
	assert.Equal(t, "", strPtr(nil))
	s := "hello"
	assert.Equal(t, "hello", strPtr(&s))
}

func TestRunSkillsPush_validationOnly(t *testing.T) {
	// Tests that runSkillsPush fails before any API call (config/network).
	// printErrorMessage returns "command failed"; detailed text is only on stderr.
	tests := []struct {
		name  string
		setup func(t *testing.T) string // returns skill dir path
	}{
		{"directory does not exist", func(t *testing.T) string {
			return filepath.Join(t.TempDir(), "missing")
		}},
		{"not a directory", func(t *testing.T) string {
			f := filepath.Join(t.TempDir(), "file")
			require.NoError(t, os.WriteFile(f, []byte("x"), 0644))
			return f
		}},
		{"SKILL.md not found", func(t *testing.T) string {
			return t.TempDir()
		}},
		{"invalid frontmatter missing name", func(t *testing.T) string {
			d := t.TempDir()
			require.NoError(t, os.WriteFile(filepath.Join(d, skillFileName), []byte("description: only\n"), 0644))
			return d
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			skillDir := tt.setup(t)
			err := runSkillsPush(skillsPushCmd, []string{skillDir})
			require.Error(t, err)
			// printErrorMessage returns "command failed"; other paths return wrapped errors
			assert.True(t, err.Error() == "command failed" || strings.Contains(err.Error(), "skill directory") ||
				strings.Contains(err.Error(), "reading") || strings.Contains(err.Error(), "load config"),
				"unexpected error: %s", err.Error())
		})
	}
}

func TestRunSkillsList(t *testing.T) {
	// runSkillsList is a placeholder that prints to stderr and returns nil.
	err := runSkillsList(skillsListCmd, nil)
	require.NoError(t, err)
}

func TestRunSkillsGroupShow(t *testing.T) {
	// runSkillsGroupShow is a placeholder that prints to stderr and returns nil.
	err := runSkillsGroupShow(skillsGroupShowCmd, []string{"some-group-id"})
	require.NoError(t, err)
}
