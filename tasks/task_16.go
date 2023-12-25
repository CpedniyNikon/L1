package tasks

import (
	"fmt"
	"math/rand"
)

func swap(a, b *int) {
	*a = *a ^ *b
	*b = *b ^ *a
	*a = *a ^ *b
}

func Partition(arr []int, l, r int) int {
	x := arr[l+rand.Int()%(r-l)]
	i := l
	j := r
	for i <= j {
		for arr[i] < x {
			i++
		}
		for arr[j] > x {
			j--
		}
		if i >= j {
			break
		}
		swap(&arr[i], &arr[j])
		i++
		j--
	}
	return j
}

func QuickSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	p := Partition(arr, l, r)
	QuickSort(arr, l, p)
	QuickSort(arr, p+1, r)
}

func Task16() {
	arr := []int{1, 6, 78, 32, 8, 9, 2, 35, 9, 0, 2342, 4, 7567, 2, 5, 8, 0, 2, 7, -42, 1, 6}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
