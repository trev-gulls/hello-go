---
status: pending
created: 2026-03-05
updated: 2026-03-05
blocked-by: []
---

# Configure Claude Code sandbox for Go toolchain

## Goal

Allow `go build`, `go test`, and `go vet` to run inside the Claude Code sandbox without requiring `dangerouslyDisableSandbox`.

## Context

Go commands fail in sandbox mode because the build cache at `~/Library/Caches/go-build/` is outside the sandbox write allowlist. `git init` also fails because it writes to `.git/hooks/`.

## Acceptance Criteria

- [ ] Go commands (`go build`, `go test`, `go vet`) run without sandbox bypass
- [ ] `git init` runs without sandbox bypass
- [ ] Sandbox config documented in `.claude/CLAUDE.md` or settings

## Notes

- Sandbox write allowlist is configured in `.claude/settings.json` or `~/.claude/settings.json`
- May need to add `~/Library/Caches/go-build/` and `.git/hooks/` to allowed write paths
