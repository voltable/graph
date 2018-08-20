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

func (s *Filter) Next(i Iterator) (interface{}, bool) {
	for value, ok := s.iterator.Next(i); ok; value, ok = s.iterator.Next(i) {
		if s.predicate(value) {
			return value, true
		}
	}

	return nil, false
}
