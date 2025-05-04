# Exercise: Arrays II

Arrays are FIXED in size, meaning that they cannot grow bigger than it's original value. This also implies that operations with arrays should be done carefully as we may hit out of bounds errors. Dynamically sized arrays are called slices, which we will learn from them soon.

There is no need to initialize all the elements from an array, but remember that the remaining elements will contain it's zero value

Exercise: Create an array of 5 strings, and initialize it's 2 first values with some random strings.

```golang
package main

import "fmt"

func main () {
	// Here goes your code
	fmt.Printf("...")
}
```

<details>
<summary> Solution: </summary>

```golang
package main

import "fmt"

func main () {
	var arr = [5]string{"thomas","phillip"}

	// Print the array
	for i := 0; i < len(arr); i++ {
        fmt.Printf("Element number %d: %v \n", i, arr[i])
    }
}
```

</details>
