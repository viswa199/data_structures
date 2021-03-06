package stack

import "fmt"

type Stack struct {
	t []interface{}
}

func (s *Stack) Push(key interface{}) {
	s.t = append(s.t, key)
}

func (s *Stack) Length() int {
	return len(s.t)
}

func (s *Stack) Pop() interface{} {
	deleting_element := s.t[len(s.t)-1]
	s.t = s.t[:len(s.t)-1]
	return deleting_element
}

func (s *Stack) Display() {
	for _, v := range s.t {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}

func (s *Stack) Contains(key interface{}) bool {
	for _, v := range s.t {
		if v == key {
			return true
		}
	}
	return false
}

func (s *Stack) Empty() bool {
	if len(s.t) == 0 {
		return true
	}
	return false
}

func (s *Stack) Top() interface{} {
	return s.t[len(s.t)-1]
}
