// Exercise: JSON!!

// Create a struct "Human", with 2 string attributes: Name and Description

// Find out about the unmarshal function (json decode), and decode the humanJson object into a variable of type Human called "human"
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
  humanJson := `{"name": "Rick","description": "has a grandson called Morty"}`

  var human Human
  json.Unmarshal([]byte(humanJson),&human)

  fmt.Println(human.Name + " is old and " + human.Description)
}
