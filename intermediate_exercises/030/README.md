# Exercise: Functions III

Create a variadic function that sums all it's numbers passed as arguments.
Variadic functions can be called with any number of arguments.


```golang
package main

import "fmt"

// Complete the function signature
func sum() int {
	total := 0
    // Your code goes here
}

func main () {
	// Your code goes here
	
}

```

<details>
<summary> Solution: </summary>

```golang
package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _,num := range numbers{
		total += num
	}
	fmt.Println(total)
	return total
}

func main () {
	// Your code goes here
	sum(2,3,4,5,6,7)
}

```

</details>
