package main

import "fmt"

func main() {
	// ch := make(chan int, 2)
	// go func() {
	// 	ch <- 42
	// 	ch <- 43

	// }()

	// value := <-ch

	// fmt.Println("Received:", value)
	// fmt.Println("Received:", <-ch)

	c := make(chan int)
	//send
	go foo(c)
	//receive
	bar(c)

	fmt.Println("about to exit")

}

//send
func foo(c chan<- int) {
	c <- 42
}

//receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}
