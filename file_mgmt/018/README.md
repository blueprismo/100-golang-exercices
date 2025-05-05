# Exercise: Write to a file

In this exercise, we will create and write into a file, thanks to the [os.Create()](https://pkg.go.dev/os#Create) function.

```go
// os.Create function signature
func Create(name string) (*File, error)
```

It will return the pointer to a `File` type. Then we can use the [os.WriteString()](https://pkg.go.dev/os#File.WriteString) function, to write a string into the file we have just created above.

```go
// os.WriteString function signature
func (f *File) WriteString(s string) (n int, err error)
```

## Methods

A method is a function with a special receiver argument.

In this function signature above we see the `(f *File)` part between the `func` word and the function name. 

This means that that function `WriteString` can only be called in variables of type `*File`. The receiver of the function is `f` in the function signature above.

Another example:

```go
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
  // the receiver of the Abs() function will be a Vertex.
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
```

In this case, we can call `Abs()` into the `Vertex` "v".

Exercise: Write a list of 5 countries to a file
Tip: "\n" can be used as a newline in a string.

```golang
package main

import "os"


func main () {
    // Here goes your code

}
```

<details>
<summary> Solution: </summary>

```golang
package main

import "os"

func main () {
  // Here goes your code
  file, err := os.Create("countries.txt")

  if err != nil {
    return
  }

  defer file.Close()

  file.WriteString("Germany\nFrance\nUSA\nSpain\nUK\n")
}
```

</details>
