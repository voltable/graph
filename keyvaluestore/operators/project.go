package operators

// Project is a set operator that projects a set of tuples onto the specified attributes
type Project struct {
}

func (s *Project) Next() (interface{}, bool) {
	return nil, false
}
