package tasks

import (
	"fmt"
)

func BinarySearch(arr []int, value int) int {
	l := 0
	r := len(arr) - 1
	var middle int
	for l <= r {
		middle = l + (r-l)/2
		if arr[middle] == value {
			return middle
		} else if arr[middle] > value {
			r = middle - 1
		} else if arr[middle] < value {
			l = middle + 1
		}
	}
	return -1
}

func Task17() {
	arr := []int{-1, 0, 3, 5, 9, 12}

	idx := BinarySearch(arr, 9)
	fmt.Println(idx)
}
