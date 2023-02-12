// Contexts - Background
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
	// create a variable called ctx, this variable will have a context.Background(), this is used for developers
	// to let know that it's initializing a context
	ctx := context.Background()

	// call the function with the empty context created before as it's only argument
	doSomething(ctx)
}