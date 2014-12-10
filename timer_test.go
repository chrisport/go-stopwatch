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

		Convey("It should return 5, when time elapsed by 5 nanoseconds", func() {
			timerUnderTest := Timer{}

			timerUnderTest.Start()
			currentTimeInTest = currentTimeInTest.Add(5 * time.Nanosecond)
			measuredTimeAfter5 := timerUnderTest.GetAndContinue()

			So(measuredTimeAfter5, ShouldEqual, 5)
		})

		Convey("It should first return 3, then 5 when GetAndContinue is called after 3 and then after 2 nanoseconds", func() {
			timerUnderTest := Timer{}

			timerUnderTest.Start()
			currentTimeInTest = currentTimeInTest.Add(2 * time.Nanosecond)
			measuredTimeAfter2 := timerUnderTest.GetAndContinue()
			currentTimeInTest = currentTimeInTest.Add(3 * time.Nanosecond)
			measuredTimeAfter2Plus3 := timerUnderTest.GetAndContinue()

			So(measuredTimeAfter2, ShouldEqual, 2)
			So(measuredTimeAfter2Plus3, ShouldEqual, 5)
		})

		Convey("It should first return 3, then 2 when GetAndRestart is called after 3 and then after 2 nanoseconds", func() {
			timerUnderTest := Timer{}

			timerUnderTest.Start()
			currentTimeInTest = currentTimeInTest.Add(2 * time.Nanosecond)
			measuredTimeAfter2 := timerUnderTest.GetAndRestart()
			currentTimeInTest = currentTimeInTest.Add(3 * time.Nanosecond)
			measuredTimeAfter3 := timerUnderTest.GetAndRestart()

			So(measuredTimeAfter2, ShouldEqual, 2)
			So(measuredTimeAfter3, ShouldEqual, 3)
		})

		Convey("NewAndStart should start the returned timer", func() {
			timerUnderTest := NewAndStart()

			timerUnderTest.Start()
			currentTimeInTest = currentTimeInTest.Add(5 * time.Nanosecond)
			measuredTimeAfter5 := timerUnderTest.GetAndRestart()

			So(measuredTimeAfter5, ShouldEqual, 5)
		})

		timeNow = timeNowOriginal

		Convey("Timer should return around 100ms +/- 5ms after sleeping for 500ms should", func() {
			timerUnderTest := NewAndStart()
			time.Sleep(100 * time.Millisecond)

			measuredTimeAfter100millis := timerUnderTest.GetAndRestart()
			factor := 1000000
			So(measuredTimeAfter100millis, ShouldBeBetween, 95*factor, 105*factor)
		})

	})
}
