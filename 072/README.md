# Exercise 72: Basic HTTP Server

## Introduction to net/http

Go's `net/http` package makes it incredibly easy to create web servers. You can build a complete HTTP server with just a few lines of code.

## Simple HTTP Server

```go
package main

import (
    "net/http"
    "io"
)

func handler(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello, Web!")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

## Key Components

**Handler Function:**
- Takes `http.ResponseWriter` and `*http.Request`
- Writes response back to client

**http.HandleFunc():**
- Registers a handler for a URL pattern

**http.ListenAndServe():**
- Starts the server on specified port

## Testing Your Server

Once running, test with:
- **Browser**: http://localhost:8080
- **curl**: `curl http://localhost:8080`
- **wget**: `wget -qO- http://localhost:8080`

## Your Task

Create a simple HTTP server that:
1. Responds with "Hello web World" 
2. Uses the `net/http` package
3. Runs on any port you choose
4. Can be tested with a browser or curl

This is the foundation of web development in Go!