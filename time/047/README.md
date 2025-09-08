# Exercise 47: Time Arithmetic

## Adding Time to Dates

Time arithmetic is common in programming - you often need to calculate times in the future or past relative to a given time.

## The Add() Method

Go's `time.Time` type has an `Add()` method that takes a `time.Duration`:

```go
future := now.Add(30 * time.Minute)     // 30 minutes from now
past := now.Add(-2 * time.Hour)        // 2 hours ago
```

## Duration Types

Go provides convenient duration constants:

```go
time.Nanosecond   // 1 nanosecond
time.Microsecond  // 1000 nanoseconds  
time.Millisecond  // 1000 microseconds
time.Second       // 1000 milliseconds
time.Minute       // 60 seconds
time.Hour         // 60 minutes
```

## Creating Durations

You can create durations in several ways:

```go
// Using constants
duration1 := 5 * time.Minute
duration2 := 2 * time.Hour + 30 * time.Minute

// Using ParseDuration
duration3, _ := time.ParseDuration("1h30m")
duration4, _ := time.ParseDuration("45s")
```

## Common Duration Examples

```go
5 * time.Second                    // 5 seconds
30 * time.Minute                   // 30 minutes  
2 * time.Hour                      // 2 hours
24 * time.Hour                     // 1 day
7 * 24 * time.Hour                 // 1 week
```

## Your Task

Add 20 minutes to the current UTC time. You need to:
1. Use the `Add()` method on the `current` time variable
2. Create a duration representing 20 minutes
3. Store the result in the `inTenMinutes` variable (note: the variable name says "ten" but you should add 20 minutes as per the exercise comment)

## Expected Behavior

The program should display a time that is exactly 20 minutes after the current time.

## Hint

Remember that `time.Minute` is a duration constant, so `20 * time.Minute` creates a 20-minute duration.