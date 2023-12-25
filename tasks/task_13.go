package tasks

import "fmt"

func Task13_1() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Print(a)
	fmt.Print(" ")
	fmt.Println(b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Print(a)
	fmt.Print(" ")
	fmt.Println(b)
}

func Task13_2() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Print(a)
	fmt.Print(" ")
	fmt.Println(b)
	a = a ^ b
	b = b ^ a
	a = a ^ b
	fmt.Print(a)
	fmt.Print(" ")
	fmt.Println(b)
}
