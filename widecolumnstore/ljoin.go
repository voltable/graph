package widecolumnstore

// LJoin a set operator that defines a nested-loop join.
type LJoin struct {
}

func (s *LJoin) Next() (interface{}, bool) {
	return nil, false
}

func (s *LJoin) op() {}
