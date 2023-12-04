package stackandqueue

import "errors"

type Queue struct {
	items []interface{}
}

func (s *Queue) EnQueue(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Queue) DeQueue() (interface{}, error) {
	if len(s.items) > 0 {
		item := s.items[0]
		s.items = s.items[1:]
		return item, nil
	}

	return nil, errors.New("queue is empty")
}

func (s *Queue) Front() (interface{}, error) {
	if len(s.items) > 0 {
		item := s.items[0]
		return item, nil
	}

	return nil, errors.New("queue is empty")
}

func (s *Queue) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Queue) Size() int {
	return len(s.items)
}
