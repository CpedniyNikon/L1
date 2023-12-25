package tasks

import "fmt"

func Task8() {

	var digit int64
	var index int
	fmt.Scan(&digit)
	fmt.Scan(&index)

	digit = digit | (1 << (index - 1))

	fmt.Println(digit)
}
