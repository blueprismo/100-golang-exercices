# Exercise 43: Range Over Unbuffered Channels

## Range with Unbuffered Channels

While the previous exercise used buffered channels with `range`, this exercise explores using `range` with unbuffered channels and goroutines producing data in real-time.

## The Key Difference

- **Buffered channels**: Fill buffer, close, then range
- **Unbuffered channels**: Goroutine produces data while main goroutine ranges

## Pattern for Unbuffered Channel Range

```go
func producer(ch chan int) {
    for i := 1; i <= 5; i++ {
        ch <- i
    }
    close(ch)  // Important: close when done producing
}

func main() {
    ch := make(chan int)  // Unbuffered
    go producer(ch)       // Producer runs concurrently
    
    for value := range ch {  // Range receives as producer sends
        fmt.Println(value)
    }
}
```

## Real-Time Data Processing

This pattern is excellent for:
- **Streaming data**: Process data as it arrives
- **Pipeline architectures**: One stage produces, another consumes
- **Real-time systems**: Handle data with minimal buffering

## Why It Works

1. **Range blocks**: Waits for next value when channel is empty
2. **Producer sends**: Goroutine sends values one by one
3. **Synchronization**: Range and send operations coordinate automatically
4. **Clean exit**: When producer closes channel, range loop exits

## Your Task

Create a system where:
1. A goroutine produces data continuously and sends it to an unbuffered channel
2. The main function uses `range` to receive and process this data
3. The producer runs in an infinite loop (never closes the channel)
4. You'll see real-time data flow between goroutines

## Expected Behavior

Since this is an infinite loop, the program will run forever, printing data as it's produced. This demonstrates real-time communication between goroutines.