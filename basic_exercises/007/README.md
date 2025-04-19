# Exercise: User input

- Using only the fmt package, ask a user for it's name and then for it's surname
- Store it in 2 variables called "name" and "surname"
- After user has entered the data, print it out

```golang
// Tip: https://pkg.go.dev/fmt#hdr-Scanning
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
	// Here goes your code
	var name, surname string
	fmt.Println("Please enter your name")
	fmt.Scanln(&name)
	fmt.Println("Please enter your surname")
	fmt.Scanln(&surname)

	fmt.Printf("Your name is: " + name + " " + surname)
}

```

</details>
