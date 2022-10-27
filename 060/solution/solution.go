// Exercise: Set up a simple HTTP Server
package main

import "net/http"
import "log"

// In this exercise we will register the handler function for a given pattern!
// We will use this HandleFunc BEWARE IT'S NOT THE SAME THAN HandlerFunc!!!!!
// We will register these patterns (uri) in the defaultServerMux, we will talk about it later

// Create a function named handler_1 that will write "Hello from Handlefunc #1"
func handler_1(w http.ResponseWriter, r *http.Request){
  w.Write([]byte("Hello from Handlefunc #1"))
}

//Create a function named handler_2 that will write "Hello from handlefunc #2"
func handler_2(w http.ResponseWriter, r *http.Request){
  w.Write([]byte("Hello from Handlefunc #2"))
}

func main() {  
  // This ListenAndServe will get 2 parameters: 1 - Address (:8080) in this case
  //                                            2 - nil
  
  // Now, use the http.HandleFunc() to register the handler_1 function to the "/handler1" pattern
  // And use the same method to register the handler_2 function to "/handler2" pattern
  http.HandleFunc("/handler1", handler_1)
  http.HandleFunc("/handler2", handler_2)

  server := http.ListenAndServe(":8080", nil)
  if (server != nil){
    log.Print("Cannot start sever")
  }
}