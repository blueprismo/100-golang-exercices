// Exercise: Channels - Range

// In this exercise we will use the range keyword to iterate over a buffered (async) channel.
// Create a buffered channel (type int) of a dimension of 5

// TIP: but buffered channels need to be closed before iterating over them!!! 

package main

import "fmt"


func main () {
	var c chan int = make(chan int,5)

	c <- 3
	c <- 6
	c <- 8
	c <- 22
	c <- 1
	close(c)
	
	for element := range c {
		fmt.Println(element)
	}
}
