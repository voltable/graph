package operators

type Sort struct {
}

func (s *Sort) Next() (interface{}, bool) {
	return nil, false
}

func (s *Sort) Op() {}
