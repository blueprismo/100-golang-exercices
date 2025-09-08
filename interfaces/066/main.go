// Interfaces -  Empty interfaces, bear in mind...
package main

import (
	"fmt"
)

type human map[string]interface{}

// But there's an important thing to point out when it comes to retrieving and using a value from this map
// let's say that we want to get the "age" value and increment it by 1.
func main() {
	person := make(human)
	person["name"] = "Alice"
    person["age"] = 21
    person["height"] = 167.64

	// Try to update the person age, the new value will be the existing one plus 1
	// (When you run the program, see what happens), what's the error it arises?
	person["age"] = 

    fmt.Printf("%+v", person)
}