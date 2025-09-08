// Exercise: Let's talk about ServeMux!
package main

import "net/http"
import "log"

// In this exercise I want you to talk about serveMux, arm yourself with documentation first please!
// In a nutshell, servemux is an http request multiplexer (https://pkg.go.dev/net/http#ServeMux)
// In the last exercise, notice that we passed the 'nil' value to the .ListenAndServe() function. When a nil value is passed, the DefaultServeMux value is gathered.

// Let's supose we have the previous handler function:
func customHandler (w http.ResponseWriter, r *http.Request){
  w.Write([]byte("Hello from customHandler!"))
}

func main() {  
  // Create a ServeMux variable called 'mux'
  // And assign it the value of a new servemux
  mux := http.NewServeMux()

  //Use the HandleFunc() function, as the 1st argument you will have to put the root path "/" and the 2nd argument will be the customHandler function
  mux.HandleFunc("/", customHandler)

  // Start the server with ListenAndServe() function, and as the second parameter use the "mux" servemux you have created!
  server := http.ListenAndServe(":8080", mux)
  if (server != nil){
    log.Print("Cannot start sever")
  }
}