package operators

import (
	"github.com/voltable/graph/widecolumnstore"
)

var _ Unary = (*Filter)(nil)

// Filter is a set operator that returns the subset of those tuples satisfying the prefix
type Filter struct {
	predicate widecolumnstore.Predicate
}

// NewFilter returns a Filter
func NewFilter(predicate widecolumnstore.Predicate) Unary {
	return &Filter{
		predicate: predicate,
	}
}

func (s *Filter) Next(iterator widecolumnstore.Iterator) (widecolumnstore.Iterator, error) {
	return func() (widecolumnstore.KeyValue, bool) {
		for keyValue, ok := iterator(); ok; keyValue, ok = iterator() {
			if ok {
				if s.predicate(keyValue) {
					return keyValue, ok
				}
			}
		}
		return widecolumnstore.KeyValue{}, false
	}, nil
}

func (s *Filter) Op() {}
