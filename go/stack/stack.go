package stack

import "errors"

type Stack[T comparable] struct {
	items []T
}

func (s *Stack[T]) empty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) pop() (T, error) {
	if s.empty() {
		return *new(T), errors.New("empty stack")
	}

	topIndex := len(s.items) - 1
	top := s.items[topIndex]

	s.items = s.items[:topIndex]

	return top, nil
}

func (s *Stack[T]) top() (T, error) {
	if s.empty() {
		return *new(T), errors.New("empty stack")
	}

	topIndex := len(s.items) - 1

	return s.items[topIndex], nil
}

func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{[]T{}}
}
