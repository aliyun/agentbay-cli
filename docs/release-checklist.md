# Release Pipeline Checklist

End-to-end checklist for cutting a new release of `agentbay-cli` and publishing
it through the Homebrew tap so that users get a pre-built bottle (秒装).

The pipeline has two GitHub Actions workflows:

| Workflow file | Trigger | What it does |
|---|---|---|
| `.github/workflows/homebrew.yml` | `git push` of a `v*` tag, or manual `workflow_dispatch` with a version input | Builds cross-platform binaries, packages bottles, creates the GitHub Release, regenerates `homebrew/agentbay.rb` with real SHA256s, commits it back to `master`, and deploys the Windows install script to GitHub Pages. |
| `.github/workflows/push-to-homebrew-tap.yml` | Manual `workflow_dispatch` with `push_to_tap=true` | Copies `homebrew/agentbay.rb` into the `aliyun/homebrew-agentbay` tap repo so `brew install agentbay` picks up the new version. |

The two workflows are intentionally separated so you can verify the formula on
`master` before pushing it to the public tap.

---

## Pre-flight (before tagging)

- [ ] `CHANGELOG.md` is current, or you are happy for `git-cliff` to regenerate it automatically as part of the release.
- [ ] `homebrew/agentbay.rb` is committed (its contents will be overwritten by the workflow — the file just needs to exist for validation).
- [ ] All required secrets exist on the repo:
  - `GITHUB_TOKEN` — provided automatically; needs `contents: write` (already declared in `homebrew.yml`).
  - `HOMEBREW_TAP_TOKEN` — PAT with `repo` scope on `aliyun/homebrew-agentbay`. Required by `push-to-homebrew-tap.yml`.
- [ ] Tag does not already exist locally or on the remote (`git tag -l vX.Y.Z`, `gh release view vX.Y.Z` should both be empty).
- [ ] You have write access to `aliyun/agentbay-cli` (for the commit-back step) and `aliyun/homebrew-agentbay` (for the tap push).

---

## Step 1 — Cut the release

Pick one of:

**Option A: tag-driven (recommended for production releases)**

```bash
git checkout master
git pull
git tag v0.4.0
git push origin v0.4.0
```

**Option B: manual dispatch (for re-runs or test releases)**

GitHub → Actions → **Agentbay CLI Official Homebrew Release** → **Run workflow** → enter version (no `v` prefix, e.g. `0.4.0`).

Either path triggers `homebrew.yml`.

---

## Step 2 — Watch the workflow run

In the Actions tab, open the run and verify each step:

- [ ] **Build Multi-platform Binaries** — `bin/` contains all 4 expected files:
  - `agentbay-darwin-arm64`
  - `agentbay-darwin-amd64`
  - `agentbay-linux-amd64`
  - `agentbay-linux-arm64`
  Plus any Windows variants (`-windows-amd64.exe`, `-windows-arm64.exe`). If any darwin/linux binary is missing, the next step's fail-fast check will abort the workflow.
- [ ] **Create Homebrew Bottles** — log shows one `✅ Created:` line per bottle tag. With the hardened pipeline you should see seven tagged tarballs total:
  - `arm64_sonoma`, `arm64_ventura`, `arm64_sequoia` (darwin-arm64)
  - `sonoma`, `ventura` (darwin-amd64)
  - `x86_64_linux`, `aarch64_linux`
  Ends with `✅ All primary bottles present`. If you see `❌ Missing primary bottle for ...`, see Troubleshooting → "Missing primary bottle".
- [ ] **Create GitHub Release** — release `vX.Y.Z` exists with all bottle `.bottle.tar.gz` files, source archive assets, and Windows installers attached.
- [ ] **Prepare Homebrew Core Submission** — commits `feat: add official Homebrew formula for vX.Y.Z` to `master`. Open the commit and confirm the new `homebrew/agentbay.rb` contains real SHA256s (no `PLACEHOLDER_*`).
- [ ] **Deploy to GitHub Pages** — `https://aliyun.github.io/agentbay-cli/windows` returns the latest PowerShell installer (curl/visit to confirm).

---

## Step 3 — Smoke-test the bottle locally (optional but recommended)

Before pushing to the tap, you can manually verify the new formula against the
GitHub Release using a local copy:

```bash
# Re-download the just-committed formula
curl -fsSL https://raw.githubusercontent.com/aliyun/agentbay-cli/master/homebrew/agentbay.rb -o /tmp/agentbay.rb

# On macOS or Linux, install directly from the file (bypasses the tap)
brew uninstall agentbay 2>/dev/null
brew install --formula /tmp/agentbay.rb -v 2>&1 | tee /tmp/install.log

# Expected: log contains "==> Pouring agentbay--X.Y.Z.<tag>.bottle.tar.gz"
grep -E "Pouring|Cloning|go build" /tmp/install.log
agentbay version
```

If you see `==> Pouring` → 🎉 the bottle works.
If you see `==> Cloning` or `go build` → the bottle for your OS tag wasn't matched; see Troubleshooting → "Falling back to source build".

---

## Step 4 — Push the formula to the tap

GitHub → Actions → **Push to Homebrew Tap** → **Run workflow** → set
`push_to_tap = true`, leave `version` empty (it will read the version straight
from the freshly committed `homebrew/agentbay.rb`).

Verify:

- [ ] Workflow finishes green.
- [ ] `aliyun/homebrew-agentbay` has a new commit `Update agentbay formula and README to version X.Y.Z`.
- [ ] `Formula/agentbay.rb` in the tap repo matches `homebrew/agentbay.rb` from this repo.

---

## Step 5 — Final user-facing verification

On a clean machine (or after fully removing your local install):

```bash
brew untap aliyun/agentbay 2>/dev/null
brew uninstall agentbay 2>/dev/null
brew cleanup

brew tap aliyun/agentbay
brew install agentbay -v 2>&1 | tee /tmp/install.log
agentbay version
```

- [ ] Install completes in seconds (no Go toolchain download, no `go build`).
- [ ] `agentbay version` prints the version you just released.
- [ ] `/tmp/install.log` contains `==> Pouring` for an `agentbay--X.Y.Z.<your-os-tag>.bottle.tar.gz` file.

Windows (PowerShell):

```powershell
powershell -Command "irm https://aliyun.github.io/agentbay-cli/windows | iex"
agentbay version
```

- [ ] PS1 installer reports the new version, places `agentbay.exe` in `%LOCALAPPDATA%\agentbay`, and `agentbay version` works after restarting the terminal.

---

## Troubleshooting

### "Missing primary bottle" abort

`Create Homebrew Bottles` exits with `❌ N primary bottle(s) missing`. The
fail-fast was added precisely so this never silently ships a half-broken
formula.

Likely causes:
1. **`make dist` didn't build a target** — check the `Makefile`'s `dist` target. It must produce `bin/agentbay-darwin-arm64`, `bin/agentbay-darwin-amd64`, `bin/agentbay-linux-amd64`, `bin/agentbay-linux-arm64`. Add the missing target and re-run.
2. **A cross-compile step failed silently** — search the `Build Multi-platform Binaries` log for the affected GOOS/GOARCH, fix the error, push, re-trigger.

### Falling back to source build instead of pouring a bottle

Local `brew install` prints `==> Cloning` or runs `go build`. Reasons:

- **Your macOS version isn't in the bottle tag list.** Current tags cover Sonoma, Ventura and Sequoia. If you're on Big Sur/Monterey you'll source-build by design. Either upgrade macOS or add the tag to `BOTTLE_MAP` in `homebrew.yml` (and rerun the release pipeline).
- **Bottle URL 404.** Open the URL in the formula's `root_url` + filename in a browser. If 404, the Release Assets upload step failed — re-run the release.
- **Formula and bottle filename mismatch.** Compare the `sha256 cellar: :any_skip_relocation, <tag>: "..."` lines in `homebrew/agentbay.rb` against the Release asset filenames. Tags must match `<tag>` in `agentbay-<ver>.<tag>.bottle.tar.gz`.

### `push-to-homebrew-tap.yml` fails to push

- Check the `HOMEBREW_TAP_TOKEN` secret hasn't expired (PATs without an expiry are recommended for this token, or use a fine-grained token with explicit `Contents: Read & Write` on `aliyun/homebrew-agentbay`).
- The default branch on the tap is `main`; if you renamed it, update the `git push origin main` fallback in the workflow.

### Formula committed back with PLACEHOLDER SHA256s

This means the bottle creation step ran but every `BOTTLE_SHA[$tag]` was empty.
Cross-check `bottles/` listing in the Actions log against the workflow's
`BOTTLE_TAGS` array — if they don't match (e.g., a typo in a tag name), the
formula will be empty. The fail-fast check should catch this; if it didn't,
something upstream of the validation skipped silently — open an issue.

### CGO becomes a dependency

If at any point the CLI starts linking against C libraries (cgo, sqlite, etc.),
the "one binary per OS family" assumption breaks. You **must** then switch
`Create Homebrew Bottles` from "duplicate the same tarball under multiple tags"
to "build a separate bottle on a real runner for each OS tag" (e.g. matrix
build across `macos-14`, `macos-13`, `ubuntu-latest`, ...). At that point the
`cellar: :any_skip_relocation` qualifier is no longer safe either; review the
[Homebrew bottle docs](https://docs.brew.sh/Bottles) before changing.

---

## Rollback

If a release goes out broken:

1. Delete the bad GitHub Release **and** the underlying tag:
   ```bash
   gh release delete vX.Y.Z --yes --cleanup-tag
   ```
2. In the tap repo, revert the offending commit on `Formula/agentbay.rb` so users start getting the previous version on `brew upgrade`:
   ```bash
   git -C path/to/homebrew-agentbay revert <commit-sha>
   git -C path/to/homebrew-agentbay push
   ```
3. Cut a new patch release (`vX.Y.(Z+1)`) following this checklist.

Do **not** force-push a fixed bottle under the same tag — Homebrew users will see SHA256 mismatches when their cached metadata is out of date.
