# Exercise 61: Context with Timeout

## What is context.WithTimeout()?

`context.WithTimeout()` creates a context that automatically cancels after a specified duration. It's the most commonly used timeout mechanism in Go.

```go
ctx, cancel := context.WithTimeout(parentCtx, 5*time.Second)
defer cancel()
```

## WithTimeout vs WithDeadline

```go
// These are equivalent:
deadline := time.Now().Add(5 * time.Second)
ctx1, cancel1 := context.WithDeadline(parentCtx, deadline)

ctx2, cancel2 := context.WithTimeout(parentCtx, 5*time.Second)
```

`WithTimeout` is just a convenience function that calculates the deadline for you.

## Common Timeout Patterns

**HTTP requests:**
```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
req = req.WithContext(ctx)
```

**Database queries:**
```go
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
rows, err := db.QueryContext(ctx, "SELECT * FROM users")
```

**API calls:**
```go
ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
defer cancel()
response, err := client.GetWithContext(ctx, url)
```

## Why Always defer cancel()?

Even if the timeout hasn't been reached, you should always call `cancel()` to:
1. **Free resources**: Clean up the timeout timer
2. **Signal completion**: Tell child contexts to stop
3. **Best practice**: Ensures proper cleanup regardless of how function exits

## Timeout Behavior

1. **Normal completion**: Operation finishes before timeout
2. **Timeout**: Context cancels itself after duration
3. **Manual cancel**: You can still call cancel() early
4. **Error reporting**: `ctx.Err()` returns `context.DeadlineExceeded`

## Your Task

1. **Create a context with 1.5-second timeout**
2. **Complete the select statement** to handle both sending and timeout
3. **Observe the race**: Will the 3-second loop complete, or will the 1.5-second timeout win?
4. **Handle the timeout gracefully** by breaking out of the loop

## Expected Behavior

Since the timeout (1.5s) is shorter than the time needed to complete the loop (3+ seconds), the context should cancel before all numbers are sent.