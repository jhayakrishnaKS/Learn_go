package main

import "fmt"

func main() {
	s := []int{1, 3, 2, 1, 2, 3, 45, 5, 3, 99, 22, 21, 2}
	m := make(map[int]int) 

	for _, num := range s { 
		m[num]++
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
