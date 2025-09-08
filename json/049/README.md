# Exercise 49: JSON Unmarshaling Basics

## Understanding JSON and Go

JSON (JavaScript Object Notation) is a lightweight, text-based data interchange format. Go has excellent built-in support for JSON through the `encoding/json` package.

## What is Unmarshaling?

**Unmarshaling** is the process of converting JSON data into Go data structures. It's like "decoding" JSON text into Go variables you can work with.

```
JSON string → Go struct/variable
```

## The json.Unmarshal() Function

The main function for unmarshaling is:

```go
err := json.Unmarshal(jsonData, &variable)
```

Where:
- `jsonData` is `[]byte` containing JSON
- `&variable` is a pointer to the Go variable to store the result
- Returns an `error` if unmarshaling fails

## JSON to Go Type Mapping

JSON maps to Go types as follows:

| JSON Type | Go Type |
|-----------|---------|
| `string` | `string` |
| `number` | `float64`, `int`, etc. |
| `boolean` | `bool` |
| `array` | `[]T` (slice) |
| `object` | `struct` or `map` |
| `null` | `nil` |

## Struct Field Matching

By default, JSON field names must match Go struct field names (case-insensitive):

```go
type Person struct {
    Name string  // matches "name" in JSON
    Age  int     // matches "age" in JSON  
}
```

## Your Task

1. **Define the Human struct** with `Name` and `Description` string fields
2. **Declare a Human variable** to store the unmarshaled data
3. **Use json.Unmarshal()** to convert the JSON string into your Go struct
4. **Handle any errors** that might occur during unmarshaling

## Important Notes

- `json.Unmarshal()` expects `[]byte`, so you may need to convert strings
- Always check for errors when unmarshaling
- The struct fields must be **exported** (capitalized) to be accessible by the JSON package

## Expected Behavior

The program should successfully unmarshal the JSON and print: "Rick is old and has a grandson called Morty"