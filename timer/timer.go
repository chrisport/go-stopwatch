package timer

import (
	"fmt"
	"time"
)

const (
	Micros = 1000
	Millis = 1000000
)

// Timer is a simple timer that prints elapsed time to console or returns it as nanoseconds
type Timer struct {
	lastCheckpoint time.Time
	accuracyFactor int64
}

// NewTimerAndStart returns a new timer and starts it immediately.
func NewTimer() *Timer {
	t := Timer{}
	t.accuracyFactor = 1
	t.Restart()
	return &t
}

// NewTimerAndStart returns a new timer and starts it immediately.
func NewTimerWithAccuracy(accuracy time.Duration) *Timer {
	t := Timer{}
	t.SetAccuracy(accuracy)
	t.Restart()
	return &t
}

// Start starts the timer, if it has already been started, the timer will restart
func (t *Timer) Restart() {
	t.lastCheckpoint = timeNow()
}

// Start starts the timer, if it has already been started, the timer will restart
func (t *Timer) SetAccuracy(accuracy time.Duration) {
	t.accuracyFactor = int64(accuracy)
}

// LogAndRestart logs the elapsed time since last start and restarts the timer
func (t *Timer) GetAndRestart() int64 {
	elapsed := timeNow().Sub(t.lastCheckpoint)
	t.lastCheckpoint = timeNow()
	return elapsed.Nanoseconds() / t.accuracyFactor
}

//LogAndContinue logs the elapsed time since last start. The timer will continue.
func (t *Timer) GetAndContinue() int64 {
	elapsed := timeNow().Sub(t.lastCheckpoint)
	return elapsed.Nanoseconds() / t.accuracyFactor
}

// LogAndRestart logs the elapsed time since last start and restarts the timer
func (t *Timer) LogAndRestart(markerMessage string) {
	fmt.Printf("\n[Timer] %vns for %s", t.GetAndRestart(), markerMessage)
}

// LogAndContinue logs the elapsed time since last start. The timer will continue.
func (t *Timer) LogAndContinue(markerMessage string) {
	fmt.Printf("\n[Timer] %v ns for %s", t.GetAndContinue(), markerMessage)
}

var timeNow = func() time.Time {
	return time.Now()
}
