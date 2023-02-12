// Interfaces -  Empty interfaces, SWITCH
package main

import "fmt"

// Create a function called do, the only argument it will have will be an empty interface
// Inside that function, create a switch case and use the type assertion to get the interface's type!
// Our goal here is to evaluate the type and proceed according to it.

func do(i interface{}) {
	// Switch statement, use type assertion to get the type (type keyword!) (("i.(type)""))
	switch v := i.(type) {
		// for integers, print the integer value
		case int:
			fmt.Printf("Twice %v is %v\n", v, v*2)
		// for strings, print the length of the string
		case string:
			fmt.Printf("%q is %v bytes long\n", v, len(v))
		// As default, print something about the type that reaches this zone!
		default:
			fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	// call the do function with 3 different types :) 
	do(21)
	do("hello")
	do(true)
}