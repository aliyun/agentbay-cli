#!/bin/bash
# Backfill GitHub Release Notes with auto-generated changelog content
#
# This is a one-time script to update existing GitHub Releases (v0.1.0 ~ v0.2.8)
# that have placeholder release notes with actual changelog content.
#
# Prerequisites:
#   - git-cliff installed (brew install git-cliff)
#   - gh CLI authenticated with repo write access
#   - Run from the project root directory
#
# Usage:
#   ./scripts/backfill-release-notes.sh           # Update all releases
#   ./scripts/backfill-release-notes.sh --dry-run  # Preview without making changes
#   ./scripts/backfill-release-notes.sh --tag v0.2.8  # Update a single release

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
if ! command -v git-cliff >/dev/null 2>&1; then
    echo "ERROR: git-cliff is required. Install with: brew install git-cliff"
    exit 1
fi

if ! command -v gh >/dev/null 2>&1; then
    echo "ERROR: gh CLI is required. Install with: brew install gh"
    exit 1
fi

cd "$PROJECT_DIR"

# Get tags to process
if [[ -n "$SINGLE_TAG" ]]; then
    TAGS=("$SINGLE_TAG")
else
    # Get all version tags sorted by version
    mapfile -t TAGS < <(git tag -l 'v*' --sort=v:refname)
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

    # Generate changelog content for this version
    NOTES_FILE="/tmp/release-notes-${tag}.md"
    if ! git-cliff --tag "$tag" --strip header > "$NOTES_FILE" 2>/dev/null; then
        echo "  ERROR: Failed to generate changelog for $tag"
        FAILED=$((FAILED + 1))
        continue
    fi

    # Check if content was generated
    if [[ ! -s "$NOTES_FILE" ]]; then
        echo "  SKIP: No changelog content for $tag"
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
