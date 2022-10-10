// Exercise: SLICES

// Create a slice (substring) from a string that contains at least 4 words

package main

import "fmt"

func main () {
    var mystring = "I like how the rain sofly touches the window when I'm reading inside"
	// Your code goes here
	var substring = mystring[2:19]
	fmt.Println(substring)
}