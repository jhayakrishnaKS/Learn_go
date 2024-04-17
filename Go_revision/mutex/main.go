package main

import (
	"fmt"
	"sync"
	"time"
)

type counter struct {
	value int
	mutex sync.Mutex
}

func (c *counter) Increment() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value++
}

func (c *counter) Decrement() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value--
}

func (c *counter) GetValue() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.value
}

func main() {
	counter := counter{}

	go func() {
		for i := 0; i < 5; i++ {
			counter.Increment()
			time.Sleep(time.Millisecond * 100)
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			counter.Decrement()
			time.Sleep(time.Millisecond * 200)
		}
	}()

	go func() {
		for {
			fmt.Println("Counter value:", counter.GetValue())
			time.Sleep(time.Millisecond * 300)
		}
	}()

	time.Sleep(time.Second * 2)
}
