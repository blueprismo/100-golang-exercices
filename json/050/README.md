# Exercise 50: JSON Arrays and Slices

## Unmarshaling JSON Arrays

JSON arrays map to Go slices. When you have a JSON array of objects, you unmarshal it into a slice of structs.

## JSON Array Format

JSON arrays are enclosed in square brackets `[]` and contain comma-separated elements:

```json
[
  {"name": "Alice", "age": 30},
  {"name": "Bob", "age": 25}
]
```

## Go Slice for JSON Arrays

To unmarshal a JSON array, use a Go slice:

```go
type Person struct {
    Name string
    Age  int
}

var people []Person  // Slice of Person structs
```

## Unmarshaling Array Example

```go
jsonData := `[{"name":"Alice","age":30},{"name":"Bob","age":25}]`
var people []Person

err := json.Unmarshal([]byte(jsonData), &people)
if err != nil {
    fmt.Println("Error:", err)
    return
}

// Now people[0] is Alice, people[1] is Bob
for _, person := range people {
    fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
}
```

## Key Concepts

1. **Array Structure**: JSON arrays become Go slices
2. **Element Type**: Array elements must match the slice element type  
3. **Indexing**: Access elements using `slice[index]`
4. **Iteration**: Use `range` to iterate over all elements

## Your Task

The exercise provides:
- A `Human` struct (already defined)
- A JSON string containing an array of 2 humans
- You need to create a slice variable and unmarshal the JSON array into it

## Steps to Complete

1. **Create a slice variable** of type `[]Human` called `humans`
2. **Unmarshal the JSON array** into the slice
3. **Print the values** from each human in the slice

## Expected Behavior

You should be able to access and print information about both Rick and Cactus from the unmarshaled slice.