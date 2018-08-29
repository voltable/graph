package operators

import "github.com/RossMerr/Caudex.Graph/widecolumnstore"

var _ widecolumnstore.Unary = (*Filter)(nil)

// Filter is a set operator that returns the subset of those tuples satisfying the predicate
type Filter struct {
	storage  widecolumnstore.Storage
	operator widecolumnstore.Operator
	prefix   widecolumnstore.Prefix
}

// NewFilter returns a Filter
func NewFilter(storage widecolumnstore.Storage, operator widecolumnstore.Operator, prefix widecolumnstore.Prefix) *Filter {
	return &Filter{
		prefix:   prefix,
		operator: operator,
		storage:  storage,
	}
}

func (s *Filter) Next(i widecolumnstore.Iterator) widecolumnstore.Iterator {
	unary := s.operator.(widecolumnstore.Unary)
	iterator := unary.Next(i)
	var prefixIterator widecolumnstore.Iterator
	return func() (widecolumnstore.KeyValue, bool) {
		for value, ok := prefixIterator(); ok; {
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
		return widecolumnstore.KeyValue{}, false
	}
}

func (s *Filter) Op() {}
