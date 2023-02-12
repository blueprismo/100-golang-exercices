// Exercise: Set up a simple HTTP Server

// To understand the http library better, we will do another server and handle it's router, one step at a time!

package main

import "net/http"
import "log"

// First of all, we are going to create an HTTP handler
  // The http.Handler is an interface, according to the docs like this>
  /*
    type Handler interface {
      ServeHTTP(ResponseWriter, *Request)
    }
    // This means that an object to be a handler MUST HAVE a ServeHTTP method with that signature [ServeHTTP(ResponseWriter, *Request)]
  */

// To put things simple, we will create handlers as functions, and them we will wrap around the ServeHTTP Method signature
// Create here a function called 'myCustomHandler' That has the same arguments as the ServeHTTP method:
func myCustomHandler(w http.ResponseWriter, r *http.Request){
  w.Write([]byte("Hello Web!"))
}

// Great! But how do we wrap around the serveHTTP method now?
// That will be easy, we will use the HandlerFunc: https://cs.opensource.google/go/go/+/refs/tags/go1.19.2:src/net/http/server.go;l=2105
// This wrapper is all we need in order to convert the function in an HTTP handler!! Easy isn't it? 

func main() {  
  // Create a server with the .ListenAndServe() function (more info here: https://pkg.go.dev/net/http#Server.ListenAndServe)
  // Also here is an excellent resource for http dissection: https://echorand.me/posts/golang-dissecting-listen-and-serve/#what-is-defaultservemux

  // This ListenAndServe will get 2 parameters: 1 - the uri (root in this case "/" and)
  //                                            2 - an http.Handler(). Or a custom function wrapped by the http.HandlerFunc(function)
  
  // This http.ListenAndServe should be called with port 8080 and use the function you created before as the Handler!
  server := http.ListenAndServe(":8080", http.HandlerFunc(myCustomHandler))
  if (server != nil){
    log.Print("Cannot start sever")
  }
}