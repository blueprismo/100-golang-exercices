# Exercise: Arrays

## What is an array

An array fixed-size collection of elements of the same type.

Arrays are always bound to a single type (`int`,`float64`,`bool`, `string`, etc.) meaning that all the elements inside an array are of the same type.

Arrays are FIXED in size, they cannot grow bigger than they are.

## Working with arrays

The syntax to declare an array is `[<size>]<type>`

- If we want to declare an array of 5 integers, we will do `var arrayOfInts = [5]int
- If we want to declare an array of 4 booleans, we will do `var arrayOfBools = [4]bool
- You know how it goes!
  
We can create arrays of almost everything, numbers, booleans, even arrays of arrays (matrices)!
In this exercise we are going to learn how to initialize an array, that is, to populate it's initial values.

## Zero-values

In golang, when we declare a variable and do not make any explicit assignment to it (we don't store any value in it) by default it takes the zero or null value of it's type.

The zero value is:

- 0 for numeric types,
- false for the boolean type, and
- "" (the empty string) for strings.

Say we declare a 4 integers array `var fourIntArray [4]int`, as the array type is of `int` and we have not assigned any value, the array will contain a `0` in every element of the array.

![array_example](array.png)

There are multiple ways of populating an array, but if we know which values are going to be in the array we can initialize those values in the variable declaration by adding the values between braces `{}`.

Example:

```go
var newArray = [4]int{3,2,1,0}
```

This will initialize an array of 4 integers and assign them the values between the curly braces.
You can access the elements of an array thanks to their index value.

```go
// newArray anatomy:
Value: 3 - 2 - 1 - 0
Index: 0 - 1 - 2 - 3
```

For example, to access the second position in the array, you would do so with `newArray[1]`, giving you the value of `2`. But we will practice this in the next exercise. For now...

Exercise: Create an array of 10 "int8" values, and fill those values from 0 to 9

```go
package main

import "fmt"

func main () {
  // Here goes your code
  fmt.Printf("...")
}
```

<details>
<summary> Solution: </summary>

```go
package main

import "fmt"

func main () {
  // common initialization when the values are deterministic
  var init_arr = [10]int{0,1,2,3,4,5,6,7,8,9}

  /*
  Simple but unefficient solution
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

  // this would also work!
  //var init_arr = [10]int{1:9}
}
```

</details>
