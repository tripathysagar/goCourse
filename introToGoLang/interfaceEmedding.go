package main

import "fmt"

type Putter interface {
	Put(id int, val any)
}

type simplePutter struct{}

func (s *simplePutter) Put(id int, val any) {
	fmt.Println("Inside putter Put func")

}

type Storage interface {
	Get(id int) (any, error)
	Putter
}

type FooStorage struct{}

func (s *FooStorage) Get(id int) (any, error) {
	fmt.Println("Inside storage Get func")
	return nil, nil
}

func (s *FooStorage) Put(id int, val any) {
	fmt.Println("Inside storage Put func")

}

func updateValue(id int, val any, p Putter) {
	p.Put(id, val)
}

type Server struct {
	store Storage
}

func main() {
	s := &Server{
		store: &FooStorage{},
	}

	s.store.Get(1)
	s.store.Put(1, "Punk")
	updateValue(1, "up town", s.store)
	// this will not work as it is a trying to create obj with the server type will thow an compile time error "missing method GEt"
	//sp := &Server{
	//	store: &simplePutter{},
	//}
	//updateValue(1, "up town", sp.store)
	sputter := &simplePutter{}

	updateValue(1, "up town", sputter)

}
