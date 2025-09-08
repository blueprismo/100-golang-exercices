# Exercise 52: JSON Struct Tags

## What are Struct Tags?

Struct tags are metadata attached to struct fields that provide instructions to packages like `encoding/json` on how to handle the fields during marshaling and unmarshaling.

## The Problem: Mismatched Field Names

Sometimes JSON field names don't match Go conventions:

```json
{
  "family name": "Smith",
  "phone number": "555-1234"
}
```

Go struct field names must be valid identifiers (no spaces), but JSON can have any string as a field name.

## Solution: JSON Struct Tags

Struct tags map JSON field names to Go struct fields:

```go
type Person struct {
    FamilyName  string `json:"family name"`
    PhoneNumber string `json:"phone number"`
}
```

## Struct Tag Syntax

The basic syntax is:
```go
FieldName DataType `json:"json_field_name"`
```

## Common JSON Struct Tag Options

- **Custom name**: `json:"custom_name"`
- **Ignore field**: `json:"-"`
- **Omit if empty**: `json:",omitempty"`
- **Omit with custom name**: `json:"custom_name,omitempty"`

## Examples

```go
type User struct {
    ID       int    `json:"id"`
    FullName string `json:"full_name"`
    Email    string `json:"email"`
    Password string `json:"-"`              // Never include in JSON
    Bio      string `json:"bio,omitempty"`  // Omit if empty
}
```

## Why Use Struct Tags?

1. **API Compatibility**: Match existing JSON APIs
2. **Naming Conventions**: JSON uses snake_case, Go uses PascalCase
3. **Field Control**: Hide sensitive fields, omit empty values
4. **Legacy Data**: Work with existing JSON formats you can't change

## Your Task

The JSON has field names with spaces: `"family name"` and `"what does he say"`. You need to:

1. **Add struct tags** to map these JSON field names to your Go struct fields
2. **Ensure the unmarshaling works** with the custom field names
3. **Verify the output** shows the correct values

## Key Learning

Struct tags bridge the gap between JSON naming conventions and Go naming conventions, allowing you to work with any JSON format regardless of its field naming style.