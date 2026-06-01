#!/bin/bash
# Extract a single version section from CHANGELOG.md.
#
# A "section" is the block from `## [VERSION]` to (but not including) the
# next `## [` line, matching the Keep a Changelog header convention used in
# this project.
#
# Usage:
#   bash scripts/extract-changelog-section.sh <VERSION> [CHANGELOG_PATH]
#     VERSION         e.g. 0.3.0  (no leading "v")
#     CHANGELOG_PATH  defaults to CHANGELOG.md (resolved from current dir)
#
# Behavior:
#   - Section found      → print to stdout, exit 0
#   - Section not found  → empty stdout, error message to stderr, exit 1
#   - File not found     → error message to stderr, exit 2

set -euo pipefail

VERSION="${1:-}"
CHANGELOG="${2:-CHANGELOG.md}"

if [[ -z "$VERSION" ]]; then
    echo "Usage: $(basename "$0") <VERSION> [CHANGELOG_PATH]" >&2
    echo "  e.g. $(basename "$0") 0.3.0" >&2
    exit 2
fi

if [[ ! -f "$CHANGELOG" ]]; then
    echo "ERROR: changelog file not found: $CHANGELOG" >&2
    exit 2
fi

# awk extracts everything from `## [VERSION]` (inclusive) up to the next
# `## [` line (exclusive). If the version is never found, awk produces no
# output and we exit 1.
output="$(awk -v ver="$VERSION" '
    /^## \[/ {
        if ($0 ~ "^## \\[" ver "\\]") {
            printing = 1
        } else if (printing) {
            exit
        }
    }
    printing { print }
' "$CHANGELOG")"

if [[ -z "$output" ]]; then
    echo "ERROR: no section found for version [$VERSION] in $CHANGELOG" >&2
    exit 1
fi

printf '%s\n' "$output"
