// Exercise: Routines

// Document yourself and ask what a goroutine is.
// Make a go routine of the counter() function
// right after calling the go routine, in the next line. call the same routine with another different number


package main

import "fmt"
import "time"

func counter(x int) {
	for i := 0; i < x; i++ {
		fmt.Println(i)
	}
}

func main () {
	// Your code goes here
	

	// this sleep is in order to not exit the program sooner than the routine lifetime :)
	time.Sleep(5 * time.Second)
}