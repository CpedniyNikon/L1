package tasks

import "fmt"

func Task23() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr = append(arr[:2], arr[2+1:]...)
	fmt.Println(arr)
}
