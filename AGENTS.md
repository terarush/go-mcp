# AGENTS.md

## Quick Reference

| Task | Command |
|------|---------|
| Install deps | `go mod download` |
| Build | `go build ./...` |
| Test | `go test -v ./...` |
| Vet | `go vet ./...` |
| Format | `gofmt -w .` |
| Run (stdio) | `go run ./cmd/stdio` |
| Run (HTTP) | `go run ./cmd/http` |

## Project Structure

- `cmd/stdio/main.go`: stdio transport entrypoint
- `cmd/http/main.go`: Streamable HTTP transport entrypoint (`/mcp` and `/health`)
- `internal/server/server.go`: MCP server setup and capability configuration
- `internal/server/tools.go`: MCP tool registrations
- `internal/server/resources.go`: MCP resource registrations
- `internal/server/prompts.go`: MCP prompt registrations
