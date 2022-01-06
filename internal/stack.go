package internal

type Stack struct {
	values []interface{}
	size   int
}

func NewStack() *Stack {
	s := &Stack{}
	s.values = make([]interface{}, 0)
	return s
}

func (s *Stack) IsEmpty() bool {
	return (s.size == 0)
}

func (s *Stack) Push(val interface{}) {
	s.values = append(s.values, val)
	s.size++
}

func (s *Stack) Pop() interface{} {
	s.size--
	value := s.values[s.size]
	s.values = s.values[:s.size]
	return value
}
