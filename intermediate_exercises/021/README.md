# Exercise: MAP

Create a map of ints to strings
  1 should map to "A"
  2 should resolve "B"
  3 should resolve "C"

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
