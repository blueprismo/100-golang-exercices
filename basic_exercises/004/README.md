# Exercise: Two string variables

You can operate with strings. For instance we can build a quote by adding 2 string variables together:
`fmt.Println("My pet is called " + dog_name + ", and he loves" +  dog_favourite_food)`

There are multiple [conventions](https://en.wikipedia.org/wiki/Naming_convention_(programming)#Go) for variable naming, like `thisVariable` (camelCase), `this_variable` (snake_case), `this-variable` (kebab-case)...

In golang, camelCase and UpperCamelCase are preferred; making the first letter uppercase exports that piece of code, while lowercase makes it only usable within the current scope.

Exercise: Concatenate two string variables and print it's result

- One variable should be called "helloWorld" and the other "itsMeMario"
- There should be a space between them

```golang
package main

func main() {
  // Here goes your code
  var helloWorld, itsMeMario string
  // ...
}
```

<details>
<summary> Solution: </summary>

```golang
package main

import "fmt"

func main() {
  // Creating new variable called helloWorld
  var helloWorld, itsMeMario string
  helloWorld = "Hello World!"
  itsMeMario = "It's a me, Mario"
  // Print the variable
  fmt.Println(helloWorld + " " + itsMeMario)
}

// To run the program:
// - go run solution.go
```

</details>
