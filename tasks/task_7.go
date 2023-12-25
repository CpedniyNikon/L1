package tasks

import (
	"fmt"
	"sync"
)

type Map struct {
	m     map[int]struct{}
	mutex sync.Mutex
}

func NewMap() *Map {
	return &Map{
		m:     make(map[int]struct{}),
		mutex: sync.Mutex{},
	}
}

func (m *Map) Add(value int) {
	m.mutex.Lock()
	m.m[value] = struct{}{}
	m.mutex.Unlock()
}

func (m *Map) Print() {
	for v := range m.m {
		fmt.Println(v)
	}
}

func Task7() {
	m := NewMap()

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(value int) {
			defer wg.Done()
			m.Add(value)
		}(i)
	}
	wg.Wait()
	m.Print()
}
