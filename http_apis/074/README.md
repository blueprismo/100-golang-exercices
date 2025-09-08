# Exercise 74: HTTP Handler Interface

## Understanding http.Handler Interface

The `http.Handler` interface is at the core of Go's HTTP handling. Any type that implements this interface can handle HTTP requests.

## The Handler Interface

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

Any type with a `ServeHTTP` method automatically implements this interface.

## Creating Custom Handlers

**Method 1: Custom Type**
```go
type MyHandler struct{}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from custom handler!")
}
```

**Method 2: Handler Function**
```go
func myHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from handler function!")
}

// Convert function to Handler using http.HandlerFunc
handler := http.HandlerFunc(myHandler)
```

## Handler vs HandlerFunc

- **http.Handler**: Interface requiring ServeHTTP method
- **http.HandlerFunc**: Type that converts functions to handlers
- **http.HandleFunc()**: Convenience function for registering handlers

## Your Task

This exercise focuses on understanding the `http.Handler` interface:
1. Create a custom type that implements `http.Handler`
2. Implement the `ServeHTTP` method
3. Understand how this interface enables flexible HTTP handling

This builds foundation knowledge for more advanced HTTP middleware and routing.