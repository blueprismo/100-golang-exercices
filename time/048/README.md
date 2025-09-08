# Exercise 48: Time Comparison and Operations

## Comparing Times

Go provides several methods to compare and work with time values. This is essential for determining relationships between different times.

## Time Comparison Methods

The `time.Time` type provides these comparison methods:

### 1. Sub() - Calculate Difference
```go
duration := end.Sub(start)  // Returns time.Duration
fmt.Println(duration)       // e.g., "8760h0m0s" (1 year)
```

### 2. Equal() - Test Equality  
```go
same := time1.Equal(time2)  // Returns bool
```

**Note**: Use `Equal()` instead of `==` because times might have different locations but represent the same instant.

### 3. After() - Test if Later
```go
isLater := time1.After(time2)  // Returns bool
```

### 4. Before() - Test if Earlier
```go
isEarlier := time1.Before(time2)  // Returns bool
```

## Understanding Duration

When you subtract times with `Sub()`, you get a `time.Duration`:

```go
diff := end.Sub(start)
fmt.Println(diff.Hours())    // Duration in hours
fmt.Println(diff.Minutes())  // Duration in minutes  
fmt.Println(diff.Seconds())  // Duration in seconds
```

## Practical Examples

```go
now := time.Now()
birthday := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)

age := now.Sub(birthday)
fmt.Printf("Age: %.0f days\n", age.Hours()/24)

if now.After(birthday) {
    fmt.Println("Birthday has passed")
}
```

## Your Task

Using the two provided time variables (`start` and `end`), you need to:

1. Calculate the difference between them using `Sub()`
2. Test if they are equal using `Equal()`  
3. Determine which comes after the other using `After()`

Display the results to understand the relationship between these two times.

## Expected Behavior

You should see:
- The duration between the two times (should be about 1 year)
- Whether they are equal (should be false)
- Which time comes after the other

## The Time Values

- `start`: February 1, 2020 at 03:00 UTC
- `end`: February 1, 2021 at 12:00 UTC

The `end` time is about 1 year and 9 hours after the `start` time.