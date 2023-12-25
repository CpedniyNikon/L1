package tasks

import "fmt"

type Set struct {
	s      map[string]struct{}
	length uint
}

func NewSet() *Set {
	return &Set{
		s:      make(map[string]struct{}),
		length: 0,
	}
}

func (m *Set) Insert(key string) {
	_, ok := m.Find(key)
	if ok {
		return
	}
	m.s[key] = struct{}{}
	m.length++
}

func (m *Set) Find(key string) (struct{}, bool) {
	val, ok := m.s[key]
	return val, ok
}

func (m *Set) Erase(val string) {
	_, ok := m.s[val]
	if !ok {
		return
	}
	delete(m.s, val)
	m.length--
}

func (m *Set) Size() uint {
	return m.length
}

func (m *Set) Empty() bool {
	return m.Size() == 0
}

func Task12() {
	Strings := []string{"cat", "cat", "dog", "cat", "tree"}
	Set := NewSet()
	for _, val := range Strings {
		Set.Insert(val)
	}
	fmt.Println(Set.Size())
}
