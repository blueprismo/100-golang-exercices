# Exercise 60: Context with Deadline

## What is context.WithDeadline()?

`context.WithDeadline()` creates a context that automatically cancels at a specific time. It's perfect when you know exactly when an operation should stop.

```go
deadline := time.Now().Add(5 * time.Second)
ctx, cancel := context.WithDeadline(parentCtx, deadline)
defer cancel()
```

## Deadline vs Timeout

| WithDeadline | WithTimeout |
|-------------|-------------|
| Specific time point | Duration from now |
| `time.Now().Add(duration)` | Just the duration |
| More control | Simpler to use |

## How Deadlines Work

1. **Set absolute time**: "Cancel at exactly 3:30 PM"
2. **Automatic cancellation**: Context cancels itself when deadline passes
3. **Done channel closes**: `ctx.Done()` receives notification
4. **Error available**: `ctx.Err()` returns deadline exceeded error

## Context Errors

When a context is cancelled, `ctx.Err()` returns the reason:

```go
case <-ctx.Done():
    err := ctx.Err()
    if err == context.DeadlineExceeded {
        fmt.Println("Operation timed out")
    } else if err == context.Canceled {
        fmt.Println("Operation was cancelled")
    }
```

## Racing Against Time

This exercise demonstrates a race condition:
- **Loop tries to send** 3 numbers, sleeping 1 second between each
- **Context cancels after 1.5 seconds**
- **Which wins?** The loop (3+ seconds) or the deadline (1.5 seconds)?

## Your Task

1. **Create context with deadline** using the provided 1.5-second deadline
2. **Modify the loop** to check for context cancellation
3. **Handle both cases** in select:
   - Sending numbers to channel
   - Receiving cancellation signal
4. **Observe the behavior**: Does the loop complete or get cancelled?

## Key Learning

Deadlines are perfect for operations where you know exactly when to give up. The context will automatically cancel, making your code responsive to time constraints.