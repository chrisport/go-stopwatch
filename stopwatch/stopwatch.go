package stopwatch

import (
	"fmt"
	"time"
)

// Stopwatch is a simple timer that prints elapsed time to console or returns it as nanoseconds
type Stopwatch struct {
	lastCheckpoint     time.Time
	lastCheckpointPrec int64
	accuracyFactor     time.Duration
}

// NewStopwatchAndStart returns a new timer and starts it immediately.
func NewStopwatch() *Stopwatch {
	t := Stopwatch{}
	t.accuracyFactor = time.Nanosecond
	t.Restart()
	return &t
}

// NewStopwatchAndStart returns a new timer and starts it immediately.
func NewStopwatchWithAccuracy(accuracy time.Duration) *Stopwatch {
	t := Stopwatch{}
	t.SetAccuracy(accuracy)
	t.Restart()
	return &t
}

// Start starts the timer, if it has already been started, the timer will restart
func (t *Stopwatch) Restart() {
	t.lastCheckpoint = timeNow()
}

// Sets the accuracy of the Log output
func (t *Stopwatch) SetAccuracy(accuracy time.Duration) {
	t.accuracyFactor = accuracy
}

// GetAndRestart calculates the elapsed time since last start and returns it as duration. The timer will restart.
func (t *Stopwatch) GetAndRestart() time.Duration {
	elapsed := timeNow().Sub(t.lastCheckpoint)
	t.Restart()
	return elapsed
}

// GetAndContinue calculates the elapsed time since last start and returns it as duration. The timer will continue.
func (t *Stopwatch) Get() time.Duration {
	elapsed := timeNow().Sub(t.lastCheckpoint)
	return elapsed
}

// GetPrecise returns the elapsed time in nanoseconds.
// Compared to GetAndContinue it saves calculation time (in range of 10-30 nanoseconds)
func (t *Stopwatch) GetPrecise() int {
	elapsed := timeNow().Nanosecond() - t.lastCheckpoint.Nanosecond()
	return elapsed
}

// GetPreciseAndRestart returns the elapsed time in nanoseconds.
// Compared to GetandRestart it saves calculation time (in range of 10-30 nanoseconds)
func (t *Stopwatch) GetPreciseAndRestart() int {
	elapsed := timeNow().Nanosecond() - t.lastCheckpoint.Nanosecond()
	t.lastCheckpoint = timeNow()
	return elapsed
}

// LogAndRestart logs the elapsed time since last start and restarts the timer
func (t *Stopwatch) LogAndRestart(markerMessage string) {
	fmt.Printf("\n[Stopwatch] %v for %s", t.GetAndRestart(), markerMessage)
}

// LogAndContinue logs the elapsed time since last start. The timer will continue.
func (t *Stopwatch) Log(markerMessage string) {
	fmt.Printf("\n[Stopwatch] %v for %s", t.GetAndContinue(), markerMessage)
}

var timeNow = func() time.Time {
	return time.Now()
}
