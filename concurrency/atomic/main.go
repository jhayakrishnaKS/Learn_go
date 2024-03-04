package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	// Print the number of CPUs available
	fmt.Println("CPUs:", runtime.NumCPU())

	// Print the number of goroutines currently running
	fmt.Println("Goroutines-start:", runtime.NumGoroutine())

	// Declare a counter variable of type int64 to be used with atomic operations
	var counter int64

	// Define the number of goroutines to be launched
	const gs = 100

	// Create a WaitGroup to synchronize goroutines
	var wg sync.WaitGroup
	wg.Add(gs)

	// Loop to launch goroutines
	for i := 0; i < gs; i++ {
		// Launch a goroutine
		go func() {
			// Atomically increment the counter by 1
			atomic.AddInt64(&counter, 1)

			// Yield the processor to allow other goroutines to run
			runtime.Gosched()

			// Atomically load and print the current value of the counter
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))

			// Notify the WaitGroup that this goroutine is done
			wg.Done()
		}()

		// Print the number of goroutines currently running (may vary)
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Print the final number of goroutines
	fmt.Println("Goroutines-last:", runtime.NumGoroutine())

	// Print the final value of the counter
	fmt.Println("count:", counter)
}
