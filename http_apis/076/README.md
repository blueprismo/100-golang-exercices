# Exercise 76: Understanding ServeMux

## What is ServeMux?

`ServeMux` (HTTP request multiplexer) is Go's built-in HTTP router. It matches incoming requests to registered patterns and calls the appropriate handler.

## DefaultServeMux vs Custom ServeMux

**DefaultServeMux** (what you've been using):
```go
http.HandleFunc("/path", handler)  // Registers with DefaultServeMux
http.ListenAndServe(":8080", nil)  // nil = use DefaultServeMux
```

**Custom ServeMux**:
```go
mux := http.NewServeMux()
mux.HandleFunc("/path", handler)
http.ListenAndServe(":8080", mux)  // Use custom mux
```

## Why Use Custom ServeMux?

- **Isolation**: Multiple servers with different routes
- **Testing**: Easier to test with isolated mux
- **Control**: More explicit about routing
- **Security**: DefaultServeMux is global and can be modified by other packages

## Your Task

Create and use a custom ServeMux instead of the DefaultServeMux:
1. Create a new ServeMux with `http.NewServeMux()`
2. Register your handlers with the custom mux
3. Pass the custom mux to `ListenAndServe`

This gives you explicit control over HTTP routing.