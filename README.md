[![Coverage Status](https://coveralls.io/repos/chrisport/go-stopwatch/badge.png)](https://coveralls.io/r/chrisport/go-stopwatch)
[![Build Status](https://drone.io/github.com/chrisport/go-timer/status.png)](https://drone.io/github.com/chrisport/go-timer/latest)
[![Coverage Status](https://img.shields.io/coveralls/chrisport/go-stopwatch.svg)](https://coveralls.io/r/chrisport/go-stopwatch)

Stopwatch
==========

Simple stopwatch to measure elapsed time. Very useful for quick debugging.

#### Basic Usage:
Start new stopwatch, get elapsed time after doing stuff and restart stopwatch

```go
stopwatch := stopwatch.Stopwatch()

// do stuff

elapsed := stopwatch.Get()
```

## Get time elapsed
#### Get
Get returns the elapsed time as time.Duration.

```go
stopwatch := stopwatch.NewStopwatch()

// do stuff for 5 nanoseconds

elapsed := stopwatch.Get() // = 5ns

// do more stuff for 11 nanoseconds

elapsedInTotal := stopwatch.Get() // =  5 + 11 = 16ns
```

#### GetAndRestart
Get and restart returns the elapsed time as time.Duration and restarts the stopwatch.

```go
stopwatch := stopwatch.NewStopwatch()

// do stuff for 5 nanoseconds

elapsed := stopwatch.GetAndRestart() // = 5ns

// do more stuff for 11 nanoseconds

elapsed2 := stopwatch.GetAndRestart() // = 11ns
```

## Print to console
#### Log & LogAndRestart
Log and LogAndRestart do same as Get, but print the result to console, together with a marker message. Sample console output:

```go
[Stopwatch] 220 ns for Stuff that must be done
```

#### Adjust accuracy
Accuracy for console log can be adjusted.

```go
stopwatch := stopwatch.NewStopwatchWithAccuracy(time.Milliseconds)

// do stuff called StuffThatMustBeDone for 500 ms

stopwatch.Log("Stuff that must be done")

Output:
    [Stopwatch] 500 ms for Stuff that must be done
```

## GetPrecise - when nanoseconds matter
GetPrecise and GetPreciseAndRestart calculate the time directly on the int representing the nanoseconds. Benchmark on MacBook 2,6 GHz Intel Core i5 shows a
time advantage of about 27% when restarting stopwatch and 47% otherwise:

```go
BenchmarkGet                    50000000                32.5 ns/op
BenchmarkGetPrecise             100000000               17.2 ns/op

BenchmarkGetAndRestart          50000000                47.3 ns/op
BenchmarkGetPreciseAndRestart   50000000                34.3 ns/op

```
