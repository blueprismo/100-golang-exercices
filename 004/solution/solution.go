package main

import "fmt"

func main () {
	// Creating new variable called helloworld
	var helloworld, itsmemario string
	helloworld = "Hello World!"
	itsmemario = "It's a me, Maaario"
	// Print the variable
	fmt.Println(helloworld + " " + itsmemario)
}

// To run the program:
// - go run solution.go