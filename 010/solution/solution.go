// Exercise: Arrays
// Create an array of 10 "int8" values, in it's initialization, fill those values from 0 to 9

package main

import "fmt"

func main () {
	var arr = [5]string{"thomas","phillip"}

	// Print the array
	for i := 0; i < len(arr); i++ {
        fmt.Println(arr[i])
    }
}