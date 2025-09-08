// Exercise: SLICES

// Create a slice of the first 5 numbers from a list of 10 numbers

package main

import "fmt"

func main () {
    var myset = []int32{0,1,2,3,4,5,6,7,8,9}
	// Your code goes here
	var slice = myset[0:5]
	fmt.Println("Slice :", slice)
}