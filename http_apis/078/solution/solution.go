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
  animals := []Animal{
      Animal{"Alice","Cat"},
      Animal{"Bob","Cat"},
      Animal{"Trinity","Dog"},
  }
  // use the json NewEncoder() function, to encode the animals array within the responsewriter element
  json.NewEncoder(w).Encode(animals)
}

func main() {  
  // Register the handler into the defaultservemux (remember which function was used!) at the root path "/"
  http.HandleFunc("/", AnimalsHandler)
  // Start the server with the DefaultServeMux!
  server := http.ListenAndServe(":8080", nil)
  if (server != nil){
    log.Print("Cannot start sever")
  }
}