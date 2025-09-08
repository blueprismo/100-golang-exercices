# Exercise 75: Multiple HTTP Routes with HandleFunc

## Multiple Route Handling

Real web applications need multiple routes. Go's `http.HandleFunc()` lets you register different handlers for different URL patterns.

## Registering Multiple Routes

```go
func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/about", aboutHandler)
    http.HandleFunc("/contact", contactHandler)
    
    http.ListenAndServe(":8080", nil)
}
```

## Handler Function Signature

All handler functions must match this signature:
```go
func handlerName(w http.ResponseWriter, r *http.Request) {
    // Handle the request
}
```

## DefaultServeMux

When you use `http.HandleFunc()` and `http.ListenAndServe()` with `nil`, Go uses the `DefaultServeMux` (default HTTP request multiplexer) to route requests.

## Route Patterns

- **Exact match**: `/users` matches only `/users`
- **Prefix match**: `/static/` matches `/static/css/style.css`
- **Root pattern**: `/` matches everything not matched by other patterns

## Your Task

Create an HTTP server with multiple routes:
1. Create `handler_1` function that writes "Hello from HandleFunc #1"
2. Create `handler_2` function that writes "Hello from HandleFunc #2"
3. Register both handlers for different URL patterns
4. Test that each route works correctly

## Testing Multiple Routes

Visit different URLs to test:
- Each route should respond with its specific message
- Different handlers for different paths

This demonstrates building multi-page web applications.