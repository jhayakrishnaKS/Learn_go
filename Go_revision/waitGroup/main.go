package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	fmt.Println("waitgroup")
	wg.Add(2)
	go func(){
		fmt.Println("ye1")
		wg.Done()
	}()
	go func ()  {
		fmt.Println("go go go go go")
		wg.Done()	
	}()
	wg.Wait()
	fmt.Println("carnival is done")
}