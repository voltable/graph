package operators

import "github.com/voltable/graph/widecolumnstore"

var _ widecolumnstore.Unary = (*Scan)(nil)

// Scan is a source operator that accesses a table and returns the set of all its tuples
type Scan struct {
	iterator widecolumnstore.Iterator
}

// NewScan returns a Scan
func NewScan(storage widecolumnstore.Storage) *Scan {
	return &Scan{
		iterator: storage.Each(),
	}
}

func (s *Scan) Next(i widecolumnstore.Iterator) widecolumnstore.Iterator {
	return s.iterator
}

func (s *Scan) Op() {}
