// Exercise: Time

// import the "time" package
// create a variable named "current" that has the current time in UTC
// Print the Date and time with the format [YYYY Mo dd], then with [YYYY/Mo/dd] and finally [HH:MM:SS]

package main

import (
    "fmt"
    "time"
)

func main() {
  current := time.Now().UTC()
  fmt.Println("Date: " + current.Format("2006 Jan 02")) //(*)
  fmt.Println("Date: " + current.Format("2006/Jan/02")) //(*)
  fmt.Println("Time: " + current.Format("03:04:05")) // (*)
}

// (*)
    // The layout string used by the Parse function and Format method
	// shows by example how the reference time should be represented.
	// We stress that one must show how the reference time is formatted,
	// not a time of the user's choosing. Thus each layout string is a
	// representation of the time stamp,
	//	Jan 2 15:04:05 2006 MST
	// An easy way to remember this value is that it holds, when presented
	// in this order, the values (lined up with the elements above):
	//	  1 2  3  4  5    6  -7