// Exercise: JSON!!

// There is a json with custom fields names.
// modify the Human struct to map the custom name, to a single-named variable in our GO code.

package main

import (
    "fmt"
    "encoding/json"
)

type Human struct {
  Name string `json:"family name"`
  Description string `json:"what does he say"`
  Dimensions Dimensions
}

type Dimensions struct {
  Height string
  Weight string
}

func main() {
  humanJson := `{ "family name": "Rick",
                  "what does he say": "has a grandson called Morty",
                  "dimensions": { 
                    "height": "1.80m",
                    "weight": "50kg"
                  }
                }`

  var human Human
  json.Unmarshal([]byte(humanJson),&human)
  fmt.Println(human.Name + " is old and " + human.Description)
}