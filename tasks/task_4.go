package tasks

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

func Task4() {
	var count int
	_, err := fmt.Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	osChan := make(chan os.Signal)
	signal.Notify(osChan, os.Interrupt)
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(count)
	defer wg.Wait()
	for i := 1; i <= count; i++ {
		go func(i int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					log.Printf("end working worker %d", i)
					return
				case inputData := <-ch:
					log.Printf("worker %d, get data from main channel %d", i, inputData)
				}
			}
		}(i)
	}

	for {
		select {
		case <-osChan:
			cancel()
		case <-time.After(1 * time.Second):
			num := rand.Int()
			log.Printf("recording main goroutine value %d\n", num)
			ch <- num
		}
	}
}
