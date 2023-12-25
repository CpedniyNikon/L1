package tasks

import (
	"fmt"
	"sync"
)

func Task3() {
	numbers := []int{2, 4, 6, 8, 10}
	sum := 0
	wg := sync.WaitGroup{}

	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			sum += n * n
		}(num)
	}

	wg.Wait()
	fmt.Println(sum)
}
