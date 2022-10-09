// Exercise: RANGE
// Range is used to iterate over data structures
// use range to print the values and index of the array


package main

import "fmt"

func main () {
	
	var arr = [10]int{1,2,3,4,5,6,7,8,9,10}
	
	for _, value := range arr {
		fmt.Print(value, "\n")
	}
}