package main

// Testing!! 
// In this file we are going to create the beforementioned test function! (read the main.go file first!!)

import "fmt"
// Import the "testing" package here
import "testing"


// Create the test function called TestHello with the signature (t *testing.T)
func TestHello (t *testing.T) {
	// Inside this function, create 2 variables:
	// "have" with the value as the function HelloTest()
	have := HelloTest()
	// "want" with the hardcoded string "Hello, Test"
	want := "Hello, Test"

	// conditional: if have != want, raise t.Error() with a message
	if (have != want){
		t.Error("Test failed, strings not matching")
	}
	// else print Test passed!
	fmt.Println("Test passed!")
}
