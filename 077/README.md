# Exercise 77: ServeMux with Stateful Handlers

## Stateful HTTP Handlers

Sometimes handlers need to maintain state (counters, databases, configuration). You can create struct-based handlers that implement the `http.Handler` interface.

## Struct-Based Handlers

```go
type Counter struct {
    mu sync.Mutex
    n  int
}

func (c *Counter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.n++
    fmt.Fprintf(w, "Count: %d", c.n)
}
```

## Thread Safety

When multiple requests access shared state, use synchronization:
- **sync.Mutex**: For exclusive access
- **sync.RWMutex**: For read-heavy workloads
- **atomic**: For simple counters

## Using with ServeMux

```go
counter := &Counter{}
mux := http.NewServeMux()
mux.Handle("/count", counter)  // Note: Handle, not HandleFunc
```

## Your Task

Create a Counter struct that:
1. Implements the `http.Handler` interface
2. Maintains an internal counter
3. Increments the counter on each request
4. Returns the current count
5. Uses proper synchronization for thread safety

This demonstrates stateful web handlers and concurrent safety.