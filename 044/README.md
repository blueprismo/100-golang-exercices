# Exercise 44: Multiple Goroutines with Channels

## Scaling Up: Multiple Producers

This exercise extends the previous concepts by using multiple goroutines sending to the same channel. This is a common pattern in concurrent programming called **fan-in**.

## Fan-In Pattern

**Fan-in** means multiple producers (goroutines) send data to a single channel, which is consumed by one or more receivers.

```go
     Producer 1 ──┐
                  ├─── Channel ─── Consumer
     Producer 2 ──┘
```

## Why Multiple Producers?

- **Parallel processing**: Multiple workers can produce data simultaneously
- **Load distribution**: Spread work across multiple goroutines
- **Resource utilization**: Make use of multiple CPU cores
- **Scalability**: Easy to add more producers as needed

## Natural Ordering

When multiple goroutines send to the same channel, messages arrive in the order they're sent, but since goroutines run concurrently, this creates a natural interleaving of messages.

## Key Concepts

1. **Shared channel**: Multiple goroutines can send to the same channel safely
2. **Concurrent execution**: All producer goroutines run simultaneously
3. **Single consumer**: One receiver (like your main function) gets all messages
4. **Non-deterministic order**: Message arrival order depends on goroutine scheduling

## Your Task

Create a system with multiple goroutines that:
1. Each runs a function that sends time-related data to channels
2. All goroutines run concurrently
3. Main function receives from all channels and displays the results
4. Observe how messages from different producers are interleaved

## Expected Behavior

You should see messages from different goroutines appearing in an interleaved, non-deterministic order. This demonstrates how multiple concurrent producers can feed data into a single processing pipeline.

## Real-World Applications

This pattern is used in:
- **Web servers**: Multiple request handlers sending responses
- **Data processing**: Multiple workers processing different parts of a dataset
- **Monitoring systems**: Multiple sensors sending data to a central collector