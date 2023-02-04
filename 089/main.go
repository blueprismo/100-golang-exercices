// Contexts - Chaining contexts
package main

import (
	"context" 
	"fmt"
)

// Create another function called doAnother, as a context as it's only argument. Like the last exercise, it will print the same key as the doSomething function


// Modify the doSomething function
func doSomething(ctx context.Context) {
	// Here, print the ctx.Value() call with the key variable you added in the main function
	fmt.Println("doSomething: Name's value is", ctx.Value("Name"))
	// Create a varibale anotherCtx and use the context.WithValue(a,b,c) function to pass in the father context as A (ctx), the same key name as b, and a different value as c

	// call the doAnother function with the anotherCtx as it's only argument.
	// See it's behaviour: Have the values chainged? Are context mutable?

}

func main() {
	ctx := context.Background()

	ctx = context.WithValue(ctx,"Name","John")

	doSomething(ctx)
}