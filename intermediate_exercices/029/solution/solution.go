// Exercise: Functions

// Create a function that returns 2 integer values
// There will be 2 arguments (int)
// The first returned value will be the sum of the arguments, and the second the substraction of them

package main

import "fmt"

func operation(x, y int) (int, int) {
	var sum, substraction int
	sum = x + y
	substraction = x - y 
	return sum, substraction
}

func main () {
	// Your code goes here
	sum, subs := operation(10,5)
	fmt.Println(sum, subs)
}