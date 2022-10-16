// Exercise: Final exercise: Concurrency

// Create 2 goroutines, one called number() and the other called letter(), without any arguments, and without any returns
// The number() function will print a number from 0 to 9, waiting for 200 milliseconds between each iteration
// The letter() function will print the 10th first alphabet words, waiting for 200 millisecs between each iteration
// In the main program, call the go routines and print a meessage, see how it behaves.

package main

import (
    "fmt"
    "time"
)

func number() {                                                                  
    for i := 1; i <= 10; i++ {                                                  
        fmt.Printf("%d ", i)                                                     
        time.Sleep(200 * time.Millisecond)                                      
    }                                                                           
}  

func letter() {                                                                    
     for i := 0; i < 10; i++ {
       fmt.Printf("%c ", ('a'+i) )
       time.Sleep(200 * time.Millisecond)
     }
}  

func main() {
    go number()
    go letter()
    time.Sleep( 1 * time.Second)
    fmt.Println("main")
}
