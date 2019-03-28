package operators

import (
	"github.com/pkg/errors"

	"github.com/RossMerr/Caudex.Graph/widecolumnstore"
)

var _ widecolumnstore.Unary = (*Filter)(nil)

// Filter is a set operator that returns the subset of those tuples satisfying the predicate
type Filter struct {
	storage  widecolumnstore.HasPrefix
	operator widecolumnstore.Unary
	prefix   widecolumnstore.Prefix
}

// NewFilter returns a Filter
func NewFilter(storage widecolumnstore.HasPrefix, operator widecolumnstore.Operator, prefix widecolumnstore.Prefix) (widecolumnstore.Unary, error) {
	unary, ok := operator.(widecolumnstore.Unary)
	if !ok {
		return nil, errors.Errorf("Filter: operator not unary found %+v", operator)
	}

	return &Filter{
		prefix:   prefix,
		operator: unary,
		storage:  storage,
	}, nil
}

func (s *Filter) Next(i widecolumnstore.Iterator) widecolumnstore.Iterator {
	iterator := s.operator.Next(i)
	var prefixIterator widecolumnstore.Iterator
	return func() (widecolumnstore.KeyValue, bool) {
		if prefixIterator != nil {
			for value, ok := prefixIterator(); ok; {
				return value, ok
			}
		}
		for value, ok := iterator(); ok; value, ok = iterator() {
			key := widecolumnstore.Key{}
			key.Unmarshal(value.Key)

			prefix := s.prefix(key)

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