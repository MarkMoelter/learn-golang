package countdown_test

import (
	"bytes"
	"reflect"
	"testing"
	"time"

	"github.com/MarkMoelter/learn-golang/countdown"
)

func TestCountdown(t *testing.T) {
	t.Run("use buffer", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &countdown.SpySleeper{}

		countdown.Countdown(buffer, spySleeper)

		got := buffer.String()
		want := "3\n2\n1\nGo!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		if spySleeper.Calls != 3 {
			t.Errorf("not enough calls to sleeper, want 3 got %d", spySleeper.Calls)
		}
	})
	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &countdown.SpyCountdownOperations{}
		countdown.Countdown(spySleepPrinter,spySleepPrinter)

		want := []string{
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	t.Run("", func(t *testing.T) {
		sleepTime := 5 * time.Second

		spyTime := &countdown.SpyTime{}
		sleeper := countdown.ConfigurableSleeper{sleepTime, spyTime.Sleep}
		sleeper.Sleep()

		if spyTime.DurationSlept != sleepTime {
			t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.DurationSlept)
		}
	})
}
