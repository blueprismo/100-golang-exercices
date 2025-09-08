// Exercise: Channels synchronisation

// There are two goroutines now
// Call the first goroutine (named 'task') and run the second goroutine 'task2' ONLY after the first one has finished

package main

import "fmt"
import "time"

func task(done chan bool) {
    fmt.Print("running Task1 goroutine...")
    time.Sleep(time.Second)
    fmt.Println("done")
    done <- true
}

func task2(){
	fmt.Println("Task2 goroutine...")
}

func main () {
	var done chan bool = make(chan bool,1)	
	fmt.Println("I am running in the main thread concurrently")
	// Your code goes here
	
}