// Package main runs the MCP server over stdio transport.
package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"mcp/internal/server"
)

func main() {
	if err := run(); err != nil {
		log.Printf("Failed to connect: %v", err)
		os.Exit(1)
	}
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		cancel()
	}()

	srv := server.NewServer()
	transport := &mcp.StdioTransport{}

	log.SetOutput(os.Stderr)
	log.Println("MCP Server running on stdio")

	if _, err := srv.Connect(ctx, transport, nil); err != nil {
		return err
	}

	<-ctx.Done()
	log.Println("Server shutting down")
	return nil
}
