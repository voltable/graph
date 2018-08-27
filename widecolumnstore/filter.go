package widecolumnstore

// Filter is a set operator that returns the subset of those tuples satisfying the predicate
type Filter struct {
	storage  Storage
	iterator Operator
	prefix   Prefix
}

// NewFilter returns a Filter
func NewFilter(storage Storage, iterator Operator, prefix Prefix) *Filter {
	return &Filter{
		prefix:   prefix,
		iterator: iterator,
		storage:  storage,
	}
}

func (s *Filter) Next(i ...Iterator) Iterator {
	first := i[0]
	iterator := s.iterator.Next(first)
	var prefixIterator Iterator
	return func() (KeyValue, bool) {
		for value, ok := prefixIterator(); ok; value, ok = prefixIterator() {
			return value, ok
		}
		for value, ok := iterator(); ok; value, ok = iterator() {
			prefix := s.prefix(value)

			prefixIterator = s.storage.HasPrefix(prefix)
			value, ok := prefixIterator()
			if ok {
				return value, ok
			}
		}
		return KeyValue{}, false
	}
}

func (s *Filter) op() {}
