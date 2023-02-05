// Interfaces -  Empty interfaces, SWITCH
package main

import "fmt"

// Create a function called do, the only argument it will have will be an empty interface
// Inside that function, create a switch case and use the type assertion to get the interface's type!
// Our goal here is to evaluate the type and proceed according to it.

func  {
	// Switch statement, use type assertion to get the type (type keyword!) (("i.(type)""))
	switch {
		// for integers, print the integer value
		case int:
			
		// for strings, print the length of the string
		case string:
			
		// As default, print something about the type that reaches this zone!
		default:
		
	}
}

func main() {
	// call the do function with 3 different types :) 
	do(21)
	do("hello")
	do(true)
}