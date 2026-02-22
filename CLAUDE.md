# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Go learning/study project for mastering Go programming.

## Documentation Style

- Korean language for comments and documentation (학습 기록용)
- Commit messages in Korean

## Commands

```bash
# Run the application
go run main.go

# Build the application
go build -o bin/main main.go

# Format code
go fmt ./...

# Add a dependency
go get <package-path>

# Tidy dependencies (remove unused, add missing)
go mod tidy
```

## Go Module

- Module name: `study/master`
- Go version: 1.25.0
- Dependencies are managed via `go.mod` (similar to package.json in Node.js)
- Use `go get` to add external packages - they are automatically recorded in go.mod
