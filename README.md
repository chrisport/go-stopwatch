[![Build Status](https://drone.io/github.com/chrisport/go-stopwatch/status.png)](https://drone.io/github.com/chrisport/go-stopwatch/latest)

Outdated readme, will be updated within 2 days. See tests for usage examples.

stopwatch
==========

Simple stopwatch to measure elapsed time. Very useful for quick debugging and

### Usage:
Start new stopwatch, get elapsed time after doing stuff and restart stopwatch
```go
stopwatch := stopwatch.Stopwatch()

// do stuffg

elapsedNanos := stopwatch.GetAndRestart()
```

#### GetAndContinue
```go
stopwatch := stopwatch.NewStopwatch()

// do stuff for 5 nanoseconds

elapsedNanos := stopwatch.GetAndContinue() // = 5ns

// do more stuff for 11 nanoseconds

elapsedNanosTotally := stopwatch.GetAndContinue() // =  5 + 11 = 16ns
```

#### GetAndRestart
```go
stopwatch := stopwatch.NewStopwatch()

// do stuff for 5 nanoseconds

elapsedNanos := stopwatch.GetAndRestart() // = 5ns

// do more stuff for 11 nanoseconds

elapsedNanosTotally := stopwatch.GetAndRestart() // = 11ns
```

#### Print to console
```go
stopwatch := stopwatch.NewStopwatch()

// do stuff called StuffThatMustBeDone for 220 nanoseconds

stopwatch.LogAndContinue("Stuff that must be done")

Output:
    [Stopwatch] 220 ns for Stuff that must be done
```

#### Adjust accuracy
```go
stopwatch := stopwatch.NewStopwatchWithAccuracy(time.Milliseconds)

// do stuff called StuffThatMustBeDone for 500 ms

stopwatch.LogAndContinue("Stuff that must be done")

Output:
    [Stopwatch] 500 for Stuff that must be done
```

### GetPrecise
Get precise calculates the time directly on the int representing the nanoseconds. Benchmark on MacBook Pro shows a
time advantage of about 27% when restarting stopwatch and 47% otherwise:
```golang
BenchmarkGetAndRestart  50000000                47.3 ns/op
BenchmarkGetAndContinue 50000000                32.5 ns/op
BenchmarkGetPrecise     100000000               17.2 ns/op
BenchmarkGetPreciseAndRestart   50000000                34.3 ns/op
```
