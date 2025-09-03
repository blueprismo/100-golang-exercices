# Exercise 51: Nested JSON Objects

## Understanding Nested JSON Structures

JSON objects can contain other objects, creating nested structures. This is common in real-world data where entities have complex relationships.

## Nested JSON Example

```json
{
  "name": "John",
  "address": {
    "street": "123 Main St",
    "city": "Boston"
  }
}
```

## Mapping to Go Structs

To handle nested JSON, you create nested Go structs:

```go
type Address struct {
    Street string `json:"street"`
    City   string `json:"city"`
}

type Person struct {
    Name    string  `json:"name"`
    Address Address `json:"address"`
}
```

## How Unmarshaling Works with Nested Objects

When `json.Unmarshal()` encounters a nested object:
1. It looks for a field in your struct that matches the JSON field name
2. If that field is another struct, it recursively unmarshals into that struct
3. All nested fields must be **exported** (capitalized) to be accessible

## Accessing Nested Data

After unmarshaling, you access nested data using dot notation:

```go
var person Person
json.Unmarshal(jsonData, &person)

fmt.Println(person.Name)           // "John"
fmt.Println(person.Address.Street) // "123 Main St"
fmt.Println(person.Address.City)   // "Boston"
```

## Your Task

1. **Create the Dimensions struct** with Height and Weight string fields
2. **Add Dimensions field** to the Human struct 
3. **Unmarshal the nested JSON** into your Human struct
4. **Print only the values from the nested Dimensions object**

## Key Concepts

- Nested structs map to nested JSON objects
- Field names must match (or use struct tags)
- Access nested data with dot notation: `human.Dimensions.Height`

## Expected Behavior

You should be able to print the height and weight values from the nested dimensions object.