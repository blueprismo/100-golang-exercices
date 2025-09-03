# Exercise 67: Type Assertions with Error Handling

## Handling Type Assertion Failures

When working with empty interfaces, type assertions can fail. It's important to handle these failures gracefully instead of letting your program panic.

## The Safe Approach

Always use the two-value form of type assertion:

```go
if value, ok := someInterface.(int); ok {
    // Successfully got an int
    fmt.Println("Value:", value)
} else {
    // Not an int, handle gracefully
    log.Println("Expected int, got something else")
}
```

## Why This Matters

In real applications, data might be:
- From external sources (APIs, files)
- User input
- Configuration files
- Dynamic at runtime

You can't always guarantee the types, so defensive programming is essential.

## Your Task

Build on the previous exercise but add proper error handling:
1. Use safe type assertions with the ok pattern
2. Handle cases where the type assertion fails
3. Use the `log` package to report errors
4. Prevent panics from invalid type assumptions

This teaches defensive programming practices essential for robust Go applications.