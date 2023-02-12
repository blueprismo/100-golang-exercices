package main

import "fmt"

// Testing!! 
// We will use the go test command and the testing package
// With go, to unit test your code you must create a file suffixed with "_test.go" where your tests will be
// For the tests to work, inside the test file the functions will be named TestXXX with signature func (t *testing.T). 
// if the function calls a failure function such as t.Error or t.Fail, the test is considered to have failed.  

// Exercise: Create a function that returns a string "Hello, Test"
func HelloTest() string {
  
}

// Print the function
func main() {
  fmt.Println(HelloTest())
}