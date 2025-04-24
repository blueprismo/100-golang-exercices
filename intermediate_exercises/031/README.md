# Exercise: Functions IV

Create a recursive function that returns the sequence of fibonacci up until the nth number

```golang
package main

import "fmt"

// Complete the function signature
func fibonacci() int {
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

func fibonacci(x int) int{
	if (x <= 1) {
		return x
	}
	return fibonacci(x-1) + fibonacci(x-2)
}

func main () {
	// Your code goes here
	fmt.Println(fibonacci(9))
}
```

</details>
