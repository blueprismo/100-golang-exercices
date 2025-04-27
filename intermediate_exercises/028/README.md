# Exercise: Functions inception

A function is a block of statements that can be used repeatedly in a program.

Functions are very convenient and widely used in programming, making it easier to read, understand, and maintain. They also promote code reuse [DRY principle](https://en.wikipedia.org/wiki/Don%27t_repeat_yourself)

## The anatomy of a function

```go
func functionName(param1 type1, param2 type2) returnType { // This is called the function signature
    // Inside the braces we have the function body, where our function logic resides.
}
```

Example:

```go
func subtract(x,y int) int { // function signature
	return x - y // function body, will return to the function caller the operation of substracting x - y.
}

func main() {
	result := subtract(5,4) // the function is called and it's return value assigned to the `result` variable
	fmt.Println(result) // the result will be 1
}
```

Create a function that sums 2 numbers, then call that function from another function that sums another number.

```golang
package main

import "fmt"

func sum() int {
	
}

func secondsum() int {
	
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

func sum(x int, y int) int {
	return x + y
}

func secondsum(x, y, z int) int {
	return sum(x,y) + z 
}

func main () {
	// Your code goes here
	var result int
	result = secondsum(2,3,5)

	fmt.Println(result)
}

```

</details>
