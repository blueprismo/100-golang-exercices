# Exercise: SLICES II

Slices are arrays that are dynamically-sized, meaning they have no fixed size!
Like arrays, they store multiple values of the same type in a single variable.

```go
// ARRAY
// Example of an array containing 5 integer numbers.
var array = [5]int32{0,1,2,3,4}
array = append(array, 3) // Compile-time error: first argument to append must be a slice

// SLICE
// Example of a slice containing 5 integer numbers
var slice = []int32{0,1,2,3,4}
slice = append(slice, 3) // ✅ OK
```

In the snippet above we can see that the difference in the declaration between array and slice is the absence of the size.

Exercise: From the slice `myslice`, create a new slice containing only the middle 3 elements.
Hint: use slice[start:end] syntax

```go
package main

import "fmt"

func main () {
  var slice = []int32{0,1,2,3,4}
  // Your code goes here

}
```

<details>
<summary> Solution: </summary>

```go
package main

import "fmt"

func main() {
  var slice = []int32{0, 1, 2, 3, 4}
  // Your code goes here
  // remember the slice index begins at number 0!
  new_slice := slice[1:3]
  fmt.Printf("substring:%v, type: %T", new_slice, new_slice)
}
```

</details>
