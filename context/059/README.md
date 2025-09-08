# Exercise 59: Context Cancellation

## What is Context Cancellation?

Context cancellation is a mechanism to signal that work should be stopped. It's essential for:
- **Resource cleanup**: Stop operations when they're no longer needed
- **Timeout handling**: Cancel slow operations
- **User cancellation**: Stop work when users cancel requests
- **Graceful shutdown**: Clean up when the application is stopping

## context.WithCancel()

`context.WithCancel()` creates a cancellable context:

```go
ctx, cancel := context.WithCancel(parentCtx)
defer cancel() // Always call cancel to free resources
```

Returns:
- **ctx**: The new cancellable context
- **cancel**: A function that cancels the context when called

## The Done Channel

Cancelled contexts signal through their `Done()` channel:

```go
select {
case <-ctx.Done():
    fmt.Println("Context cancelled!")
    return ctx.Err() // Returns cancellation reason
case result := <-workChannel:
    fmt.Println("Work completed:", result)
}
```

## Select with Context

The common pattern for cancellable operations:

```go
for {
    select {
    case <-ctx.Done():
        fmt.Println("Stopping due to cancellation")
        return
    case data := <-dataChannel:
        fmt.Println("Processing:", data)
    }
}
```

## Why Use Select with Context?

- **Responsive cancellation**: Check for cancellation between operations
- **Clean shutdown**: Exit loops and functions when cancelled
- **Resource management**: Stop goroutines that are no longer needed

## Your Task

This is a complex exercise involving:

1. **Create a cancellable context** using `context.WithCancel()`
2. **Create an unbuffered int channel** for communication
3. **Start a goroutine** (`doAnother`) that waits for cancellation or data
4. **Send data through the channel** (numbers 1-3)
5. **Cancel the context** to stop the goroutine
6. **Use select in the goroutine** to handle both cancellation and data

## Key Concepts

- `ctx.Done()` returns a channel that closes when context is cancelled
- Always call the cancel function to clean up resources
- Use select to make your goroutines responsive to cancellation