// Exercise: Create an API
package main

import (
  "encoding/json"
  "net/http"
  "log"
)
// Let's create a struct 'Animal' with a variable 'Name' and another called 'Type' both type string
type Animal struct {
  Name  string `json:"Name"`
  Type  string `json:"Type"`
}

// Then, let's create a handler function ((a normal function, which then we will wrap :) ))
func AnimalsHandler(w http.ResponseWriter, r *http.Request) {
  // Create an array of 3 animals called "animals" and put 3 animals in there with it's name and type:
  
  // use the json NewEncoder() function, to encode the animals array within the responsewriter element
  
}

func main() {  
  // Register the handler into the defaultservemux (remember which function was used!) at the root path "/"
  
  // Start the server with the DefaultServeMux!
  server := http.ListenAndServe(":8080", )
  if (server != nil){
    log.Print("Cannot start sever")
  }
}