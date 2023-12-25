package tasks

import (
	"fmt"
	"strconv"
	"sync"
)

func PutDigitsInChannel(nums []int) chan int {
	Channel := make(chan int)
	go func() {
		wg := sync.WaitGroup{}
		for i := 0; i < len(nums); i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				fmt.Println("value " + strconv.Itoa(nums[i]) + " added to channel")
				Channel <- nums[i]
			}(i)
		}
		wg.Wait()
		close(Channel)
	}()
	return Channel
}

func TransmitToChannel(c chan int) chan int {
	Channel := make(chan int)
	go func() {
		for val := range c {
			fmt.Println("value " + strconv.Itoa(val*2) + " transmitted to channel")
			Channel <- val * 2
		}
		close(Channel)
	}()
	return Channel
}

func Task9() {
	Arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	First := PutDigitsInChannel(Arr)
	Second := TransmitToChannel(First)
	for val := range Second {
		fmt.Println("value " + strconv.Itoa(val) + " read from channel")
	}
}
