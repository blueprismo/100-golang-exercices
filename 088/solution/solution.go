// Contexts - Add data 
package main

import (
	"context" 
	"fmt"
)

// Create a function called doSomething with an argument ctx of type context.Context (this is an interface)
func doSomething(ctx context.Context) {
	// Here, print the ctx.Value() call with the key variable you added in the main function
	fmt.Println("doSomething: Name's value is", ctx.Value("Name"))
}

func main() {
	ctx := context.Background()

	// here we will assign to ctx the context.WithValue() method, the first argument will be the context, the second your key
	// (this key will be referenced as the "ctx.Value(key)" in the function doSomething above, and the third the value of your key (for example key = Name, value = John)
	ctx = context.WithValue(ctx,"Name","John")
	// call the function with the empty context created before as it's only argument
	doSomething(ctx)
}