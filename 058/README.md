# Exercise 58: Testing Context Immutability

## Proving Context Immutability

This exercise demonstrates that contexts are truly immutable. When you create a new context with `context.WithValue()`, the original context remains unchanged.

## Understanding the Flow

The exercise shows three contexts in action:
1. **Original context** in main: Has "Name" = "John"
2. **Modified context** in doSomething: Has "Name" = "Mary" 
3. **Original context again**: Still has "Name" = "John"

## Why This Matters

Context immutability is crucial for:
- **Thread safety**: Multiple goroutines can safely use the same context
- **Predictable behavior**: Functions can't accidentally modify contexts from callers
- **Isolation**: Changes in one part of your code don't affect others

## The Context Tree Structure

```
Background Context
    |
    └── WithValue(ctx, "Name", "John")  ← main's context
            |
            └── WithValue(ctx, "Name", "Mary")  ← doSomething's child context
```

Each level maintains its own value, and parent contexts are never modified.

## Real-World Implications

```go
func handleRequest(ctx context.Context) {
    // Add request ID to context
    reqCtx := context.WithValue(ctx, "requestID", "12345")
    
    // Call other functions with enhanced context
    processData(reqCtx)
    
    // Original ctx is unchanged - other parts of the app
    // won't see the requestID
}
```

## Your Task

1. **Add a print statement** in `doSomething` after calling `doAnother`
2. **Print the original context's value** to see if it changed
3. **Observe the output**: The original context should still have "John"
4. **Understand the immutability**: The child context's "Mary" doesn't affect the parent

## Expected Behavior

You should see:
- doSomething prints "John" (original value)
- doAnother prints "Mary" (modified value)  
- doSomething prints "John" again (proving immutability)