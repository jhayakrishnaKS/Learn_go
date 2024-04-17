package main

import (
	"fmt"
	"sort"
)

// custom sort
type person struct {
	first string
	age   int
}

type ByAge []person

func (s ByAge) Len() int {
	return len(s)
}

func (s ByAge) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByAge) Less(i, j int) bool {
	return s[i].age < s[j].age
}

type ByName []person

func (b ByName) Len() int           { return len(b) }
func (b ByName) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByName) Less(i, j int) bool { return b[i].first < b[j].first }

func main() {
	xs := []string{"hi", "hello", "aloha", "jumbo"}
	xi := []int{9, 3, 33, 4, 67, 11}
	fmt.Println("-----before sort------")
	fmt.Println(xs)
	fmt.Println(xi)
	fmt.Println("-----after sort------")
	sort.Ints(xi)
	fmt.Println(xi)
	sort.Strings(xs)
	fmt.Println(xs)

	// Custom sort
	fmt.Println("-----custom sort------")
	p1 := person{"zames", 32}
	p2 := person{"bond", 22}
	p3 := person{"kaka", 42}
	p4 := person{"lava", 62}

	people := []person{p1, p2, p3, p4}
	fmt.Println(people)
	fmt.Println("-----after custom sort------")
	sort.Sort(ByAge(people))
	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)

}
