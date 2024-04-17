package main

import "fmt"

func RemoveDuplicatesFromSlice(nums []int) []int {
	seen := make(map[int]bool)
	result := []int{}
	for _, num := range nums {
		if !seen[num] {
			result = append(result, num)
			seen[num] = true
		}
	}
	return result
}

func main() {
	slice := []int{1, 3, 2, 2, 1, 34, 67, 21, 1, 44, 234, 55}
	fmt.Println("Original slice:", slice)
	uniqueSlice := RemoveDuplicatesFromSlice(slice)
	fmt.Println("Slice with duplicates removed:", uniqueSlice)

}