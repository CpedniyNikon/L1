package tasks

import (
	"fmt"
	"sync"
)

func Task2() {
	numbers := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}

	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			square := n * n

			fmt.Println(square)
		}(num)
	}

	wg.Wait()
}
