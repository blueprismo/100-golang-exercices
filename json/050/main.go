// Exercise: JSON!!

// Create a struct "Human", with 2 string attributes: Name and Description

// Keep the Human struct, now in the json object there are going to be 2 humans
// create a "humans" array of type Human and print it's values.

package main

import (
    "fmt"
    "encoding/json"
)

type Human struct {
  Name string
  Description string
}

func main() {
  humansJson := `[{"name": "Rick",
                 "description": "has a grandson called Morty"},
                {"name": "Cactus",
                 "description": "is one of the powerpuff girls' name"}]`

  

  
}
