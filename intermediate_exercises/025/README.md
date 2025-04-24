# Exercise: Pointers II

- Declare a pointer variable of type int32
- Declare the address of the var "x"
- Save the address of the var "x" in the pointer variable

```golang
package main

import "fmt"

func main () {
    var x int32 = 5
	// Your code goes here

}
```

<details>
<summary> Solution: </summary>

```golang
package main

import "fmt"

func main () {
    var x int32 = 5
	// Your code goes here
	var pointerX *int32 = &x

	fmt.Println("Value of x: %d", x)
	fmt.Println("Memory address of x: %d", &x)
	fmt.Println("Pointer value: %d", pointerX)
}
```

</details>
