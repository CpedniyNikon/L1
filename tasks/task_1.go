package tasks

import "fmt"

type Human struct {
	Age  uint
	Name string
}

func (h Human) sayHello() {
	fmt.Println("i am only human after all")
}

type Action struct {
	Human
}

func (a Action) sayHello() {
	fmt.Println("i am only action after all")
}

func Task1() {
	Human := new(Human)
	Human.sayHello()

	Action := new(Action)
	Action.sayHello()
}
