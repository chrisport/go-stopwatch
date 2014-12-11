[![Build Status](https://drone.io/github.com/chrisport/go-timer/status.png)](https://drone.io/github.com/chrisport/go-timer/latest)

timer
==========

Simple timer to measure elapsed time. Very useful for quick debugging and

### Usage:
Start new timer, get elapsed time after doing stuff and restart timer
```go
timer := timer.NewAndStart()

// do stuff

elapsedNanos := timer.GetAndRestart()
```

#### GetAndContinue
```go
timer := timer.NewAndStart()

// do stuff for 5 nanoseconds

elapsedNanos := timer.GetAndContinue() // = 5ns

// do more stuff for 11 nanoseconds

elapsedNanosTotally := timer.GetAndContinue() // =  5 + 11 = 16ns
```

#### GetAndRestart
```go
timer := timer.NewAndStart()

// do stuff for 5 nanoseconds

elapsedNanos := timer.GetAndRestart() // = 5ns

// do more stuff for 11 nanoseconds

elapsedNanosTotally := timer.GetAndRestart() // = 11ns
```

#### Print to console
```go
timer := timer.NewAndStart()

// do stuff called StuffThatMustBeDone

timer.LogAndContinue("Stuff that must be done")

Output:
    [Timer] 220 ns for Stuff that must be done
```

### TODO
- adjustable precision (ns, ms, s, ...)
