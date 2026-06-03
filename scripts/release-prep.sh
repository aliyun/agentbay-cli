#!/bin/bash
# Prepare a new release locally before tagging.
#
# This script is the local-side counterpart of the bilingual changelog
# pipeline (see docs/internal/bilingual-changelog-proposal.md, v3).
# It generates the bilingual section skeleton for VERSION and prepends
# (or rather, replaces the [Unreleased] section with) it in CHANGELOG.md.
# The Chinese sub-section is left as a TRANSLATE_ME placeholder so that
# the developer can translate it (typically via Claude Code) before
# committing and tagging.
#
# Usage:
#   bash scripts/release-prep.sh <VERSION>      # e.g. 0.4.0
#
# Dependencies:
#   - git-cliff (brew install git-cliff)
#   - awk, grep, head, tail, mktemp (POSIX)
#   - git
#
# Exit codes:
#   0 success, 1 validation failure, 2 internal error.

set -euo pipefail

VERSION="${1:-}"
TAG="v${VERSION}"

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
CHANGELOG="$PROJECT_DIR/CHANGELOG.md"

cd "$PROJECT_DIR"

# ---- 1. version format ----
if [[ -z "$VERSION" ]]; then
    echo "Usage: $(basename "$0") <VERSION>      # e.g. 0.4.0" >&2
    exit 1
fi
if [[ ! "$VERSION" =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    echo "❌ Invalid version: '$VERSION' (expected X.Y.Z, no leading 'v')" >&2
    exit 1
fi

# ---- 2. tooling ----
if ! command -v git-cliff >/dev/null 2>&1; then
    echo "❌ git-cliff is required. Install: brew install git-cliff" >&2
    exit 1
fi

# ---- 3. clean working tree ----
if [[ -n "$(git status --porcelain)" ]]; then
    echo "❌ Working tree is not clean. Commit or stash changes first." >&2
    git status --short >&2
    exit 1
fi

# ---- 4. branch check ----
CURRENT_BRANCH="$(git rev-parse --abbrev-ref HEAD)"
if [[ "$CURRENT_BRANCH" != "master" ]]; then
    echo "⚠️  Not on master branch (current: $CURRENT_BRANCH)" >&2
    read -r -p "Continue anyway? [y/N] " reply
    if [[ ! "$reply" =~ ^[Yy]$ ]]; then
        echo "Aborted." >&2
        exit 1
    fi
fi

# ---- 5. pull latest ----
echo "📥 Pulling latest from origin/$CURRENT_BRANCH..."
if ! git pull --ff-only origin "$CURRENT_BRANCH"; then
    echo "❌ git pull --ff-only failed. Resolve manually and retry." >&2
    exit 1
fi

# ---- 6. tag does not exist ----
if git rev-parse "$TAG" >/dev/null 2>&1; then
    echo "❌ Tag $TAG already exists. Choose a different version." >&2
    exit 1
fi

# ---- 7. generate bilingual skeleton via git-cliff ----
TMP_SECTION="$(mktemp -t release-prep-${VERSION}-XXXXXX.md)"
trap 'rm -f "$TMP_SECTION"' EXIT

echo "🪄 Generating bilingual changelog skeleton for $TAG..."
if ! git-cliff --tag "$TAG" --unreleased --strip header 2>/dev/null > "$TMP_SECTION"; then
    echo "❌ git-cliff failed. See its error output above." >&2
    exit 2
fi

if [[ ! -s "$TMP_SECTION" ]]; then
    echo "❌ git-cliff produced no output for $TAG (no commits since last tag?)." >&2
    exit 1
fi

# ---- 8. replace [Unreleased] section in CHANGELOG.md ----
if [[ ! -f "$CHANGELOG" ]]; then
    echo "❌ $CHANGELOG not found." >&2
    exit 2
fi

START="$(grep -n '^## \[Unreleased\]' "$CHANGELOG" | head -1 | cut -d: -f1 || true)"
if [[ -z "$START" ]]; then
    echo "❌ Could not find '## [Unreleased]' line in $CHANGELOG." >&2
    echo "   The CHANGELOG is expected to keep an [Unreleased] anchor at the top." >&2
    exit 2
fi

# End of [Unreleased] section: line of the next `## [` after START, or EOF.
END_REL="$(awk -v s="$START" 'NR > s && /^## \[/ { print NR; exit }' "$CHANGELOG")"
if [[ -z "$END_REL" ]]; then
    # No following section; tail off after EOF.
    END_REL=$(($(wc -l < "$CHANGELOG") + 1))
fi

# Splice: keep lines 1..(START-1), insert new section, then lines END_REL..EOF.
NEW_CHANGELOG="$(mktemp -t CHANGELOG-new-XXXXXX.md)"
trap 'rm -f "$TMP_SECTION" "$NEW_CHANGELOG"' EXIT

{
    if [[ "$START" -gt 1 ]]; then
        head -n "$((START - 1))" "$CHANGELOG"
    fi
    cat "$TMP_SECTION"
    echo ""
    # Re-insert an empty [Unreleased] anchor so the next release-prep run
    # has something to replace.
    echo "## [Unreleased]"
    echo ""
    if [[ "$END_REL" -le "$(wc -l < "$CHANGELOG")" ]]; then
        tail -n "+${END_REL}" "$CHANGELOG"
    fi
} > "$NEW_CHANGELOG"

mv "$NEW_CHANGELOG" "$CHANGELOG"
trap 'rm -f "$TMP_SECTION"' EXIT

# ---- 9. show diff ----
echo ""
echo "✅ CHANGELOG.md updated. Preview (first ~80 lines of diff):"
echo "================================================================"
git --no-pager diff CHANGELOG.md | head -80
echo "================================================================"

# ---- 10. next-step instructions ----
cat <<EOF

📋 Next steps:

  1. Translate the [$VERSION] section in CHANGELOG.md.
     In Claude Code, say:
       翻译 CHANGELOG.md 顶部 [$VERSION] 那段的 English 子段为中文，
       写到 ### 中文 段下，并删除 TRANSLATE_ME 注释行

  2. Review:
       vim CHANGELOG.md

  3. Commit + tag + push:
       git add CHANGELOG.md
       git commit -m "docs: changelog for $TAG"
       git tag $TAG
       git push origin $CURRENT_BRANCH $TAG

  4. The tag push triggers .github/workflows/homebrew.yml,
     which extracts the [$VERSION] section from CHANGELOG.md
     and creates the GitHub Release.

EOF
