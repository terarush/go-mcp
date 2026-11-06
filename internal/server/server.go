// Package server defines and configures the MCP server.
package server

import "github.com/modelcontextprotocol/go-sdk/mcp"

// NewServer creates and configures the MCP server.
func NewServer() *mcp.Server {
	srv := mcp.NewServer(
		&mcp.Implementation{
			Name:    "mcp-server",
			Version: "1.0.0",
		},
		&mcp.ServerOptions{
			Capabilities: &mcp.ServerCapabilities{
				Tools:     &mcp.ToolCapabilities{},
				Resources: &mcp.ResourceCapabilities{},
				Prompts:   &mcp.PromptCapabilities{},
			},
		},
	)

	registerTools(srv)
	registerResources(srv)
	registerPrompts(srv)

	return srv
}
