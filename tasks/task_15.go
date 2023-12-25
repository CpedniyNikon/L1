package tasks

import (
	"strings"
)

var justString string

func createHugeString(length int) string {
	return strings.Repeat("A", length)
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func Task15() {
	someFunc()
}
