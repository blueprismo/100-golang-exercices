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

  var result map[string]interface{}
  json.Unmarshal([]byte(birdJson), &result)
  
  // The object stored in the "birds" key is also stored as 
  // a map[string]interface{} type, and its type is asserted from
  // the interface{} type
  birds := result["birds"].(map[string]interface{})
  
  for key, value := range birds {
    // Each value is an interface{} type, that is type asserted as a string
    fmt.Println(key, value.(string))
  }
}