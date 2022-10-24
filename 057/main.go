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
  conn, _ := 

  // Infinite loop again 
  for { 
    // Sending...
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Text to send: ")
    // in a variable called "text, use the ReadString function from the reader variable, with the newline character as it's argument"
    

    // send to server (use Fprintf for this!) 
    
    // wait for reply from server (must accept the connection)
    message, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print("Message from server: "+message)
  }
}