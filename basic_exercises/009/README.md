# Exercise: For loops

Loops are a crucial element in software engineering. They are usually used to fetch each element of a data structure, to populate multiple values while adding some logic into your code without repeating yourself.

The anatomy of a for loop is as follows, you need 3 parameters in this order:

1. The initial value (sometimes called iterator variable, because it iterates at every loop)
2. The condition in which the loop will exit. Infinite loops should be avoided.
3. The alteration in the iterator variable, always bringing it closer to exiting the loop.

The syntax to create such a loop would be like this:

```go
var iterator = 0

// for <iterator variable initial state>; <condition for loop>; <iterator alteration>
for iterator ; iterator < 3 ; iterator++ {
  fmt.Println(iterator)
  fmt.Println("I am inside the loop!") 
}
```

The above code snippet will print:
0
"I am inside the loop!"
1
"I am inside the loop!"
2
"I am inside the loop!"

Exercise: Iterate through all the array and add two (+2) to each value

- Hint: The iterator variable can be used as an index to fetch all the elements of an array!
- Tip: If you somehow create an infinite loop, hit `Control + C` to send a SIGINT signal to your program, forcing it to terminate the process, and therefore stop looping.

```go
package main

import "fmt"

func main () {
  // initialized array of 10 int values [0..9]
  var arr = [10]int{0,1,2,3,4,5,6,7,8,9}
  // Here goes your code

}
```

<details>
<summary> Solution: </summary>

```go
package main

import "fmt"

func main () {
  var arr = [10]int{0,1,2,3,4,5,6,7,8,9}
  for i := 0; i < len(arr); i++ {
    arr[i] = arr[i] + 2
  }
  fmt.Printf("%v", arr)
}
```

</details>
