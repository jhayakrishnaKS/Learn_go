package main

import "fmt"

func main() {
	//ex6
	// c := make(chan int)

	// go func() {
	// 	for i := 0; i < 100; i++ {
	// 		c <- i
	// 	}
	// 	close(c)
	// }()
	// for b := range c {
	// 	fmt.Println(b)

	// }
	// fmt.Print("about to exit")

	//ex7
	c := make(chan int)

	for j := 0; j < 10; j++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()
	}

	for k := 0; k < 100; k++ {
		fmt.Println(k, <-c)
	}

	fmt.Println("about to exit")
}
