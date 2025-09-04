# Exercise: Structs

## What are structs?

In golang, there is no 'class' keyword to define an object, in golang the concept to create a collection of members of different data types into a single variable are called structs (structures).

A struct is a typed collection of fields.
It can hold any type, even other structs.
An example of a struct:

```go
type Person struct {
  name string
  age int
  job string
  salary int
}
```

You can see that the Person object contains 4 members: `name`,`age`,`job` and `salary`.
The members can be from different data types.

## Accessing struct members

To access any member of a structure, use the dot operator (.) between the structure variable name and the structure member:

```go
func main() {
  var pers1 Person

  // Pers1 specification
  pers1.name = "Hege"
  pers1.age = 45
  pers1.job = "Teacher"
  pers1.salary = 1000

  // Access and print Pers1 info
  fmt.Println("Name: ", pers1.name) // Output: 'Name: Hege'
  fmt.Println("Age: ", pers1.age)
  fmt.Println("Job: ", pers1.job)
  fmt.Println("Salary: ", pers1.salary)
}
```

## Anonymous structs

A struct that doesn't have a name is called an anonymous struct, you can see lots of them in [table-driven tests](https://go.dev/wiki/TableDrivenTests).

Anonymous structs are defined without a name and therefore cannot be referenced elsewhere in the code.
To create an anonymous struct, just instantiate the instance immediately using a second pair of brackets after declaring the type:

```go
newCar := struct {
    make    string
    model   string
    mileage int
}{
    make:    "Ford",
    model:   "Taurus",
    mileage: 200000,
}
```

Exercise: Create a structure named `Hotel` with three members `numRooms`,`streetName` and `hasPool`

- numRooms int32
- streetName string
- hasPool bool

Then assign a value to each of those attributes

References:  
[https://go.dev/tour/moretypes/2](https://go.dev/tour/moretypes/2)  
[https://www.w3schools.com/go/go_struct.php](https://www.w3schools.com/go/go_struct.php)  
[https://gobyexample.com/structs](https://gobyexample.com/structs)  

```go
package main

import "fmt"

type Hotel struct {
  // Your code goes here

}

func main () {
  var myHotel Hotel
  // Your code goes here
}
```

<details>
<summary> Solution: </summary>

```go
package main

import "fmt"

type Hotel struct {
  // Your code goes here
  numRooms int32
  streetName string
  hasPool bool
}

func main () {
  var myHotel Hotel
  // Your code goes here
  myHotel.numRooms = 30
  myHotel.streetName = "Thaerstrasse"
  myHotel.hasPool    = true
  fmt.Printf("My hotel in %v has %d rooms and it's %t that has a Pool", myHotel.streetName, myHotel.numRooms, myHotel.hasPool)
}
```

</details>
