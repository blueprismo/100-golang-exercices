# Exercise: Create an array of 5 strings, and initialize it's 2 first values with some random names

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
        fmt.Println(arr[i])
    }
}
```

</details>
