---
status: pending
created: 2026-03-05
updated: 2026-03-05
blocked-by: []
---

# Create Go module version of @tgulls/hello

## Goal

Publish a Go equivalent of the `@tgulls/hello` CLI as a Go module, so users can run via `go install github.com/trev-gulls/hello-go/cmd/tgulls-hello@latest` or `go run github.com/trev-gulls/hello-go/cmd/tgulls-hello@latest`.

## Acceptance Criteria

- [ ] Go module with `cmd/tgulls-hello/` containing a `main` package that prints "Hello, Go!"
- [ ] Uses `go.mod` with appropriate module path
- [ ] CLI test that runs the command in a subprocess and verifies output
- [ ] Tagged and pushed to GitHub (makes it available on pkg.go.dev)
- [ ] README with `go install` and `go run` instructions

## Notes

- Mirror the structure of the npm and PyPI packages — keep it minimal
- Go modules are fetched from VCS, not a registry — tagging a release on GitHub is the equivalent of `npm publish` / `uv publish`
- Module path should match the GitHub repo: `github.com/trev-gulls/hello-go`
