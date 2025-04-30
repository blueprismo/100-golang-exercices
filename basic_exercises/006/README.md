# Exercise: Format Verbs

Go provides an easy way to format verbs into your output. This is called formatting verbs.
You tell the program how do you want the variable to be expressed, according to it's variable type.

The formatting verb acts like a placeholder for the variable that you pass into it, here are two examples

The verb `%v` prints the value in it's default format
`fmt.Printf("my string variable contains %v", dog_name)`
Will print "my string variable contains Scooby"

The verb `%T` prints the variable type
`fmt.Printf("my string variable is of type %T", dog_name)`
Will print "my string variable is of type string"

You can check all the format verbs in the documentation: https://pkg.go.dev/fmt#hdr-Printing

Exercise: Print some verbs with their corresponding letter

- Format verbs are followed by a '%' character.


```golang

package main

import "fmt"

func main () {
	var name string
	var age	int64
	var legal bool
	var weight float32

	name = "Anna"
	age  = 29
	legal = false
	weight = 70.12

	// Here goes your code
	fmt.Printf("My name is __, I am __ years old and it's __ that I can drive a car, my pet weights __ kilograms", , , , )
}
```

<details>
<summary> Solution: </summary>

```golang
package main

import "fmt"

func main () {
	// Here goes your code
	var name string
	var age	int64
	var legal bool
	var weight float32

	name = "Anna"
	age  = 20
	legal = true
	weight = 70.12

	fmt.Printf("My name is %s, I am %d years old and it's %t that I can drive a car, my pet weights %f kilograms",name, age, legal, weight)
}
```

</details>
