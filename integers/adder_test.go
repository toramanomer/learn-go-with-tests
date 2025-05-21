package integers

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestAdder(t *testing.T) {
	want, got := 5, Add(3, 2)
	if want != got {
		t.Errorf("expected %d, got %d", want, got)
	}
}
