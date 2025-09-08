# Exercise 33: Basic Channels

## What are Channels?

Channels are Go's way of allowing goroutines to communicate and synchronize. They follow Go's concurrency principle: **"Don't communicate by sharing memory; share memory by communicating."**

Think of channels as pipes through which you can send and receive values between goroutines.

## Channel Basics

**Creating channels:**
```go
ch := make(chan string)  // Channel that carries strings
ch := make(chan int)     // Channel that carries integers
```

**Channel operations:**
- **Send**: `ch <- value` (send value into channel)
- **Receive**: `value := <-ch` (receive value from channel)

## Unbuffered Channels (Synchronous)

By default, channels are unbuffered, which means:
- **Sends block** until another goroutine is ready to receive
- **Receives block** until another goroutine is ready to send
- This creates a **synchronization point** between goroutines

## Communication Pattern

The typical pattern is:
1. Create a channel
2. Start one or more goroutines that will send/receive on the channel
3. Use the channel operations to coordinate between goroutines

## Why Channels are Better than Sleep

Instead of guessing how long to wait with `time.Sleep()`, channels provide **proper synchronization**:
- Goroutines wait exactly as long as needed
- No race conditions
- Clear communication intent

## Your Task

Look at the `main.go` file and complete the exercise. You need to:
1. Understand how to send messages through channels
2. Understand how to receive messages from channels
3. Set up communication between two goroutines

## Expected Behavior

When completed, you should see `f2` printing a message that includes the message sent from `f1`. The program will coordinate properly without needing to guess timing.

## Key Points to Remember

- Channels are typed - a `chan string` only carries strings
- Unbuffered channel operations are synchronous
- The arrow `<-` shows the direction of data flow
- Channel operations will block until both sides are ready