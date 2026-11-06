# MCP Server

Minimal Model Context Protocol (MCP) server starter in Go using `github.com/modelcontextprotocol/go-sdk`.

## Project Structure

```
.
├── cmd/
│   ├── stdio/
│   │   └── main.go        # stdio transport entrypoint
│   └── http/
│       ├── main.go        # Streamable HTTP transport entrypoint
│       └── main_test.go   # HTTP handler tests
├── internal/
│   └── server/
│       ├── server.go      # Server initialization
│       ├── tools.go       # Tool definitions
│       ├── resources.go   # Resource definitions
│       └── prompts.go     # Prompt definitions
├── tools/
│   └── go.mod             # Dev tool dependencies (air, goimports)
├── .air.toml              # Live reload configuration
├── .golangci.yml          # Linter configuration
├── go.mod
└── Makefile
```

## Quick Start

```bash
# Run over stdio
make run-stdio

# Run over HTTP (default port 3000)
make run-http

# Build binaries (bin/stdio, bin/http)
make build

# Run tests
make test
```
