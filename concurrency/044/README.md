# Exercise 44: Multiple Goroutines with Channels

## Range with Unbuffered Channels

This exercise explores using `range` with unbuffered channels and goroutines producing data in real-time. Unlike buffered channels, unbuffered channels require a goroutine to be actively sending data.

## Unbuffered vs Buffered Channels

**Buffered channels** (previous exercise):
- Can hold multiple values before blocking
- Must be closed before using `range`
- Values are stored in the buffer

**Unbuffered channels** (this exercise):
- Block on send until receiver is ready
- Block on receive until sender is ready
- Perfect for real-time data streaming
- Don't need to be closed for `range` (but goroutine must keep sending)

## Real-Time Data Streaming

When using `range` with unbuffered channels:
- The goroutine continuously produces data
- `range` blocks waiting for each new value
- Data flows in real-time as it's produced
- The loop continues until the channel is closed or the goroutine stops

## Key Concepts

1. **Unbuffered channel**: Synchronous communication, no buffer
2. **Goroutine producer**: Continuously sends data in a loop
3. **Range consumer**: Receives data as it's produced
4. **Real-time processing**: Data flows immediately without buffering

## Your Task

Create a goroutine that:
1. Has an infinite loop that sends the current second to a channel
2. Waits one second between each send
3. Use `range` in main to receive and print each second value
4. Observe how data flows in real-time from the goroutine

## Expected Behavior

You should see the current second value printed continuously, updating every second. This demonstrates real-time data streaming with unbuffered channels.

```go
// Exercise: Channels - Range (unbuffered)

// In this exercise we will use the range keyword to iterate over a UNbuffered (sync) channel.
// Create a unbuffered channel (type int)
// Create a goroutine that:
// Has an infinite loop -> prints the current second and waits for a second
// Use the range iterator in the main function to see each second

package main

import "fmt"
import "time"

func second(c chan int){
	
}

func main () {
	var c chan int = make(chan int)
	go second(c)
	
	
}
```

<details>
<summary> Solution: </summary>

```go
// Exercise: Channels - Range (unbuffered)

// In this exercise we will use the range keyword to iterate over a UNbuffered (sync) channel.
// Create a unbuffered channel (type int)
// Create a goroutine that:
// Has an infinite loop -> prints the current second and waits for a second
// Use the range iterator in the main function to see each second

package main

import "fmt"
import "time"

func second(c chan int){
	for {
		c <- time.Now().Second()
		time.Sleep(time.Second)
	}
}

func main () {
	var c chan int = make(chan int)

	go second(c)
	
	for element := range c {
		fmt.Println(element)
	}
}
```

</details>
