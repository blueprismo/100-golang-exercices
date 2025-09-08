# Exercise 56: Context with Values

## Storing Data in Context

Context can carry request-scoped values across API boundaries. This is useful for passing data like user IDs, request IDs, or authentication tokens through your call chain.

## context.WithValue()

The `context.WithValue()` function creates a new context with a key-value pair:

```go
newCtx := context.WithValue(parentCtx, key, value)
```

Parameters:
- **parentCtx**: The parent context to derive from
- **key**: The key to store (should be comparable type)
- **value**: The value to store (can be any type)

## Retrieving Values

Use `ctx.Value(key)` to retrieve stored values:

```go
value := ctx.Value("userID")
if userID, ok := value.(string); ok {
    fmt.Println("User ID:", userID)
}
```

## Key Best Practices

1. **Use unexported key types**: Prevents key collisions
2. **Define key constants**: Make keys reusable and type-safe
3. **Use type assertions**: Values are returned as `interface{}`

## Example Key Pattern

```go
type contextKey string

const (
    userIDKey contextKey = "userID"
    requestIDKey contextKey = "requestID"
)

// Usage
ctx = context.WithValue(ctx, userIDKey, "user123")
userID := ctx.Value(userIDKey).(string)
```

## What Values to Store

**Good candidates**:
- User authentication info
- Request IDs for tracing
- API keys or tokens
- Request metadata

**Bad candidates**:
- Optional parameters (use function parameters instead)
- Large objects (consider passing directly)
- Frequently changing data

## Your Task

1. **Start with a background context**
2. **Use `context.WithValue()`** to add a key-value pair to the context
3. **In your `doSomething` function**, retrieve and print the value using `ctx.Value()`
4. **Choose meaningful key and value** (e.g., "name" -> "John")

This demonstrates how context can carry data through your function call chain.