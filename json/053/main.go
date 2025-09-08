// Exercise: JSON!! - Unknown beforehand data

// Now the JSON will be unstructured (we won't know the values beforehand)
// When data is unstructured, we will create a map of strings to empty interfaces.
// Create a map of string to an empty interface called "result"
// Unmarshall (decode) the json into the result
// Get the "birds" object into a variable called "birds" and print it's values

package main

import (
    "fmt"
    "encoding/json"
)



func main() {
  birdJson := `{
                "birds": {
                  "pigeon":"likes to perch on rocks",
                  "eagle":"bird of prey"
                },
                "animals": "none"
              }`

  
}