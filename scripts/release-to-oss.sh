#!/bin/bash
# AgentBay CLI - Release to Production OSS
# Builds all platforms with a semantic version and uploads to public OSS bucket.
#
# Usage:
#   ./scripts/release-to-oss.sh --version 0.2.4 --description "新增镜像删除功能"
#   ./scripts/release-to-oss.sh --version 0.2.4 --description "desc" --dry-run
#   ./scripts/release-to-oss.sh --version 0.2.4 --description "desc" --force

set -euo pipefail

# ============================================================================
# Constants
# ============================================================================
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
BINARY_NAME="agentbay"
DEFAULT_CONFIG="$HOME/.agentbay-release.env"
PLATFORMS=("darwin-amd64" "darwin-arm64" "linux-amd64" "linux-arm64" "windows-amd64" "windows-arm64")

# ============================================================================
# Variables (set by argument parsing)
# ============================================================================
VERSION=""
DESCRIPTION=""
CONFIG_FILE="$DEFAULT_CONFIG"
DRY_RUN=false
FORCE=false
SKIP_BUILD=false
CHANNEL="stable"

# ============================================================================
# Helper Functions
# ============================================================================
log_info() { echo "  [INFO] $*"; }
log_ok() { echo "  [OK] $*"; }
log_warn() { echo "  [WARN] $*"; }
log_error() { echo "  [ERROR] $*" >&2; }
log_step() { echo ""; echo "== $* =="; }

usage() {
    cat <<EOF
Usage: $(basename "$0") --version <VERSION> [--description <DESC>] [OPTIONS]

Builds all platforms and uploads production artifacts to public OSS.

Required:
  --version VERSION       Semantic version (e.g., 0.2.4, no 'v' prefix)

Optional:
  --description DESC      Release description (for version_manifest.json)
                          If omitted, auto-generates from git commits diff
                          between aliyun/master and current branch
  --config PATH           Config file path (default: ~/.agentbay-release.env)
  --dry-run               Show what would happen without uploading
  --force                 Overwrite if version already exists on OSS
  --skip-build            Skip rebuild (use existing packages/ contents)
  --channel CHANNEL       Release channel (default: stable)
  --help                  Show this help

Examples:
  $(basename "$0") --version 0.2.4
  $(basename "$0") --version 0.2.4 --description "新增镜像删除功能"
  $(basename "$0") --version 0.2.4 --dry-run
  $(basename "$0") --version 0.2.4 --force
EOF
    exit 0
}

# ============================================================================
# OSS Tool Abstraction (aligned with .aoneci/cicd.yml approach)
# Uses standalone ossutil binary with global config file
# ============================================================================
OSS_TOOL=""
OSS_CONFIG_FILE="$PROJECT_DIR/.ossutil_config_tmp"

oss_detect_tool() {
    if command -v ossutil >/dev/null 2>&1; then
        OSS_TOOL="ossutil"
    elif command -v ossutil64 >/dev/null 2>&1; then
        OSS_TOOL="ossutil64"
    elif [[ -x "$PROJECT_DIR/.ossutil_bin/ossutil" ]]; then
        OSS_TOOL="$PROJECT_DIR/.ossutil_bin/ossutil"
    else
        log_warn "ossutil not found, attempting to install..."
        oss_install_tool
    fi
    log_info "Using OSS tool: $OSS_TOOL"
}

oss_install_tool() {
    local os_type arch_type download_url
    os_type=$(uname -s | tr '[:upper:]' '[:lower:]')
    arch_type=$(uname -m)

    # Map architecture
    case "$arch_type" in
        x86_64|amd64) arch_type="amd64" ;;
        arm64|aarch64) arch_type="arm64" ;;
        *) log_error "Unsupported architecture: $arch_type"; exit 1 ;;
    esac

    # Map OS (ossutil uses "mac" for macOS)
    case "$os_type" in
        darwin) os_type="mac" ;;
        linux) os_type="linux" ;;
        *) log_error "Unsupported OS: $os_type"; exit 1 ;;
    esac

    download_url="https://gosspublic.alicdn.com/ossutil/1.7.19/ossutil-v1.7.19-${os_type}-${arch_type}.zip"
    log_info "Downloading ossutil from: $download_url"

    local temp_dir
    temp_dir=$(mktemp -d)
    if ! curl -sL "$download_url" -o "$temp_dir/ossutil.zip"; then
        log_error "Failed to download ossutil"
        rm -rf "$temp_dir"
        exit 1
    fi

    if ! unzip -q "$temp_dir/ossutil.zip" -d "$temp_dir/extracted"; then
        log_error "Failed to extract ossutil"
        rm -rf "$temp_dir"
        exit 1
    fi

    # Find the ossutil binary in extracted files
    local ossutil_bin
    ossutil_bin=$(find "$temp_dir/extracted" -name "ossutil*" -type f | head -1)
    if [[ -z "$ossutil_bin" ]]; then
        log_error "Failed to find ossutil binary after extraction"
        rm -rf "$temp_dir"
        exit 1
    fi

    mkdir -p "$PROJECT_DIR/.ossutil_bin"
    chmod +x "$ossutil_bin"
    mv "$ossutil_bin" "$PROJECT_DIR/.ossutil_bin/ossutil"
    rm -rf "$temp_dir"

    OSS_TOOL="$PROJECT_DIR/.ossutil_bin/ossutil"
    log_ok "ossutil installed to $PROJECT_DIR/.ossutil_bin/ossutil"
}

oss_configure() {
    # Same approach as .aoneci/cicd.yml: use ossutil config for global auth
    log_info "Configuring ossutil..."
    if "$OSS_TOOL" config \
        --endpoint="$OSS_ENDPOINT" \
        --access-key-id="$OSS_ACCESS_KEY_ID" \
        --access-key-secret="$OSS_ACCESS_KEY_SECRET" \
        -c "$OSS_CONFIG_FILE" >/dev/null 2>&1; then
        log_ok "ossutil configured"
    else
        log_error "ossutil configuration failed"
        exit 1
    fi
}

oss_cleanup() {
    rm -f "$OSS_CONFIG_FILE"
}

oss_cp() {
    local src="$1"
    local dst="$2"
    "$OSS_TOOL" cp "$src" "$dst" --force -c "$OSS_CONFIG_FILE"
}

oss_ls() {
    local path="$1"
    "$OSS_TOOL" ls "$path" -c "$OSS_CONFIG_FILE" 2>/dev/null
}

oss_download() {
    local src="$1"
    local dst="$2"
    "$OSS_TOOL" cp "$src" "$dst" --force -c "$OSS_CONFIG_FILE"
}

oss_set_acl() {
    local path="$1"
    local acl="$2"
    "$OSS_TOOL" set-acl "$path" "$acl" -c "$OSS_CONFIG_FILE" 2>/dev/null || true
}

# ============================================================================
# SHA256 Helper (macOS/Linux compatible)
# ============================================================================
compute_sha256() {
    local file="$1"
    if command -v sha256sum >/dev/null 2>&1; then
        sha256sum "$file" | awk '{print $1}'
    elif command -v shasum >/dev/null 2>&1; then
        shasum -a 256 "$file" | awk '{print $1}'
    else
        log_error "No SHA256 tool available"
        exit 1
    fi
}

generate_sha256_file() {
    local file="$1"
    local hash
    hash=$(compute_sha256 "$file")
    echo "$hash  $(basename "$file")" > "${file}.sha256"
}

# ============================================================================
# Auto-generate Description from Git Diff
# ============================================================================
auto_generate_description() {
    log_info "No --description provided, auto-generating from git diff..."

    # Fetch aliyun remote
    if ! git remote get-url aliyun >/dev/null 2>&1; then
        log_error "Remote 'aliyun' not found. Cannot auto-generate description."
        log_error "Please provide --description manually, or add the aliyun remote:"
        echo "  git remote add aliyun git@github.com:aliyun/agentbay-cli.git"
        exit 1
    fi

    log_info "Fetching aliyun remote..."
    if ! git fetch aliyun >/dev/null 2>&1; then
        log_error "Failed to fetch aliyun remote. Check network and SSH key."
        log_error "Please provide --description manually."
        exit 1
    fi

    # Get commit log between aliyun/master and current HEAD
    local commits
    commits=$(git log aliyun/master..HEAD --oneline --no-merges 2>/dev/null || true)

    if [[ -z "$commits" ]]; then
        log_error "No commits found between aliyun/master and HEAD."
        log_error "Current branch might already be up to date with aliyun/master."
        log_error "Please provide --description manually."
        exit 1
    fi

    # Count commits
    local commit_count
    commit_count=$(echo "$commits" | wc -l | tr -d ' ')

    # Generate description from commit subjects
    if [[ "$commit_count" -eq 1 ]]; then
        # Single commit: use its subject directly (strip hash prefix)
        DESCRIPTION=$(echo "$commits" | sed 's/^[a-f0-9]* //')
    else
        # Multiple commits: extract feat/fix subjects, combine them
        local subjects
        subjects=$(echo "$commits" | sed 's/^[a-f0-9]* //' | \
            grep -v "^Merge " | \
            grep -v "^chore" | \
            grep -v "^style" | \
            grep -v "^test" | \
            head -5)

        if [[ -z "$subjects" ]]; then
            # Fallback: just use all subjects
            subjects=$(echo "$commits" | sed 's/^[a-f0-9]* //' | head -5)
        fi

        # Join with "; " and truncate if too long
        DESCRIPTION=$(echo "$subjects" | tr '\n' ';' | sed 's/;$//' | sed 's/;/; /g')

        # Truncate to 100 chars if too long
        if [[ ${#DESCRIPTION} -gt 100 ]]; then
            DESCRIPTION="${DESCRIPTION:0:97}..."
        fi
    fi

    log_info "Auto-generated description ($commit_count commits):"
    log_info "  \"$DESCRIPTION\""
    echo ""

    # Ask for confirmation
    read -r -p "  Use this description? [Y/n] " confirm
    if [[ "$confirm" =~ ^[Nn] ]]; then
        read -r -p "  Enter custom description: " DESCRIPTION
        if [[ -z "$DESCRIPTION" ]]; then
            log_error "Description cannot be empty."
            exit 1
        fi
    fi
}

# ============================================================================
# Argument Parsing
# ============================================================================
parse_args() {
    while [[ $# -gt 0 ]]; do
        case "$1" in
            --version)
                VERSION="$2"; shift 2 ;;
            --description)
                DESCRIPTION="$2"; shift 2 ;;
            --config)
                CONFIG_FILE="$2"; shift 2 ;;
            --dry-run)
                DRY_RUN=true; shift ;;
            --force)
                FORCE=true; shift ;;
            --skip-build)
                SKIP_BUILD=true; shift ;;
            --channel)
                CHANNEL="$2"; shift 2 ;;
            --help|-h)
                usage ;;
            *)
                log_error "Unknown option: $1"
                usage ;;
        esac
    done

    if [[ -z "$VERSION" ]]; then
        log_error "--version is required"
        exit 1
    fi
    # Validate version format
    if ! echo "$VERSION" | grep -qE '^[0-9]+\.[0-9]+\.[0-9]+(-[a-zA-Z0-9.]+)?$'; then
        log_error "Invalid version format: $VERSION (expected: X.Y.Z)"
        exit 1
    fi
    # If description not provided, auto-generate from git diff
    if [[ -z "$DESCRIPTION" ]]; then
        auto_generate_description
    fi
}

# ============================================================================
# Phase 1: Initialization & Validation
# ============================================================================
phase_init() {
    log_step "Phase 1: Initialization"

    # Load config
    if [[ ! -f "$CONFIG_FILE" ]]; then
        log_error "Config file not found: $CONFIG_FILE"
        echo ""
        echo "  Setup instructions:"
        echo "    cp scripts/release-to-oss.sample.env ~/.agentbay-release.env"
        echo "    vim ~/.agentbay-release.env  # Fill in AK/SK"
        exit 1
    fi

    # shellcheck disable=SC1090
    source "$CONFIG_FILE"

    # Validate credentials
    if [[ -z "${OSS_ACCESS_KEY_ID:-}" ]]; then
        log_error "OSS_ACCESS_KEY_ID not set in $CONFIG_FILE"
        exit 1
    fi
    if [[ -z "${OSS_ACCESS_KEY_SECRET:-}" ]]; then
        log_error "OSS_ACCESS_KEY_SECRET not set in $CONFIG_FILE"
        exit 1
    fi
    if [[ -z "${OSS_BUCKET:-}" ]]; then
        log_error "OSS_BUCKET not set in $CONFIG_FILE"
        exit 1
    fi
    if [[ -z "${OSS_ENDPOINT:-}" ]]; then
        log_error "OSS_ENDPOINT not set in $CONFIG_FILE"
        exit 1
    fi

    log_info "Config loaded from: $CONFIG_FILE"
    log_info "Bucket: $OSS_BUCKET"
    log_info "Endpoint: $OSS_ENDPOINT"
    log_info "Version: $VERSION"
    log_info "Channel: $CHANNEL"
    if $DRY_RUN; then
        log_warn "DRY RUN mode - no uploads will be performed"
    fi

    # Check prerequisites
    local missing=()
    command -v go >/dev/null 2>&1 || missing+=("go")
    command -v make >/dev/null 2>&1 || missing+=("make")
    command -v jq >/dev/null 2>&1 || missing+=("jq (brew install jq)")
    command -v zip >/dev/null 2>&1 || missing+=("zip")
    if ! command -v sha256sum >/dev/null 2>&1 && ! command -v shasum >/dev/null 2>&1; then
        missing+=("sha256sum or shasum")
    fi

    if [[ ${#missing[@]} -gt 0 ]]; then
        log_error "Missing prerequisites: ${missing[*]}"
        exit 1
    fi

    # Detect OSS tool
    oss_detect_tool
    oss_configure

    # Verify project directory
    if [[ ! -f "$PROJECT_DIR/go.mod" ]]; then
        log_error "Not in project root. Run from agentbay-cli directory."
        exit 1
    fi

    # Check if version exists on OSS
    if ! $FORCE; then
        local ls_output
        ls_output=$(oss_ls "oss://$OSS_BUCKET/$VERSION/" 2>&1 || true)
        if echo "$ls_output" | grep -q "$BINARY_NAME-$VERSION"; then
            log_error "Version $VERSION already exists on OSS. Use --force to overwrite."
            oss_cleanup
            exit 1
        fi
    fi

    log_ok "Initialization complete"
}

# ============================================================================
# Phase 2: Build
# ============================================================================
phase_build() {
    log_step "Phase 2: Build"

    if $SKIP_BUILD; then
        log_warn "Skipping build (--skip-build)"
        if [[ ! -d "$PROJECT_DIR/packages" ]] || [[ -z "$(ls "$PROJECT_DIR/packages/"*.tar.gz 2>/dev/null)" ]]; then
            log_error "No packages found. Remove --skip-build to rebuild."
            oss_cleanup
            exit 1
        fi
        return
    fi

    cd "$PROJECT_DIR"

    # Clean
    log_info "Cleaning previous artifacts..."
    rm -rf bin/ packages/

    # Build
    log_info "Building all platforms with VERSION=$VERSION..."
    VERSION="$VERSION" make build-all-optimized

    # Verify binaries
    local expected_bins=("$BINARY_NAME-darwin-amd64" "$BINARY_NAME-darwin-arm64" \
                         "$BINARY_NAME-linux-amd64" "$BINARY_NAME-linux-arm64" \
                         "$BINARY_NAME-windows-amd64.exe" "$BINARY_NAME-windows-arm64.exe")
    for bin in "${expected_bins[@]}"; do
        if [[ ! -f "bin/$bin" ]]; then
            log_error "Missing binary: bin/$bin"
            oss_cleanup
            exit 1
        fi
    done

    # Verify version on native binary
    local native_arch
    if [[ "$(uname -m)" == "arm64" ]]; then
        native_arch="arm64"
    else
        native_arch="amd64"
    fi
    local native_bin="bin/$BINARY_NAME-darwin-$native_arch"
    if [[ -f "$native_bin" ]]; then
        chmod +x "$native_bin"
        local version_output embedded_version
        version_output=$("$native_bin" version 2>/dev/null || true)
        embedded_version=$(echo "$version_output" | grep -oE '[0-9]+\.[0-9]+\.[0-9]+' | head -1 || true)
        if [[ "$embedded_version" != "$VERSION" ]]; then
            log_warn "Version mismatch: binary reports '$embedded_version', expected '$VERSION'"
        else
            log_ok "Version verified: $embedded_version"
        fi
    fi

    log_ok "Build complete ($(ls bin/ | wc -l | tr -d ' ') binaries)"
}

# ============================================================================
# Phase 3: Package
# ============================================================================
phase_package() {
    log_step "Phase 3: Package"

    if $SKIP_BUILD && [[ -d "$PROJECT_DIR/packages" ]]; then
        log_warn "Using existing packages/"
        return
    fi

    cd "$PROJECT_DIR"
    mkdir -p packages

    for platform in "${PLATFORMS[@]}"; do
        local os="${platform%-*}"
        local arch="${platform#*-}"

        if [[ "$os" == "windows" ]]; then
            local src_bin="bin/$BINARY_NAME-$platform.exe"
            local pkg_zip="packages/$BINARY_NAME-$VERSION-$platform.zip"
            local pkg_exe="packages/$BINARY_NAME-$VERSION-$platform.exe"

            # Create zip
            local temp_dir
            temp_dir=$(mktemp -d)
            cp "$src_bin" "$temp_dir/$BINARY_NAME.exe"
            (cd "$temp_dir" && zip -q "$PROJECT_DIR/$pkg_zip" "$BINARY_NAME.exe")
            rm -rf "$temp_dir"
            generate_sha256_file "$pkg_zip"
            log_info "Created: $(basename "$pkg_zip")"

            # Copy standalone exe
            cp "$src_bin" "$pkg_exe"
            generate_sha256_file "$pkg_exe"
            log_info "Created: $(basename "$pkg_exe")"
        else
            local src_bin="bin/$BINARY_NAME-$platform"
            local pkg_tar="packages/$BINARY_NAME-$VERSION-$platform.tar.gz"

            # Create tar.gz
            local temp_dir
            temp_dir=$(mktemp -d)
            cp "$src_bin" "$temp_dir/$BINARY_NAME"
            chmod +x "$temp_dir/$BINARY_NAME"
            tar -czf "$pkg_tar" -C "$temp_dir" "$BINARY_NAME"
            rm -rf "$temp_dir"
            generate_sha256_file "$pkg_tar"
            log_info "Created: $(basename "$pkg_tar")"
        fi
    done

    local pkg_count
    pkg_count=$(ls packages/ | wc -l | tr -d ' ')
    log_ok "Packaging complete ($pkg_count files)"
}

# ============================================================================
# Phase 4: Upload
# ============================================================================
phase_upload() {
    log_step "Phase 4: Upload to OSS"

    cd "$PROJECT_DIR"

    local oss_prefix="oss://$OSS_BUCKET/$VERSION"
    local file_count=0

    for file in packages/*; do
        if [[ -f "$file" ]]; then
            local filename
            filename=$(basename "$file")
            if $DRY_RUN; then
                log_info "[DRY RUN] Would upload: $filename -> $oss_prefix/$filename"
            else
                log_info "Uploading: $filename"
                local retry=0
                while [[ $retry -lt 3 ]]; do
                    if oss_cp "$file" "$oss_prefix/$filename" >/dev/null 2>&1; then
                        oss_set_acl "$oss_prefix/$filename" "public-read"
                        break
                    fi
                    retry=$((retry + 1))
                    if [[ $retry -lt 3 ]]; then
                        log_warn "Upload failed, retrying ($retry/3)..."
                        sleep 5
                    else
                        log_error "Failed to upload $filename after 3 attempts"
                        oss_cleanup
                        exit 1
                    fi
                done
            fi
            file_count=$((file_count + 1))
        fi
    done

    if $DRY_RUN; then
        log_ok "[DRY RUN] Would upload $file_count files to $oss_prefix/"
    else
        log_ok "Uploaded $file_count files to $oss_prefix/"
    fi
}

# ============================================================================
# Phase 5: Update version_manifest.json
# ============================================================================
phase_update_manifest() {
    log_step "Phase 5: Update version_manifest.json"

    cd "$PROJECT_DIR"

    local manifest_oss_path="oss://$OSS_BUCKET/version_manifest.json"
    local manifest_local="/tmp/version_manifest_$$.json"
    local manifest_new="/tmp/version_manifest_new_$$.json"
    local released_at
    released_at=$(date -u +"%Y-%m-%dT%H:%M:%SZ")

    # Download current manifest
    if oss_download "$manifest_oss_path" "$manifest_local" >/dev/null 2>&1; then
        log_info "Downloaded existing version_manifest.json from OSS"
    else
        if $DRY_RUN; then
            log_warn "Cannot download existing version_manifest.json from OSS"
            log_warn "Preview will only show the new version entry (existing versions missing)"
            log_warn "This is normal if credentials are not configured or network is unavailable"
            log_warn "The real 'make release' will download and merge correctly"
        else
            log_info "No existing manifest on OSS, creating new one"
        fi
        echo '{"channels":{"stable":{"latest_version":"","versions":[]}}}' > "$manifest_local"
    fi

    # Update manifest with jq
    local new_entry
    new_entry=$(jq -n \
        --arg ver "$VERSION" \
        --arg desc "$DESCRIPTION" \
        --arg released "$released_at" \
        '{version: $ver, description: $desc, is_latest: true}')

    jq --arg channel "$CHANNEL" \
       --arg version "$VERSION" \
       --argjson entry "$new_entry" \
       '
       .channels[$channel].latest_version = $version |
       .channels[$channel].versions = (
         [.channels[$channel].versions[] | .is_latest = false] |
         [$entry] + .
       )
       ' "$manifest_local" > "$manifest_new"

    if $DRY_RUN; then
        log_info "[DRY RUN] Updated manifest preview:"
        jq '.' "$manifest_new"
        # Save to packages/ for local review
        cp "$manifest_new" "$PROJECT_DIR/packages/version_manifest.json"
        log_ok "[DRY RUN] Saved to packages/version_manifest.json for review"
    else
        log_info "Uploading updated version_manifest.json"
        oss_cp "$manifest_new" "$manifest_oss_path" >/dev/null 2>&1
        oss_set_acl "$manifest_oss_path" "public-read"
        log_ok "version_manifest.json updated"
    fi

    # Cleanup temp files
    rm -f "$manifest_local" "$manifest_new"
}

# ============================================================================
# Phase 6: Summary
# ============================================================================
phase_summary() {
    log_step "Release Summary"

    local base_url="https://$OSS_BUCKET.$OSS_ENDPOINT/$VERSION"

    if $DRY_RUN; then
        echo ""
        echo "  *** DRY RUN - No files were uploaded ***"
        echo ""
    fi

    echo "  Version: $VERSION"
    echo "  Description: $DESCRIPTION"
    echo "  Channel: $CHANNEL"
    echo ""
    echo "  Download URLs:"
    echo ""
    echo "  macOS:"
    echo "    Intel:         $base_url/$BINARY_NAME-$VERSION-darwin-amd64.tar.gz"
    echo "    Apple Silicon: $base_url/$BINARY_NAME-$VERSION-darwin-arm64.tar.gz"
    echo ""
    echo "  Linux:"
    echo "    x64:           $base_url/$BINARY_NAME-$VERSION-linux-amd64.tar.gz"
    echo "    ARM64:         $base_url/$BINARY_NAME-$VERSION-linux-arm64.tar.gz"
    echo ""
    echo "  Windows:"
    echo "    x64 (exe):     $base_url/$BINARY_NAME-$VERSION-windows-amd64.exe"
    echo "    ARM64 (exe):   $base_url/$BINARY_NAME-$VERSION-windows-arm64.exe"
    echo "    x64 (zip):     $base_url/$BINARY_NAME-$VERSION-windows-amd64.zip"
    echo "    ARM64 (zip):   $base_url/$BINARY_NAME-$VERSION-windows-arm64.zip"
    echo ""
    echo "  Manifest: https://$OSS_BUCKET.$OSS_ENDPOINT/version_manifest.json"
    echo ""
}

# ============================================================================
# Main
# ============================================================================
main() {
    echo "============================================"
    echo " AgentBay CLI - Release to Production OSS"
    echo "============================================"

    parse_args "$@"
    phase_init
    phase_build
    phase_package
    phase_upload
    phase_update_manifest
    phase_summary

    # Cleanup OSS config
    oss_cleanup

    if $DRY_RUN; then
        echo "  [DRY RUN complete - no changes were made]"
    else
        echo "  [Release $VERSION published successfully]"
    fi
}

main "$@"
