package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("krish")
	if s != "welcome my bro:krish" {
		t.Error("got", s, "expected", "welcome my bro:krish")
	}
}
func ExampleGreet() {
	fmt.Println(Greet("krish"))
	//output:
	//welcome my bro:krish
}
func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("krish")
	}
}
