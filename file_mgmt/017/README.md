# Exercise: Read a file

To read a file, there is the [os.ReadFile()](https://pkg.go.dev/os#ReadFile) function.
The function signature is:

```go
func ReadFile(name string) ([]byte, error)
```

see that it returns an array of `byte`, how do we read that as humans? you may ask.
Here's were we are going to learn about type conversions.

## Type conversion

Type conversion is the process of converting one type into another. For example an `int32` to a `float64` or a `[]byte` to a `string` and *Vice Versa*.

Not to be confused with type casting. There is no such thing in golang
> In type casting, a data type is converted into another data type by a programmer using casting operator. Whereas in type conversion, a data type is converted into another data type by a compiler.

## Converting a type

To convert a type in golang, you add the type you want the variable to be converted to and then wrap the variable in parentheses

```go
var i int = 42
var f float64 = float64(i) // now var f will contain the value 42, but won't be an integer!
```

Remember golang is a statically typed language! meaning that if we tried to assign the variable `i` (of type `int`) directly to `f` (of type `float64`) we would have the `cannot use i (variable of type int) as float64 value in variable declaration` error!

Exercise: Read the `read.txt` file in this directory.

Tip: You can convert an array of bytes to a string with `string(myArrayOfBytesVariable)`

```go
package main

import (
  "log"
  "fmt"
)

func main() {
	data, err := os.ReadFile()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
}

```

<details>
<summary> Solution: </summary>

```go
package main

import (
	"log"
  "fmt"
)

func main() {
	data, err := os.ReadFile("/tmp/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data)) // or os.Stdout.Write(data)
}
```

</details>
