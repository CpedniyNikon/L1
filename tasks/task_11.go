package tasks

import "fmt"

func Task11() {
	FirstSet := []string{"Dog", "Cat", "Bird", "Pig", "Raven", "Sparrow"}
	SecondSet := []string{"Bird", "Cat", "Ant", "Deer", "Beer", "Donkey", "Sparrow"}

	s := make(map[string]struct{})

	var IntersectionSet []string

	for _, val := range FirstSet {
		s[val] = struct{}{}
	}

	for _, val := range SecondSet {
		_, ok := s[val]
		if ok {
			IntersectionSet = append(IntersectionSet, val)
		}
	}

	fmt.Println(IntersectionSet)
}
