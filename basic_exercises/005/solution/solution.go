// Exercise: Learn FORMAT STRING!! Print some verbs with their corresponding letter
// Format string / verbs are followed by a '%' character.
// Example: fmt.Printf("Hello, my name is %s", name) // will substitute %s for the content in the variable name called "name"

// Tip: You can head into the documentation for reference: https://pkg.go.dev/fmt#hdr-Printing

package main

import "fmt"

func main () {
	// Here goes your code
	var name string
	var age	int64
	var legal bool
	var weight float32

	name = "Anna"
	age  = 20
	legal = true
	weight = 70.12

	fmt.Printf("My name is %s, I am %d years old and it's %t that I can drive a car, my pet weights %f kilograms",name, age, legal, weight)
}
