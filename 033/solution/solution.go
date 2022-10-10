// Exercise: Channels

// It's important to get this concept, document yourself first! :) 
// There are two concurrent routines (f1 and f2)
// Send a message "Hello from f1" from f1 function

package main

import "fmt"
import "time"

func f1 (c chan string){
	c <- "Hello from f1"
}

func f2 (c chan string){
	msg := <- c
	fmt.Println("I am f2 and...", msg)
}
func main () {
	// Your code goes here
	var c chan string = make(chan string)
	go f1(c)
	go f2(c)

	// this sleep is in order to not exit the program sooner than the routine lifetime :)
	time.Sleep(1 * time.Second)
}