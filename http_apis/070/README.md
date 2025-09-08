# Exercise 70: TCP Server with net Package

## Introduction to Network Programming

Go's `net` package provides excellent support for network programming. You can create TCP/UDP servers and clients with minimal code.

## TCP Server Basics

A TCP server:
1. **Listens** on a specific port
2. **Accepts** incoming connections  
3. **Handles** client communications
4. **Reads/Writes** data to/from clients

## Key Functions

**net.Listen()** - Start listening:
```go
listener, err := net.Listen("tcp", ":8080")
```

**listener.Accept()** - Accept connections:
```go
conn, err := listener.Accept()
```

**Reading from connection:**
```go
reader := bufio.NewReader(conn)
message, err := reader.ReadString('\n')
```

## Basic Server Structure

```go
// Listen on port
listener, err := net.Listen("tcp", ":8000")

// Accept connection  
conn, err := listener.Accept()

// Read messages in loop
for {
    reader := bufio.NewReader(conn)
    message, err := reader.ReadString('\n')
    fmt.Println("Received:", message)
}
```

## Testing Your Server

Once your server runs:
- **netcat**: `nc localhost 8000`
- **telnet**: `telnet localhost 8000`  
- **curl**: For HTTP-like testing

## Your Task

Create a basic TCP server that:
1. Listens on port 8000
2. Accepts a connection
3. Reads messages from the client in a loop
4. Prints received messages

This is foundational for building web servers, chat servers, and other networked applications.