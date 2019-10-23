package operators

import (
	"github.com/voltable/graph"
	"github.com/voltable/graph/widecolumnstore"
)

var _ Nullary = (*Scan)(nil)

// Scan is a source operator that accesses a table and returns the set of all its tuples
type Scan struct {
	iterator widecolumnstore.Iterator
	prefix   []byte
}

// NewScan returns a Scan
func NewScan(storage widecolumnstore.Storage, prefix []byte) *Scan {
	scan := &Scan{}

	if prefix == nil {
		scan.iterator = storage.Each()
	} else {
		scan.iterator = storage.HasPrefix(prefix)
	}
	return scan
}

func (s *Scan) Next() (widecolumnstore.Iterator, graph.Statistics) {
	statistics := graph.NewStatistics()

	return s.iterator, statistics
}

func (s *Scan) Op() {}
