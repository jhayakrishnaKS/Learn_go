package main

import "fmt"

func unique(slice []int) *[]int {
	us := []int{}

	for i := 0; i < len(slice); i++ {
		isUnique := true
		for j := 0; j < i; j++ {
			if slice[i] == slice[j] {
				isUnique = false
				break
			}
		}
		if isUnique {
			us = append(us, slice[i])
		}
	}
	return &us
}

func main() {
	og := []int{1, 23, 4, 55, 3, 2, 1, 22, 33, 22, 2}
	us := unique(og)
	fmt.Println(og)
	fmt.Println(us)
}
