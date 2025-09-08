// Exercise: Channels - Range (unbuffered)

// In this exercise we will use the range keyword to iterate over a UNbuffered (sync) channel.
// Create a unbuffered channel (type int)
// Create a goroutine that:
// Has an infinite loop -> prints the current second and waits for a second
// Use the range iterator in the main function to see each second

package main

import "fmt"
import "time"

func second(c chan int){
	for {
		c <- time.Now().Second()
		time.Sleep(time.Second)
	}
}

func main () {
	var c chan int = make(chan int)

	go second(c)
	
	for element := range c {
		fmt.Println(element)
	}
}
