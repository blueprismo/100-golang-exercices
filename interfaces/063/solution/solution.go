// Interfaces - Intro - Implicit implementation

// In golang, we don't use the "implements" keyword when using an interface.
// We don't have to say that explicitly, instead we need some type to satisfy the area() method.
// We say something implements an interface if it has a method with the exact signature.
// Inside our geometry interface, we have a signature called "area()"
// In this exercise, we are going to create a function with that signature!

package main

import (
	"fmt"
)

type geometry interface{
	area() float64
}

// A rectangle struct
type rect struct {
    width, height float64
}

// We are going to create a function with the area() signature.
// The function will be a receiver function (the receiver will be a rectangle "r")
// The function will be named area, without any arguments and will return a float64 value
func (r rect) area() float64 {
	// The return value will be the area of the rectangle
    return r.width * r.height
}

func main() {
	// Create a new rectangle and store it in a variable
	r := rect{width: 3, height: 4}
	fmt.Println(r.area())
}