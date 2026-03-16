# Skills Documentation Update Plan

> **For Claude:** Use verification-before-completion before claiming docs update complete.

**Goal:** Ensure all skills-related markdown (README, USER_GUIDE) reflect the current CLI behavior and are consistent.

**Scope:** README.md, docs/USER_GUIDE.md. Plans under docs/plans/ are implementation history; no edits unless explicitly requested.

---

## Task 1: README.md

**File:** `README.md`

**Requirements:**
- Features section MUST mention skills (push, list/show, group create/list, add/remove skills in groups).
- Quick Start MUST include at least one skills block: e.g. `agentbay skills push`, `agentbay skills group create`, `agentbay skills group list`.
- Link to User Guide for detailed usage unchanged.

**Verification:** Read README.md; confirm "Skills" in Features and "Skills (optional)" block in Quick Start exist.

---

## Task 2: docs/USER_GUIDE.md — Section 3 Skills

**File:** `docs/USER_GUIDE.md`

**Requirements:**
- Section "3. Skills (requirement a)" MUST list: push, list, show, group create, group list, group show, group add-skill, group remove-skill.
- `group create` MUST note: success prints group-id; use `-v` to see raw API response for debugging.
- Placeholder notes for list / group show kept as-is.

**Verification:** Read Section 3; confirm commands and -v note for group create.

---

## Task 3: docs/USER_GUIDE.md — FAQ

**File:** `docs/USER_GUIDE.md`

**Requirements:**
- FAQ "Enable detailed logs?" MUST show skills example alongside image example, e.g. `agentbay -v skills group create my-group` or `agentbay -v skills group list`.

**Verification:** Grep for "Enable detailed logs" and "-v"; confirm skills example present.

---

## Task 4: Verification before completion

**Commands:**
1. `go build ./...` — must exit 0 (docs don’t affect build; sanity check).
2. Read README.md and USER_GUIDE.md sections above; confirm content matches plan.

**Claim completion only after:** Both checks done and output confirmed.
