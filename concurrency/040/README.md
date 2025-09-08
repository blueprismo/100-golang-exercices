# Exercise 40: Channels and Timeouts

## Implementing Timeouts with Select

Timeouts are crucial in concurrent programming to avoid waiting indefinitely for slow or unresponsive operations. Go makes timeouts elegant using the `select` statement combined with `time.After()`.

## The time.After() Function

`time.After(duration)` returns a channel that will receive a value after the specified duration:

```go
timeout := time.After(3 * time.Second)
<-timeout  // Blocks for 3 seconds
```

## Timeout Pattern with Select

Combine `time.After()` with `select` to create timeouts:

```go
select {
case result := <-workChannel:
    fmt.Println("Work completed:", result)
case <-time.After(5 * time.Second):
    fmt.Println("Operation timed out!")
}
```

## Why Timeouts Matter

- **Prevent hanging**: Don't wait forever for slow operations
- **Resource management**: Free up goroutines that are waiting
- **User experience**: Provide feedback when operations take too long
- **System reliability**: Detect and handle unresponsive components

## Your Task

Create a timeout mechanism where:
1. A goroutine sleeps for 10 seconds then sends a message
2. The main program waits for either:
   - The message from the channel, OR
   - A timeout after 3 seconds
3. Since the timeout (3s) is shorter than the work time (10s), the timeout should occur first

## Expected Behavior

You should see "Timeout!!!!" printed because the 3-second timeout occurs before the 10-second work completes.

## Key Concepts

- `time.After()` creates a timeout channel
- `select` waits for whichever happens first
- This pattern is non-blocking and deterministic