# Exercise 45: Tickers

## Understanding Tickers

A **ticker** is a mechanism that repeatedly sends events at regular intervals. Think of it like a metronome - it "ticks" at a steady pace.

## time.Tick() Function

Go provides `time.Tick(duration)` which returns a channel that delivers "tick" events at regular intervals:

```go
ticker := time.Tick(time.Second)  // Ticks every second
```

## Using Tickers with Range

You can use `range` to receive from a ticker channel:

```go
for tick := range time.Tick(time.Second) {
    fmt.Println("Tick at:", tick)
}
```

## Important Notes about time.Tick()

- **Returns a channel**: The channel receives time values at each tick
- **Cannot be stopped**: `time.Tick()` runs forever and cannot be stopped
- **Memory considerations**: For long-running programs, consider `time.NewTicker()` instead, which can be stopped

## Alternative: time.NewTicker()

For production code, `time.NewTicker()` is often better because it can be stopped:

```go
ticker := time.NewTicker(time.Second)
defer ticker.Stop()  // Important: stop the ticker to free resources

for tick := range ticker.C {
    fmt.Println("Tick at:", tick)
}
```

## Common Use Cases

- **Periodic tasks**: Running maintenance operations
- **Status updates**: Refreshing displays or sending heartbeats
- **Timeouts**: Implementing regular checks
- **Animation**: Updating visual elements at regular intervals

## Your Task

Create a goroutine that uses `time.Tick()` to print "Tick" every second. The main function should let it run for 5 seconds before the program exits.

## Expected Behavior

You should see "Tick" printed approximately every second for 5 seconds total.