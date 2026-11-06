package server

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type pingInput struct {
	Message string `json:"message,omitempty" jsonschema:"Optional message to echo back"`
}

func registerTools(srv *mcp.Server) {
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "ping",
		Description: "Health check tool to verify MCP server connectivity",
	}, pingHandler)
}

func pingHandler(_ context.Context, _ *mcp.CallToolRequest, input pingInput) (*mcp.CallToolResult, any, error) {
	msg := "pong"
	if input.Message != "" {
		msg = fmt.Sprintf("pong: %s", input.Message)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: msg},
		},
	}, nil, nil
}
