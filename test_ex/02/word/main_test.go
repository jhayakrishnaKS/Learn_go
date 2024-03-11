package word

import (
	"common/test_ex/02/quote"
	"fmt"
	"testing"
)

func TestUseCount(t *testing.T) {
	n := UseCount("one two three three")
	for k, v := range n {
		switch k {
		case "one":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "two":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "three":
			if v != 2 {
				t.Error("got", v, "want", 2)
			}
		}
	}

}

func TestCount(t *testing.T) {
	n := Count("one two three")
	if n != 3 {
		t.Error("got", n, "want", 3)
	}
}

func ExampleCount() {
	fmt.Println(Count("one two three"))
	//output:
	//3
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}
func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}
