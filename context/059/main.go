// Contexts - Canceling a context
package main

import (
	"context" 
	"fmt"
	"time"
)

// Modify the doSomething function
func doSomething(ctx context.Context) {
	// use the context.WithCancel() function with the ctx as the only argument. Assign it to two values, first the ctx(this will be the parent function) and then the cancelCtx which will be a new empty function.

	// Make a new unbuffered channel of integers and assign it to printCh

	// call the doAnother function as goroutine

	// create a loop of 3 integer iterations from 1 to 3 which will be sent to the printCh channel!
	for num := 1; num <= 3; num++ {
		printCh <- num
	}
	// call the cancelCtx() function

	// sleep for 100 ms
	time.Sleep(100 * time.Millisecond)
	// print that doSomething has finished
	fmt.Printf("doSomething: finished\n")
}


// Modify the doAnother function,
// First add a second argument that will receive a int value from a channel into a variable called printCh
func doAnother(ctx context.Context, printCh <-chan int) {
	// For and select in conjuntion is a great way to operate with channels
	// The select statement lets a goroutine wait on multiple communication operations.
	// A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready. 
	for {
		select {
		// first case will be reciving a ctx.Done() call, if this happens we will handle errors and abort the doAnother function.
		case :
			
		// Second case will receive a value from printCh and assign that to a variable called num, after that print the num variable
		case :
			fmt.Printf()
		}
	}
}

func main() {
	ctx := context.Background()
	doSomething(ctx)
}