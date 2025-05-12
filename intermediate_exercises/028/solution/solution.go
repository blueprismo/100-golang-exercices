// Exercise: Functions

// Create a function that sums 2 numbers, then call that function from another function that sums another number.

package main

import "fmt"

func sum(x int, y int) int {
	return x + y
}

func secondsum(x, y, z int) int {
	return sum(x,y) + z 
}

func main () {
	// Your code goes here
	var result int
	result = secondsum(2,3,5)

	fmt.Println(result)
}