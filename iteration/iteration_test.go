package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestRepeatWithStringBuilder(t *testing.T) {
	repeated := RepeatWithStringsBuilder("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func BenchmarkRepeatWithStringsBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatWithStringsBuilder("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 3)
	fmt.Println(repeated)
	// Output: aaa
}

func ExampleRepeatWithStringsBuilder() {
	repeated := RepeatWithStringsBuilder("a", 3)
	fmt.Println(repeated)
	// Output: aaa
}
