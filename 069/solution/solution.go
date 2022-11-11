package main

import "fmt"

// Testing!! 
// Now let's see what happens when a test fails!
// Exercise: Create a function that returns a string "Hello, Test. THIS WILL FAIL"
func HelloTest() string {
  return "Hello, Test. THIS WILL FAIL"
}

// Print the function
func main() {
  fmt.Println(HelloTest())
}