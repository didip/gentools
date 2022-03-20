package stack

import (
	"sync"
)

type Stack[T any] struct {
	lock   sync.Mutex
	values []T
}

func (s *Stack[T]) Push(value T) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.values = append(s.values, value)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack[T]) Pop() T {
	s.lock.Lock()
	defer s.lock.Unlock()

	el, index := s.peekInternal()
	if index == -1 {
		return el
	}

	s.values = s.values[:index]
	return el
}

func (s *Stack[T]) peekInternal() (T, int) {
	if s.IsEmpty() {
		var zero T
		return zero, -1
	}

	index := len(s.values) - 1
	el := s.values[index]
	return el, index
}

func (s *Stack[T]) Peek() T {
	el, _ := s.peekInternal()
	return el
}

func (s *Stack[T]) Size() int {
	return len(s.values)
}
