# Exercise: while loop

- There is no "while" keyword in GOlang!
- With a for loop, print the numbers from 30 to 50

```golang
package main

import "fmt"

func main () {
	init := 30
	end  := 50
	// Here goes your code
	
}
```

<details>
<summary> Solution: </summary>

```golang
package main

import "fmt"

func main () {
	init := 30
	end := 50
	
	for (init <= end) {
		fmt.Println(init)
		init++
	}
}
```

</details>
