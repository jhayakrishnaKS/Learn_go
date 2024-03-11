package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}
func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("Arch\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	// foo()
	bar()
	wg.Wait()

	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
}
