// Exercise: For loops
// Iterate through the array and add two (+2) to each value

package main

import "fmt"

func main () {
	
	var arr = [10]int{1,2,3,4,5,6,7,8,9,10}
	
	for index, value := range arr {
		fmt.Print(index , ") " , value, "\n")
	}
}