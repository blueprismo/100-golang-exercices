# Exercise: SLICES

Slices are arrays that are dynamically-sized, meaning they have no fixed size!
Like arrays, they store multiple values of the same type in a single variable.

```go
// ARRAY
// Example of an array containing 5 integer numbers.
var array = [5]int32{0,1,2,3,4}
array = append(array, 3) // Compile-time ERROR: first argument to append must be a slice

// SLICE
// Example of a slice containing 5 integer numbers
var slice = []int32{0,1,2,3,4}
slice = append(slice, 3) // ✅ OK
```

In the snippet above we can see that the difference in the declaration between array and slice is the absence of the size.

Exercise: Create a slice of 5 integers from an already existing slice called 'myset'
(This is called reslicing)

```go
package main

import "fmt"

func main () {
    var myset = []int32{0,1,2,3,4,5,6,7,8,9}
    // Your code goes here

}
```

<details>
<summary> Solution: </summary>

```go
package main

import "fmt"

func main () {
    var myset = []int32{0,1,2,3,4,5,6,7,8,9}
    // Your code goes here
    var slice = myset[0:5]
    fmt.Println("Slice :", slice)
}
```

</details>
