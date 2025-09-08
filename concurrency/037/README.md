# Exercise 37: Channel Directions - Receive-Only Channels

## What are Channel Directions?

By default, channels are **bidirectional** - you can both send and receive on them. But Go allows you to restrict channels to only send or only receive, providing better type safety and clearer code intent.

## Receive-Only Channels

A receive-only channel can only be used to receive values:

**Declaration syntax:**
```go
func myFunction(ch <-chan string) {  // <-chan = receive-only
    value := <-ch  // ✓ OK - can receive
    // ch <- "hello"  // ✗ ERROR - cannot send
}
```

## Why Use Directional Channels?

1. **Type Safety**: Prevents accidental sends/receives
2. **Clear Intent**: Function signature shows exactly what the function does
3. **Better APIs**: Makes function contracts explicit
4. **Compiler Help**: Catches misuse at compile time

## Channel Direction Syntax

- **Bidirectional**: `chan string` (default)
- **Send-only**: `chan<- string` (arrow points into chan)
- **Receive-only**: `<-chan string` (arrow points out of chan)

## Automatic Conversion

Go automatically converts bidirectional channels to directional channels when needed:

```go
ch := make(chan string)        // Bidirectional
go sender(ch)                  // Converts to send-only automatically
go receiver(ch)                // Converts to receive-only automatically

func sender(out chan<- string) { ... }    // Send-only parameter
func receiver(in <-chan string) { ... }   // Receive-only parameter
```

## Your Task

Create a function called `receive` that:
1. Takes a receive-only channel as its first and only argument
2. Cannot send data to the channel (compiler will prevent this)
3. Can only receive data from the channel

You'll also need to:
- Create the channel in main
- Send some data to the channel
- Start the receive goroutine

## Expected Behavior

The program should demonstrate that the `receive` function can only receive from the channel, not send to it.

## Hint

Pay attention to the arrow direction in the channel type: `<-chan string` means "receive-only channel of strings".