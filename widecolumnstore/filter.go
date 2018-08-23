package widecolumnstore

// Filter is a set operator that returns the subset of those tuples satisfying the predicate
type Filter struct {
	iterator  Operator
	predicate Predicate
}

// NewFilter returns a Filter
func NewFilter(predicate Predicate, iterator Operator) *Filter {
	return &Filter{
		predicate: predicate,
		iterator:  iterator,
	}
}

func (s *Filter) Next(i Iterator) Iterator {
	iterator := s.iterator.Next(i)
	return func() (interface{}, bool) {
		for value, ok := iterator(); ok; value, ok = iterator() {
			if s.predicate(value) {
				return value, true
			}
		}
		return nil, false
	}
}

func (s *Filter) op() {}
