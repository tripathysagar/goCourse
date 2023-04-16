package main

import "fmt"

type CustomMap[K comparable, V any] struct {
	data map[K]V
}

func NewCustomMap[K comparable, V any]() *CustomMap[K, V] {
	return &CustomMap[K, V]{
		data: make(map[K]V),
	}
}

func (m *CustomMap[K, V]) Insert(k K, v V) error {
	m.data[k] = v
	return nil
}

func main() {
	M := NewCustomMap[string, int]()
	M.Insert("a", 1)
	M.Insert("b", 2)
	fmt.Printf("%+v\n", M)

	M2 := NewCustomMap[any, any]()
	M2.Insert("a", 1)
	M2.Insert(2, "b")
	fmt.Printf("%+v\n", M2)
}
