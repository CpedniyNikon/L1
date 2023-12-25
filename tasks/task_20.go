package tasks

import (
	"fmt"
	"strings"
)

func Task20() {
	String := "snow dog sun"

	Words := strings.Split(String, " ")

	lenWords := len(Words)
	for i := 0; i < lenWords/2; i++ {
		Words[i], Words[lenWords-i-1] = Words[lenWords-i-1], Words[i]
	}

	String = strings.Join(Words, " ")
	fmt.Println(String)
}
