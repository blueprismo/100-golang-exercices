# Exercise: MAP

A map is a key-value store where each key is associated with a single value.
We could imagine a dictionary, where the word `bee` holds the index for it's definition "a stinging winged insect which collects nectar and pollen, produces wax and honey, and lives in large communities."

Like in the dictionary we only have one entry for `bee`, maps do not allow duplicates.

A map is a very important data structure that we often see in software engineering, as usually provide very nice performance when accessing or storing data, even for huge data sets.

## Syntax for creating maps

In golang, maps are created with

```go
var a = map[KeyType]ValueType{key1:value1, key2:value2,...}

// examples:
var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}
```

Exercise: Create a map of ints to strings

`1` should map to `A`
`2` should map to `B`
`3` should map to `C`

Reference: https://gobyexample.com/maps

```go
package main

import "fmt"

func main () {
  // Your code goes here
}
```

<details>
<summary> Solution: </summary>

```go
package main

import "fmt"

func main () {
  // Your code goes here

  mymap := map[int]string{1: "A", 2: "B", 3: "C"}
  // Clumsier code (but works)
  // amap := make(map[int32]string)
  // amap[1] = "A"
  // amap[2] = "B"
  // amap[3] = "C"

  fmt.Println(mymap[2])
}
```

</details>
