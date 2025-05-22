package main

import (
	"fmt"
	"io"
	"iter"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (s *ConfigurableSleeper) Sleep() {
	s.sleep(s.duration)
}

func countDownFrom(from int) iter.Seq[int] {
	return func(yield func(i int) bool) {
		for i := from; i > 0; i-- {
			if !yield(i) {
				return
			}
		}
	}
}

func Countdown(w io.Writer, sleeper Sleeper) {
	for i := range countDownFrom(3) {
		fmt.Fprintln(w, i)
		sleeper.Sleep()
	}

	fmt.Fprintln(w, finalWord)
}

type DefaultSleeper struct{}

func (sleeper *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	sleeper := &ConfigurableSleeper{
		duration: 1 * time.Second,
		sleep:    time.Sleep,
	}
	Countdown(os.Stdout, sleeper)
}
