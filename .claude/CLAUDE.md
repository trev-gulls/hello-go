# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

`hello-go` — a minimal CLI Go module that prints "Hello, Go!" Part of the `tgulls` package family alongside `@tgulls/hello` (npm) and `tgulls-hello` (PyPI).

Module path: `github.com/trev-gulls/hello-go`

## Commands

```sh
go build ./cmd/tgulls-hello        # Build binary (outputs to current dir)
go run ./cmd/tgulls-hello          # Run without building
go test ./...                      # Run all tests
go test -v -run TestName ./...     # Run a single test by name
```

## Architecture

Uses the `cmd/` layout so the repo name (`hello-go`) and binary name (`tgulls-hello`) can differ, and additional commands can be added later.

```
cmd/
  tgulls-hello/
    main.go          # CLI entry point
    main_test.go     # Subprocess-based CLI test
go.mod               # Module definition
```

## Publishing

Go modules are published by pushing a tagged commit to GitHub. No separate registry upload needed — pkg.go.dev indexes automatically.

```sh
git tag v0.1.0
git push origin v0.1.0
```

Users install with:

```sh
go install github.com/trev-gulls/hello-go/cmd/tgulls-hello@latest
```
