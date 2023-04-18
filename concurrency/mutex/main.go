package main

import (
	"fmt"
	"sync"
)

type State struct {
	mu    sync.RWMutex
	count int
}

func (s *State) setState(i int) {
	s.mu.Lock()
	defer s.mu.Unlock() // it unlocks mutex just before exiting from the function done via
	s.count = i
	//s.mu.Unlock()
}

func main() {
	state := State{}
	for i := 0; i < 10; i++ {
		state.count = i
	}

	fmt.Println(state)
}
