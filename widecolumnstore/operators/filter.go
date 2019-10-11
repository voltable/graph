package operators

import (
	"github.com/voltable/graph/widecolumnstore"
)

var _ widecolumnstore.Unary = (*Filter)(nil)

// Filter is a set operator that returns the subset of those tuples satisfying the prefix
type Filter struct {
	storage   widecolumnstore.HasPrefix
	prefix    widecolumnstore.Prefix
	predicate widecolumnstore.Predicate
}

// NewFilter returns a Filter
func NewFilter(storage widecolumnstore.HasPrefix, prefix widecolumnstore.Prefix, predicate widecolumnstore.Predicate) (widecolumnstore.Unary, error) {
	return &Filter{
		prefix:    prefix,
		storage:   storage,
		predicate: predicate,
	}, nil
}

func (s *Filter) Next(iterator widecolumnstore.Iterator) widecolumnstore.Iterator {
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
			keyValue, ok := prefixIterator()
			if ok {
				if s.predicate(keyValue) {
					return keyValue, ok
				}
			}
		}
		return widecolumnstore.KeyValue{}, false
	}
}

func (s *Filter) Op() {}
