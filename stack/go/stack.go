package stack

import "errors"

type Item int

type Stack struct {
	items []Item
}

func (s *Stack) empty() bool {
	return len(s.items) == 0
}

func (s *Stack) push(item Item) {
	s.items = append(s.items, item)
}

func (s *Stack) pop() (Item, error) {
	if s.empty() {
		return 0, errors.New("empty stack")
	}

	topIndex := len(s.items) - 1
	top := s.items[topIndex]

	s.items = s.items[:topIndex]

	return top, nil
}

func (s *Stack) top() (Item, error) {
	if s.empty() {
		return 0, errors.New("empty stack")
	}

	topIndex := len(s.items) - 1

	return s.items[topIndex], nil
}

func New() *Stack {
	return &Stack{[]Item{}}
}
