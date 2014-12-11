package timer

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	Convey("Subject: Test Timer with approxiamte durations", t, func() {
		currentTimeInTest := time.Now()
		timeNowOriginal := timeNow
		timeNow = func() time.Time {
			return currentTimeInTest
		}

		Convey("NewStopwatch should start the returned timer", func() {
			timerUnderTest := NewStopwatch()

			currentTimeInTest = currentTimeInTest.Add(5 * time.Nanosecond)
			measuredTimeAfter5 := timerUnderTest.GetAndRestart()

			So(measuredTimeAfter5, ShouldEqual, 5*time.Nanosecond)
		})

		Convey("It should return 5, when time elapsed by 5 nanoseconds", func() {
			timerUnderTest := NewStopwatch()

			timerUnderTest.Restart()
			currentTimeInTest = currentTimeInTest.Add(5 * time.Nanosecond)
			measuredTimeAfter5 := timerUnderTest.GetAndContinue()

			So(measuredTimeAfter5, ShouldEqual, 5*time.Nanosecond)
		})

		Convey("It should first return 3, then 5 when GetAndContinue is called after 3 and then after 2 nanoseconds", func() {
			timerUnderTest := NewStopwatch()

			timerUnderTest.Restart()
			currentTimeInTest = currentTimeInTest.Add(2 * time.Nanosecond)
			measuredTimeAfter2 := timerUnderTest.GetAndContinue()
			currentTimeInTest = currentTimeInTest.Add(3 * time.Nanosecond)
			measuredTimeAfter2Plus3 := timerUnderTest.GetAndContinue()

			So(measuredTimeAfter2, ShouldEqual, 2)
			So(measuredTimeAfter2Plus3, ShouldEqual, 5*time.Nanosecond)
		})

		Convey("It should first return 3, then 2 when GetAndRestart is called after 3 and then after 2 nanoseconds", func() {
			timerUnderTest := NewStopwatch()

			timerUnderTest.Restart()
			currentTimeInTest = currentTimeInTest.Add(2 * time.Nanosecond)
			measuredTimeAfter2 := timerUnderTest.GetAndRestart()
			currentTimeInTest = currentTimeInTest.Add(3 * time.Nanosecond)
			measuredTimeAfter3 := timerUnderTest.GetAndRestart()

			So(measuredTimeAfter2, ShouldEqual, 2*time.Nanosecond)
			So(measuredTimeAfter3, ShouldEqual, 3*time.Nanosecond)
		})

		Convey("When accuracy is set to microseconds, timer should return 7 when 7ms are elapsed", func() {
			timerUnderTest := NewStopwatch()
			timerUnderTest.SetAccuracy(time.Millisecond)
			currentTimeInTest = currentTimeInTest.Add(7 * time.Millisecond)
			measuredTime := timerUnderTest.GetAndRestart()

			So(measuredTime, ShouldEqual, 7*time.Millisecond)
		})

		Convey("When accuracy is set to seconds, timer should return 13 when 13s are elapsed", func() {
			timerUnderTest := NewStopwatch()
			timerUnderTest.SetAccuracy(time.Second)
			currentTimeInTest = currentTimeInTest.Add(13 * time.Second)
			measuredTime := timerUnderTest.GetAndRestart()

			So(measuredTime, ShouldEqual, 13*time.Second)
		})

		timeNow = timeNowOriginal

		Convey("Timer should return around 100ms +/- 5ms after sleeping for 500ms should", func() {
			timerUnderTest := NewStopwatch()
			time.Sleep(100 * time.Millisecond)

			measuredTime := timerUnderTest.GetAndRestart()
			factor := 1000000
			So(measuredTime, ShouldBeBetween, 95*factor, 105*factor)
		})

	})
}

func BenchmarkGetAndRestart(b *testing.B) {
	stopwatch := NewStopwatch()
	for n := 0; n < b.N; n++ {
		stopwatch.GetAndRestart()
	}
}

func BenchmarkGetAndContinue(b *testing.B) {
	stopwatch := NewStopwatch()
	for n := 0; n < b.N; n++ {
		stopwatch.GetAndContinue()
	}
}
