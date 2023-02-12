// Contexts - Chaining contexts, mutable or immutable?
package main

import (
	"context" 
	"fmt"
)

func doAnother(ctx context.Context) {
	fmt.Println("doSomething: Name's value is", ctx.Value("Name"))
}

// Modify the doSomething function
func doSomething(ctx context.Context) {
	fmt.Println("doSomething: Name's value is", ctx.Value("Name"))
	anotherCtx := context.WithValue(ctx,"Name","Mary")
	doAnother(anotherCtx)
	// Add another print statement to show the ctx.Value for your key again, does it change?
	
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx,"Name","John")
	doSomething(ctx)
}