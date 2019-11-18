package stack

// Stack a simple stack for the AST
type Stack struct {
	Data []interface{}
	Index int
}

// Push add's a item the the Stack
func (s *Stack) Push(v interface{})  {
	s.Index++
	s.Data = append(s.Data, v)
}

// Pop removes the last item on the Stack and returns it
func (s *Stack) Pop() interface{} {
	s.Index--
	if s.Index >= 0 {
		expr := s.Data[s.Index]
		s.Data = s.Data[:s.Index]
		return expr
	}
	return nil
}


// Peek returns the last item on the Stack without removing it
func (s Stack) Peek() interface{} {
	top := s.Index - 1
	if top >= 0 {
		return s.Data[top]
	}
	return nil
}
