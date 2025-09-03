// Exercise: Functions

// Create a recursive function that returns the sequence of fibonacci up until the nth number


package main

import "fmt"

func fibonacci(x int) int{
	if (x <= 1) {
		return x
	}
	return fibonacci(x-1) + fibonacci(x-2)
}

func main () {
	// Your code goes here
	fmt.Println(fibonacci(9))
}