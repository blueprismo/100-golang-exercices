// Exercise: Multiple Goroutines with Channels


// Create a goroutine that:
// 1. Has an infinite loop that sends the current second to a channel
// 2. Waits one second between each send
// 3. Use `range` in main to receive and print each second value
// 4. Observe how data flows in real-time from the goroutine

package main


func second(c chan int){
	
}

func main () {
	var c chan int = make(chan int)
	go second(c)
	
	
}
