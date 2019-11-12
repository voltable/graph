package operators

import (
	"github.com/voltable/graph"
	"github.com/voltable/graph/operators/ir"
	"github.com/voltable/graph/widecolumnstore"
)

var _ Unary = (*Projection)(nil)

type Projection struct {
	storage widecolumnstore.Storage
	statistics *graph.Statistics
	r  *ir.Return
	variables *ir.Variable
}

// NewProjection returns a Projection
func NewProjection(storage widecolumnstore.Storage, statistics *graph.Statistics, r *ir.Return) (*Projection, error) {
	return &Projection{
		storage: storage,
		statistics: statistics,
		r: r,
	}, nil
}

func (s *Projection) Next(iterator widecolumnstore.Iterator) widecolumnstore.Iterator {
	return func() (widecolumnstore.KeyValue, bool) {
		return widecolumnstore.KeyValue{}, false
	}
}

func (s *Projection) Op() {}