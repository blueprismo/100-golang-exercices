// Interfaces -  Empty interfaces

// The empty interface type essentially describes no methods. It has no rules.
// And because of that, it follows that any and every object satisfies the empty interface.
// Wherever you see an empty interface in a declaration (such as a variable, function parameter or struct field)
// This means you can use an object of any type.
package main

import (
	"fmt"
)

// Create an type called human, it will be a map of a string as a key, and an empty interface as the value


func main() {
	// make a new human, assign it to the person variable

	// set the person "name" to "Alice"

	// Set the persons age to any int

	// Set the person height to any float64

	// print the person
    fmt.Printf("%+v", person)
}