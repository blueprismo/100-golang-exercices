# Exercise: MAP

A map is a key-value store where each key is associated with a single value.
We could imagine a dictionary, where the word "Bee" holds the index for it's definition "a stinging winged insect which collects nectar and pollen, produces wax and honey, and lives in large communities."

Create a map of ints to strings
  1 should map to "A"
  2 should resolve "B"
  3 should resolve "C"

Reference: https://gobyexample.com/maps

```golang
package main

import "fmt"

func main () {
	// Your code goes here
}

```

<details>
<summary> Solution: </summary>

```golang
package main

import "fmt"

func main () {
	// Your code goes here
	amap := make(map[int32]string)
	amap[1] = "A"
	amap[2] = "B"
	amap[3] = "C"

	fmt.Println(amap[2])
}
```

</details>
