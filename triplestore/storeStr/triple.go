package storeStr

import (
	"github.com/voltable/graph/container/table"
)

// NewTriple returns a storeStr.Triple
func NewTriple(r, c, value string) *Triple {
	return &Triple{Row: r, Column: c, Value: value}
}

// NewTriplesFromTable returns a []*Triple
func NewTriplesFromTable(t table.Table) []*Triple {
	tt := make([]*Triple, 0)

	t.ReadAll()

	t.Iterator(func(r, c string, v interface{}) {
		if str, ok := v.(string); ok {
			triple := &Triple{Row: c, Column: r, Value: str}
			tt = append(tt, triple)
		}
	})

	return tt
}

// NewTripleTransposeFromTable returns a []*Triple transposed
func NewTripleTransposeFromTable(t table.Table) []*Triple {
	tt := make([]*Triple, 0)

	t.ReadAll()

	t.Iterator(func(r, c string, v interface{}) {
		if str, ok := v.(string); ok {
			triple := &Triple{Row: c, Column: r, Value: str}
			tt = append(tt, triple)
		}
	})

	return tt
}

// Transpose swap the row's and column's
func Transpose(tt []*Triple) []*Triple {
	triples := make([]*Triple, 0)

	for _, t := range tt {
		triple := &Triple{Row: t.Column, Column: t.Row, Value: t.Value}
		triples = append(triples, triple)
	}

	return triples
}
