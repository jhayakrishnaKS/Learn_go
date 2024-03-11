package main

import (
	"fmt"
	// "sync"
)

func main() {
	c := make(chan int)
	// var wg sync.WaitGroup
	// wg.Add(2)
	done := make(chan bool)
	n := 2

	for i := 0; i <= n; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
			// wg.Done()
			done <- true
		}()

	}

	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		c <- i
	// 	}
	// 	// wg.Done()
	// 	done <- true

	// }()

	// wg.Wait()
	go func() {
		<-done
		<-done
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
