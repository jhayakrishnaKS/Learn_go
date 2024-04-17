package main

import "fmt"

func main() {
	// var a, b int
	// a = 20
	// b = 30
	// found := false

	// fmt.Print("Numbers between ", a, " and ", b, " that are divisible by 7 but not by 5: ")
	// for i := a; i <= b; i++ {
	//     if i%7 == 0 && i%5 != 0 {
	//         if found {
	//             fmt.Print(", ")
	//         }
	//         fmt.Print(i)
	//         found = true
	//     }
	// }
	// fmt.Println()
	// a := 10
	// m := make(map[int]int)
	// for i := 1; i <= a; i++ {
	// 	m[i] = i * i
	// }
	// fmt.Println(m)

	// for i := 2; i <= 5; i++ {
	// 	for j := i; j <= i+3; j++ {
	// 		fmt.Print(j)
	// 	}
	// 	fmt.Println()
	// }
	a := 10
	b := 20
	Swap(&a, &b)
	fmt.Println(a, b)
}

func Swap(a, b *int) {
	*a, *b = *b, *a
	fmt.Println(*a, *b)
}
