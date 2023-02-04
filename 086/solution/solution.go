// Contexts
// I feel contexts are a really powerful feature in GO and I haven't given the importance that it deserves.

// In this episodes, we are going to talk about contexts!
package main

import (
	"context" // import the "context" package
	"fmt"
)

// Create a function called doSomething with an argument ctx of type context.Context (this is an interface)
func doSomething(ctx context.Context) {
	fmt.Println("Doing something!")
}

func main() {
	// create a variable called ctx, this variable will have a context.TODO() value, which returns an empty Context.
	ctx := context.TODO()

	// call the function with the empty context created before as it's only argument
	doSomething(ctx)
}