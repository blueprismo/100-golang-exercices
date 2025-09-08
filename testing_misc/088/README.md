# Exercise 88: HTTP Client with Heimdall

## What is Heimdall?

Heimdall is an enhanced HTTP client for Go that provides additional features like automatic retries, circuit breakers, timeouts, and better error handling. It wraps Go's standard `net/http` client with production-ready features.

## Why Use Enhanced HTTP Clients?

Standard `http.Client` is basic. Enhanced clients provide:
- **Automatic Retries**: Retry failed requests with backoff
- **Circuit Breakers**: Prevent cascading failures
- **Timeouts**: Fine-grained timeout control
- **Metrics**: Built-in request/response metrics
- **Middleware**: Request/response interceptors

## Heimdall Features

```go
// Basic client with timeout
client := httpclient.NewClient(httpclient.WithHTTPTimeout(10*time.Second))

// Client with retries
client := httpclient.NewClient(
    httpclient.WithHTTPTimeout(10*time.Second),
    httpclient.WithRetryCount(3),
)
```

## Installation and Import

```bash
go get -u github.com/gojek/heimdall/v7
```

```go
import "github.com/gojek/heimdall/v7/httpclient"
```

## Basic Usage Pattern

1. Create client with desired configuration
2. Use standard HTTP methods (GET, POST, etc.)
3. Handle response like standard `*http.Response`
4. Process body and handle errors

## HTTP Client Best Practices

- Always set timeouts to prevent hanging requests
- Handle network errors gracefully
- Close response bodies to prevent resource leaks
- Use context for cancellation when needed

## Your Task

Look at the main.go file and complete the exercise by:
1. Installing the Heimdall package
2. Creating an HTTP client with timeout configuration
3. Making a GET request to a webpage
4. Handling the response and reading the body

This exercise introduces you to enhanced HTTP clients and production-ready web service communication patterns.