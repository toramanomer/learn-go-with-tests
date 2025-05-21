package iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	repeated := Repeat("hi", 3)
	fmt.Println(repeated)
	// Output: hihihi
}

func TestRepeat(t *testing.T) {
	var (
		want = "aaa"
		got  = Repeat("a", 3)
	)
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}
