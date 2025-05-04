# Exercise: Range

Ranges are an idiomatic way for iterating over data structures, it is widely used to iterate over arrays, maps or slices.

The syntax for range is:

```go
for index, value := range <data_structure> {
  // index and value are populated automatically for each element of the data structure.
}
```

- Use range to print the values and index of the array

```golang
package main

import "fmt"

func main () {
  // initialized array of 10 int values [1..10]
  var arr = [10]int{1,2,3,4,5,6,7,8,9,10}
  // Here goes your code

}
```

<details>
<summary> Solution: </summary>

```golang
package main

import "fmt"

func main () {

  var arr = [10]int{1,2,3,4,5,6,7,8,9,10}

  for index, value := range arr {
    fmt.Print(index , ") " , value, "\n")
  }
}
```

</details>
