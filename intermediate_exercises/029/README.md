# Exercise: Functions II

Create a function that returns 2 integer values
The function will take 2 integers as input (int)
The first returned value will be the sum of the arguments, and the second the substraction of them


```golang
package main

import "fmt"

func operation() () {
	
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

func operation(x, y int) (int, int) {
	var sum, substraction int
	sum = x + y
	substraction = x - y 
	return sum, substraction
}

func main () {
	// Your code goes here
	sum, subs := operation(10,5)
	fmt.Println(sum, subs)
}

```

</details>
