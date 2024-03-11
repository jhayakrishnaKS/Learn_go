package dog

import (
	"fmt"
	"testing"
)

func TestDog(t *testing.T) {
	i := Years(10)
	if i != 70 {
		t.Error("got", i, "expected", 70)
	}
}
func ExampleYears() {
	fmt.Println(Years(10))
	//output:
	//70
}

func BenchmarkYear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func TestDog2(t *testing.T) {
	j := YearTwo(20)
	if j != 140 {
		t.Error("got", j, "expected", 140)
	}
}

func ExampleYearTwo() {
	fmt.Println(YearTwo(20))
	//output:
	//140
}

func BenchmarkYearTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(20)
	}
}
