// Exercise: Socket implementation (client side)

// We will use the net library, and create a really basic client
// Please read the docs before anything!
// https://pkg.go.dev/net

package main

import "net"
import "fmt"
import "bufio"
import "os"

func main() {

  // connect to server (use the Dial function!)
  conn, _ := net.Dial("tcp", "127.0.0.1:8000")

  // Infinite loop again 
  for { 
    // what to send?
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Text to send: ")
    text, _ := reader.ReadString('\n')
    // send to server
    fmt.Fprintf(conn, text + "\n")
    // wait for reply
    message, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print("Message from server: "+message)
  }
}