# Exercise: Arrays

- Create an array of 10 "int8" values, in it's initialization, fill those values from 0 to 9

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
	var arr = new([10]int)
	/*
	Simple but handcrafted solution
	arr[0] = 0
	arr[1] = 1
	arr[2] = 2
	arr[3] = 3
	arr[4] = 4
	arr[5] = 5
	arr[6] = 6
	arr[7] = 7
	arr[8] = 8
	arr[9] = 9
	*/

	// Another solution :) 
	for i := 0; i < len(arr); i++ {
		arr[i] = i
        fmt.Println(arr[i])
    }
}
```

</details>
