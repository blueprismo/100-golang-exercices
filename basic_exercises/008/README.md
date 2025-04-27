# Exercise: User input

- Get a number from the console and check if it's odd
- You can create another function or do it everything in the main func

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
	var number int32
	fmt.Println("Enter a number: ")
	fmt.Scanln(&number)

	iseven(number)

	/*
	Possible solution without an additional function
	if (number % 2 == 0) { 
		fmt.Println("It's even")
	} else { 
		fmt.Println("It's odd")
	}*/
}

func iseven(number int32) bool {
	if (number % 2 == 0) {
		fmt.Println("is even")
		return true
	}  else { 
		fmt.Println("is odd")
		return false
	}
}
```

</details>
