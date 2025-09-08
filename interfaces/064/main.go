// Interfaces -  Implicit implementation

// Extend the functionality of our interface

package main

import (
	"fmt"
)

type geometry interface{
	area() float64
	// add a new signature: perim() of type float64
	
}

// A rectangle struct
type rect struct {
    width, height float64
}

func (r rect) area() float64 {
    return r.width * r.height
}

// Create the func for perim(). The return value will be a rectangle again.
// To calculate a rectangle's perimeter, you should remember that it's 2*rwidth + 2*rheight


func main() {
	r := rect{width: 3, height: 4}
	fmt.Println(r.area())
	// show the perimeter:
	
}