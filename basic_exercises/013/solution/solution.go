// Exercise: Conditional
// Check if the range of a number is between 20 and 30
// If the number is below 20 print : too cold
// If the number is inbetween print: perfect
// If the number is above 30 print : so hot

// Use if and a else if!

package main

import "fmt"

func main () {
	var number int
	fmt.Scanln(&number)
	
	if (number < 20){
		fmt.Println("Too cold!")
	} else if (number > 20 && number < 30) {
		fmt.Println("perfect")
	} else {
		fmt.Println("so hot!")
	}
	
}