# sacr E2E matrix — PR mode

This repo audits `sacr` (the code reviewer at `shubam-disseqt/shubam-ai-code-reviewer`) by running its real PR-review flow against a set of persistent case PRs on THIS repo.

Runs nightly at 03:00 UTC + on demand via `workflow_dispatch`.

## Why the audit lives here (not on the product repo)

The product repo is treated as a shipped product other repos consume via the sacr GitHub Action. It has no E2E scaffolding. This repo — the matrix — owns:

- Seeded-bug branches (`pr-N-branch`)
- Ground-truth expected findings (`.matrix/cases/<branch>.json` on `main`, never on the case branch — anything in the PR diff gets indexed and quoted back to the model, which contaminated earlier recall numbers)
- The runner workflow (`.github/workflows/e2e-matrix.yml`)
- Assertion + aggregator tools (`.e2e/`)

## What "PR mode" means

Every case is a **persistent open PR** on THIS repo (e.g. `pr-3-branch → main`). Each nightly run:

1. For each case: check out the branch, build `sacr` from product-repo source
2. Run `sacr review --pr N --format github` → posts inline comments to the persistent PR
3. Run `sacr review --commit HEAD --format json` `RUNS` times (default 3) → `actual-1..N.json`
4. `.e2e/assert` compares every run against `.matrix/cases/<branch>.json` from `main`. A finding counts only when **every** run caught it (intersection); union recall is reported next to it
5. `.e2e/aggregate` produces a matrix-wide report and gates on `intersection recall >= 80%` AND no missed CRITICAL

`sacr`'s fingerprint markers (`<!-- sacr:fp:HEX -->`) auto-clean stale comments across runs.

## Case catalog

| Case branch | PR # | Seeded bugs |
|---|---|---|
| `pr-11-branch` | 13 | 1 (SQL injection) |
| `pr-19-branch` | 19 | 12 (index blast radius, slice race, goroutine leak, SQLi, AWS creds, md5, InsecureSkipVerify, unclosed body, path traversal, ignored Atoi, off-by-one, vulnerable dep) — full-feature audit case |
| `pr-1-branch` | 1 | 0 (docs-only overview, no ground truth yet) |
| `pr-3-branch` | 6 | 5 (no ground truth yet) |
| `pr-5-branch` | 3 | 2 (no ground truth yet) |
| `pr-10-branch` | 8 | 0 (deps-only bump, no ground truth yet) |

PR numbers are discovered at runtime via `gh pr list --head <branch>` — branches can be renamed or PRs recreated without workflow edits.

## expected.json schema

```json
{
  "case": "pr-3-branch",
  "description": "human-readable archetype summary",
  "findings": [
    {
      "path": "pricing.go",
      "start_line": 9,
      "end_line": 9,
      "category": "security",
      "min_severity": "critical",
      "description": "hardcoded API key",
      "soft": false
    }
  ]
}
```

Match rules (see `.e2e/assert/assert.go`):

1. Category equal (case-insensitive)
2. Path equal (exact)
3. Actual line range overlaps `[start_line - 3, end_line + 3]` (`LineTolerance = 3`)
4. Actual severity `>=` `min_severity`

Extras (no expected match) are reported but do not fail the case.

## Required secrets on this repo

| Secret | Purpose |
|---|---|
| `OPENAI_API_KEY` **or** `ANTHROPIC_API_KEY` | LLM provider for sacr |

That's the only one. The product repo is public, so `actions/checkout@v4` reads it with the default `GITHUB_TOKEN` — no PAT needed. Once sacr publishes a `v*` release, we'll swap `go build ./cmd/sacr` for `uses: shubam-disseqt/shubam-ai-code-reviewer@v0.2.0` and even the product-repo checkout goes away.

No secrets are needed on the product repo — cross-repo triggers are not used; this workflow is self-contained.

## Adding a case

1. Create branch `pr-N-branch` off `main` with the seeded bug(s). Do NOT add the ground truth to the branch
2. Add `.matrix/cases/pr-N-branch.json` on `main` describing every seeded finding
3. Open a persistent PR `pr-N-branch → main` — keep it open forever
4. Add the branch name to the `CASES` env in `.github/workflows/e2e-matrix.yml`
5. `gh workflow run e2e-matrix.yml` for a manual first check

## Local dry-run

```bash
# Clone once
git clone git@github.com:shubam-disseqt/zreview-e2e-matrix.git
cd zreview-e2e-matrix && git checkout pr-3-branch

# In another shell — build sacr from the product repo
git clone git@github.com:shubam-disseqt/shubam-ai-code-reviewer.git /tmp/product
cd /tmp/product && go build -o /tmp/sacr ./cmd/sacr

# Review the case
cd - && /tmp/sacr review --commit HEAD --format json --output /tmp/actual.json --repo .

# Build + run assert
cd .e2e && go build -o /tmp/assert ./assert
/tmp/assert -expected ../.matrix/cases/pr-3-branch.json -actual /tmp/actual.json -actual /tmp/actual-2.json -case pr-3-branch
```

## Gate

- `hard recall >= 80%` (aggregate across all cases' non-soft expected findings)
- **AND** no seeded `critical` finding is missed on any case

Below either bar → workflow fails.

## Cost expectations

Two `sacr review` calls per case (post + capture-json), 4 cases per run:

- ~5-15k input tokens × 2 calls × 4 cases = ~$0.10 per matrix run at `gpt-4o-mini`
- Nightly + on-demand ≈ 400 runs/year ≈ **~$40/year**

## Deferred

- `pr-2/pr-4/pr-6/pr-7/pr-8/pr-9` persistent PRs exist on this repo but their `.matrix/expected.json` isn't written yet — ~15 min per branch to derive from the feat commit
- Swap `go build ./cmd/sacr` → `uses: shubam-disseqt/shubam-ai-code-reviewer@v0.2.0` once the Action is published
- Multi-provider matrix (Sonnet vs gpt-4o-mini) side-by-side
