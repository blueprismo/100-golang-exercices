// Interfaces -  Empty interfaces, bear in mind...
package main

import (
	"fmt"
	"log"
)

type human map[string]interface{}

// But there's an important thing to point out when it comes to retrieving and using a value from this map
// let's say that we want to get the "age" value and increment it by 1.
func main() {
	person := make(human)
	person["name"] = "Alice"
    person["age"] = 21
    person["height"] = 167.64

	// Type assertion for the integer. use .(int) for type assertion!
	
	// Check that the assertion was alright
	

	person["age"] = age + 1

    fmt.Printf("%+v", person)
}