package sum

import (
	"fmt"
	"reflect"
	"testing"
)

func ExampleSum() {
	sum := Sum([]int{1, 2, 3})
	fmt.Println(sum)
	// Output: 6
}

func assertCorrectMessage(t testing.TB, got, want int, numbers []int) {
	t.Helper()

	if want != got {
		t.Errorf("got: %d, want: %d, given: %v", got, want, numbers)
	}
}

func TestSum(t *testing.T) {
	var (
		numbers = []int{1, 2, 3}
		want    = 6
		got     = Sum(numbers)
	)

	assertCorrectMessage(t, got, want, numbers)
}

func BenchmarkSum(b *testing.B) {
	for b.Loop() {
		Sum([]int{1, 2, 3, 4, 5, 6})
	}
}

func TestSumAll(t *testing.T) {
	t.Run("with nil slice", func(t *testing.T) {
		var (
			nilSlice []int
			want     = []int{0}
			got      = SumAll(nilSlice)
		)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got: %v, given: %v", want, got, nilSlice)
		}
	})

	t.Run("with empty slice", func(t *testing.T) {
		var (
			emptySlice = []int{}
			want       = []int{0}
			got        = SumAll(emptySlice)
		)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got: %v, given: %v", want, got, emptySlice)
		}
	})

	t.Run("with no args", func(t *testing.T) {
		var (
			want = []int{}
			got  = SumAll()
		)
		if !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got: %v", want, got)
		}
	})

	t.Run("with collection of numbers", func(t *testing.T) {
		var (
			numbers = [][]int{
				{1, 2},
				{3},
				{4, 5},
			}
			want = []int{3, 3, 9}
			got  = SumAll(numbers...)
		)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want: %v, got: %v, given: %v", want, got, numbers)
		}
	})

}

func BenchmarkSumAll(b *testing.B) {
	for b.Loop() {
		SumAll([]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8})
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int, numbers [][]int) {
		t.Helper()
		if !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got: %v, given: %v", want, got, numbers)
		}
	}

	t.Run("with at least two elements per slice", func(t *testing.T) {
		var (
			numbers = [][]int{{1, 2}, {0, 9}}
			want    = []int{2, 9}
			got     = SumAllTails(numbers...)
		)
		checkSums(t, got, want, numbers)

	})

	t.Run("with empty slice", func(t *testing.T) {
		var (
			numbers = [][]int{{}}
			want    = []int{0}
			got     = SumAllTails(numbers...)
		)
		checkSums(t, got, want, numbers)
	})
}
