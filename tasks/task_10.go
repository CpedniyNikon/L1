package tasks

import (
	"fmt"
)

func Task10() {

	Arr := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	m := make(map[int][]float32)

	for _, val := range Arr {
		m[int(val)/10*10] = append(m[int(val)/10*10], val)
	}

	fmt.Println(m)

}
