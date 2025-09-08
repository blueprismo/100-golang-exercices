# Exercise 69: JSON Marshaling (Encoding)

## What is JSON Marshaling?

**Marshaling** is the process of converting Go data structures into JSON format. It's the opposite of unmarshaling - instead of JSON → Go, it's Go → JSON.

## json.Marshal() Function

```go
data, err := json.Marshal(goStruct)
if err != nil {
    // Handle error
}
jsonString := string(data) // Convert []byte to string
```

## Marshaling Examples

**Simple struct:**
```go
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

person := Person{Name: "Alice", Age: 30}
jsonData, _ := json.Marshal(person)
// Result: {"name":"Alice","age":30}
```

**Array of structs:**
```go
people := []Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
}
jsonData, _ := json.Marshal(people)
// Result: [{"name":"Alice","age":30},{"name":"Bob","age":25}]
```

## Key Points

- `json.Marshal()` returns `[]byte`, not string
- Use `string(data)` to convert to readable format
- Struct fields must be exported (capitalized) to be marshaled
- Use struct tags to control JSON field names

## Your Task

1. Create a `user` struct with Username and Password fields
2. Create an array of 3 users
3. Marshal the array to JSON using `json.Marshal()`
4. Print the resulting JSON string

This demonstrates how to convert Go data to JSON format for APIs, file storage, or data transmission.