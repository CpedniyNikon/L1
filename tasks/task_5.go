package tasks

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func Task5() {
	N := 5

	channel := make(chan int)

	wg := sync.WaitGroup{}
	defer wg.Wait()
	wg.Add(1)
	go func() {
		defer wg.Done()

		go func() {
			i := 0
			for {
				channel <- i
				fmt.Println("value " + strconv.Itoa(i) + " added to channel")
				time.Sleep(100 * time.Millisecond)
				i++
			}
		}()

		time.Sleep(time.Duration(N) * time.Second)
	}()

	for v := range channel {
		fmt.Println("value " + strconv.Itoa(v) + " is read from channel")
	}

}
