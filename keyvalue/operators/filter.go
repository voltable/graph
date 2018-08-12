package operators

// Filter is a set operator that returns the subset of those tuples satisfying the predicate
type Filter struct {
	iterator  Iterator
	predicate Predicate
}

// NewFilter returns a Filter
func NewFilter(predicate Predicate, iterator Iterator) *Filter {
	return &Filter{
		predicate: predicate,
		iterator:  iterator,
	}
}

func (s *Filter) Next() (interface{}, bool) {
	for i, ok := s.iterator.Next(); ok; i, ok = s.iterator.Next() {
		if s.predicate(i) {
			return i, true
		}
	}

	return nil, false
}
