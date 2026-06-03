#!/bin/bash
# Backfill GitHub Release Notes from CHANGELOG.md.
#
# Under the v3 bilingual changelog pipeline, CHANGELOG.md is the single
# upstream source of truth. This script extracts a per-version section
# from CHANGELOG.md and pushes it to the corresponding GitHub Release as
# the release body. Use this when:
#   - You modified CHANGELOG.md (e.g. refined a translation) and want to
#     refresh an already-published release body.
#   - You historically released with the old single-language pipeline and
#     want to refresh release bodies after backfilling Chinese into
#     CHANGELOG.md.
#
# Prerequisites:
#   - gh CLI authenticated with repo write access
#   - Run from the project root directory
#   - CHANGELOG.md must contain a section for each tag you want to refresh
#     (lines starting with `## [<VERSION>]`). Tags without a matching
#     section are SKIPPED, not failed.
#
# Usage:
#   ./scripts/backfill-release-notes.sh                 # Refresh all v* releases
#   ./scripts/backfill-release-notes.sh --dry-run       # Preview without writing
#   ./scripts/backfill-release-notes.sh --tag v0.2.8    # Refresh a single release

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"

DRY_RUN=false
SINGLE_TAG=""

# Parse arguments
while [[ $# -gt 0 ]]; do
    case "$1" in
        --dry-run)
            DRY_RUN=true
            shift
            ;;
        --tag)
            SINGLE_TAG="$2"
            shift 2
            ;;
        -h|--help)
            echo "Usage: $(basename "$0") [--dry-run] [--tag TAG]"
            echo ""
            echo "  --dry-run    Preview changes without updating releases"
            echo "  --tag TAG    Update a single release (e.g., --tag v0.2.8)"
            echo "  -h, --help   Show this help"
            exit 0
            ;;
        *)
            echo "Unknown option: $1"
            exit 1
            ;;
    esac
done

# Check prerequisites
if ! command -v gh >/dev/null 2>&1; then
    echo "ERROR: gh CLI is required. Install with: brew install gh"
    exit 1
fi

EXTRACT_SCRIPT="$SCRIPT_DIR/extract-changelog-section.sh"
if [[ ! -x "$EXTRACT_SCRIPT" ]]; then
    echo "ERROR: $EXTRACT_SCRIPT is missing or not executable"
    exit 1
fi

CHANGELOG="$PROJECT_DIR/CHANGELOG.md"
if [[ ! -f "$CHANGELOG" ]]; then
    echo "ERROR: $CHANGELOG not found"
    exit 1
fi

cd "$PROJECT_DIR"

# Get tags to process
if [[ -n "$SINGLE_TAG" ]]; then
    TAGS=("$SINGLE_TAG")
else
    # Get all version tags sorted by version.
    # Use a read loop instead of `mapfile` for bash 3.2 (macOS default) compat.
    TAGS=()
    while IFS= read -r line; do
        TAGS+=("$line")
    done < <(git tag -l 'v*' --sort=v:refname)
fi

echo "=========================================="
echo " Backfill GitHub Release Notes"
echo "=========================================="
echo ""
echo "Tags to process: ${#TAGS[@]}"
echo "Dry run: $DRY_RUN"
echo ""

UPDATED=0
SKIPPED=0
FAILED=0

for tag in "${TAGS[@]}"; do
    echo "--- Processing $tag ---"

    # Check if release exists
    if ! gh release view "$tag" >/dev/null 2>&1; then
        echo "  SKIP: Release $tag does not exist on GitHub"
        SKIPPED=$((SKIPPED + 1))
        continue
    fi

    # Extract changelog section for this version from CHANGELOG.md
    NOTES_FILE="/tmp/release-notes-${tag}.md"
    VERSION_NO_V="${tag#v}"
    if ! bash "$EXTRACT_SCRIPT" "$VERSION_NO_V" "$CHANGELOG" > "$NOTES_FILE" 2>/dev/null; then
        echo "  SKIP: CHANGELOG.md has no section for $tag"
        SKIPPED=$((SKIPPED + 1))
        rm -f "$NOTES_FILE"
        continue
    fi

    if [[ ! -s "$NOTES_FILE" ]]; then
        echo "  SKIP: extracted section is empty for $tag"
        SKIPPED=$((SKIPPED + 1))
        rm -f "$NOTES_FILE"
        continue
    fi

    # Append installation instructions
    cat >> "$NOTES_FILE" << 'INSTALL_EOF'

## Installation
Once merged into Homebrew core, install with:
```bash
brew install agentbay
```

## Manual Installation
Download the appropriate binary for your platform from the assets below.
INSTALL_EOF

    if $DRY_RUN; then
        echo "  [DRY RUN] Would update release $tag with:"
        echo "  ---"
        cat "$NOTES_FILE"
        echo "  ---"
    else
        if gh release edit "$tag" --notes-file "$NOTES_FILE"; then
            echo "  OK: Updated release notes for $tag"
            UPDATED=$((UPDATED + 1))
        else
            echo "  ERROR: Failed to update release $tag"
            FAILED=$((FAILED + 1))
        fi
    fi

    rm -f "$NOTES_FILE"
done

echo ""
echo "=========================================="
echo " Summary"
echo "=========================================="
echo "  Updated: $UPDATED"
echo "  Skipped: $SKIPPED"
echo "  Failed:  $FAILED"
echo "  Total:   ${#TAGS[@]}"
if $DRY_RUN; then
    echo ""
    echo "  [DRY RUN - No changes were made]"
fi
