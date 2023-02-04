// Contexts - Chaining contexts
package main

import (
	"context" 
	"fmt"
)

// Create another function called doAnother, as a context as it's only argument. Like the last exercise, it will print the same key as the doSomething function
func doAnother(ctx context.Context) {
	// Here, print the ctx.Value() call with the key variable you added in the main function
	fmt.Println("doSomething: Name's value is", ctx.Value("Name"))
}

// Modify the doSomething function
func doSomething(ctx context.Context) {
	// Here, print the ctx.Value() call with the key variable you added in the main function
	fmt.Println("doSomething: Name's value is", ctx.Value("Name"))
	// Create a varibale anotherCtx and use the context.WithValue(a,b,c) function to pass in the father context as A (ctx), the same key name as b, and a different value as c
	anotherCtx := context.WithValue(ctx,"Name","Mary")
	// call the doAnother function with the anotherCtx as it's only argument.
	// See it's behaviour: Have the values chainged? Are context mutable?
	doAnother(anotherCtx)
}

func main() {
	ctx := context.Background()

	// here we will assign to ctx the context.WithValue() method, the first argument will be the context, the second your key
	// (this key will be referenced as the "ctx.Value(key)" in the function doSomething above, and the third the value of your key (for example key = Name, value = John)
	ctx = context.WithValue(ctx,"Name","John")
	// call the function with the empty context created before as it's only argument
	doSomething(ctx)
}