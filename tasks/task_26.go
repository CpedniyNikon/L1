package tasks

import (
	"fmt"
	"strings"
)

func Task26() {
	str := "abCdec"

	str = strings.ToLower(str)

	s := make(map[byte]int)

	for i := 0; i < len(str); i++ {
		if s[str[i]] != 0 {
			fmt.Println("есть повторение")
			break
		}
		s[str[i]]++
	}
}
