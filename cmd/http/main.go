// Package main runs the MCP server over streamable HTTP transport.
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"mcp/internal/server"
)

func main() {
	if err := run(); err != nil {
		log.Printf("Server error: %v", err)
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	addr := fmt.Sprintf(":%s", port)

	log.Printf("MCP Server running on http://localhost%s", addr)
	log.Printf("  MCP endpoint: http://localhost%s/mcp", addr)
	log.Printf("  Health check: http://localhost%s/health", addr)

	httpServer := &http.Server{
		Addr:    addr,
		Handler: newHandler(),
	}

	go func() {
		<-ctx.Done()
		log.Println("Shutting down server...")
		_ = httpServer.Shutdown(context.Background())
	}()

	if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	return nil
}

func newHandler() http.Handler {
	handler := mcp.NewStreamableHTTPHandler(func(_ *http.Request) *mcp.Server {
		return server.NewServer()
	}, nil)

	mux := http.NewServeMux()
	mux.Handle("/mcp", http.NewCrossOriginProtection().Handler(handler))
	mux.HandleFunc("/health", healthHandler)
	return mux
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"ok","server":"mcp-server","version":"1.0.0"}`))
}
