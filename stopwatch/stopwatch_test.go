package stopwatch

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestStopwatch(t *testing.T) {
	Convey("Subject: Test Stopwatch with approxiamte durations", t, func() {
		currentTimeInTest := time.Now()
		timeNowOriginal := timeNow
		timeNow = func() time.Time {
			return currentTimeInTest
		}

		Convey("NewStopwatch should start the returned stopwatch", func() {
			stopwatchUnderTest := NewStopwatch()

			currentTimeInTest = currentTimeInTest.Add(5 * time.Nanosecond)
			measuredTimeAfter5 := stopwatchUnderTest.GetAndRestart()

			So(measuredTimeAfter5, ShouldEqual, 5*time.Nanosecond)
		})

		Convey("It should return 5, when time elapsed by 5 nanoseconds", func() {
			stopwatchUnderTest := NewStopwatch()

			stopwatchUnderTest.Restart()
			currentTimeInTest = currentTimeInTest.Add(5 * time.Nanosecond)
			measuredTimeAfter5 := stopwatchUnderTest.GetAndContinue()

			So(measuredTimeAfter5, ShouldEqual, 5*time.Nanosecond)
		})

		Convey("It should first return 3, then 5 when GetAndContinue is called after 3 and then after 2 nanoseconds", func() {
			stopwatchUnderTest := NewStopwatch()

			stopwatchUnderTest.Restart()
			currentTimeInTest = currentTimeInTest.Add(2 * time.Nanosecond)
			measuredTimeAfter2 := stopwatchUnderTest.GetAndContinue()
			currentTimeInTest = currentTimeInTest.Add(3 * time.Nanosecond)
			measuredTimeAfter2Plus3 := stopwatchUnderTest.GetAndContinue()

			So(measuredTimeAfter2, ShouldEqual, 2)
			So(measuredTimeAfter2Plus3, ShouldEqual, 5*time.Nanosecond)
		})

		Convey("It should first return 3, then 2 when GetAndRestart is called after 3 and then after 2 nanoseconds", func() {
			stopwatchUnderTest := NewStopwatch()

			stopwatchUnderTest.Restart()
			currentTimeInTest = currentTimeInTest.Add(2 * time.Nanosecond)
			measuredTimeAfter2 := stopwatchUnderTest.GetAndRestart()
			currentTimeInTest = currentTimeInTest.Add(3 * time.Nanosecond)
			measuredTimeAfter3 := stopwatchUnderTest.GetAndRestart()

			So(measuredTimeAfter2, ShouldEqual, 2*time.Nanosecond)
			So(measuredTimeAfter3, ShouldEqual, 3*time.Nanosecond)
		})

		Convey("When accuracy is set to microseconds, stopwatch should return 7 when 7ms are elapsed", func() {
			stopwatchUnderTest := NewStopwatch()
			stopwatchUnderTest.SetAccuracy(time.Millisecond)
			currentTimeInTest = currentTimeInTest.Add(7 * time.Millisecond)
			measuredTime := stopwatchUnderTest.GetAndRestart()

			So(measuredTime, ShouldEqual, 7*time.Millisecond)
		})

		Convey("When accuracy is set to seconds, stopwatch should return 13 when 13s are elapsed", func() {
			stopwatchUnderTest := NewStopwatchWithAccuracy(time.Second)

			currentTimeInTest = currentTimeInTest.Add(13 * time.Second)
			measuredTime := stopwatchUnderTest.GetAndRestart()

			So(measuredTime, ShouldEqual, 13*time.Second)
		})

		timeNow = timeNowOriginal

		Convey("Stopwatch should return around 100ms +/- 5ms after sleeping for 500ms should", func() {
			stopwatchUnderTest := NewStopwatch()
			time.Sleep(100 * time.Millisecond)

			measuredTime := stopwatchUnderTest.GetAndRestart()
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
