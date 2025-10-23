// Tencent is pleased to support the open source community by making trpc-mcp-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-mcp-go is licensed under the Apache License Version 2.0.

package mcp

import (
	"go.uber.org/zap/zapcore"
	"trpc.group/trpc-go/trpc-mcp-go/internal/log"
)

// Logger defines the logging interface used throughout MCP framework.
type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
}

// LogLevel represents log level constants for user-friendly configuration.
type LogLevel int

const (
	// LogLevelDebug enables debug and above logs.
	LogLevelDebug LogLevel = iota - 1
	// LogLevelInfo enables info and above logs (default).
	LogLevelInfo
	// LogLevelWarn enables warn and above logs.
	LogLevelWarn
	// LogLevelError enables error and above logs.
	LogLevelError
)

// toZapLevel converts LogLevel to zapcore.Level.
func (l LogLevel) toZapLevel() zapcore.Level {
	switch l {
	case LogLevelDebug:
		return zapcore.DebugLevel
	case LogLevelInfo:
		return zapcore.InfoLevel
	case LogLevelWarn:
		return zapcore.WarnLevel
	case LogLevelError:
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

// NewZapLogger returns a Logger interface with Info level (default), hiding zap details.
func NewZapLogger() Logger {
	return log.NewZapLogger()
}

// NewZapLoggerWithLevel returns a Logger interface with custom log level.
// Example:
//
//	debugLogger := mcp.NewZapLoggerWithLevel(mcp.LogLevelDebug)
//	mcp.SetDefaultLogger(debugLogger)
func NewZapLoggerWithLevel(level LogLevel) Logger {
	return log.NewZapLoggerWithLevel(level.toZapLevel())
}

var (
	// TEMPORARY: Set default logger to Debug level for debugging
	// TODO: Revert to NewZapLogger() (Info level) before release
	defaultLogger Logger = NewZapLoggerWithLevel(LogLevelDebug)
)

// SetDefaultLogger sets the global default logger.
func SetDefaultLogger(l Logger) {
	defaultLogger = l
}

// GetDefaultLogger returns the global default logger.
func GetDefaultLogger() Logger {
	return defaultLogger
}
