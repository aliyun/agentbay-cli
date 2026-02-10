package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ParseCOPYADDSources(dockerfileContent []byte, contextDir string) ([]string, error) {
	lines := SplitDockerfileLines(dockerfileContent)
	seen := make(map[string]struct{})
	var out []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		upper := strings.ToUpper(line)
		if !strings.HasPrefix(upper, "COPY ") && !strings.HasPrefix(upper, "ADD ") {
			continue
		}
		idx := strings.Index(line, " ")
		if idx < 0 {
			continue
		}
		rest := strings.TrimSpace(line[idx+1:])
		tokens, err := TokenizeInstruction(rest)
		if err != nil {
			continue
		}
		if len(tokens) < 2 {
			continue
		}
		dest := tokens[len(tokens)-1]
		sources := tokens[:len(tokens)-1]
		_ = dest

		var i int
		for _, t := range sources {
			if strings.HasPrefix(t, "--") {
				if strings.HasPrefix(strings.ToLower(t), "--from=") {
					break
				}
				continue
			}
			sources[i] = t
			i++
		}
		sources = sources[:i]
		if len(sources) == 0 {
			continue
		}
		if strings.HasPrefix(strings.ToUpper(line), "ADD ") && IsURL(sources[0]) {
			continue
		}
		for _, src := range sources {
			absPaths, err := ExpandSource(contextDir, src)
			if err != nil {
				return nil, err
			}
			for _, p := range absPaths {
				if _, ok := seen[p]; ok {
					continue
				}
				seen[p] = struct{}{}
				out = append(out, p)
			}
		}
	}
	return out, nil
}

func SplitDockerfileLines(content []byte) []string {
	var lines []string
	s := string(content)
	for {
		idx := strings.IndexAny(s, "\n\r")
		if idx < 0 {
			line := strings.TrimSpace(s)
			if line != "" {
				lines = append(lines, line)
			}
			break
		}
		line := s[:idx]
		s = s[idx+1:]
		if len(s) > 0 && (s[0] == '\n' || s[0] == '\r') {
			s = s[1:]
		}
		line = strings.TrimRight(line, "\r\n")
		for strings.HasSuffix(line, "\\") {
			line = strings.TrimSuffix(line, "\\")
			line = strings.TrimRight(line, " \t")
			next := strings.IndexAny(s, "\n\r")
			if next < 0 {
				line = line + " " + strings.TrimSpace(s)
				s = ""
				break
			}
			line = line + " " + strings.TrimSpace(s[:next])
			s = s[next+1:]
			if len(s) > 0 && (s[0] == '\n' || s[0] == '\r') {
				s = s[1:]
			}
		}
		lines = append(lines, strings.TrimSpace(line))
	}
	return lines
}

func TokenizeInstruction(rest string) ([]string, error) {
	if strings.HasPrefix(rest, "[") {
		return tokenizeJSONArray(rest)
	}
	var tokens []string
	for rest != "" {
		rest = strings.TrimLeft(rest, " \t")
		if rest == "" {
			break
		}
		if rest[0] == '"' || rest[0] == '\'' {
			end := strings.IndexByte(rest[1:], rest[0])
			if end < 0 {
				return nil, fmt.Errorf("unclosed quote")
			}
			tokens = append(tokens, rest[1:end+1])
			rest = rest[end+2:]
			continue
		}
		i := 0
		for i < len(rest) && rest[i] != ' ' && rest[i] != '\t' {
			i++
		}
		tokens = append(tokens, rest[:i])
		rest = rest[i:]
	}
	return tokens, nil
}

func tokenizeJSONArray(rest string) ([]string, error) {
	rest = strings.TrimSpace(rest)
	if !strings.HasPrefix(rest, "[") {
		return nil, fmt.Errorf("not json array")
	}
	rest = strings.TrimSpace(rest[1:])
	var tokens []string
	for {
		rest = strings.TrimLeft(rest, " \t,")
		if rest == "" || rest[0] == ']' {
			break
		}
		if rest[0] != '"' && rest[0] != '\'' {
			return nil, fmt.Errorf("expected quoted string in array")
		}
		quote := rest[0]
		end := strings.IndexByte(rest[1:], quote)
		if end < 0 {
			return nil, fmt.Errorf("unclosed quote")
		}
		tokens = append(tokens, rest[1:end+1])
		rest = rest[end+2:]
	}
	return tokens, nil
}

func ExpandSource(contextDir, source string) ([]string, error) {
	source = filepath.Clean(source)
	if filepath.IsAbs(source) {
		return nil, fmt.Errorf("absolute source path not supported: %s", source)
	}
	pattern := filepath.Clean(filepath.Join(contextDir, source))
	rel, err := filepath.Rel(contextDir, pattern)
	if err != nil || strings.HasPrefix(rel, "..") {
		return nil, fmt.Errorf("source path escapes context: %s", source)
	}
	if strings.Contains(source, "*") {
		matches, err := filepath.Glob(pattern)
		if err != nil {
			return nil, err
		}
		var files []string
		for _, m := range matches {
			info, err := os.Stat(m)
			if err != nil {
				continue
			}
			if info.IsDir() {
				sub, err := walkFiles(m)
				if err != nil {
					return nil, err
				}
				files = append(files, sub...)
			} else {
				files = append(files, m)
			}
		}
		return files, nil
	}
	info, err := os.Stat(pattern)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("source not found: %s", source)
		}
		return nil, err
	}
	if info.IsDir() {
		return walkFiles(pattern)
	}
	return []string{pattern}, nil
}

func walkFiles(dir string) ([]string, error) {
	var out []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		out = append(out, path)
		return nil
	})
	return out, err
}

func RelativePathForUpload(contextDir, absolutePath string) (string, error) {
	rel, err := filepath.Rel(contextDir, absolutePath)
	if err != nil {
		return "", err
	}
	return filepath.ToSlash(rel), nil
}

func IsURL(s string) bool {
	return strings.HasPrefix(s, "http://") || strings.HasPrefix(s, "https://")
}
