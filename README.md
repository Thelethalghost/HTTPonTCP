# HTTP Parser from TCP

A lightweight Go project that parses raw HTTP/1.1 requests directly from TCP connections. This demonstrates core networking concepts and protocol implementation from first principles.

## Overview

This project implements a complete HTTP/1.1 request parser that reads raw TCP streams and extracts request components without relying on Go's built-in HTTP libraries. It includes two command-line tools: a TCP listener that parses and displays incoming requests, and an HTTP server that responds to connections.

## Features

- **Full HTTP/1.1 Request Parsing**: Parses request lines, headers, and message bodies according to HTTP/1.1 specification
- **Streaming Parser**: Processes data incrementally as it arrives over the network, handling variable-sized chunks
- **Header Validation**: Validates header field names according to RFC 7230 token rules
- **Content-Length Support**: Correctly handles request bodies using the Content-Length header
- **Case-Insensitive Header Lookup**: Headers are stored and retrieved in a case-insensitive manner
- **Comprehensive Test Coverage**: Includes unit tests for request parsing and header validation

## Project Structure
```
├── cmd/
│   ├── httpserver/    # HTTP server that responds to connections on port 42069
│   └── tcplistener/   # TCP listener that parses and displays raw HTTP requests
├── internal/
│   ├── request/       # HTTP request parsing logic
│   ├── headers/       # HTTP header parsing and management
│   └── server/        # Server implementation
├── go.mod            # Go module definition
└── makefile          # Build automation
```

## Building & Running

### Build
```bash
make build
```

### Run HTTP Server
```bash
make run
```
The server listens on port 42069 and responds with a simple "Hello World!" response.

### Run Tests
```bash
make test
```
Tests cover edge cases including malformed headers, variable chunk sizes, and incomplete requests.

## Technical Highlights

**Request Parsing State Machine**: The parser uses a finite state machine with four states (Init, Headers, Body, Done) to handle the sequential nature of HTTP request parsing.

**Incremental Parsing**: The `RequestFromReader` function demonstrates non-blocking, incremental parsing that works with network streams of arbitrary chunk sizes—critical for real network scenarios.

**Header Field Validation**: Implements RFC 7230 compliant validation for header field names using regex pattern matching.

**Dual-Value Headers**: Properly handles headers that appear multiple times by concatenating values with commas, per HTTP specification.

## Example Usage

### Sending a Request
```bash
curl http://localhost:42069/path
```

### TCP Listener Output
```
Request Line: 
 - Method: GET
 - Target: /path
 - Version: 1.1
Headers:
- host: localhost:42069
- user-agent: curl/7.81.0
- accept: */*
Body:
```

## Learning Outcomes

This project demonstrates:
- Low-level network programming with Go's `net` package
- Protocol implementation and RFC compliance
- State machine design patterns
- Streaming data processing
- Comprehensive unit testing with variable input scenarios
- Graceful signal handling and server lifecycle management

## Technologies

- **Language**: Go 1.24.6
- **Testing**: testify (assertions and requirements)
- **Networking**: Go standard library (`net`, `io`)

---

