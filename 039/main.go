// Exercise: Multiple different channels

// We will make use of the 'select' statement
// Create three different channels "c1,c2,c3" for a goroutine called "name"
// The go routine will have 2 arguments, first the channel (type string) and then a name (type string)
// Inside the goroutine (function) will be:
// 1- An infinite loop:
// 1.1- The second argument "name" will be ingested into the channel
// 1.2- Wait for 2 seconds

// In the main program
// call 3 times the 'name' goroutine with a different name each
// make an infinite loop containing the select statement
// Inside that select, put 3 'case' keywords. In each case print the name for each different channel.

package main

import "fmt"
import "time"

func name(c chan string, name string){
	
}

func main () {
	var c1 chan string = make(chan string)
	var c2 chan string = make(chan string)
	var c3 chan string = make(chan string)

	for {
		select {
		
		}
	}
	fmt.Println("Goroutines finished.") // You shouldn't see this message as the goroutines run forever!
}