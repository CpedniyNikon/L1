package tasks

import (
	"fmt"
	"reflect"
)

func ShowMeTheType(i interface{}) string {
	return reflect.TypeOf(i).String()
}

func Task14() {
	fmt.Println(ShowMeTheType(1))
	fmt.Println(ShowMeTheType([]int{1, 2, 3}))
	fmt.Println(ShowMeTheType("abc"))
	fmt.Println(ShowMeTheType(true))
	fmt.Println(ShowMeTheType(2.5))
	fmt.Println(ShowMeTheType(make(chan int)))

}
