# Exercise: RANGE

- Range is used to iterate over data structures
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
