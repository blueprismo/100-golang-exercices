// Exercise: Channels - Closing

// Create a string channel "c" (make it a buffered channel)
// Add 2 different strings directly into that channel.
// Close the channel with the close() statement and read a quote from the channel, Can you read it?

package main

import "fmt"


func main () {
	var c chan string = make(chan string,2)

	c <- "Hello"
	c <- "how are you?"

	fmt.Println( <- c)
	close(c)
	fmt.Println( <- c)
}
