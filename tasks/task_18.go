package tasks

import (
	"fmt"
	"sync"
)

type Counter struct {
	mutex sync.Mutex
	Count uint
}

func NewCounter() *Counter {
	return &Counter{mutex: sync.Mutex{},
		Count: 0}
}

func (c *Counter) Increment() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.Count++
}

func Task18() {
	Counter := NewCounter()
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			Counter.Increment()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(Counter.Count)
}
