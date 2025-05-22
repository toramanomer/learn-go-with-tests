package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const (
	write = "write"
	sleep = "sleep"
)

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		var (
			writer  = bytes.NewBuffer(nil)
			sleeper = &SpyCountdownOperations{}
		)
		Countdown(writer, sleeper)
		var (
			wantWrites = "3\n2\n1\nGo!\n"
			gotWrites  = writer.String()
		)

		if wantWrites != gotWrites {
			t.Errorf("wanted: %q, got: %q", wantWrites, gotWrites)
		}
	})

	t.Run("sleep before every countdown", func(t *testing.T) {
		spySleeperWriter := &SpyCountdownOperations{}
		Countdown(spySleeperWriter, spySleeperWriter)
		var (
			wantCalls = []string{
				// 3
				write, sleep,
				// 2
				write, sleep,
				// 1
				write, sleep,
				// Go!
				write,
			}
			gotCalls = spySleeperWriter.Calls
		)

		if !reflect.DeepEqual(wantCalls, gotCalls) {
			t.Errorf("wanted calls: %v, got: %v", wantCalls, gotCalls)
		}
	})
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestConfigurableSleep(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{
		duration: sleepTime,
		sleep:    spyTime.Sleep,
	}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for: %v, but slept for: %v", sleepTime, spyTime.durationSlept)
	}
}
