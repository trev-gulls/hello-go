# tgulls-hello

A simple CLI that prints "Hello, Go!"

Part of the `tgulls` package family alongside [`@tgulls/hello`](https://www.npmjs.com/package/@tgulls/hello) (npm) and [`tgulls-hello`](https://pypi.org/project/tgulls-hello/) (PyPI).

## Install

```sh
go install github.com/trev-gulls/hello-go/cmd/tgulls-hello@latest
```

## Run without installing

```sh
go run github.com/trev-gulls/hello-go/cmd/tgulls-hello@latest
```

## Development

```sh
go run ./cmd/tgulls-hello        # Run without building
go build ./cmd/tgulls-hello      # Build binary
go test ./...                    # Run tests
```
