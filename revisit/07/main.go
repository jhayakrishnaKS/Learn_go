package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter = 0
	var wg sync.WaitGroup
	var m sync.Mutex
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			m.Lock()
			defer m.Unlock()
			counter++
			wg.Done() 
		}()
	}
	wg.Wait()
	fmt.Println(counter)
	fmt.Println("finished")
}
// →_→
// o(一︿一+)o
// X﹏X
