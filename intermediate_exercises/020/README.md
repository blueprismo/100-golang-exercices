# Exercise: STRUCT

In golang, there is no 'class' keyword to define an object for example, a parallel for classes in golang are structs.

A struct is a typed collection of fields.
It can hold any type, even other structs.
You can associate methods for structs
A struct that doesn't have a name is called an anonymous struct, you can see lots of them in [table-driven tests](https://go.dev/wiki/TableDrivenTests).

Create a Hotel structure with:

- numRooms int32
- streetName string
- hasPool bool

Then assign a value to each of those attributes

References / more info:
https://go.dev/tour/moretypes/2
https://www.w3schools.com/go/go_struct.php
https://gobyexample.com/structs

```golang
package main

import "fmt"

type Hotel struct {
	// Your code goes here

}

func main () {
	var myHotel Hotel
	// Your code goes here
}
```

<details>
<summary> Solution: </summary>

```golang
package main

import "fmt"

type Hotel struct {
  // Your code goes here
  numRooms int32
  streetName string
  hasPool bool
}

func main () {
  var myHotel Hotel
  // Your code goes here
  myHotel.numRooms = 30
  myHotel.streetName = "Thaerstrasse"
  myHotel.hasPool    = true
  fmt.Printf("My hotel in %v has %d rooms and it's %t that has a Pool", myHotel.streetName, myHotel.numRooms, myHotel.hasPool)
}
```

</details>
