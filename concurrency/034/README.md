# Exercise 34: Buffered Channels

## Understanding Buffered vs Unbuffered Channels

So far you've worked with **unbuffered channels** (synchronous). Now let's explore **buffered channels** (asynchronous).

## Unbuffered Channels (Review)
- Capacity: 0
- **Synchronous**: Send blocks until receive happens
- **Direct handoff**: Values pass directly from sender to receiver
- Created with: `make(chan type)`

## Buffered Channels
- **Capacity**: Can hold multiple values
- **Asynchronous**: Send doesn't block until buffer is full
- **Internal storage**: Values are queued in the buffer
- Created with: `make(chan type, capacity)`

## How Buffered Channels Work

```go
ch := make(chan string, 3) // Buffer can hold 3 strings

ch <- "first"   // Doesn't block - goes into buffer
ch <- "second"  // Doesn't block - goes into buffer  
ch <- "third"   // Doesn't block - goes into buffer
ch <- "fourth"  // BLOCKS - buffer is full!
```

## Key Differences

**Unbuffered Channel Behavior:**
- Every send must have a matching receive ready
- Tight synchronization between goroutines
- Good for signaling and coordination

**Buffered Channel Behavior:**
- Sends don't block until buffer is full
- Receives don't block until buffer is empty
- Good for producer-consumer scenarios with different speeds

## When to Use Buffered Channels

- **Producer faster than consumer**: Buffer smooths out the difference
- **Batch processing**: Accumulate several values before processing
- **Decoupling**: Producer and consumer don't need to be synchronized

## Your Task

Complete the exercise in `main.go`. You need to:
1. Create a buffered channel with the specified capacity
2. Send multiple messages without blocking
3. Use the provided function to retrieve and display messages

## Expected Behavior

You should be able to send all 4 strings immediately (no blocking), then retrieve them one by one using the `pop_message` function.

## Important Notes

- Buffer capacity is fixed at creation time
- Buffered channels still block when full (sends) or empty (receives)
- The order of messages is preserved (FIFO - First In, First Out)