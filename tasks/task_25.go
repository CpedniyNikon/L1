package tasks

import (
	"fmt"
	"time"
)

func Sleep(duration time.Duration) {
	<-time.After(duration)
}

func Task25() {
	fmt.Println("start")
	Sleep(3 * time.Second)
	fmt.Println("end")
}
