// Exercise: Set up a simple HTTP Server

// We will use the net/http library
// https://pkg.go.dev/net/http

// There are multiple ways of creating a http server in GO
// In this task, we are going to crete a http in your preferred way, so run free with the documentation :)
// I will provide the simplest solution, the only requirement for this exercise is that the server must 
// reply with a "Hello web World" response.

package main

import "net/http"
import "io"
import "log"

func main() {
   // Set routing rules
   http.HandleFunc("/", WebFunc)

   //Use the default DefaultServeMux.
   err := http.ListenAndServe(":8080", nil)
   if err != nil {
           log.Fatal(err)
   }
}

func WebFunc(w http.ResponseWriter, r *http.Request) {
   io.WriteString(w, "Hello web World!")
}
  
