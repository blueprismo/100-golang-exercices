// Exercise: Socket implementation (server side)

// We will use the net library, and create a really basic server
// Please read the docs before anything!
// https://pkg.go.dev/net
// When finished, try to connect with netcat (or telnet, or a client) and see if it works
package main

import (
    "fmt"
    "net"
    "bufio"
)

func main() {
  fmt.Println("Start server")
  // Now make the server listen at the 8000 port (tcp protocol)
  
  // Accept the connection
  

  // Run a loop forever (unless interrupted by signal)
  for {
    // Recive a message with the bufio.NewReader(connection).ReadString function
    
    fmt.Print("Message Received:", string(message))
  }
  
}