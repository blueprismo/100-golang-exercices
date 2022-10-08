package main

import "fmt"

func main () {
	// Creating new variable called helloworld
	var helloworld string
	helloworld = "Hello World!"
	// Print the first letter
	fmt.Println(helloworld[0])
}

// To run the program:
// - go run solution.go