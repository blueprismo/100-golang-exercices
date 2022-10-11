// Exercise: Channels and timeouts

// Create a go routine that uses time.Sleep for 10 seconds, and then add the string "10 seconds passed" into the channel (of type string)
// in the main program, inside the select block, have 2 cases:
// 1- The message from the channel
// 2- A timeout with the "time.After(3 * time.Second)" statement. After timeout happens, print "Timeout!!!!"


package main

import "fmt"
import "time"

func timeout(c chan string){
	for {
		time.Sleep(10 * time.Second)
		c <- "10 seconds passed"
	}
}

func main () {
	var c1 chan string = make(chan string)

	go timeout(c1)
	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case <- time.After(3 * time.Second):
			fmt.Println("Timeout!!!!")
		}
	}
	fmt.Println("Goroutines finished.") // You shouldn't see this message as the goroutines run forever!
}
