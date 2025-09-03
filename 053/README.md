# Exercise 53: Unstructured JSON with interface{}

## The Challenge of Unknown JSON Structure

Sometimes you receive JSON data where you don't know the structure beforehand. The fields, types, or nesting might vary, making it impossible to define a fixed struct.

## What is interface{}?

`interface{}` (empty interface) is Go's way of representing "any type." It can hold values of any type, making it perfect for unknown JSON data.

## Using map[string]interface{} for JSON

For unstructured JSON objects, use:
```go
var result map[string]interface{}
```

This creates a map where:
- **Keys** are strings (JSON field names)
- **Values** can be anything (interface{})

## How JSON Types Map to Go Types

When unmarshaling into `interface{}`, JSON types become:

| JSON Type | Go Type |
|-----------|---------|
| `string` | `string` |
| `number` | `float64` |
| `boolean` | `bool` |
| `array` | `[]interface{}` |
| `object` | `map[string]interface{}` |
| `null` | `nil` |

## Accessing Nested Data

To access nested objects, you need type assertions:

```go
// JSON: {"user": {"name": "Alice", "age": 30}}
var result map[string]interface{}
json.Unmarshal(jsonData, &result)

// Access nested object
user := result["user"].(map[string]interface{})
name := user["name"].(string)
age := user["age"].(float64)  // JSON numbers are float64
```

## Type Assertions

When getting values from `interface{}`, use type assertions:
```go
// Basic type assertion
name := result["name"].(string)

// Safe type assertion with ok check
if name, ok := result["name"].(string); ok {
    fmt.Println("Name:", name)
} else {
    fmt.Println("Name is not a string")
}
```

## Your Task

Working with the provided bird JSON:
1. **Create a map[string]interface{}** called "result"
2. **Unmarshal the JSON** into this map
3. **Extract the "birds" nested object**
4. **Print the values** from the birds object

## Key Concepts

- `map[string]interface{}` handles any JSON object structure
- Nested objects become nested maps
- Type assertions are needed to access specific types
- Always consider type safety when working with `interface{}`

## When to Use This Approach

- **Unknown JSON structure**: APIs that return different formats
- **Flexible data**: Configuration files with varying structures  
- **Exploration**: When you need to examine JSON before defining structs
- **Partial parsing**: When you only need specific fields from complex JSON