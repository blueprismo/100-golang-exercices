# Exercise: STRUCT

Create a Hotel structure with:

- numRooms int32
- streetName string
- hasPool bool

Then assign a value to each of those attributes

References:
https://go.dev/tour/moretypes/2
https://www.w3schools.com/go/go_struct.php

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
