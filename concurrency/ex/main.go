package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("cpu:", runtime.NumCPU())
	fmt.Println("go routine:", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("ye's carnival")
		wg.Done()
	}()
	go func() {
		fmt.Println("Go Go Go Go Go")
		wg.Done()
	}()
	fmt.Println("mid-cpu:", runtime.NumCPU())
	fmt.Println("mid-go routine:", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("oh oh oh oh")
	fmt.Println("end-cpu:", runtime.NumCPU())
	fmt.Println("end-go routine:", runtime.NumGoroutine())

}
