# Exercise 71: TCP Client with net Package

## Creating TCP Clients

A TCP client connects to a server and can send/receive data. This complements the server from the previous exercise.

## Key Client Functions

**net.Dial()** - Connect to server:
```go
conn, err := net.Dial("tcp", "localhost:8000")
```

**Writing to connection:**
```go
fmt.Fprintf(conn, "Hello server\n")
```

**Reading from stdin:**
```go
reader := bufio.NewReader(os.Stdin)
text, err := reader.ReadString('\n')
```

## Basic Client Structure

```go
// Connect to server
conn, err := net.Dial("tcp", "localhost:8000")

for {
    // Get user input
    reader := bufio.NewReader(os.Stdin) 
    fmt.Print("Enter message: ")
    text, _ := reader.ReadString('\n')
    
    // Send to server
    fmt.Fprintf(conn, text)
    
    // Read server response
    message, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print("Server:", message)
}
```

## Testing Client-Server

1. **Start server** (exercise 70) in one terminal
2. **Start client** (this exercise) in another terminal
3. **Type messages** in client - they appear on server
4. **Bidirectional communication** possible

## Your Task

Create a TCP client that:
1. Connects to localhost:8000 (your server from previous exercise)
2. Prompts user for text input
3. Sends input to server
4. Waits for server response
5. Repeats in a loop

This demonstrates the client side of network communication and completes your client-server pair.