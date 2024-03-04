package main

import "fmt"

func main() {
	// normal
	// c := make(chan int, 2)
	// go func() {
	// 	c <- 42
	// 	c <- 45
	// }()

	// fmt.Println(<-c)
	// fmt.Println(<-c)

	// buffered
	c := make(chan int, 2)

	c <- 42
	c <- 45
	fmt.Println(<-c)
	fmt.Println(<-c)

}
