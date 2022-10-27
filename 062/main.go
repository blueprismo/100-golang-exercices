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
// Bear in mind that this first part of the function
//   (ctr *Counter) is called a receiver, this means that the function caller WILL be a Counter element
func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
  ctr.n++
  fmt.Fprintf(w, "counter = %d\n", ctr.n)
}

func main() {  
  // Create a ServeMux variable called 'mux'
  // And assign it the value of a new servemux
  mux := http.NewServeMux()

  // Register a new Counter element in a variable called 'ctr' (tip: use the new() function!)

  // Now we will use the Handle() function with the mux variable.
  // The first argument will be the "/counter" pattern
  // The second argument will be the ctr function, in this case we can see very clearly how the handler acts as middleware. 
  

  // Start the server with ListenAndServe() function, and as the second parameter use the "mux" servemux you have created!
  server := http.ListenAndServe(":8080", mux)
  if (server != nil){
    log.Print("Cannot start sever")
  }
}