package stackandqueue

import "errors"

type Stack struct {
	items []interface{}
}

func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (interface{}, error) {
	if len(s.items) > 0 {
		item := s.items[len(s.items)-1]
		s.items = s.items[:len(s.items)-1]
		return item, nil
	}

	return nil, errors.New("stack is empty")
}

func (s *Stack) Peak() (interface{}, error) {
	if len(s.items) > 0 {
		item := s.items[len(s.items)-1]
		return item, nil
	}

	return nil, errors.New("stack is empty")
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}
