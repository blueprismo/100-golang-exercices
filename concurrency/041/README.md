# Exercise 41: Closing Channels

## Understanding Channel Closing

Channels can be closed to signal that no more values will be sent. This is a powerful mechanism for coordinating between goroutines and signaling completion.

## Why Close Channels?

- **Signal completion**: Tell receivers no more data is coming
- **Enable range loops**: `range` over channels requires closing
- **Resource cleanup**: Properly manage channel resources
- **Prevent goroutine leaks**: Allow goroutines to exit cleanly

## How to Close Channels

```go
ch := make(chan string)
ch <- "hello"
close(ch)  // No more values can be sent
```

## What Happens After Closing?

1. **Sending**: Panic if you try to send to a closed channel
2. **Receiving**: You can still receive remaining buffered values
3. **Empty closed channel**: Returns zero value for the type
4. **Range**: Loop exits when channel is closed and empty

## Detecting Closed Channels

Use the "comma ok" idiom:
```go
value, ok := <-ch
if !ok {
    // Channel is closed
}
```

## Your Task

Work with a buffered channel that will be closed. Learn how:
1. You can still read from a closed channel
2. What happens when you read from a closed, empty channel
3. The behavior difference between open and closed channels