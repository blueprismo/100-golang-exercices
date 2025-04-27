# Exercise: RANGE and Blank identifier

- use range to print the values of the array
- use the blank identifier variable to supress the index, DON'T show the index, just the values of the array!

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
	
	for _, value := range arr {
		fmt.Print(value, "\n")
	}
}
```

</details>
