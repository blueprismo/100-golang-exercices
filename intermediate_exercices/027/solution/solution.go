// Exercise: SLICES

// Create a slice (substring) from a string that contains at least 4 words

package main

import "fmt"

func main() {
	var slice = []int32{0, 1, 2, 3, 4}
	// Your code goes here
	// remember the slice index begins at number 0!
	new_slice := slice[1:3]
	fmt.Printf("substring:%v, type: %T", new_slice, new_slice)
}