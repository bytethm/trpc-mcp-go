// Tencent is pleased to support the open source community by making trpc-mcp-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-mcp-go is licensed under the Apache License Version 2.0.

package main

import (
	"context"
	"fmt"
	"log"

	mcp "trpc.group/trpc-go/trpc-mcp-go"
)

func main() {
	// Example 1: Set global default logger to Debug level
	// This affects ALL servers and clients created afterwards
	debugLogger := mcp.NewZapLoggerWithLevel(mcp.LogLevelDebug)
	mcp.SetDefaultLogger(debugLogger)

	log.Println("🔧 Creating MCP server with Debug logging...")

	// Create server - it will automatically use the debug logger
	mcpServer := mcp.NewServer(
		"Debug-Example-Server",
		"1.0.0",
		mcp.WithServerAddress(":3000"),
		mcp.WithServerPath("/mcp"),
		// Note: No need to pass logger explicitly if using global default
	)

	// Example 2: Set logger explicitly for a specific server
	// This overrides the global default for this server only
	warnLogger := mcp.NewZapLoggerWithLevel(mcp.LogLevelWarn)
	anotherServer := mcp.NewServer(
		"Warn-Only-Server",
		"1.0.0",
		mcp.WithServerAddress(":3001"),
		mcp.WithServerLogger(warnLogger), // Only warn and above
	)

	// Register a simple tool
	greetTool := mcp.NewTool("greet",
		mcp.WithDescription("A simple greeting tool."),
		mcp.WithString("name", mcp.Description("Name to greet.")))

	mcpServer.RegisterTool(greetTool, func(ctx context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name := "World"
		if nameArg, ok := req.Params.Arguments["name"]; ok {
			if nameStr, ok := nameArg.(string); ok && nameStr != "" {
				name = nameStr
			}
		}

		// These debug logs will now be visible!
		debugLogger.Debugf("Processing greet tool for name: %s", name)

		content := []mcp.Content{
			mcp.NewTextContent(fmt.Sprintf("Hello, %s!", name)),
		}

		return &mcp.CallToolResult{Content: content}, nil
	})

	log.Println("✅ Server configured with debug logging")
	log.Println("Available log levels:")
	log.Println("  - mcp.LogLevelDebug (most verbose)")
	log.Println("  - mcp.LogLevelInfo  (default)")
	log.Println("  - mcp.LogLevelWarn")
	log.Println("  - mcp.LogLevelError (least verbose)")
	log.Println("")
	log.Println("To change log level:")
	log.Println("  1. Global: mcp.SetDefaultLogger(mcp.NewZapLoggerWithLevel(mcp.LogLevelDebug))")
	log.Println("  2. Per-server: mcp.WithServerLogger(mcp.NewZapLoggerWithLevel(mcp.LogLevelDebug))")
	log.Println("  3. Per-client: mcp.WithClientLogger(mcp.NewZapLoggerWithLevel(mcp.LogLevelDebug))")
	log.Println("")
	log.Printf("Starting server at :3000...")

	if err := mcpServer.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	_ = anotherServer // Unused in this example, just to show the pattern
}
