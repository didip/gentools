package queue

import (
	"sync"
)

type Queue[T any] struct {
	lock   sync.Mutex
	values []T
}

func (s *Queue[T]) Enqueue(value T) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.values = append(s.values, value)
}

func (s *Queue[T]) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Queue[T]) Dequeue() T {
	s.lock.Lock()
	defer s.lock.Unlock()

	el := s.Peek()

	s.values = s.values[1:len(s.values)]
	return el
}

func (s *Queue[T]) Peek() T {
	if s.IsEmpty() {
		var zero T
		return zero
	}

	el := s.values[0]
	return el
}

func (s *Queue[T]) Size() int {
	return len(s.values)
}
