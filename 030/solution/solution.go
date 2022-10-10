// Exercise: Functions

// Create a variadic function that sums many numbers passed as arguments
// Variadic functions can be called with any number of arguments.


package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _,num := range numbers{
		total += num
	}
	fmt.Println(total)
	return total
}

func main () {
	// Your code goes here
	sum(2,3,4,5,6,7)
}