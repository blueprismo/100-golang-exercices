# Exercise 55: Context Background

## Understanding context.Background()

`context.Background()` returns a non-nil, empty context. It's the root of any context tree and is typically used at the top level of your application.

## When to Use context.Background()

Use `context.Background()` when:
- **Starting a new context tree**: At the beginning of request handling
- **Main functions**: As the root context for your application
- **Top-level operations**: When you need a context but don't have one from above
- **Tests**: As a starting point for test contexts

## Background vs TODO

| context.Background() | context.TODO() |
|---------------------|----------------|
| Production code | Development/placeholder |
| Clear intent | Temporary/uncertain |
| Root of context tree | "I'll figure this out later" |

## Example Usage

```go
func main() {
    // Root context for the application
    ctx := context.Background()
    
    // Pass it down to other functions
    processRequest(ctx)
}

func processRequest(ctx context.Context) {
    // This function receives context from above
    // and can create child contexts if needed
}
```

## Context Best Practices

1. **Pass context as first parameter**: `func DoWork(ctx context.Context, ...)`
2. **Don't store contexts**: Pass them through function calls
3. **Don't pass nil**: Use `context.Background()` instead
4. **Accept contexts in APIs**: Even if you don't use them immediately

## Your Task

This exercise builds on the previous one:
1. **Use `context.Background()`** instead of `context.TODO()`
2. **Understand the semantic difference**: Background indicates this is a proper root context
3. **Call your function** with this background context

The functionality is identical to the previous exercise, but the intent is clearer - you're establishing a root context for your operation.