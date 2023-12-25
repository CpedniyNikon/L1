package tasks

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func StopInterrupt1() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	wg := sync.WaitGroup{}
	wg.Add(1)
	defer wg.Wait()
	go func() {
		defer wg.Done()

		go func() {
			for i := 0; i < 10; i++ {
				fmt.Println(i)
				time.Sleep(time.Second)
			}
		}()

		select {
		case <-ctx.Done():
			cancel()
			fmt.Println("context canceled")
			close(c)
		}

	}()
	for range c {
		cancel()
	}
}

func StopInterrupt2() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(1)
	defer wg.Wait()
	go func() {
		defer wg.Done()

		go func() {
			for i := 0; i < 10; i++ {
				fmt.Println(i)
				time.Sleep(time.Second)
			}
		}()

		select {
		case <-ctx.Done():
			fmt.Println("context canceled")
		}
	}()

}

func Task6() {
	StopInterrupt1()
}
