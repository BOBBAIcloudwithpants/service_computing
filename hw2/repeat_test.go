package hw2

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 7)
	expected := "aaaaaaa"

	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i<b.N; i++ {
		Repeat("a", 7)
	}
}

func ExampleRepeat_10() {
	fmt.Println(Repeat("a", 10))
	// Output: aaaaaaaaaa
}

func ExampleRepeat_20() {
	fmt.Println(Repeat("ab", 20))
	// Output: abababababababababababababababababababab
}

