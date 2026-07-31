// Exercise: Routines

// Document yourself and ask what a goroutine is.
// Create a go routine named `start()` and call it inside our main() block
// Notice how goroutines do not wait for the execution to finish and follows with the next code statements

package main

import "fmt"

func main() {
	// call our go routine here

	// Immediately after calling our go routine, we will print "Started"
	// Note that possibly "Started" can be printed BEFORE our go routine string
	fmt.Println("Started")
	// add a sleep of 1 second here

	fmt.Println("Finished")
}

func start() {
	// Print something
}

