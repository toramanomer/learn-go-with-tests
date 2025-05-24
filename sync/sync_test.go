package main

import (
	"sync"
	"testing"
)

func assertCounter(t testing.TB, counter *Counter, want int) {
	t.Helper()

	if got := counter.Value(); got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
}

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, &counter, 3)
	})

	t.Run("it runes safely concurrently", func(t *testing.T) {
		counter := Counter{}
		wantedCount := 1000

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for range wantedCount {
			go func() {
				defer wg.Done()
				counter.Inc()
			}()
		}

		wg.Wait()

		assertCounter(t, &counter, wantedCount)
	})
}

func BenchmarkCounter(b *testing.B) {
	counter := Counter{}

	for b.Loop() {
		counter.Inc()
	}
}
