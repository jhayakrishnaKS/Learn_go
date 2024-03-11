// Package dog allows us to more fully understand dogs.
package dog

// YearsTwo converts human years to dog years.
func YearTwo(n int) int {
	count := 0
	for i := 0; i < n; i++ {
		count += 7
	}
	return count
}

// Years converts human years to dog years.
func Years(n int) int {
	return n * 7
}
