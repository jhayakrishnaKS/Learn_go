package acdc

// Sum calculates the sum of a variable number of integers.
// It takes an arbitrary number of integers as input and returns their sum.
func Sum(xi ...int) int {
	s := 0
	for _, v := range xi {
		s += v
	}
	return s
}
