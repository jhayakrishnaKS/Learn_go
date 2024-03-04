package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// Print the number of CPUs available
	fmt.Println("CPUs:", runtime.NumCPU())

	// Print the number of goroutines currently running before launching new ones
	fmt.Println("Goroutines-start:", runtime.NumGoroutine())

	// Initialize a counter variable
	counter := 0

	// Define the number of goroutines to be launched
	const gs = 10

	// Create a WaitGroup to synchronize goroutines
	var wg sync.WaitGroup
	wg.Add(gs)

	// Create a mutex to synchronize access to the counter variable
	var mu sync.Mutex

	// Loop to launch goroutines
	for i := 0; i < gs; i++ {
		// Launch a goroutine
		go func() {
			// Lock the mutex to ensure exclusive access to the counter variable
			mu.Lock()

			// Read the current value of counter
			v := counter
			// Yield the processor to allow other goroutines to run
			runtime.Gosched()
			// Increment the value read from counter
			v++
			// Write the incremented value back to counter
			counter = v

			// Unlock the mutex to release the lock
			mu.Unlock()

			// Notify the WaitGroup that this goroutine is done
			wg.Done()
		}()
		// Print the number of goroutines currently running (may vary)
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	// Wait for all goroutines to finish
	wg.Wait() 

	// Print the final value of the counter
	fmt.Println("count:", counter)
}
