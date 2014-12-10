package timer

import (
	"fmt"
	"time"
)

// Timer is a simple timer that prints elapsed time to console or returns it as nanoseconds
type Timer struct {
	lastCheckpoint time.Time
}

// NewAndStart returns a new timer and starts it immediately.
func NewAndStart() *Timer {
	t := Timer{}
	t.Start()
	return &t
}

// Start starts the timer, if it has already been started, the timer will restart
func (t *Timer) Start() {
	t.lastCheckpoint = timeNow()
}

// LogAndRestart logs the elapsed time since last start and restarts the timer
func (t *Timer) GetAndRestart() int64 {
	elapsed := timeNow().Sub(t.lastCheckpoint)
	t.lastCheckpoint = timeNow()
	return elapsed.Nanoseconds()
}

//LogAndContinue logs the elapsed time since last start. The timer will continue.
func (t *Timer) GetAndContinue() int64 {
	elapsed := timeNow().Sub(t.lastCheckpoint)
	return elapsed.Nanoseconds()
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
