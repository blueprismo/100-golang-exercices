# Exercise 46: Time Basics and Formatting

## Working with Time in Go

Go's `time` package provides comprehensive functionality for working with dates and times. Understanding time formatting is essential for displaying times in human-readable formats.

## Getting Current Time

```go
now := time.Now()        // Current local time
utc := time.Now().UTC()  // Current time in UTC
```

## Go's Unique Time Formatting

Go uses a unique approach to time formatting. Instead of using format codes like `%Y %m %d`, Go uses a **reference time** to show the desired format.

## The Reference Time

Go's reference time is: **`Mon Jan 2 15:04:05 MST 2006`**

This represents:
- `Mon` - Monday (weekday)
- `Jan` - January (month name)  
- `2` - 2nd (day)
- `15:04:05` - 3:04:05 PM (time)
- `MST` - Mountain Standard Time (timezone)
- `2006` - Year

## Memory Aid: 1 2 3 4 5 6 -7

To remember the reference time, use the sequence:
- `1` = Month (January)
- `2` = Day
- `3` = Hour (15 = 3 PM)
- `4` = Minute  
- `5` = Second
- `6` = Year (2006)
- `-7` = Timezone offset

## Common Format Patterns

```go
time.Format("2006-01-02")           // 2023-12-07
time.Format("2006 Jan 02")          // 2023 Dec 07  
time.Format("15:04:05")             // 15:30:45
time.Format("3:04 PM")              // 3:30 PM
time.Format("January 2, 2006")      // December 7, 2023
```

## Your Task

Complete the exercise by:
1. Importing the `time` package
2. Creating a `current` variable with the current UTC time
3. Filling in the missing Format() calls to display the time in the requested formats

## Expected Output

You should see the current date and time displayed in the specified formats.

## Key Concept

Remember: Show the reference time (`Jan 2 15:04:05 2006 MST`) in the format you want your actual time to appear.