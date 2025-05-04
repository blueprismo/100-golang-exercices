# Exercise: while loop

While in other programs there is a "while" loop, for instance python, in go there is no such a keyword!
The idiomatic equivalent in golang is a simplified for loop where you put the condition.

Bear in mind, that this can create infinite loops if the logic inside the loop is not properly handled. That is, the condition that is being checked must change towards exiting the loop.

A simple example of this loop that prints the numbers from 0 to 100

```go
init := 0
end := 100
for (init <= end) {
	fmt.Println(init)
	init++
	}
```

Exercise: With a simple "for" loop, print the numbers from 30 to 50

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
