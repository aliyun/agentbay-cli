// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd_test

import (
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/agentbay/agentbay-cli/cmd"
)

func TestParseCOPYADDSources(t *testing.T) {
	tempDir := t.TempDir()

	// Create test files
	appPath := filepath.Join(tempDir, "app.py")
	require.NoError(t, os.WriteFile(appPath, []byte("print('hello')"), 0644))
	reqPath := filepath.Join(tempDir, "requirements.txt")
	require.NoError(t, os.WriteFile(reqPath, []byte("flask"), 0644))
	subDir := filepath.Join(tempDir, "code")
	require.NoError(t, os.MkdirAll(subDir, 0755))
	codePath := filepath.Join(subDir, "main.py")
	require.NoError(t, os.WriteFile(codePath, []byte("pass"), 0644))

	tests := []struct {
		name        string
		dockerfile  string
		wantPaths   []string
		wantErr     bool
		errContains string
	}{
		{
			name: "no COPY or ADD",
			dockerfile: `FROM ubuntu:20.04
RUN echo "test"
`,
			wantPaths: nil,
		},
		{
			name: "single COPY file",
			dockerfile: `FROM ubuntu:20.04
COPY app.py /app/
`,
			wantPaths: []string{appPath},
		},
		{
			name: "single ADD file",
			dockerfile: `FROM ubuntu:20.04
ADD requirements.txt /app/
`,
			wantPaths: []string{reqPath},
		},
		{
			name: "multiple COPY sources",
			dockerfile: `FROM ubuntu:20.04
COPY app.py requirements.txt /app/
`,
			wantPaths: []string{appPath, reqPath},
		},
		{
			name: "COPY with subdirectory",
			dockerfile: `FROM ubuntu:20.04
COPY code/main.py /app/
`,
			wantPaths: []string{codePath},
		},
		{
			name: "multiple COPY instructions",
			dockerfile: `FROM ubuntu:20.04
COPY app.py /app/
COPY requirements.txt /app/
`,
			wantPaths: []string{appPath, reqPath},
		},
		{
			name: "COPY with --chown",
			dockerfile: `FROM ubuntu:20.04
COPY --chown=root:root app.py /app/
`,
			wantPaths: []string{appPath},
		},
		{
			name: "ignores comments",
			dockerfile: `FROM ubuntu:20.04
# COPY ignored.py /app/
COPY app.py /app/
`,
			wantPaths: []string{appPath},
		},
		{
			name: "source not found",
			dockerfile: `FROM ubuntu:20.04
COPY nonexistent.py /app/
`,
			wantErr:     true,
			errContains: "source not found",
		},
		{
			name: "absolute path not supported",
			dockerfile: `FROM ubuntu:20.04
COPY /absolute/path /app/
`,
			wantErr:     true,
			errContains: "absolute source path not supported",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cmd.ParseCOPYADDSources([]byte(tt.dockerfile), tempDir)
			if tt.wantErr {
				require.Error(t, err)
				if tt.errContains != "" {
					assert.Contains(t, err.Error(), tt.errContains)
				}
				return
			}
			require.NoError(t, err)
			sort.Strings(got)
			sort.Strings(tt.wantPaths)
			assert.Equal(t, tt.wantPaths, got)
		})
	}
}

func TestParseCOPYADDSources_WithWildcard(t *testing.T) {
	tempDir := t.TempDir()

	require.NoError(t, os.WriteFile(filepath.Join(tempDir, "a.py"), []byte("a"), 0644))
	require.NoError(t, os.WriteFile(filepath.Join(tempDir, "b.py"), []byte("b"), 0644))
	require.NoError(t, os.WriteFile(filepath.Join(tempDir, "c.txt"), []byte("c"), 0644))

	dockerfile := `FROM ubuntu:20.04
COPY *.py /app/
`
	got, err := cmd.ParseCOPYADDSources([]byte(dockerfile), tempDir)
	require.NoError(t, err)
	require.Len(t, got, 2)
	sort.Strings(got)
	assert.Contains(t, got, filepath.Join(tempDir, "a.py"))
	assert.Contains(t, got, filepath.Join(tempDir, "b.py"))
	assert.NotContains(t, got, filepath.Join(tempDir, "c.txt"))
}

func TestParseCOPYADDSources_ADDWithURL(t *testing.T) {
	tempDir := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(tempDir, "local.txt"), []byte("x"), 0644))

	dockerfile := `FROM ubuntu:20.04
ADD https://example.com/file.tar.gz /tmp/
COPY local.txt /app/
`
	got, err := cmd.ParseCOPYADDSources([]byte(dockerfile), tempDir)
	require.NoError(t, err)
	assert.Equal(t, []string{filepath.Join(tempDir, "local.txt")}, got)
}

func TestExpandSource(t *testing.T) {
	tempDir := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(tempDir, "file.txt"), []byte("x"), 0644))
	subDir := filepath.Join(tempDir, "dir")
	require.NoError(t, os.MkdirAll(subDir, 0755))
	require.NoError(t, os.WriteFile(filepath.Join(subDir, "nested.txt"), []byte("y"), 0644))

	tests := []struct {
		name        string
		source      string
		wantErr     bool
		errContains string
	}{
		{name: "single file", source: "file.txt"},
		{name: "single file in subdir", source: "dir/nested.txt"},
		{name: "absolute path", source: "/absolute/path", wantErr: true, errContains: "absolute source path not supported"},
		{name: "path escapes context", source: "../outside", wantErr: true, errContains: "source path escapes context"},
		{name: "nonexistent file", source: "nonexistent.txt", wantErr: true, errContains: "source not found"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cmd.ExpandSource(tempDir, tt.source)
			if tt.wantErr {
				require.Error(t, err)
				if tt.errContains != "" {
					assert.Contains(t, err.Error(), tt.errContains)
				}
				return
			}
			require.NoError(t, err)
			require.NotEmpty(t, got)
		})
	}
}

func TestExpandSource_Directory(t *testing.T) {
	tempDir := t.TempDir()
	subDir := filepath.Join(tempDir, "mydir")
	require.NoError(t, os.MkdirAll(subDir, 0755))
	require.NoError(t, os.WriteFile(filepath.Join(subDir, "a.txt"), []byte("a"), 0644))
	require.NoError(t, os.WriteFile(filepath.Join(subDir, "b.txt"), []byte("b"), 0644))

	got, err := cmd.ExpandSource(tempDir, "mydir")
	require.NoError(t, err)
	require.Len(t, got, 2)
	sort.Strings(got)
	assert.Equal(t, filepath.Join(subDir, "a.txt"), got[0])
	assert.Equal(t, filepath.Join(subDir, "b.txt"), got[1])
}

func TestRelativePathForUpload(t *testing.T) {
	tempDir := t.TempDir()
	subDir := filepath.Join(tempDir, "code")
	require.NoError(t, os.MkdirAll(subDir, 0755))
	filePath := filepath.Join(subDir, "app.py")
	require.NoError(t, os.WriteFile(filePath, []byte("x"), 0644))

	tests := []struct {
		name       string
		contextDir string
		absPath    string
		want       string
		wantErr    bool
	}{
		{name: "file in subdir", contextDir: tempDir, absPath: filePath, want: "code/app.py"},
		{name: "file in context root", contextDir: tempDir, absPath: filepath.Join(tempDir, "root.txt"), want: "root.txt"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cmd.RelativePathForUpload(tt.contextDir, tt.absPath)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSplitDockerfileLines(t *testing.T) {
	tests := []struct {
		name    string
		content string
		want    []string
	}{
		{name: "simple lines", content: "FROM a\nRUN b\n", want: []string{"FROM a", "RUN b"}},
		{name: "line continuation", content: "COPY a b \\\n  c /dest/", want: []string{"COPY a b c /dest/"}},
		{name: "empty and comments", content: "# comment\n\nFROM x\n", want: []string{"# comment", "FROM x"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cmd.SplitDockerfileLines([]byte(tt.content))
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTokenizeInstruction(t *testing.T) {
	tests := []struct {
		name    string
		rest    string
		want    []string
		wantErr bool
	}{
		{name: "simple tokens", rest: "app.py requirements.txt /app/", want: []string{"app.py", "requirements.txt", "/app/"}},
		{name: "quoted path with spaces", rest: `"path with spaces" /dest/`, want: []string{"path with spaces", "/dest/"}},
		{name: "json array", rest: `["a", "b", "c"]`, want: []string{"a", "b", "c"}},
		{name: "unclosed quote", rest: `"unclosed`, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cmd.TokenizeInstruction(tt.rest)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.True(t, reflect.DeepEqual(tt.want, got), "got %v want %v", got, tt.want)
		})
	}
}

func TestIsURL(t *testing.T) {
	assert.True(t, cmd.IsURL("http://example.com"))
	assert.True(t, cmd.IsURL("https://example.com"))
	assert.False(t, cmd.IsURL("local/file.txt"))
	assert.False(t, cmd.IsURL("./relative"))
}
