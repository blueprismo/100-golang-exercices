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

	ctx, cancelCtx := context.WithCancel(ctx)
	// Make a new unbuffered channel of integers and assign it to printCh
	printCh := make(chan int)
	// call the doAnother function as goroutine
	go doAnother(ctx, printCh)
	// create a loop of 3 integer iterations from 1 to 3 which will be sent to the printCh channel!
	for num := 1; num <= 3; num++ {
		printCh <- num
	}
	// call the cancelCtx() function
	cancelCtx() // cancel when we are done counting numbers
	// sleep for 100 ms
	time.Sleep(1000 * time.Millisecond)
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
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Printf("doAnother err: %s\n", err)
			}
			fmt.Printf("doAnother: finished\n")
			return
		// Second case will receive a value from printCh and assign that to a variable called num, after that print the num variable
		case num := <-printCh:
			fmt.Printf("doAnother: %d\n", num)
		}
	}
}

func main() {
	ctx := context.Background()
	doSomething(ctx)
}