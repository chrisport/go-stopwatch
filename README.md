[![Build Status](https://drone.io/github.com/chrisport/go-timer/status.png)](https://drone.io/github.com/chrisport/go-timer/latest)

Outdated readme, will be updated within 2 days. See tests for usage examples.

timer
==========

Simple timer to measure elapsed time. Very useful for quick debugging and

### Usage:
Start new timer, get elapsed time after doing stuff and restart timer
```go
timer := timer.Timer()

// do stuffg

elapsedNanos := timer.GetAndRestart()
```

#### GetAndContinue
```go
timer := timer.NewTimer()

// do stuff for 5 nanoseconds

elapsedNanos := timer.GetAndContinue() // = 5ns

// do more stuff for 11 nanoseconds

elapsedNanosTotally := timer.GetAndContinue() // =  5 + 11 = 16ns
```

#### GetAndRestart
```go
timer := timer.NewTimer()

// do stuff for 5 nanoseconds

elapsedNanos := timer.GetAndRestart() // = 5ns

// do more stuff for 11 nanoseconds

elapsedNanosTotally := timer.GetAndRestart() // = 11ns
```

#### Print to console
```go
timer := timer.NewTimer()

// do stuff called StuffThatMustBeDone for 220 nanoseconds

timer.LogAndContinue("Stuff that must be done")

Output:
    [Timer] 220 ns for Stuff that must be done
```

#### Adjust accuracy
```go
timer := timer.NewTimerWithAccuracy(time.Milliseconds)

// do stuff called StuffThatMustBeDone for 500 ms

timer.LogAndContinue("Stuff that must be done")

Output:
    [Timer] 500 for Stuff that must be done
```
