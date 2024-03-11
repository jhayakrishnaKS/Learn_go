package main

import (
	"testing"
)

//table test
func TestSum(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}
	tests := []test{
		{[]int{21, 21}, 42},
		{[]int{1, 21}, 22},
		{[]int{2, 1, 2, 1}, 6},
		{[]int{12, 21}, 33},
	}

	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}

}
