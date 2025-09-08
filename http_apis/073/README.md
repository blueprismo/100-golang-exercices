# Exercise 73: Dynamic HTTP Routes

## Dynamic Route Responses

Instead of hardcoding responses, you can make your server dynamically respond based on the requested URL path.

## Accessing Request Information

The `*http.Request` parameter contains information about the request:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path        // "/some/path"
    method := r.Method        // "GET", "POST", etc.
    host := r.Host           // "localhost:8080"
}
```

## HTML Escaping

When including user input (like URL paths) in responses, use `html.EscapeString()` to prevent XSS attacks:

```go
safeOutput := html.EscapeString(userInput)
```

## Dynamic Response Example

```go
func handler(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path
    response := fmt.Sprintf("Hello, %s", html.EscapeString(path))
    fmt.Fprintf(w, response)
}
```

## Your Task

Create an HTTP server where:
1. The handler responds with "Hello, [path]" 
2. The [path] part is the actual URL path requested
3. If you visit `/bar`, it responds "Hello, /bar"
4. If you visit `/newpath`, it responds "Hello, /newpath"
5. Use proper HTML escaping for security

## Testing Different Paths

- http://localhost:8080/bar → "Hello, /bar"
- http://localhost:8080/test → "Hello, /test"  
- http://localhost:8080/anything → "Hello, /anything"

This demonstrates how to create dynamic web responses based on the request.