# Exercise 83: Introduction to Testing

## What is Testing in Go?

Testing is a crucial part of software development that ensures your code works correctly. Go has built-in testing support through the `testing` package, making it easy to write and run tests.

## Go Testing Conventions

- Test files must end with `_test.go`
- Test functions must start with `Test` and accept `*testing.T` parameter
- Test functions signature: `func TestXxx(t *testing.T)`
- Use `t.Error()` or `t.Fail()` to mark tests as failed
- Run tests with `go test` command

## Basic Testing Structure

```go
func TestExample(t *testing.T) {
    expected := "expected result"
    actual := YourFunction()
    
    if actual != expected {
        t.Errorf("Expected %s, got %s", expected, actual)
    }
}
```

## Common Testing Methods

- **t.Error(args)**: Log error and mark test as failed
- **t.Errorf(format, args)**: Log formatted error and fail
- **t.Fatal(args)**: Log error, fail test, and stop execution
- **t.Fatalf(format, args)**: Log formatted error, fail, and stop

## Your Task

Look at the main.go file and complete the exercise by:
1. Creating a function that returns "Hello, Test"
2. Understanding how to structure test functions
3. Learning the basic testing workflow in Go

This exercise introduces you to Go's testing framework and establishes the foundation for writing reliable code.