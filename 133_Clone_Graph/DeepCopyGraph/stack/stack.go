package stack

import "container/list"

//LIFO last in first out
type Stack struct {
	list.List
}

//O(1) operation
func (s *Stack) Push(item interface{}) {
	s.PushFront(item)
}

func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

//O(1) operation
func (s *Stack) Pop() interface{} {
	//always get the first item since its the most recently inserted item
	if !s.IsEmpty() {
		val := s.Front()
		s.Remove(val)
		return val.Value
	}
	return nil
}
