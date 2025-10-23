# Debug Logging Example

This example demonstrates how to configure log levels in trpc-mcp-go.

## Problem

By default, the framework uses **Info** level logging. If you want to see debug logs (e.g., for troubleshooting), you need to explicitly configure the logger.

## Solution

The framework now provides `NewZapLoggerWithLevel()` to create loggers with custom log levels.

## Available Log Levels

```go
mcp.LogLevelDebug  // Most verbose - shows all logs
mcp.LogLevelInfo   // Default - shows info and above
mcp.LogLevelWarn   // Shows warnings and errors only
mcp.LogLevelError  // Least verbose - errors only
```

## Usage

### 1. Set Global Default Logger (Affects All Components)

```go
// Enable debug logging globally
debugLogger := mcp.NewZapLoggerWithLevel(mcp.LogLevelDebug)
mcp.SetDefaultLogger(debugLogger)

// All servers and clients created after this will use debug logging
server := mcp.NewServer("My-Server", "1.0.0")
client := mcp.NewClient("http://localhost:3000", clientInfo)
```

### 2. Set Logger Per-Server

```go
// This server will only show warnings and errors
warnLogger := mcp.NewZapLoggerWithLevel(mcp.LogLevelWarn)
server := mcp.NewServer(
    "My-Server",
    "1.0.0",
    mcp.WithServerLogger(warnLogger),
)
```

### 3. Set Logger Per-Client

```go
// This client will show debug logs
debugLogger := mcp.NewZapLoggerWithLevel(mcp.LogLevelDebug)
client := mcp.NewClient(
    "http://localhost:3000",
    clientInfo,
    mcp.WithClientLogger(debugLogger),
)
```

## Running the Example

```bash
cd examples/debug-logging
go run main.go
```

You'll see debug logs from the framework internals!

## When to Use Debug Logging

- **Development**: See detailed request/response flow
- **Troubleshooting**: Diagnose connection issues
- **Integration**: Understand middleware execution order
- **Performance**: Identify bottlenecks

## When to Use Other Levels

- **Production**: Use `LogLevelInfo` or `LogLevelWarn` to reduce noise
- **CI/CD**: Use `LogLevelError` for clean test output
- **Performance-critical**: Use `LogLevelWarn` to minimize overhead

## Example Output

With `LogLevelDebug`:
```
2025-10-23 16:30:45.123 DEBUG Initializing MCP handler
2025-10-23 16:30:45.124 DEBUG Registering tool: greet
2025-10-23 16:30:45.125 INFO  MCP server started on :3000
2025-10-23 16:30:50.456 DEBUG Received request: tools/call
2025-10-23 16:30:50.457 DEBUG Processing greet tool for name: Alice
2025-10-23 16:30:50.458 DEBUG Sending response
```

With `LogLevelInfo` (default):
```
2025-10-23 16:30:45.125 INFO  MCP server started on :3000
```

## API Reference

```go
// Create logger with custom level
func NewZapLoggerWithLevel(level LogLevel) Logger

// Set as global default
func SetDefaultLogger(l Logger)

// Use in server options
func WithServerLogger(logger Logger) ServerOption

// Use in client options
func WithClientLogger(logger Logger) ClientOption
```

