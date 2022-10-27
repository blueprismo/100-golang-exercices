// Exercise: Let's talk about ServeMux! And receivers in functions.
package main

import "net/http"
import "log"
import "fmt"

// We will have this Counter struct
type Counter struct {
  n int
}

// Imagine we have this counter.
func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
  ctr.n++
  fmt.Fprintf(w, "counter = %d\n", ctr.n)
}

func main() {  
  // Create a ServeMux variable called 'mux'
  // And assign it the value of a new servemux
  mux := http.NewServeMux()

  // Register a new Counter element in a variable called 'ctr'
  ctr := new(Counter) // This will initialize the value to 0!

  // Now we will use the Handle() function
  // The first argument will be the "/counter" pattern
  // The second argument will be the ctr function, in this case we can see very clearly how the handler acts as middleware. 
  mux.Handle("/counter", ctr)

  // Start the server with ListenAndServe() function, and as the second parameter use the "mux" servemux you have created!
  server := http.ListenAndServe(":8080", mux)
  if (server != nil){
    log.Print("Cannot start sever")
  }
}